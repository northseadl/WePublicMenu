package svc

import (
	"WePublicMenu/internal/config"
	"WePublicMenu/internal/middleware"
	"WePublicMenu/pkg/powerx"
	"WePublicMenu/pkg/powerx/client"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/zeromicro/go-zero/rest"
	"log/slog"
)

type ServiceContext struct {
	Config           *config.Config
	PluginMiddleware rest.Middleware
	*client.PClient
	PowerX *powerx.PowerX
	App    *officialAccount.OfficialAccount
}

func NewServiceContext(c *config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		PluginMiddleware: middleware.NewPluginMiddleware().Handle,
	}
}

func (c *ServiceContext) Setup() {
	c.PowerX = powerx.NewPowerX(c.Config.Endpoint, c.Config.Mode == "dev")

	// 初始化微信公众号API SDK
	app, err := officialAccount.NewOfficialAccount(&officialAccount.UserConfig{
		AppID:  c.Config.AppId,
		Secret: c.Config.Secret,
		OAuth: officialAccount.OAuth{
			Callback: c.Config.OAuth.Callback,
			Scopes:   c.Config.OAuth.Scopes,
		},
		AESKey:    c.Config.AESKey,
		HttpDebug: true,
	})

	c.App = app

	if err != nil {
		slog.Error(fmt.Sprintf("official account init failed %v", err))
	}
}
