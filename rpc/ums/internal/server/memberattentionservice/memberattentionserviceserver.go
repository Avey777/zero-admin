// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/logic/memberattentionservice"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
)

type MemberAttentionServiceServer struct {
	svcCtx *svc.ServiceContext
	umsclient.UnimplementedMemberAttentionServiceServer
}

func NewMemberAttentionServiceServer(svcCtx *svc.ServiceContext) *MemberAttentionServiceServer {
	return &MemberAttentionServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加品牌关注
func (s *MemberAttentionServiceServer) MemberBrandAttentionAdd(ctx context.Context, in *umsclient.MemberBrandAttentionAddReq) (*umsclient.MemberBrandAttentionAddResp, error) {
	l := memberattentionservicelogic.NewMemberBrandAttentionAddLogic(ctx, s.svcCtx)
	return l.MemberBrandAttentionAdd(in)
}

// 取消品牌关注/清空当前用户品牌关注列表
func (s *MemberAttentionServiceServer) MemberBrandAttentionDelete(ctx context.Context, in *umsclient.MemberBrandAttentionDeleteReq) (*umsclient.MemberBrandAttentionDeleteResp, error) {
	l := memberattentionservicelogic.NewMemberBrandAttentionDeleteLogic(ctx, s.svcCtx)
	return l.MemberBrandAttentionDelete(in)
}

// 分页查询当前用户品牌关注列表
func (s *MemberAttentionServiceServer) MemberBrandAttentionList(ctx context.Context, in *umsclient.MemberBrandAttentionListReq) (*umsclient.MemberBrandAttentionListResp, error) {
	l := memberattentionservicelogic.NewMemberBrandAttentionListLogic(ctx, s.svcCtx)
	return l.MemberBrandAttentionList(in)
}
