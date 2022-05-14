package logic

import (
	"context"

	"go_zero_limit/api/internal/svc"
	"go_zero_limit/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLimitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsLimitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLimitLogic {
	return &IsLimitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsLimitLogic) IsLimit(req *types.IsLimitReq) (resp *types.IsLimitResp, err error) {
	// todo: add your logic here and delete this line

	return
}
