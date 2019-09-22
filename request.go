package gosms

import (
	"bytes"
	"net/url"
	"time"
)

const (
	delimeter   = "&"
	signLinaJie = "GET&%2F&"
)

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

func (this *QueryString) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(this.getAccessKeyId())
	buf.WriteString(this.getAction())
	buf.WriteString(this.getFormat())
	buf.WriteString(this.getOutId())
	buf.WriteString(this.getPhoneNumbers())
	buf.WriteString(this.getRegionId())
	buf.WriteString(this.getSignName())
	buf.WriteString(this.getSignatureMethod())
	buf.WriteString(this.getSignatureNonce())
	buf.WriteString(this.getSignatureVersion())
	buf.WriteString(this.getTemplateCode())
	buf.WriteString(this.getTemplateParam())
	buf.WriteString(this.getTimestamp())
	buf.WriteString(this.getVersion())
	signString := buf.String()
	return signString
}

func (this *QueryString) getAccessKeyId() string {
	return "AccessKeyId=" + this.AccessKeyId
}

func (this *QueryString) getOutId() string {
	return "&OutId=" + this.OutId
}

func (this *QueryString) getFormat() string {
	return "&Format=" + this.Format
}

func (this *QueryString) getAction() string {
	return "&Action=" + this.Action
}

func (this *QueryString) getSignatureNonce() string {
	return "&SignatureNonce=" + this.SignatureNonce
}

func (this *QueryString) getSignatureMethod() string {
	return "&SignatureMethod=" + this.SignatureMethod
}

func (this *QueryString) getSignName() string {
	return "&SignName=" + url.QueryEscape(this.SignName)
}

func (this *QueryString) getRegionId() string {
	return "&RegionId=" + this.RegionId
}

func (this *QueryString) getPhoneNumbers() string {
	return "&PhoneNumbers=" + this.PhoneNumbers
}

func (this *QueryString) getTemplateParam() string {
	return "&TemplateParam=" + url.QueryEscape(this.TemplateParam)
}

func (this *QueryString) getTemplateCode() string {
	return "&TemplateCode=" + this.TemplateCode
}

func (this *QueryString) getSignatureVersion() string {
	return "&SignatureVersion="+this.SignatureVersion
}

func (this *QueryString) getTimestamp() string {
	var cstZone = time.FixedZone("GMT", 0)
	keyTime := time.Now().In(cstZone).Format("2006-01-02T15:04:05Z")
	timstapm := "&Timestamp=" + url.QueryEscape(keyTime)
	return timstapm
}

func (this *QueryString) getVersion() string {
	return "&Version="+this.Version
}
