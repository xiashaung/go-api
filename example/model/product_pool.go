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


CREATE TABLE `product_pool` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主播商品池信息表',
  `anchor_id` int(11) NOT NULL DEFAULT '0' COMMENT '主播id',
  `talent_id` int(11) NOT NULL DEFAULT '0' COMMENT '达人id',
  `product_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品id',
  `stock` int(11) NOT NULL DEFAULT '0' COMMENT '库存',
  `merchant_id` int(10) NOT NULL DEFAULT '0' COMMENT '商家ID',
  `shop_id` bigint(11) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `cate_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属类目三级',
  `brand_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '品牌id 0：没有品牌',
  `tender_id` int(11) NOT NULL DEFAULT '0' COMMENT '招商计划id 0：主播自己添加的意向商品',
  `platform_id` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '所属平台:1 抖音 2 快手 ',
  `pool_type` int(11) NOT NULL DEFAULT '0' COMMENT '商品池类型 0：精选池 1：报名池 2：通用意向池 3：通用选品池 4 招商意向池  5 招商选品池',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '样品状态 0 未申请  1 待付款  2 待发货  3 待收货   4已完成',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '佣金率',
  `commission_type` int(11) NOT NULL DEFAULT '0' COMMENT '佣金率类型 0 通用 1 定向  2 专属',
  `commission_status` int(1) NOT NULL DEFAULT '1' COMMENT '佣金状态 0 关闭  1 开始',
  `product_price` int(11) NOT NULL DEFAULT '0' COMMENT '商品售卖价单位 厘',
  `sale_price` int(11) NOT NULL DEFAULT '0' COMMENT '商品直播价 单位 厘',
  `plan_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快手定向计划ID',
  `activity_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动id',
  `up_commission_rate_time` int(10) NOT NULL DEFAULT '0' COMMENT '定向佣金最近更新时间',
  `pool_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '选品池状态 1有效 2可恢复失效 3删除',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加商品时间',
  `last_time` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tender_pro` (`tender_id`,`product_id`,`talent_id`,`pool_type`) USING BTREE,
  KEY `idx_mt_id` (`merchant_id`,`tender_id`),
  KEY `idx_pid_talent_platform` (`product_id`,`talent_id`,`platform_id`) USING BTREE,
  KEY `idx_pap_id` (`plan_id`,`product_id`,`anchor_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=247 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 77,    "anchor_id": 58,    "talent_id": 16,    "product_id": 99,    "stock": 86,    "merchant_id": 92,    "shop_id": 33,    "cate_id": 32,    "brand_id": 73,    "tender_id": 40,    "platform_id": 84,    "pool_type": 70,    "status": 72,    "commission_rate": 17,    "commission_type": 49,    "commission_status": 83,    "product_price": 19,    "sale_price": 11,    "plan_id": 91,    "activity_id": 88,    "up_commission_rate_time": 83,    "pool_status": 8,    "add_time": 11,    "last_time": 31}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[10] column is set for unsigned
[19] column is set for unsigned



*/

// ProductPool struct is a row record of the product_pool table in the yx database
type ProductPool struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 主播商品池信息表
	//[ 1] anchor_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AnchorID int32 `gorm:"column:anchor_id;type:int;default:0;" json:"anchor_id"` // 主播id
	//[ 2] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"` // 达人id
	//[ 3] product_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductID int32 `gorm:"column:product_id;type:int;default:0;" json:"product_id"` // 商品id
	//[ 4] stock                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Stock int32 `gorm:"column:stock;type:int;default:0;" json:"stock"` // 库存
	//[ 5] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"` // 商家ID
	//[ 6] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ShopID int64 `gorm:"column:shop_id;type:bigint;default:0;" json:"shop_id"` // 店铺ID
	//[ 7] cate_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateID uint32 `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"` // 所属类目三级
	//[ 8] brand_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	BrandID uint32 `gorm:"column:brand_id;type:uint;default:0;" json:"brand_id"` // 品牌id 0：没有品牌
	//[ 9] tender_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderID int32 `gorm:"column:tender_id;type:int;default:0;" json:"tender_id"` // 招商计划id 0：主播自己添加的意向商品
	//[10] platform_id                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	PlatformID uint32 `gorm:"column:platform_id;type:utinyint;default:0;" json:"platform_id"` // 所属平台:1 抖音 2 快手
	//[11] pool_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PoolType int32 `gorm:"column:pool_type;type:int;default:0;" json:"pool_type"` // 商品池类型 0：精选池 1：报名池 2：通用意向池 3：通用选品池 4 招商意向池  5 招商选品池
	//[12] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status int32 `gorm:"column:status;type:int;default:0;" json:"status"` // 样品状态 0 未申请  1 待付款  2 待发货  3 待收货   4已完成
	//[13] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 佣金率
	//[14] commission_type                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionType int32 `gorm:"column:commission_type;type:int;default:0;" json:"commission_type"` // 佣金率类型 0 通用 1 定向  2 专属
	//[15] commission_status                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	CommissionStatus int32 `gorm:"column:commission_status;type:int;default:1;" json:"commission_status"` // 佣金状态 0 关闭  1 开始
	//[16] product_price                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductPrice int32 `gorm:"column:product_price;type:int;default:0;" json:"product_price"` // 商品售卖价单位 厘
	//[17] sale_price                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SalePrice int32 `gorm:"column:sale_price;type:int;default:0;" json:"sale_price"` // 商品直播价 单位 厘
	//[18] plan_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PlanID int64 `gorm:"column:plan_id;type:bigint;default:0;" json:"plan_id"` // 快手定向计划ID
	//[19] activity_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ActivityID uint32 `gorm:"column:activity_id;type:uint;default:0;" json:"activity_id"` // 活动id
	//[20] up_commission_rate_time                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UpCommissionRateTime int32 `gorm:"column:up_commission_rate_time;type:int;default:0;" json:"up_commission_rate_time"` // 定向佣金最近更新时间
	//[21] pool_status                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	PoolStatus int32 `gorm:"column:pool_status;type:tinyint;default:1;" json:"pool_status"` // 选品池状态 1有效 2可恢复失效 3删除
	//[22] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加商品时间
	//[23] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var product_poolTableInfo = &TableInfo{
	Name: "product_pool",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主播商品池信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "int32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_id",
			Comment:            `商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "int32",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "stock",
			Comment:            `库存`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Stock",
			GoFieldType:        "int32",
			JSONFieldName:      "stock",
			ProtobufFieldName:  "stock",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "merchant_id",
			Comment:            `商家ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "shop_id",
			Comment:            `店铺ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "int64",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "int64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "cate_id",
			Comment:            `所属类目三级`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "tender_id",
			Comment:            `招商计划id 0：主播自己添加的意向商品`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderID",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_id",
			ProtobufFieldName:  "tender_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "platform_id",
			Comment:            `所属平台:1 抖音 2 快手 `,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "uint32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "pool_type",
			Comment:            `商品池类型 0：精选池 1：报名池 2：通用意向池 3：通用选品池 4 招商意向池  5 招商选品池`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PoolType",
			GoFieldType:        "int32",
			JSONFieldName:      "pool_type",
			ProtobufFieldName:  "pool_type",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "status",
			Comment:            `样品状态 0 未申请  1 待付款  2 待发货  3 待收货   4已完成`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "commission_rate",
			Comment:            `佣金率`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommissionRate",
			GoFieldType:        "int32",
			JSONFieldName:      "commission_rate",
			ProtobufFieldName:  "commission_rate",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "commission_type",
			Comment:            `佣金率类型 0 通用 1 定向  2 专属`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommissionType",
			GoFieldType:        "int32",
			JSONFieldName:      "commission_type",
			ProtobufFieldName:  "commission_type",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "commission_status",
			Comment:            `佣金状态 0 关闭  1 开始`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "product_price",
			Comment:            `商品售卖价单位 厘`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "sale_price",
			Comment:            `商品直播价 单位 厘`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "plan_id",
			Comment:            `快手定向计划ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PlanID",
			GoFieldType:        "int64",
			JSONFieldName:      "plan_id",
			ProtobufFieldName:  "plan_id",
			ProtobufType:       "int64",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "activity_id",
			Comment:            `活动id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ActivityID",
			GoFieldType:        "uint32",
			JSONFieldName:      "activity_id",
			ProtobufFieldName:  "activity_id",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "up_commission_rate_time",
			Comment:            `定向佣金最近更新时间`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "pool_status",
			Comment:            `选品池状态 1有效 2可恢复失效 3删除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "PoolStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "pool_status",
			ProtobufFieldName:  "pool_status",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "add_time",
			Comment:            `添加商品时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "int32",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "last_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "int32",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProductPool) TableName() string {
	return "product_pool"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductPool) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductPool) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductPool) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductPool) TableInfo() *TableInfo {
	return product_poolTableInfo
}
