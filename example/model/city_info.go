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


CREATE TABLE `city_info` (
  `city_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '世界城市信息表',
  `city_pid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父级城市id',
  `city_name` char(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '城市名称',
  `city_code` char(12) NOT NULL COMMENT '城市代码(身份证)',
  `region_code` char(12) NOT NULL DEFAULT '' COMMENT '行政区域码',
  `city_level` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '城市层级',
  `city_year` varchar(200) DEFAULT NULL COMMENT '城市变更年份',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`city_id`) USING BTREE,
  UNIQUE KEY `city_code_index` (`city_code`) USING BTREE,
  UNIQUE KEY `region_code_uq` (`region_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18443 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "city_id": 51,    "city_pid": 8,    "city_name": "gMsOGSbvexOODkOpghTfnQTVp",    "city_code": "iMMbOlKeHPvNcsyqeQvLSbBlI",    "region_code": "NeNvISGMNAHasuLHeXqGQpqiF",    "city_level": 35,    "city_year": "yDsrqpZeyrdwMydWnPDHxqEHg",    "add_time": 6,    "last_time": 20}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 5] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned



*/

// CityInfo struct is a row record of the city_info table in the yx database
type CityInfo struct {
	//[ 0] city_id                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CityID int32 `gorm:"primary_key;AUTO_INCREMENT;column:city_id;type:int;" json:"city_id"` // 世界城市信息表
	//[ 1] city_pid                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CityPid uint32 `gorm:"column:city_pid;type:uint;default:0;" json:"city_pid"` // 父级城市id
	//[ 2] city_name                                      char(30)             null: false  primary: false  isArray: false  auto: false  col: char            len: 30      default: []
	CityName string `gorm:"column:city_name;type:char;size:30;" json:"city_name"` // 城市名称
	//[ 3] city_code                                      char(12)             null: false  primary: false  isArray: false  auto: false  col: char            len: 12      default: []
	CityCode string `gorm:"column:city_code;type:char;size:12;" json:"city_code"` // 城市代码(身份证)
	//[ 4] region_code                                    char(12)             null: false  primary: false  isArray: false  auto: false  col: char            len: 12      default: []
	RegionCode string `gorm:"column:region_code;type:char;size:12;" json:"region_code"` // 行政区域码
	//[ 5] city_level                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CityLevel uint32 `gorm:"column:city_level;type:uint;default:0;" json:"city_level"` // 城市层级
	//[ 6] city_year                                      varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	CityYear sql.NullString `gorm:"column:city_year;type:varchar;size:200;" json:"city_year"` // 城市变更年份
	//[ 7] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 8] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 更新时间

}

var city_infoTableInfo = &TableInfo{
	Name: "city_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "city_id",
			Comment:            `世界城市信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CityID",
			GoFieldType:        "int32",
			JSONFieldName:      "city_id",
			ProtobufFieldName:  "city_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "city_pid",
			Comment:            `父级城市id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CityPid",
			GoFieldType:        "uint32",
			JSONFieldName:      "city_pid",
			ProtobufFieldName:  "city_pid",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "city_name",
			Comment:            `城市名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       30,
			GoFieldName:        "CityName",
			GoFieldType:        "string",
			JSONFieldName:      "city_name",
			ProtobufFieldName:  "city_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "city_code",
			Comment:            `城市代码(身份证)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       12,
			GoFieldName:        "CityCode",
			GoFieldType:        "string",
			JSONFieldName:      "city_code",
			ProtobufFieldName:  "city_code",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "region_code",
			Comment:            `行政区域码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       12,
			GoFieldName:        "RegionCode",
			GoFieldType:        "string",
			JSONFieldName:      "region_code",
			ProtobufFieldName:  "region_code",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "city_level",
			Comment:            `城市层级`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CityLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "city_level",
			ProtobufFieldName:  "city_level",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "city_year",
			Comment:            `城市变更年份`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "CityYear",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city_year",
			ProtobufFieldName:  "city_year",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "last_time",
			Comment:            `更新时间`,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CityInfo) TableName() string {
	return "city_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CityInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CityInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CityInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CityInfo) TableInfo() *TableInfo {
	return city_infoTableInfo
}
