package message

import (
	"errors"
)

type Target int

const (
	TargetUser     Target = iota // 针对用户
	TargetGroup                  // 针对群组
	TargetChatroom               // 针对聊天室
)

const (
	txt      = "txt"
	image    = "img"
	audio    = "audio"
	video    = "video"
	file     = "file"
	location = "loc"
	cmd      = "cmd"
	custom   = "custom"
)

type Message struct {
	err        error
	target     Target      // 消息目标
	sender     string      // 发送方username
	receivers  []string    // 接收方，接收方则为username
	msgType    string      // 消息类型
	msgBody    interface{} // 消息内容
	syncDevice bool        // 消息发送成功后，是否将消息同步到发送方。
	onlyOnline bool        // 只有接收方在线时，消息才能成功发送
	ext        interface{} // 消息扩展字段
}

func NewMessage(target Target) *Message {
	return &Message{target: target}
}

// AddReceivers 添加接收方
func (m *Message) AddReceivers(receivers ...string) {
	m.receivers = append(m.receivers, receivers...)
}

// SetReceivers 设置接收方
func (m *Message) SetReceivers(receivers ...string) {
	m.receivers = m.receivers[0:0]
	m.receivers = append(m.receivers, receivers...)
}

// GetReceivers 获取接收方
func (m *Message) GetReceivers() []string {
	return m.receivers
}

// SetSender 设置发送方
func (m *Message) SetSender(sender string) {
	m.sender = sender
}

// GetSender 获取发送方
func (m *Message) GetSender() string {
	return m.sender
}

// GetType 获取消息类型
func (m *Message) GetType() string {
	return m.msgType
}

// GetBody 获取消息体
func (m *Message) GetBody() interface{} {
	return m.msgBody
}

// SetBody 设置消息体
func (m *Message) SetBody(body interface{}) {
	switch body.(type) {
	case MsgTxt, *MsgTxt:
		m.msgType = txt
		m.msgBody = body
	case MsgImage, *MsgImage:
		m.msgType = image
	case MsgAudio, *MsgAudio:
		m.msgType = audio
	case MsgVideo, *MsgVideo:
		m.msgType = video
	case MsgFile, *MsgFile:
		m.msgType = file
	case MsgLocation, *MsgLocation:
		m.msgType = location
	case MsgCMD, *MsgCMD:
		m.msgType = cmd
	case MsgCustom, *MsgCustom:
		m.msgType = custom
	default:
		m.err = errors.New("invalid msg")
	}
	m.msgBody = body
}

// GetSyncDevice 获取同步至设备
func (m *Message) GetSyncDevice() bool {
	return m.syncDevice
}

// SetSyncDevice 设置同步至设备
func (m *Message) SetSyncDevice() {
	m.syncDevice = true
}

// SetOnlyOnline 设置只有接收方在线时，消息才能成功发送
func (m *Message) SetOnlyOnline() {
	m.onlyOnline = true
}

// GetOnlyOnline 获取只有接收方在线时，消息才能成功发送
func (m *Message) GetOnlyOnline() bool {
	return m.onlyOnline
}

// SetExt 设置消息扩展字段
func (m *Message) SetExt(ext interface{}) {
	m.ext = ext
}

// GetExt 获取消息扩展字段
func (m *Message) GetExt() interface{} {
	return m.ext
}
