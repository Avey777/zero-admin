package companyaddress

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCompanyAddressDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCompanyAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCompanyAddressDetailLogic {
	return &QueryCompanyAddressDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCompanyAddressDetailLogic) QueryCompanyAddressDetail(req *types.QueryCompanyAddressDetailReq) (resp *types.QueryCompanyAddressDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
