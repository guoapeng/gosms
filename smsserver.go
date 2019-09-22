package gosms

import (
	"errors"
	"fmt"
)

type  SmsServer interface {
	GetUrl() (string, error)
}

func NewServer(protocal, host, port string) SmsServer {
	return &smsServer{protocal: protocal, host: host, port: port}
}

type smsServer struct {
	protocal string
	host     string
	port     string
}

func (server *smsServer) GetUrl() (string, error) {
	if server.protocal == "" {
		return "", errors.New("protocal should be specified")
	} else if server.host == "" {
		return "", errors.New("host should be specified")
	} else {
		if (server.protocal == "http" && server.port == "80") ||
			(server.protocal == "https" && server.port == "443") ||
			server.port == "" {
			return fmt.Sprintf("%s://%s", server.protocal, server.host), nil
		} else {
			return fmt.Sprintf("%s://%s:%s", server.protocal, server.host, server.port), nil
		}
	}
}
