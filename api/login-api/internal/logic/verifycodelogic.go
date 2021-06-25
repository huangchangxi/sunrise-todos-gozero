package logic

import (
	"context"

	"login-api/internal/svc"
	"login-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) VerifyCodeLogic {
	return VerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyCodeLogic) VerifyCode() (*types.HttpResponse, error) {
	// todo: add your logic here and delete this line

	return &types.HttpResponse{}, nil
}
