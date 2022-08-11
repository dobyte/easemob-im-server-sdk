package push

import (
	"fmt"
	"github.com/dobyte/easemob-im-server-sdk/internal/core"
)

const (
	getTemplateUri    = "/notification/template/%s"
	createTemplateUri = "/notification/template"
	deleteTemplateUri = "/notification/template/%s"
)

type API interface {
	// GetTemplate 查询离线推送模板
	// 查询离线推送消息使用的模板。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#查询离线推送模板
	GetTemplate(name string) (*Template, error)

	// CreateTemplate 创建离线推送模板
	// 创建离线推送消息模板。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#创建离线推送模板
	CreateTemplate(name, titlePattern, contentPattern string) error

	// DeleteTemplate 删除离线推送模板
	// 删除离线消息推送模板。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#删除离线推送模板
	DeleteTemplate(name string) error
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// GetTemplate 查询离线推送模板
func (a *api) GetTemplate(name string) (*Template, error) {
	resp := &getTemplateResp{}
	if err := a.client.Get(fmt.Sprintf(getTemplateUri, name), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// CreateTemplate 创建离线推送模板
func (a *api) CreateTemplate(name, titlePattern, contentPattern string) error {
	req := &createTemplateReq{Name: name, TitlePattern: titlePattern, ContentPattern: contentPattern}
	return a.client.Post(createTemplateUri, req, nil)
}

// DeleteTemplate 删除离线推送模板
func (a *api) DeleteTemplate(name string) error {
	return a.client.Delete(fmt.Sprintf(deleteTemplateUri, name), nil, nil)
}
