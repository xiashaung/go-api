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


CREATE TABLE `logistics_company` (
  `company_id` int(11) NOT NULL COMMENT '物流公司信息表',
  `company_name` varchar(20) NOT NULL COMMENT '物流公司名称',
  `company_type` int(255) DEFAULT NULL COMMENT '物流公司类别 1：央选 2：抖音',
  `ref_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '映射央选物流公司id',
  `third_id` int(11) DEFAULT NULL COMMENT '第三方id',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加物流公司时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`company_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "company_id": 44,    "company_name": "TOcmMjQYqbZVqkeYgTNjWGQvc",    "company_type": 69,    "ref_id": 27,    "third_id": 39,    "add_time": 39,    "last_time": 94}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// LogisticsCompany struct is a row record of the logistics_company table in the yx database
type LogisticsCompany struct {
	//[ 0] company_id                                     int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CompanyID int32 `gorm:"primary_key;column:company_id;type:int;" json:"company_id"` // 物流公司信息表
	//[ 1] company_name                                   varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	CompanyName string `gorm:"column:company_name;type:varchar;size:20;" json:"company_name"` // 物流公司名称
	//[ 2] company_type                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CompanyType sql.NullInt64 `gorm:"column:company_type;type:int;" json:"company_type"` // 物流公司类别 1：央选 2：抖音
	//[ 3] ref_id                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RefID uint32 `gorm:"column:ref_id;type:uint;default:0;" json:"ref_id"` // 映射央选物流公司id
	//[ 4] third_id                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ThirdID sql.NullInt64 `gorm:"column:third_id;type:int;" json:"third_id"` // 第三方id
	//[ 5] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加物流公司时间
	//[ 6] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var logistics_companyTableInfo = &TableInfo{
	Name: "logistics_company",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "company_id",
			Comment:            `物流公司信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CompanyID",
			GoFieldType:        "int32",
			JSONFieldName:      "company_id",
			ProtobufFieldName:  "company_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "company_name",
			Comment:            `物流公司名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "CompanyName",
			GoFieldType:        "string",
			JSONFieldName:      "company_name",
			ProtobufFieldName:  "company_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "company_type",
			Comment:            `物流公司类别 1：央选 2：抖音`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CompanyType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "company_type",
			ProtobufFieldName:  "company_type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ref_id",
			Comment:            `映射央选物流公司id`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "third_id",
			Comment:            `第三方id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ThirdID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "third_id",
			ProtobufFieldName:  "third_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "add_time",
			Comment:            `添加物流公司时间`,
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
func (l *LogisticsCompany) TableName() string {
	return "logistics_company"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LogisticsCompany) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LogisticsCompany) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LogisticsCompany) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LogisticsCompany) TableInfo() *TableInfo {
	return logistics_companyTableInfo
}
