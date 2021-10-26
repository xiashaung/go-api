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


CREATE TABLE `tender_info` (
  `tender_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '招商计划信息表',
  `talent_id` int(11) NOT NULL DEFAULT '0' COMMENT '达人id',
  `platform_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属平台 1：抖音 2：快手',
  `anchor_id` int(11) NOT NULL DEFAULT '0' COMMENT '主播id',
  `tender_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '招商计划名称',
  `tender_type` int(11) NOT NULL DEFAULT '0' COMMENT '1 专场   2 拼场',
  `tender_start` int(10) NOT NULL DEFAULT '0' COMMENT '招商计划的开始时间，即报名开始时间',
  `tender_end` int(10) NOT NULL DEFAULT '0' COMMENT '招商计划的结束时间，即报名结束时间',
  `live_start` int(10) NOT NULL DEFAULT '0' COMMENT '直播开始时间',
  `live_end` int(10) NOT NULL DEFAULT '0' COMMENT '直播结束时间',
  `tender_require` text COMMENT '招商计划的其他要求',
  `tender_desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '招商计划备注',
  `tender_status` int(11) NOT NULL DEFAULT '0' COMMENT '1 草稿 2 待报名 3 报名中 4 报名截止',
  `is_deleted` int(1) NOT NULL DEFAULT '0' COMMENT '是否删除，0 否，正常   1 是，已删除 ',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '创建招商计划的时间',
  `last_time` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`tender_id`),
  UNIQUE KEY `idx_plan_tender_id` (`tender_id`),
  KEY `idx_tender_time` (`tender_start`,`tender_end`)
) ENGINE=InnoDB AUTO_INCREMENT=156 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='招商计划信息表'

JSON Sample
-------------------------------------
{    "tender_id": 0,    "talent_id": 68,    "platform_id": 2,    "anchor_id": 21,    "tender_name": "tYObQfmFNignPAirjVBPjIUtH",    "tender_type": 36,    "tender_start": 19,    "tender_end": 20,    "live_start": 66,    "live_end": 24,    "tender_require": "aNGVFOyHtLMULsZsqALeHkUsA",    "tender_desc": "KXTuEuELEbGLrLiWcaNCENMnp",    "tender_status": 23,    "is_deleted": 34,    "add_time": 4,    "last_time": 87}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// TenderInfo struct is a row record of the tender_info table in the yx database
type TenderInfo struct {
	//[ 0] tender_id                                      ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	TenderID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:tender_id;type:ubigint;" json:"tender_id"` // 招商计划信息表
	//[ 1] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"` // 达人id
	//[ 2] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"` // 所属平台 1：抖音 2：快手
	//[ 3] anchor_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AnchorID int32 `gorm:"column:anchor_id;type:int;default:0;" json:"anchor_id"` // 主播id
	//[ 4] tender_name                                    varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	TenderName string `gorm:"column:tender_name;type:varchar;size:30;" json:"tender_name"` // 招商计划名称
	//[ 5] tender_type                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderType int32 `gorm:"column:tender_type;type:int;default:0;" json:"tender_type"` // 1 专场   2 拼场
	//[ 6] tender_start                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderStart int32 `gorm:"column:tender_start;type:int;default:0;" json:"tender_start"` // 招商计划的开始时间，即报名开始时间
	//[ 7] tender_end                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderEnd int32 `gorm:"column:tender_end;type:int;default:0;" json:"tender_end"` // 招商计划的结束时间，即报名结束时间
	//[ 8] live_start                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LiveStart int32 `gorm:"column:live_start;type:int;default:0;" json:"live_start"` // 直播开始时间
	//[ 9] live_end                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LiveEnd int32 `gorm:"column:live_end;type:int;default:0;" json:"live_end"` // 直播结束时间
	//[10] tender_require                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	TenderRequire sql.NullString `gorm:"column:tender_require;type:text;size:65535;" json:"tender_require"` // 招商计划的其他要求
	//[11] tender_desc                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	TenderDesc sql.NullString `gorm:"column:tender_desc;type:text;size:65535;" json:"tender_desc"` // 招商计划备注
	//[12] tender_status                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderStatus int32 `gorm:"column:tender_status;type:int;default:0;" json:"tender_status"` // 1 草稿 2 待报名 3 报名中 4 报名截止
	//[13] is_deleted                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsDeleted int32 `gorm:"column:is_deleted;type:int;default:0;" json:"is_deleted"` // 是否删除，0 否，正常   1 是，已删除
	//[14] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 创建招商计划的时间
	//[15] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var tender_infoTableInfo = &TableInfo{
	Name: "tender_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "tender_id",
			Comment:            `招商计划信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TenderID",
			GoFieldType:        "uint64",
			JSONFieldName:      "tender_id",
			ProtobufFieldName:  "tender_id",
			ProtobufType:       "uint64",
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
			Name:               "platform_id",
			Comment:            `所属平台 1：抖音 2：快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "int32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "tender_name",
			Comment:            `招商计划名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "TenderName",
			GoFieldType:        "string",
			JSONFieldName:      "tender_name",
			ProtobufFieldName:  "tender_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "tender_type",
			Comment:            `1 专场   2 拼场`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderType",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_type",
			ProtobufFieldName:  "tender_type",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "tender_start",
			Comment:            `招商计划的开始时间，即报名开始时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderStart",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_start",
			ProtobufFieldName:  "tender_start",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "tender_end",
			Comment:            `招商计划的结束时间，即报名结束时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderEnd",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_end",
			ProtobufFieldName:  "tender_end",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "live_start",
			Comment:            `直播开始时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LiveStart",
			GoFieldType:        "int32",
			JSONFieldName:      "live_start",
			ProtobufFieldName:  "live_start",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "live_end",
			Comment:            `直播结束时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LiveEnd",
			GoFieldType:        "int32",
			JSONFieldName:      "live_end",
			ProtobufFieldName:  "live_end",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "tender_require",
			Comment:            `招商计划的其他要求`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "TenderRequire",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tender_require",
			ProtobufFieldName:  "tender_require",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "tender_desc",
			Comment:            `招商计划备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "TenderDesc",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tender_desc",
			ProtobufFieldName:  "tender_desc",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "tender_status",
			Comment:            `1 草稿 2 待报名 3 报名中 4 报名截止`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_status",
			ProtobufFieldName:  "tender_status",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "is_deleted",
			Comment:            `是否删除，0 否，正常   1 是，已删除 `,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsDeleted",
			GoFieldType:        "int32",
			JSONFieldName:      "is_deleted",
			ProtobufFieldName:  "is_deleted",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "add_time",
			Comment:            `创建招商计划的时间`,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TenderInfo) TableName() string {
	return "tender_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TenderInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TenderInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TenderInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TenderInfo) TableInfo() *TableInfo {
	return tender_infoTableInfo
}
