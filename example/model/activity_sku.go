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


CREATE TABLE `activity_sku` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `activity_product_id` int(11) NOT NULL DEFAULT '0' COMMENT '活动产品id',
  `product_id` int(11) NOT NULL DEFAULT '0' COMMENT '央选产品id',
  `original_price` int(11) NOT NULL DEFAULT '0' COMMENT '原价',
  `sku_id` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '第三方平台sku id',
  `seller_price` int(11) NOT NULL DEFAULT '0' COMMENT '推广售价',
  `exposure_talent_num` int(11) NOT NULL DEFAULT '0' COMMENT '曝光达人数',
  `commerce_talent_num` int(11) NOT NULL DEFAULT '0' COMMENT '带货达人数',
  `valid_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '有效期,必须活动结束有效期',
  `estimated_revenue` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '预计收入',
  `sell_num` int(11) NOT NULL DEFAULT '0' COMMENT '活动销售数量',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='活动sku信息'

JSON Sample
-------------------------------------
{    "id": 13,    "activity_product_id": 1,    "product_id": 16,    "original_price": 22,    "sku_id": 64,    "seller_price": 36,    "exposure_talent_num": 32,    "commerce_talent_num": 56,    "valid_time": 78,    "estimated_revenue": 46,    "sell_num": 27,    "add_time": 11,    "last_time": 41}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 4] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned



*/

// ActivitySku struct is a row record of the activity_sku table in the yx database
type ActivitySku struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] activity_product_id                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ActivityProductID int32 `gorm:"column:activity_product_id;type:int;default:0;" json:"activity_product_id"` // 活动产品id
	//[ 2] product_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductID int32 `gorm:"column:product_id;type:int;default:0;" json:"product_id"` // 央选产品id
	//[ 3] original_price                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OriginalPrice int32 `gorm:"column:original_price;type:int;default:0;" json:"original_price"` // 原价
	//[ 4] sku_id                                         ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SkuID uint64 `gorm:"column:sku_id;type:ubigint;default:0;" json:"sku_id"` // 第三方平台sku id
	//[ 5] seller_price                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SellerPrice int32 `gorm:"column:seller_price;type:int;default:0;" json:"seller_price"` // 推广售价
	//[ 6] exposure_talent_num                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExposureTalentNum int32 `gorm:"column:exposure_talent_num;type:int;default:0;" json:"exposure_talent_num"` // 曝光达人数
	//[ 7] commerce_talent_num                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommerceTalentNum int32 `gorm:"column:commerce_talent_num;type:int;default:0;" json:"commerce_talent_num"` // 带货达人数
	//[ 8] valid_time                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ValidTime uint32 `gorm:"column:valid_time;type:uint;default:0;" json:"valid_time"` // 有效期,必须活动结束有效期
	//[ 9] estimated_revenue                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EstimatedRevenue uint32 `gorm:"column:estimated_revenue;type:uint;default:0;" json:"estimated_revenue"` // 预计收入
	//[10] sell_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SellNum int32 `gorm:"column:sell_num;type:int;default:0;" json:"sell_num"` // 活动销售数量
	//[11] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"`
	//[12] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var activity_skuTableInfo = &TableInfo{
	Name: "activity_sku",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "activity_product_id",
			Comment:            `活动产品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ActivityProductID",
			GoFieldType:        "int32",
			JSONFieldName:      "activity_product_id",
			ProtobufFieldName:  "activity_product_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "product_id",
			Comment:            `央选产品id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "original_price",
			Comment:            `原价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OriginalPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "original_price",
			ProtobufFieldName:  "original_price",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "sku_id",
			Comment:            `第三方平台sku id`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "seller_price",
			Comment:            `推广售价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SellerPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "seller_price",
			ProtobufFieldName:  "seller_price",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "exposure_talent_num",
			Comment:            `曝光达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExposureTalentNum",
			GoFieldType:        "int32",
			JSONFieldName:      "exposure_talent_num",
			ProtobufFieldName:  "exposure_talent_num",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "commerce_talent_num",
			Comment:            `带货达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommerceTalentNum",
			GoFieldType:        "int32",
			JSONFieldName:      "commerce_talent_num",
			ProtobufFieldName:  "commerce_talent_num",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "valid_time",
			Comment:            `有效期,必须活动结束有效期`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ValidTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "valid_time",
			ProtobufFieldName:  "valid_time",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "estimated_revenue",
			Comment:            `预计收入`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedRevenue",
			GoFieldType:        "uint32",
			JSONFieldName:      "estimated_revenue",
			ProtobufFieldName:  "estimated_revenue",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "sell_num",
			Comment:            `活动销售数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SellNum",
			GoFieldType:        "int32",
			JSONFieldName:      "sell_num",
			ProtobufFieldName:  "sell_num",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "add_time",
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
			GoFieldName:        "AddTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ActivitySku) TableName() string {
	return "activity_sku"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ActivitySku) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ActivitySku) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ActivitySku) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ActivitySku) TableInfo() *TableInfo {
	return activity_skuTableInfo
}
