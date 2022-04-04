package logic

import (
	"context"
	"fmt"

	"github.com/dfang/shorturl/rpc/internal/svc"
	"github.com/dfang/shorturl/rpc/model"
	"github.com/dfang/shorturl/rpc/transform"

	"github.com/zeromicro/go-zero/core/hash"
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
	key := hash.Md5Hex([]byte(in.Url))[:6]

	// 数据库有这条记录，不用插入了
	shorturl, err := l.svcCtx.Model.FindOne(l.ctx, key)
	if shorturl != nil && err == nil {
		fmt.Println("数据库有这条记录，不用插入了")
		return &transform.ShortenResponse{
			Shorten: key,
		}, nil
	}

	result, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	})
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	logx.Info("last insert id: ", id)

	return &transform.ShortenResponse{
		Shorten: key,
	}, nil
}
