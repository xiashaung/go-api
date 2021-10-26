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


CREATE TABLE `white_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `seller_id` int(11) DEFAULT NULL COMMENT '商家id',
  `type` int(11) DEFAULT NULL COMMENT '业务类型',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '原因',
  `start_time` int(11) DEFAULT NULL COMMENT '开始时间',
  `end_time` int(11) DEFAULT NULL COMMENT '结束时间',
  `platform_id` int(11) DEFAULT NULL COMMENT '平台类型1 抖音 2 快手',
  `status` int(11) DEFAULT '1' COMMENT '1正常 2  中止',
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='白名单'

JSON Sample
-------------------------------------
{    "id": 84,    "seller_id": 99,    "type": 20,    "remark": "QJtDapFTfGlEeHKZUQldVrJdX",    "start_time": 43,    "end_time": 52,    "platform_id": 57,    "status": 87,    "add_time": 2,    "last_time": 15}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// WhiteList struct is a row record of the white_list table in the yx database
type WhiteList struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] seller_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SellerID sql.NullInt64 `gorm:"column:seller_id;type:int;" json:"seller_id"` // 商家id
	//[ 2] type                                           int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	Type sql.NullInt64 `gorm:"column:type;type:int;" json:"type"` // 业务类型
	//[ 3] remark                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Remark sql.NullString `gorm:"column:remark;type:varchar;size:255;" json:"remark"` // 原因
	//[ 4] start_time                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	StartTime sql.NullInt64 `gorm:"column:start_time;type:int;" json:"start_time"` // 开始时间
	//[ 5] end_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	EndTime sql.NullInt64 `gorm:"column:end_time;type:int;" json:"end_time"` // 结束时间
	//[ 6] platform_id                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PlatformID sql.NullInt64 `gorm:"column:platform_id;type:int;" json:"platform_id"` // 平台类型1 抖音 2 快手
	//[ 7] status                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Status sql.NullInt64 `gorm:"column:status;type:int;default:1;" json:"status"` // 1正常 2  中止
	//[ 8] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 9] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var white_listTableInfo = &TableInfo{
	Name: "white_list",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "seller_id",
			Comment:            `商家id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SellerID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "seller_id",
			ProtobufFieldName:  "seller_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "type",
			Comment:            `业务类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "remark",
			Comment:            `原因`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "start_time",
			Comment:            `开始时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "end_time",
			Comment:            `结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "platform_id",
			Comment:            `平台类型1 抖音 2 快手`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "status",
			Comment:            `1正常 2  中止`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "add_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "last_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WhiteList) TableName() string {
	return "white_list"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WhiteList) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WhiteList) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WhiteList) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WhiteList) TableInfo() *TableInfo {
	return white_listTableInfo
}
