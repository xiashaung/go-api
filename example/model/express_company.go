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


CREATE TABLE `express_company` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `platform_id` smallint(6) NOT NULL DEFAULT '2' COMMENT '1-抖音,2-快手',
  `express_code` bigint(20) NOT NULL DEFAULT '0' COMMENT '快速公司 code',
  `express_name` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递公司名称',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='快递公司'

JSON Sample
-------------------------------------
{    "id": 27,    "platform_id": 34,    "express_code": 55,    "express_name": "NlIGYgcByXlWKfMqxpVmMIHnA",    "add_time": 53,    "last_time": 60}



*/

// ExpressCompany struct is a row record of the express_company table in the yx database
type ExpressCompany struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 主键
	//[ 1] platform_id                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [2]
	PlatformID int32 `gorm:"column:platform_id;type:smallint;default:2;" json:"platform_id"` // 1-抖音,2-快手
	//[ 2] express_code                                   bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ExpressCode int64 `gorm:"column:express_code;type:bigint;default:0;" json:"express_code"` // 快速公司 code
	//[ 3] express_name                                   varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	ExpressName string `gorm:"column:express_name;type:varchar;size:125;" json:"express_name"` // 快递公司名称
	//[ 4] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 5] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var express_companyTableInfo = &TableInfo{
	Name: "express_company",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "platform_id",
			Comment:            `1-抖音,2-快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "express_code",
			Comment:            `快速公司 code`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressCode",
			GoFieldType:        "int64",
			JSONFieldName:      "express_code",
			ProtobufFieldName:  "express_code",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "express_name",
			Comment:            `快递公司名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "ExpressName",
			GoFieldType:        "string",
			JSONFieldName:      "express_name",
			ProtobufFieldName:  "express_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "int32",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "last_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "int32",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (e *ExpressCompany) TableName() string {
	return "express_company"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *ExpressCompany) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *ExpressCompany) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *ExpressCompany) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (e *ExpressCompany) TableInfo() *TableInfo {
	return express_companyTableInfo
}
