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


CREATE TABLE `mcn_info` (
  `mcn_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'mcn机构信息表',
  `company_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '公司id',
  `mcn_name` varchar(30) NOT NULL COMMENT 'mcn名称',
  `add_time` int(10) unsigned NOT NULL COMMENT 'mcn入驻时间',
  `last_time` int(10) unsigned NOT NULL,
  PRIMARY KEY (`mcn_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "mcn_id": 15,    "company_id": 92,    "mcn_name": "lTVgcUSkyyvxpdZpdPBLwTUjA",    "add_time": 56,    "last_time": 65}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// McnInfo struct is a row record of the mcn_info table in the yx database
type McnInfo struct {
	//[ 0] mcn_id                                         int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	McnID int32 `gorm:"primary_key;AUTO_INCREMENT;column:mcn_id;type:int;" json:"mcn_id"` // mcn机构信息表
	//[ 1] company_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CompanyID uint32 `gorm:"column:company_id;type:uint;default:0;" json:"company_id"` // 公司id
	//[ 2] mcn_name                                       varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	McnName string `gorm:"column:mcn_name;type:varchar;size:30;" json:"mcn_name"` // mcn名称
	//[ 3] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	AddTime uint32 `gorm:"column:add_time;type:uint;" json:"add_time"` // mcn入驻时间
	//[ 4] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	LastTime uint32 `gorm:"column:last_time;type:uint;" json:"last_time"`
}

var mcn_infoTableInfo = &TableInfo{
	Name: "mcn_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "mcn_id",
			Comment:            `mcn机构信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "McnID",
			GoFieldType:        "int32",
			JSONFieldName:      "mcn_id",
			ProtobufFieldName:  "mcn_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "company_id",
			Comment:            `公司id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CompanyID",
			GoFieldType:        "uint32",
			JSONFieldName:      "company_id",
			ProtobufFieldName:  "company_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "mcn_name",
			Comment:            `mcn名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "McnName",
			GoFieldType:        "string",
			JSONFieldName:      "mcn_name",
			ProtobufFieldName:  "mcn_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "add_time",
			Comment:            `mcn入驻时间`,
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
func (m *McnInfo) TableName() string {
	return "mcn_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *McnInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *McnInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *McnInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *McnInfo) TableInfo() *TableInfo {
	return mcn_infoTableInfo
}
