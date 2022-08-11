package core

import (
	"errors"
	"github.com/dobyte/http"
	"log"
	nethttp "net/http"
	"reflect"
	"strings"
)

type Options struct {
	Host                string
	AppKey              string
	ClientID            string
	ClientSecret        string
	TTL                 int64
	unauthorizedHandler func(c *client) error
}

type Client interface {
	// Use 设置中间件
	Use(middlewares ...http.MiddlewareFunc)
	// BaseUrl 获取基础url
	BaseUrl() string
	// Get GET请求
	Get(uri string, data interface{}, resp interface{}) error
	// Post POST请求
	Post(uri string, data interface{}, resp interface{}) error
	// Put PUT请求
	Put(uri string, data interface{}, resp interface{}) error
	// Patch PATCH请求
	Patch(uri string, data interface{}, resp interface{}) error
	// Delete DELETE请求
	Delete(uri string, data interface{}, resp interface{}) error
}

type client struct {
	opts    *Options
	client  *http.Client
	baseUrl string
}

func NewClient(opts *Options) *client {
	args := strings.Split(opts.AppKey, "#")
	if len(args) != 2 {
		log.Fatal("invalid appKey")
	}

	c := new(client)
	c.opts = opts
	c.baseUrl = "https://" + opts.Host + "/" + args[0] + "/" + args[1]
	c.client = http.NewClient()
	c.client.SetContentType(http.ContentTypeJson)
	c.client.SetHeader("Accept", http.ContentTypeJson)
	c.client.SetBaseUrl(c.baseUrl)

	return c
}

// Use 设置中间件
func (c *client) Use(middlewares ...http.MiddlewareFunc) {
	c.client.Use(middlewares...)
}

// BaseUrl 获取基础url
func (c *client) BaseUrl() string {
	return c.baseUrl
}

// Get GET请求
func (c *client) Get(uri string, data interface{}, resp interface{}) error {
	return c.request(http.MethodGet, uri, data, resp)
}

// Post POST请求
func (c *client) Post(uri string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPost, uri, data, resp)
}

// Put PUT请求
func (c *client) Put(uri string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPut, uri, data, resp)
}

// Patch PATCH请求
func (c *client) Patch(uri string, data interface{}, resp interface{}) error {
	return c.request(http.MethodPatch, uri, data, resp)
}

// Delete DELETE请求
func (c *client) Delete(uri string, data interface{}, resp interface{}) error {
	return c.request(http.MethodDelete, uri, data, resp)
}

// HTTP请求
func (c *client) request(method string, uri string, data interface{}, resp interface{}) error {
	for i := 0; i < 2; i++ {
		res, err := c.client.Request(method, uri, data)
		if err != nil {
			return err
		}

		if res.Response.StatusCode == nethttp.StatusOK {
			if resp == nil || reflect.ValueOf(resp).IsNil() {
				return nil
			}

			return res.Scan(resp)
		}

		if res.Response.StatusCode == nethttp.StatusUnauthorized {
			if c.opts.unauthorizedHandler != nil && i < 1 {
				if err = c.opts.unauthorizedHandler(c); err != nil {
					return err
				}
				continue
			}
		}

		errResp := &errorResp{}
		if err = res.Scan(errResp); err != nil {
			return err
		}

		return errors.New(errResp.ErrorDescription)
	}

	return nil
}
