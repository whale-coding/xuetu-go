package handler

import (
	"fmt"
	"log"
	"math/rand"
	"time"
	"xuetu-wx/global"
	"xuetu-wx/utils"
)

const (
	keyword = "验证码"
)

// HandleTextMessage 处理文本消息
func HandleTextMessage(msg *utils.RequestMessage) *utils.ReplyTextMessage {
	if msg.Content == keyword {
		code := generateVerificationCode() // 验证码
		key := fmt.Sprintf("wx:verify:code:%s", msg.FromUserName)
		// 将验证码放到redis中，key为 wx:verify:code:<openid>  微信用户登录的openId， value为验证码code
		if err := global.RedisDB.Set(key, code, 5*time.Minute); err != nil {
			log.Printf("Failed to save verification code to Redis: %v", err)
			// Optionally, return an error message to the user
			return utils.NewReplyTextMessage(msg.FromUserName, msg.ToUserName, "服务器繁忙，请稍后再试")
		}
		// 将验证码返回给微信公众号
		content := fmt.Sprintf("您当前的验证码是：%d！ 5分钟内有效", code)
		return utils.NewReplyTextMessage(msg.FromUserName, msg.ToUserName, content)
	}
	return nil
}

// 生成6位验证码
func generateVerificationCode() int {
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	return src.Intn(900000) + 100000
}
