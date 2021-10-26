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


CREATE TABLE `statistics_temp` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1.昨日2.本周3.上周4.本月5.累计',
  `sale_talent` int(8) NOT NULL DEFAULT '0' COMMENT '动销达人数',
  `sale_merchant` int(8) NOT NULL DEFAULT '0' COMMENT '动销商家数',
  `sale_product` int(8) NOT NULL DEFAULT '0' COMMENT '动销商品数',
  `sale` int(8) NOT NULL DEFAULT '0' COMMENT '成交销量',
  `gmv` int(11) NOT NULL DEFAULT '0' COMMENT '成交GMV（分）',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='运营统计表'

JSON Sample
-------------------------------------
{    "id": 14,    "type": 3,    "sale_talent": 3,    "sale_merchant": 63,    "sale_product": 90,    "sale": 53,    "gmv": 25,    "add_time": 51,    "last_time": 97}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// StatisticsTemp struct is a row record of the statistics_temp table in the yx database
type StatisticsTemp struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] type                                           tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	Type int32 `gorm:"column:type;type:tinyint;default:0;" json:"type"` // 1.昨日2.本周3.上周4.本月5.累计
	//[ 2] sale_talent                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SaleTalent int32 `gorm:"column:sale_talent;type:int;default:0;" json:"sale_talent"` // 动销达人数
	//[ 3] sale_merchant                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SaleMerchant int32 `gorm:"column:sale_merchant;type:int;default:0;" json:"sale_merchant"` // 动销商家数
	//[ 4] sale_product                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SaleProduct int32 `gorm:"column:sale_product;type:int;default:0;" json:"sale_product"` // 动销商品数
	//[ 5] sale                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Sale int32 `gorm:"column:sale;type:int;default:0;" json:"sale"` // 成交销量
	//[ 6] gmv                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Gmv int32 `gorm:"column:gmv;type:int;default:0;" json:"gmv"` // 成交GMV（分）
	//[ 7] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加时间
	//[ 8] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 更新时间

}

var statistics_tempTableInfo = &TableInfo{
	Name: "statistics_temp",
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
			Name:               "type",
			Comment:            `1.昨日2.本周3.上周4.本月5.累计`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "sale_talent",
			Comment:            `动销达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_talent",
			ProtobufFieldName:  "sale_talent",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "sale_merchant",
			Comment:            `动销商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_merchant",
			ProtobufFieldName:  "sale_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "sale_product",
			Comment:            `动销商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleProduct",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_product",
			ProtobufFieldName:  "sale_product",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "sale",
			Comment:            `成交销量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Sale",
			GoFieldType:        "int32",
			JSONFieldName:      "sale",
			ProtobufFieldName:  "sale",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "gmv",
			Comment:            `成交GMV（分）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Gmv",
			GoFieldType:        "int32",
			JSONFieldName:      "gmv",
			ProtobufFieldName:  "gmv",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "add_time",
			Comment:            `添加时间`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "last_time",
			Comment:            `更新时间`,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *StatisticsTemp) TableName() string {
	return "statistics_temp"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *StatisticsTemp) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *StatisticsTemp) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *StatisticsTemp) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *StatisticsTemp) TableInfo() *TableInfo {
	return statistics_tempTableInfo
}
