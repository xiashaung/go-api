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


CREATE TABLE `cate_industry_relate` (
  `cate_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '一级类目id',
  `cate_industry_id` int(11) NOT NULL COMMENT '大类目id',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`cate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100098 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='行业类目与一级类目关联表'

JSON Sample
-------------------------------------
{    "cate_id": 21,    "cate_industry_id": 78,    "add_time": 19,    "last_time": 1}


Comments
-------------------------------------
[ 2] column is set for unsigned
[ 3] column is set for unsigned



*/

// CateIndustryRelate struct is a row record of the cate_industry_relate table in the yx database
type CateIndustryRelate struct {
	//[ 0] cate_id                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CateID int32 `gorm:"primary_key;AUTO_INCREMENT;column:cate_id;type:int;" json:"cate_id"` // 一级类目id
	//[ 1] cate_industry_id                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CateIndustryID int32 `gorm:"column:cate_industry_id;type:int;" json:"cate_industry_id"` // 大类目id
	//[ 2] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 3] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var cate_industry_relateTableInfo = &TableInfo{
	Name: "cate_industry_relate",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "cate_id",
			Comment:            `一级类目id`,
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
			Name:               "cate_industry_id",
			Comment:            `大类目id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateIndustryID",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_industry_id",
			ProtobufFieldName:  "cate_industry_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CateIndustryRelate) TableName() string {
	return "cate_industry_relate"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CateIndustryRelate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CateIndustryRelate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CateIndustryRelate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CateIndustryRelate) TableInfo() *TableInfo {
	return cate_industry_relateTableInfo
}
