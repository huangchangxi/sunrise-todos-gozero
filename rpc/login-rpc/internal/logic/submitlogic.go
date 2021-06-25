package logic

import (
	"context"

	"login-rpc/internal/svc"
	"login-rpc/proto"

	"github.com/tal-tech/go-zero/core/logx"
)

type SubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitLogic {
	return &SubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitLogic) Submit(in *proto.Request) (*proto.Response, error) {
	// todo: add your logic here and delete this line

	return &proto.Response{}, nil
}
