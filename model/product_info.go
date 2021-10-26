package model

import (
	"database/sql"
)

// ProductInfo struct is a row record of the product_info table in the yx database
type ProductInfo struct {
	ProductID            uint64        `gorm:"primary_key;AUTO_INCREMENT;column:product_id;type:ubigint;" json:"product_id"`      // 央选商品信息表
	MerchantID           int64         `gorm:"column:merchant_id;type:bigint;default:0;" json:"merchant_id"`                      // 商家id
	ShopID               uint64        `gorm:"column:shop_id;type:ubigint;default:0;" json:"shop_id"`                             // 店铺id
	BrandID              uint32        `gorm:"column:brand_id;type:uint;default:0;" json:"brand_id"`                              // 品牌id 0：没有品牌
	CateRootID           int32         `gorm:"column:cate_root_id;type:int;default:0;" json:"cate_root_id"`                       // 根类目
	CatePid              int32         `gorm:"column:cate_pid;type:int;default:0;" json:"cate_pid"`                               // 上一级类目
	CateThird            int32         `gorm:"column:cate_third;type:int;default:0;" json:"cate_third"`                           // 三级类目
	CateFourth           int32         `gorm:"column:cate_fourth;type:int;default:0;" json:"cate_fourth"`                         // 四级类目
	CateID               uint32        `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"`                                // 叶子节点类目
	PlatformID           int32         `gorm:"column:platform_id;type:smallint;default:0;" json:"platform_id"`                    // 所属平台 1：抖音 2：快手
	ThirdProductID       int64         `gorm:"column:third_product_id;type:bigint;default:0;" json:"third_product_id"`            // 三方商品id
	ProductName          string        `gorm:"column:product_name;type:varchar;size:200;" json:"product_name"`                    // 商品名称
	ProductPic           string        `gorm:"column:product_pic;type:varchar;size:1000;" json:"product_pic"`                     // 商品主图
	ProductDesc          string        `gorm:"column:product_desc;type:text;size:65535;" json:"product_desc"`                     // 商品详情
	ProductPrice         int32         `gorm:"column:product_price;type:int;default:0;" json:"product_price"`                     // 商品原价 扩大1000
	SalePrice            int32         `gorm:"column:sale_price;type:int;default:0;" json:"sale_price"`                           // 商品售价 扩展1000
	CommissionRate       uint32        `gorm:"column:commission_rate;type:uint;default:0;" json:"commission_rate"`                // 商品普通佣金率 扩大1000倍存储
	Volume               uint32        `gorm:"column:volume;type:uint;default:0;" json:"volume"`                                  // 销量
	ProductWeight        int32         `gorm:"column:product_weight;type:int;default:0;" json:"product_weight"`                   // 商品重量 单位：克
	ProductStatus        uint32        `gorm:"column:product_status;type:usmallint;default:0;" json:"product_status"`             // 商品状态0-下架 1-上架
	AuditStatus          int32         `gorm:"column:audit_status;type:smallint;default:0;" json:"audit_status"`                  // 审核状态0待审核 1审核待修改 2审核通过 3审核拒绝
	IsPopular            sql.NullInt64 `gorm:"column:is_popular;type:smallint;default:0;" json:"is_popular"`                      // 是否精选  1是  0否
	IntentionNum         uint64        `gorm:"column:intention_num;type:ubigint;default:0;" json:"intention_num"`                 // 带货量：意向池
	UpCommissionRateTime int32         `gorm:"column:up_commission_rate_time;type:int;default:0;" json:"up_commission_rate_time"` // 佣金最近更新时间
	IsFreeSample         int32         `gorm:"column:is_free_sample;type:int;default:1;" json:"is_free_sample"`                   // 是否免费样品 0 否  1 是
	LivePrice            int32         `gorm:"column:live_price;type:int;default:0;" json:"live_price"`                           // 直播价 扩大1000倍
	LivePriceInfo        string        `gorm:"column:live_price_info;type:varchar;size:20;" json:"live_price_info"`               // 直播价实现方式
	SellingPoints        string        `gorm:"column:selling_points;type:varchar;size:800;" json:"selling_points"`                // 卖点
	CommissionStatus     int32         `gorm:"column:commission_status;type:int;default:1;" json:"commission_status"`             // 佣金状态 0 关闭 1开启
	ModelTime
}

// TableName sets the insert table name for this struct type
func (p *ProductInfo) TableName() string {
	return "product_info"
}
