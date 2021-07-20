package logic

import (
	"context"
	"zeroService/rpc/userService"

	"zeroService/api/internal/svc"
	"zeroService/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (*types.HttpResponse, error) {
	in := &userService.RegisterRequest{
		Username: req.UserName,
		Nickname: req.NickName,
		Pwd:      req.Pwd,
		Age:      int64(req.Age),
	}
	regRsp, err := l.svcCtx.UserServiceRpc.Register(l.ctx, in)
	if err != nil {
		return nil, err
	}
	rsp := types.HttpResponse{}

	rsp.Data = types.RegisterRsp{
		Rid: int(regRsp.Rid),
	}
	return &rsp, nil
}
