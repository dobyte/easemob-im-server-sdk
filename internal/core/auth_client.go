package core

import "golang.org/x/sync/singleflight"

const (
	defaultGrantType = "client_credentials"
	defaultAuthUri   = "/token"
	defaultTokenTTL  = 7200
)

type authClient struct {
	*client
	sfg singleflight.Group
}

type getTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	TTL          int64  `json:"ttl"`
}

type getTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	Application string `json:"application"`
}

func NewAuthClient(opts *Options) Client {
	ac := &authClient{}
	opts.unauthorizedHandler = func(c *client) error {
		_, err, _ := ac.sfg.Do(opts.ClientID, func() (interface{}, error) {
			req := &getTokenReq{
				GrantType:    defaultGrantType,
				ClientID:     opts.ClientID,
				ClientSecret: opts.ClientSecret,
				TTL:          opts.TTL,
			}
			resp := &getTokenResp{}

			if req.TTL <= 0 {
				req.TTL = defaultTokenTTL
			}

			if err := c.Post(defaultAuthUri, req, resp); err != nil {
				return nil, err
			}

			c.client.SetBearerToken(resp.AccessToken)

			return nil, nil
		})

		return err
	}
	ac.client = NewClient(opts)

	return ac
}
