package logic

import (
	"context"

	"github.com/dfang/shorturl/api/internal/svc"
	"github.com/dfang/shorturl/api/internal/types"
	"github.com/dfang/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (*types.ExpandResp, error) {
	resp, err := l.svcCtx.Transformer.Expand(l.ctx, &transform.ExpandRequest{
		Shorten: req.Shorten,
	})
	if err != nil {
		return &types.ExpandResp{}, err
	}

	return &types.ExpandResp{
		Url: resp.Url,
	}, nil
}
