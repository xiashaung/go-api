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


CREATE TABLE `cate_tb` (
  `cate_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `yx_cate_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '央选类目id',
  `cate_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类目名称',
  `cate_level` smallint(6) NOT NULL DEFAULT '0' COMMENT '类目级别：1，2，3级类目',
  `cate_pid` int(11) NOT NULL DEFAULT '0' COMMENT '父类目id',
  `cate_step` text NOT NULL COMMENT '层级类目信息',
  `is_leaf` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否叶子节点',
  `enable` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否有效',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `tips` text,
  `is_brand` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否是品牌',
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1497800250 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC

JSON Sample
-------------------------------------
{    "cate_id": 68,    "yx_cate_id": 20,    "cate_name": "pnmfZGDABcOrcKGIwVCUtCSLv",    "cate_level": 38,    "cate_pid": 28,    "cate_step": "EBiuiauTSvyrGZfebBwAGKsZC",    "is_leaf": 54,    "enable": 49,    "add_time": 78,    "last_time": 94,    "tips": "XeFQqDwLXVnvCbgNbxlUMvwmY",    "is_brand": 28}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned



*/

// CateTb struct is a row record of the cate_tb table in the yx database
type CateTb struct {
	//[ 0] cate_id                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CateID int32 `gorm:"primary_key;AUTO_INCREMENT;column:cate_id;type:int;" json:"cate_id"` // 主键
	//[ 1] yx_cate_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	YxCateID uint32 `gorm:"column:yx_cate_id;type:uint;default:0;" json:"yx_cate_id"` // 央选类目id
	//[ 2] cate_name                                      varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:32;" json:"cate_name"` // 类目名称
	//[ 3] cate_level                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	CateLevel int32 `gorm:"column:cate_level;type:smallint;default:0;" json:"cate_level"` // 类目级别：1，2，3级类目
	//[ 4] cate_pid                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CatePid int32 `gorm:"column:cate_pid;type:int;default:0;" json:"cate_pid"` // 父类目id
	//[ 5] cate_step                                      text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CateStep string `gorm:"column:cate_step;type:text;size:65535;" json:"cate_step"` // 层级类目信息
	//[ 6] is_leaf                                        smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsLeaf int32 `gorm:"column:is_leaf;type:smallint;default:0;" json:"is_leaf"` // 是否叶子节点
	//[ 7] enable                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Enable int32 `gorm:"column:enable;type:smallint;default:0;" json:"enable"` // 是否有效
	//[ 8] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 9] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 修改时间
	//[10] tips                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Tips sql.NullString `gorm:"column:tips;type:text;size:65535;" json:"tips"`
	//[11] is_brand                                       tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsBrand int32 `gorm:"column:is_brand;type:tinyint;default:0;" json:"is_brand"` // 是否是品牌

}

var cate_tbTableInfo = &TableInfo{
	Name: "cate_tb",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "cate_id",
			Comment:            `主键`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "yx_cate_id",
			Comment:            `央选类目id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "YxCateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "yx_cate_id",
			ProtobufFieldName:  "yx_cate_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_name",
			Comment:            `类目名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "CateName",
			GoFieldType:        "string",
			JSONFieldName:      "cate_name",
			ProtobufFieldName:  "cate_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_level",
			Comment:            `类目级别：1，2，3级类目`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "cate_pid",
			Comment:            `父类目id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CatePid",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_pid",
			ProtobufFieldName:  "cate_pid",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "cate_step",
			Comment:            `层级类目信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
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
			Comment:            `是否叶子节点`,
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
			Name:               "enable",
			Comment:            `是否有效`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Enable",
			GoFieldType:        "int32",
			JSONFieldName:      "enable",
			ProtobufFieldName:  "enable",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "last_time",
			Comment:            `修改时间`,
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

		&ColumnInfo{
			Index:              10,
			Name:               "tips",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Tips",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tips",
			ProtobufFieldName:  "tips",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "is_brand",
			Comment:            `是否是品牌`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsBrand",
			GoFieldType:        "int32",
			JSONFieldName:      "is_brand",
			ProtobufFieldName:  "is_brand",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CateTb) TableName() string {
	return "cate_tb"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CateTb) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CateTb) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CateTb) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CateTb) TableInfo() *TableInfo {
	return cate_tbTableInfo
}
