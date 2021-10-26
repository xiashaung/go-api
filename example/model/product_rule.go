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


CREATE TABLE `product_rule` (
  `product_rule_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品样品规则表',
  `merchant_id` int(10) NOT NULL DEFAULT '0' COMMENT '账号id',
  `platform_id` int(10) NOT NULL COMMENT '平台类型',
  `product_id` bigint(20) NOT NULL COMMENT '产品id',
  `day_type` int(10) NOT NULL COMMENT '日期类型',
  `is_free_sample` int(10) NOT NULL DEFAULT '1' COMMENT '是否免费申样 0 否 1 是',
  `fans_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝数量门槛 0：代表无门槛',
  `sale_num` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '30天商品销量',
  `sale_fee` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '30天销售额 扩大1000倍存储',
  `free_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '免审商品数量 0：需要手动审核',
  `free_used` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '已经自动审核的数量',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`product_rule_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "product_rule_id": 95,    "merchant_id": 16,    "platform_id": 58,    "product_id": 38,    "day_type": 27,    "is_free_sample": 35,    "fans_num": 54,    "sale_num": 98,    "sale_fee": 7,    "free_num": 50,    "free_used": 35,    "add_time": 57,    "last_time": 6}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned



*/

// ProductRule struct is a row record of the product_rule table in the yx database
type ProductRule struct {
	//[ 0] product_rule_id                                ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ProductRuleID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:product_rule_id;type:ubigint;" json:"product_rule_id"` // 商品样品规则表
	//[ 1] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"` // 账号id
	//[ 2] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PlatformID int32 `gorm:"column:platform_id;type:int;" json:"platform_id"` // 平台类型
	//[ 3] product_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	ProductID int64 `gorm:"column:product_id;type:bigint;" json:"product_id"` // 产品id
	//[ 4] day_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DayType int32 `gorm:"column:day_type;type:int;" json:"day_type"` // 日期类型
	//[ 5] is_free_sample                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	IsFreeSample int32 `gorm:"column:is_free_sample;type:int;default:1;" json:"is_free_sample"` // 是否免费申样 0 否 1 是
	//[ 6] fans_num                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FansNum uint32 `gorm:"column:fans_num;type:uint;default:0;" json:"fans_num"` // 粉丝数量门槛 0：代表无门槛
	//[ 7] sale_num                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SaleNum uint32 `gorm:"column:sale_num;type:uint;default:0;" json:"sale_num"` // 30天商品销量
	//[ 8] sale_fee                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SaleFee uint32 `gorm:"column:sale_fee;type:uint;default:0;" json:"sale_fee"` // 30天销售额 扩大1000倍存储
	//[ 9] free_num                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FreeNum uint32 `gorm:"column:free_num;type:uint;default:0;" json:"free_num"` // 免审商品数量 0：需要手动审核
	//[10] free_used                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FreeUsed uint32 `gorm:"column:free_used;type:uint;default:0;" json:"free_used"` // 已经自动审核的数量
	//[11] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"`
	//[12] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var product_ruleTableInfo = &TableInfo{
	Name: "product_rule",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "product_rule_id",
			Comment:            `商品样品规则表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductRuleID",
			GoFieldType:        "uint64",
			JSONFieldName:      "product_rule_id",
			ProtobufFieldName:  "product_rule_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "merchant_id",
			Comment:            `账号id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "platform_id",
			Comment:            `平台类型`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_id",
			Comment:            `产品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "day_type",
			Comment:            `日期类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DayType",
			GoFieldType:        "int32",
			JSONFieldName:      "day_type",
			ProtobufFieldName:  "day_type",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_free_sample",
			Comment:            `是否免费申样 0 否 1 是`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "fans_num",
			Comment:            `粉丝数量门槛 0：代表无门槛`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "FansNum",
			GoFieldType:        "uint32",
			JSONFieldName:      "fans_num",
			ProtobufFieldName:  "fans_num",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sale_num",
			Comment:            `30天商品销量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SaleNum",
			GoFieldType:        "uint32",
			JSONFieldName:      "sale_num",
			ProtobufFieldName:  "sale_num",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "sale_fee",
			Comment:            `30天销售额 扩大1000倍存储`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SaleFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "sale_fee",
			ProtobufFieldName:  "sale_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "free_num",
			Comment:            `免审商品数量 0：需要手动审核`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "FreeNum",
			GoFieldType:        "uint32",
			JSONFieldName:      "free_num",
			ProtobufFieldName:  "free_num",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "free_used",
			Comment:            `已经自动审核的数量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "FreeUsed",
			GoFieldType:        "uint32",
			JSONFieldName:      "free_used",
			ProtobufFieldName:  "free_used",
			ProtobufType:       "uint32",
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
func (p *ProductRule) TableName() string {
	return "product_rule"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductRule) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductRule) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductRule) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductRule) TableInfo() *TableInfo {
	return product_ruleTableInfo
}
