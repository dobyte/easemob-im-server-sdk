package group

type Group struct {
	Name              string   `json:"name"`                // 群组名称，最大长度为 128 字符。如果有空格，则使用 “+” 代替。
	Description       string   `json:"description"`         // 群组描述，最大长度为 512 字符。如果有空格，则使用 “+” 代替。
	Public            bool     `json:"public"`              // 是否是公开群。- true：公开群；- false：私有群。
	MaxUsers          int      `json:"maxusers"`            // 群组最大成员数（包括群主），值为数值类型，默认值 200，具体上限请参考 环信即时通讯云控制台。
	AllowInvites      bool     `json:"allowinvites"`        // 是否允许群成员邀请别人加入此群：- true：允许群成员邀请人加入此群;- （默认）false：只有群主或者管理员才可以往群里加人。 注：该参数仅对私有群有效，因为公开群不允许群成员邀请其他用户入群。
	MembersOnly       bool     `json:"members_only"`        // 用户申请入群是否需要群主或者群管理员审批。- true：是；- （默认）false：否。
	InviteNeedConfirm bool     `json:"invite_need_confirm"` // 邀请用户入群时是否需要被邀用户同意。- （默认）true：是；- false：否。
	AffiliationsCount int      `json:"affiliations_count"`  // 聊天室现有成员总数。
	Owner             string   `json:"owner"`               // 群组的管理员。
	Members           []string `json:"members"`             // 群组成员。若传入该参数，数组元素至少一个，不能超过 100。（注：群主 user1 不需要写入到 members 里面）
	Custom            string   `json:"custom"`              // 群组扩展信息，例如可以给群组添加业务相关的标记，不要超过 1,024 字符。
	Mute              bool     `json:"mute"`                // 是否处于全员禁言状态。- true：是；- （默认）false：否。
	Created           int64    `json:"created"`             // 创建该群组的 Unix 时间戳。
	Permission        string   `json:"permission"`          // 群组成员角色：- owner：群主；- admin：管理员；- member：普通成员。
}

type CreateGroupArg struct {
	Name              string   `json:"groupname"`                     // （必填）群组名称，最大长度为 128 字符。如果有空格，则使用 “+” 代替。
	Description       string   `json:"desc"`                          // （必填）群组描述，最大长度为 512 字符。如果有空格，则使用 “+” 代替。
	Public            bool     `json:"public"`                        // （必填）是否是公开群。- true：公开群；- false：私有群。
	MaxUsers          int      `json:"maxusers,omitempty"`            // （选填）群组最大成员数（包括群主），值为数值类型，默认值 200，具体上限请参考 环信即时通讯云控制台。
	AllowInvites      bool     `json:"allowinvites"`                  // （必填）是否允许群成员邀请别人加入此群：- true：允许群成员邀请人加入此群;- （默认）false：只有群主或者管理员才可以往群里加人。 注：该参数仅对私有群有效，因为公开群不允许群成员邀请其他用户入群。
	MembersOnly       bool     `json:"members_only,omitempty"`        // （选填）用户申请入群是否需要群主或者群管理员审批。- true：是；- （默认）false：否。
	InviteNeedConfirm bool     `json:"invite_need_confirm,omitempty"` // （选填）邀请用户入群时是否需要被邀用户同意。- （默认）true：是；- false：否。
	Owner             string   `json:"owner"`                         // （必填）群组的管理员。
	Members           []string `json:"members,omitempty"`             // （选填）群组成员。若传入该参数，数组元素至少一个，不能超过 100。（注：群主 user1 不需要写入到 members 里面）
	Custom            string   `json:"custom,omitempty"`              // （选填）群组扩展信息，例如可以给群组添加业务相关的标记，不要超过 1,024 字符。
}

type createGroupResp struct {
	Data struct {
		ID string `json:"groupid"`
	} `json:"data"`
}

type UpdateGroupArg struct {
	ID           string  `json:"-"`                      // （必填）群组ID
	Name         *string `json:"groupname,omitempty"`    // （选填）群组名称，最大长度为 128 字符。如果有空格，则使用 “+” 代替。
	Description  *string `json:"description,omitempty"`  // （选填）群组描述，最大长度为 512 字符。如果有空格，则使用 “+” 代替。
	MaxUsers     *int    `json:"maxusers,omitempty"`     // （选填）群组最大成员数（包括群主），值为数值类型，默认值 200，具体上限请参考 环信即时通讯云控制台。
	AllowInvites *bool   `json:"allowinvites,omitempty"` // （选填）是否允许群成员邀请别人加入此群：- true：允许群成员邀请人加入此群;- （默认）false：只有群主或者管理员才可以往群里加人。 注：该参数仅对私有群有效，因为公开群不允许群成员邀请其他用户入群。
	MembersOnly  *bool   `json:"membersonly,omitempty"`  // （选填）用户申请入群是否需要群主或者群管理员审批。- true：是；- （默认）false：否。
	Custom       *string `json:"custom,omitempty"`       // （选填）群组扩展信息，例如可以给群组添加业务相关的标记，不要超过 1,024 字符。
}

type UpdateGroupRet struct {
	Name         bool `json:"groupname"`
	Description  bool `json:"description"`
	MaxUsers     bool `json:"maxusers"`
	MembersOnly  bool `json:"membersonly"`
	AllowInvites bool `json:"allowinvites"`
	Custom       bool `json:"custom"`
}

type updateGroupResp struct {
	Data *UpdateGroupRet `json:"data"`
}

type getGroupResp struct {
	Data []*Group `json:"data"`
}

type getAllGroupsResp struct {
	Data []*ListedGroup `json:"data"`
}

type ListedGroup struct {
	ID           string `json:"groupid"`      // 群组 ID。
	Name         string `json:"groupname"`    // 群组名称。
	Type         string `json:"type"`         // “group” 群组类型。
	Owner        string `json:"owner"`        // 群主的 ID。
	Affiliations int    `json:"affiliations"` // 群组现有成员数。
	Created      int64  `json:"created"`      // 群组创建时间，单位为毫秒。
	LastModified int64  `json:"lastModified"` // 最近一次修改的时间戳，单位为毫秒。
	Disabled     bool   `json:"disabled"`     // 群组是否被禁用
}

type FetchGroupsArg struct {
	Limit  int    `json:"limit"`  // （必填）每次期望返回的群组数量。 该参数仅在分页获取时为必需。
	Cursor string `json:"cursor"` // （选填）数据查询的起始位置。该参数仅在分页获取时为必需。
}

type FetchGroupsRet struct {
	List    []*ListedGroup `json:"list"`
	HasMore bool           `json:"has_more"`
	Cursor  string         `json:"cursor"`
}

type fetchGroupsResp struct {
	Data   []*ListedGroup `json:"data"`
	Cursor string         `json:"cursor"`
}

type getAnnouncementResp struct {
	Data struct {
		Announcement string `json:"announcement"`
	} `json:"data"`
}

type updateAnnouncementReq struct {
	Announcement string `json:"announcement"`
}

type getAllShareFilesResp struct {
	Data []*ShareFile `json:"data"`
}

type ShareFile struct {
	FileID    string `json:"file_id"`
	FileName  string `json:"file_name"`
	FileOwner string `json:"file_owner"`
	FileSize  int    `json:"file_size"`
	Created   int64  `json:"created"`
}

type FetchShareFilesArg struct {
	ID       string // （必填）群组ID。
	PageNum  int    // （必填）请求查询的页码。
	PageSize int    // （必填）请求查询每页显示的禁言用户的数量。
}

type FetchShareFilesRet struct {
	List    []*ShareFile `json:"list"`
	HasMore bool         `json:"has_more"`
}

type getShareFileResp struct {
	Data *ShareFile `json:"data"`
}

type fetchShareFilesResp struct {
	Data []*ShareFile `json:"data"`
}

type FetchMembersArg struct {
	ID       string `json:"id"`       // （必填）聊天室ID
	PageNum  int    `json:"pagenum"`  // （选填）请求查询的页码。
	PageSize int    `json:"pagesize"` // （选填）请求查询每页显示的禁言用户的数量。
}

type FetchMembersRet struct {
	List    []string `json:"list"`
	HasMore bool     `json:"has_more"`
}

type fetchMembersResp struct {
	Data []*Affiliation `json:"data"`
}

type Affiliation struct {
	Owner  string `json:"owner"`
	Member string `json:"member"`
}

type addMembersReq struct {
	Usernames []string `json:"usernames"`
}

type addMembersResp struct {
	Data struct {
		Action     string   `json:"action"`
		ID         string   `json:"groupid"`
		NewMembers []string `json:"newmembers"`
	} `json:"data"`
}

type removeMemberResp struct {
	Data *ActionResult `json:"data"`
}

type removeMembersResp struct {
	Data []*ActionResult `json:"data"`
}

type ActionResult struct {
	Result bool   `json:"result"`
	Action string `json:"action"`
	ID     string `json:"groupid"`
	User   string `json:"user"`
	Reason string `json:"reason"`
}

type getAdminResp struct {
	Data []string `json:"data"`
}

type addAdminReq struct {
	NewAdmin string `json:"newadmin"`
}

type transferGroupReq struct {
	NewOwner string `json:"newowner"`
}

type getBlacklistsResp struct {
	Data []string `json:"data"`
}

type addBlacklistsReq struct {
	Usernames []string `json:"usernames"`
}

type addBlacklistsResp struct {
	Data []*ActionResult `json:"data"`
}

type removeBlacklistResp struct {
	Data *ActionResult `json:"data"`
}

type removeBlacklistsResp struct {
	Data []*ActionResult `json:"data"`
}

type getWhitelistsResp struct {
	Data []string `json:"data"`
}

type addWhitelistsReq struct {
	Usernames []string `json:"usernames"`
}

type addWhitelistsResp struct {
	Data []*ActionResult `json:"data"`
}

type removeWhitelistResp struct {
	Data *ActionResult `json:"data"`
}

type removeWhitelistsResp struct {
	Data []*ActionResult `json:"data"`
}

type getMutesResp struct {
	Data []*Mute `json:"data"`
}

type Mute struct {
	Username string `json:"username"`
	Expire   int64  `json:"expire"`
}

type addMutesReq struct {
	MuteDuration int64    `json:"mute_duration"`
	Usernames    []string `json:"usernames"`
}

type addMutesResp struct {
	Data []*AddMuteResult `json:"data"`
}

type AddMuteResult struct {
	Result   bool   `json:"result"`
	Username string `json:"user"`
	Expire   int64  `json:"expire"`
}

type removeMutesResp struct {
	Data []*RemoveMuteResult `json:"data"`
}

type RemoveMuteResult struct {
	Result   bool   `json:"result"`
	Username string `json:"user"`
}

type CreateThreadArg struct {
	GroupID string `json:"group_id"` // （必填）子区所在的群组 ID。
	Name    string `json:"name"`     // （必填）子区的名称，最大长度为 64 字符。
	MsgID   string `json:"msg_id"`   // （必填）子区所在的消息 ID。
	Owner   string `json:"owner"`    // （必填）子区的所有者，即创建子区的成员。
}

type createThreadResp struct {
	Data struct {
		ID string `json:"thread_id"`
	} `json:"data"`
}

type updateThreadReq struct {
	Name string `json:"name"`
}

type FetchThreadsArg struct {
	Sort   string `json:"sort"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type fetchThreadsResp struct {
	Entities []struct {
		ID string `json:"id"`
	} `json:"entities"`
	Properties struct {
		Cursor string `json:"cursor"`
	} `json:"properties"`
}

type FetchThreadsRet struct {
	List    []string `json:"list"`
	HasMore bool     `json:"has_more"`
	Cursor  string   `json:"cursor"`
}

type FetchGroupUserThreadsArg struct {
	GroupID  string `json:"group_id"`
	Username string `json:"username"`
	Sort     string `json:"sort"`
	Limit    int    `json:"limit"`
	Cursor   string `json:"cursor"`
}

type fetchGroupUserThreadsResp struct {
	Entities   []*Thread `json:"entities"`
	Properties struct {
		Cursor string `json:"cursor"`
	} `json:"properties"`
}

type Thread struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	MsgID   string `json:"msgId"`
	GroupID string `json:"groupId"`
	Created int64  `json:"created"`
}

type FetchGroupUserThreadsRet struct {
	List    []*Thread `json:"list"`
	HasMore bool      `json:"has_more"`
	Cursor  string    `json:"cursor"`
}
