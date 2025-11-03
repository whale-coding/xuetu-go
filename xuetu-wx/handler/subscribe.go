package handler

import (
	"xuetu-wx/utils"
)

const (
	subscribeReply = "欢迎关注公众号！🎉"
)

// HandleSubscribeEvent 处理订阅消息
func HandleSubscribeEvent(msg *utils.RequestMessage) *utils.ReplyTextMessage {
	return utils.NewReplyTextMessage(msg.FromUserName, msg.ToUserName, subscribeReply)
}
