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


CREATE TABLE `prop_value` (
  `id` bigint(20) NOT NULL COMMENT 'prop属性值信息表',
  `prop_id` bigint(20) NOT NULL COMMENT 'prop属性id',
  `prop_value` varchar(20) NOT NULL COMMENT '属性名称',
  `prop_type` int(255) unsigned NOT NULL DEFAULT '1' COMMENT '类目类型 1：系统	2：自定义',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'prop值添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 98,    "prop_id": 11,    "prop_value": "VtnALvmaXMmAMtUEMZlGIENDj",    "prop_type": 23,    "add_time": 30,    "last_time": 51}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// PropValue struct is a row record of the prop_value table in the yx database
type PropValue struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"` // prop属性值信息表
	//[ 1] prop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PropID int64 `gorm:"column:prop_id;type:bigint;" json:"prop_id"` // prop属性id
	//[ 2] prop_value                                     varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	PropValue string `gorm:"column:prop_value;type:varchar;size:20;" json:"prop_value"` // 属性名称
	//[ 3] prop_type                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	PropType uint32 `gorm:"column:prop_type;type:uint;default:1;" json:"prop_type"` // 类目类型 1：系统	2：自定义
	//[ 4] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // prop值添加时间
	//[ 5] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var prop_valueTableInfo = &TableInfo{
	Name: "prop_value",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `prop属性值信息表`,
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
			Name:               "prop_id",
			Comment:            `prop属性id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PropID",
			GoFieldType:        "int64",
			JSONFieldName:      "prop_id",
			ProtobufFieldName:  "prop_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "prop_value",
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
			GoFieldName:        "PropValue",
			GoFieldType:        "string",
			JSONFieldName:      "prop_value",
			ProtobufFieldName:  "prop_value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "prop_type",
			Comment:            `类目类型 1：系统	2：自定义`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PropType",
			GoFieldType:        "uint32",
			JSONFieldName:      "prop_type",
			ProtobufFieldName:  "prop_type",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_time",
			Comment:            `prop值添加时间`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PropValue) TableName() string {
	return "prop_value"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PropValue) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PropValue) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PropValue) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PropValue) TableInfo() *TableInfo {
	return prop_valueTableInfo
}
