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


CREATE TABLE `talent_right` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `talent_id` int(11) NOT NULL COMMENT '达人id',
  `right_type` int(11) NOT NULL DEFAULT '0' COMMENT '权益类型',
  `right_info` varchar(255) NOT NULL DEFAULT '' COMMENT '权益内容',
  `source_type` int(11) NOT NULL DEFAULT '0' COMMENT '来源类型 1 开业活动',
  `source_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '来源描述',
  `start_time` int(11) NOT NULL COMMENT '生效开始时间',
  `end_time` int(11) unsigned NOT NULL COMMENT '生效结束时间',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='达人权益表'

JSON Sample
-------------------------------------
{    "id": 13,    "talent_id": 49,    "right_type": 25,    "right_info": "VaRMhGUqqOCKmLDKlWxopLCes",    "source_type": 42,    "source_desc": "EMkxLWlUglMUJqwZIuKYQUaCr",    "start_time": 67,    "end_time": 60,    "add_time": 46,    "last_time": 92}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 7] column is set for unsigned



*/

// TalentRight struct is a row record of the talent_right table in the yx database
type TalentRight struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TalentID int32 `gorm:"column:talent_id;type:int;" json:"talent_id"` // 达人id
	//[ 2] right_type                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RightType int32 `gorm:"column:right_type;type:int;default:0;" json:"right_type"` // 权益类型
	//[ 3] right_info                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	RightInfo string `gorm:"column:right_info;type:varchar;size:255;" json:"right_info"` // 权益内容
	//[ 4] source_type                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SourceType int32 `gorm:"column:source_type;type:int;default:0;" json:"source_type"` // 来源类型 1 开业活动
	//[ 5] source_desc                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SourceDesc string `gorm:"column:source_desc;type:varchar;size:255;" json:"source_desc"` // 来源描述
	//[ 6] start_time                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	StartTime int32 `gorm:"column:start_time;type:int;" json:"start_time"` // 生效开始时间
	//[ 7] end_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	EndTime uint32 `gorm:"column:end_time;type:uint;" json:"end_time"` // 生效结束时间
	//[ 8] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 9] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var talent_rightTableInfo = &TableInfo{
	Name: "talent_right",
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
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "right_type",
			Comment:            `权益类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RightType",
			GoFieldType:        "int32",
			JSONFieldName:      "right_type",
			ProtobufFieldName:  "right_type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "right_info",
			Comment:            `权益内容`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "RightInfo",
			GoFieldType:        "string",
			JSONFieldName:      "right_info",
			ProtobufFieldName:  "right_info",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "source_type",
			Comment:            `来源类型 1 开业活动`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SourceType",
			GoFieldType:        "int32",
			JSONFieldName:      "source_type",
			ProtobufFieldName:  "source_type",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "source_desc",
			Comment:            `来源描述`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SourceDesc",
			GoFieldType:        "string",
			JSONFieldName:      "source_desc",
			ProtobufFieldName:  "source_desc",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "start_time",
			Comment:            `生效开始时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "int32",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "end_time",
			Comment:            `生效结束时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TalentRight) TableName() string {
	return "talent_right"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TalentRight) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TalentRight) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TalentRight) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TalentRight) TableInfo() *TableInfo {
	return talent_rightTableInfo
}
