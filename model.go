package gosms

import "strings"

type Message struct {
	PhoneNumbers  string
	TemplateParam string
	TemplateCode  string
	SignName      string
	OutId         string ////可选:outId为提供给业务方扩展字段,最终在短信回执消息中将此值带回给调用者
}

type Action struct {
	Action     string
	ApiVersion string
	RegionId   string
	Msg        Message
}

//SimpleMessage ...
type Response struct {
	Message string `json:"Message"`
	Code    string `json:"Code"`
}

type RequestBody struct {
	Reader *strings.Reader
}

