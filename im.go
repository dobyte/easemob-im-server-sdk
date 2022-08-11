package im

import (
	"github.com/dobyte/easemob-im-server-sdk/chatroom"
	"github.com/dobyte/easemob-im-server-sdk/group"
	"github.com/dobyte/easemob-im-server-sdk/internal/core"
	"github.com/dobyte/easemob-im-server-sdk/message"
	"github.com/dobyte/easemob-im-server-sdk/push"
	"github.com/dobyte/easemob-im-server-sdk/user"
	"sync"
)

type IM interface {
	// User 获取用户管理接口
	User() user.API
	// Push 获取推送设置接口
	Push() push.API
	// Message 获取消息管理接口
	Message() message.API
	// Group 获取群组管理接口
	Group() group.API
	// Chatroom 获取聊天室管理接口
	Chatroom() chatroom.API
}

type Options struct {
	Host         string
	AppKey       string
	ClientID     string
	ClientSecret string
	TokenTTL     int64
}

type im struct {
	client     core.Client
	authClient core.Client
	user       struct {
		once     sync.Once
		instance user.API
	}
	push struct {
		once     sync.Once
		instance push.API
	}
	message struct {
		once     sync.Once
		instance message.API
	}
	group struct {
		once     sync.Once
		instance group.API
	}
	chatroom struct {
		once     sync.Once
		instance chatroom.API
	}
}

func NewIM(opts *Options) IM {
	return &im{
		client: core.NewClient(&core.Options{
			Host:   opts.Host,
			AppKey: opts.AppKey,
		}),
		authClient: core.NewAuthClient(&core.Options{
			Host:         opts.Host,
			AppKey:       opts.AppKey,
			ClientID:     opts.ClientID,
			ClientSecret: opts.ClientSecret,
			TTL:          opts.TokenTTL,
		}),
	}
}

// User 获取用户管理接口
func (i *im) User() user.API {
	i.user.once.Do(func() {
		i.user.instance = user.NewAPI(i.authClient)
	})
	return i.user.instance
}

// Push 获取推送设置接口
func (i *im) Push() push.API {
	i.push.once.Do(func() {
		i.push.instance = push.NewAPI(i.authClient)
	})
	return i.push.instance
}

// Message 获取消息管理接口
func (i *im) Message() message.API {
	i.message.once.Do(func() {
		i.message.instance = message.NewAPI(i.authClient)
	})
	return i.message.instance
}

// Group 获取群组管理接口
func (i *im) Group() group.API {
	i.group.once.Do(func() {
		i.group.instance = group.NewAPI(i.authClient)
	})
	return i.group.instance
}

// Chatroom 获取聊天室管理接口
func (i *im) Chatroom() chatroom.API {
	i.chatroom.once.Do(func() {
		i.chatroom.instance = chatroom.NewAPI(i.authClient)
	})
	return i.chatroom.instance
}
