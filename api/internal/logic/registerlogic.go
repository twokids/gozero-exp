package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &types.HttpResponse{}, nil
}
