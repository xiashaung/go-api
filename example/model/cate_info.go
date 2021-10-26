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


CREATE TABLE `cate_info` (
  `cate_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '类目信息表',
  `cate_pid` int(11) NOT NULL COMMENT '父级类目id',
  `cate_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '类目名称',
  `cate_level` int(11) NOT NULL COMMENT '类目层级',
  `cate_step` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]',
  `is_system` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1:系统字段0:非系统字段（自定义）',
  `weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重   排序按照权重的大小的正序',
  `is_leaf` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '是否叶子节点 0：非叶子节点 1：叶子节点',
  `root_cate_id` int(11) DEFAULT NULL COMMENT '根节点id',
  `platform_limit_dy` int(11) NOT NULL DEFAULT '0' COMMENT '平台扣点(抖音)',
  `platform_limit_ks` int(11) NOT NULL DEFAULT '0' COMMENT '平台扣点(快手)',
  `industry_commission` int(11) NOT NULL DEFAULT '0' COMMENT '行业佣金',
  `super_commission_min` int(11) NOT NULL DEFAULT '0' COMMENT '最小超佣比例',
  `super_commission_max` int(11) NOT NULL DEFAULT '0' COMMENT '最大超佣比例',
  `min_stock` int(11) NOT NULL DEFAULT '0' COMMENT '最小允许库存',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加类目时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`cate_id`) USING BTREE,
  KEY `cate_pid` (`cate_pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100099 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "cate_id": 28,    "cate_pid": 60,    "cate_name": "gtoFOPOKhcJpdVRJfYKNSKiTs",    "cate_level": 51,    "cate_step": "XBroKRHZbaaEiQMjhvJLLKxIu",    "is_system": 35,    "weight": 34,    "is_leaf": 68,    "root_cate_id": 71,    "platform_limit_dy": 40,    "platform_limit_ks": 93,    "industry_commission": 50,    "super_commission_min": 57,    "super_commission_max": 19,    "min_stock": 91,    "add_time": 8,    "last_time": 34}


Comments
-------------------------------------
[ 7] column is set for unsigned
[15] column is set for unsigned
[16] column is set for unsigned



*/

// CateInfo struct is a row record of the cate_info table in the yx database
type CateInfo struct {
	//[ 0] cate_id                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CateID int32 `gorm:"primary_key;AUTO_INCREMENT;column:cate_id;type:int;" json:"cate_id"` // 类目信息表
	//[ 1] cate_pid                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CatePid int32 `gorm:"column:cate_pid;type:int;" json:"cate_pid"` // 父级类目id
	//[ 2] cate_name                                      varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:50;" json:"cate_name"` // 类目名称
	//[ 3] cate_level                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CateLevel int32 `gorm:"column:cate_level;type:int;" json:"cate_level"` // 类目层级
	//[ 4] cate_step                                      text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CateStep string `gorm:"column:cate_step;type:text;size:65535;" json:"cate_step"` // 一级类目到当前类目的路径，例如：[1:"服装",11:"女装"]
	//[ 5] is_system                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsSystem int32 `gorm:"column:is_system;type:tinyint;default:0;" json:"is_system"` // 1:系统字段0:非系统字段（自定义）
	//[ 6] weight                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Weight int32 `gorm:"column:weight;type:int;default:0;" json:"weight"` // 权重   排序按照权重的大小的正序
	//[ 7] is_leaf                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	IsLeaf uint32 `gorm:"column:is_leaf;type:uint;default:0;" json:"is_leaf"` // 是否叶子节点 0：非叶子节点 1：叶子节点
	//[ 8] root_cate_id                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RootCateID sql.NullInt64 `gorm:"column:root_cate_id;type:int;" json:"root_cate_id"` // 根节点id
	//[ 9] platform_limit_dy                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformLimitDy int32 `gorm:"column:platform_limit_dy;type:int;default:0;" json:"platform_limit_dy"` // 平台扣点(抖音)
	//[10] platform_limit_ks                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformLimitKs int32 `gorm:"column:platform_limit_ks;type:int;default:0;" json:"platform_limit_ks"` // 平台扣点(快手)
	//[11] industry_commission                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IndustryCommission int32 `gorm:"column:industry_commission;type:int;default:0;" json:"industry_commission"` // 行业佣金
	//[12] super_commission_min                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SuperCommissionMin int32 `gorm:"column:super_commission_min;type:int;default:0;" json:"super_commission_min"` // 最小超佣比例
	//[13] super_commission_max                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SuperCommissionMax int32 `gorm:"column:super_commission_max;type:int;default:0;" json:"super_commission_max"` // 最大超佣比例
	//[14] min_stock                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MinStock int32 `gorm:"column:min_stock;type:int;default:0;" json:"min_stock"` // 最小允许库存
	//[15] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加类目时间
	//[16] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var cate_infoTableInfo = &TableInfo{
	Name: "cate_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "cate_id",
			Comment:            `类目信息表`,
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
			Name:               "cate_pid",
			Comment:            `父级类目id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_name",
			Comment:            `类目名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
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
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
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
			Name:               "is_system",
			Comment:            `1:系统字段0:非系统字段（自定义）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsSystem",
			GoFieldType:        "int32",
			JSONFieldName:      "is_system",
			ProtobufFieldName:  "is_system",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "weight",
			Comment:            `权重   排序按照权重的大小的正序`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Weight",
			GoFieldType:        "int32",
			JSONFieldName:      "weight",
			ProtobufFieldName:  "weight",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "is_leaf",
			Comment:            `是否叶子节点 0：非叶子节点 1：叶子节点`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "root_cate_id",
			Comment:            `根节点id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RootCateID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "root_cate_id",
			ProtobufFieldName:  "root_cate_id",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "platform_limit_dy",
			Comment:            `平台扣点(抖音)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformLimitDy",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_limit_dy",
			ProtobufFieldName:  "platform_limit_dy",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "platform_limit_ks",
			Comment:            `平台扣点(快手)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformLimitKs",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_limit_ks",
			ProtobufFieldName:  "platform_limit_ks",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "industry_commission",
			Comment:            `行业佣金`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IndustryCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "industry_commission",
			ProtobufFieldName:  "industry_commission",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "super_commission_min",
			Comment:            `最小超佣比例`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SuperCommissionMin",
			GoFieldType:        "int32",
			JSONFieldName:      "super_commission_min",
			ProtobufFieldName:  "super_commission_min",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "super_commission_max",
			Comment:            `最大超佣比例`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SuperCommissionMax",
			GoFieldType:        "int32",
			JSONFieldName:      "super_commission_max",
			ProtobufFieldName:  "super_commission_max",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "min_stock",
			Comment:            `最小允许库存`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MinStock",
			GoFieldType:        "int32",
			JSONFieldName:      "min_stock",
			ProtobufFieldName:  "min_stock",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "add_time",
			Comment:            `添加类目时间`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CateInfo) TableName() string {
	return "cate_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CateInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CateInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CateInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CateInfo) TableInfo() *TableInfo {
	return cate_infoTableInfo
}
