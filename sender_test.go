package gosms_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gosms"
	"testing"
)

func TestGoSMSSenderSuite(t *testing.T) {
	suite.Run(t, new(GoSMSSenderSuite))
}

type GoSMSSenderSuite struct {
	suite.Suite
	sender gosms.Sender
}

func (suite *GoSMSSenderSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.sender = gosms.NewSender("your_accessKeyId", "your_AccessSecret", "your_SignName")
}

func (suite *GoSMSSenderSuite) TestSend() {
	msgJSON := `{code:"1234"}`
	if response, err := suite.sender.Send("13288888888", msgJSON, "SMS_9999999"); err ==nil {
		suite.Assertions.Equal("OK", response.Code )
	} else {
		suite.Fail("failed to send message with error", err)
	}
}