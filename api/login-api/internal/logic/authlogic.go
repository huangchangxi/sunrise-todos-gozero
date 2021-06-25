package logic

import (
	"context"

	"login-api/internal/svc"
	"login-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) AuthLogic {
	return AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLogic) Auth(req types.AuthReq) (*types.HttpResponse, error) {
	// todo: add your logic here and delete this line

	return &types.HttpResponse{}, nil
}
