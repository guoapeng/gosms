package gosms_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"gosms"
	"reflect"
	"testing"
)

func TestSsmServerSuite(t *testing.T) {
	suite.Run(t, new(SsmServerSuite))
}

type SsmServerSuite struct {
	suite.Suite
	server gosms.SmsServer
}

func (suite *SsmServerSuite) SetupTest() {
	mockCtrl := gomock.NewController(suite.T())
	defer mockCtrl.Finish()
	suite.server = gosms.NewServer("", "", "")
}

func (suite *SsmServerSuite) TestGetUrl() {

	for i, unit := range []struct {
		protocal string
		host     string
		port     string
		expected string
		Error    error
	}{
		{"http", "localhost", "80", "http://localhost", nil},
		{"https", "localhost", "443", "https://localhost", nil},
		{"http", "localhost", "8080", "http://localhost:8080", nil},
		{"https", "localhost", "80", "https://localhost:80", nil},
		{protocal: "http", host: "localhost", port: "", expected: "http://localhost"},
		{protocal: "http", host: "localhost", expected: "http://localhost"},
		{protocal: "http", host: "192.168.1.1", expected: "http://192.168.1.1"},
		{protocal: "http",  Error: errors.New("host should be specified")},
		{ host: "192.168.1.1",  Error: errors.New("protocal should be specified")},
	} {
		actually, err := gosms.NewServer(unit.protocal, unit.host, unit.port).GetUrl()
		if !reflect.DeepEqual(err, unit.Error) {
			suite.Fail("error", "Case %d GetUrl() error:\ngot  %v\nwant %v", i, err, unit.Error)
		} else if err== nil && !reflect.DeepEqual(actually, unit.expected) {
			suite.Fail("value", "Case %d  ReadAll() output:\ngot  %q\nwant %q", i, actually, unit.expected)
		}
	}
}
