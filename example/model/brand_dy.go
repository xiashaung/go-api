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


CREATE TABLE `brand_dy` (
  `id` int(11) NOT NULL COMMENT '抖音品牌信息表',
  `yx_brand_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '央选平台品牌id',
  `brand_chinese_name` varchar(30) NOT NULL COMMENT '品牌中文名',
  `brand_english_name` varchar(30) NOT NULL COMMENT '品牌英文名称',
  `brand_reg_num` int(11) DEFAULT NULL COMMENT '品牌注册号',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 54,    "yx_brand_id": 31,    "brand_chinese_name": "TOmcfkMFAoaFmCmWeYfXsUwED",    "brand_english_name": "LahuOoAZkBaXuywQHwvdpontH",    "brand_reg_num": 77,    "add_time": 38,    "last_time": 24}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// BrandDy struct is a row record of the brand_dy table in the yx database
type BrandDy struct {
	//[ 0] id                                             int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ID int32 `gorm:"primary_key;column:id;type:int;" json:"id"` // 抖音品牌信息表
	//[ 1] yx_brand_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	YxBrandID uint32 `gorm:"column:yx_brand_id;type:uint;default:0;" json:"yx_brand_id"` // 央选平台品牌id
	//[ 2] brand_chinese_name                             varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	BrandChineseName string `gorm:"column:brand_chinese_name;type:varchar;size:30;" json:"brand_chinese_name"` // 品牌中文名
	//[ 3] brand_english_name                             varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	BrandEnglishName string `gorm:"column:brand_english_name;type:varchar;size:30;" json:"brand_english_name"` // 品牌英文名称
	//[ 4] brand_reg_num                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BrandRegNum sql.NullInt64 `gorm:"column:brand_reg_num;type:int;" json:"brand_reg_num"` // 品牌注册号
	//[ 5] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 6] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var brand_dyTableInfo = &TableInfo{
	Name: "brand_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `抖音品牌信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
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
			Name:               "yx_brand_id",
			Comment:            `央选平台品牌id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "YxBrandID",
			GoFieldType:        "uint32",
			JSONFieldName:      "yx_brand_id",
			ProtobufFieldName:  "yx_brand_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "brand_chinese_name",
			Comment:            `品牌中文名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "BrandChineseName",
			GoFieldType:        "string",
			JSONFieldName:      "brand_chinese_name",
			ProtobufFieldName:  "brand_chinese_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "brand_english_name",
			Comment:            `品牌英文名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "BrandEnglishName",
			GoFieldType:        "string",
			JSONFieldName:      "brand_english_name",
			ProtobufFieldName:  "brand_english_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "brand_reg_num",
			Comment:            `品牌注册号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BrandRegNum",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "brand_reg_num",
			ProtobufFieldName:  "brand_reg_num",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BrandDy) TableName() string {
	return "brand_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BrandDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BrandDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BrandDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BrandDy) TableInfo() *TableInfo {
	return brand_dyTableInfo
}
