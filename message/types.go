package message

type MsgType string

type sendReq struct {
	From       string   `json:"from"`
	To         []string `json:"to"`
	Type       string   `json:"type"`
	Body       string   `json:"body"`
	SyncDevice bool     `json:"sync_device,omitempty"`
	RouteType  string   `json:"routetype,omitempty"`
	Ext        string   `json:"ext,omitempty"`
}

type sendResp struct {
	Data map[string]string `json:"data"`
}

type MsgTxt struct {
	Msg string `json:"msg"` // 消息内容。
}

type MsgImage struct {
	Filename string `json:"filename"` // 图片名称。
	Secret   string `json:"secret"`   // 图片的访问密钥。成功上传图片后，从 文件上传 的响应 body 中获取的 share-secret。如果图片文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	Width    int    `json:"width"`    // 图片宽度。
	Height   int    `json:"height"`   // 图片高度。
	UUID     string `json:"uuid"`     // 文件的UUID，其中 uuid 为文件 ID，成功上传文件后，从 文件上传 的响应 body 中获取。
}

type msgImageBody struct {
	Filename string `json:"filename"`
	Secret   string `json:"secret"`
	Size     string `json:"size"`
	Url      string `json:"url"`
}

type MsgAudio struct {
	Filename string `json:"filename"` // 语音文件的名称。
	Secret   string `json:"secret"`   // 语音文件访问密钥，成功上传语音文件后，从 文件上传 的响应 body 中获取的 share-secret。 如果语音文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	Length   int    `json:"length"`   // 语音时长，单位为秒。
	UUID     string `json:"uuid"`     // 文件的UUID，其中 uuid 为文件 ID，成功上传文件后，从 文件上传 的响应 body 中获取。
}

type msgAudioBody struct {
	Filename string `json:"filename"`
	Secret   string `json:"secret"`
	Length   int    `json:"length"`
	Url      string `json:"url"`
}

type MsgVideo struct {
	ThumbUUID   string `json:"thumb_uuid"`   // 视频缩略图 URL 地址：https://{host}/{org_name}/{app_name}/chatfiles/{uuid}。uuid 为视频缩略图唯一标识，成功上传缩略图文件后，从 文件上传 的响应 body 中获取。
	ThumbSecret string `json:"thumb_secret"` // 视频缩略图访问密钥，成功上传视频文件后，从 文件上传 的响应 body 中获取的 share-secret。如果缩略图文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	VideoLength int    `json:"video_length"` // 视频时长，单位为秒。
	VideoSecret string `json:"video_secret"` // 视频文件访问密钥，成功上传视频文件后，从 文件上传 的响应 body 中获取的 share-secret。如果视频文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	VideoSize   int64  `json:"video_size"`   // 视频文件大小，单位为字节。
	VideoUUID   string `json:"video_uuid"`   // 文件的UUID，其中 uuid 为文件 ID，成功上传文件后，从 文件上传 的响应 body 中获取。
}

type msgVideoBody struct {
	Thumb       string `json:"thumb"`
	Length      int    `json:"length"`
	Secret      string `json:"secret"`
	FileLength  int64  `json:"file_length"`
	ThumbSecret string `json:"thumb_secret"`
	Url         string `json:"url"`
}

type MsgFile struct {
	Filename string `json:"filename"` // 文件名称。
	Secret   string `json:"secret"`   // 文件访问密钥，成功上传文件后，从 文件上传 的响应 body 中获取的 share-secret。如果文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	UUID     string `json:"uuid"`     // 文件的UUID，其中 uuid 为文件 ID，成功上传文件后，从 文件上传 的响应 body 中获取。
}

type msgFileBody struct {
	Filename string `json:"filename"`
	Secret   string `json:"secret"`
	Url      string `json:"url"`
}

type MsgLocation struct {
	Lat  float64 `json:"lat"`  // 位置的纬度，单位为度。
	Lng  float64 `json:"lng"`  // 位置的经度，单位为度。
	Addr string  `json:"addr"` // 位置的文字描述。
}

type MsgCMD struct {
	Action string `json:"action"` // 命令内容。
}

type MsgCustom struct {
	CustomEvent string            `json:"customEvent"` // 用户自定义的事件类型。该参数的值必须满足正则表达式 [a-zA-Z0-9-_/\.]{1,32}，长度为 1-32 个字符。
	CustomExts  map[string]string `json:"customExts"`  // 用户自定义的事件属性，类型必须是 Map<String,String>，最多可以包含 16 个元素。customExts 是可选的，不需要可以不传。
	From        string            `json:"from"`        // 消息发送者。若未传值，默认值为admin；若传了空字符串 (“”)，请求失败。
	Ext         interface{}       `json:"ext"`         // 扩展属性，支持 app 自定义内容。若传入该字段，值不能为 “ext:null” 这种形式，否则会发生错误。
}

type msgCustomBody struct {
	CustomEvent string `json:"customEvent,omitempty"`
	CustomExts  string `json:"CustomExts,omitempty"`
	From        string `json:"from"`
	Ext         string `json:"ext,omitempty"`
}
