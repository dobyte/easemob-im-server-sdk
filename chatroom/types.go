package chatroom

type (
	addSuperAdminReq struct {
		SuperAdmin string `json:"superadmin"`
	}

	addSuperAdminResp struct {
		Data struct {
			Result string `json:"result"` // 添加是否成功
		} `json:"data"`
	}

	FetchSuperAdminsArg struct {
		PageNum  int `json:"pagenum"`  // （必填）当前页码，默认值为 1。
		PageSize int `json:"pagesize"` // （必填）每页返回的超级管理员数量，默认值为 10。
	}

	FetchSuperAdminsRet struct {
		List    []string `json:"list"`
		HasMore bool     `json:"has_more"`
	}

	fetchSuperAdminsResp struct {
		Data  []string `json:"data"`
		Count int      `json:"count"`
	}

	getAllChatroomsResp struct {
		Data []*ListedChatroom `json:"data"`
	}

	getChatroomsResp struct {
		Data []*Chatroom `json:"data"`
	}

	ListedChatroom struct {
		ID                string `json:"id"`                 // 聊天室 ID，聊天室唯一标识，由环信即时通讯 IM 服务器生成。
		Name              string `json:"name"`               // 聊天室名称，最大长度为 128 字符。
		Owner             string `json:"owner"`              // 聊天室的管理员。
		AffiliationsCount int    `json:"affiliations_count"` // 聊天室现有成员总数。
	}

	Chatroom struct {
		ID                string        `json:"id"`                 // 聊天室 ID，聊天室唯一标识，由环信即时通讯 IM 服务器生成。
		Name              string        `json:"name"`               // 聊天室名称，最大长度为 128 字符。
		Description       string        `json:"description"`        // 聊天室描述，最大长度为 512 字符。
		MaxUsers          int           `json:"maxusers"`           // 聊天室成员最大数（包括聊天室所有者），值为数值类型，默认可设置的最大人数为 10,000，如需调整请联系商务。
		Owner             string        `json:"owner"`              // 聊天室的管理员。
		Members           []string      `json:"members"`            // 聊天室成员。若传该参数，数组元素至少一个。
		Custom            string        `json:"custom"`             // 聊天室自定义属性，例如可以给聊天室添加业务相关的标记，不要超过 1,024 字符。
		AffiliationsCount int           `json:"affiliations_count"` // 聊天室现有成员总数。
		MembersOnly       bool          `json:"membersonly"`        // 加入聊天室是否需要群主或者群管理员审批
		AllowInvites      bool          `json:"allowinvites"`       // 是否允许聊天室成员邀请其他用户加入该聊天室
		Created           int64         `json:"created"`            // 创建聊天室时间，Unix 时间戳，单位为毫秒。
		Affiliations      []Affiliation `json:"affiliations"`       // 现有成员列表，包含聊天室所有者和成员。例如：“affiliations”:[{“owner”: “user1”},{“member”:”user2”},{“member”:”user3”}]。
	}

	Affiliation struct {
		Owner  string `json:"owner"`
		Member string `json:"member"`
	}

	CreateChatRoomArg struct {
		Name        string   `json:"name"`        // （必填）聊天室名称，最大长度为 128 字符。
		Description string   `json:"description"` // （必填）聊天室描述，最大长度为 512 字符。
		MaxUsers    int      `json:"maxusers"`    // （选填）聊天室成员最大数（包括聊天室所有者），值为数值类型，默认可设置的最大人数为 10,000，如需调整请联系商务。
		Owner       string   `json:"owner"`       // （必填）聊天室的管理员。
		Members     []string `json:"members"`     // （选填）聊天室成员。若传该参数，数组元素至少一个。
		Custom      string   `json:"custom"`      // （选填）聊天室自定义属性，例如可以给聊天室添加业务相关的标记，不要超过 1,024 字符。
	}

	createChatRoomResp struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	UpdateChatroomArg struct {
		ID          string `json:"-"`                     // （必填）聊天室ID
		Name        string `json:"name,omitempty"`        // （必填）聊天室名称，最大长度为 128 字符。
		Description string `json:"description,omitempty"` // （必填）聊天室描述，最大长度为 512 字符。
		MaxUsers    int    `json:"maxusers,omitempty"`    // （选填）聊天室成员最大数（包括聊天室所有者），值为数值类型，默认可设置的最大人数为 10,000，如需调整请联系商务。
	}

	UpdateChatroomRet struct {
		Name        bool `json:"groupname"`
		Description bool `json:"description"`
		MaxUsers    bool `json:"maxusers"`
	}

	deleteChatroomResp struct {
		Data struct {
			ID      string `json:"id"`
			Success bool   `json:"success"`
		} `json:"data"`
	}

	getAnnouncementResp struct {
		Data struct {
			Announcement string `json:"announcement"`
		} `json:"data"`
	}

	updateAnnouncementReq struct {
		Announcement string `json:"announcement"`
	}

	FetchMembersArg struct {
		ID       string `json:"id"`       // （必填）聊天室ID
		PageNum  int    `json:"pagenum"`  // （选填）请求查询的页码。
		PageSize int    `json:"pagesize"` // （选填）请求查询每页显示的禁言用户的数量。
	}

	FetchMembersRet struct {
		List    []string `json:"list"`
		HasMore bool     `json:"has_more"`
	}

	fetchMembersResp struct {
		Data []*Affiliation `json:"data"`
	}

	addMemberResp struct {
		Data struct {
			Result bool   `json:"result"`
			Action string `json:"action"`
			ID     string `json:"id"`
			User   string `json:"user"`
		} `json:"data"`
	}

	addMembersReq struct {
		Usernames []string `json:"usernames"`
	}

	addMembersResp struct {
		Data struct {
			Action     string   `json:"action"`
			ID         string   `json:"id"`
			NewMembers []string `json:"newmembers"`
		} `json:"data"`
	}

	removeMemberResp struct {
		Data *ActionResult `json:"data"`
	}

	removeMembersResp struct {
		Data []*ActionResult `json:"data"`
	}

	ActionResult struct {
		Result bool   `json:"result"`
		Action string `json:"action"`
		ID     string `json:"id"`
		User   string `json:"user"`
		Reason string `json:"reason"`
	}

	getAdminResp struct {
		Data []string `json:"data"`
	}

	addAdminReq struct {
		NewAdmin string `json:"newadmin"`
	}

	addAdminResp struct {
		Data struct {
			Result string `json:"result"`
		} `json:"data"`
	}

	removeAdminResp struct {
		Data struct {
			Result string `json:"result"`
		} `json:"data"`
	}

	getBlacklistsResp struct {
		Data []string `json:"data"`
	}

	addBlacklistResp struct {
		Data struct {
			Result bool `json:"result"`
		} `json:"data"`
	}

	addBlacklistsReq struct {
		Usernames []string `json:"usernames"`
	}

	addBlacklistsResp struct {
		Data []*ActionResult `json:"data"`
	}

	removeBlacklistResp struct {
		Data *ActionResult `json:"data"`
	}

	removeBlacklistsResp struct {
		Data []*ActionResult `json:"data"`
	}

	getWhitelistsResp struct {
		Data []string `json:"data"`
	}

	addWhitelistsReq struct {
		Usernames []string `json:"usernames"`
	}

	addWhitelistsResp struct {
		Data []*ActionResult `json:"data"`
	}

	removeWhitelistResp struct {
		Data *ActionResult `json:"data"`
	}

	removeWhitelistsResp struct {
		Data []*ActionResult `json:"data"`
	}

	Mute struct {
		Username string `json:"username"`
		Expire   int64  `json:"expire"`
	}

	getMutesResp struct {
		Data []*Mute `json:"data"`
	}

	addMutesReq struct {
		MuteDuration int64    `json:"mute_duration"`
		Usernames    []string `json:"usernames"`
	}

	addMutesResp struct {
		Data []*AddMuteResult `json:"data"`
	}

	AddMuteResult struct {
		Result   bool   `json:"result"`
		Username string `json:"user"`
		Expire   int64  `json:"expire"`
	}

	removeMutesResp struct {
		Data []*RemoveMuteResult `json:"data"`
	}

	RemoveMuteResult struct {
		Result   bool   `json:"result"`
		Username string `json:"user"`
	}
)
