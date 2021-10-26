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


CREATE TABLE `agreement_info` (
  `agreement_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `source_type` tinyint(3) DEFAULT '1' COMMENT '履约类型：1样品订单',
  `source_id` int(10) NOT NULL DEFAULT '0' COMMENT 'type值：1、样品订单sample_id',
  `agreement_type` tinyint(3) DEFAULT '0' COMMENT '履约方式：0未选择履约方式，1短视频带货，2直播带货',
  `link_url` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '不履约图片链接或履约视频、图片链接',
  `unagree_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '不履约原因',
  `start_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '直播开始时间',
  `end_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '直播结束时间',
  `status` tinyint(3) DEFAULT '1' COMMENT '状态值，0不需要履约，1未履约，2履约待验收，3不履约待审核，4已履约，5不履约',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`agreement_id`),
  KEY `idx_source_id` (`source_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='履约表'

JSON Sample
-------------------------------------
{    "agreement_id": 17,    "source_type": 79,    "source_id": 42,    "agreement_type": 57,    "link_url": "nVaRLkOQCqgMZrKsTpItUvhnk",    "unagree_reason": "AsTibacYKkHsUHKisKoqBfHas",    "start_time": 39,    "end_time": 16,    "status": 95,    "add_time": 52,    "last_time": 19}


Comments
-------------------------------------
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// AgreementInfo struct is a row record of the agreement_info table in the yx database
type AgreementInfo struct {
	//[ 0] agreement_id                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	AgreementID int32 `gorm:"primary_key;AUTO_INCREMENT;column:agreement_id;type:int;" json:"agreement_id"` // 自增id
	//[ 1] source_type                                    tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	SourceType sql.NullInt64 `gorm:"column:source_type;type:tinyint;default:1;" json:"source_type"` // 履约类型：1样品订单
	//[ 2] source_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SourceID int32 `gorm:"column:source_id;type:int;default:0;" json:"source_id"` // type值：1、样品订单sample_id
	//[ 3] agreement_type                                 tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	AgreementType sql.NullInt64 `gorm:"column:agreement_type;type:tinyint;default:0;" json:"agreement_type"` // 履约方式：0未选择履约方式，1短视频带货，2直播带货
	//[ 4] link_url                                       text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	LinkURL string `gorm:"column:link_url;type:text;size:65535;" json:"link_url"` // 不履约图片链接或履约视频、图片链接
	//[ 5] unagree_reason                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	UnagreeReason string `gorm:"column:unagree_reason;type:varchar;size:255;" json:"unagree_reason"` // 不履约原因
	//[ 6] start_time                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	StartTime uint32 `gorm:"column:start_time;type:uint;default:0;" json:"start_time"` // 直播开始时间
	//[ 7] end_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EndTime uint32 `gorm:"column:end_time;type:uint;default:0;" json:"end_time"` // 直播结束时间
	//[ 8] status                                         tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Status sql.NullInt64 `gorm:"column:status;type:tinyint;default:1;" json:"status"` // 状态值，0不需要履约，1未履约，2履约待验收，3不履约待审核，4已履约，5不履约
	//[ 9] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[10] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var agreement_infoTableInfo = &TableInfo{
	Name: "agreement_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "agreement_id",
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
			GoFieldName:        "AgreementID",
			GoFieldType:        "int32",
			JSONFieldName:      "agreement_id",
			ProtobufFieldName:  "agreement_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "source_type",
			Comment:            `履约类型：1样品订单`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "SourceType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "source_type",
			ProtobufFieldName:  "source_type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "source_id",
			Comment:            `type值：1、样品订单sample_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SourceID",
			GoFieldType:        "int32",
			JSONFieldName:      "source_id",
			ProtobufFieldName:  "source_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "agreement_type",
			Comment:            `履约方式：0未选择履约方式，1短视频带货，2直播带货`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "AgreementType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "agreement_type",
			ProtobufFieldName:  "agreement_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "link_url",
			Comment:            `不履约图片链接或履约视频、图片链接`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "LinkURL",
			GoFieldType:        "string",
			JSONFieldName:      "link_url",
			ProtobufFieldName:  "link_url",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "unagree_reason",
			Comment:            `不履约原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "UnagreeReason",
			GoFieldType:        "string",
			JSONFieldName:      "unagree_reason",
			ProtobufFieldName:  "unagree_reason",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "start_time",
			Comment:            `直播开始时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "end_time",
			Comment:            `直播结束时间`,
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
			Name:               "status",
			Comment:            `状态值，0不需要履约，1未履约，2履约待验收，3不履约待审核，4已履约，5不履约`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AgreementInfo) TableName() string {
	return "agreement_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AgreementInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AgreementInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AgreementInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AgreementInfo) TableInfo() *TableInfo {
	return agreement_infoTableInfo
}
