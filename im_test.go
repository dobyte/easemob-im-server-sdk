package im_test

import (
	"github.com/dobyte/easemob-im-server-sdk"
	"github.com/dobyte/easemob-im-server-sdk/chatroom"
	"github.com/dobyte/easemob-im-server-sdk/group"
	"github.com/dobyte/easemob-im-server-sdk/user"
	"testing"
)

var sdk im.IM

const (
	defaultChatroomID  = "188688613048322"
	defaultUsername1   = "test1"
	defaultUsername2   = "test2"
	defaultOldPassword = "123456"
	defaultNewPassword = "456123"
	defaultTemplate    = "test"
	defaultGroupID     = "188864710901761"
)

func init() {
	sdk = im.NewIM(&im.Options{
		Host:         "a1.easemob.com",
		AppKey:       "",
		ClientID:     "",
		ClientSecret: "",
	})
}

func TestIM_User_Register(t *testing.T) {
	entity, err := sdk.User().RegisterUsers(user.User{
		Username: defaultUsername1,
		Password: defaultOldPassword,
		Nickname: "test-1",
	}, user.User{
		Username: defaultUsername2,
		Password: defaultOldPassword,
		Nickname: "test-2",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", entity)
}

func TestIM_User_Get(t *testing.T) {
	entity, err := sdk.User().GetUser(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", entity)
}

func TestIM_User_Delete(t *testing.T) {
	err := sdk.User().DeleteUser(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_BatchDeleteUsers(t *testing.T) {
	entities, err := sdk.User().DeleteUsers(2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", entities)
}

func TestIM_User_BatchDeleteAllUsers(t *testing.T) {
	entities, err := sdk.User().DeleteAllUsers()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", entities)
}

func TestIM_User_ModifyUserPassword(t *testing.T) {
	err := sdk.User().UpdatePassword(defaultUsername1, defaultNewPassword)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetOnlineStatus(t *testing.T) {
	status, err := sdk.User().GetOnlineStatus(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(status)
}

func TestIM_User_BatchGetOnlineStatus(t *testing.T) {
	statuses, err := sdk.User().GetOnlineStatuses(defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(statuses)
}

func TestIM_User_SetMutes(t *testing.T) {
	chat := 10
	groupchat := 10

	err := sdk.User().SetMutes(user.Mutes{
		Username:  defaultUsername1,
		Chat:      &chat,
		Groupchat: &groupchat,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetMutes(t *testing.T) {
	ret, err := sdk.User().GetMutes(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ret)
}

func TestIM_User_FetchMutes(t *testing.T) {
	ret, err := sdk.User().FetchMutes(user.FetchMutesArg{
		PageNum:  1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ret)
}

func TestIM_User_GetOfflineMsgCount(t *testing.T) {
	count, err := sdk.User().GetOfflineMsgCount(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%d", count)
}

func TestIM_User_GetOfflineMsgStatus(t *testing.T) {
	status, err := sdk.User().GetOfflineMsgStatus(defaultUsername1, "123")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", status)
}

func TestIM_User_DeactivateUser(t *testing.T) {
	entity, err := sdk.User().DeactivateUser(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", entity)
}

func TestIM_User_ActivateUser(t *testing.T) {
	err := sdk.User().ActivateUser(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_OfflineUser(t *testing.T) {
	ok, err := sdk.User().OfflineUser(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIM_User_AddFriend(t *testing.T) {
	err := sdk.User().AddFriend(defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_RemoveFriend(t *testing.T) {
	err := sdk.User().RemoveFriend(defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetFriends(t *testing.T) {
	friends, err := sdk.User().GetFriends(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(friends)
}

func TestIM_User_AddBlacklists(t *testing.T) {
	err := sdk.User().AddBlacklists(defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_RemoveBlacklist(t *testing.T) {
	err := sdk.User().RemoveBlacklist(defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetBlacklists(t *testing.T) {
	blacklists, err := sdk.User().GetBlacklists(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blacklists)
}

func TestIM_User_SetMetadata(t *testing.T) {
	err := sdk.User().SetMetadata(defaultUsername1, map[string]string{
		"avatarurl": "http://www.baidu.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetMetadata(t *testing.T) {
	metadata, err := sdk.User().GetMetadata(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(metadata)
}

func TestIM_User_BatchGetMetadata(t *testing.T) {
	metadata, err := sdk.User().BatchGetMetadata([]string{
		"avatarurl",
	}, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(metadata)
}

func TestIM_User_DeleteMetadata(t *testing.T) {
	ok, err := sdk.User().DeleteMetadata(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIM_User_GetCapacity(t *testing.T) {
	capacity, err := sdk.User().GetCapacity()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(capacity)
}

func TestIM_User_SetOfflinePushNickname(t *testing.T) {
	err := sdk.User().SetOfflinePushNickname(defaultUsername1, "test")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_SetOfflinePushDisplayStyle(t *testing.T) {
	err := sdk.User().SetOfflinePushDisplayStyle(defaultUsername1, 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_EnableOfflinePushNoDisturbing(t *testing.T) {
	err := sdk.User().EnableOfflinePushNoDisturbing(defaultUsername1, 8, 23)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_DisableOfflinePushNoDisturbing(t *testing.T) {
	err := sdk.User().DisableOfflinePushNoDisturbing(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_SetOfflinePushTargetedNoDisturbing(t *testing.T) {
	err := sdk.User().SetOfflinePushTargetedNoDisturbing(&user.SetOfflinePushTargetedNoDisturbingArg{
		Username:       defaultUsername1,
		ToType:         "user",
		ToKey:          defaultUsername2,
		Type:           "NONE",
		IgnoreInterval: "21:30-08:00",
		IgnoreDuration: 1647590149924,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetOfflinePushTargetedNoDisturbing(t *testing.T) {
	noDisturbing, err := sdk.User().GetOfflinePushTargetedNoDisturbing(defaultUsername1, "user", defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", noDisturbing)
}

func TestIM_User_SetOfflinePushLanguage(t *testing.T) {
	err := sdk.User().SetOfflinePushLanguage(defaultUsername1, "EU")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIM_User_GetOfflinePushLanguage(t *testing.T) {
	language, err := sdk.User().GetOfflinePushLanguage(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(language)
}

func TestIm_User_GetJoinedChatrooms(t *testing.T) {
	chatrooms, err := sdk.User().GetJoinedChatrooms(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range chatrooms {
		t.Logf("%+v", item)
	}
}

func TestIm_User_GetJoinedGroups(t *testing.T) {
	groups, err := sdk.User().GetJoinedGroups(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range groups {
		t.Logf("%+v", item)
	}
}

func TestIm_Push_GetTemplate(t *testing.T) {
	template, err := sdk.Push().GetTemplate(defaultTemplate)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", template)
}

func TestIm_Push_CreateTemplate(t *testing.T) {
	err := sdk.Push().CreateTemplate(defaultTemplate, "你好,{0}", "推送测试,{0}")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Push_DeleteTemplate(t *testing.T) {
	err := sdk.Push().DeleteTemplate(defaultTemplate)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Chatroom_AddSuperAdmin(t *testing.T) {
	ok, err := sdk.Chatroom().AddSuperAdmin(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_RevokeSuperAdmin(t *testing.T) {
	err := sdk.Chatroom().RevokeSuperAdmin(defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Chatroom_FetchSuperAdmins(t *testing.T) {
	ret, err := sdk.Chatroom().FetchSuperAdmins(chatroom.FetchSuperAdminsArg{
		PageNum:  3,
		PageSize: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ret)
}

func TestIm_Chatroom_GetAllChatrooms(t *testing.T) {
	chatrooms, err := sdk.Chatroom().GetAllChatrooms()
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range chatrooms {
		t.Logf("%+v", item)
	}
}

func TestIm_Chatroom_GetChatrooms(t *testing.T) {
	chatrooms, err := sdk.Chatroom().GetChatrooms(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range chatrooms {
		t.Logf("%+v", item)
	}
}

func TestIm_Chatroom_CreateChatroom(t *testing.T) {
	id, err := sdk.Chatroom().CreateChatroom(&chatroom.CreateChatRoomArg{
		Name:        "testchatroom1",
		Description: "This is a chat room for test",
		MaxUsers:    100,
		Owner:       defaultUsername1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)
}

func TestIm_Chatroom_ModifyChatroom(t *testing.T) {
	ok, err := sdk.Chatroom().UpdateChatroom(chatroom.UpdateChatroomArg{
		ID:          defaultChatroomID,
		Name:        "testchatroom2",
		Description: "This is a chat room for test",
		MaxUsers:    200,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_DeleteChatroom(t *testing.T) {
	ok, err := sdk.Chatroom().DeleteChatroom(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_GetAnnouncement(t *testing.T) {
	announcement, err := sdk.Chatroom().GetAnnouncement(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(announcement)
}

func TestIm_Chatroom_UpdateAnnouncement(t *testing.T) {
	err := sdk.Chatroom().UpdateAnnouncement(defaultChatroomID, "aaa")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Chatroom_FetchMembers(t *testing.T) {
	members, err := sdk.Chatroom().FetchMembers(chatroom.FetchMembersArg{
		ID:       defaultChatroomID,
		PageNum:  1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(members)
}

func TestIm_Chatroom_AddMember(t *testing.T) {
	ok, err := sdk.Chatroom().AddMember(defaultChatroomID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_AddMembers(t *testing.T) {
	members, err := sdk.Chatroom().AddMembers(defaultChatroomID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(members)
}

func TestIm_Chatroom_RemoveMember(t *testing.T) {
	ok, err := sdk.Chatroom().RemoveMember(defaultChatroomID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_RemoveMembers(t *testing.T) {
	rets, err := sdk.Chatroom().RemoveMembers(defaultChatroomID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Chatroom_GetAdmins(t *testing.T) {
	admins, err := sdk.Chatroom().GetAdmins(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(admins)
}

func TestIm_Chatroom_AddAdmin(t *testing.T) {
	ok, err := sdk.Chatroom().AddAdmin(defaultChatroomID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_RemoveAdmin(t *testing.T) {
	ok, err := sdk.Chatroom().RemoveAdmin(defaultChatroomID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_GetBlacklists(t *testing.T) {
	blacklists, err := sdk.Chatroom().GetBlacklists(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blacklists)
}

func TestIm_Chatroom_AddBlacklist(t *testing.T) {
	ok, err := sdk.Chatroom().AddBlacklist(defaultChatroomID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_AddBlacklists(t *testing.T) {
	results, err := sdk.Chatroom().AddBlacklists(defaultChatroomID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range results {
		t.Logf("%+v", result)
	}
}

func TestIm_Chatroom_RemoveBlacklist(t *testing.T) {
	ok, err := sdk.Chatroom().RemoveBlacklist(defaultChatroomID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_RemoveBlacklists(t *testing.T) {
	rets, err := sdk.Chatroom().RemoveBlacklists(defaultChatroomID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Chatroom_GetWhitelists(t *testing.T) {
	blacklists, err := sdk.Chatroom().GetWhitelists(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blacklists)
}

func TestIm_Chatroom_AddWhitelist(t *testing.T) {
	ok, err := sdk.Chatroom().AddWhitelist(defaultChatroomID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_AddWhitelists(t *testing.T) {
	results, err := sdk.Chatroom().AddWhitelists(defaultChatroomID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range results {
		t.Logf("%+v", result)
	}
}

func TestIm_Chatroom_RemoveWhitelist(t *testing.T) {
	ok, err := sdk.Chatroom().RemoveWhitelist(defaultChatroomID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ok)
}

func TestIm_Chatroom_RemoveWhitelists(t *testing.T) {
	rets, err := sdk.Chatroom().RemoveWhitelists(defaultChatroomID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Chatroom_GetMutes(t *testing.T) {
	mutes, err := sdk.Chatroom().GetMutes(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	for _, mute := range mutes {
		t.Logf("%+v", mute)
	}
}

func TestIm_Chatroom_AddMutes(t *testing.T) {
	rets, err := sdk.Chatroom().AddMutes(defaultChatroomID, 5000, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Chatroom_RemoveMutes(t *testing.T) {
	rets, err := sdk.Chatroom().RemoveMutes(defaultChatroomID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Chatroom_AddAllMutes(t *testing.T) {
	err := sdk.Chatroom().AddAllMutes(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Chatroom_RemoveAllMutes(t *testing.T) {
	err := sdk.Chatroom().RemoveAllMutes(defaultChatroomID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_GetGroup(t *testing.T) {
	detail, err := sdk.Group().GetGroup(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", detail)
}

func TestIm_Group_CreateGroup(t *testing.T) {
	id, err := sdk.Group().CreateGroup(&group.CreateGroupArg{
		Name:        "test-group",
		Description: "this is a desc of group",
		Public:      true,
		Owner:       defaultUsername1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(id)
}

func TestIm_Group_UpdateGroup(t *testing.T) {
	name := "test-group-new"
	description := "this is a desc of group"
	maxUsers := 300
	allowInvites := true
	membersOnly := false
	custom := "aaa"

	ret, err := sdk.Group().UpdateGroup(&group.UpdateGroupArg{
		ID:           defaultGroupID,
		Name:         &name,
		Description:  &description,
		MaxUsers:     &maxUsers,
		AllowInvites: &allowInvites,
		MembersOnly:  &membersOnly,
		Custom:       &custom,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ret)
}

func TestIm_Group_DeleteGroup(t *testing.T) {
	err := sdk.Group().DeleteGroup(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_GetAllGroups(t *testing.T) {
	groups, err := sdk.Group().GetAllGroups()
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range groups {
		t.Logf("%+v", item)
	}
}

func TestIm_Group_FetchGroups(t *testing.T) {
	ret, err := sdk.Group().FetchGroups(group.FetchGroupsArg{
		Limit:  10,
		Cursor: "",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", ret)

	for _, item := range ret.List {
		t.Logf("%+v", item)
	}
}

func TestIm_Group_GetAnnouncement(t *testing.T) {
	announcement, err := sdk.Group().GetAnnouncement(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(announcement)
}

func TestIm_Group_UpdateAnnouncement(t *testing.T) {
	err := sdk.Group().UpdateAnnouncement(defaultGroupID, "aaa")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_GetAllShareFiles(t *testing.T) {
	files, err := sdk.Group().GetAllShareFiles(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		t.Logf("%+v", file)
	}
}

func TestIm_Group_FetchMembers(t *testing.T) {
	members, err := sdk.Group().FetchMembers(group.FetchMembersArg{
		ID:       defaultGroupID,
		PageNum:  1,
		PageSize: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(members)
}

func TestIm_Group_AddMember(t *testing.T) {
	err := sdk.Group().AddMember(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_AddMembers(t *testing.T) {
	members, err := sdk.Group().AddMembers(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(members)
}

func TestIm_Group_RemoveMember(t *testing.T) {
	err := sdk.Group().RemoveMember(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_RemoveMembers(t *testing.T) {
	rets, err := sdk.Group().RemoveMembers(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Group_GetAdmins(t *testing.T) {
	admins, err := sdk.Group().GetAdmins(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(admins)
}

func TestIm_Group_AddAdmin(t *testing.T) {
	err := sdk.Group().AddAdmin(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_RemoveAdmin(t *testing.T) {
	err := sdk.Group().RemoveAdmin(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_TransferGroup(t *testing.T) {
	err := sdk.Group().TransferGroup(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_GetBlacklists(t *testing.T) {
	blacklists, err := sdk.Group().GetBlacklists(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blacklists)
}

func TestIm_Group_AddBlacklist(t *testing.T) {
	err := sdk.Group().AddBlacklist(defaultGroupID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_AddBlacklists(t *testing.T) {
	results, err := sdk.Group().AddBlacklists(defaultGroupID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	for _, result := range results {
		t.Logf("%+v", result)
	}
}

func TestIm_Group_RemoveBlacklist(t *testing.T) {
	err := sdk.Group().RemoveBlacklist(defaultGroupID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_RemoveBlacklists(t *testing.T) {
	rets, err := sdk.Group().RemoveBlacklists(defaultGroupID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}

func TestIm_Group_GetWhitelists(t *testing.T) {
	blacklists, err := sdk.Group().GetWhitelists(defaultGroupID)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blacklists)
}

func TestIm_Group_AddWhitelist(t *testing.T) {
	err := sdk.Group().AddWhitelist(defaultGroupID, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_RemoveWhitelist(t *testing.T) {
	err := sdk.Group().RemoveWhitelist(defaultGroupID, defaultUsername1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

func TestIm_Group_RemoveWhitelists(t *testing.T) {
	rets, err := sdk.Group().RemoveWhitelists(defaultGroupID, defaultUsername1, defaultUsername2)
	if err != nil {
		t.Fatal(err)
	}

	for _, ret := range rets {
		t.Logf("%+v", ret)
	}
}
