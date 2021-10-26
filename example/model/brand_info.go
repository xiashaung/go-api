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


CREATE TABLE `brand_info` (
  `brand_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '品牌信息表',
  `brand_type` int(255) unsigned NOT NULL DEFAULT '1' COMMENT '品牌类型 1：抖音 2：快手',
  `ref_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '对应的央选品牌id',
  `third_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '第三方平台品牌id',
  `brand_cname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '品牌中文名称',
  `brand_ename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '品牌英文名称',
  `brand_logo` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '品牌logo',
  `brand_sn` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '品牌注册号',
  `brand_desc` text COMMENT '品牌描述',
  `is_auto` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:系统识别加入 1:统一品牌接口导入',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加品牌时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`brand_id`)
) ENGINE=InnoDB AUTO_INCREMENT=181 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='品牌信息表'

JSON Sample
-------------------------------------
{    "brand_id": 81,    "brand_type": 74,    "ref_id": 60,    "third_id": 83,    "brand_cname": "BsjBuAUJCBUfseYJcLouNcBvO",    "brand_ename": "QHuJIGcRXwWahPsXMNpUaeMjP",    "brand_logo": "KvoIFpHPPhdeSLFWmvSkeTGkY",    "brand_sn": "aWwJddwAeWaRJhDZoDjwUSHuN",    "brand_desc": "uqoamRCSdUdkMWtsPcyupsHQF",    "is_auto": 88,    "add_time": 78,    "last_time": 5}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned



*/

// BrandInfo struct is a row record of the brand_info table in the yx database
type BrandInfo struct {
	//[ 0] brand_id                                       int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	BrandID int32 `gorm:"primary_key;AUTO_INCREMENT;column:brand_id;type:int;" json:"brand_id"` // 品牌信息表
	//[ 1] brand_type                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	BrandType uint32 `gorm:"column:brand_type;type:uint;default:1;" json:"brand_type"` // 品牌类型 1：抖音 2：快手
	//[ 2] ref_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RefID uint32 `gorm:"column:ref_id;type:uint;default:0;" json:"ref_id"` // 对应的央选品牌id
	//[ 3] third_id                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ThirdID uint32 `gorm:"column:third_id;type:uint;default:0;" json:"third_id"` // 第三方平台品牌id
	//[ 4] brand_cname                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	BrandCname string `gorm:"column:brand_cname;type:varchar;size:30;" json:"brand_cname"` // 品牌中文名称
	//[ 5] brand_ename                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BrandEname string `gorm:"column:brand_ename;type:varchar;size:255;" json:"brand_ename"` // 品牌英文名称
	//[ 6] brand_logo                                     varchar(300)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 300     default: []
	BrandLogo string `gorm:"column:brand_logo;type:varchar;size:300;" json:"brand_logo"` // 品牌logo
	//[ 7] brand_sn                                       varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	BrandSn string `gorm:"column:brand_sn;type:varchar;size:60;" json:"brand_sn"` // 品牌注册号
	//[ 8] brand_desc                                     text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	BrandDesc sql.NullString `gorm:"column:brand_desc;type:text;size:65535;" json:"brand_desc"` // 品牌描述
	//[ 9] is_auto                                        tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsAuto int32 `gorm:"column:is_auto;type:tinyint;default:0;" json:"is_auto"` // 0:系统识别加入 1:统一品牌接口导入
	//[10] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加品牌时间
	//[11] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var brand_infoTableInfo = &TableInfo{
	Name: "brand_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "brand_id",
			Comment:            `品牌信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BrandID",
			GoFieldType:        "int32",
			JSONFieldName:      "brand_id",
			ProtobufFieldName:  "brand_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "brand_type",
			Comment:            `品牌类型 1：抖音 2：快手`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "BrandType",
			GoFieldType:        "uint32",
			JSONFieldName:      "brand_type",
			ProtobufFieldName:  "brand_type",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ref_id",
			Comment:            `对应的央选品牌id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "RefID",
			GoFieldType:        "uint32",
			JSONFieldName:      "ref_id",
			ProtobufFieldName:  "ref_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "third_id",
			Comment:            `第三方平台品牌id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdID",
			GoFieldType:        "uint32",
			JSONFieldName:      "third_id",
			ProtobufFieldName:  "third_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "brand_cname",
			Comment:            `品牌中文名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "BrandCname",
			GoFieldType:        "string",
			JSONFieldName:      "brand_cname",
			ProtobufFieldName:  "brand_cname",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "brand_ename",
			Comment:            `品牌英文名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "BrandEname",
			GoFieldType:        "string",
			JSONFieldName:      "brand_ename",
			ProtobufFieldName:  "brand_ename",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "brand_logo",
			Comment:            `品牌logo`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       300,
			GoFieldName:        "BrandLogo",
			GoFieldType:        "string",
			JSONFieldName:      "brand_logo",
			ProtobufFieldName:  "brand_logo",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "brand_sn",
			Comment:            `品牌注册号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "BrandSn",
			GoFieldType:        "string",
			JSONFieldName:      "brand_sn",
			ProtobufFieldName:  "brand_sn",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "brand_desc",
			Comment:            `品牌描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "BrandDesc",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "brand_desc",
			ProtobufFieldName:  "brand_desc",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "is_auto",
			Comment:            `0:系统识别加入 1:统一品牌接口导入`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsAuto",
			GoFieldType:        "int32",
			JSONFieldName:      "is_auto",
			ProtobufFieldName:  "is_auto",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "add_time",
			Comment:            `添加品牌时间`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BrandInfo) TableName() string {
	return "brand_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BrandInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BrandInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BrandInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BrandInfo) TableInfo() *TableInfo {
	return brand_infoTableInfo
}
