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


CREATE TABLE `product_attr` (
  `id` bigint(20) NOT NULL COMMENT '商品属性关联表',
  `product_id` bigint(20) NOT NULL COMMENT '商品id',
  `attr_id` bigint(20) NOT NULL COMMENT '商品属性id',
  `attr_name` varchar(20) NOT NULL COMMENT '属性名称',
  `attr_value_id` bigint(20) NOT NULL COMMENT '属性值id',
  `attr_value` varchar(20) NOT NULL COMMENT '属性值名称',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 69,    "product_id": 44,    "attr_id": 29,    "attr_name": "DyYnSpQdvhbhxeIRpCwoVakuh",    "attr_value_id": 80,    "attr_value": "pTpNJgHsfiJcVlPRMEHfCItrH",    "add_time": 85,    "last_time": 46}


Comments
-------------------------------------
[ 6] column is set for unsigned
[ 7] column is set for unsigned



*/

// ProductAttr struct is a row record of the product_attr table in the yx database
type ProductAttr struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"` // 商品属性关联表
	//[ 1] product_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	ProductID int64 `gorm:"column:product_id;type:bigint;" json:"product_id"` // 商品id
	//[ 2] attr_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AttrID int64 `gorm:"column:attr_id;type:bigint;" json:"attr_id"` // 商品属性id
	//[ 3] attr_name                                      varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AttrName string `gorm:"column:attr_name;type:varchar;size:20;" json:"attr_name"` // 属性名称
	//[ 4] attr_value_id                                  bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AttrValueID int64 `gorm:"column:attr_value_id;type:bigint;" json:"attr_value_id"` // 属性值id
	//[ 5] attr_value                                     varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AttrValue string `gorm:"column:attr_value;type:varchar;size:20;" json:"attr_value"` // 属性值名称
	//[ 6] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[ 7] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var product_attrTableInfo = &TableInfo{
	Name: "product_attr",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `商品属性关联表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "product_id",
			Comment:            `商品id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "attr_id",
			Comment:            `商品属性id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AttrID",
			GoFieldType:        "int64",
			JSONFieldName:      "attr_id",
			ProtobufFieldName:  "attr_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "attr_name",
			Comment:            `属性名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "AttrName",
			GoFieldType:        "string",
			JSONFieldName:      "attr_name",
			ProtobufFieldName:  "attr_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "attr_value_id",
			Comment:            `属性值id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AttrValueID",
			GoFieldType:        "int64",
			JSONFieldName:      "attr_value_id",
			ProtobufFieldName:  "attr_value_id",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "attr_value",
			Comment:            `属性值名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "AttrValue",
			GoFieldType:        "string",
			JSONFieldName:      "attr_value",
			ProtobufFieldName:  "attr_value",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProductAttr) TableName() string {
	return "product_attr"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductAttr) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductAttr) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductAttr) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductAttr) TableInfo() *TableInfo {
	return product_attrTableInfo
}
