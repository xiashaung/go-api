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


CREATE TABLE `attr_info` (
  `attr_id` bigint(20) NOT NULL COMMENT '属性id',
  `cate_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类目id',
  `attr_name` varchar(10) NOT NULL COMMENT '属性名称',
  `attr_sort` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '属性排序',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '属性创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`attr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "attr_id": 74,    "cate_id": 44,    "attr_name": "OZEfTJfWZYrrrZMWtUsyuKvtu",    "attr_sort": 61,    "add_time": 44,    "last_time": 1}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// AttrInfo struct is a row record of the attr_info table in the yx database
type AttrInfo struct {
	//[ 0] attr_id                                        bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	AttrID int64 `gorm:"primary_key;column:attr_id;type:bigint;" json:"attr_id"` // 属性id
	//[ 1] cate_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateID uint32 `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"` // 类目id
	//[ 2] attr_name                                      varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AttrName string `gorm:"column:attr_name;type:varchar;size:10;" json:"attr_name"` // 属性名称
	//[ 3] attr_sort                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AttrSort uint32 `gorm:"column:attr_sort;type:uint;default:0;" json:"attr_sort"` // 属性排序
	//[ 4] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 属性创建时间
	//[ 5] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var attr_infoTableInfo = &TableInfo{
	Name: "attr_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "attr_id",
			Comment:            `属性id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AttrID",
			GoFieldType:        "int64",
			JSONFieldName:      "attr_id",
			ProtobufFieldName:  "attr_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "cate_id",
			Comment:            `类目id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "attr_name",
			Comment:            `属性名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AttrName",
			GoFieldType:        "string",
			JSONFieldName:      "attr_name",
			ProtobufFieldName:  "attr_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "attr_sort",
			Comment:            `属性排序`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AttrSort",
			GoFieldType:        "uint32",
			JSONFieldName:      "attr_sort",
			ProtobufFieldName:  "attr_sort",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_time",
			Comment:            `属性创建时间`,
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
func (a *AttrInfo) TableName() string {
	return "attr_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AttrInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AttrInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AttrInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AttrInfo) TableInfo() *TableInfo {
	return attr_infoTableInfo
}
