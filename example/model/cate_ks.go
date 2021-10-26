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


CREATE TABLE `cate_ks` (
  `cate_id` bigint(11) NOT NULL COMMENT '快手类目信息表',
  `yx_cate_id` bigint(11) NOT NULL COMMENT '央选类目id',
  `cate_pid` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级类目id',
  `cate_name` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类目名称',
  `cate_level` smallint(11) NOT NULL DEFAULT '0' COMMENT '类目层级',
  `cate_step` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]',
  `is_leaf` smallint(11) NOT NULL COMMENT '是否叶子节点 0：非叶子节点 1：叶子节点',
  `refund_rule` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类目退款规则列表，多个规则以英文逗号分隔',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "cate_id": 22,    "yx_cate_id": 12,    "cate_pid": 87,    "cate_name": "kURDMxyHxqiTfjCmZfxwcfHWZ",    "cate_level": 67,    "cate_step": "MpgPJVtBOfxFbtWPAwdKIxdEy",    "is_leaf": 20,    "refund_rule": "ZmJUbESCaAWvGxvttcbBahfaK",    "add_time": 11,    "last_time": 76}


Comments
-------------------------------------
[ 2] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned



*/

// CateKs struct is a row record of the cate_ks table in the yx database
type CateKs struct {
	//[ 0] cate_id                                        bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	CateID int64 `gorm:"primary_key;column:cate_id;type:bigint;" json:"cate_id"` // 快手类目信息表
	//[ 1] yx_cate_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	YxCateID int64 `gorm:"column:yx_cate_id;type:bigint;" json:"yx_cate_id"` // 央选类目id
	//[ 2] cate_pid                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	CatePid uint64 `gorm:"column:cate_pid;type:ubigint;default:0;" json:"cate_pid"` // 父级类目id
	//[ 3] cate_name                                      varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:125;" json:"cate_name"` // 类目名称
	//[ 4] cate_level                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	CateLevel int32 `gorm:"column:cate_level;type:smallint;default:0;" json:"cate_level"` // 类目层级
	//[ 5] cate_step                                      varchar(2000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2000    default: []
	CateStep string `gorm:"column:cate_step;type:varchar;size:2000;" json:"cate_step"` // 一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]
	//[ 6] is_leaf                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: []
	IsLeaf int32 `gorm:"column:is_leaf;type:smallint;" json:"is_leaf"` // 是否叶子节点 0：非叶子节点 1：叶子节点
	//[ 7] refund_rule                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	RefundRule string `gorm:"column:refund_rule;type:varchar;size:255;" json:"refund_rule"` // 类目退款规则列表，多个规则以英文逗号分隔
	//[ 8] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[ 9] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var cate_ksTableInfo = &TableInfo{
	Name: "cate_ks",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "cate_id",
			Comment:            `快手类目信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "int64",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "yx_cate_id",
			Comment:            `央选类目id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "YxCateID",
			GoFieldType:        "int64",
			JSONFieldName:      "yx_cate_id",
			ProtobufFieldName:  "yx_cate_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_pid",
			Comment:            `父级类目id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "CatePid",
			GoFieldType:        "uint64",
			JSONFieldName:      "cate_pid",
			ProtobufFieldName:  "cate_pid",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_name",
			Comment:            `类目名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "CateName",
			GoFieldType:        "string",
			JSONFieldName:      "cate_name",
			ProtobufFieldName:  "cate_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "cate_level",
			Comment:            `类目层级`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "CateLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_level",
			ProtobufFieldName:  "cate_level",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "cate_step",
			Comment:            `一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2000,
			GoFieldName:        "CateStep",
			GoFieldType:        "string",
			JSONFieldName:      "cate_step",
			ProtobufFieldName:  "cate_step",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "is_leaf",
			Comment:            `是否叶子节点 0：非叶子节点 1：叶子节点`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsLeaf",
			GoFieldType:        "int32",
			JSONFieldName:      "is_leaf",
			ProtobufFieldName:  "is_leaf",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "refund_rule",
			Comment:            `类目退款规则列表，多个规则以英文逗号分隔`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "RefundRule",
			GoFieldType:        "string",
			JSONFieldName:      "refund_rule",
			ProtobufFieldName:  "refund_rule",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
func (c *CateKs) TableName() string {
	return "cate_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CateKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CateKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CateKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CateKs) TableInfo() *TableInfo {
	return cate_ksTableInfo
}
