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


CREATE TABLE `live_info` (
  `live_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '直播场次信息表',
  `plan_id` int(11) NOT NULL DEFAULT '0' COMMENT '直播计划id',
  `platform_id` int(1) NOT NULL DEFAULT '0' COMMENT '所属平台 1：抖音 2：快手',
  `anchor_id` int(11) NOT NULL DEFAULT '0' COMMENT '主播id',
  `live_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '直播场次名称',
  `live_type` int(1) NOT NULL DEFAULT '0' COMMENT '直播场次类型	1：专场 2：拼场',
  `live_start` int(10) NOT NULL DEFAULT '0' COMMENT '直播开始时间',
  `live_end` int(10) NOT NULL DEFAULT '0' COMMENT '直播结束时间',
  `live_status` int(1) NOT NULL DEFAULT '1' COMMENT '直播场次状态	0 取消  1 待确定 2 待直播 3 直播中 4 已经结束',
  `brand_id` int(11) DEFAULT '0' COMMENT '品牌id  0：待确定',
  `product_num` int(11) NOT NULL DEFAULT '0' COMMENT '本场次直播确认上橱窗的商品数量',
  `live_desc` text COMMENT '直播场次备注信息',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '直播场次创建时间',
  `last_time` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`live_id`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='直播场次信息表'

JSON Sample
-------------------------------------
{    "live_id": 93,    "plan_id": 19,    "platform_id": 31,    "anchor_id": 98,    "live_name": "TJnriotJKbKghpabeCrmebmCF",    "live_type": 1,    "live_start": 35,    "live_end": 17,    "live_status": 28,    "brand_id": 61,    "product_num": 59,    "live_desc": "XGYrTuwHbENnPHERYwJKApVoI",    "add_time": 42,    "last_time": 37}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// LiveInfo struct is a row record of the live_info table in the yx database
type LiveInfo struct {
	//[ 0] live_id                                        ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	LiveID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:live_id;type:ubigint;" json:"live_id"` // 直播场次信息表
	//[ 1] plan_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlanID int32 `gorm:"column:plan_id;type:int;default:0;" json:"plan_id"` // 直播计划id
	//[ 2] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"` // 所属平台 1：抖音 2：快手
	//[ 3] anchor_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AnchorID int32 `gorm:"column:anchor_id;type:int;default:0;" json:"anchor_id"` // 主播id
	//[ 4] live_name                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	LiveName string `gorm:"column:live_name;type:varchar;size:255;" json:"live_name"` // 直播场次名称
	//[ 5] live_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LiveType int32 `gorm:"column:live_type;type:int;default:0;" json:"live_type"` // 直播场次类型	1：专场 2：拼场
	//[ 6] live_start                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LiveStart int32 `gorm:"column:live_start;type:int;default:0;" json:"live_start"` // 直播开始时间
	//[ 7] live_end                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LiveEnd int32 `gorm:"column:live_end;type:int;default:0;" json:"live_end"` // 直播结束时间
	//[ 8] live_status                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	LiveStatus int32 `gorm:"column:live_status;type:int;default:1;" json:"live_status"` // 直播场次状态	0 取消  1 待确定 2 待直播 3 直播中 4 已经结束
	//[ 9] brand_id                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	BrandID sql.NullInt64 `gorm:"column:brand_id;type:int;default:0;" json:"brand_id"` // 品牌id  0：待确定
	//[10] product_num                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductNum int32 `gorm:"column:product_num;type:int;default:0;" json:"product_num"` // 本场次直播确认上橱窗的商品数量
	//[11] live_desc                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	LiveDesc sql.NullString `gorm:"column:live_desc;type:text;size:65535;" json:"live_desc"` // 直播场次备注信息
	//[12] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 直播场次创建时间
	//[13] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var live_infoTableInfo = &TableInfo{
	Name: "live_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "live_id",
			Comment:            `直播场次信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "LiveID",
			GoFieldType:        "uint64",
			JSONFieldName:      "live_id",
			ProtobufFieldName:  "live_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "plan_id",
			Comment:            `直播计划id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlanID",
			GoFieldType:        "int32",
			JSONFieldName:      "plan_id",
			ProtobufFieldName:  "plan_id",
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
			Name:               "live_name",
			Comment:            `直播场次名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "LiveName",
			GoFieldType:        "string",
			JSONFieldName:      "live_name",
			ProtobufFieldName:  "live_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "live_type",
			Comment:            `直播场次类型	1：专场 2：拼场`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LiveType",
			GoFieldType:        "int32",
			JSONFieldName:      "live_type",
			ProtobufFieldName:  "live_type",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "live_status",
			Comment:            `直播场次状态	0 取消  1 待确定 2 待直播 3 直播中 4 已经结束`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LiveStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "live_status",
			ProtobufFieldName:  "live_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "brand_id",
			Comment:            `品牌id  0：待确定`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BrandID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "brand_id",
			ProtobufFieldName:  "brand_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "product_num",
			Comment:            `本场次直播确认上橱窗的商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductNum",
			GoFieldType:        "int32",
			JSONFieldName:      "product_num",
			ProtobufFieldName:  "product_num",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "live_desc",
			Comment:            `直播场次备注信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "LiveDesc",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "live_desc",
			ProtobufFieldName:  "live_desc",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "add_time",
			Comment:            `直播场次创建时间`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LiveInfo) TableName() string {
	return "live_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LiveInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LiveInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LiveInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LiveInfo) TableInfo() *TableInfo {
	return live_infoTableInfo
}
