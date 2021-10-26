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


CREATE TABLE `talent_cate_industry` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `talent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `cate_industry_id` int(11) NOT NULL COMMENT '意向大类id',
  `cate_industry_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '意向大类目名称',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=171 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='达人意向类目表（大类）'

JSON Sample
-------------------------------------
{    "id": 6,    "talent_id": 34,    "cate_industry_id": 30,    "cate_industry_name": "aCkKXPMGCQxiegbSEbkXVwISu",    "add_time": 89,    "last_time": 34}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// TalentCateIndustry struct is a row record of the talent_cate_industry table in the yx database
type TalentCateIndustry struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:int;" json:"id"` // 自增id
	//[ 1] talent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TalentID uint32 `gorm:"column:talent_id;type:uint;default:0;" json:"talent_id"` // 达人id
	//[ 2] cate_industry_id                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CateIndustryID int32 `gorm:"column:cate_industry_id;type:int;" json:"cate_industry_id"` // 意向大类id
	//[ 3] cate_industry_name                             varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CateIndustryName string `gorm:"column:cate_industry_name;type:varchar;size:50;" json:"cate_industry_name"` // 意向大类目名称
	//[ 4] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 5] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var talent_cate_industryTableInfo = &TableInfo{
	Name: "talent_cate_industry",
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
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_industry_id",
			Comment:            `意向大类id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_industry_name",
			Comment:            `意向大类目名称`,
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
func (t *TalentCateIndustry) TableName() string {
	return "talent_cate_industry"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TalentCateIndustry) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TalentCateIndustry) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TalentCateIndustry) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TalentCateIndustry) TableInfo() *TableInfo {
	return talent_cate_industryTableInfo
}
