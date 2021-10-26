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


CREATE TABLE `oauth_task` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号ID 商家 达人主键ID',
  `anchor_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播ID',
  `shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `open_id` varchar(50) NOT NULL DEFAULT '' COMMENT 'openId 第三方openID',
  `platform_id` smallint(2) NOT NULL DEFAULT '0' COMMENT '平台ID 1:抖音 2:快手',
  `oauth_type` smallint(2) NOT NULL DEFAULT '0' COMMENT '授权类型 1:商家 2:达人',
  `seller_order_status` smallint(4) NOT NULL DEFAULT '0' COMMENT '商家订单任务',
  `seller_status` smallint(2) NOT NULL DEFAULT '0' COMMENT '商家是否处理 0:未处理 1:已处理',
  `dr_status` smallint(2) NOT NULL DEFAULT '0' COMMENT '达人是否处理 0:未处理 1:已处理',
  `task_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_open_type` (`open_id`,`oauth_type`) USING BTREE,
  KEY `idx_add_time` (`add_time`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 9,    "account_id": 21,    "anchor_id": 42,    "shop_id": 20,    "open_id": "lQLqhOSGpAjXQUEDrTBUXVxGY",    "platform_id": 58,    "oauth_type": 56,    "seller_order_status": 36,    "seller_status": 10,    "dr_status": 53,    "task_params": "KglegNPGcvclMMvYKWILYgmOE",    "add_time": 91,    "last_time": 93}



*/

// OauthTask struct is a row record of the oauth_task table in the yx database
type OauthTask struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] account_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AccountID int64 `gorm:"column:account_id;type:bigint;default:0;" json:"account_id"` // 账号ID 商家 达人主键ID
	//[ 2] anchor_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AnchorID int64 `gorm:"column:anchor_id;type:bigint;default:0;" json:"anchor_id"` // 主播ID
	//[ 3] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ShopID int64 `gorm:"column:shop_id;type:bigint;default:0;" json:"shop_id"` // 店铺ID
	//[ 4] open_id                                        varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	OpenID string `gorm:"column:open_id;type:varchar;size:50;" json:"open_id"` // openId 第三方openID
	//[ 5] platform_id                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:smallint;default:0;" json:"platform_id"` // 平台ID 1:抖音 2:快手
	//[ 6] oauth_type                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	OauthType int32 `gorm:"column:oauth_type;type:smallint;default:0;" json:"oauth_type"` // 授权类型 1:商家 2:达人
	//[ 7] seller_order_status                            smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	SellerOrderStatus int32 `gorm:"column:seller_order_status;type:smallint;default:0;" json:"seller_order_status"` // 商家订单任务
	//[ 8] seller_status                                  smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	SellerStatus int32 `gorm:"column:seller_status;type:smallint;default:0;" json:"seller_status"` // 商家是否处理 0:未处理 1:已处理
	//[ 9] dr_status                                      smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	DrStatus int32 `gorm:"column:dr_status;type:smallint;default:0;" json:"dr_status"` // 达人是否处理 0:未处理 1:已处理
	//[10] task_params                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	TaskParams sql.NullString `gorm:"column:task_params;type:text;size:65535;" json:"task_params"`
	//[11] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[12] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var oauth_taskTableInfo = &TableInfo{
	Name: "oauth_task",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "account_id",
			Comment:            `账号ID 商家 达人主键ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AccountID",
			GoFieldType:        "int64",
			JSONFieldName:      "account_id",
			ProtobufFieldName:  "account_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "anchor_id",
			Comment:            `主播ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "int64",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "open_id",
			Comment:            `openId 第三方openID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "OpenID",
			GoFieldType:        "string",
			JSONFieldName:      "open_id",
			ProtobufFieldName:  "open_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "platform_id",
			Comment:            `平台ID 1:抖音 2:快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "oauth_type",
			Comment:            `授权类型 1:商家 2:达人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "OauthType",
			GoFieldType:        "int32",
			JSONFieldName:      "oauth_type",
			ProtobufFieldName:  "oauth_type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "seller_order_status",
			Comment:            `商家订单任务`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "SellerOrderStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "seller_order_status",
			ProtobufFieldName:  "seller_order_status",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "seller_status",
			Comment:            `商家是否处理 0:未处理 1:已处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "SellerStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "seller_status",
			ProtobufFieldName:  "seller_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "dr_status",
			Comment:            `达人是否处理 0:未处理 1:已处理`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "DrStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "dr_status",
			ProtobufFieldName:  "dr_status",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "task_params",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "TaskParams",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "task_params",
			ProtobufFieldName:  "task_params",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OauthTask) TableName() string {
	return "oauth_task"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OauthTask) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OauthTask) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OauthTask) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OauthTask) TableInfo() *TableInfo {
	return oauth_taskTableInfo
}
