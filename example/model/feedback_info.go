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


CREATE TABLE `feedback_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '反馈人id根据user_type判断',
  `user_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '用户类型 0游客 1达人 2商家',
  `type` smallint(6) NOT NULL DEFAULT '1' COMMENT '反馈类型 1 功能异常 2 优化建议 3 其他',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '反馈信息',
  `imgs` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '图片信息数组         ',
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户反馈信息'

JSON Sample
-------------------------------------
{    "id": 15,    "user_id": 28,    "user_type": 90,    "type": 87,    "msg": "VkpOUGnnKtiVqcArAtjfNLKGD",    "imgs": "CjZnbFLQiMNQMBFIaMViMVvFB",    "add_time": 29,    "last_time": 93}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// FeedbackInfo struct is a row record of the feedback_info table in the yx database
type FeedbackInfo struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] user_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UserID int32 `gorm:"column:user_id;type:int;default:0;" json:"user_id"` // 反馈人id根据user_type判断
	//[ 2] user_type                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	UserType int32 `gorm:"column:user_type;type:tinyint;default:0;" json:"user_type"` // 用户类型 0游客 1达人 2商家
	//[ 3] type                                           smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [1]
	Type int32 `gorm:"column:type;type:smallint;default:1;" json:"type"` // 反馈类型 1 功能异常 2 优化建议 3 其他
	//[ 4] msg                                            varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Msg string `gorm:"column:msg;type:varchar;size:255;" json:"msg"` // 反馈信息
	//[ 5] imgs                                           text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Imgs string `gorm:"column:imgs;type:text;size:65535;" json:"imgs"` // 图片信息数组
	//[ 6] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 7] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var feedback_infoTableInfo = &TableInfo{
	Name: "feedback_info",
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
			Name:               "user_id",
			Comment:            `反馈人id根据user_type判断`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "user_type",
			Comment:            `用户类型 0游客 1达人 2商家`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "UserType",
			GoFieldType:        "int32",
			JSONFieldName:      "user_type",
			ProtobufFieldName:  "user_type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "type",
			Comment:            `反馈类型 1 功能异常 2 优化建议 3 其他`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "int32",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "msg",
			Comment:            `反馈信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Msg",
			GoFieldType:        "string",
			JSONFieldName:      "msg",
			ProtobufFieldName:  "msg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "imgs",
			Comment:            `图片信息数组         `,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Imgs",
			GoFieldType:        "string",
			JSONFieldName:      "imgs",
			ProtobufFieldName:  "imgs",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FeedbackInfo) TableName() string {
	return "feedback_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FeedbackInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FeedbackInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FeedbackInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FeedbackInfo) TableInfo() *TableInfo {
	return feedback_infoTableInfo
}
