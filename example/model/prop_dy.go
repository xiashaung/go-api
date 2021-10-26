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


CREATE TABLE `prop_dy` (
  `id` bigint(20) NOT NULL COMMENT '抖音prop信息表',
  `name` varchar(20) NOT NULL COMMENT 'prop名称',
  `pid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '父级节点id',
  `is_leaf` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '是否叶子节点',
  `status` int(255) unsigned NOT NULL DEFAULT '0' COMMENT 'prop状态',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 90,    "name": "kMNmtZfroGUuVCppKcJueyWIt",    "pid": 98,    "is_leaf": 17,    "status": 6,    "add_time": 58,    "last_time": 40}


Comments
-------------------------------------
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// PropDy struct is a row record of the prop_dy table in the yx database
type PropDy struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"` // 抖音prop信息表
	//[ 1] name                                           varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Name string `gorm:"column:name;type:varchar;size:20;" json:"name"` // prop名称
	//[ 2] pid                                            ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	Pid uint64 `gorm:"column:pid;type:ubigint;default:0;" json:"pid"` // 父级节点id
	//[ 3] is_leaf                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	IsLeaf uint32 `gorm:"column:is_leaf;type:uint;default:0;" json:"is_leaf"` // 是否叶子节点
	//[ 4] status                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Status uint32 `gorm:"column:status;type:uint;default:0;" json:"status"` // prop状态
	//[ 5] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[ 6] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var prop_dyTableInfo = &TableInfo{
	Name: "prop_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `抖音prop信息表`,
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
			Name:               "name",
			Comment:            `prop名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "pid",
			Comment:            `父级节点id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "Pid",
			GoFieldType:        "uint64",
			JSONFieldName:      "pid",
			ProtobufFieldName:  "pid",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "is_leaf",
			Comment:            `是否叶子节点`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "IsLeaf",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_leaf",
			ProtobufFieldName:  "is_leaf",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "status",
			Comment:            `prop状态`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "uint32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PropDy) TableName() string {
	return "prop_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PropDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PropDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PropDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PropDy) TableInfo() *TableInfo {
	return prop_dyTableInfo
}
