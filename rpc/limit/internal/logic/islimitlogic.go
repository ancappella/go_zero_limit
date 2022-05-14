package logic

import (
	"context"

	"go_zero_limit/rpc/limit/internal/svc"
	"go_zero_limit/rpc/limit/limit"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLimitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLimitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLimitLogic {
	return &IsLimitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsLimitLogic) IsLimit(in *limit.LimitReq) (*limit.LimitResp, error) {
	// todo: add your logic here and delete this line

	return &limit.LimitResp{}, nil
}
