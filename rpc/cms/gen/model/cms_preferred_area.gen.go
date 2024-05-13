// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCmsPreferredArea = "cms_preferred_area"

// CmsPreferredArea 优选专区
type CmsPreferredArea struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string     `gorm:"column:name;not null" json:"name"`
	SubTitle   string     `gorm:"column:sub_title;not null" json:"sub_title"`
	Pic        string     `gorm:"column:pic;not null;comment:展示图片" json:"pic"` // 展示图片
	Sort       int32      `gorm:"column:sort;not null" json:"sort"`
	ShowStatus int32      `gorm:"column:show_status;not null" json:"show_status"`
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   *string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                                         // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName CmsPreferredArea's table name
func (*CmsPreferredArea) TableName() string {
	return TableNameCmsPreferredArea
}