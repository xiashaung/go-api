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


CREATE TABLE `sample_order_product` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sample_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '样品订单号',
  `product_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '央选商品id',
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '商品名称',
  `product_pic` varchar(255) NOT NULL COMMENT '商品头图',
  `talent_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `anchor_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '主播id',
  `shop_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '店铺id',
  `merchant_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '第三方商家id',
  `sku_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT 'skuid',
  `sku_info` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '订单sku信息',
  `cate_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类目id',
  `brand_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '品牌id',
  `platform_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '平台id',
  `sale_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '售卖价',
  `quantity` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '购买数量',
  `total_fee` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总价',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid_talentid` (`product_id`,`talent_id`) USING BTREE,
  KEY `idx_sn` (`sample_sn`,`talent_id`) USING BTREE,
  KEY `idx_sn_merchantid` (`sample_sn`,`merchant_id`) USING BTREE,
  KEY `idx_pid_merchantid` (`product_id`,`merchant_id`) USING BTREE,
  KEY `idx_talentid_status` (`talent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=133 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 6,    "sample_sn": "MfdsTLRjhtHpDULMFQvhffPWA",    "product_id": 51,    "product_name": "MhKLSqMAgaTShYwRnGqytZAnI",    "product_pic": "aXmkSkFOduoAbrQajETeDYnQP",    "talent_id": 34,    "anchor_id": 71,    "shop_id": 44,    "merchant_id": 25,    "sku_id": 61,    "sku_info": "BRdGbAbXuepTgabLpDgQhwdAO",    "cate_id": 79,    "brand_id": 9,    "platform_id": 69,    "sale_price": 35,    "quantity": 70,    "total_fee": 42,    "add_time": 43,    "last_time": 25}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned



*/

// SampleOrderProduct struct is a row record of the sample_order_product table in the yx database
type SampleOrderProduct struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] sample_sn                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SampleSn string `gorm:"column:sample_sn;type:varchar;size:255;" json:"sample_sn"` // 样品订单号
	//[ 2] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 央选商品id
	//[ 3] product_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:255;" json:"product_name"` // 商品名称
	//[ 4] product_pic                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductPic string `gorm:"column:product_pic;type:varchar;size:255;" json:"product_pic"` // 商品头图
	//[ 5] talent_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TalentID uint64 `gorm:"column:talent_id;type:ubigint;default:0;" json:"talent_id"` // 达人id
	//[ 6] anchor_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	AnchorID uint64 `gorm:"column:anchor_id;type:ubigint;default:0;" json:"anchor_id"` // 主播id
	//[ 7] shop_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ShopID uint32 `gorm:"column:shop_id;type:uint;default:0;" json:"shop_id"` // 店铺id
	//[ 8] merchant_id                                    ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	MerchantID uint64 `gorm:"column:merchant_id;type:ubigint;default:0;" json:"merchant_id"` // 第三方商家id
	//[ 9] sku_id                                         ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SkuID uint64 `gorm:"column:sku_id;type:ubigint;default:0;" json:"sku_id"` // skuid
	//[10] sku_info                                       varchar(512)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	SkuInfo string `gorm:"column:sku_info;type:varchar;size:512;" json:"sku_info"` // 订单sku信息
	//[11] cate_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateID uint32 `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"` // 类目id
	//[12] brand_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	BrandID uint32 `gorm:"column:brand_id;type:uint;default:0;" json:"brand_id"` // 品牌id
	//[13] platform_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PlatformID uint32 `gorm:"column:platform_id;type:uint;default:0;" json:"platform_id"` // 平台id
	//[14] sale_price                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SalePrice uint32 `gorm:"column:sale_price;type:uint;default:0;" json:"sale_price"` // 售卖价
	//[15] quantity                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Quantity uint32 `gorm:"column:quantity;type:uint;default:0;" json:"quantity"` // 购买数量
	//[16] total_fee                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TotalFee uint32 `gorm:"column:total_fee;type:uint;default:0;" json:"total_fee"` // 总价
	//[17] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[18] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 修改时间

}

var sample_order_productTableInfo = &TableInfo{
	Name: "sample_order_product",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
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
			Name:               "sample_sn",
			Comment:            `样品订单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SampleSn",
			GoFieldType:        "string",
			JSONFieldName:      "sample_sn",
			ProtobufFieldName:  "sample_sn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "product_id",
			Comment:            `央选商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ProductName",
			GoFieldType:        "string",
			JSONFieldName:      "product_name",
			ProtobufFieldName:  "product_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "product_pic",
			Comment:            `商品头图`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ProductPic",
			GoFieldType:        "string",
			JSONFieldName:      "product_pic",
			ProtobufFieldName:  "product_pic",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint64",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "uint64",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "shop_id",
			Comment:            `店铺id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "uint32",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "merchant_id",
			Comment:            `第三方商家id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "uint64",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "sku_id",
			Comment:            `skuid`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SkuID",
			GoFieldType:        "uint64",
			JSONFieldName:      "sku_id",
			ProtobufFieldName:  "sku_id",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "sku_info",
			Comment:            `订单sku信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       512,
			GoFieldName:        "SkuInfo",
			GoFieldType:        "string",
			JSONFieldName:      "sku_info",
			ProtobufFieldName:  "sku_info",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "cate_id",
			Comment:            `类目id`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "brand_id",
			Comment:            `品牌id`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "platform_id",
			Comment:            `平台id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "uint32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "sale_price",
			Comment:            `售卖价`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SalePrice",
			GoFieldType:        "uint32",
			JSONFieldName:      "sale_price",
			ProtobufFieldName:  "sale_price",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "quantity",
			Comment:            `购买数量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Quantity",
			GoFieldType:        "uint32",
			JSONFieldName:      "quantity",
			ProtobufFieldName:  "quantity",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "total_fee",
			Comment:            `总价`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TotalFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "total_fee",
			ProtobufFieldName:  "total_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "add_time",
			Comment:            `创建时间`,
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
			Comment:            `修改时间`,
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
	},
}

// TableName sets the insert table name for this struct type
func (s *SampleOrderProduct) TableName() string {
	return "sample_order_product"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SampleOrderProduct) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SampleOrderProduct) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SampleOrderProduct) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SampleOrderProduct) TableInfo() *TableInfo {
	return sample_order_productTableInfo
}
