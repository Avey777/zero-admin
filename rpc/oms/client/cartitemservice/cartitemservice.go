// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package cartitemservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CartItemAddReq                    = omsclient.CartItemAddReq
	CartItemAddResp                   = omsclient.CartItemAddResp
	CartItemClearReq                  = omsclient.CartItemClearReq
	CartItemClearResp                 = omsclient.CartItemClearResp
	CartItemDeleteReq                 = omsclient.CartItemDeleteReq
	CartItemDeleteResp                = omsclient.CartItemDeleteResp
	CartItemListData                  = omsclient.CartItemListData
	CartItemListReq                   = omsclient.CartItemListReq
	CartItemListResp                  = omsclient.CartItemListResp
	CartItemUpdateReq                 = omsclient.CartItemUpdateReq
	CartItemUpdateResp                = omsclient.CartItemUpdateResp
	CompanyAddressAddReq              = omsclient.CompanyAddressAddReq
	CompanyAddressAddResp             = omsclient.CompanyAddressAddResp
	CompanyAddressDeleteReq           = omsclient.CompanyAddressDeleteReq
	CompanyAddressDeleteResp          = omsclient.CompanyAddressDeleteResp
	CompanyAddressListData            = omsclient.CompanyAddressListData
	CompanyAddressListReq             = omsclient.CompanyAddressListReq
	CompanyAddressListResp            = omsclient.CompanyAddressListResp
	CompanyAddressUpdateReq           = omsclient.CompanyAddressUpdateReq
	CompanyAddressUpdateResp          = omsclient.CompanyAddressUpdateResp
	OrderAddReq                       = omsclient.OrderAddReq
	OrderAddResp                      = omsclient.OrderAddResp
	OrderCancelReq                    = omsclient.OrderCancelReq
	OrderCancelResp                   = omsclient.OrderCancelResp
	OrderConfirmReq                   = omsclient.OrderConfirmReq
	OrderConfirmResp                  = omsclient.OrderConfirmResp
	OrderDeleteByIdReq                = omsclient.OrderDeleteByIdReq
	OrderDeleteReq                    = omsclient.OrderDeleteReq
	OrderDeleteResp                   = omsclient.OrderDeleteResp
	OrderItemAddReq                   = omsclient.OrderItemAddReq
	OrderItemAddResp                  = omsclient.OrderItemAddResp
	OrderItemDeleteReq                = omsclient.OrderItemDeleteReq
	OrderItemDeleteResp               = omsclient.OrderItemDeleteResp
	OrderItemListData                 = omsclient.OrderItemListData
	OrderItemListReq                  = omsclient.OrderItemListReq
	OrderItemListResp                 = omsclient.OrderItemListResp
	OrderItemUpdateReq                = omsclient.OrderItemUpdateReq
	OrderItemUpdateResp               = omsclient.OrderItemUpdateResp
	OrderListByMemberIdReq            = omsclient.OrderListByMemberIdReq
	OrderListByMemberIdResp           = omsclient.OrderListByMemberIdResp
	OrderListData                     = omsclient.OrderListData
	OrderListReq                      = omsclient.OrderListReq
	OrderListResp                     = omsclient.OrderListResp
	OrderOperateHistoryAddReq         = omsclient.OrderOperateHistoryAddReq
	OrderOperateHistoryAddResp        = omsclient.OrderOperateHistoryAddResp
	OrderOperateHistoryDeleteReq      = omsclient.OrderOperateHistoryDeleteReq
	OrderOperateHistoryDeleteResp     = omsclient.OrderOperateHistoryDeleteResp
	OrderOperateHistoryListData       = omsclient.OrderOperateHistoryListData
	OrderOperateHistoryListReq        = omsclient.OrderOperateHistoryListReq
	OrderOperateHistoryListResp       = omsclient.OrderOperateHistoryListResp
	OrderOperateHistoryUpdateReq      = omsclient.OrderOperateHistoryUpdateReq
	OrderOperateHistoryUpdateResp     = omsclient.OrderOperateHistoryUpdateResp
	OrderRefundReq                    = omsclient.OrderRefundReq
	OrderRefundResp                   = omsclient.OrderRefundResp
	OrderReturnApplyAddReq            = omsclient.OrderReturnApplyAddReq
	OrderReturnApplyAddResp           = omsclient.OrderReturnApplyAddResp
	OrderReturnApplyDeleteReq         = omsclient.OrderReturnApplyDeleteReq
	OrderReturnApplyDeleteResp        = omsclient.OrderReturnApplyDeleteResp
	OrderReturnApplyListData          = omsclient.OrderReturnApplyListData
	OrderReturnApplyListReq           = omsclient.OrderReturnApplyListReq
	OrderReturnApplyListResp          = omsclient.OrderReturnApplyListResp
	OrderReturnApplyUpdateReq         = omsclient.OrderReturnApplyUpdateReq
	OrderReturnApplyUpdateResp        = omsclient.OrderReturnApplyUpdateResp
	OrderReturnReasonAddReq           = omsclient.OrderReturnReasonAddReq
	OrderReturnReasonAddResp          = omsclient.OrderReturnReasonAddResp
	OrderReturnReasonDeleteReq        = omsclient.OrderReturnReasonDeleteReq
	OrderReturnReasonDeleteResp       = omsclient.OrderReturnReasonDeleteResp
	OrderReturnReasonListData         = omsclient.OrderReturnReasonListData
	OrderReturnReasonListReq          = omsclient.OrderReturnReasonListReq
	OrderReturnReasonListResp         = omsclient.OrderReturnReasonListResp
	OrderReturnReasonUpdateReq        = omsclient.OrderReturnReasonUpdateReq
	OrderReturnReasonUpdateResp       = omsclient.OrderReturnReasonUpdateResp
	OrderReturnReasonUpdateStatusReq  = omsclient.OrderReturnReasonUpdateStatusReq
	OrderSettingAddReq                = omsclient.OrderSettingAddReq
	OrderSettingAddResp               = omsclient.OrderSettingAddResp
	OrderSettingDeleteReq             = omsclient.OrderSettingDeleteReq
	OrderSettingDeleteResp            = omsclient.OrderSettingDeleteResp
	OrderSettingListData              = omsclient.OrderSettingListData
	OrderSettingListReq               = omsclient.OrderSettingListReq
	OrderSettingListResp              = omsclient.OrderSettingListResp
	OrderSettingUpdateReq             = omsclient.OrderSettingUpdateReq
	OrderSettingUpdateResp            = omsclient.OrderSettingUpdateResp
	OrderUpdateReq                    = omsclient.OrderUpdateReq
	OrderUpdateResp                   = omsclient.OrderUpdateResp
	QueryOrderListReq                 = omsclient.QueryOrderListReq
	ReleaseSkuStockLockData           = omsclient.ReleaseSkuStockLockData
	UpdateOrderStatusByOutTradeNoReq  = omsclient.UpdateOrderStatusByOutTradeNoReq
	UpdateOrderStatusByOutTradeNoResp = omsclient.UpdateOrderStatusByOutTradeNoResp

	CartItemService interface {
		CartItemAdd(ctx context.Context, in *CartItemAddReq, opts ...grpc.CallOption) (*CartItemAddResp, error)
		CartItemList(ctx context.Context, in *CartItemListReq, opts ...grpc.CallOption) (*CartItemListResp, error)
		CartItemUpdate(ctx context.Context, in *CartItemUpdateReq, opts ...grpc.CallOption) (*CartItemUpdateResp, error)
		CartItemUpdateQuantity(ctx context.Context, in *CartItemUpdateReq, opts ...grpc.CallOption) (*CartItemUpdateResp, error)
		CartItemDelete(ctx context.Context, in *CartItemDeleteReq, opts ...grpc.CallOption) (*CartItemDeleteResp, error)
		CartItemClear(ctx context.Context, in *CartItemClearReq, opts ...grpc.CallOption) (*CartItemClearResp, error)
	}

	defaultCartItemService struct {
		cli zrpc.Client
	}
)

func NewCartItemService(cli zrpc.Client) CartItemService {
	return &defaultCartItemService{
		cli: cli,
	}
}

func (m *defaultCartItemService) CartItemAdd(ctx context.Context, in *CartItemAddReq, opts ...grpc.CallOption) (*CartItemAddResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.CartItemAdd(ctx, in, opts...)
}

func (m *defaultCartItemService) CartItemList(ctx context.Context, in *CartItemListReq, opts ...grpc.CallOption) (*CartItemListResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.CartItemList(ctx, in, opts...)
}

func (m *defaultCartItemService) CartItemUpdate(ctx context.Context, in *CartItemUpdateReq, opts ...grpc.CallOption) (*CartItemUpdateResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.CartItemUpdate(ctx, in, opts...)
}

func (m *defaultCartItemService) CartItemUpdateQuantity(ctx context.Context, in *CartItemUpdateReq, opts ...grpc.CallOption) (*CartItemUpdateResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.CartItemUpdateQuantity(ctx, in, opts...)
}

func (m *defaultCartItemService) CartItemDelete(ctx context.Context, in *CartItemDeleteReq, opts ...grpc.CallOption) (*CartItemDeleteResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.CartItemDelete(ctx, in, opts...)
}

func (m *defaultCartItemService) CartItemClear(ctx context.Context, in *CartItemClearReq, opts ...grpc.CallOption) (*CartItemClearResp, error) {
	client := omsclient.NewCartItemServiceClient(m.cli.Conn())
	return client.CartItemClear(ctx, in, opts...)
}
