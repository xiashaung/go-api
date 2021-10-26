package example

import (
	"database/sql"
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `shop_info` (
  `shop_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '店铺信息表',
  `merchant_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商家id',
  `cate_id` int(11) NOT NULL DEFAULT '0' COMMENT '店铺经营类目',
  `platform_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属平台 1：抖音 2：快手',
  `third_shopid` bigint(20) NOT NULL DEFAULT '0' COMMENT '第三方平台id',
  `open_id` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三方平台的openid',
  `access_token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '访问token',
  `refresh_token` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '访问令牌，快手默认30天，抖音为14天',
  `expires_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'oken过期时间',
  `refresh_expires_time` int(10) NOT NULL DEFAULT '0' COMMENT 'refresh_token过期时间',
  `seller_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '平台ID',
  `shop_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `sex` tinyint(2) NOT NULL DEFAULT '0' COMMENT '性别  1:男 2：女 3：保密',
  `shop_headpic` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '店铺图片',
  `contact_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '店铺联系人',
  `contact_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '店铺联系电话',
  `contact_addr` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '店铺联系地址',
  `shop_customer` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '客服号码',
  `auth_status` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '授权状态 1 是正常，0 是已过期',
  `sync_order_time` int(10) NOT NULL DEFAULT '0' COMMENT '同步订单时间(日期时间戳)',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '店铺入驻时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`shop_id`) USING BTREE,
  UNIQUE KEY `idx_open_platform_id` (`open_id`,`platform_id`) USING BTREE,
  KEY `idx_merchant_id` (`merchant_id`) USING BTREE,
  KEY `idx_seller_id` (`seller_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='店铺信息表'

JSON Sample
-------------------------------------
{    "shop_id": 1,    "merchant_id": 39,    "cate_id": 46,    "platform_id": 15,    "third_shopid": 64,    "open_id": "xhttEEcHLyvVagHFpgCxwUxWI",    "access_token": "ZHOjAQppYCqOKAwOpcuhnLeWK",    "refresh_token": "IecMKtZlqktZYZALfBthahOWY",    "expires_time": 58,    "refresh_expires_time": 16,    "seller_id": 29,    "shop_name": "wZvTYjKIBvKSSqmnQvfZbmDbu",    "sex": 5,    "shop_headpic": "xAdSYLCChoNJyhNLEDuTdgHwk",    "contact_name": "pNAZTWUwgvidbniulgDMjQCYm",    "contact_mobile": "JjHtisrLEqPrypCObJRDTtbgw",    "contact_addr": "rPQWDleYpCUySSKlDLvRomNCo",    "shop_customer": "eEZBQLkTHkpgNJCIgZujBphJV",    "auth_status": 40,    "sync_order_time": 86,    "add_time": 18,    "last_time": 97}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 8] column is set for unsigned
[18] column is set for unsigned
[20] column is set for unsigned
[21] column is set for unsigned



*/

// ShopInfo struct is a row record of the shop_info table in the yx database
type ShopInfo struct {
	//[ 0] shop_id                                        uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ShopID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:shop_id;type:uint;" json:"shop_id"` // 店铺信息表
	//[ 1] merchant_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	MerchantID uint32 `gorm:"column:merchant_id;type:uint;default:0;" json:"merchant_id"` // 商家id
	//[ 2] cate_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateID int32 `gorm:"column:cate_id;type:int;default:0;" json:"cate_id"` // 店铺经营类目
	//[ 3] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"` // 所属平台 1：抖音 2：快手
	//[ 4] third_shopid                                   bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ThirdShopid int64 `gorm:"column:third_shopid;type:bigint;default:0;" json:"third_shopid"` // 第三方平台id
	//[ 5] open_id                                        char(32)             null: false  primary: false  isArray: false  auto: false  col: char            len: 32      default: []
	OpenID string `gorm:"column:open_id;type:char;size:32;" json:"open_id"` // 第三方平台的openid
	//[ 6] access_token                                   varchar(500)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 500     default: []
	AccessToken sql.NullString `gorm:"column:access_token;type:varchar;size:500;" json:"access_token"` // 访问token
	//[ 7] refresh_token                                  varchar(500)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 500     default: []
	RefreshToken sql.NullString `gorm:"column:refresh_token;type:varchar;size:500;" json:"refresh_token"` // 访问令牌，快手默认30天，抖音为14天
	//[ 8] expires_time                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ExpiresTime uint32 `gorm:"column:expires_time;type:uint;default:0;" json:"expires_time"` // oken过期时间
	//[ 9] refresh_expires_time                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RefreshExpiresTime int32 `gorm:"column:refresh_expires_time;type:int;default:0;" json:"refresh_expires_time"` // refresh_token过期时间
	//[10] seller_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SellerID int64 `gorm:"column:seller_id;type:bigint;default:0;" json:"seller_id"` // 平台ID
	//[11] shop_name                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ShopName string `gorm:"column:shop_name;type:varchar;size:100;" json:"shop_name"` // 店铺名称
	//[12] sex                                            tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Sex int32 `gorm:"column:sex;type:tinyint;default:0;" json:"sex"` // 性别  1:男 2：女 3：保密
	//[13] shop_headpic                                   varchar(300)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 300     default: []
	ShopHeadpic sql.NullString `gorm:"column:shop_headpic;type:varchar;size:300;" json:"shop_headpic"` // 店铺图片
	//[14] contact_name                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	ContactName sql.NullString `gorm:"column:contact_name;type:varchar;size:10;" json:"contact_name"` // 店铺联系人
	//[15] contact_mobile                                 char(11)             null: true   primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	ContactMobile sql.NullString `gorm:"column:contact_mobile;type:char;size:11;" json:"contact_mobile"` // 店铺联系电话
	//[16] contact_addr                                   varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	ContactAddr sql.NullString `gorm:"column:contact_addr;type:varchar;size:200;" json:"contact_addr"` // 店铺联系地址
	//[17] shop_customer                                  char(11)             null: true   primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	ShopCustomer sql.NullString `gorm:"column:shop_customer;type:char;size:11;" json:"shop_customer"` // 客服号码
	//[18] auth_status                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	AuthStatus uint32 `gorm:"column:auth_status;type:utinyint;default:1;" json:"auth_status"` // 授权状态 1 是正常，0 是已过期
	//[19] sync_order_time                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SyncOrderTime int32 `gorm:"column:sync_order_time;type:int;default:0;" json:"sync_order_time"` // 同步订单时间(日期时间戳)
	//[20] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 店铺入驻时间
	//[21] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var shop_infoTableInfo = &TableInfo{
	Name: "shop_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "shop_id",
			Comment:            `店铺信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "uint32",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "merchant_id",
			Comment:            `商家id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "uint32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_id",
			Comment:            `店铺经营类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "platform_id",
			Comment:            `所属平台 1：抖音 2：快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "third_shopid",
			Comment:            `第三方平台id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdShopid",
			GoFieldType:        "int64",
			JSONFieldName:      "third_shopid",
			ProtobufFieldName:  "third_shopid",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "open_id",
			Comment:            `第三方平台的openid`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       32,
			GoFieldName:        "OpenID",
			GoFieldType:        "string",
			JSONFieldName:      "open_id",
			ProtobufFieldName:  "open_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "access_token",
			Comment:            `访问token`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(500)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       500,
			GoFieldName:        "AccessToken",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "access_token",
			ProtobufFieldName:  "access_token",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "refresh_token",
			Comment:            `访问令牌，快手默认30天，抖音为14天`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(500)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       500,
			GoFieldName:        "RefreshToken",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "refresh_token",
			ProtobufFieldName:  "refresh_token",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "expires_time",
			Comment:            `oken过期时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpiresTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "expires_time",
			ProtobufFieldName:  "expires_time",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "refresh_expires_time",
			Comment:            `refresh_token过期时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RefreshExpiresTime",
			GoFieldType:        "int32",
			JSONFieldName:      "refresh_expires_time",
			ProtobufFieldName:  "refresh_expires_time",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "seller_id",
			Comment:            `平台ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SellerID",
			GoFieldType:        "int64",
			JSONFieldName:      "seller_id",
			ProtobufFieldName:  "seller_id",
			ProtobufType:       "int64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "shop_name",
			Comment:            `店铺名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ShopName",
			GoFieldType:        "string",
			JSONFieldName:      "shop_name",
			ProtobufFieldName:  "shop_name",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "sex",
			Comment:            `性别  1:男 2：女 3：保密`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Sex",
			GoFieldType:        "int32",
			JSONFieldName:      "sex",
			ProtobufFieldName:  "sex",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "shop_headpic",
			Comment:            `店铺图片`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       300,
			GoFieldName:        "ShopHeadpic",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "shop_headpic",
			ProtobufFieldName:  "shop_headpic",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "contact_name",
			Comment:            `店铺联系人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "ContactName",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "contact_name",
			ProtobufFieldName:  "contact_name",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "contact_mobile",
			Comment:            `店铺联系电话`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "ContactMobile",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "contact_mobile",
			ProtobufFieldName:  "contact_mobile",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "contact_addr",
			Comment:            `店铺联系地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "ContactAddr",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "contact_addr",
			ProtobufFieldName:  "contact_addr",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "shop_customer",
			Comment:            `客服号码`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "ShopCustomer",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "shop_customer",
			ProtobufFieldName:  "shop_customer",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "auth_status",
			Comment:            `授权状态 1 是正常，0 是已过期`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "AuthStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "auth_status",
			ProtobufFieldName:  "auth_status",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "sync_order_time",
			Comment:            `同步订单时间(日期时间戳)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SyncOrderTime",
			GoFieldType:        "int32",
			JSONFieldName:      "sync_order_time",
			ProtobufFieldName:  "sync_order_time",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "add_time",
			Comment:            `店铺入驻时间`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
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
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShopInfo) TableName() string {
	return "shop_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopInfo) TableInfo() *TableInfo {
	return shop_infoTableInfo
}
