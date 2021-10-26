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


CREATE TABLE `shop_rule` (
  `shop_rule_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '店铺规则表',
  `is_free_sample` int(10) NOT NULL DEFAULT '0' COMMENT '是否免费申请样品',
  `merchant_id` int(10) DEFAULT NULL COMMENT '账号id',
  `platform_id` int(10) NOT NULL COMMENT '平台类型',
  `shop_id` bigint(20) NOT NULL COMMENT '店铺id',
  `day_type` int(10) NOT NULL DEFAULT '30' COMMENT '时间类型',
  `fans_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝数量门槛 0：代表无门槛',
  `sale_num` int(255) NOT NULL COMMENT '30天商品销量',
  `sale_fee` int(255) NOT NULL COMMENT '30天销售额 扩大1000倍存储',
  `free_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '免审商品数量 0：需要手动审核',
  `free_used` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '已经自动审核的数量',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`shop_rule_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "shop_rule_id": 72,    "is_free_sample": 60,    "merchant_id": 80,    "platform_id": 24,    "shop_id": 88,    "day_type": 18,    "fans_num": 31,    "sale_num": 7,    "sale_fee": 98,    "free_num": 7,    "free_used": 39,    "add_time": 31,    "last_time": 81}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 6] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned



*/

// ShopRule struct is a row record of the shop_rule table in the yx database
type ShopRule struct {
	//[ 0] shop_rule_id                                   ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ShopRuleID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:shop_rule_id;type:ubigint;" json:"shop_rule_id"` // 店铺规则表
	//[ 1] is_free_sample                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsFreeSample int32 `gorm:"column:is_free_sample;type:int;default:0;" json:"is_free_sample"` // 是否免费申请样品
	//[ 2] merchant_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MerchantID sql.NullInt64 `gorm:"column:merchant_id;type:int;" json:"merchant_id"` // 账号id
	//[ 3] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PlatformID int32 `gorm:"column:platform_id;type:int;" json:"platform_id"` // 平台类型
	//[ 4] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	ShopID int64 `gorm:"column:shop_id;type:bigint;" json:"shop_id"` // 店铺id
	//[ 5] day_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [30]
	DayType int32 `gorm:"column:day_type;type:int;default:30;" json:"day_type"` // 时间类型
	//[ 6] fans_num                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FansNum uint32 `gorm:"column:fans_num;type:uint;default:0;" json:"fans_num"` // 粉丝数量门槛 0：代表无门槛
	//[ 7] sale_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SaleNum int32 `gorm:"column:sale_num;type:int;" json:"sale_num"` // 30天商品销量
	//[ 8] sale_fee                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SaleFee int32 `gorm:"column:sale_fee;type:int;" json:"sale_fee"` // 30天销售额 扩大1000倍存储
	//[ 9] free_num                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FreeNum uint32 `gorm:"column:free_num;type:uint;default:0;" json:"free_num"` // 免审商品数量 0：需要手动审核
	//[10] free_used                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FreeUsed uint32 `gorm:"column:free_used;type:uint;default:0;" json:"free_used"` // 已经自动审核的数量
	//[11] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"`
	//[12] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var shop_ruleTableInfo = &TableInfo{
	Name: "shop_rule",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "shop_rule_id",
			Comment:            `店铺规则表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ShopRuleID",
			GoFieldType:        "uint64",
			JSONFieldName:      "shop_rule_id",
			ProtobufFieldName:  "shop_rule_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "is_free_sample",
			Comment:            `是否免费申请样品`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "merchant_id",
			Comment:            `账号id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "shop_id",
			Comment:            `店铺id`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "day_type",
			Comment:            `时间类型`,
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
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleNum",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_num",
			ProtobufFieldName:  "sale_num",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "sale_fee",
			Comment:            `30天销售额 扩大1000倍存储`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleFee",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_fee",
			ProtobufFieldName:  "sale_fee",
			ProtobufType:       "int32",
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
func (s *ShopRule) TableName() string {
	return "shop_rule"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShopRule) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShopRule) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShopRule) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShopRule) TableInfo() *TableInfo {
	return shop_ruleTableInfo
}
