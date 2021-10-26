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


CREATE TABLE `notify_product_ks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `event_id` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息唯一id',
  `biz_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '业务id如订单id、退款单id、商品id',
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商家id',
  `open_id` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '买家openId',
  `app_key` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '应用id',
  `event` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '消息标示',
  `info` json NOT NULL COMMENT '消息内容，业务内容Json串',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间-ks',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间-ks',
  `error_msg` varchar(255) NOT NULL DEFAULT '' COMMENT '错误原因',
  `status` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否消费:0-待消费,1-已消费,2-消费异常',
  `try_count` smallint(6) NOT NULL DEFAULT '0' COMMENT '重试次数',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_event_id` (`event_id`),
  KEY `idx_add_time` (`add_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='快手订单通知消息表'

JSON Sample
-------------------------------------
{    "id": 43,    "event_id": "FrBEJqHRdAoKNJRjeFYVgHOqN",    "biz_id": 2,    "user_id": 47,    "open_id": "SgEVStQiBLGUlZSbKSdwakKZF",    "app_key": "dNQgaOyxZSePUHhrwwvgCrEKL",    "event": "aNCOqEXPmJjsdKkrqvxHoIGEJ",    "info": "pFCjhDIQVMXnBITqTdYmIkKht",    "create_time": 99,    "update_time": 45,    "error_msg": "rkUTTQwwudIpYcCxMfdHecuxq",    "status": 31,    "try_count": 78,    "add_time": 31,    "last_time": 77}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned



*/

// NotifyProductKs struct is a row record of the notify_product_ks table in the yx database
type NotifyProductKs struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 主键
	//[ 1] event_id                                       varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	EventID string `gorm:"column:event_id;type:varchar;size:125;" json:"event_id"` // 消息唯一id
	//[ 2] biz_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	BizID int64 `gorm:"column:biz_id;type:bigint;default:0;" json:"biz_id"` // 业务id如订单id、退款单id、商品id
	//[ 3] user_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	UserID uint64 `gorm:"column:user_id;type:ubigint;default:0;" json:"user_id"` // 商家id
	//[ 4] open_id                                        varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	OpenID string `gorm:"column:open_id;type:varchar;size:125;" json:"open_id"` // 买家openId
	//[ 5] app_key                                        varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	AppKey string `gorm:"column:app_key;type:varchar;size:125;" json:"app_key"` // 应用id
	//[ 6] event                                          varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	Event string `gorm:"column:event;type:varchar;size:125;" json:"event"` // 消息标示
	//[ 7] info                                           json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Info string `gorm:"column:info;type:json;" json:"info"` // 消息内容，业务内容Json串
	//[ 8] create_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	CreateTime int64 `gorm:"column:create_time;type:bigint;default:0;" json:"create_time"` // 创建时间-ks
	//[ 9] update_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	UpdateTime int64 `gorm:"column:update_time;type:bigint;default:0;" json:"update_time"` // 更新时间-ks
	//[10] error_msg                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ErrorMsg string `gorm:"column:error_msg;type:varchar;size:255;" json:"error_msg"` // 错误原因
	//[11] status                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Status int32 `gorm:"column:status;type:smallint;default:0;" json:"status"` // 是否消费:0-待消费,1-已消费,2-消费异常
	//[12] try_count                                      smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	TryCount int32 `gorm:"column:try_count;type:smallint;default:0;" json:"try_count"` // 重试次数
	//[13] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加时间
	//[14] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 修改时间

}

var notify_product_ksTableInfo = &TableInfo{
	Name: "notify_product_ks",
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
			Name:               "event_id",
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
			GoFieldName:        "EventID",
			GoFieldType:        "string",
			JSONFieldName:      "event_id",
			ProtobufFieldName:  "event_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "biz_id",
			Comment:            `业务id如订单id、退款单id、商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "BizID",
			GoFieldType:        "int64",
			JSONFieldName:      "biz_id",
			ProtobufFieldName:  "biz_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "user_id",
			Comment:            `商家id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "uint64",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "open_id",
			Comment:            `买家openId`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "OpenID",
			GoFieldType:        "string",
			JSONFieldName:      "open_id",
			ProtobufFieldName:  "open_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "app_key",
			Comment:            `应用id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "AppKey",
			GoFieldType:        "string",
			JSONFieldName:      "app_key",
			ProtobufFieldName:  "app_key",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "event",
			Comment:            `消息标示`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "Event",
			GoFieldType:        "string",
			JSONFieldName:      "event",
			ProtobufFieldName:  "event",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "info",
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
			GoFieldName:        "Info",
			GoFieldType:        "string",
			JSONFieldName:      "info",
			ProtobufFieldName:  "info",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "create_time",
			Comment:            `创建时间-ks`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "int64",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "int64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "update_time",
			Comment:            `更新时间-ks`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "int64",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "int64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "status",
			Comment:            `是否消费:0-待消费,1-已消费,2-消费异常`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (n *NotifyProductKs) TableName() string {
	return "notify_product_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NotifyProductKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NotifyProductKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NotifyProductKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (n *NotifyProductKs) TableInfo() *TableInfo {
	return notify_product_ksTableInfo
}
