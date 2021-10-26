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


CREATE TABLE `trade_info` (
  `trade_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `trade_sn` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '交易流水号',
  `order_sn` varchar(50) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL COMMENT '订单号',
  `trans_action` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1:收入 2:支出',
  `trans_status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '0未完成 1已完成',
  `amount` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '金额(分)',
  `pay_type` tinyint(1) unsigned NOT NULL COMMENT '1微信2支付宝',
  `desc` text CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '备注',
  `merchant_id` bigint(10) unsigned NOT NULL COMMENT '商家id',
  `talent_id` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `admin_id` bigint(10) unsigned NOT NULL DEFAULT '0' COMMENT '运营后台id',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`trade_id`),
  UNIQUE KEY `uniq_tradesn` (`trade_sn`) USING BTREE,
  KEY `idx_ordersn_talentid` (`order_sn`,`talent_id`) USING BTREE,
  KEY `idx_ordersn_adminid` (`order_sn`,`admin_id`) USING BTREE,
  KEY `idx_adminid` (`admin_id`,`trans_status`) USING BTREE,
  KEY `idx_merchantid` (`merchant_id`,`trans_status`) USING BTREE,
  KEY `idx_ordersn_merchantid` (`order_sn`,`merchant_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "trade_id": 75,    "trade_sn": "WeIZSegnhRMVOCCtsGVcWvVOM",    "order_sn": "cxVCeSHUjABDISnChSMaAMllD",    "trans_action": 12,    "trans_status": 26,    "amount": 21,    "pay_type": 92,    "desc": "eJLgqlwhiQLssmDkmtGbiKWuV",    "merchant_id": 2,    "talent_id": 12,    "admin_id": 33,    "add_time": 97,    "last_time": 69}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned



*/

// TradeInfo struct is a row record of the trade_info table in the yx database
type TradeInfo struct {
	//[ 0] trade_id                                       ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	TradeID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:trade_id;type:ubigint;" json:"trade_id"`
	//[ 1] trade_sn                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	TradeSn string `gorm:"column:trade_sn;type:varchar;size:50;" json:"trade_sn"` // 交易流水号
	//[ 2] order_sn                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	OrderSn string `gorm:"column:order_sn;type:varchar;size:50;" json:"order_sn"` // 订单号
	//[ 3] trans_action                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [1]
	TransAction uint32 `gorm:"column:trans_action;type:utinyint;default:1;" json:"trans_action"` // 1:收入 2:支出
	//[ 4] trans_status                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	TransStatus uint32 `gorm:"column:trans_status;type:utinyint;default:0;" json:"trans_status"` // 0未完成 1已完成
	//[ 5] amount                                         ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	Amount uint64 `gorm:"column:amount;type:ubigint;default:0;" json:"amount"` // 金额(分)
	//[ 6] pay_type                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: []
	PayType uint32 `gorm:"column:pay_type;type:utinyint;" json:"pay_type"` // 1微信2支付宝
	//[ 7] desc                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Desc sql.NullString `gorm:"column:desc;type:text;size:65535;" json:"desc"` // 备注
	//[ 8] merchant_id                                    ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: []
	MerchantID uint64 `gorm:"column:merchant_id;type:ubigint;" json:"merchant_id"` // 商家id
	//[ 9] talent_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TalentID uint64 `gorm:"column:talent_id;type:ubigint;default:0;" json:"talent_id"` // 达人id
	//[10] admin_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	AdminID uint64 `gorm:"column:admin_id;type:ubigint;default:0;" json:"admin_id"` // 运营后台id
	//[11] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[12] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 修改时间

}

var trade_infoTableInfo = &TableInfo{
	Name: "trade_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "trade_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TradeID",
			GoFieldType:        "uint64",
			JSONFieldName:      "trade_id",
			ProtobufFieldName:  "trade_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "trade_sn",
			Comment:            `交易流水号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "TradeSn",
			GoFieldType:        "string",
			JSONFieldName:      "trade_sn",
			ProtobufFieldName:  "trade_sn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "order_sn",
			Comment:            `订单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "OrderSn",
			GoFieldType:        "string",
			JSONFieldName:      "order_sn",
			ProtobufFieldName:  "order_sn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "trans_action",
			Comment:            `1:收入 2:支出`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "TransAction",
			GoFieldType:        "uint32",
			JSONFieldName:      "trans_action",
			ProtobufFieldName:  "trans_action",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "trans_status",
			Comment:            `0未完成 1已完成`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "TransStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "trans_status",
			ProtobufFieldName:  "trans_status",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "amount",
			Comment:            `金额(分)`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "Amount",
			GoFieldType:        "uint64",
			JSONFieldName:      "amount",
			ProtobufFieldName:  "amount",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "pay_type",
			Comment:            `1微信2支付宝`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "PayType",
			GoFieldType:        "uint32",
			JSONFieldName:      "pay_type",
			ProtobufFieldName:  "pay_type",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "desc",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Desc",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "desc",
			ProtobufFieldName:  "desc",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "merchant_id",
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
			GoFieldName:        "MerchantID",
			GoFieldType:        "uint64",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "uint64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint64",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "admin_id",
			Comment:            `运营后台id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "AdminID",
			GoFieldType:        "uint64",
			JSONFieldName:      "admin_id",
			ProtobufFieldName:  "admin_id",
			ProtobufType:       "uint64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "add_time",
			Comment:            `创建时间`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "last_time",
			Comment:            `修改时间`,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TradeInfo) TableName() string {
	return "trade_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TradeInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TradeInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TradeInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TradeInfo) TableInfo() *TableInfo {
	return trade_infoTableInfo
}
