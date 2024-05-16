// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package memberlevelservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GrowthChangeHistoryAddReq               = umsclient.GrowthChangeHistoryAddReq
	GrowthChangeHistoryAddResp              = umsclient.GrowthChangeHistoryAddResp
	GrowthChangeHistoryDeleteReq            = umsclient.GrowthChangeHistoryDeleteReq
	GrowthChangeHistoryDeleteResp           = umsclient.GrowthChangeHistoryDeleteResp
	GrowthChangeHistoryListData             = umsclient.GrowthChangeHistoryListData
	GrowthChangeHistoryListReq              = umsclient.GrowthChangeHistoryListReq
	GrowthChangeHistoryListResp             = umsclient.GrowthChangeHistoryListResp
	GrowthChangeHistoryUpdateReq            = umsclient.GrowthChangeHistoryUpdateReq
	GrowthChangeHistoryUpdateResp           = umsclient.GrowthChangeHistoryUpdateResp
	IntegrationChangeHistoryAddReq          = umsclient.IntegrationChangeHistoryAddReq
	IntegrationChangeHistoryAddResp         = umsclient.IntegrationChangeHistoryAddResp
	IntegrationChangeHistoryDeleteReq       = umsclient.IntegrationChangeHistoryDeleteReq
	IntegrationChangeHistoryDeleteResp      = umsclient.IntegrationChangeHistoryDeleteResp
	IntegrationChangeHistoryListData        = umsclient.IntegrationChangeHistoryListData
	IntegrationChangeHistoryListReq         = umsclient.IntegrationChangeHistoryListReq
	IntegrationChangeHistoryListResp        = umsclient.IntegrationChangeHistoryListResp
	IntegrationChangeHistoryUpdateReq       = umsclient.IntegrationChangeHistoryUpdateReq
	IntegrationChangeHistoryUpdateResp      = umsclient.IntegrationChangeHistoryUpdateResp
	IntegrationConsumeSettingAddReq         = umsclient.IntegrationConsumeSettingAddReq
	IntegrationConsumeSettingAddResp        = umsclient.IntegrationConsumeSettingAddResp
	IntegrationConsumeSettingDeleteReq      = umsclient.IntegrationConsumeSettingDeleteReq
	IntegrationConsumeSettingDeleteResp     = umsclient.IntegrationConsumeSettingDeleteResp
	IntegrationConsumeSettingListData       = umsclient.IntegrationConsumeSettingListData
	IntegrationConsumeSettingListReq        = umsclient.IntegrationConsumeSettingListReq
	IntegrationConsumeSettingListResp       = umsclient.IntegrationConsumeSettingListResp
	IntegrationConsumeSettingUpdateReq      = umsclient.IntegrationConsumeSettingUpdateReq
	IntegrationConsumeSettingUpdateResp     = umsclient.IntegrationConsumeSettingUpdateResp
	MemberAddReq                            = umsclient.MemberAddReq
	MemberAddResp                           = umsclient.MemberAddResp
	MemberBrandAttentionAddReq              = umsclient.MemberBrandAttentionAddReq
	MemberBrandAttentionAddResp             = umsclient.MemberBrandAttentionAddResp
	MemberBrandAttentionDeleteReq           = umsclient.MemberBrandAttentionDeleteReq
	MemberBrandAttentionDeleteResp          = umsclient.MemberBrandAttentionDeleteResp
	MemberBrandAttentionListData            = umsclient.MemberBrandAttentionListData
	MemberBrandAttentionListReq             = umsclient.MemberBrandAttentionListReq
	MemberBrandAttentionListResp            = umsclient.MemberBrandAttentionListResp
	MemberByIdReq                           = umsclient.MemberByIdReq
	MemberDeleteReq                         = umsclient.MemberDeleteReq
	MemberDeleteResp                        = umsclient.MemberDeleteResp
	MemberLevelAddReq                       = umsclient.MemberLevelAddReq
	MemberLevelAddResp                      = umsclient.MemberLevelAddResp
	MemberLevelDeleteReq                    = umsclient.MemberLevelDeleteReq
	MemberLevelDeleteResp                   = umsclient.MemberLevelDeleteResp
	MemberLevelListData                     = umsclient.MemberLevelListData
	MemberLevelListReq                      = umsclient.MemberLevelListReq
	MemberLevelListResp                     = umsclient.MemberLevelListResp
	MemberLevelUpdateReq                    = umsclient.MemberLevelUpdateReq
	MemberLevelUpdateResp                   = umsclient.MemberLevelUpdateResp
	MemberListData                          = umsclient.MemberListData
	MemberListReq                           = umsclient.MemberListReq
	MemberListResp                          = umsclient.MemberListResp
	MemberLoginLogAddReq                    = umsclient.MemberLoginLogAddReq
	MemberLoginLogAddResp                   = umsclient.MemberLoginLogAddResp
	MemberLoginLogDeleteReq                 = umsclient.MemberLoginLogDeleteReq
	MemberLoginLogDeleteResp                = umsclient.MemberLoginLogDeleteResp
	MemberLoginLogListData                  = umsclient.MemberLoginLogListData
	MemberLoginLogListReq                   = umsclient.MemberLoginLogListReq
	MemberLoginLogListResp                  = umsclient.MemberLoginLogListResp
	MemberLoginReq                          = umsclient.MemberLoginReq
	MemberLoginResp                         = umsclient.MemberLoginResp
	MemberMemberTagRelationAddReq           = umsclient.MemberMemberTagRelationAddReq
	MemberMemberTagRelationAddResp          = umsclient.MemberMemberTagRelationAddResp
	MemberMemberTagRelationDeleteReq        = umsclient.MemberMemberTagRelationDeleteReq
	MemberMemberTagRelationDeleteResp       = umsclient.MemberMemberTagRelationDeleteResp
	MemberMemberTagRelationListData         = umsclient.MemberMemberTagRelationListData
	MemberMemberTagRelationListReq          = umsclient.MemberMemberTagRelationListReq
	MemberMemberTagRelationListResp         = umsclient.MemberMemberTagRelationListResp
	MemberMemberTagRelationUpdateReq        = umsclient.MemberMemberTagRelationUpdateReq
	MemberMemberTagRelationUpdateResp       = umsclient.MemberMemberTagRelationUpdateResp
	MemberProductCategoryRelationAddReq     = umsclient.MemberProductCategoryRelationAddReq
	MemberProductCategoryRelationAddResp    = umsclient.MemberProductCategoryRelationAddResp
	MemberProductCategoryRelationDeleteReq  = umsclient.MemberProductCategoryRelationDeleteReq
	MemberProductCategoryRelationDeleteResp = umsclient.MemberProductCategoryRelationDeleteResp
	MemberProductCategoryRelationListData   = umsclient.MemberProductCategoryRelationListData
	MemberProductCategoryRelationListReq    = umsclient.MemberProductCategoryRelationListReq
	MemberProductCategoryRelationListResp   = umsclient.MemberProductCategoryRelationListResp
	MemberProductCategoryRelationUpdateReq  = umsclient.MemberProductCategoryRelationUpdateReq
	MemberProductCategoryRelationUpdateResp = umsclient.MemberProductCategoryRelationUpdateResp
	MemberProductCollectionAddReq           = umsclient.MemberProductCollectionAddReq
	MemberProductCollectionAddResp          = umsclient.MemberProductCollectionAddResp
	MemberProductCollectionDeleteReq        = umsclient.MemberProductCollectionDeleteReq
	MemberProductCollectionDeleteResp       = umsclient.MemberProductCollectionDeleteResp
	MemberProductCollectionListData         = umsclient.MemberProductCollectionListData
	MemberProductCollectionListReq          = umsclient.MemberProductCollectionListReq
	MemberProductCollectionListResp         = umsclient.MemberProductCollectionListResp
	MemberReadHistoryAddReq                 = umsclient.MemberReadHistoryAddReq
	MemberReadHistoryAddResp                = umsclient.MemberReadHistoryAddResp
	MemberReadHistoryDeleteReq              = umsclient.MemberReadHistoryDeleteReq
	MemberReadHistoryDeleteResp             = umsclient.MemberReadHistoryDeleteResp
	MemberReadHistoryListData               = umsclient.MemberReadHistoryListData
	MemberReadHistoryListReq                = umsclient.MemberReadHistoryListReq
	MemberReadHistoryListResp               = umsclient.MemberReadHistoryListResp
	MemberReceiveAddressAddReq              = umsclient.MemberReceiveAddressAddReq
	MemberReceiveAddressAddResp             = umsclient.MemberReceiveAddressAddResp
	MemberReceiveAddressDeleteReq           = umsclient.MemberReceiveAddressDeleteReq
	MemberReceiveAddressDeleteResp          = umsclient.MemberReceiveAddressDeleteResp
	MemberReceiveAddressListData            = umsclient.MemberReceiveAddressListData
	MemberReceiveAddressListReq             = umsclient.MemberReceiveAddressListReq
	MemberReceiveAddressListResp            = umsclient.MemberReceiveAddressListResp
	MemberReceiveAddressQueryDetailReq      = umsclient.MemberReceiveAddressQueryDetailReq
	MemberReceiveAddressQueryDetailResp     = umsclient.MemberReceiveAddressQueryDetailResp
	MemberReceiveAddressUpdateReq           = umsclient.MemberReceiveAddressUpdateReq
	MemberReceiveAddressUpdateResp          = umsclient.MemberReceiveAddressUpdateResp
	MemberRuleSettingAddReq                 = umsclient.MemberRuleSettingAddReq
	MemberRuleSettingAddResp                = umsclient.MemberRuleSettingAddResp
	MemberRuleSettingDeleteReq              = umsclient.MemberRuleSettingDeleteReq
	MemberRuleSettingDeleteResp             = umsclient.MemberRuleSettingDeleteResp
	MemberRuleSettingListData               = umsclient.MemberRuleSettingListData
	MemberRuleSettingListReq                = umsclient.MemberRuleSettingListReq
	MemberRuleSettingListResp               = umsclient.MemberRuleSettingListResp
	MemberRuleSettingUpdateReq              = umsclient.MemberRuleSettingUpdateReq
	MemberRuleSettingUpdateResp             = umsclient.MemberRuleSettingUpdateResp
	MemberStatisticsInfoAddReq              = umsclient.MemberStatisticsInfoAddReq
	MemberStatisticsInfoAddResp             = umsclient.MemberStatisticsInfoAddResp
	MemberStatisticsInfoDeleteReq           = umsclient.MemberStatisticsInfoDeleteReq
	MemberStatisticsInfoDeleteResp          = umsclient.MemberStatisticsInfoDeleteResp
	MemberStatisticsInfoListData            = umsclient.MemberStatisticsInfoListData
	MemberStatisticsInfoListReq             = umsclient.MemberStatisticsInfoListReq
	MemberStatisticsInfoListResp            = umsclient.MemberStatisticsInfoListResp
	MemberStatisticsInfoUpdateReq           = umsclient.MemberStatisticsInfoUpdateReq
	MemberStatisticsInfoUpdateResp          = umsclient.MemberStatisticsInfoUpdateResp
	MemberTagAddReq                         = umsclient.MemberTagAddReq
	MemberTagAddResp                        = umsclient.MemberTagAddResp
	MemberTagDeleteReq                      = umsclient.MemberTagDeleteReq
	MemberTagDeleteResp                     = umsclient.MemberTagDeleteResp
	MemberTagListData                       = umsclient.MemberTagListData
	MemberTagListReq                        = umsclient.MemberTagListReq
	MemberTagListResp                       = umsclient.MemberTagListResp
	MemberTagUpdateReq                      = umsclient.MemberTagUpdateReq
	MemberTagUpdateResp                     = umsclient.MemberTagUpdateResp
	MemberTaskAddReq                        = umsclient.MemberTaskAddReq
	MemberTaskAddResp                       = umsclient.MemberTaskAddResp
	MemberTaskDeleteReq                     = umsclient.MemberTaskDeleteReq
	MemberTaskDeleteResp                    = umsclient.MemberTaskDeleteResp
	MemberTaskListData                      = umsclient.MemberTaskListData
	MemberTaskListReq                       = umsclient.MemberTaskListReq
	MemberTaskListResp                      = umsclient.MemberTaskListResp
	MemberTaskUpdateReq                     = umsclient.MemberTaskUpdateReq
	MemberTaskUpdateResp                    = umsclient.MemberTaskUpdateResp
	MemberUpdatePasswordReq                 = umsclient.MemberUpdatePasswordReq
	MemberUpdateReq                         = umsclient.MemberUpdateReq
	MemberUpdateResp                        = umsclient.MemberUpdateResp
	QueryIntegrationConsumeSettingByIdReq   = umsclient.QueryIntegrationConsumeSettingByIdReq
	UpdateMemberIntegrationReq              = umsclient.UpdateMemberIntegrationReq
	UpdateMemberIntegrationResp             = umsclient.UpdateMemberIntegrationResp

	MemberLevelService interface {
		MemberLevelAdd(ctx context.Context, in *MemberLevelAddReq, opts ...grpc.CallOption) (*MemberLevelAddResp, error)
		MemberLevelList(ctx context.Context, in *MemberLevelListReq, opts ...grpc.CallOption) (*MemberLevelListResp, error)
		MemberLevelUpdate(ctx context.Context, in *MemberLevelUpdateReq, opts ...grpc.CallOption) (*MemberLevelUpdateResp, error)
		MemberLevelDelete(ctx context.Context, in *MemberLevelDeleteReq, opts ...grpc.CallOption) (*MemberLevelDeleteResp, error)
	}

	defaultMemberLevelService struct {
		cli zrpc.Client
	}
)

func NewMemberLevelService(cli zrpc.Client) MemberLevelService {
	return &defaultMemberLevelService{
		cli: cli,
	}
}

func (m *defaultMemberLevelService) MemberLevelAdd(ctx context.Context, in *MemberLevelAddReq, opts ...grpc.CallOption) (*MemberLevelAddResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.MemberLevelAdd(ctx, in, opts...)
}

func (m *defaultMemberLevelService) MemberLevelList(ctx context.Context, in *MemberLevelListReq, opts ...grpc.CallOption) (*MemberLevelListResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.MemberLevelList(ctx, in, opts...)
}

func (m *defaultMemberLevelService) MemberLevelUpdate(ctx context.Context, in *MemberLevelUpdateReq, opts ...grpc.CallOption) (*MemberLevelUpdateResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.MemberLevelUpdate(ctx, in, opts...)
}

func (m *defaultMemberLevelService) MemberLevelDelete(ctx context.Context, in *MemberLevelDeleteReq, opts ...grpc.CallOption) (*MemberLevelDeleteResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.MemberLevelDelete(ctx, in, opts...)
}
