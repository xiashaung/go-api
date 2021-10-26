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


CREATE TABLE `notify_order_dy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `tag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息种类，订单创建消息的tag值为"111"',
  `msg_id` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息唯一id',
  `p_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父订单ID',
  `s_ids` json NOT NULL COMMENT '子订单ID列表',
  `shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `notify_info` json NOT NULL COMMENT '消息内容，业务内容Json串',
  `error_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '错误原因',
  `task_status` smallint(6) NOT NULL DEFAULT '0' COMMENT '任务状态:0-待消费,1-已消费,2-消费异常',
  `try_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '重试次数',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_add_time` (`add_time`),
  KEY `idx_task_status` (`task_status`)
) ENGINE=InnoDB AUTO_INCREMENT=7917875 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='抖音订单通知消息表'

JSON Sample
-------------------------------------
{    "id": 35,    "tag": "SEOnRguoAAFUjlbgESBOZQArl",    "msg_id": "EchvFbyPKgwjKkSsjhdmBcRRO",    "p_id": 9,    "s_ids": "YlDQHdrcVHTAxeqSakRfhOkDD",    "shop_id": 38,    "notify_info": "eylmfiSGARWlogaafYMddfXIa",    "error_msg": "RTVhcPJHHFyWcUBdPuPSlsXZv",    "task_status": 2,    "try_count": 42,    "add_time": 59,    "last_time": 46}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// NotifyOrderDy struct is a row record of the notify_order_dy table in the yx database
type NotifyOrderDy struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 主键
	//[ 1] tag                                            varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Tag string `gorm:"column:tag;type:varchar;size:32;" json:"tag"` // 消息种类，订单创建消息的tag值为"111"
	//[ 2] msg_id                                         varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	MsgID string `gorm:"column:msg_id;type:varchar;size:125;" json:"msg_id"` // 消息唯一id
	//[ 3] p_id                                           bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PID int64 `gorm:"column:p_id;type:bigint;default:0;" json:"p_id"` // 父订单ID
	//[ 4] s_ids                                          json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	SIds string `gorm:"column:s_ids;type:json;" json:"s_ids"` // 子订单ID列表
	//[ 5] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ShopID int64 `gorm:"column:shop_id;type:bigint;default:0;" json:"shop_id"` // 店铺ID
	//[ 6] notify_info                                    json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	NotifyInfo string `gorm:"column:notify_info;type:json;" json:"notify_info"` // 消息内容，业务内容Json串
	//[ 7] error_msg                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ErrorMsg string `gorm:"column:error_msg;type:varchar;size:255;" json:"error_msg"` // 错误原因
	//[ 8] task_status                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	TaskStatus int32 `gorm:"column:task_status;type:smallint;default:0;" json:"task_status"` // 任务状态:0-待消费,1-已消费,2-消费异常
	//[ 9] try_count                                      smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	TryCount int32 `gorm:"column:try_count;type:smallint;default:0;" json:"try_count"` // 重试次数
	//[10] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加时间
	//[11] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 修改时间

}

var notify_order_dyTableInfo = &TableInfo{
	Name: "notify_order_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "tag",
			Comment:            `消息种类，订单创建消息的tag值为"111"`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Tag",
			GoFieldType:        "string",
			JSONFieldName:      "tag",
			ProtobufFieldName:  "tag",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "msg_id",
			Comment:            `消息唯一id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "MsgID",
			GoFieldType:        "string",
			JSONFieldName:      "msg_id",
			ProtobufFieldName:  "msg_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "p_id",
			Comment:            `父订单ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PID",
			GoFieldType:        "int64",
			JSONFieldName:      "p_id",
			ProtobufFieldName:  "p_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "s_ids",
			Comment:            `子订单ID列表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "SIds",
			GoFieldType:        "string",
			JSONFieldName:      "s_ids",
			ProtobufFieldName:  "s_ids",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "shop_id",
			Comment:            `店铺ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "int64",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "notify_info",
			Comment:            `消息内容，业务内容Json串`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "NotifyInfo",
			GoFieldType:        "string",
			JSONFieldName:      "notify_info",
			ProtobufFieldName:  "notify_info",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "error_msg",
			Comment:            `错误原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ErrorMsg",
			GoFieldType:        "string",
			JSONFieldName:      "error_msg",
			ProtobufFieldName:  "error_msg",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "task_status",
			Comment:            `任务状态:0-待消费,1-已消费,2-消费异常`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "TaskStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "task_status",
			ProtobufFieldName:  "task_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "try_count",
			Comment:            `重试次数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "TryCount",
			GoFieldType:        "int32",
			JSONFieldName:      "try_count",
			ProtobufFieldName:  "try_count",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "add_time",
			Comment:            `添加时间`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "last_time",
			Comment:            `修改时间`,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NotifyOrderDy) TableName() string {
	return "notify_order_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NotifyOrderDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NotifyOrderDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NotifyOrderDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NotifyOrderDy) TableInfo() *TableInfo {
	return notify_order_dyTableInfo
}
