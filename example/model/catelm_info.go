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


CREATE TABLE `catelm_info` (
  `cate_id` int(255) unsigned NOT NULL AUTO_INCREMENT COMMENT '前台类目信息id',
  `cate_pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级类目id',
  `cate_name` varchar(255) NOT NULL COMMENT '类目名称',
  `cate_level` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '类目层级',
  `cate_step` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]',
  `cate_weight` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '显示权重',
  `is_leaf` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否叶子节点 0：非叶子节点 1：叶子节点',
  `root_cate_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '根节点id',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最近修改时间',
  `cate_img` text NOT NULL COMMENT '类目图片',
  PRIMARY KEY (`cate_id`),
  KEY `idx_catepid` (`cate_pid`) USING BTREE,
  KEY `idx_rootcateid` (`root_cate_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='前台类目信息'

JSON Sample
-------------------------------------
{    "cate_id": 63,    "cate_pid": 70,    "cate_name": "OoYKZjRaXmrZutBEsJpyxNqOu",    "cate_level": 22,    "cate_step": "YpUuSubFwabKpEfDjXDtpeKYx",    "cate_weight": 62,    "is_leaf": 38,    "root_cate_id": 5,    "add_time": 56,    "last_time": 73,    "cate_img": "BFFHxedVpomJTcmKZXNfFSycJ"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned



*/

// CatelmInfo struct is a row record of the catelm_info table in the yx database
type CatelmInfo struct {
	//[ 0] cate_id                                        uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	CateID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:cate_id;type:uint;" json:"cate_id"` // 前台类目信息id
	//[ 1] cate_pid                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CatePid uint32 `gorm:"column:cate_pid;type:uint;default:0;" json:"cate_pid"` // 父级类目id
	//[ 2] cate_name                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:255;" json:"cate_name"` // 类目名称
	//[ 3] cate_level                                     utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	CateLevel uint32 `gorm:"column:cate_level;type:utinyint;default:1;" json:"cate_level"` // 类目层级
	//[ 4] cate_step                                      text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CateStep string `gorm:"column:cate_step;type:text;size:65535;" json:"cate_step"` // 一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]
	//[ 5] cate_weight                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateWeight uint32 `gorm:"column:cate_weight;type:uint;default:0;" json:"cate_weight"` // 显示权重
	//[ 6] is_leaf                                        utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsLeaf uint32 `gorm:"column:is_leaf;type:utinyint;default:0;" json:"is_leaf"` // 是否叶子节点 0：非叶子节点 1：叶子节点
	//[ 7] root_cate_id                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RootCateID uint32 `gorm:"column:root_cate_id;type:uint;default:0;" json:"root_cate_id"` // 根节点id
	//[ 8] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 9] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 最近修改时间
	//[10] cate_img                                       text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CateImg string `gorm:"column:cate_img;type:text;size:65535;" json:"cate_img"` // 类目图片

}

var catelm_infoTableInfo = &TableInfo{
	Name: "catelm_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "cate_id",
			Comment:            `前台类目信息id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "cate_pid",
			Comment:            `父级类目id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CatePid",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_pid",
			ProtobufFieldName:  "cate_pid",
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
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
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
			Comment:            `类目层级`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "CateLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_level",
			ProtobufFieldName:  "cate_level",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "cate_step",
			Comment:            `一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "cate_weight",
			Comment:            `显示权重`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateWeight",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_weight",
			ProtobufFieldName:  "cate_weight",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "is_leaf",
			Comment:            `是否叶子节点 0：非叶子节点 1：叶子节点`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsLeaf",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_leaf",
			ProtobufFieldName:  "is_leaf",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "root_cate_id",
			Comment:            `根节点id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "RootCateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "root_cate_id",
			ProtobufFieldName:  "root_cate_id",
			ProtobufType:       "uint32",
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
			Comment:            `最近修改时间`,
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
			Name:               "cate_img",
			Comment:            `类目图片`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "CateImg",
			GoFieldType:        "string",
			JSONFieldName:      "cate_img",
			ProtobufFieldName:  "cate_img",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CatelmInfo) TableName() string {
	return "catelm_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CatelmInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CatelmInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CatelmInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CatelmInfo) TableInfo() *TableInfo {
	return catelm_infoTableInfo
}
