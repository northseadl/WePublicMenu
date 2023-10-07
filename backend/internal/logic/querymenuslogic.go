package logic

import (
	"WePublicMenu/internal/svc"
	"WePublicMenu/internal/types"
	"context"
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v3/object"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenusLogic {
	return &QueryMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenusLogic) QueryMenus() (resp *types.QueryMenusReply, err error) {
	res, err := l.svcCtx.App.Menu.Get(l.ctx)
	if err != nil {
		return nil, err
	}
	if res.ErrCode != 0 {
		return nil, errors.New(res.Message)
	}

	return &types.QueryMenusReply{
		Button:    res.Menus.Buttons,
		MatchRule: object.HashMap{},
	}, nil
}
