package model

import (
	"database/sql"
)

// MerchantInfo struct is a row record of the merchant_info table in the yx database
type MerchantInfo struct {
	MerchantID      int32         `gorm:"primary_key;AUTO_INCREMENT;column:merchant_id;type:int;" json:"merchant_id"` // 商家信息表
	MerchantMobile  string        `gorm:"column:merchant_mobile;type:char;size:11;" json:"merchant_mobile"`           // 商家手机号
	MerchantPasswd  string        `gorm:"column:merchant_passwd;type:char;size:32;" json:"merchant_passwd"`           // 商家密码
	MerchantName    string        `gorm:"column:merchant_name;type:varchar;size:100;" json:"merchant_name"`           // 商家姓名
	MerchantLevel   int32         `gorm:"column:merchant_level;type:int;default:1;" json:"merchant_level"`            // 商家级别
	CompanyID       int32         `gorm:"column:company_id;type:int;default:0;" json:"company_id"`                    // 公司信息
	CityID          int32         `gorm:"column:city_id;type:int;default:0;" json:"city_id"`                          // 商家所在城市
	MerchantStatus  uint32        `gorm:"column:merchant_status;type:uint;default:0;" json:"merchant_status"`         // 商家状态0:正常
	IsPublicContact int32         `gorm:"column:is_public_contact;type:tinyint;default:0;" json:"is_public_contact"`  // 公开联系方式 1是 0否
	ContactName     string        `gorm:"column:contact_name;type:varchar;size:255;" json:"contact_name"`             // 联系人
	JobTitle        string        `gorm:"column:job_title;type:varchar;size:255;" json:"job_title"`                   // 联系人职位
	ContactMobile   string        `gorm:"column:contact_mobile;type:varchar;size:20;" json:"contact_mobile"`          // 联系人电话
	PeLevel         uint32        `gorm:"column:pe_level;type:uint;default:0;" json:"pe_level"`                       // 开业活动 产品体验官等级
	Wechat          string        `gorm:"column:wechat;type:varchar;size:30;" json:"wechat"`                          // 联系人微信
	FromType        int32         `gorm:"column:from_type;type:int;default:0;" json:"from_type"`                      // 来源类型 1 pc 2小程序
	RegisterType    int32         `gorm:"column:register_type;type:int;default:0;" json:"register_type"`              // 注册类型 0 未知 1 商家 2 达人
	IsBlack         sql.NullInt64 `gorm:"column:is_black;type:tinyint;default:0;" json:"is_black"`                    // 商家是否进入了黑名单，0没有，1是
	LastAccessTime  int32         `gorm:"column:last_access_time;type:int;default:0;" json:"last_access_time"`        // 最后访问时间
	MerchantDesc    string        `gorm:"column:merchant_desc;type:varchar;size:255;" json:"merchant_desc"`           // 备注
	ModelTime
}

// TableName sets the insert table name for this struct type
func (m *MerchantInfo) TableName() string {
	return "merchant_info"
}
