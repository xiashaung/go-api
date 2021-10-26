package example

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `merchant_info` (
  `merchant_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '商家信息表',
  `merchant_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商家手机号',
  `merchant_passwd` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商家密码',
  `merchant_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商家姓名',
  `merchant_level` int(255) NOT NULL DEFAULT '1' COMMENT '商家级别',
  `company_id` int(11) NOT NULL DEFAULT '0' COMMENT '公司信息',
  `city_id` int(11) NOT NULL DEFAULT '0' COMMENT '商家所在城市',
  `merchant_status` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '商家状态0:正常',
  `is_public_contact` tinyint(1) NOT NULL DEFAULT '0' COMMENT '公开联系方式 1是 0否',
  `contact_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `job_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人职位',
  `contact_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人电话',
  `pe_level` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开业活动 产品体验官等级',
  `wechat` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人微信',
  `from_type` int(11) NOT NULL DEFAULT '0' COMMENT '来源类型 1 pc 2小程序',
  `register_type` int(11) NOT NULL DEFAULT '0' COMMENT '注册类型 0 未知 1 商家 2 达人',
  `is_black` tinyint(3) DEFAULT '0' COMMENT '商家是否进入了黑名单，0没有，1是',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商家入驻时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_access_time` int(10) NOT NULL DEFAULT '0' COMMENT '最后访问时间',
  `merchant_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`merchant_id`) USING BTREE,
  UNIQUE KEY `idx_mobile` (`merchant_mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=284 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商家信息表'

JSON Sample
-------------------------------------
{    "merchant_id": 41,    "merchant_mobile": "riVZmswlnArDoDFbYkcpOIpPh",    "merchant_passwd": "EEYTSiufmNVnAukiBxNvNGBem",    "merchant_name": "ZBKAYfWAxNBksidpAGTYgqGkY",    "merchant_level": 93,    "company_id": 27,    "city_id": 79,    "merchant_status": 98,    "is_public_contact": 10,    "contact_name": "FxacQfYSresnUTemlXLOUXDQV",    "job_title": "gwAQrfVgujUlkOkcQBLFTQDft",    "contact_mobile": "lgibNkTwvkLFIsLKvSsaVlnAh",    "pe_level": 88,    "wechat": "HjcZJwxxoZQHUsiRdmtRpHQTs",    "from_type": 91,    "register_type": 76,    "is_black": 21,    "add_time": 66,    "last_time": 11,    "last_access_time": 57,    "merchant_desc": "OvEVIjXqlTepLqOtQysrZkpyq"}


Comments
-------------------------------------
[ 7] column is set for unsigned
[12] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned



*/

// MerchantInfo struct is a row record of the merchant_info table in the yx database
type MerchantInfo struct {
	//[ 0] merchant_id                                    int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MerchantID int32 `gorm:"primary_key;AUTO_INCREMENT;column:merchant_id;type:int;" json:"merchant_id"` // 商家信息表
	//[ 1] merchant_mobile                                char(11)             null: false  primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	MerchantMobile string `gorm:"column:merchant_mobile;type:char;size:11;" json:"merchant_mobile"` // 商家手机号
	//[ 2] merchant_passwd                                char(32)             null: false  primary: false  isArray: false  auto: false  col: char            len: 32      default: []
	MerchantPasswd string `gorm:"column:merchant_passwd;type:char;size:32;" json:"merchant_passwd"` // 商家密码
	//[ 3] merchant_name                                  varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	MerchantName string `gorm:"column:merchant_name;type:varchar;size:100;" json:"merchant_name"` // 商家姓名
	//[ 4] merchant_level                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	MerchantLevel int32 `gorm:"column:merchant_level;type:int;default:1;" json:"merchant_level"` // 商家级别
	//[ 5] company_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CompanyID int32 `gorm:"column:company_id;type:int;default:0;" json:"company_id"` // 公司信息
	//[ 6] city_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CityID int32 `gorm:"column:city_id;type:int;default:0;" json:"city_id"` // 商家所在城市
	//[ 7] merchant_status                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	MerchantStatus uint32 `gorm:"column:merchant_status;type:uint;default:0;" json:"merchant_status"` // 商家状态0:正常
	//[ 8] is_public_contact                              tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsPublicContact int32 `gorm:"column:is_public_contact;type:tinyint;default:0;" json:"is_public_contact"` // 公开联系方式 1是 0否
	//[ 9] contact_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactName string `gorm:"column:contact_name;type:varchar;size:255;" json:"contact_name"` // 联系人
	//[10] job_title                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	JobTitle string `gorm:"column:job_title;type:varchar;size:255;" json:"job_title"` // 联系人职位
	//[11] contact_mobile                                 varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	ContactMobile string `gorm:"column:contact_mobile;type:varchar;size:20;" json:"contact_mobile"` // 联系人电话
	//[12] pe_level                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PeLevel uint32 `gorm:"column:pe_level;type:uint;default:0;" json:"pe_level"` // 开业活动 产品体验官等级
	//[13] wechat                                         varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Wechat string `gorm:"column:wechat;type:varchar;size:30;" json:"wechat"` // 联系人微信
	//[14] from_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FromType int32 `gorm:"column:from_type;type:int;default:0;" json:"from_type"` // 来源类型 1 pc 2小程序
	//[15] register_type                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RegisterType int32 `gorm:"column:register_type;type:int;default:0;" json:"register_type"` // 注册类型 0 未知 1 商家 2 达人
	//[16] is_black                                       tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsBlack sql.NullInt64 `gorm:"column:is_black;type:tinyint;default:0;" json:"is_black"` // 商家是否进入了黑名单，0没有，1是
	//[17] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 商家入驻时间
	//[18] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
	//[19] last_access_time                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastAccessTime int32 `gorm:"column:last_access_time;type:int;default:0;" json:"last_access_time"` // 最后访问时间
	//[20] merchant_desc                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MerchantDesc string `gorm:"column:merchant_desc;type:varchar;size:255;" json:"merchant_desc"` // 备注

}

var merchant_infoTableInfo = &TableInfo{
	Name: "merchant_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "merchant_id",
			Comment:            `商家信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "merchant_mobile",
			Comment:            `商家手机号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "MerchantMobile",
			GoFieldType:        "string",
			JSONFieldName:      "merchant_mobile",
			ProtobufFieldName:  "merchant_mobile",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "merchant_passwd",
			Comment:            `商家密码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       32,
			GoFieldName:        "MerchantPasswd",
			GoFieldType:        "string",
			JSONFieldName:      "merchant_passwd",
			ProtobufFieldName:  "merchant_passwd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "merchant_name",
			Comment:            `商家姓名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "MerchantName",
			GoFieldType:        "string",
			JSONFieldName:      "merchant_name",
			ProtobufFieldName:  "merchant_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "merchant_level",
			Comment:            `商家级别`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MerchantLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_level",
			ProtobufFieldName:  "merchant_level",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "company_id",
			Comment:            `公司信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CompanyID",
			GoFieldType:        "int32",
			JSONFieldName:      "company_id",
			ProtobufFieldName:  "company_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "city_id",
			Comment:            `商家所在城市`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CityID",
			GoFieldType:        "int32",
			JSONFieldName:      "city_id",
			ProtobufFieldName:  "city_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "merchant_status",
			Comment:            `商家状态0:正常`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "merchant_status",
			ProtobufFieldName:  "merchant_status",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "is_public_contact",
			Comment:            `公开联系方式 1是 0否`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsPublicContact",
			GoFieldType:        "int32",
			JSONFieldName:      "is_public_contact",
			ProtobufFieldName:  "is_public_contact",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "contact_name",
			Comment:            `联系人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ContactName",
			GoFieldType:        "string",
			JSONFieldName:      "contact_name",
			ProtobufFieldName:  "contact_name",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "job_title",
			Comment:            `联系人职位`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "JobTitle",
			GoFieldType:        "string",
			JSONFieldName:      "job_title",
			ProtobufFieldName:  "job_title",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "contact_mobile",
			Comment:            `联系人电话`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "ContactMobile",
			GoFieldType:        "string",
			JSONFieldName:      "contact_mobile",
			ProtobufFieldName:  "contact_mobile",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "pe_level",
			Comment:            `开业活动 产品体验官等级`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PeLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "pe_level",
			ProtobufFieldName:  "pe_level",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "wechat",
			Comment:            `联系人微信`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Wechat",
			GoFieldType:        "string",
			JSONFieldName:      "wechat",
			ProtobufFieldName:  "wechat",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "from_type",
			Comment:            `来源类型 1 pc 2小程序`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FromType",
			GoFieldType:        "int32",
			JSONFieldName:      "from_type",
			ProtobufFieldName:  "from_type",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "register_type",
			Comment:            `注册类型 0 未知 1 商家 2 达人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RegisterType",
			GoFieldType:        "int32",
			JSONFieldName:      "register_type",
			ProtobufFieldName:  "register_type",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "is_black",
			Comment:            `商家是否进入了黑名单，0没有，1是`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsBlack",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "is_black",
			ProtobufFieldName:  "is_black",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "add_time",
			Comment:            `商家入驻时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "last_time",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "last_access_time",
			Comment:            `最后访问时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastAccessTime",
			GoFieldType:        "int32",
			JSONFieldName:      "last_access_time",
			ProtobufFieldName:  "last_access_time",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "merchant_desc",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "MerchantDesc",
			GoFieldType:        "string",
			JSONFieldName:      "merchant_desc",
			ProtobufFieldName:  "merchant_desc",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MerchantInfo) TableName() string {
	return "merchant_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MerchantInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MerchantInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MerchantInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MerchantInfo) TableInfo() *TableInfo {
	return merchant_infoTableInfo
}
