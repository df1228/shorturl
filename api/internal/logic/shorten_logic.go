package logic

import (
	"context"

	"github.com/dfang/shorturl/api/internal/svc"
	"github.com/dfang/shorturl/api/internal/types"
	"github.com/dfang/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (*types.ShortenResp, error) {
	resp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transform.ShortenRequest{
		Url: req.Url,
	})
	if err != nil {
		return &types.ShortenResp{}, err
	}

	return &types.ShortenResp{
		Shorten: resp.Shorten,
	}, nil
}
