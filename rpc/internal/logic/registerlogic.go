package logic

import (
	"context"
	"zeroService/api/model"

	"zeroService/rpc/internal/svc"
	"zeroService/rpc/userService"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *userService.RegisterRequest) (*userService.RegisterResponse, error) {
	user := model.Userinfo{
		Username: in.Username,
		Nickname: in.Nickname,
		Pwd:      in.Pwd,
		Age:      in.Age,
	}
	res, err := l.svcCtx.Model.Insert(user)
	if err != nil {
		return nil, err
	}
	rid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &userService.RegisterResponse{Rid: rid}, nil
}
