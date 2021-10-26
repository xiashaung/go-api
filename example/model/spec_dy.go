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


CREATE TABLE `spec_dy` (
  `spec_id` bigint(20) NOT NULL COMMENT '规格信息表，只有抖音使用',
  `spec_name` varchar(30) NOT NULL COMMENT '规则名称',
  `spec_prop` text NOT NULL COMMENT '商品规格属性',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`spec_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "spec_id": 13,    "spec_name": "cfSHfliBlaCXExNOVRFDeYXvE",    "spec_prop": "ufHyuYIYOLoqpNIaOhCeOQqSE",    "add_time": 91,    "last_time": 51}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// SpecDy struct is a row record of the spec_dy table in the yx database
type SpecDy struct {
	//[ 0] spec_id                                        bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	SpecID int64 `gorm:"primary_key;column:spec_id;type:bigint;" json:"spec_id"` // 规格信息表，只有抖音使用
	//[ 1] spec_name                                      varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	SpecName string `gorm:"column:spec_name;type:varchar;size:30;" json:"spec_name"` // 规则名称
	//[ 2] spec_prop                                      text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	SpecProp string `gorm:"column:spec_prop;type:text;size:65535;" json:"spec_prop"` // 商品规格属性
	//[ 3] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"`
	//[ 4] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var spec_dyTableInfo = &TableInfo{
	Name: "spec_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "spec_id",
			Comment:            `规格信息表，只有抖音使用`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SpecID",
			GoFieldType:        "int64",
			JSONFieldName:      "spec_id",
			ProtobufFieldName:  "spec_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "spec_name",
			Comment:            `规则名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "SpecName",
			GoFieldType:        "string",
			JSONFieldName:      "spec_name",
			ProtobufFieldName:  "spec_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "spec_prop",
			Comment:            `商品规格属性`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "SpecProp",
			GoFieldType:        "string",
			JSONFieldName:      "spec_prop",
			ProtobufFieldName:  "spec_prop",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SpecDy) TableName() string {
	return "spec_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SpecDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SpecDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SpecDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SpecDy) TableInfo() *TableInfo {
	return spec_dyTableInfo
}
