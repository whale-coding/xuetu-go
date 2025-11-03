package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

/**
XML 解析与回复封装
*/

// RequestMessage 表示微信发送的消息
type RequestMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int64    `xml:"MsgId"`
	Event        string   `xml:"Event"`
}

// ReplyTextMessage 表示微信文本回复
type ReplyTextMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

// ParseXML 从 *http.Request 中解析 XML 到结构体
func ParseXML(r *http.Request) (*RequestMessage, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var msg RequestMessage
	if err := xml.Unmarshal(body, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

// NewReplyTextMessage 构造文本回复信息
func NewReplyTextMessage(toUserName, fromUserName, content string) *ReplyTextMessage {
	return &ReplyTextMessage{
		ToUserName:   toUserName,
		FromUserName: fromUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}
}
