package logic

import (
	"context"

	"zeroService/api/internal/svc"
	"zeroService/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)


type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserinfoLogic {
	return UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req types.InfoReq) (*types.HttpResponse, error) {
	// todo: add your logic here and delete this line
	userinfo,err:=l.svcCtx.Model.FindOne(int64(req.Rid))
	rsp:=types.HttpResponse{}
	if err!=nil{
		rsp.Code=2
		rsp.Msg="not find data"
		return &rsp,err
	}
	rsp.Data=types.InfoRsp{
		Rid:      int(userinfo.Id),
		UserName: userinfo.Username,
		Pwd:      userinfo.Pwd,
		NickName: userinfo.Nickname,
		Age:      int(userinfo.Age),
	}

	return &rsp, nil
}
