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


CREATE TABLE `sample_order` (
  `sample_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '样品订单信息表',
  `sample_sn` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '样品订单编号',
  `talent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `platform_id` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '平台1抖音 2快手',
  `anchor_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '主播id',
  `pool_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '选品池id',
  `total_fee` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单总金额',
  `pay_fee` int(11) unsigned NOT NULL DEFAULT '0' COMMENT ' 支付金额',
  `discount_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '优惠金额',
  `express_fee` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '运费',
  `pay_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '支付方式1微信 2支付宝',
  `delivery_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '收货地址ID',
  `shipping_sn` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '物流单号',
  `shipping_code` int(11) NOT NULL DEFAULT '0' COMMENT '快递公司-code',
  `shipping_company` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递公司',
  `delivery_address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '收货地址',
  `buyer_note_list` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '买家备注',
  `sample_status` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '订单状态;1-待付款,2-待发货,3-已发货,4-确认收货5-订单关闭',
  `agreement_status` tinyint(3) DEFAULT '1' COMMENT '履约状态值，0不需要履约，1未履约，2履约待验收，3不履约待审核，4已履约，5不履约',
  `close_type` tinyint(3) DEFAULT '0' COMMENT '订单关闭类型0未关闭，1未知类型，2自动关闭，3手动关闭',
  `pay_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
  `delivery_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '发货时间',
  `complete_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单完成时间',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '下单时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0',
  `sample_type` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '订单类型  1 付费 2 免费',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '商家备注',
  `estimate_live_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '预估直播时间 0未设置',
  `refuse_type` int(10) NOT NULL DEFAULT '0' COMMENT '拒绝类型 0 其他  1 粉丝画像 2 口碑匹配不高',
  `pass_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '免费样品订单审核通过时间',
  PRIMARY KEY (`sample_id`),
  UNIQUE KEY `uniq_sn` (`sample_sn`) USING BTREE,
  KEY `idx_talent_pid` (`talent_id`,`add_time`) USING BTREE,
  KEY `idx_talent_status` (`talent_id`,`sample_status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=117 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='样品订单信息表'

JSON Sample
-------------------------------------
{    "sample_id": 26,    "sample_sn": "LeBPsOGiMVLXrhrMfWSOCZTJO",    "talent_id": 84,    "platform_id": 18,    "anchor_id": 52,    "pool_id": 74,    "total_fee": 58,    "pay_fee": 12,    "discount_amount": 52,    "express_fee": 50,    "pay_type": 84,    "delivery_id": 55,    "shipping_sn": "RXbsQTYlPKvgNowQuiPurWLnQ",    "shipping_code": 55,    "shipping_company": "WkWXDeeNvdVlTpHnVBhRIpPBt",    "delivery_address": "WBPQPAdKPbrRWqcViQmhlROuH",    "buyer_note_list": "CxpRrwjxJdRPtZwtwMOXIvVLJ",    "sample_status": 46,    "agreement_status": 89,    "close_type": 86,    "pay_time": 67,    "delivery_time": 70,    "complete_time": 19,    "add_time": 37,    "last_time": 7,    "sample_type": 6,    "remark": "KCCZNShdPJRSJaXipttvVJHxh",    "estimate_live_time": 49,    "refuse_type": 46,    "pass_time": 38}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[17] column is set for unsigned
[20] column is set for unsigned
[21] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned
[24] column is set for unsigned
[25] column is set for unsigned
[29] column is set for unsigned



*/

// SampleOrder struct is a row record of the sample_order table in the yx database
type SampleOrder struct {
	//[ 0] sample_id                                      ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	SampleID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:sample_id;type:ubigint;" json:"sample_id"` // 样品订单信息表
	//[ 1] sample_sn                                      varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	SampleSn string `gorm:"column:sample_sn;type:varchar;size:32;" json:"sample_sn"` // 样品订单编号
	//[ 2] talent_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TalentID uint64 `gorm:"column:talent_id;type:ubigint;default:0;" json:"talent_id"` // 达人id
	//[ 3] platform_id                                    usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [1]
	PlatformID uint32 `gorm:"column:platform_id;type:usmallint;default:1;" json:"platform_id"` // 平台1抖音 2快手
	//[ 4] anchor_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	AnchorID uint64 `gorm:"column:anchor_id;type:ubigint;default:0;" json:"anchor_id"` // 主播id
	//[ 5] pool_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	PoolID uint64 `gorm:"column:pool_id;type:ubigint;default:0;" json:"pool_id"` // 选品池id
	//[ 6] total_fee                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TotalFee uint32 `gorm:"column:total_fee;type:uint;default:0;" json:"total_fee"` // 订单总金额
	//[ 7] pay_fee                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayFee uint32 `gorm:"column:pay_fee;type:uint;default:0;" json:"pay_fee"` //  支付金额
	//[ 8] discount_amount                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DiscountAmount uint32 `gorm:"column:discount_amount;type:uint;default:0;" json:"discount_amount"` // 优惠金额
	//[ 9] express_fee                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ExpressFee uint32 `gorm:"column:express_fee;type:uint;default:0;" json:"express_fee"` // 运费
	//[10] pay_type                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	PayType int32 `gorm:"column:pay_type;type:smallint;default:0;" json:"pay_type"` // 支付方式1微信 2支付宝
	//[11] delivery_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	DeliveryID int64 `gorm:"column:delivery_id;type:bigint;default:0;" json:"delivery_id"` // 收货地址ID
	//[12] shipping_sn                                    varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	ShippingSn string `gorm:"column:shipping_sn;type:varchar;size:64;" json:"shipping_sn"` // 物流单号
	//[13] shipping_code                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShippingCode int32 `gorm:"column:shipping_code;type:int;default:0;" json:"shipping_code"` // 快递公司-code
	//[14] shipping_company                               varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	ShippingCompany string `gorm:"column:shipping_company;type:varchar;size:125;" json:"shipping_company"` // 快递公司
	//[15] delivery_address                               text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	DeliveryAddress sql.NullString `gorm:"column:delivery_address;type:text;size:65535;" json:"delivery_address"` // 收货地址
	//[16] buyer_note_list                                text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	BuyerNoteList sql.NullString `gorm:"column:buyer_note_list;type:text;size:65535;" json:"buyer_note_list"` // 买家备注
	//[17] sample_status                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	SampleStatus uint32 `gorm:"column:sample_status;type:uint;default:1;" json:"sample_status"` // 订单状态;1-待付款,2-待发货,3-已发货,4-确认收货5-订单关闭
	//[18] agreement_status                               tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	AgreementStatus sql.NullInt64 `gorm:"column:agreement_status;type:tinyint;default:1;" json:"agreement_status"` // 履约状态值，0不需要履约，1未履约，2履约待验收，3不履约待审核，4已履约，5不履约
	//[19] close_type                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	CloseType sql.NullInt64 `gorm:"column:close_type;type:tinyint;default:0;" json:"close_type"` // 订单关闭类型0未关闭，1未知类型，2自动关闭，3手动关闭
	//[20] pay_time                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	PayTime uint64 `gorm:"column:pay_time;type:ubigint;default:0;" json:"pay_time"` // 支付时间
	//[21] delivery_time                                  ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	DeliveryTime uint64 `gorm:"column:delivery_time;type:ubigint;default:0;" json:"delivery_time"` // 发货时间
	//[22] complete_time                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CompleteTime uint32 `gorm:"column:complete_time;type:uint;default:0;" json:"complete_time"` // 订单完成时间
	//[23] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 下单时间
	//[24] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
	//[25] sample_type                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	SampleType uint32 `gorm:"column:sample_type;type:uint;default:1;" json:"sample_type"` // 订单类型  1 付费 2 免费
	//[26] remark                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Remark string `gorm:"column:remark;type:varchar;size:255;" json:"remark"` // 商家备注
	//[27] estimate_live_time                             bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	EstimateLiveTime int64 `gorm:"column:estimate_live_time;type:bigint;default:0;" json:"estimate_live_time"` // 预估直播时间 0未设置
	//[28] refuse_type                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RefuseType int32 `gorm:"column:refuse_type;type:int;default:0;" json:"refuse_type"` // 拒绝类型 0 其他  1 粉丝画像 2 口碑匹配不高
	//[29] pass_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PassTime uint32 `gorm:"column:pass_time;type:uint;default:0;" json:"pass_time"` // 免费样品订单审核通过时间

}

var sample_orderTableInfo = &TableInfo{
	Name: "sample_order",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "sample_id",
			Comment:            `样品订单信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SampleID",
			GoFieldType:        "uint64",
			JSONFieldName:      "sample_id",
			ProtobufFieldName:  "sample_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "sample_sn",
			Comment:            `样品订单编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "SampleSn",
			GoFieldType:        "string",
			JSONFieldName:      "sample_sn",
			ProtobufFieldName:  "sample_sn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "platform_id",
			Comment:            `平台1抖音 2快手`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "uint32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "uint64",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "pool_id",
			Comment:            `选品池id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "PoolID",
			GoFieldType:        "uint64",
			JSONFieldName:      "pool_id",
			ProtobufFieldName:  "pool_id",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "total_fee",
			Comment:            `订单总金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TotalFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "total_fee",
			ProtobufFieldName:  "total_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "pay_fee",
			Comment:            ` 支付金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PayFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "pay_fee",
			ProtobufFieldName:  "pay_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "discount_amount",
			Comment:            `优惠金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DiscountAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "discount_amount",
			ProtobufFieldName:  "discount_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "express_fee",
			Comment:            `运费`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "express_fee",
			ProtobufFieldName:  "express_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "pay_type",
			Comment:            `支付方式1微信 2支付宝`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "PayType",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_type",
			ProtobufFieldName:  "pay_type",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "delivery_id",
			Comment:            `收货地址ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "DeliveryID",
			GoFieldType:        "int64",
			JSONFieldName:      "delivery_id",
			ProtobufFieldName:  "delivery_id",
			ProtobufType:       "int64",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "shipping_sn",
			Comment:            `物流单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "ShippingSn",
			GoFieldType:        "string",
			JSONFieldName:      "shipping_sn",
			ProtobufFieldName:  "shipping_sn",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "shipping_code",
			Comment:            `快递公司-code`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ShippingCode",
			GoFieldType:        "int32",
			JSONFieldName:      "shipping_code",
			ProtobufFieldName:  "shipping_code",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "shipping_company",
			Comment:            `快递公司`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "ShippingCompany",
			GoFieldType:        "string",
			JSONFieldName:      "shipping_company",
			ProtobufFieldName:  "shipping_company",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "delivery_address",
			Comment:            `收货地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "DeliveryAddress",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "delivery_address",
			ProtobufFieldName:  "delivery_address",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "buyer_note_list",
			Comment:            `买家备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "BuyerNoteList",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "buyer_note_list",
			ProtobufFieldName:  "buyer_note_list",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "sample_status",
			Comment:            `订单状态;1-待付款,2-待发货,3-已发货,4-确认收货5-订单关闭`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SampleStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "sample_status",
			ProtobufFieldName:  "sample_status",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "agreement_status",
			Comment:            `履约状态值，0不需要履约，1未履约，2履约待验收，3不履约待审核，4已履约，5不履约`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "AgreementStatus",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "agreement_status",
			ProtobufFieldName:  "agreement_status",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "close_type",
			Comment:            `订单关闭类型0未关闭，1未知类型，2自动关闭，3手动关闭`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "CloseType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "close_type",
			ProtobufFieldName:  "close_type",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "pay_time",
			Comment:            `支付时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "PayTime",
			GoFieldType:        "uint64",
			JSONFieldName:      "pay_time",
			ProtobufFieldName:  "pay_time",
			ProtobufType:       "uint64",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "delivery_time",
			Comment:            `发货时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "DeliveryTime",
			GoFieldType:        "uint64",
			JSONFieldName:      "delivery_time",
			ProtobufFieldName:  "delivery_time",
			ProtobufType:       "uint64",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "complete_time",
			Comment:            `订单完成时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CompleteTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "complete_time",
			ProtobufFieldName:  "complete_time",
			ProtobufType:       "uint32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "add_time",
			Comment:            `下单时间`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "sample_type",
			Comment:            `订单类型  1 付费 2 免费`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SampleType",
			GoFieldType:        "uint32",
			JSONFieldName:      "sample_type",
			ProtobufFieldName:  "sample_type",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "remark",
			Comment:            `商家备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "estimate_live_time",
			Comment:            `预估直播时间 0未设置`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "EstimateLiveTime",
			GoFieldType:        "int64",
			JSONFieldName:      "estimate_live_time",
			ProtobufFieldName:  "estimate_live_time",
			ProtobufType:       "int64",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "refuse_type",
			Comment:            `拒绝类型 0 其他  1 粉丝画像 2 口碑匹配不高`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RefuseType",
			GoFieldType:        "int32",
			JSONFieldName:      "refuse_type",
			ProtobufFieldName:  "refuse_type",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "pass_time",
			Comment:            `免费样品订单审核通过时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PassTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "pass_time",
			ProtobufFieldName:  "pass_time",
			ProtobufType:       "uint32",
			ProtobufPos:        30,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SampleOrder) TableName() string {
	return "sample_order"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SampleOrder) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SampleOrder) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SampleOrder) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SampleOrder) TableInfo() *TableInfo {
	return sample_orderTableInfo
}
