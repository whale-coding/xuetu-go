package controller

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"xuetu-wx/handler"

	"xuetu-wx/utils"

	"github.com/gin-gonic/gin"
)

// Token Token常量: 微信公众号的token(需要与微信公众号里面保持一致)
const Token = "adwidhaidwoaid"

// VerifySignature 微信回调消息校验 GET请求
func VerifySignature(c *gin.Context) {
	// 获取查询路径参数
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	log.Printf("signature=%s timestamp=%s nonce=%s echostr=%s", signature, timestamp, nonce, echostr)

	// 校验微信加密签名
	// 验证成功返回echostr，否则返回unknown
	if utils.CheckSignature(signature, timestamp, nonce, Token) {
		c.String(200, echostr)
		return
	} else {
		c.String(200, "unknown")
	}
}

// CallbackHandler 处理微信的回调 POST请求
func CallbackHandler(c *gin.Context) {
	// 读原始 XML
	body, _ := c.GetRawData()
	log.Printf("接收到微信消息：%s\n", string(body))

	// 调用封装的 ParseXML 方法，解析XML消息
	msg, err := utils.ParseXML(c.Request)

	if err != nil {
		log.Println("解析微信消息失败:", err)
		c.String(http.StatusOK, "success") // 微信要求返回 success 表示已接收
		return
	}
	log.Printf("收到微信消息: 类型=%s, 内容=%s, 事件=%s\n", msg.MsgType, msg.Content, msg.Event)

	// 需要回复的消息
	var replyMsg *utils.ReplyTextMessage

	// 根据消息类型处理
	switch msg.MsgType {
	case "text":
		// 仅处理文本消息
		replyMsg = handler.HandleTextMessage(msg)
	case "event":
		// 处理订阅事件 msg.MsgType == "event" && msg.Event == "subscribe"
		if msg.Event == "subscribe" {
			replyMsg = handler.HandleSubscribeEvent(msg)
		}
	}

	// 设置请求头等信息
	if replyMsg != nil {
		c.Header("Content-Type", "application/xml; charset=utf-8")

		if err := xml.NewEncoder(c.Writer).Encode(replyMsg); err != nil {
			log.Printf("Error encoding XML: %v", err)
		}
	} else {
		fmt.Fprint(c.Writer, "unknown")
	}
}
