package message

import (
	"encoding/json"
	"fmt"
	"github.com/dobyte/easemob-im-server-sdk/internal/core"
	"reflect"
)

const (
	fileUrlFormat      = "%s/chatfiles/%s"
	sendPrivateMsgUri  = "/messages/users"
	sendGroupMsgUri    = "/messages/chatgroups"
	sendChatroomMsgUri = "/messages/chatrooms"
)

type API interface {
}

type api struct {
	client core.Client
}

func NewAPI(client core.Client) API {
	return &api{client: client}
}

// Send 发送单聊消息
func (a *api) Send(msg *Message) (map[string]string, error) {
	if msg.err != nil {
		return nil, msg.err
	}

	var (
		buf []byte
		err error
	)
	switch msg.msgType {
	case txt:
		buf, err = toTxtBody(msg.msgBody.(*MsgTxt))
	case image:
		buf, err = toImageBody(a.client.BaseUrl(), msg.msgBody.(*MsgImage))
	case audio:
		buf, err = toAudioBody(a.client.BaseUrl(), msg.msgBody.(*MsgAudio))
	case video:
		buf, err = toVideoBody(a.client.BaseUrl(), msg.msgBody.(*MsgVideo))
	case file:
		buf, err = toFileBody(a.client.BaseUrl(), msg.msgBody.(*MsgFile))
	case location:
		buf, err = toLocationBody(msg.msgBody.(*MsgLocation))
	case cmd:
		buf, err = toCMDBody(msg.msgBody.(*MsgCMD))
	case custom:
		buf, err = toCustomBody(msg.msgBody.(*MsgCustom))
	}
	if err != nil {
		return nil, err
	}

	req := &sendReq{
		From:       msg.sender,
		To:         msg.receivers,
		Type:       msg.msgType,
		Body:       string(buf),
		SyncDevice: msg.syncDevice,
	}
	resp := &sendResp{}

	if msg.onlyOnline {
		req.RouteType = "ROUTE_ONLINE"
	}

	if msg.ext != nil && !reflect.ValueOf(msg.ext).IsNil() {
		buf, err := json.Marshal(msg.ext)
		if err != nil {
			return nil, err
		}
		req.Ext = string(buf)
	}

	var uri string
	switch msg.target {
	case TargetUser:
		uri = sendPrivateMsgUri
	case TargetGroup:
		uri = sendGroupMsgUri
	case TargetChatroom:
		uri = sendChatroomMsgUri
	}

	if err = a.client.Post(uri, req, resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func toTxtBody(msg *MsgTxt) ([]byte, error) {
	return json.Marshal(msg)
}

func toImageBody(baseUrl string, msg *MsgImage) ([]byte, error) {
	size, err := json.Marshal(struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{
		Width:  msg.Width,
		Height: msg.Height,
	})
	if err != nil {
		return nil, err
	}

	return json.Marshal(&msgImageBody{
		Filename: msg.Filename,
		Secret:   msg.Secret,
		Size:     string(size),
		Url:      fmt.Sprintf(fileUrlFormat, baseUrl, msg.UUID),
	})
}

func toAudioBody(baseUrl string, msg *MsgAudio) ([]byte, error) {
	return json.Marshal(&msgAudioBody{
		Filename: msg.Filename,
		Secret:   msg.Secret,
		Length:   msg.Length,
		Url:      fmt.Sprintf(fileUrlFormat, baseUrl, msg.UUID),
	})
}

func toVideoBody(baseUrl string, msg *MsgVideo) ([]byte, error) {
	return json.Marshal(&msgVideoBody{
		Thumb:       fmt.Sprintf(fileUrlFormat, baseUrl, msg.ThumbUUID),
		Secret:      msg.VideoSecret,
		Length:      msg.VideoLength,
		FileLength:  msg.VideoSize,
		ThumbSecret: msg.ThumbSecret,
		Url:         fmt.Sprintf(fileUrlFormat, baseUrl, msg.VideoUUID),
	})
}

func toFileBody(baseUrl string, msg *MsgFile) ([]byte, error) {
	return json.Marshal(&msgFileBody{
		Filename: msg.Filename,
		Secret:   msg.Secret,
		Url:      fmt.Sprintf(fileUrlFormat, baseUrl, msg.UUID),
	})
}

func toLocationBody(msg *MsgLocation) ([]byte, error) {
	return json.Marshal(msg)
}

func toCMDBody(msg *MsgCMD) ([]byte, error) {
	return json.Marshal(msg)
}

func toCustomBody(msg *MsgCustom) ([]byte, error) {
	var (
		buf []byte
		ext []byte
		err error
	)

	if len(msg.CustomExts) > 0 {
		if buf, err = json.Marshal(msg.CustomExts); err != nil {
			return nil, err
		}
	}

	if msg.Ext != nil && !reflect.ValueOf(msg.Ext).IsNil() {
		if ext, err = json.Marshal(msg.Ext); err != nil {
			return nil, err
		}
	}

	return json.Marshal(&msgCustomBody{
		CustomEvent: msg.CustomEvent,
		CustomExts:  string(buf),
		From:        msg.From,
		Ext:         string(ext),
	})
}
