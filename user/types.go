package user

type User struct {
	Username string `json:"username"` // （必填）用户 ID，长度不可超过 64 个字节长度。
	Password string `json:"password"` // （必填）用户的登录密码，长度不可超过 64 个字符。
	Nickname string `json:"nickname"` // （选填）推送消息时，在消息推送通知栏内显示的用户昵称，并非用户个人信息的昵称。长度不可超过 100 个字符。
}

type registerUsersResp struct {
	Entities []*Entity `json:"entities"`
}

type getResp struct {
	Entities []map[string]interface{} `json:"entities"`
}

type Entity struct {
	UUID                          string   `json:"uuid"`
	Type                          string   `json:"type"`
	Created                       int64    `json:"created"`
	Modified                      int64    `json:"modified"`
	Username                      string   `json:"username"`
	Activated                     bool     `json:"activated"`
	Nickname                      string   `json:"nickname"`
	NotificationDisplayStyle      int      `json:"notification_display_style"`
	NotificationNoDisturbing      bool     `json:"notification_no_disturbing"`
	NotificationNoDisturbingStart string   `json:"notification_no_disturbing_start"`
	NotificationNoDisturbingEnd   string   `json:"notification_no_disturbing_end"`
	NotificationIgnoreGroups      []string `json:"notification_ignore_groups"`
	NotifierName                  string   `json:"notifier_name"`
	DeviceToken                   string   `json:"device_token"`
}

type FetchUserArg struct {
	Limit  int    `json:"limit"`  // （选填）请求查询用户的数量。取值范围[1,100]，默认值为 10。若实际用户数量超过 100，返回 100 个用户。
	Cursor string `json:"cursor"` // （选填）开始获取数据的游标位置，用于分页显示用户列表。第一次发起批量查询用户请求时若不设置 cursor，请求成功后会获得第一页用户列表。从响应 body 中获取 cursor，并在下一次请求的 URL 中传入该 cursor，直到响应 body 中不再有 cursor 字段，则表示已查询到 app 中所有用户。
}

type FetchUsersRet struct {
	List    []*Entity `json:"list"`
	HasMore bool      `json:"has_more"`
	Cursor  string    `json:"cursor"`
}

type fetchUsersResp struct {
	Entities []map[string]interface{} `json:"entities"`
	Cursor   string                   `json:"cursor"`
	Count    int                      `json:"count"`
}

type deleteUsersResp struct {
	Entities []map[string]interface{} `json:"entities"`
}

type updatePasswordReq struct {
	NewPassword string `json:"newpassword"` // （必填）新密码
}

type getOnlineStatusResp struct {
	Data map[string]string `json:"data"`
}

type batchGetOnlineStatusReq struct {
	Usernames []string `json:"usernames"` // （必填）要查询状态的用户 ID，以数组方式提交，最多不能超过 100 个。
}
type batchGetOnlineStatusResp struct {
	Data []map[string]string `json:"data"`
}

type Mutes struct {
	Username  string `json:"username"`            // （必填）设置禁言配置的用户 ID。
	Chat      *int   `json:"chat,omitempty"`      // （选填）单聊消息禁言时长，单位为秒，最大值为 2147483647。(> 0：该用户 ID 具体的单聊消息禁言时长。0：取消该用户的单聊消息禁言。-1：该用户被设置永久单聊消息禁言。其他负值无效。)
	Groupchat *int   `json:"groupchat,omitempty"` // （选填）群组消息禁言时长，单位为秒，规则同上。
	Chatroom  *int   `json:"chatroom,omitempty"`  // （选填）聊天室消息禁言时长，单位为秒，规则同上。
	Unixtime  int    `json:"-"`
}

type getMutesResp struct {
	Data getMutesRet `json:"data"`
}

type getMutesRet struct {
	Username  string `json:"userid"`    // 设置禁言配置的用户 ID。
	Chat      int    `json:"chat"`      // 单聊消息禁言时长，单位为秒，最大值为 2147483647。(> 0：该用户 ID 具体的单聊消息禁言时长。0：取消该用户的单聊消息禁言。-1：该用户被设置永久单聊消息禁言。其他负值无效。)
	Groupchat int    `json:"groupchat"` // 群组消息禁言时长，单位为秒，规则同上。
	Chatroom  int    `json:"chatroom"`  // 聊天室消息禁言时长，单位为秒，规则同上。
	Unixtime  int    `json:"unixtime"`  // 当前操作的 Unix 时间戳。
}

type FetchMutesArg struct {
	PageNum  int `json:"pageNum"`  // （必填）请求查询的页码。
	PageSize int `json:"pageSize"` // （必填）请求查询每页显示的禁言用户的数量。
}

type FetchMutesRet struct {
	List     []*MutesRet `json:"list"`     // 列表
	HasMore  bool        `json:"has_more"` // 是否还有更多数据
	Unixtime int         `json:"unixtime"` // 当前操作的 Unix 时间戳。
}

type fetchMutesResp struct {
	Data struct {
		Data     []*MutesRet `json:"data"`     // 数据
		Unixtime int         `json:"unixtime"` // 当前操作的 Unix 时间戳。
	} `json:"data"`
}

type MutesRet struct {
	Username  string `json:"username"`  // 设置禁言配置的用户 ID。
	Chat      int    `json:"chat"`      // 单聊消息禁言时长，单位为秒，最大值为 2147483647。(> 0：该用户 ID 具体的单聊消息禁言时长。0：取消该用户的单聊消息禁言。-1：该用户被设置永久单聊消息禁言。其他负值无效。)
	Groupchat int    `json:"groupchat"` // 群组消息禁言时长，单位为秒，规则同上。
	Chatroom  int    `json:"chatroom"`  // 聊天室消息禁言时长，单位为秒，规则同上。
	Unixtime  int    `json:"unixtime"`  // 当前操作的 Unix 时间戳。
}

type getOfflineMsgCountResp struct {
	Data map[string]int `json:"data"`
}

type getOfflineMsgStatusResp struct {
	Data map[string]string `json:"data"`
}

type deactivateResp struct {
	Entities []map[string]interface{} `json:"entities"`
}

type offlineResp struct {
	Data struct {
		Result bool `json:"result"`
	} `json:"data"`
}

type getFriendsResp struct {
	Data []string `json:"data"`
}

type addBlacklistsReq struct {
	Usernames []string `json:"usernames"`
}

type getBlacklistsResp struct {
	Data []string `json:"data"`
}

type getMetadataResp struct {
	Data map[string]string `json:"data"`
}

type deleteMetadataResp struct {
	Data bool `json:"data"`
}

type batchGetMetadataReq struct {
	Properties []string `json:"properties"`
	Targets    []string `json:"targets"`
}

type batchGetMetadataResp struct {
	Data map[string]map[string]string `json:"data"`
}

type getCapacityResp struct {
	Data int64 `json:"data"`
}

type setOfflinePushNicknameReq struct {
	Nickname string `json:"nickname"`
}

type setOfflinePushDisplayStyleReq struct {
	NotificationDisplayStyle int `json:"notification_display_style"`
}

type setOfflinePushNoDisturbingReq struct {
	NotificationNoDisturbing      bool   `json:"notification_no_disturbing"`
	NotificationNoDisturbingStart string `json:"notification_no_disturbing_start"`
	NotificationNoDisturbingEnd   string `json:"notification_no_disturbing_end"`
}

type setOfflinePushLanguageReq struct {
	TranslationLanguage string `json:"translationLanguage"`
}

type getOfflinePushLanguageResp struct {
	Data struct {
		Language string `json:"language"`
	} `json:"data"`
}

type SetOfflinePushTargetedNoDisturbingArg struct {
	Username       string // （必填）用户名
	ToType         string // （必填）对象类型，即会话类型：- user：用户，表示单聊；- chatgroup：群组，表示群聊。
	ToKey          string // （必填）对象名称：- 单聊时为对端用户的用户 ID；- 群聊时为群组 ID。
	Type           string // （选填）离线推送通知方式：- DEFAULT：默认值，采用全局配置；- ALL：接收全部离线消息的推送通知；- AT：只接收提及当前用户的离线消息的推送通知；- NONE：不接收离线消息的推送通知。
	IgnoreInterval string // （选填）离线推送免打扰时间段，精确到分钟，格式为 HH:MM-HH:MM，例如 08:30-10:00。该时间为 24 小时制，免打扰时间段的开始时间和结束时间中的小时数和分钟数的取值范围分别为 [00,23] 和 [00,59]。免打扰时段的设置仅针对 app 生效，对单聊或群聊不生效。如需设置 app 的免打扰时段，type 指定为 user，key 指定为当前用户 ID。
	IgnoreDuration int64  // （选填）离线推送免打扰时长，单位为毫秒。该参数的取值范围为 [0,604800000]，0 表示该参数无效，604800000 表示免打扰模式持续 7 天。
}

type setOfflinePushTargetedNoDisturbingReq struct {
	Type           string `json:"type"`
	IgnoreInterval string `json:"ignoreInterval"`
	IgnoreDuration int64  `json:"ignoreDuration"`
}

type getOfflinePushTargetedNoDisturbingResp struct {
	Data *NoDisturbing `json:"data"`
}

type NoDisturbing struct {
	Type           string `json:"type"`
	IgnoreInterval string `json:"ignoreInterval"`
	IgnoreDuration int64  `json:"ignoreDuration"`
}

type getJoinedChatroomsResp struct {
	Data []*JoinedChatroom `json:"data"`
}

type JoinedChatroom struct {
	ID   string `json:"id"`   // 聊天室 ID，聊天室唯一标识，由环信即时通讯 IM 服务器生成。
	Name string `json:"name"` // 聊天室名称，最大长度为 128 字符。
}

type getJoinedGroupsResp struct {
	Data []*JoinedGroup `json:"data"`
}

type JoinedGroup struct {
	ID   string `json:"groupid"`   // 群组 ID。
	Name string `json:"groupname"` // 群组名称。
}

type FetchJoinedThreadsArg struct {
	Username string `json:"username"`
	Limit    int    `json:"limit"`
	Cursor   string `json:"cursor"`
	Sort     string `json:"sort"`
}

type FetchJoinedThreadsRet struct {
	List    []*Thread `json:"list"`
	HasMore bool      `json:"has_more"`
	Cursor  string    `json:"cursor"`
}

type Thread struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	MsgID   string `json:"msgId"`
	GroupID string `json:"groupId"`
	Created int64  `json:"created"`
}

type fetchJoinedThreadsResp struct {
	Entities   []*Thread `json:"entities"`
	Properties struct {
		Cursor string `json:"cursor"`
	} `json:"properties"`
}
