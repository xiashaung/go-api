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


CREATE TABLE `product_info` (
  `product_id` bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '央选商品信息表',
  `merchant_id` bigint(11) NOT NULL DEFAULT '0' COMMENT '商家id',
  `shop_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '店铺id',
  `brand_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '品牌id 0：没有品牌',
  `cate_root_id` int(11) NOT NULL DEFAULT '0' COMMENT '根类目',
  `cate_pid` int(11) NOT NULL DEFAULT '0' COMMENT '上一级类目',
  `cate_third` int(11) NOT NULL DEFAULT '0' COMMENT '三级类目',
  `cate_fourth` int(11) NOT NULL DEFAULT '0' COMMENT '四级类目',
  `cate_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '叶子节点类目',
  `platform_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '所属平台 1：抖音 2：快手',
  `third_product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '三方商品id',
  `product_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `product_pic` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品主图',
  `product_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品详情',
  `product_price` int(10) NOT NULL DEFAULT '0' COMMENT '商品原价 扩大1000',
  `sale_price` int(10) NOT NULL DEFAULT '0' COMMENT '商品售价 扩展1000',
  `commission_rate` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品普通佣金率 扩大1000倍存储',
  `volume` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '销量',
  `product_weight` int(11) NOT NULL DEFAULT '0' COMMENT '商品重量 单位：克',
  `product_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '商品状态0-下架 1-上架',
  `audit_status` smallint(6) NOT NULL DEFAULT '0' COMMENT '审核状态0待审核 1审核待修改 2审核通过 3审核拒绝',
  `is_popular` smallint(6) DEFAULT '0' COMMENT '是否精选  1是  0否',
  `intention_num` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '带货量：意向池',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  `up_commission_rate_time` int(10) NOT NULL DEFAULT '0' COMMENT '佣金最近更新时间',
  `is_free_sample` int(11) NOT NULL DEFAULT '1' COMMENT '是否免费样品 0 否  1 是',
  `live_price` int(10) NOT NULL DEFAULT '0' COMMENT '直播价 扩大1000倍',
  `live_price_info` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '直播价实现方式',
  `selling_points` varchar(800) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '卖点',
  `commission_status` int(10) NOT NULL DEFAULT '1' COMMENT '佣金状态 0 关闭 1开启',
  PRIMARY KEY (`product_id`),
  UNIQUE KEY `idx_third_pid` (`third_product_id`,`platform_id`) USING BTREE,
  UNIQUE KEY `uiq_productid` (`product_id`) USING BTREE,
  KEY `idx_merchant_platform` (`merchant_id`,`platform_id`),
  KEY `idx_add_time` (`add_time`),
  KEY `idx_sale_price` (`sale_price`),
  KEY `idx_cateid` (`cate_id`) USING BTREE,
  KEY `idx_cate_third` (`cate_third`),
  KEY `idx_cate_level` (`cate_root_id`,`cate_pid`,`cate_third`,`cate_fourth`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "product_id": 63,    "merchant_id": 35,    "shop_id": 79,    "brand_id": 2,    "cate_root_id": 57,    "cate_pid": 30,    "cate_third": 39,    "cate_fourth": 57,    "cate_id": 61,    "platform_id": 82,    "third_product_id": 52,    "product_name": "EZbKtiITbWPHOMrnLXCSdNdbN",    "product_pic": "EbAeFYPxDIAXOESebCtaJoIFR",    "product_desc": "uDEqfMHlBIvPVMQpYueOgFqMx",    "product_price": 19,    "sale_price": 98,    "commission_rate": 84,    "volume": 47,    "product_weight": 60,    "product_status": 48,    "audit_status": 76,    "is_popular": 61,    "intention_num": 63,    "add_time": 76,    "last_time": 46,    "up_commission_rate_time": 97,    "is_free_sample": 93,    "live_price": 25,    "live_price_info": "SLfLJMfMnSJGTBPNMiaMySdao",    "selling_points": "ukehFtcVWehcLsxPbYewyTevt",    "commission_status": 73}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 8] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[19] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned
[24] column is set for unsigned



*/

// ProductInfo struct is a row record of the product_info table in the yx database
type ProductInfo struct {
	//[ 0] product_id                                     ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ProductID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:product_id;type:ubigint;" json:"product_id"` // 央选商品信息表
	//[ 1] merchant_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	MerchantID int64 `gorm:"column:merchant_id;type:bigint;default:0;" json:"merchant_id"` // 商家id
	//[ 2] shop_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ShopID uint64 `gorm:"column:shop_id;type:ubigint;default:0;" json:"shop_id"` // 店铺id
	//[ 3] brand_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	BrandID uint32 `gorm:"column:brand_id;type:uint;default:0;" json:"brand_id"` // 品牌id 0：没有品牌
	//[ 4] cate_root_id                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateRootID int32 `gorm:"column:cate_root_id;type:int;default:0;" json:"cate_root_id"` // 根类目
	//[ 5] cate_pid                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CatePid int32 `gorm:"column:cate_pid;type:int;default:0;" json:"cate_pid"` // 上一级类目
	//[ 6] cate_third                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateThird int32 `gorm:"column:cate_third;type:int;default:0;" json:"cate_third"` // 三级类目
	//[ 7] cate_fourth                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateFourth int32 `gorm:"column:cate_fourth;type:int;default:0;" json:"cate_fourth"` // 四级类目
	//[ 8] cate_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateID uint32 `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"` // 叶子节点类目
	//[ 9] platform_id                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:smallint;default:0;" json:"platform_id"` // 所属平台 1：抖音 2：快手
	//[10] third_product_id                               bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ThirdProductID int64 `gorm:"column:third_product_id;type:bigint;default:0;" json:"third_product_id"` // 三方商品id
	//[11] product_name                                   varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:200;" json:"product_name"` // 商品名称
	//[12] product_pic                                    varchar(1000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1000    default: []
	ProductPic string `gorm:"column:product_pic;type:varchar;size:1000;" json:"product_pic"` // 商品主图
	//[13] product_desc                                   text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ProductDesc string `gorm:"column:product_desc;type:text;size:65535;" json:"product_desc"` // 商品详情
	//[14] product_price                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductPrice int32 `gorm:"column:product_price;type:int;default:0;" json:"product_price"` // 商品原价 扩大1000
	//[15] sale_price                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SalePrice int32 `gorm:"column:sale_price;type:int;default:0;" json:"sale_price"` // 商品售价 扩展1000
	//[16] commission_rate                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CommissionRate uint32 `gorm:"column:commission_rate;type:uint;default:0;" json:"commission_rate"` // 商品普通佣金率 扩大1000倍存储
	//[17] volume                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Volume uint32 `gorm:"column:volume;type:uint;default:0;" json:"volume"` // 销量
	//[18] product_weight                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductWeight int32 `gorm:"column:product_weight;type:int;default:0;" json:"product_weight"` // 商品重量 单位：克
	//[19] product_status                                 usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	ProductStatus uint32 `gorm:"column:product_status;type:usmallint;default:0;" json:"product_status"` // 商品状态0-下架 1-上架
	//[20] audit_status                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	AuditStatus int32 `gorm:"column:audit_status;type:smallint;default:0;" json:"audit_status"` // 审核状态0待审核 1审核待修改 2审核通过 3审核拒绝
	//[21] is_popular                                     smallint             null: true   primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsPopular sql.NullInt64 `gorm:"column:is_popular;type:smallint;default:0;" json:"is_popular"` // 是否精选  1是  0否
	//[22] intention_num                                  ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	IntentionNum uint64 `gorm:"column:intention_num;type:ubigint;default:0;" json:"intention_num"` // 带货量：意向池
	//[23] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 商品创建时间
	//[24] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
	//[25] up_commission_rate_time                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UpCommissionRateTime int32 `gorm:"column:up_commission_rate_time;type:int;default:0;" json:"up_commission_rate_time"` // 佣金最近更新时间
	//[26] is_free_sample                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	IsFreeSample int32 `gorm:"column:is_free_sample;type:int;default:1;" json:"is_free_sample"` // 是否免费样品 0 否  1 是
	//[27] live_price                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LivePrice int32 `gorm:"column:live_price;type:int;default:0;" json:"live_price"` // 直播价 扩大1000倍
	//[28] live_price_info                                varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	LivePriceInfo string `gorm:"column:live_price_info;type:varchar;size:20;" json:"live_price_info"` // 直播价实现方式
	//[29] selling_points                                 varchar(800)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 800     default: []
	SellingPoints string `gorm:"column:selling_points;type:varchar;size:800;" json:"selling_points"` // 卖点
	//[30] commission_status                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	CommissionStatus int32 `gorm:"column:commission_status;type:int;default:1;" json:"commission_status"` // 佣金状态 0 关闭 1开启

}

var product_infoTableInfo = &TableInfo{
	Name: "product_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "product_id",
			Comment:            `央选商品信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "merchant_id",
			Comment:            `商家id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int64",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "shop_id",
			Comment:            `店铺id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "uint64",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "brand_id",
			Comment:            `品牌id 0：没有品牌`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "BrandID",
			GoFieldType:        "uint32",
			JSONFieldName:      "brand_id",
			ProtobufFieldName:  "brand_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "cate_root_id",
			Comment:            `根类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateRootID",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_root_id",
			ProtobufFieldName:  "cate_root_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "cate_pid",
			Comment:            `上一级类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CatePid",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_pid",
			ProtobufFieldName:  "cate_pid",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "cate_third",
			Comment:            `三级类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateThird",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_third",
			ProtobufFieldName:  "cate_third",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "cate_fourth",
			Comment:            `四级类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateFourth",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_fourth",
			ProtobufFieldName:  "cate_fourth",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "cate_id",
			Comment:            `叶子节点类目`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "platform_id",
			Comment:            `所属平台 1：抖音 2：快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "third_product_id",
			Comment:            `三方商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "third_product_id",
			ProtobufFieldName:  "third_product_id",
			ProtobufType:       "int64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "product_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "ProductName",
			GoFieldType:        "string",
			JSONFieldName:      "product_name",
			ProtobufFieldName:  "product_name",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "product_pic",
			Comment:            `商品主图`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1000,
			GoFieldName:        "ProductPic",
			GoFieldType:        "string",
			JSONFieldName:      "product_pic",
			ProtobufFieldName:  "product_pic",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "product_desc",
			Comment:            `商品详情`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "ProductDesc",
			GoFieldType:        "string",
			JSONFieldName:      "product_desc",
			ProtobufFieldName:  "product_desc",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "product_price",
			Comment:            `商品原价 扩大1000`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "product_price",
			ProtobufFieldName:  "product_price",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "sale_price",
			Comment:            `商品售价 扩展1000`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SalePrice",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_price",
			ProtobufFieldName:  "sale_price",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "commission_rate",
			Comment:            `商品普通佣金率 扩大1000倍存储`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CommissionRate",
			GoFieldType:        "uint32",
			JSONFieldName:      "commission_rate",
			ProtobufFieldName:  "commission_rate",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "volume",
			Comment:            `销量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Volume",
			GoFieldType:        "uint32",
			JSONFieldName:      "volume",
			ProtobufFieldName:  "volume",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "product_weight",
			Comment:            `商品重量 单位：克`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductWeight",
			GoFieldType:        "int32",
			JSONFieldName:      "product_weight",
			ProtobufFieldName:  "product_weight",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "product_status",
			Comment:            `商品状态0-下架 1-上架`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "ProductStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "product_status",
			ProtobufFieldName:  "product_status",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "audit_status",
			Comment:            `审核状态0待审核 1审核待修改 2审核通过 3审核拒绝`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "AuditStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "audit_status",
			ProtobufFieldName:  "audit_status",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "is_popular",
			Comment:            `是否精选  1是  0否`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsPopular",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "is_popular",
			ProtobufFieldName:  "is_popular",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "intention_num",
			Comment:            `带货量：意向池`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "IntentionNum",
			GoFieldType:        "uint64",
			JSONFieldName:      "intention_num",
			ProtobufFieldName:  "intention_num",
			ProtobufType:       "uint64",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "add_time",
			Comment:            `商品创建时间`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "up_commission_rate_time",
			Comment:            `佣金最近更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UpCommissionRateTime",
			GoFieldType:        "int32",
			JSONFieldName:      "up_commission_rate_time",
			ProtobufFieldName:  "up_commission_rate_time",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "is_free_sample",
			Comment:            `是否免费样品 0 否  1 是`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsFreeSample",
			GoFieldType:        "int32",
			JSONFieldName:      "is_free_sample",
			ProtobufFieldName:  "is_free_sample",
			ProtobufType:       "int32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "live_price",
			Comment:            `直播价 扩大1000倍`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LivePrice",
			GoFieldType:        "int32",
			JSONFieldName:      "live_price",
			ProtobufFieldName:  "live_price",
			ProtobufType:       "int32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "live_price_info",
			Comment:            `直播价实现方式`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "LivePriceInfo",
			GoFieldType:        "string",
			JSONFieldName:      "live_price_info",
			ProtobufFieldName:  "live_price_info",
			ProtobufType:       "string",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "selling_points",
			Comment:            `卖点`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(800)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       800,
			GoFieldName:        "SellingPoints",
			GoFieldType:        "string",
			JSONFieldName:      "selling_points",
			ProtobufFieldName:  "selling_points",
			ProtobufType:       "string",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "commission_status",
			Comment:            `佣金状态 0 关闭 1开启`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommissionStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "commission_status",
			ProtobufFieldName:  "commission_status",
			ProtobufType:       "int32",
			ProtobufPos:        31,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProductInfo) TableName() string {
	return "product_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductInfo) TableInfo() *TableInfo {
	return product_infoTableInfo
}
