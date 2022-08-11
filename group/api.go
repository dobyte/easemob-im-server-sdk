package group

import (
	"errors"
	"fmt"
	"github.com/dobyte/easemob-im-server-sdk/internal/core"
	"strings"
)

const (
	getGroupUri              = "/chatgroups/%s"
	createGroupUri           = "/chatgroups"
	updateGroupUri           = "/chatgroups/%s"
	deleteGroupUri           = "/chatgroups/%s"
	getAllGroupsUri          = "/chatgroups"
	fetchGroupsUri           = "/chatgroups?limit=%d&cursor=%s"
	getAnnouncementUri       = "/chatgroups/%s/announcement"
	updateAnnouncementUri    = "/chatgroups/%s/announcement"
	getAllShareFilesUri      = "/chatgroups/%s/share_files"
	fetchShareFilesUri       = "/chatgroups/%s/share_files?pagenum=%d&pagesize=%d"
	getShareFileUri          = "/chatgroups/%s/share_files/%s"
	deleteShareFileUri       = "/chatgroups/%s/share_files/%s"
	fetchMembersUri          = "/chatgroups/%s/users?pagenum=%d&pagesize=%d"
	addMemberUri             = "/chatgroups/%s/users/%s"
	addMembersUri            = "/chatgroups/%s/users"
	removeMembersUri         = "/chatgroups/%s/users/%s"
	getAdminsUri             = "/chatgroups/%s/admin"
	addAdminUri              = "/chatgroups/%s/admin"
	removeAdminUri           = "/chatgroups/%s/admin/%s"
	transferGroupUri         = "/chatgroups/%s"
	getBlacklistsUri         = "/chatgroups/%s/blocks/users"
	addBlacklistUri          = "/chatgroups/%s/blocks/users/%s"
	addBlacklistsUri         = "/chatgroups/%s/blocks/users"
	removeBlacklistUri       = "/chatgroups/%s/blocks/users/%s"
	removeBlacklistsUri      = "/chatgroups/%s/blocks/users/%s"
	getWhitelistsUri         = "/chatgroups/%s/white/users"
	addWhitelistUri          = "/chatgroups/%s/white/users/%s"
	addWhitelistsUri         = "/chatgroups/%s/white/users"
	removeWhitelistsUri      = "/chatgroups/%s/white/users/%s"
	getMutesUri              = "/chatgroups/%s/mute"
	addMutesUri              = "/chatgroups/%s/mute"
	removeMutesUri           = "/chatgroups/%s/mute/%s"
	addAllMutesUri           = "/chatgroups/%s/ban"
	removeAllMutesUri        = "/chatgroups/%s/ban"
	createThreadUri          = "/thread"
	updateThreadUri          = "/thread/%s"
	deleteThreadUri          = "/thread/%s"
	fetchThreadsUri          = "/thread?limit=%d&cursor=%s&sort=%s"
	fetchGroupUserThreadsUri = "/threads/chatgroups/%s/user/%s?limit=%d&cursor=%s&sort=%s"
)

type API interface {
	// GetGroup 获取群组详情
	// 可以获取一个或多个群组的详情。当获取多个群组的详情时，返回所有存在的群组的详情；对于不存在的群组，返回 “group id doesn’t exist”。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取群组详情
	GetGroup(id string) (*Group, error)

	// CreateGroup 创建群组
	// 创建一个群组，并设置群组名称、群组描述、公开群/私有群属性、群成员最大人数（包括群主）、加入公开群是否需要批准、群主、群成员、群组扩展信息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#创建群组
	CreateGroup(arg *CreateGroupArg) (string, error)

	// UpdateGroup 修改群组信息
	// 修改指定的群组信息。仅支持修改 groupname、description、maxusers、membersonly、allowinvites、custom 六个属性。如果传入其他字段，或传入的字段不存在，则不能修改的字段会抛出异常。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#修改群组信息
	UpdateGroup(arg *UpdateGroupArg) (*UpdateGroupRet, error)

	// DeleteGroup 删除群组
	// 删除一个群组的接口。删除群组时会同时删除群组下所有的子区（Thread）。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#删除群组
	DeleteGroup(id string) error

	// GetAllGroups 获取 App 中所有的群组
	// 获取应用下全部的群组信息的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取_app_中所有的群组_可分页
	GetAllGroups() ([]*ListedGroup, error)

	// FetchGroups 分页拉取群组
	// 获取应用下全部的群组信息的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取_app_中所有的群组_可分页
	FetchGroups(arg FetchGroupsArg) (*FetchGroupsRet, error)

	// GetAnnouncement 获取群组公告
	// 获取指定群组 ID 的群组公告。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取群组公告
	GetAnnouncement(id string) (string, error)

	// UpdateAnnouncement 修改聊天室公告
	// 修改指定群组 ID 的群组公告，注意群组公告的内容不能超过 512 个字符。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#修改群组公告
	UpdateAnnouncement(id, announcement string) error

	// GetAllShareFiles 获取群组共享文件
	// 分页获取指定群组 ID 的群组共享文件，之后可以根据 response 中返回的 file_id，file_id 是群组共享文件的唯一标识，调用 下载群组共享文件 接口下载文件，或调用 删除群组共享文件接口删除文件。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取群组共享文件
	GetAllShareFiles(id string) ([]*ShareFile, error)

	// FetchShareFiles 分页拉取群组共享文件
	// 分页获取指定群组 ID 的群组共享文件，之后可以根据 response 中返回的 file_id，file_id 是群组共享文件的唯一标识，调用 下载群组共享文件 接口下载文件，或调用 删除群组共享文件接口删除文件。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取群组共享文件
	FetchShareFiles(arg FetchShareFilesArg) (*FetchShareFilesRet, error)

	// GetShareFile 下载群组共享文件
	// 根据指定的群组 ID 与 file_id 下载群组共享文件，file_id 是通过 获取群组共享文件 接口获取。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#下载群组共享文件
	GetShareFile(groupID, fileID string) (*ShareFile, error)

	// DeleteShareFile 删除群组共享文件
	// 根据指定的群组 ID 与 file_id 删除群组共享文件，file_id 是通过 获取群组共享文件 接口获取。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#删除群组共享文件
	DeleteShareFile(groupID, fileID string) error

	// FetchMembers 分页获取群组成员
	// 可以分页获取群组成员列表的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#分页获取群组成员
	FetchMembers(arg FetchMembersArg) (*FetchMembersRet, error)

	// AddMember 添加单个群组成员
	// 一次给群添加一个成员，不能重复添加同一个成员。如果用户已经是群成员，将添加失败，并返回错误。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#添加单个群组成员
	AddMember(id, username string) error

	// AddMembers 批量添加群组成员
	// 为群组添加多个成员，一次最多可以添加 60 位成员。如果所有用户均已是群成员，将添加失败，并返回错误。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#批量添加群组成员
	AddMembers(id string, usernames ...string) ([]string, error)

	// RemoveMember 移除单个群组成员
	// 从群中移除指定成员。如果被移除用户不是群成员，将移除失败，并返回错误。群成员移除时还会移除他在群下所有加入的子区。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#移除单个群组成员
	RemoveMember(id, username string) error

	// RemoveMembers 批量移除群组成员
	// 移除多名群成员。如果所有被移除用户均不是群成员，将移除失败，并返回错误。移除后也会将用户从该群里加入的子区中移除。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#批量移除群组成员
	RemoveMembers(id string, usernames ...string) ([]*ActionResult, error)

	// GetAdmins 获取群管理员列表
	// 获取群组管理员列表的接口。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取群管理员列表
	GetAdmins(id string) ([]string, error)

	// AddAdmin 添加群管理员
	// 将一个群成员角色权限提升为群管理员。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#添加群管理员
	AddAdmin(id, username string) error

	// RemoveAdmin 移除群管理员
	// 将用户的角色从群管理员降为群普通成员。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#移除群管理员
	RemoveAdmin(id, username string) error

	// TransferGroup 转让群组
	// 修改群主为同一群组中的其他成员。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#转让群组
	TransferGroup(id string, username string) error

	// GetBlacklists 查询群组黑名单
	// 查询一个群组黑名单中的用户列表。位于黑名单中的用户查看不到该群组的信息，也无法收到该群组的消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#查询群组黑名单
	GetBlacklists(id string) ([]string, error)

	// AddBlacklist 添加单个用户至群组黑名单
	// 添加一个用户进入一个群组的黑名单。群主无法被加入群组的黑名单。
	// 用户进入群组黑名单后，会收到消息：You are kicked out of the group xxx。之后，该用户查看不到该群组的信息，也收不到该群组的消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#添加单个用户至群组黑名单
	AddBlacklist(id, username string) error

	// AddBlacklists 批量添加用户至群组黑名单
	// 将多个用户添加一个群组的黑名单。一次最多可以添加 60 个用户至群组黑名单。群主无法被加入群组的黑名单。
	// 用户进入群组黑名单后，会收到消息：You are kicked out of the group xxx。之后，该用户查看不到该群组的信息，也收不到该群组的消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#批量添加用户至群组黑名单
	AddBlacklists(id string, usernames ...string) ([]*ActionResult, error)

	// RemoveBlacklist 从群组黑名单移除单个用户
	// 将指定用户移出群组黑名单。对于群组黑名单中的用户，如果需要将其再次加入群组，需要先将其从群组黑名单中移除。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#从群组黑名单移除单个用户
	RemoveBlacklist(id, username string) error

	// RemoveBlacklists 从群组黑名单批量移除用户
	// 将多名指定用户从群组黑名单中移除。对于群组黑名单中的用户，如果需要将其再次加入群组，需要先将其从群组黑名单中移除。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#从群组黑名单批量移除用户
	RemoveBlacklists(id string, usernames ...string) ([]*ActionResult, error)

	// GetWhitelists 查询群组白名单
	// 查询一个群组白名单中的用户列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#查询群组白名单
	GetWhitelists(id string) ([]string, error)

	// AddWhitelist 添加单个用户至群组白名单
	// 将指定的单个用户添加至群组白名单。用户在添加至群组白名单后，当群组全员被禁言时，仍可以在群组中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#添加单个用户至群组白名单
	AddWhitelist(id, username string) error

	// AddWhitelists 批量添加用户至群组白名单
	// 添加多个用户至群组白名单。你一次最多可添加 60 个用户。用户添加至白名单后在群组全员禁言时仍可以在群组中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#批量添加用户至群组白名单
	AddWhitelists(id string, usernames ...string) ([]*ActionResult, error)

	// RemoveWhitelist 将单个用户移除群组白名单
	// 将指定用户从群组白名单中移除。你每次最多可移除 60 个用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#将用户移除群组白名单
	RemoveWhitelist(id, username string) error

	// RemoveWhitelists 将用户批量移除群组白名单
	// 将指定用户从群组白名单中移除。你每次最多可移除 60 个用户。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#将用户移除群组白名单
	RemoveWhitelists(id string, usernames ...string) ([]*ActionResult, error)

	// GetMutes 获取禁言列表
	// 获取当前群组的禁言用户列表。
	// 将用户从禁言列表中移除。移除后，用户可以正常在群中发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取禁言列表
	GetMutes(id string) ([]*Mute, error)

	// AddMute 禁言指定群成员
	// 对指定群成员禁言。群成员被禁言后，将无法在群中发送消息，也无法在子区里面发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#禁言指定群成员
	AddMute(id string, duration int64, username string) error

	// AddMutes 禁言指定群成员
	// 对指定群成员禁言。群成员被禁言后，将无法在群中发送消息，也无法在子区里面发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#禁言指定群成员
	AddMutes(id string, duration int64, usernames ...string) ([]*AddMuteResult, error)

	// RemoveMute 解除单个成员禁言
	// 将一个或多个群成员移除禁言列表。移除后，群成员可以在群组中正常发送消息。同时也可以在子区里面发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#解除成员禁言
	RemoveMute(id string, username string) error

	// RemoveMutes 批量解除成员禁言
	// 将一个或多个群成员移除禁言列表。移除后，群成员可以在群组中正常发送消息。同时也可以在子区里面发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#批量解除成员禁言
	RemoveMutes(id string, usernames ...string) ([]*RemoveMuteResult, error)

	// AddAllMutes 禁言全体成员
	// 对所有群组成员一键禁言，即将群组的所有成员均加入禁言列表。设置群组全员禁言后，仅群组白名单中的用户可在群组内发消息。同样在子区中发消息也需要在群组的白名单里面。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#禁言全体成员
	AddAllMutes(id string) error

	// RemoveAllMutes 解除全员禁言
	// 一键取消对群组全体成员的禁言。移除后，群成员可以在群组和子区中正常发送消息。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#解除全员禁言
	RemoveAllMutes(id string) error

	// CreateThread 创建子区
	// 创建子区。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#创建子区
	CreateThread(arg CreateThreadArg) (string, error)

	// UpdateThread 修改子区
	// 修改指定子区。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#修改子区
	UpdateThread(id, name string) error

	// DeleteThread 删除子区
	// 删除指定子区。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#删除子区
	DeleteThread(id string) error

	// FetchThreads 分页拉取所有的子区
	// 获取应用下全部的子区列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取_app_中所有的子区_分页获取
	FetchThreads(arg FetchThreadsArg) (*FetchThreadsRet, error)

	// FetchGroupUserThreads 获取一个用户某个群组下加入的所有子区
	// 根据用户 ID 获取该用户在某个群组下加入的子区列表。
	// 点击查看详细文档:
	// https://docs-im.easemob.com/ccim/rest/group#获取一个用户某个群组下加入的所有子区_分页获取
	FetchGroupUserThreads(arg FetchGroupUserThreadsArg) (*FetchGroupUserThreadsRet, error)
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// GetGroup 获取群组详情
func (a *api) GetGroup(id string) (*Group, error) {
	resp := &getGroupResp{}
	if err := a.client.Get(fmt.Sprintf(getGroupUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data[0], nil
}

// CreateGroup 创建群组
func (a *api) CreateGroup(arg *CreateGroupArg) (string, error) {
	resp := &createGroupResp{}
	if err := a.client.Post(createGroupUri, arg, resp); err != nil {
		return "", err
	}

	return resp.Data.ID, nil
}

// UpdateGroup 修改群组信息
func (a *api) UpdateGroup(arg *UpdateGroupArg) (*UpdateGroupRet, error) {
	resp := &updateGroupResp{}
	if err := a.client.Put(fmt.Sprintf(updateGroupUri, arg.ID), arg, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// DeleteGroup 删除群组
func (a *api) DeleteGroup(id string) error {
	return a.client.Delete(fmt.Sprintf(deleteGroupUri, id), nil, nil)
}

// GetAllGroups 获取 App 中所有的群组
func (a *api) GetAllGroups() ([]*ListedGroup, error) {
	resp := &getAllGroupsResp{}
	if err := a.client.Get(getAllGroupsUri, nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// FetchGroups 分页拉取群组
func (a *api) FetchGroups(arg FetchGroupsArg) (*FetchGroupsRet, error) {
	resp := &fetchGroupsResp{}
	if err := a.client.Get(fmt.Sprintf(fetchGroupsUri, arg.Limit, arg.Cursor), nil, resp); err != nil {
		return nil, err
	}

	return &FetchGroupsRet{
		List:    resp.Data,
		Cursor:  resp.Cursor,
		HasMore: resp.Cursor != "",
	}, nil
}

// GetAnnouncement 获取群组公告
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

// GetAllShareFiles 获取群组共享文件
func (a *api) GetAllShareFiles(id string) ([]*ShareFile, error) {
	resp := &getAllShareFilesResp{}
	if err := a.client.Get(fmt.Sprintf(getAllShareFilesUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// FetchShareFiles 分页拉取群组共享文件
func (a *api) FetchShareFiles(arg FetchShareFilesArg) (*FetchShareFilesRet, error) {
	uri := fmt.Sprintf(fetchShareFilesUri, arg.ID, arg.PageNum, arg.PageSize)
	resp := &fetchShareFilesResp{}
	if err := a.client.Get(uri, nil, resp); err != nil {
		return nil, err
	}

	return &FetchShareFilesRet{
		List:    resp.Data,
		HasMore: len(resp.Data) == arg.PageSize,
	}, nil
}

// UploadShareFile 上传群组共享文件
func (a *api) UploadShareFile() (*ShareFile, error) {
	return nil, nil
}

// GetShareFile 下载群组共享文件
func (a *api) GetShareFile(groupID, fileID string) (*ShareFile, error) {
	resp := &getShareFileResp{}
	if err := a.client.Get(fmt.Sprintf(getShareFileUri, groupID, fileID), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// DeleteShareFile 删除群组共享文件
func (a *api) DeleteShareFile(groupID, fileID string) error {
	return a.client.Delete(fmt.Sprintf(deleteShareFileUri, groupID, fileID), nil, nil)
}

// FetchMembers 分页获取群组成员
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

// AddMember 添加单个群组成员
func (a *api) AddMember(id, username string) error {
	return a.client.Post(fmt.Sprintf(addMemberUri, id, username), nil, nil)
}

// AddMembers 批量添加群组成员
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
func (a *api) RemoveMember(id, username string) error {
	if _, err := a.RemoveMembers(id, username); err != nil {
		return err
	}

	return nil
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

// GetAdmins 获取群管理员列表
func (a *api) GetAdmins(id string) ([]string, error) {
	resp := &getAdminResp{}
	if err := a.client.Get(fmt.Sprintf(getAdminsUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddAdmin 添加群管理员
func (a *api) AddAdmin(id, username string) error {
	req := &addAdminReq{NewAdmin: username}
	return a.client.Post(fmt.Sprintf(addAdminUri, id), req, nil)
}

// RemoveAdmin 移除群管理员
func (a *api) RemoveAdmin(id, username string) error {
	return a.client.Delete(fmt.Sprintf(removeAdminUri, id, username), nil, nil)
}

// TransferGroup 转让群组
func (a *api) TransferGroup(id string, username string) error {
	req := &transferGroupReq{NewOwner: username}
	return a.client.Put(fmt.Sprintf(transferGroupUri, id), req, nil)
}

// GetBlacklists 查询群组黑名单
func (a *api) GetBlacklists(id string) ([]string, error) {
	resp := &getBlacklistsResp{}
	if err := a.client.Get(fmt.Sprintf(getBlacklistsUri, id), nil, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// AddBlacklist 添加单个用户至群组黑名单
func (a *api) AddBlacklist(id, username string) error {
	return a.client.Post(fmt.Sprintf(addBlacklistUri, id, username), nil, nil)
}

// AddBlacklists 批量添加用户至群组黑名单
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

// RemoveBlacklist 从群组黑名单移除单个用户
func (a *api) RemoveBlacklist(id, username string) error {
	return a.client.Delete(fmt.Sprintf(removeBlacklistUri, id, username), nil, nil)
}

// RemoveBlacklists 从群组黑名单批量移除用户
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
func (a *api) AddWhitelist(id, username string) error {
	return a.client.Post(fmt.Sprintf(addWhitelistUri, id, username), nil, nil)
}

// AddWhitelists 批量添加用户至群组白名单
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

// RemoveWhitelist 将单个用户移除群组白名单
func (a *api) RemoveWhitelist(id, username string) error {
	_, err := a.RemoveWhitelists(id, username)
	return err
}

// RemoveWhitelists 将用户批量移除群组白名单
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

// AddMute 禁言指定群成员
func (a *api) AddMute(id string, duration int64, username string) error {
	_, err := a.AddMutes(id, duration, username)
	return err
}

// AddMutes 禁言指定群成员
func (a *api) AddMutes(id string, duration int64, usernames ...string) ([]*AddMuteResult, error) {
	req := &addMutesReq{MuteDuration: duration, Usernames: usernames}
	resp := &addMutesResp{}
	if err := a.client.Post(fmt.Sprintf(addMutesUri, id), req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// RemoveMute 解除单个成员禁言
func (a *api) RemoveMute(id string, username string) error {
	_, err := a.RemoveMutes(id, username)
	return err
}

// RemoveMutes 批量解除成员禁言
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

// CreateThread 创建子区
func (a *api) CreateThread(arg CreateThreadArg) (string, error) {
	resp := &createThreadResp{}
	if err := a.client.Post(createThreadUri, arg, resp); err != nil {
		return "", err
	}

	return resp.Data.ID, nil
}

// UpdateThread 修改子区
func (a *api) UpdateThread(id, name string) error {
	req := &updateThreadReq{Name: name}
	return a.client.Put(fmt.Sprintf(updateThreadUri, id), req, nil)
}

// DeleteThread 删除子区
func (a *api) DeleteThread(id string) error {
	return a.client.Put(fmt.Sprintf(deleteThreadUri, id), nil, nil)
}

// FetchThreads 分页拉取所有的子区
func (a *api) FetchThreads(arg FetchThreadsArg) (*FetchThreadsRet, error) {
	uri := fmt.Sprintf(fetchThreadsUri, arg.Limit, arg.Cursor, arg.Sort)
	resp := &fetchThreadsResp{}
	if err := a.client.Get(uri, nil, resp); err != nil {
		return nil, err
	}

	ret := &FetchThreadsRet{
		List:    make([]string, 0, len(resp.Entities)),
		HasMore: resp.Properties.Cursor != "",
		Cursor:  resp.Properties.Cursor,
	}
	for _, entity := range resp.Entities {
		ret.List = append(ret.List, entity.ID)
	}

	return ret, nil
}

// FetchGroupUserThreads 获取一个用户某个群组下加入的所有子区
func (a *api) FetchGroupUserThreads(arg FetchGroupUserThreadsArg) (*FetchGroupUserThreadsRet, error) {
	uri := fmt.Sprintf(fetchGroupUserThreadsUri, arg.GroupID, arg.Username, arg.Limit, arg.Cursor, arg.Sort)
	resp := &fetchGroupUserThreadsResp{}
	if err := a.client.Get(uri, nil, resp); err != nil {
		return nil, err
	}

	return &FetchGroupUserThreadsRet{
		List:    resp.Entities,
		HasMore: resp.Properties.Cursor != "",
		Cursor:  resp.Properties.Cursor,
	}, nil
}
