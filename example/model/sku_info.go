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


CREATE TABLE `sku_info` (
  `sku_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品sku信息表',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `sku_prop` text NOT NULL COMMENT 'sku的属性信息',
  `sku_price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品价格，单位：分',
  `sku_sale_price` int(10) NOT NULL COMMENT '商品售卖价',
  `sku_stock` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '商品库存数量',
  `sku_status` int(255) NOT NULL COMMENT '商品状态',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建sku时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "sku_id": 67,    "product_id": 20,    "sku_prop": "gtgIUmpdhJaVHkhjwghIMVUwW",    "sku_price": 86,    "sku_sale_price": 69,    "sku_stock": 85,    "sku_status": 13,    "add_time": 89,    "last_time": 39}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 5] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned



*/

// SkuInfo struct is a row record of the sku_info table in the yx database
type SkuInfo struct {
	//[ 0] sku_id                                         ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	SkuID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:sku_id;type:ubigint;" json:"sku_id"` // 商品sku信息表
	//[ 1] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品id
	//[ 2] sku_prop                                       text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	SkuProp string `gorm:"column:sku_prop;type:text;size:65535;" json:"sku_prop"` // sku的属性信息
	//[ 3] sku_price                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SkuPrice uint32 `gorm:"column:sku_price;type:uint;default:0;" json:"sku_price"` // 商品价格，单位：分
	//[ 4] sku_sale_price                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SkuSalePrice int32 `gorm:"column:sku_sale_price;type:int;" json:"sku_sale_price"` // 商品售卖价
	//[ 5] sku_stock                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SkuStock uint32 `gorm:"column:sku_stock;type:uint;default:0;" json:"sku_stock"` // 商品库存数量
	//[ 6] sku_status                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SkuStatus int32 `gorm:"column:sku_status;type:int;" json:"sku_status"` // 商品状态
	//[ 7] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建sku时间
	//[ 8] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var sku_infoTableInfo = &TableInfo{
	Name: "sku_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "sku_id",
			Comment:            `商品sku信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SkuID",
			GoFieldType:        "uint64",
			JSONFieldName:      "sku_id",
			ProtobufFieldName:  "sku_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "product_id",
			Comment:            `商品id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "sku_prop",
			Comment:            `sku的属性信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "SkuProp",
			GoFieldType:        "string",
			JSONFieldName:      "sku_prop",
			ProtobufFieldName:  "sku_prop",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "sku_price",
			Comment:            `商品价格，单位：分`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SkuPrice",
			GoFieldType:        "uint32",
			JSONFieldName:      "sku_price",
			ProtobufFieldName:  "sku_price",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "sku_sale_price",
			Comment:            `商品售卖价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SkuSalePrice",
			GoFieldType:        "int32",
			JSONFieldName:      "sku_sale_price",
			ProtobufFieldName:  "sku_sale_price",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "sku_stock",
			Comment:            `商品库存数量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SkuStock",
			GoFieldType:        "uint32",
			JSONFieldName:      "sku_stock",
			ProtobufFieldName:  "sku_stock",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "sku_status",
			Comment:            `商品状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SkuStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "sku_status",
			ProtobufFieldName:  "sku_status",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "add_time",
			Comment:            `创建sku时间`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SkuInfo) TableName() string {
	return "sku_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SkuInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SkuInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SkuInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SkuInfo) TableInfo() *TableInfo {
	return sku_infoTableInfo
}
