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


CREATE TABLE `rule_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `service_type` varchar(255) NOT NULL DEFAULT '' COMMENT '业务类型',
  `object` varchar(255) NOT NULL DEFAULT '' COMMENT '操作对象',
  `status` tinyint(3) DEFAULT '1' COMMENT '状态值，0关闭，1正常启用',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='名单业务类型'

JSON Sample
-------------------------------------
{    "id": 77,    "service_type": "cDXqPQwxOtnLCHxohkUEOAWKS",    "object": "WelNKyQHahGYdrTNODaeHdduH",    "status": 18,    "add_time": 60,    "last_time": 7}


Comments
-------------------------------------
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// RuleInfo struct is a row record of the rule_info table in the yx database
type RuleInfo struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"` // 自增id
	//[ 1] service_type                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ServiceType string `gorm:"column:service_type;type:varchar;size:255;" json:"service_type"` // 业务类型
	//[ 2] object                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Object string `gorm:"column:object;type:varchar;size:255;" json:"object"` // 操作对象
	//[ 3] status                                         tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Status sql.NullInt64 `gorm:"column:status;type:tinyint;default:1;" json:"status"` // 状态值，0关闭，1正常启用
	//[ 4] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 5] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var rule_infoTableInfo = &TableInfo{
	Name: "rule_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `自增id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "service_type",
			Comment:            `业务类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ServiceType",
			GoFieldType:        "string",
			JSONFieldName:      "service_type",
			ProtobufFieldName:  "service_type",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "object",
			Comment:            `操作对象`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Object",
			GoFieldType:        "string",
			JSONFieldName:      "object",
			ProtobufFieldName:  "object",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
			Comment:            `状态值，0关闭，1正常启用`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_time",
			Comment:            `添加时间`,
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
func (r *RuleInfo) TableName() string {
	return "rule_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RuleInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RuleInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RuleInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RuleInfo) TableInfo() *TableInfo {
	return rule_infoTableInfo
}
