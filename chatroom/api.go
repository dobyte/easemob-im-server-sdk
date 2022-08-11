package chatroom

import (
	"errors"
	"fmt"
	"github.com/dobyte/easemob-im-server-sdk/internal/core"
	"strings"
)

const (
	addSuperAdminUri      = "/chatrooms/super_admin"
	revokeSuperAdminUri   = "/chatrooms/super_admin/%s"
	fetchSuperAdminsUri   = "/chatrooms/super_admin?pagenum=%d&pagesize=%d"
	getAllChatroomsUri    = "/chatrooms"
	getChatroomsUri       = "/chatrooms/%s"
	createChatroomUri     = "/chatrooms"
	updateChatroomUri     = "/chatrooms/%s"
	deleteChatroomUri     = "/chatrooms/%s"
	getAnnouncementUri    = "/chatrooms/%s/announcement"
	updateAnnouncementUri = "/chatrooms/%s/announcement"
	fetchMembersUri       = "/chatrooms/%s/users?pagenum=%d&pagesize=%d"
	addMemberUri          = "/chatrooms/%s/users/%s"
	addMembersUri         = "/chatrooms/%s/users"
	removeMembersUri      = "/chatrooms/%s/users/%s"
	getAdminsUri          = "/chatrooms/%s/admin"
	addAdminUri           = "/chatrooms/%s/admin"
	removeAdminUri        = "/chatrooms/%s/admin/%s"
	getBlacklistsUri      = "/chatrooms/%s/blocks/users"
	addBlacklistUri       = "/chatrooms/%s/blocks/users/%s"
	addBlacklistsUri      = "/chatrooms/%s/blocks/users"
	removeBlacklistUri    = "/chatrooms/%s/blocks/users/%s"
	removeBlacklistsUri   = "/chatrooms/%s/blocks/users/%s"
	getWhitelistsUri      = "/chatrooms/%s/white/users"
	addWhitelistUri       = "/chatrooms/%s/white/users/%s"
	addWhitelistsUri      = "/chatrooms/%s/white/users"
	removeWhitelistsUri   = "/chatrooms/%s/white/users/%s"
	getMutesUri           = "/chatrooms/%s/mute"
	addMutesUri           = "/chatrooms/%s/mute"
	addAllMutesUri        = "/chatrooms/%s/ban"
	removeMutesUri        = "/chatrooms/%s/mute/%s"
	removeAllMutesUri     = "/chatrooms/%s/ban"
)

type API interface {
	// AddSuperAdmin 添加超级管理员
	// 在即时通讯应用中，仅聊天室超级管理员具有在客户端创建聊天室的权限。
	// 环信即时通讯 IM 提供多个管理超级管理员的接口，包括获取、添加、移除等。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#添加超级管理员
	AddSuperAdmin(username string) (bool, error)

	// RevokeSuperAdmin 撤销超级管理员
	// 撤销超级管理员权限，用户将不能再创建聊天室。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#撤销超级管理员
	RevokeSuperAdmin(username string) error

	// FetchSuperAdmins 分页获取超级管理员列表
	// 可以分页获取超级管理员列表的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#分页获取超级管理员列表
	FetchSuperAdmins(arg FetchSuperAdminsArg) (*FetchSuperAdminsRet, error)

	// GetAllChatrooms 获取app中所有的聊天室
	// 获取应用下全部的聊天室列表和信息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#获取_app_中所有的聊天室
	GetAllChatrooms() ([]*ListedChatroom, error)

	// GetChatrooms 查询聊天室详情
	// 查询一个或多个聊天室的详情。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#查询聊天室详情
	GetChatrooms(id ...string) ([]*Chatroom, error)

	// CreateChatroom 创建聊天室
	// 创建一个聊天室，并设置聊天室名称、聊天室描述、公开聊天室/私有聊天室属性、聊天室成员最大人数（包括管理员）、加入公开聊天室是否需要批准、管理员、以及聊天室成员。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#创建聊天室
	CreateChatroom(arg *CreateChatRoomArg) (string, error)

	// UpdateChatroom 修改聊天室
	// 修改指定聊天室信息。仅支持修改 name、description 和 maxusers。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#修改聊天室
	UpdateChatroom(arg UpdateChatroomArg) (*UpdateChatroomRet, error)

	// DeleteChatroom 删除聊天室
	// 删除单个聊天室。如果被删除的聊天室不存在，会返回错误。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#删除聊天室
	DeleteChatroom(id string) (bool, error)

	// GetAnnouncement 获取聊天室公告
	// 获取指定聊天室 ID 的聊天室公告。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#获取聊天室公告
	GetAnnouncement(id string) (string, error)

	// UpdateAnnouncement 修改聊天室公告
	// 修改指定聊天室 ID 的聊天室公告。聊天室公告内容不能超过 512 个字符。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#修改聊天室公告
	UpdateAnnouncement(id, announcement string) error

	// FetchMembers 分页获取聊天室成员
	// 可以分页获取聊天室成员列表的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#分页获取聊天室成员
	FetchMembers(arg FetchMembersArg) (*FetchMembersRet, error)

	// AddMember 添加单个聊天室成员
	// 向聊天室添加一个成员。如果待添加的用户在 app 中不存在或已经在聊天室中，则请求失败并返回错误码 400。
	// 一个聊天室ID多次添加同一个用户，均添加成功。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#添加单个聊天室成员
	AddMember(id, username string) (bool, error)

	// AddMembers 批量添加聊天室成员
	// 向聊天室添加多位用户，一次性最多可添加 60 位用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#批量添加聊天室成员
	AddMembers(id string, usernames ...string) ([]string, error)

	// RemoveMember 移除单个聊天室成员
	// 从聊天室移除一个成员。如果被移除用户不在聊天室中，或者聊天室不存在，将返回错误。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#移除单个聊天室成员
	RemoveMember(id, username string) (bool, error)

	// RemoveMembers 批量移除聊天室成员
	// 从聊天室移除多个成员，单次请求最多移除 100 个成员。如果被移除用户不在聊天室中，或者聊天室不存在，将返回错误。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#批量移除聊天室成员
	RemoveMembers(id string, usernames ...string) ([]*ActionResult, error)

	// GetAdmins 获取聊天室管理员列表
	// 获取聊天室管理员列表的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#获取聊天室管理员列表
	GetAdmins(id string) ([]string, error)

	// AddAdmin 添加聊天室管理员
	// 将一个聊天室成员角色设置为聊天室管理员。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#添加聊天室管理员
	AddAdmin(id, username string) (bool, error)

	// RemoveAdmin 移除聊天室管理员
	// 将用户的角色从聊天室管理员降为普通聊天室成员。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#移除聊天室管理员
	RemoveAdmin(id, username string) (bool, error)

	// GetBlacklists 查询聊天室黑名单
	// 查询一个聊天室黑名单中的用户列表。黑名单中的用户无法查看或收到该聊天室的信息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#查询聊天室黑名单
	GetBlacklists(id string) ([]string, error)

	// AddBlacklist 添加单个用户至聊天室黑名单
	// 添加一个用户进入一个聊天室的黑名单。聊天室所有者无法被加入聊天室的黑名单。
	// 用户进入聊天室黑名单后，会收到消息：“You are kicked out of the chatroom xxx”。之后，该用户无法查看和收发该聊天室的信息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#添加单个用户至聊天室黑名单
	AddBlacklist(id, username string) (bool, error)

	// AddBlacklists 批量添加用户至聊天室黑名单
	// 将多个用户加入一个聊天室的黑名单。你一次最多可以添加 60 个用户至聊天室黑名单。聊天室所有者无法被加入聊天室的黑名单。
	// 用户进入聊天室黑名单后，会收到消息：“You are kicked out of the chatroom xxx”。之后，这些用户无法查看和收发该聊天室的信息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#批量添加用户至聊天室黑名单
	AddBlacklists(id string, usernames ...string) ([]*ActionResult, error)

	// RemoveBlacklist 从聊天室黑名单移除单个用户
	// 将指定用户移出聊天室黑名单。对于聊天室黑名单中的用户，如果需要将其再次加入聊天室，需要先将其从聊天室黑名单中移除。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#从聊天室黑名单移除单个用户
	RemoveBlacklist(id, username string) (bool, error)

	// RemoveBlacklists 批量添加用户至聊天室黑名单
	// 将多名指定用户从聊天室黑名单中移除。你每次最多可移除 60 个用户。对于聊天室黑名单中的用户，如果需要将其再次加入聊天室，需要先将其从聊天室黑名单中移除。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#批量添加用户至聊天室黑名单
	RemoveBlacklists(id string, usernames ...string) ([]*ActionResult, error)

	// GetWhitelists 查询聊天室白名单
	// 查询一个聊天室白名单中的用户列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#查询聊天室白名单
	GetWhitelists(id string) ([]string, error)

	// AddWhitelist 添加单个用户至聊天室黑名单
	// 将指定的单个用户添加至聊天室白名单。用户添加至聊天室白名单后，当聊天室全员禁言时，仍可以在聊天室中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#添加单个用户至聊天室黑名单
	AddWhitelist(id, username string) (bool, error)

	// AddWhitelists 批量添加用户至聊天室白名单
	// 添加多个用户至聊天室白名单。你一次最多可添加 60 个用户。用户添加至聊天室白名单后，在聊天室全员禁言时，仍可以在聊天室中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#批量添加用户至聊天室白名单
	AddWhitelists(id string, usernames ...string) ([]*ActionResult, error)

	// RemoveWhitelist 从聊天室白名单移除单个用户
	// 本方法为“将用户批量移除聊天室白名单（RemoveWhitelists）”的拓展方法。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#将用户移除聊天室白名单
	RemoveWhitelist(id, username string) (bool, error)

	// RemoveWhitelists 将用户批量移除聊天室白名单
	// 将指定用户从聊天室白名单移除。你每次最多可移除 60 个用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#将用户移除聊天室白名单
	RemoveWhitelists(id string, usernames ...string) ([]*ActionResult, error)

	// GetMutes 获取禁言列表
	// 获取当前聊天室的禁言用户列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#获取禁言列表
	GetMutes(id string) ([]*Mute, error)

	// AddMute 禁言单个聊天室成员
	// 本方法为“禁言聊天室成员（AddMutes）”的拓展方法。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#禁言聊天室成员
	AddMute(id string, duration int64, username string) (bool, error)

	// AddMutes 禁言聊天室成员
	// 将一个用户禁言。用户被禁言后，将无法在聊天室中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#禁言聊天室成员
	AddMutes(id string, duration int64, usernames ...string) ([]*AddMuteResult, error)

	// RemoveMute 解除单个聊天室禁言成员
	// 本方法为“解除聊天室禁言成员（RemoveMutes）”的拓展方法。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#解除聊天室禁言成员
	RemoveMute(id string, username string) (bool, error)

	// RemoveMutes 解除聊天室禁言成员
	// 解除一个或多个聊天室成员的禁言。解除禁言后，该成员可以正常在聊天室中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#解除聊天室禁言成员
	RemoveMutes(id string, usernames ...string) ([]*RemoveMuteResult, error)

	// AddAllMutes 禁言聊天室全体成员
	// 对所有聊天室成员一键禁言，即将聊天室的所有成员均加入禁言列表。设置聊天室全员禁言后，仅聊天室白名单中的用户可在聊天室内发消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#禁言聊天室成员
	AddAllMutes(id string) error

	// RemoveAllMutes 解除聊天室全员禁言
	// 一键取消对聊天室全体成员的禁言。解除禁言后，聊天室成员可以在聊天室中正常发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/chatroom#解除聊天室全员禁言
	RemoveAllMutes(id string) error
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// AddSuperAdmin 添加超级管理员
func (a *api) AddSuperAdmin(username string) (bool, error) {
	req := &addSuperAdminReq{SuperAdmin: username}
	resp := &addSuperAdminResp{}

	if err := a.client.Post(addSuperAdminUri, req, resp); err != nil {
		return false, err
	}

	return resp.Data.Result == "success", nil
}

// RevokeSuperAdmin 撤销超级管理员
func (a *api) RevokeSuperAdmin(username string) error {
	return a.client.Delete(fmt.Sprintf(revokeSuperAdminUri, username), nil, nil)
}

// FetchSuperAdmins 分页获取超级管理员列表
func (a *api) FetchSuperAdmins(arg FetchSuperAdminsArg) (*FetchSuperAdminsRet, error) {
	resp := &fetchSuperAdminsResp{}
	if err := a.client.Get(fmt.Sprintf(fetchSuperAdminsUri, arg.PageNum, arg.PageSize), nil, resp); err != nil {
		return nil, err
	}

	return &FetchSuperAdminsRet{
		List:    resp.Data,
		HasMore: resp.Count >= arg.PageSize,
	}, nil
}

// GetAllChatrooms 获取app中所有的聊天室
func (a *api) GetAllChatrooms() ([]*ListedChatroom, error) {
	resp := &getAllChatroomsResp{}
	if err := a.client.Get(getAllChatroomsUri, nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// GetChatrooms 查询聊天室详情
func (a *api) GetChatrooms(id ...string) ([]*Chatroom, error) {
	switch count := len(id); {
	case count == 0:
		return nil, nil
	case count > 100:
		return nil, errors.New("the number of get chatrooms exceeds the upper limit")
	}

	resp := &getChatroomsResp{}
	if err := a.client.Get(fmt.Sprintf(getChatroomsUri, strings.Join(id, ",")), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// CreateChatroom 创建聊天室
func (a *api) CreateChatroom(arg *CreateChatRoomArg) (string, error) {
	resp := &createChatRoomResp{}
	if err := a.client.Post(createChatroomUri, arg, resp); err != nil {
		return "", err
	}

	return resp.Data.ID, nil
}

// UpdateChatroom 修改聊天室
func (a *api) UpdateChatroom(arg UpdateChatroomArg) (*UpdateChatroomRet, error) {
	resp := &UpdateChatroomRet{}
	if err := a.client.Put(fmt.Sprintf(updateChatroomUri, arg.ID), arg, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteChatroom 删除聊天室
func (a *api) DeleteChatroom(id string) (bool, error) {
	resp := &deleteChatroomResp{}
	if err := a.client.Delete(fmt.Sprintf(deleteChatroomUri, id), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Success, nil
}

// GetAnnouncement 获取聊天室公告
func (a *api) GetAnnouncement(id string) (string, error) {
	resp := &getAnnouncementResp{}
	if err := a.client.Get(fmt.Sprintf(getAnnouncementUri, id), nil, resp); err != nil {
		return "", err
	}

	return resp.Data.Announcement, nil
}

// UpdateAnnouncement 修改聊天室公告
func (a *api) UpdateAnnouncement(id, announcement string) error {
	req := &updateAnnouncementReq{Announcement: announcement}
	return a.client.Post(fmt.Sprintf(updateAnnouncementUri, id), req, nil)
}

// FetchMembers 分页获取聊天室成员
func (a *api) FetchMembers(arg FetchMembersArg) (*FetchMembersRet, error) {
	resp := &fetchMembersResp{}
	if err := a.client.Get(fmt.Sprintf(fetchMembersUri, arg.ID, arg.PageNum, arg.PageSize), nil, resp); err != nil {
		return nil, err
	}

	ret := &FetchMembersRet{
		List:    make([]string, 0, len(resp.Data)),
		HasMore: len(resp.Data) == arg.PageSize,
	}
	for _, item := range resp.Data {
		if item.Owner != "" {
			ret.List = append(ret.List, item.Owner)
		}

		if item.Member != "" {
			ret.List = append(ret.List, item.Member)
		}
	}

	return ret, nil
}

// AddMember 添加单个聊天室成员
func (a *api) AddMember(id, username string) (bool, error) {
	resp := &addMemberResp{}
	if err := a.client.Post(fmt.Sprintf(addMemberUri, id, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Result, nil
}

// AddMembers 批量添加聊天室成员
func (a *api) AddMembers(id string, usernames ...string) ([]string, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 60:
		return nil, errors.New("the number of member exceeds the upper limit")
	}

	req := &addMembersReq{Usernames: usernames}
	resp := &addMembersResp{}
	if err := a.client.Post(fmt.Sprintf(addMembersUri, id), req, resp); err != nil {
		return nil, err
	}

	return resp.Data.NewMembers, nil
}

// RemoveMember 移除单个聊天室成员
func (a *api) RemoveMember(id, username string) (bool, error) {
	rets, err := a.RemoveMembers(id, username)
	if err != nil {
		return false, err
	}

	return rets[0].Result, nil
}

// RemoveMembers 批量移除聊天室成员
func (a *api) RemoveMembers(id string, usernames ...string) ([]*ActionResult, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 100:
		return nil, errors.New("the number of member exceeds the upper limit")
	}

	uri := fmt.Sprintf(removeMembersUri, id, strings.Join(usernames, ","))
	if len(usernames) > 1 {
		resp := &removeMembersResp{}
		if err := a.client.Delete(uri, nil, resp); err != nil {
			return nil, err
		}
		return resp.Data, nil
	} else {
		resp := &removeMemberResp{}
		if err := a.client.Delete(uri, nil, resp); err != nil {
			return nil, err
		}
		return []*ActionResult{resp.Data}, nil
	}
}

// GetAdmins 获取聊天室管理员列表
func (a *api) GetAdmins(id string) ([]string, error) {
	resp := &getAdminResp{}
	if err := a.client.Get(fmt.Sprintf(getAdminsUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddAdmin 添加聊天室管理员
func (a *api) AddAdmin(id, username string) (bool, error) {
	req := &addAdminReq{NewAdmin: username}
	resp := &addAdminResp{}
	if err := a.client.Post(fmt.Sprintf(addAdminUri, id), req, resp); err != nil {
		return false, err
	}

	return resp.Data.Result == "success", nil
}

// RemoveAdmin 移除聊天室管理员
func (a *api) RemoveAdmin(id, username string) (bool, error) {
	resp := &removeAdminResp{}
	if err := a.client.Delete(fmt.Sprintf(removeAdminUri, id, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Result == "success", nil
}

// GetBlacklists 查询聊天室黑名单
func (a *api) GetBlacklists(id string) ([]string, error) {
	resp := &getBlacklistsResp{}
	if err := a.client.Get(fmt.Sprintf(getBlacklistsUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddBlacklist 添加单个用户至聊天室黑名单
func (a *api) AddBlacklist(id, username string) (bool, error) {
	resp := &addBlacklistResp{}
	if err := a.client.Post(fmt.Sprintf(addBlacklistUri, id, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Result, nil
}

// AddBlacklists 批量添加用户至聊天室黑名单
func (a *api) AddBlacklists(id string, usernames ...string) ([]*ActionResult, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 60:
		return nil, errors.New("the number of user exceeds the upper limit")
	}

	req := &addBlacklistsReq{Usernames: usernames}
	resp := &addBlacklistsResp{}
	if err := a.client.Post(fmt.Sprintf(addBlacklistsUri, id), req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// RemoveBlacklist 从聊天室黑名单移除单个用户
func (a *api) RemoveBlacklist(id, username string) (bool, error) {
	resp := &removeBlacklistResp{}
	if err := a.client.Delete(fmt.Sprintf(removeBlacklistUri, id, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Result, nil
}

// RemoveBlacklists 批量添加用户至聊天室黑名单
func (a *api) RemoveBlacklists(id string, usernames ...string) ([]*ActionResult, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 60:
		return nil, errors.New("the number of user exceeds the upper limit")
	}

	uri := fmt.Sprintf(removeBlacklistsUri, id, strings.Join(usernames, ","))
	if len(usernames) > 1 {
		resp := &removeBlacklistsResp{}
		if err := a.client.Delete(uri, nil, resp); err != nil {
			return nil, err
		}
		return resp.Data, nil
	} else {
		resp := &removeBlacklistResp{}
		if err := a.client.Delete(uri, nil, resp); err != nil {
			return nil, err
		}
		return []*ActionResult{resp.Data}, nil
	}
}

// GetWhitelists 查询聊天室白名单
func (a *api) GetWhitelists(id string) ([]string, error) {
	resp := &getWhitelistsResp{}
	if err := a.client.Get(fmt.Sprintf(getWhitelistsUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddWhitelist 添加单个用户至聊天室黑名单
func (a *api) AddWhitelist(id, username string) (bool, error) {
	resp := &addBlacklistResp{}
	if err := a.client.Post(fmt.Sprintf(addWhitelistUri, id, username), nil, resp); err != nil {
		return false, err
	}

	return resp.Data.Result, nil
}

// AddWhitelists 批量添加用户至聊天室白名单
func (a *api) AddWhitelists(id string, usernames ...string) ([]*ActionResult, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 60:
		return nil, errors.New("the number of user exceeds the upper limit")
	}

	req := &addWhitelistsReq{Usernames: usernames}
	resp := &addWhitelistsResp{}
	if err := a.client.Post(fmt.Sprintf(addWhitelistsUri, id), req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// RemoveWhitelist 从聊天室白名单移除单个用户
func (a *api) RemoveWhitelist(id, username string) (bool, error) {
	rets, err := a.RemoveWhitelists(id, username)
	if err != nil {
		return false, err
	}

	return rets[0].Result, nil
}

// RemoveWhitelists 将用户批量移除聊天室白名单
func (a *api) RemoveWhitelists(id string, usernames ...string) ([]*ActionResult, error) {
	switch count := len(usernames); {
	case count == 0:
		return nil, nil
	case count > 60:
		return nil, errors.New("the number of user exceeds the upper limit")
	}

	uri := fmt.Sprintf(removeWhitelistsUri, id, strings.Join(usernames, ","))
	if len(usernames) > 1 {
		resp := &removeWhitelistsResp{}
		if err := a.client.Delete(uri, nil, resp); err != nil {
			return nil, err
		}
		return resp.Data, nil
	} else {
		resp := &removeWhitelistResp{}
		if err := a.client.Delete(uri, nil, resp); err != nil {
			return nil, err
		}
		return []*ActionResult{resp.Data}, nil
	}
}

// GetMutes 获取禁言列表
func (a *api) GetMutes(id string) ([]*Mute, error) {
	resp := &getMutesResp{}
	if err := a.client.Get(fmt.Sprintf(getMutesUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddMute 禁言单个聊天室成员
func (a *api) AddMute(id string, duration int64, username string) (bool, error) {
	results, err := a.AddMutes(id, duration, username)
	if err != nil {
		return false, err
	}

	return results[0].Result, nil
}

// AddMutes 禁言聊天室成员
func (a *api) AddMutes(id string, duration int64, usernames ...string) ([]*AddMuteResult, error) {
	req := &addMutesReq{MuteDuration: duration, Usernames: usernames}
	resp := &addMutesResp{}
	if err := a.client.Post(fmt.Sprintf(addMutesUri, id), req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// RemoveMute 解除单个聊天室禁言成员
func (a *api) RemoveMute(id string, username string) (bool, error) {
	results, err := a.RemoveMutes(id, username)
	if err != nil {
		return false, err
	}

	return results[0].Result, nil
}

// RemoveMutes 解除聊天室禁言成员
func (a *api) RemoveMutes(id string, usernames ...string) ([]*RemoveMuteResult, error) {
	resp := &removeMutesResp{}
	if err := a.client.Delete(fmt.Sprintf(removeMutesUri, id, strings.Join(usernames, ",")), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddAllMutes 禁言聊天室全体成员
func (a *api) AddAllMutes(id string) error {
	return a.client.Post(fmt.Sprintf(addAllMutesUri, id), nil, nil)
}

// RemoveAllMutes 解除聊天室全员禁言
func (a *api) RemoveAllMutes(id string) error {
	return a.client.Delete(fmt.Sprintf(removeAllMutesUri, id), nil, nil)
}
