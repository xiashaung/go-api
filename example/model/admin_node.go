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


CREATE TABLE `admin_node` (
  `node_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `node_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '节点名字',
  `node_uri` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '节点uri',
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '节点图片',
  `is_system` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0：非系统定义 1：系统定义',
  `node_level` tinyint(1) DEFAULT '1' COMMENT '菜单级别 最多三级',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1:菜单 2：节点',
  `weigh` int(10) NOT NULL DEFAULT '0' COMMENT '权重',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0:禁用 1:开启',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `add_time` int(10) DEFAULT NULL COMMENT '创建时间',
  `last_time` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`node_id`),
  UNIQUE KEY `name` (`node_name`) USING BTREE,
  KEY `pid` (`pid`),
  KEY `weigh` (`weigh`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='节点表'

JSON Sample
-------------------------------------
{    "node_id": 13,    "pid": 15,    "node_name": "quSpoESqbhaiNPvwxtfMZWkNG",    "node_uri": "PBrPcMEMPfFrwhRcZbpkeMmCR",    "icon": "XHfVlGhAUgvcFjoLZYHDxfVch",    "is_system": 81,    "node_level": 26,    "type": 39,    "weigh": 89,    "status": 17,    "remark": "DhqZNqQoYvghpYPJSjGsOVJKV",    "add_time": 4,    "last_time": 32}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 7] column is set for unsigned



*/

// AdminNode struct is a row record of the admin_node table in the yx database
type AdminNode struct {
	//[ 0] node_id                                        uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	NodeID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:node_id;type:uint;" json:"node_id"`
	//[ 1] pid                                            uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Pid uint32 `gorm:"column:pid;type:uint;default:0;" json:"pid"` // 父ID
	//[ 2] node_name                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	NodeName string `gorm:"column:node_name;type:varchar;size:100;" json:"node_name"` // 节点名字
	//[ 3] node_uri                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	NodeURI string `gorm:"column:node_uri;type:varchar;size:50;" json:"node_uri"` // 节点uri
	//[ 4] icon                                           varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Icon string `gorm:"column:icon;type:varchar;size:50;" json:"icon"` // 节点图片
	//[ 5] is_system                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsSystem int32 `gorm:"column:is_system;type:tinyint;default:0;" json:"is_system"` // 0：非系统定义 1：系统定义
	//[ 6] node_level                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	NodeLevel sql.NullInt64 `gorm:"column:node_level;type:tinyint;default:1;" json:"node_level"` // 菜单级别 最多三级
	//[ 7] type                                           utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	Type uint32 `gorm:"column:type;type:utinyint;default:1;" json:"type"` // 1:菜单 2：节点
	//[ 8] weigh                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Weigh int32 `gorm:"column:weigh;type:int;default:0;" json:"weigh"` // 权重
	//[ 9] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Status int32 `gorm:"column:status;type:tinyint;default:1;" json:"status"` // 状态 0:禁用 1:开启
	//[10] remark                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Remark string `gorm:"column:remark;type:varchar;size:255;" json:"remark"` // 备注
	//[11] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"` // 创建时间
	//[12] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"` // 更新时间

}

var admin_nodeTableInfo = &TableInfo{
	Name: "admin_node",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "node_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "NodeID",
			GoFieldType:        "uint32",
			JSONFieldName:      "node_id",
			ProtobufFieldName:  "node_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "pid",
			Comment:            `父ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Pid",
			GoFieldType:        "uint32",
			JSONFieldName:      "pid",
			ProtobufFieldName:  "pid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "node_name",
			Comment:            `节点名字`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "NodeName",
			GoFieldType:        "string",
			JSONFieldName:      "node_name",
			ProtobufFieldName:  "node_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "node_uri",
			Comment:            `节点uri`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "NodeURI",
			GoFieldType:        "string",
			JSONFieldName:      "node_uri",
			ProtobufFieldName:  "node_uri",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "icon",
			Comment:            `节点图片`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Icon",
			GoFieldType:        "string",
			JSONFieldName:      "icon",
			ProtobufFieldName:  "icon",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_system",
			Comment:            `0：非系统定义 1：系统定义`,
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
			Name:               "node_level",
			Comment:            `菜单级别 最多三级`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "NodeLevel",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "node_level",
			ProtobufFieldName:  "node_level",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "type",
			Comment:            `1:菜单 2：节点`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "uint32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "weigh",
			Comment:            `权重`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Weigh",
			GoFieldType:        "int32",
			JSONFieldName:      "weigh",
			ProtobufFieldName:  "weigh",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status",
			Comment:            `状态 0:禁用 1:开启`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "remark",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "add_time",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "last_time",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdminNode) TableName() string {
	return "admin_node"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdminNode) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdminNode) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdminNode) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdminNode) TableInfo() *TableInfo {
	return admin_nodeTableInfo
}
