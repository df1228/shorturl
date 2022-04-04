package logic

import (
	"context"

	"github.com/dfang/shorturl/rpc/internal/svc"
	"github.com/dfang/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenRequest) (*transform.ShortenResponse, error) {
	// todo: add your logic here and delete this line

	return &transform.ShortenResponse{}, nil
}
