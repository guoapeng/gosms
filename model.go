package gosms

import (
	"bytes"
	"net/url"
	"strings"
	"time"
)

const (
	DELIMITER  = "&"
	SignPrefix = "GET&%2F&"
)


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

type Credential struct {
	AccessKeyId  string
	AccessSecret string
}

type Profile struct {
	SignatureMethod string
	SignatureVersion string
	Credential Credential
}


type QueryString struct {
	MethodType       string
	AccessKeyId      string
	Action           string
	Format           string //Optional 没传默认为JSON，可选填值：XML
	OutId            string //外部流水扩展字段
	PhoneNumbers     string
	SignName         string
	RegionId         string
	SignatureMethod  string
	SignatureNonce   string  //用于请求的防重放攻击，每次请求唯一
	SignatureVersion string
	TemplateCode     string
	TemplateParam    string
	Timestamp        string
	Version          string
}

func NewRequest( profile Profile, signatureNonce string, action Action) *QueryString {
	return &QueryString{PhoneNumbers: action.Msg.PhoneNumbers,
		MethodType: "GET",
		AccessKeyId:      profile.Credential.AccessKeyId,
		OutId:            action.Msg.OutId,
		Format:           "JSON",
		Action:           action.Action,
		SignatureNonce:   signatureNonce,
		SignatureMethod:  profile.SignatureMethod,
		SignatureVersion: profile.SignatureVersion,
		SignName:         action.Msg.SignName,
		RegionId:         action.RegionId,
		TemplateCode:     action.Msg.TemplateCode,
		TemplateParam:    action.Msg.TemplateParam,
		Version: action.ApiVersion}
}

func (p *QueryString) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(p.getAccessKeyId())
	buf.WriteString(p.getAction())
	buf.WriteString(p.getFormat())
	buf.WriteString(p.getOutId())
	buf.WriteString(p.getPhoneNumbers())
	buf.WriteString(p.getRegionId())
	buf.WriteString(p.getSignName())
	buf.WriteString(p.getSignatureMethod())
	buf.WriteString(p.getSignatureNonce())
	buf.WriteString(p.getSignatureVersion())
	buf.WriteString(p.getTemplateCode())
	buf.WriteString(p.getTemplateParam())
	buf.WriteString(p.getTimestamp())
	buf.WriteString(p.getVersion())
	signString := buf.String()
	return signString
}

func (p *QueryString) getAccessKeyId() string {
	return "AccessKeyId=" + p.AccessKeyId
}

func (p *QueryString) getOutId() string {
	return "&OutId=" + p.OutId
}

func (p *QueryString) getFormat() string {
	return "&Format=" + p.Format
}

func (p *QueryString) getAction() string {
	return "&Action=" + p.Action
}

func (p *QueryString) getSignatureNonce() string {
	return "&SignatureNonce=" + p.SignatureNonce
}

func (p *QueryString) getSignatureMethod() string {
	return "&SignatureMethod=" + p.SignatureMethod
}

func (p *QueryString) getSignName() string {
	return "&SignName=" + url.QueryEscape(p.SignName)
}

func (p *QueryString) getRegionId() string {
	return "&RegionId=" + p.RegionId
}

func (p *QueryString) getPhoneNumbers() string {
	return "&PhoneNumbers=" + p.PhoneNumbers
}

func (p *QueryString) getTemplateParam() string {
	return "&TemplateParam=" + url.QueryEscape(p.TemplateParam)
}

func (p *QueryString) getTemplateCode() string {
	return "&TemplateCode=" + p.TemplateCode
}

func (p *QueryString) getSignatureVersion() string {
	return "&SignatureVersion="+ p.SignatureVersion
}

func (p *QueryString) getTimestamp() string {
	var cstZone = time.FixedZone("GMT", 0)
	keyTime := time.Now().In(cstZone).Format("2006-01-02T15:04:05Z")
	timstapm := "&Timestamp=" + url.QueryEscape(keyTime)
	return timstapm
}

func (p *QueryString) getVersion() string {
	return "&Version="+ p.Version
}
