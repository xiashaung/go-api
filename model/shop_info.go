package model

import (
	"database/sql"
)

// ShopInfo struct is a row record of the shop_info table in the yx database
type ShopInfo struct {
	ShopID             uint32         `gorm:"primary_key;AUTO_INCREMENT;column:shop_id;type:uint;" json:"shop_id"`         // 店铺信息表
	MerchantID         uint32         `gorm:"column:merchant_id;type:uint;default:0;" json:"merchant_id"`                  // 商家id
	CateID             int32          `gorm:"column:cate_id;type:int;default:0;" json:"cate_id"`                           // 店铺经营类目
	PlatformID         int32          `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"`                   // 所属平台 1：抖音 2：快手
	ThirdShopid        int64          `gorm:"column:third_shopid;type:bigint;default:0;" json:"third_shopid"`              // 第三方平台id
	OpenID             string         `gorm:"column:open_id;type:char;size:32;" json:"open_id"`                            // 第三方平台的openid
	AccessToken        sql.NullString `gorm:"column:access_token;type:varchar;size:500;" json:"access_token"`              // 访问token
	RefreshToken       sql.NullString `gorm:"column:refresh_token;type:varchar;size:500;" json:"refresh_token"`            // 访问令牌，快手默认30天，抖音为14天
	ExpiresTime        uint32         `gorm:"column:expires_time;type:uint;default:0;" json:"expires_time"`                // oken过期时间
	RefreshExpiresTime int32          `gorm:"column:refresh_expires_time;type:int;default:0;" json:"refresh_expires_time"` // refresh_token过期时间
	SellerID           int64          `gorm:"column:seller_id;type:bigint;default:0;" json:"seller_id"`                    // 平台ID
	ShopName           string         `gorm:"column:shop_name;type:varchar;size:100;" json:"shop_name"`                    // 店铺名称
	Sex                int32          `gorm:"column:sex;type:tinyint;default:0;" json:"sex"`                               // 性别  1:男 2：女 3：保密
	ShopHeadpic        sql.NullString `gorm:"column:shop_headpic;type:varchar;size:300;" json:"shop_headpic"`              // 店铺图片
	ContactName        sql.NullString `gorm:"column:contact_name;type:varchar;size:10;" json:"contact_name"`               // 店铺联系人
	ContactMobile      sql.NullString `gorm:"column:contact_mobile;type:char;size:11;" json:"contact_mobile"`              // 店铺联系电话
	ContactAddr        sql.NullString `gorm:"column:contact_addr;type:varchar;size:200;" json:"contact_addr"`              // 店铺联系地址
	ShopCustomer       sql.NullString `gorm:"column:shop_customer;type:char;size:11;" json:"shop_customer"`                // 客服号码
	AuthStatus         uint32         `gorm:"column:auth_status;type:utinyint;default:1;" json:"auth_status"`              // 授权状态 1 是正常，0 是已过期
	SyncOrderTime      int32          `gorm:"column:sync_order_time;type:int;default:0;" json:"sync_order_time"`           // 同步订单时间(日期时间戳)
	ModelTime
}

// TableName sets the insert table name for this struct type
func (s *ShopInfo) TableName() string {
	return "shop_info"
}
