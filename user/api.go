package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dobyte/easemob-im-server-sdk/internal/core"
	"github.com/dobyte/http"
	"net/url"
	"regexp"
	"strconv"
)

const (
	registerUsersUri                      = "/users"
	getUserUri                            = "/users/%s"
	fetchUsersUri                         = "/users"
	deleteUserUri                         = "/users/%s"
	deleteUsersUri                        = "/users?limit=%d"
	updatePasswordUri                     = "/users/%s/password"
	getOnlineStatusUri                    = "/users/%s/status"
	batchGetOnlineStatusUri               = "/users/batch/status"
	setMutesUri                           = "/mutes"
	getMutesUri                           = "/mutes/%s"
	fetchMutesUri                         = "/mutes?pageNum=%d&pageSize=%d"
	getOfflineMsgCountUri                 = "/users/%s/offline_msg_count"
	getOfflineMsgStatusUri                = "/users/%s/offline_msg_status/%s"
	deactivateUri                         = "/users/%s/deactivate"
	activateUri                           = "/users/%s/activate"
	offlineUri                            = "/users/%s/disconnect"
	addFriendUri                          = "/users/%s/contacts/users/%s"
	removeFriendUri                       = "/users/%s/contacts/users/%s"
	getFriendsUri                         = "/users/%s/contacts/users"
	addBlacklistsUri                      = "/users/%s/blocks/users"
	removeBlacklistUri                    = "/users/%s/blocks/users/%s"
	getBlacklistsUri                      = "/users/%s/blocks/users"
	setMetadataUri                        = "/metadata/user/%s"
	getMetadataUri                        = "/metadata/user/%s"
	deleteMetadataUri                     = "/metadata/user/%s"
	batchGetMetadataUri                   = "/metadata/user/get"
	getCapacityUri                        = "/metadata/user/capacity"
	setOfflinePushNicknameUri             = "/users/%s"
	setOfflinePushDisplayStyleUri         = "/users/%s"
	setOfflinePushNoDisturbingUri         = "/users/%s"
	setOfflinePushTargetedNoDisturbingUri = "/users/%s/notification/%s/%s"
	getOfflinePushTargetedNoDisturbingUri = "/users/%s/notification/%s/%s"
	setOfflinePushLanguageUri             = "/users/%s/notification/language"
	getOfflinePushLanguageUri             = "/users/%s/notification/language"
	getJoinedChatroomsUri                 = "/users/%s/joined_chatrooms"
	getJoinedGroupUri                     = "/users/%s/joined_chatgroups"
	fetchJoinedThreadsUri                 = "/threads/user/%s?limit=%d&cursor=%s&sort=%s"
)

type API interface {
	// RegisterUsers 批量注册用户
	// 批量注册是授权注册方式，服务端需要校验有效的 token 权限才能进行操作。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#批量注册用户
	RegisterUsers(user ...User) ([]*Entity, error)

	// GetUser 获取单个用户
	// 获取单个应用用户的详细信息接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#获取单个用户的详情
	GetUser(username string) (*Entity, error)

	// FetchUsers 批量获取用户详情
	// 该接口查询多个用户的信息列表，按照用户创建时间顺序返回。你可以指定要查询的用户数量（limit）。
	// 若数据库中的用户数量大于你要查询的用户数量（limit），返回的信息中会携带游标 “cursor” 标示下次数据获取的开始位置。你可以分页获取多个用户的详情，直到返回的信息中不再包含 cursor，即已经达到最后一页。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#批量获取用户详情
	FetchUsers(arg FetchUserArg) (*FetchUsersRet, error)

	// DeleteUser 删除单个用户
	// 删除一个用户。如果此用户是群主或者聊天室所有者，系统会同时删除对应的群组和聊天室。请在操作时进行确认。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#删除单个用户
	DeleteUser(username string) error

	// DeleteUsers 批量删除用户
	// 删除某个 App 下指定数量的用户账号。建议一次删除的用户数量不要超过 100。需要注意的是，这里只指定了要删除的用户数量，并未指定要删除的具体用户，你可以在响应中查看删除的用户。
	// 如果删除的多个用户中包含群组或者聊天室的管理员，该用户管理的群组和聊天室也会相应被删除。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#批量删除用户
	DeleteUsers(limit int) ([]*Entity, error)

	// DeleteAllUsers 删除所有用户
	// 本方法为“批量删除用户（BatchDelete）”的拓展方法。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#批量删除用户
	DeleteAllUsers() ([]*Entity, error)

	// UpdatePassword 修改用户密码
	// 可以通过服务端接口修改用户的登录密码，不需要提供原密码。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#修改用户密码
	UpdatePassword(username, password string) error

	// GetOnlineStatus 获取单个用户在线状态
	// 查看一个用户的在线状态。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#获取单个用户在线状态
	GetOnlineStatus(username string) (string, error)

	// GetOnlineStatuses 批量获取用户在线状态
	// 批量查看用户的在线状态，最多可同时查看 100 个用户的状态。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#批量获取用户在线状态
	GetOnlineStatuses(usernames ...string) (map[string]string, error)

	// SetMutes 设置用户全局禁言
	// 设置单个用户 ID 的单聊、群组、聊天室消息全局禁言。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#设置用户全局禁言
	SetMutes(mutes Mutes) error

	// GetMutes 查询单个用户全局禁言
	// 查询单个用户的单聊、群聊和聊天室消息的禁言。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#查询单个用户_id_全局禁言
	GetMutes(username string) (*MutesRet, error)

	// FetchMutes 查询app下的所有全局禁言的用户
	// 该方法查询 app 下所有全局禁言的用户及其禁言剩余时间。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#查询_app_下的所有全局禁言的用户
	FetchMutes(arg FetchMutesArg) (*FetchMutesRet, error)

	// GetOfflineMsgCount 获取用户离线消息数量
	// 获取环信 IM 用户的离线消息数量。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#获取用户离线消息数据
	GetOfflineMsgCount(username string) (int, error)

	// GetOfflineMsgStatus 获取某条离线消息状态
	// 获取用户的离线消息的状态，即查看该消息是否已投递。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#获取某条离线消息状态
	GetOfflineMsgStatus(username, msgID string) (string, error)

	// DeactivateUser 账号封禁
	// 环信即时通讯 IM 提供了对用户的禁用以及解禁接口操作，用户若被禁用将立即下线并无法登录进入环信即时通讯 IM，直到被解禁后才能恢复登录。常用在对异常用户的即时处理场景使用。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#账号封禁
	DeactivateUser(username string) (*Entity, error)

	// ActivateUser 账号解禁
	// 环信即时通讯 IM 提供了对用户的禁用以及解禁接口操作。对用户禁用后，用户将立即下线并无法登录进入环信即时通讯 IM，直到被解禁后才能恢复登录。该功能常于对异常用户的即时处理场景。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#账号解禁
	ActivateUser(username string) error

	// OfflineUser 强制下线
	// 强制用户即把用户状态改为离线，用户需要重新登录才能正常使用。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/accountsystem#强制下线
	OfflineUser(username string) (bool, error)

	// AddFriend 添加好友
	// 添加好友，好友必须是和自己在一个 App Key 下的用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/relationship#添加好友
	AddFriend(ownerUsername, friendUsername string) error

	// RemoveFriend 移除好友
	// 从用户的好友列表中移除一个用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/relationship#移除好友
	RemoveFriend(ownerUsername, friendUsername string) error

	// GetFriends 获取好友列表
	// 获取用户的好友列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/relationship#移除好友
	GetFriends(username string) ([]string, error)

	// AddBlacklists 添加黑名单
	// 向用户的黑名单列表中添加一个或者多个用户，黑名单中的用户无法给该用户发送消息，每个用户的黑名单人数上限为 500。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/relationship#添加黑名单
	AddBlacklists(ownerUsername string, otherUsernames ...string) error

	// RemoveBlacklist 移除黑名单
	// 从用户的黑名单中移除用户。将用户从黑名单移除后，恢复到好友，或者未添加好友的用户关系。可以正常的进行消息收发。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/relationship#移除黑名单
	RemoveBlacklist(ownerUsername, blackedUsername string) error

	// GetBlacklists 获取黑名单
	// 获取黑名单列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/relationship#获取黑名单
	GetBlacklists(ownerUsername string) ([]string, error)

	// SetMetadata 设置用户属性
	// 用户属性的内容为一个或多个纯文本键值对，默认单一用户的属性总长不得超过 2 KB，默认一个 app 下所有用户的所有属性总长不得超过 10 GB。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/userprofile#设置用户属性
	SetMetadata(username string, metadata map[string]string) error

	// GetMetadata 获取用户属
	// 获取指定用户的所有用户属性键值对。需要在请求时对应填写 {username}，需要获取用户属性的用户 ID。
	// 如果指定的用户或用户属性不存在，返回空数据 {}。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/userprofile#获取用户属
	GetMetadata(username string) (map[string]string, error)

	// BatchGetMetadata 批量获取用户属性
	// 根据指定的用户名列表和属性列表，查询用户属性。
	// 如果指定的用户或用户属性不存在，返回空数据 {}。 每次最多指定 100 个用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/userprofile#批量获取用户属性
	BatchGetMetadata(properties []string, usernames ...string) (map[string]map[string]string, error)

	// DeleteMetadata 删除用户属性
	// 删除指定用户的所有属性。如果指定的用户或用户属性不存在（可能已删除），也视为删除成功。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/userprofile#批量获取用户属性
	DeleteMetadata(username string) (bool, error)

	// GetCapacity 获取用户属性总量大小
	// 获取该 app 下所有用户的属性数据大小，单位为字节。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/userprofile#获取用户属性总量大小
	GetCapacity() (int64, error)

	// SetOfflinePushNickname 设置离线推送时显示的昵称
	// 设置离线推送时显示的昵称。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#设置离线推送时显示的昵称
	SetOfflinePushNickname(username, nickname string) error

	// SetOfflinePushDisplayStyle 设置离线推送通知的展示方式
	// 设置离线推送通知在客户端的展示方式，设置即时生效。服务端据此向用户推送离线消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#设置离线推送通知的展示方式
	SetOfflinePushDisplayStyle(username string, displayStyle int) error

	// EnableOfflinePushNoDisturbing 启用免打扰模式
	// 设置离线推送免打扰模式，在免打扰期间，用户将不会收到离线消息推送。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#设置免打扰模式
	EnableOfflinePushNoDisturbing(username string, start, end int) error

	// DisableOfflinePushNoDisturbing 禁用免打扰模式
	// 设置离线推送免打扰模式，在免打扰期间，用户将不会收到离线消息推送。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#设置免打扰模式
	DisableOfflinePushNoDisturbing(username string) error

	// SetOfflinePushTargetedNoDisturbing 设置离线推送设置
	// 你可以设置全局离线推送的通知方式和免打扰模式以及单个单聊或群聊会话的离线推送设置。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#设置离线推送设置
	SetOfflinePushTargetedNoDisturbing(arg *SetOfflinePushTargetedNoDisturbingArg) error

	// GetOfflinePushTargetedNoDisturbing 查询离线推送设置
	// 查询指定单聊、指定群聊或全局的离线推送设置。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#查询离线推送设置
	GetOfflinePushTargetedNoDisturbing(username string, toType, toKey string) (*NoDisturbing, error)

	// SetOfflinePushLanguage 设置推送翻译语言
	// 设置离线推送消息的翻译语言。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#设置推送翻译语言
	SetOfflinePushLanguage(username string, language string) error

	// GetOfflinePushLanguage 获取推送翻译语言
	// 查询离线推送消息的翻译语言。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/pushconfig#查询推送翻译
	GetOfflinePushLanguage(username string) (string, error)

	// GetJoinedChatrooms 获取用户加入的聊天室
	// 根据用户 ID 获取该用户加入的全部聊天室。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#获取用户加入的聊天室
	GetJoinedChatrooms(username string) ([]*JoinedChatroom, error)

	// GetJoinedGroups 获取单个用户加入的所有群组
	// 根据用户 ID 称获取该用户加入的全部群组。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取单个用户加入的所有群组_可分页
	GetJoinedGroups(username string) ([]*JoinedGroup, error)

	// FetchJoinedThreads 获取一个用户加入的所有子区
	// 根据用户 ID 获取该用户加入的子区列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取一个用户加入的所有子区_分页获取
	FetchJoinedThreads(arg FetchJoinedThreadsArg) (*FetchJoinedThreadsRet, error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// RegisterUsers 注册用户
func (a *api) RegisterUsers(users ...User) ([]*Entity, error) {
	if len(users) > 60 {
		return nil, errors.New("the number of registered users exceeds the upper limit")
	}

	resp := &registerUsersResp{}

	if err := a.client.Post(registerUsersUri, users, resp); err != nil {
		return nil, err
	}

	return resp.Entities, nil
}

// GetUser 获取单个用户
func (a *api) GetUser(username string) (*Entity, error) {
	resp := &getResp{}
	if err := a.client.Get(fmt.Sprintf(getUserUri, username), nil, resp); err != nil {
		return nil, err
	}

	return a.toEntity(resp.Entities[0])
}

// FetchUsers 批量获取用户详情
func (a *api) FetchUsers(arg FetchUserArg) (*FetchUsersRet, error) {
	resp := &fetchUsersResp{}
	if err := a.client.Get(fetchUsersUri, arg, resp); err != nil {
		return nil, err
	}

	ret := &FetchUsersRet{
		List:    make([]*Entity, 0, len(resp.Entities)),
		Cursor:  resp.Cursor,
		HasMore: resp.Cursor != "",
	}
	for _, item := range resp.Entities {
		entity, err := a.toEntity(item)
		if err != nil {
			return nil, err
		}
		ret.List = append(ret.List, entity)
	}

	return ret, nil
}

// DeleteUser 删除单个用户
func (a *api) DeleteUser(username string) error {
	return a.client.Delete(fmt.Sprintf(deleteUserUri, username), nil, nil)
}

// DeleteUsers 批量删除用户
func (a *api) DeleteUsers(limit int) ([]*Entity, error) {
	resp := &deleteUsersResp{}
	if err := a.client.Delete(fmt.Sprintf(deleteUsersUri, limit), nil, resp); err != nil {
		return nil, err
	}

	entities := make([]*Entity, 0, len(resp.Entities))
	for _, item := range resp.Entities {
		entity, err := a.toEntity(item)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// DeleteAllUsers 删除所有用户
func (a *api) DeleteAllUsers() ([]*Entity, error) {
	return a.DeleteUsers(0)
}

// UpdatePassword 修改用户密码
func (a *api) UpdatePassword(username, password string) error {
	req := &updatePasswordReq{NewPassword: password}
	return a.client.Put(fmt.Sprintf(updatePasswordUri, username), req, nil)
}

// GetOnlineStatus 获取单个用户在线状态
func (a *api) GetOnlineStatus(username string) (string, error) {
	resp := &getOnlineStatusResp{}
	if err := a.client.Get(fmt.Sprintf(getOnlineStatusUri, username), nil, resp); err != nil {
		return "", err
	}

	return resp.Data[username], nil
}

// GetOnlineStatuses 批量获取用户在线状态
func (a *api) GetOnlineStatuses(usernames ...string) (map[string]string, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 100:
		return nil, errors.New("the number of batch get users exceeds the upper limit")
	}

	req := &batchGetOnlineStatusReq{Usernames: usernames}
	resp := &batchGetOnlineStatusResp{}
	if err := a.client.Post(batchGetOnlineStatusUri, req, resp); err != nil {
		return nil, err
	}

	ret := make(map[string]string, len(resp.Data))
	for _, item := range resp.Data {
		for k, v := range item {
			ret[k] = v
		}
	}

	return ret, nil
}

// SetMutes 设置用户全局禁言
func (a *api) SetMutes(mutes Mutes) error {
	return a.client.Post(setMutesUri, mutes, nil)
}

// GetMutes 查询单个用户全局禁言
func (a *api) GetMutes(username string) (*MutesRet, error) {
	resp := &getMutesResp{}
	if err := a.client.Get(fmt.Sprintf(getMutesUri, username), nil, resp); err != nil {
		return nil, err
	}

	return &MutesRet{
		Username:  resp.Data.Username,
		Chat:      resp.Data.Chat,
		Groupchat: resp.Data.Groupchat,
		Chatroom:  resp.Data.Chatroom,
		Unixtime:  resp.Data.Unixtime,
	}, nil
}

// FetchMutes 查询app下的所有全局禁言的用户
func (a *api) FetchMutes(arg FetchMutesArg) (*FetchMutesRet, error) {
	resp := &fetchMutesResp{}
	if err := a.client.Get(fmt.Sprintf(fetchMutesUri, arg.PageNum, arg.PageSize), nil, resp); err != nil {
		return nil, err
	}

	return &FetchMutesRet{
		List:     resp.Data.Data,
		HasMore:  len(resp.Data.Data) >= arg.PageSize,
		Unixtime: resp.Data.Unixtime,
	}, nil
}

// GetOfflineMsgCount 获取用户的离线消息数量。
func (a *api) GetOfflineMsgCount(username string) (int, error) {
	resp := &getOfflineMsgCountResp{}
	if err := a.client.Get(fmt.Sprintf(getOfflineMsgCountUri, username), nil, resp); err != nil {
		return 0, err
	}

	return resp.Data[username], nil
}

// GetOfflineMsgStatus 获取某条离线消息状态
func (a *api) GetOfflineMsgStatus(username, msgID string) (string, error) {
	resp := &getOfflineMsgStatusResp{}
	if err := a.client.Get(fmt.Sprintf(getOfflineMsgStatusUri, username, msgID), nil, resp); err != nil {
		return "", err
	}

	return resp.Data[msgID], nil
}

// DeactivateUser 账号封禁
func (a *api) DeactivateUser(username string) (*Entity, error) {
	resp := &deactivateResp{}
	if err := a.client.Post(fmt.Sprintf(deactivateUri, username), nil, resp); err != nil {
		return nil, err
	}

	return a.toEntity(resp.Entities[0])
}

// ActivateUser 账号解禁
func (a *api) ActivateUser(username string) error {
	return a.client.Post(fmt.Sprintf(activateUri, username), nil, nil)
}

// OfflineUser 强制下线
func (a *api) OfflineUser(username string) (bool, error) {
	resp := &offlineResp{}
	if err := a.client.Get(fmt.Sprintf(offlineUri, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Result, nil
}

func (a *api) toEntity(data map[string]interface{}) (*Entity, error) {
	entity := &Entity{}

	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(buf, entity); err != nil {
		return nil, err
	}

	reg := regexp.MustCompile(`^notification_ignore_(\d+)$`)
	for k := range data {
		result := reg.FindStringSubmatch(k)
		if len(result) != 2 {
			continue
		}
		entity.NotificationIgnoreGroups = append(entity.NotificationIgnoreGroups, result[1])
	}

	return entity, nil
}

// AddFriend 添加好友
func (a *api) AddFriend(ownerUsername, friendUsername string) error {
	return a.client.Post(fmt.Sprintf(addFriendUri, ownerUsername, friendUsername), nil, nil)
}

// RemoveFriend 移除好友
func (a *api) RemoveFriend(ownerUsername, friendUsername string) error {
	return a.client.Delete(fmt.Sprintf(removeFriendUri, ownerUsername, friendUsername), nil, nil)
}

// GetFriends 获取好友列表
func (a *api) GetFriends(ownerUsername string) ([]string, error) {
	resp := &getFriendsResp{}
	if err := a.client.Get(fmt.Sprintf(getFriendsUri, ownerUsername), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddBlacklists 添加黑名单
func (a *api) AddBlacklists(ownerUsername string, otherUsernames ...string) error {
	switch count := len(otherUsernames); {
	case count == 0:
		return nil
	case count > 500:
		return errors.New("the number of user exceeds the upper limit")
	}

	req := &addBlacklistsReq{Usernames: otherUsernames}
	return a.client.Post(fmt.Sprintf(addBlacklistsUri, ownerUsername), req, nil)
}

// RemoveBlacklist 移除黑名单
func (a *api) RemoveBlacklist(ownerUsername, blackedUsername string) error {
	return a.client.Delete(fmt.Sprintf(removeBlacklistUri, ownerUsername, blackedUsername), nil, nil)
}

// GetBlacklists 获取黑名单
func (a *api) GetBlacklists(ownerUsername string) ([]string, error) {
	resp := &getBlacklistsResp{}
	if err := a.client.Get(fmt.Sprintf(getBlacklistsUri, ownerUsername), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// SetMetadata 设置用户属性
func (a *api) SetMetadata(username string, metadata map[string]string) error {
	query := new(url.URL).Query()
	for k, v := range metadata {
		query.Add(k, v)
	}

	a.client.Use(func(r *http.Request) (*http.Response, error) {
		r.Request.Header.Set(http.HeaderContentType, http.ContentTypeFormUrlEncoded)
		return r.Next()
	})

	return a.client.Put(fmt.Sprintf(setMetadataUri, username), query.Encode(), nil)
}

// GetMetadata 获取用户属
func (a *api) GetMetadata(username string) (map[string]string, error) {
	resp := &getMetadataResp{}
	if err := a.client.Get(fmt.Sprintf(getMetadataUri, username), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// BatchGetMetadata 批量获取用户属性
func (a *api) BatchGetMetadata(properties []string, usernames ...string) (map[string]map[string]string, error) {
	if len(properties) == 0 {
		return nil, errors.New("property to get is not set")
	}

	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 100:
		return nil, errors.New("the number of batch get users exceeds the upper limit")
	}

	req := &batchGetMetadataReq{Properties: properties, Targets: usernames}
	resp := &batchGetMetadataResp{}
	if err := a.client.Post(batchGetMetadataUri, req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// DeleteMetadata 删除用户属性
func (a *api) DeleteMetadata(username string) (bool, error) {
	resp := &deleteMetadataResp{}
	if err := a.client.Delete(fmt.Sprintf(deleteMetadataUri, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data, nil
}

// GetCapacity 获取用户属性总量大小
func (a *api) GetCapacity() (int64, error) {
	resp := &getCapacityResp{}
	if err := a.client.Get(getCapacityUri, nil, resp); err != nil {
		return 0, err
	}

	return resp.Data, nil
}

// SetOfflinePushNickname 设置离线推送时显示的昵称
func (a *api) SetOfflinePushNickname(username, nickname string) error {
	req := &setOfflinePushNicknameReq{Nickname: nickname}
	return a.client.Put(fmt.Sprintf(setOfflinePushNicknameUri, username), req, nil)
}

// SetOfflinePushDisplayStyle 设置离线推送通知的展示方式
func (a *api) SetOfflinePushDisplayStyle(username string, displayStyle int) error {
	req := &setOfflinePushDisplayStyleReq{NotificationDisplayStyle: displayStyle}
	return a.client.Put(fmt.Sprintf(setOfflinePushDisplayStyleUri, username), req, nil)
}

// 设置免打扰模式
func (a *api) setOfflinePushNoDisturbing(username string, enable bool, start, end string) error {
	req := &setOfflinePushNoDisturbingReq{
		NotificationNoDisturbing:      enable,
		NotificationNoDisturbingStart: start,
		NotificationNoDisturbingEnd:   end,
	}
	return a.client.Put(fmt.Sprintf(setOfflinePushNoDisturbingUri, username), req, nil)
}

// EnableOfflinePushNoDisturbing 启用免打扰模式
func (a *api) EnableOfflinePushNoDisturbing(username string, start, end int) error {
	return a.setOfflinePushNoDisturbing(username, true, strconv.Itoa(start), strconv.Itoa(end))
}

// DisableOfflinePushNoDisturbing 禁用免打扰模式
func (a *api) DisableOfflinePushNoDisturbing(username string) error {
	return a.setOfflinePushNoDisturbing(username, false, "", "")
}

// SetOfflinePushTargetedNoDisturbing 设置离线推送设置
func (a *api) SetOfflinePushTargetedNoDisturbing(arg *SetOfflinePushTargetedNoDisturbingArg) error {
	uri := fmt.Sprintf(setOfflinePushTargetedNoDisturbingUri, arg.Username, arg.ToType, arg.ToKey)
	req := &setOfflinePushTargetedNoDisturbingReq{
		Type:           arg.Type,
		IgnoreInterval: arg.IgnoreInterval,
		IgnoreDuration: arg.IgnoreDuration,
	}
	return a.client.Put(uri, req, nil)
}

// GetOfflinePushTargetedNoDisturbing 查询离线推送设置
func (a *api) GetOfflinePushTargetedNoDisturbing(username string, toType, toKey string) (*NoDisturbing, error) {
	uri := fmt.Sprintf(getOfflinePushTargetedNoDisturbingUri, username, toType, toKey)
	resp := &getOfflinePushTargetedNoDisturbingResp{}
	if err := a.client.Get(uri, nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// SetOfflinePushLanguage 设置推送翻译语言
func (a *api) SetOfflinePushLanguage(username string, language string) error {
	req := &setOfflinePushLanguageReq{TranslationLanguage: language}
	return a.client.Put(fmt.Sprintf(setOfflinePushLanguageUri, username), req, nil)
}

// GetOfflinePushLanguage 获取推送翻译语言
func (a *api) GetOfflinePushLanguage(username string) (string, error) {
	resp := &getOfflinePushLanguageResp{}
	if err := a.client.Get(fmt.Sprintf(getOfflinePushLanguageUri, username), nil, resp); err != nil {
		return "", err
	}

	return resp.Data.Language, nil
}

// GetJoinedChatrooms 获取用户加入的聊天室
func (a *api) GetJoinedChatrooms(username string) ([]*JoinedChatroom, error) {
	resp := &getJoinedChatroomsResp{}
	if err := a.client.Get(fmt.Sprintf(getJoinedChatroomsUri, username), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// GetJoinedGroups 获取单个用户加入的所有群组
func (a *api) GetJoinedGroups(username string) ([]*JoinedGroup, error) {
	resp := &getJoinedGroupsResp{}
	if err := a.client.Get(fmt.Sprintf(getJoinedGroupUri, username), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// FetchJoinedThreads 获取一个用户加入的所有子区
func (a *api) FetchJoinedThreads(arg FetchJoinedThreadsArg) (*FetchJoinedThreadsRet, error) {
	uri := fmt.Sprintf(fetchJoinedThreadsUri, arg.Username, arg.Limit, arg.Cursor, arg.Sort)
	resp := &fetchJoinedThreadsResp{}
	if err := a.client.Get(uri, nil, resp); err != nil {
		return nil, err
	}

	return &FetchJoinedThreadsRet{
		List:    resp.Entities,
		HasMore: resp.Properties.Cursor != "",
		Cursor:  resp.Properties.Cursor,
	}, nil
}
