由于没有注意到ali官方发布到新版的golang sdk for SMS service
快实现完成完成了才发现, 索性把它实现完成,只是没有补充足够的测试.
相比官方sdk, 个人认为实现要比官方版简洁,调用优雅, 可扩展性强于官方版.

description:
golang version short message sender underline with ali sms service

## Installation
Use `go get` to install SDK

```sh
$ go get -u github.com/guoapeng/gosms
```

## Quick Examples
Before you begin, you need to sign up for an Alibaba Cloud account and retrieve your [Credentials](https://usercenter.console.aliyun.com/#/manage/ak).

### Create sender and send short message
```go
package main

import "github.com/guoapeng/gosms"

func main() {

	sender := gosms.NewSender("your_accessKeyId", "your_AccessSecret", "your_SignName")
    
	msgJSON := `{code:"1234"}`           
	if response, err := sender.Send("13288888888", msgJSON, "SMS_9999999"); err ==nil {  // SMS_9999999 is message template id defined in aliyun's sms service
		if "OK" == response.Code {
			log.Println("sent message successfully and get response ",response)
		} else {
			log.Println("sent message with issue ",response)
		}
	} else {
		log.Fatal("failed to send message with error", err)
	}
}
```