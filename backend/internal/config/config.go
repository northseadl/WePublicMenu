package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	PowerXPlugin
	Endpoint string
	Mode     string
}

type PowerXPlugin struct {
	AppId  string
	Secret string
	AESKey string
	OAuth  struct {
		Callback string
		Scopes   []string
	}
	HttpDebug bool
}
