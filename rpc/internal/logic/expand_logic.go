package logic

import (
	"context"

	"github.com/dfang/shorturl/rpc/internal/svc"
	"github.com/dfang/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandRequest) (*transform.ExpandResponse, error) {
	// todo: add your logic here and delete this line

	return &transform.ExpandResponse{}, nil
}
