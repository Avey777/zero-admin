// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePmsProduct = "pms_product"

// PmsProduct 商品信息
type PmsProduct struct {
	ID                         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	BrandID                    int64     `gorm:"column:brand_id;not null;comment:品牌id" json:"brand_id"`                                                                      // 品牌id
	ProductCategoryID          int64     `gorm:"column:product_category_id;not null;comment:商品分类id" json:"product_category_id"`                                              // 商品分类id
	FeightTemplateID           int64     `gorm:"column:feight_template_id;not null;comment:商品运费模板id" json:"feight_template_id"`                                              // 商品运费模板id
	ProductAttributeCategoryID int64     `gorm:"column:product_attribute_category_id;not null;comment:商品属性分类id" json:"product_attribute_category_id"`                        // 商品属性分类id
	Name                       string    `gorm:"column:name;not null;comment:商品名称" json:"name"`                                                                              // 商品名称
	Pic                        string    `gorm:"column:pic;not null;comment:商品图片" json:"pic"`                                                                                // 商品图片
	ProductSn                  string    `gorm:"column:product_sn;not null;comment:货号" json:"product_sn"`                                                                    // 货号
	DeleteStatus               int32     `gorm:"column:delete_status;not null;comment:删除状态：0->未删除；1->已删除" json:"delete_status"`                                              // 删除状态：0->未删除；1->已删除
	PublishStatus              int32     `gorm:"column:publish_status;not null;comment:上架状态：0->下架；1->上架" json:"publish_status"`                                              // 上架状态：0->下架；1->上架
	NewStatus                  int32     `gorm:"column:new_status;not null;comment:新品状态:0->不是新品；1->新品" json:"new_status"`                                                    // 新品状态:0->不是新品；1->新品
	RecommandStatus            int32     `gorm:"column:recommand_status;not null;comment:推荐状态；0->不推荐；1->推荐" json:"recommand_status"`                                         // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int32     `gorm:"column:verify_status;not null;comment:审核状态：0->未审核；1->审核通过" json:"verify_status"`                                             // 审核状态：0->未审核；1->审核通过
	Sort                       int32     `gorm:"column:sort;not null;comment:排序" json:"sort"`                                                                                // 排序
	Sale                       int32     `gorm:"column:sale;not null;comment:销量" json:"sale"`                                                                                // 销量
	Price                      float64   `gorm:"column:price;not null;comment:商品价格" json:"price"`                                                                            // 商品价格
	PromotionPrice             float64   `gorm:"column:promotion_price;not null;comment:促销价格" json:"promotion_price"`                                                        // 促销价格
	GiftGrowth                 int32     `gorm:"column:gift_growth;not null;comment:赠送的成长值" json:"gift_growth"`                                                              // 赠送的成长值
	GiftPoint                  int32     `gorm:"column:gift_point;not null;comment:赠送的积分" json:"gift_point"`                                                                 // 赠送的积分
	UsePointLimit              int32     `gorm:"column:use_point_limit;not null;comment:限制使用的积分数" json:"use_point_limit"`                                                    // 限制使用的积分数
	SubTitle                   string    `gorm:"column:sub_title;not null;comment:副标题" json:"sub_title"`                                                                     // 副标题
	Description                string    `gorm:"column:description;not null;comment:商品描述" json:"description"`                                                                // 商品描述
	OriginalPrice              float64   `gorm:"column:original_price;not null;comment:市场价" json:"original_price"`                                                           // 市场价
	Stock                      int32     `gorm:"column:stock;not null;comment:库存" json:"stock"`                                                                              // 库存
	LowStock                   int32     `gorm:"column:low_stock;not null;comment:库存预警值" json:"low_stock"`                                                                   // 库存预警值
	Unit                       string    `gorm:"column:unit;not null;comment:单位" json:"unit"`                                                                                // 单位
	Weight                     float64   `gorm:"column:weight;not null;comment:商品重量，默认为克" json:"weight"`                                                                     // 商品重量，默认为克
	PreviewStatus              int32     `gorm:"column:preview_status;not null;comment:是否为预告商品：0->不是；1->是" json:"preview_status"`                                            // 是否为预告商品：0->不是；1->是
	ServiceIds                 string    `gorm:"column:service_ids;not null;comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮" json:"service_ids"`                                  // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string    `gorm:"column:keywords;not null;comment:搜索关键字" json:"keywords"`                                                                     // 搜索关键字
	Note                       string    `gorm:"column:note;not null;comment:备注" json:"note"`                                                                                // 备注
	AlbumPics                  string    `gorm:"column:album_pics;not null;comment:画册图片，连产品图片限制为5张，以逗号分割" json:"album_pics"`                                                 // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string    `gorm:"column:detail_title;not null;comment:详情标题" json:"detail_title"`                                                              // 详情标题
	DetailDesc                 string    `gorm:"column:detail_desc;not null;comment:详情描述" json:"detail_desc"`                                                                // 详情描述
	DetailHTML                 string    `gorm:"column:detail_html;not null;comment:产品详情网页内容" json:"detail_html"`                                                            // 产品详情网页内容
	DetailMobileHTML           string    `gorm:"column:detail_mobile_html;not null;comment:移动端网页详情" json:"detail_mobile_html"`                                               // 移动端网页详情
	PromotionStartTime         time.Time `gorm:"column:promotion_start_time;not null;comment:促销开始时间" json:"promotion_start_time"`                                            // 促销开始时间
	PromotionEndTime           time.Time `gorm:"column:promotion_end_time;not null;comment:促销结束时间" json:"promotion_end_time"`                                                // 促销结束时间
	PromotionPerLimit          int32     `gorm:"column:promotion_per_limit;not null;comment:活动限购数量" json:"promotion_per_limit"`                                              // 活动限购数量
	PromotionType              int32     `gorm:"column:promotion_type;not null;comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购" json:"promotion_type"` // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string    `gorm:"column:brand_name;not null;comment:品牌名称" json:"brand_name"`                                                                  // 品牌名称
	ProductCategoryName        string    `gorm:"column:product_category_name;not null;comment:商品分类名称" json:"product_category_name"`                                          // 商品分类名称
	ProductCategoryIDArray     string    `gorm:"column:product_category_id_Array;not null;comment:商品分类id字符串" json:"product_category_id_Array"`                               // 商品分类id字符串
}

// TableName PmsProduct's table name
func (*PmsProduct) TableName() string {
	return TableNamePmsProduct
}