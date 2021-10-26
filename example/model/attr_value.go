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


CREATE TABLE `attr_value` (
  `id` bigint(255) NOT NULL COMMENT '属性值信息表',
  `attr_id` bigint(20) NOT NULL COMMENT '属性id',
  `attr_value` varchar(20) NOT NULL COMMENT '属性值',
  `attr_type` int(255) NOT NULL COMMENT '属性值类型 1：系统 2：自定义',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '属性值创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 80,    "attr_id": 45,    "attr_value": "gTTWlQDgTPustwCTTSdkkDWxu",    "attr_type": 28,    "add_time": 8,    "last_time": 81}


Comments
-------------------------------------
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// AttrValue struct is a row record of the attr_value table in the yx database
type AttrValue struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"` // 属性值信息表
	//[ 1] attr_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	AttrID int64 `gorm:"column:attr_id;type:bigint;" json:"attr_id"` // 属性id
	//[ 2] attr_value                                     varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AttrValue string `gorm:"column:attr_value;type:varchar;size:20;" json:"attr_value"` // 属性值
	//[ 3] attr_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AttrType int32 `gorm:"column:attr_type;type:int;" json:"attr_type"` // 属性值类型 1：系统 2：自定义
	//[ 4] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 属性值创建时间
	//[ 5] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var attr_valueTableInfo = &TableInfo{
	Name: "attr_value",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `属性值信息表`,
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
			Name:               "attr_id",
			Comment:            `属性id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "attr_value",
			Comment:            `属性值`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "attr_type",
			Comment:            `属性值类型 1：系统 2：自定义`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AttrType",
			GoFieldType:        "int32",
			JSONFieldName:      "attr_type",
			ProtobufFieldName:  "attr_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_time",
			Comment:            `属性值创建时间`,
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
func (a *AttrValue) TableName() string {
	return "attr_value"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AttrValue) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AttrValue) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AttrValue) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AttrValue) TableInfo() *TableInfo {
	return attr_valueTableInfo
}
