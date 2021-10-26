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


CREATE TABLE `commission_info` (
  `commission_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '定向佣金信息表',
  `shop_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '店铺id',
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `talent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `anchor_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '主播id',
  `commission_rate` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '佣金率',
  `commission_start` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '佣金开始时间',
  `commission_end` int(255) NOT NULL COMMENT '佣金结束时间',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建佣金时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`commission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "commission_id": 47,    "shop_id": 81,    "product_id": 51,    "talent_id": 58,    "anchor_id": 31,    "commission_rate": 5,    "commission_start": 18,    "commission_end": 3,    "add_time": 14,    "last_time": 18}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned



*/

// CommissionInfo struct is a row record of the commission_info table in the yx database
type CommissionInfo struct {
	//[ 0] commission_id                                  bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	CommissionID int64 `gorm:"primary_key;AUTO_INCREMENT;column:commission_id;type:bigint;" json:"commission_id"` // 定向佣金信息表
	//[ 1] shop_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ShopID uint32 `gorm:"column:shop_id;type:uint;default:0;" json:"shop_id"` // 店铺id
	//[ 2] product_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ProductID uint32 `gorm:"column:product_id;type:uint;default:0;" json:"product_id"` // 商品id
	//[ 3] talent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TalentID uint32 `gorm:"column:talent_id;type:uint;default:0;" json:"talent_id"` // 达人id
	//[ 4] anchor_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AnchorID uint32 `gorm:"column:anchor_id;type:uint;default:0;" json:"anchor_id"` // 主播id
	//[ 5] commission_rate                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CommissionRate uint32 `gorm:"column:commission_rate;type:uint;default:0;" json:"commission_rate"` // 佣金率
	//[ 6] commission_start                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CommissionStart uint32 `gorm:"column:commission_start;type:uint;default:0;" json:"commission_start"` // 佣金开始时间
	//[ 7] commission_end                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CommissionEnd int32 `gorm:"column:commission_end;type:int;" json:"commission_end"` // 佣金结束时间
	//[ 8] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建佣金时间
	//[ 9] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var commission_infoTableInfo = &TableInfo{
	Name: "commission_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "commission_id",
			Comment:            `定向佣金信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CommissionID",
			GoFieldType:        "int64",
			JSONFieldName:      "commission_id",
			ProtobufFieldName:  "commission_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "product_id",
			Comment:            `商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint32",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "uint32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "commission_rate",
			Comment:            `佣金率`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "commission_start",
			Comment:            `佣金开始时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CommissionStart",
			GoFieldType:        "uint32",
			JSONFieldName:      "commission_start",
			ProtobufFieldName:  "commission_start",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "commission_end",
			Comment:            `佣金结束时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommissionEnd",
			GoFieldType:        "int32",
			JSONFieldName:      "commission_end",
			ProtobufFieldName:  "commission_end",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "add_time",
			Comment:            `创建佣金时间`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CommissionInfo) TableName() string {
	return "commission_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CommissionInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CommissionInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CommissionInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CommissionInfo) TableInfo() *TableInfo {
	return commission_infoTableInfo
}
