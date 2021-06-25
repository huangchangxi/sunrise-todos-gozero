package logic

import (
	"context"

	"login-api/internal/svc"
	"login-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) IndexLogic {
	return IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (*types.HttpResponse, error) {
	// todo: add your logic here and delete this line

	return &types.HttpResponse{}, nil
}
