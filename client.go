package gosms

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type AcsClient struct {
	TargetServer SmsServer
	SignatureBuilder SignatureBuilder
	HttpClient   *http.Client
}

func (p *AcsClient) GetResponse(queryString *QueryString, body RequestBody) (*Response, error) {
	if serverAddr, err := p.TargetServer.GetUrl(); err == nil {
		qs := queryString.String()
		signature := p.SignatureBuilder.BuildSignature(qs)
		urlstr := serverAddr + "/?Signature=" + url.QueryEscape(signature) + DELIMITER + qs
		if req, err := http.NewRequest(queryString.MethodType, urlstr, body.Reader); err == nil {
			log.Println("urlstr ", urlstr)
			if resp, err := p.HttpClient.Do(req); err == nil {
				defer resp.Body.Close()
				if response, err := p.processResponse(resp); err == nil {
					return response, nil
				} else {
					log.Println("AliSmsSender processResponse error.", err)
					return nil, err
				}
			} else {
				log.Println("http get error.", err)
				return nil, err
			}
		} else {
			log.Println("http NewRequest error.", err)
			return nil, err
		}
	} else {
		return nil, err
	}
}

type SignatureBuilder struct {
	Credential Credential
}

func (p *SignatureBuilder) BuildSignature(signString string) string {
	singstr := SignPrefix + url.QueryEscape(signString)
	signature := p.hmac4Go(singstr, p.Credential.AccessSecret+DELIMITER)
	return signature
}

func (p *SignatureBuilder) hmac4Go(name, sk string) string {
	mac := hmac.New(sha1.New, []byte(sk))
	mac.Write([]byte(name))
	encodeString := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return encodeString
}


func (p *AcsClient) processResponse(resp *http.Response) (*Response, error) {
	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("http readAll body error ", err)
		return nil, err
	}
	var msg Response
	if len(bodys) > 0 {
		err = json.Unmarshal(bodys, &msg)
		if err != nil {
			log.Println("http unmarshal json body error ", err)
			return nil, err
		}
	}
	return &msg, nil
}
