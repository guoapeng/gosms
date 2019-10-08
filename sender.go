package gosms

import (
	satoriuuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
)

type Sender interface {
	Send(phone, msgJSON, templateCode string) (*Response, error)
}

func NewSender(accessKeyId, accessSecret, signName string) Sender {
	server := NewServer("http",
		"dysmsapi.aliyuncs.com", "80")
	credential := Credential{AccessKeyId: accessKeyId, AccessSecret: accessSecret}
	profile := Profile{
		SignatureMethod:  "HMAC-SHA1",
		SignatureVersion: "1.0",
		Credential:       credential,
	}
	signatureBuilder := SignatureBuilder{ Credential: credential}
	return &AliSmsSender{Client: AcsClient{TargetServer: server, HttpClient: &http.Client{}, SignatureBuilder: signatureBuilder},
		Profile: profile,
	    SignNameProvider: SignNameProvider{SignName:signName},
	}
}

type AliSmsSender struct {
	Client  AcsClient
	Profile Profile
	SignNameProvider SignNameProvider
}

type SignNameProvider struct {
	SignName string
}

func (p *AliSmsSender) Send(phone, msgJSON, templateCode string) (*Response, error) {
	action := Action{Action: "SendSms",
		ApiVersion: "2017-05-25",
		RegionId:   "cn-hangzhou",
		Msg: Message{PhoneNumbers: phone,
			TemplateParam: msgJSON,
			TemplateCode:  templateCode,
			SignName:      p.SignNameProvider.SignName,
			OutId:         satoriuuid.NewV4().String()}}
	qs := p.buildQueryString(action)
	return p.Client.GetResponse(qs, p.buildBody("send msg"))
}

func (p *AliSmsSender) buildBody(content string) RequestBody {
	return RequestBody{Reader: strings.NewReader(content)}
}

func (p *AliSmsSender) buildQueryString(action Action) *QueryString {
	return NewRequest(p.Profile, satoriuuid.NewV4().String(), action)
}

