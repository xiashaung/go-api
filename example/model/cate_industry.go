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


CREATE TABLE `cate_industry` (
  `cate_industry_id` int(11) NOT NULL AUTO_INCREMENT,
  `cate_industry_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '大类目名称',
  `sort` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加类目时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`cate_industry_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='行业类目表'

JSON Sample
-------------------------------------
{    "cate_industry_id": 21,    "cate_industry_name": "KDvBJiflOawFKeKMKEKBAfXOP",    "sort": 62,    "add_time": 81,    "last_time": 59}


Comments
-------------------------------------
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// CateIndustry struct is a row record of the cate_industry table in the yx database
type CateIndustry struct {
	//[ 0] cate_industry_id                               int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CateIndustryID int32 `gorm:"primary_key;AUTO_INCREMENT;column:cate_industry_id;type:int;" json:"cate_industry_id"`
	//[ 1] cate_industry_name                             varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CateIndustryName string `gorm:"column:cate_industry_name;type:varchar;size:50;" json:"cate_industry_name"` // 大类目名称
	//[ 2] sort                                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Sort uint32 `gorm:"column:sort;type:uint;default:0;" json:"sort"` // 排序
	//[ 3] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加类目时间
	//[ 4] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var cate_industryTableInfo = &TableInfo{
	Name: "cate_industry",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "cate_industry_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateIndustryID",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_industry_id",
			ProtobufFieldName:  "cate_industry_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "cate_industry_name",
			Comment:            `大类目名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "CateIndustryName",
			GoFieldType:        "string",
			JSONFieldName:      "cate_industry_name",
			ProtobufFieldName:  "cate_industry_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "sort",
			Comment:            `排序`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Sort",
			GoFieldType:        "uint32",
			JSONFieldName:      "sort",
			ProtobufFieldName:  "sort",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CateIndustry) TableName() string {
	return "cate_industry"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CateIndustry) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CateIndustry) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CateIndustry) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CateIndustry) TableInfo() *TableInfo {
	return cate_industryTableInfo
}
