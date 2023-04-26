// Code generated by goctl. DO NOT EDIT!

package pmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	pmsCommentReplayFieldNames          = builder.RawFieldNames(&PmsCommentReplay{})
	pmsCommentReplayRows                = strings.Join(pmsCommentReplayFieldNames, ",")
	pmsCommentReplayRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsCommentReplayFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	pmsCommentReplayRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsCommentReplayFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	pmsCommentReplayModel interface {
		Insert(ctx context.Context, data *PmsCommentReplay) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsCommentReplay, error)
		Update(ctx context.Context, data *PmsCommentReplay) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsCommentReplayModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsCommentReplay struct {
		Id             int64          `db:"id"`
		CommentId      sql.NullInt64  `db:"comment_id"`
		MemberNickName sql.NullString `db:"member_nick_name"`
		MemberIcon     sql.NullString `db:"member_icon"`
		Content        sql.NullString `db:"content"`
		CreateTime     sql.NullTime   `db:"create_time"`
		Type           sql.NullInt64  `db:"type"` // 评论人员类型；0->会员；1->管理员
	}
)

func newPmsCommentReplayModel(conn sqlx.SqlConn) *defaultPmsCommentReplayModel {
	return &defaultPmsCommentReplayModel{
		conn:  conn,
		table: "`pms_comment_replay`",
	}
}

func (m *defaultPmsCommentReplayModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsCommentReplayModel) FindOne(ctx context.Context, id int64) (*PmsCommentReplay, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsCommentReplayRows, m.table)
	var resp PmsCommentReplay
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPmsCommentReplayModel) Insert(ctx context.Context, data *PmsCommentReplay) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, pmsCommentReplayRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CommentId, data.MemberNickName, data.MemberIcon, data.Content, data.Type)
	return ret, err
}

func (m *defaultPmsCommentReplayModel) Update(ctx context.Context, data *PmsCommentReplay) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsCommentReplayRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CommentId, data.MemberNickName, data.MemberIcon, data.Content, data.Type, data.Id)
	return err
}

func (m *defaultPmsCommentReplayModel) tableName() string {
	return m.table
}