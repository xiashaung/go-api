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


CREATE TABLE `tag_info` (
  `tag_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名称',
  `tag_pid` int(11) NOT NULL DEFAULT '0' COMMENT '父级标签',
  `tag_type` int(11) NOT NULL DEFAULT '0' COMMENT '标签类型',
  `mark_type` int(11) NOT NULL DEFAULT '2' COMMENT '打标签类型 1 自动 2 手动',
  `tag_weight` int(11) NOT NULL DEFAULT '0' COMMENT '权重 最大10000',
  `input_type` int(11) DEFAULT '0' COMMENT '输入方式',
  `product_dimension` int(11) DEFAULT NULL COMMENT '标签属性',
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

JSON Sample
-------------------------------------
{    "tag_id": 37,    "tag_name": "RFkakiyttDMvhKgZCEcKRWuZg",    "tag_pid": 40,    "tag_type": 61,    "mark_type": 83,    "tag_weight": 70,    "input_type": 20,    "product_dimension": 28,    "add_time": 65,    "last_time": 93}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// TagInfo struct is a row record of the tag_info table in the yx database
type TagInfo struct {
	//[ 0] tag_id                                         uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	TagID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:tag_id;type:uint;" json:"tag_id"`
	//[ 1] tag_name                                       varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	TagName string `gorm:"column:tag_name;type:varchar;size:60;" json:"tag_name"` // 标签名称
	//[ 2] tag_pid                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TagPid int32 `gorm:"column:tag_pid;type:int;default:0;" json:"tag_pid"` // 父级标签
	//[ 3] tag_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TagType int32 `gorm:"column:tag_type;type:int;default:0;" json:"tag_type"` // 标签类型
	//[ 4] mark_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [2]
	MarkType int32 `gorm:"column:mark_type;type:int;default:2;" json:"mark_type"` // 打标签类型 1 自动 2 手动
	//[ 5] tag_weight                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TagWeight int32 `gorm:"column:tag_weight;type:int;default:0;" json:"tag_weight"` // 权重 最大10000
	//[ 6] input_type                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	InputType sql.NullInt64 `gorm:"column:input_type;type:int;default:0;" json:"input_type"` // 输入方式
	//[ 7] product_dimension                              int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ProductDimension sql.NullInt64 `gorm:"column:product_dimension;type:int;" json:"product_dimension"` // 标签属性
	//[ 8] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 9] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var tag_infoTableInfo = &TableInfo{
	Name: "tag_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "tag_id",
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
			GoFieldName:        "TagID",
			GoFieldType:        "uint32",
			JSONFieldName:      "tag_id",
			ProtobufFieldName:  "tag_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "tag_name",
			Comment:            `标签名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "TagName",
			GoFieldType:        "string",
			JSONFieldName:      "tag_name",
			ProtobufFieldName:  "tag_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "tag_pid",
			Comment:            `父级标签`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TagPid",
			GoFieldType:        "int32",
			JSONFieldName:      "tag_pid",
			ProtobufFieldName:  "tag_pid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "tag_type",
			Comment:            `标签类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TagType",
			GoFieldType:        "int32",
			JSONFieldName:      "tag_type",
			ProtobufFieldName:  "tag_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "mark_type",
			Comment:            `打标签类型 1 自动 2 手动`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MarkType",
			GoFieldType:        "int32",
			JSONFieldName:      "mark_type",
			ProtobufFieldName:  "mark_type",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "tag_weight",
			Comment:            `权重 最大10000`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TagWeight",
			GoFieldType:        "int32",
			JSONFieldName:      "tag_weight",
			ProtobufFieldName:  "tag_weight",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "input_type",
			Comment:            `输入方式`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "InputType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "input_type",
			ProtobufFieldName:  "input_type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "product_dimension",
			Comment:            `标签属性`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductDimension",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "product_dimension",
			ProtobufFieldName:  "product_dimension",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "add_time",
			Comment:            ``,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "last_time",
			Comment:            ``,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TagInfo) TableName() string {
	return "tag_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TagInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TagInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TagInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TagInfo) TableInfo() *TableInfo {
	return tag_infoTableInfo
}
