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


CREATE TABLE `catelm_relation` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `catelm_id` bigint(20) unsigned DEFAULT '0' COMMENT 'catelm_info表的cate_id',
  `cate_id` bigint(20) unsigned DEFAULT '0' COMMENT 'cate_info表的cate_id (三级类目的id)',
  `add_time` int(11) unsigned DEFAULT '0' COMMENT '添加时间',
  `last_time` int(11) unsigned DEFAULT '0' COMMENT '更新时间',
  `is_leaf` tinyint(1) unsigned DEFAULT '0' COMMENT '是否子节点：1是 0:否',
  PRIMARY KEY (`id`),
  KEY `idx_catelmid_leafid` (`catelm_id`,`is_leaf`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='前后类目与后台类目关系表'

JSON Sample
-------------------------------------
{    "id": 79,    "catelm_id": 10,    "cate_id": 59,    "add_time": 20,    "last_time": 5,    "is_leaf": 98}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// CatelmRelation struct is a row record of the catelm_relation table in the yx database
type CatelmRelation struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"`
	//[ 1] catelm_id                                      ubigint              null: true   primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	CatelmID sql.NullInt64 `gorm:"column:catelm_id;type:ubigint;default:0;" json:"catelm_id"` // catelm_info表的cate_id
	//[ 2] cate_id                                        ubigint              null: true   primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	CateID sql.NullInt64 `gorm:"column:cate_id;type:ubigint;default:0;" json:"cate_id"` // cate_info表的cate_id (三级类目的id)
	//[ 3] add_time                                       uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime sql.NullInt64 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 4] last_time                                      uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime sql.NullInt64 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 更新时间
	//[ 5] is_leaf                                        utinyint             null: true   primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	IsLeaf sql.NullInt64 `gorm:"column:is_leaf;type:utinyint;default:0;" json:"is_leaf"` // 是否子节点：1是 0:否

}

var catelm_relationTableInfo = &TableInfo{
	Name: "catelm_relation",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "catelm_id",
			Comment:            `catelm_info表的cate_id`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "CatelmID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "catelm_id",
			ProtobufFieldName:  "catelm_id",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_id",
			Comment:            `cate_info表的cate_id (三级类目的id)`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "add_time",
			Comment:            `添加时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "last_time",
			Comment:            `更新时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "is_leaf",
			Comment:            `是否子节点：1是 0:否`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsLeaf",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "is_leaf",
			ProtobufFieldName:  "is_leaf",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CatelmRelation) TableName() string {
	return "catelm_relation"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CatelmRelation) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CatelmRelation) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CatelmRelation) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CatelmRelation) TableInfo() *TableInfo {
	return catelm_relationTableInfo
}
