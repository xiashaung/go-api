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


CREATE TABLE `order_info` (
  `order_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '订单信息表',
  `platform_id` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '所属平台 1：抖音 2：快手',
  `order_sn` varchar(64) NOT NULL COMMENT '订单编号',
  `talent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `anchor_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '主播id',
  `merchant_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商家id',
  `shop_id` int(11) NOT NULL COMMENT '店铺id',
  `pay_fee` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额',
  `pay_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
  `pay_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '支付方式',
  `cps_type` smallint(6) NOT NULL COMMENT '分销类型',
  `order_status` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态 1 待付款  2 待发货  3 待收货   4已完成',
  `express_fee` int(11) NOT NULL DEFAULT '0' COMMENT '运费',
  `discount_fee` int(11) NOT NULL DEFAULT '0' COMMENT '优惠金额',
  `deliver_time` int(11) NOT NULL DEFAULT '0' COMMENT '发货时间',
  `express_no` int(11) NOT NULL DEFAULT '0' COMMENT '快递单号',
  `express_code` int(11) NOT NULL DEFAULT '0' COMMENT '快递公司 code',
  `mobile` char(11) NOT NULL COMMENT '手机号',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '收货地址',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '买家留言',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单拉取时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单更新时间',
  PRIMARY KEY (`order_id`) USING BTREE,
  KEY `idx_shop_pay` (`shop_id`,`pay_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单信息表'

JSON Sample
-------------------------------------
{    "order_id": 20,    "platform_id": 61,    "order_sn": "GcYyfmBKVjoWJmKudpytrrZTW",    "talent_id": 90,    "anchor_id": 35,    "merchant_id": 24,    "shop_id": 55,    "pay_fee": 40,    "pay_time": 69,    "pay_type": 13,    "cps_type": 57,    "order_status": 54,    "express_fee": 49,    "discount_fee": 8,    "deliver_time": 7,    "express_no": 9,    "express_code": 46,    "mobile": "scujVjbGIwCkIdvrhRtojvZZx",    "address": "yMxSSTnwnLvfeLUjhRrZESwfI",    "remark": "fyWnIcIbwBKjaPSDjZHCbIJsm",    "add_time": 53,    "last_time": 83}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[11] column is set for unsigned
[20] column is set for unsigned
[21] column is set for unsigned



*/

// OrderInfo struct is a row record of the order_info table in the yx database
type OrderInfo struct {
	//[ 0] order_id                                       bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	OrderID int64 `gorm:"primary_key;AUTO_INCREMENT;column:order_id;type:bigint;" json:"order_id"` // 订单信息表
	//[ 1] platform_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	PlatformID uint32 `gorm:"column:platform_id;type:uint;default:1;" json:"platform_id"` // 所属平台 1：抖音 2：快手
	//[ 2] order_sn                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	OrderSn string `gorm:"column:order_sn;type:varchar;size:64;" json:"order_sn"` // 订单编号
	//[ 3] talent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TalentID uint32 `gorm:"column:talent_id;type:uint;default:0;" json:"talent_id"` // 达人id
	//[ 4] anchor_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AnchorID uint32 `gorm:"column:anchor_id;type:uint;default:0;" json:"anchor_id"` // 主播id
	//[ 5] merchant_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	MerchantID uint32 `gorm:"column:merchant_id;type:uint;default:0;" json:"merchant_id"` // 商家id
	//[ 6] shop_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ShopID int32 `gorm:"column:shop_id;type:int;" json:"shop_id"` // 店铺id
	//[ 7] pay_fee                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayFee uint32 `gorm:"column:pay_fee;type:uint;default:0;" json:"pay_fee"` // 支付金额
	//[ 8] pay_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayTime uint32 `gorm:"column:pay_time;type:uint;default:0;" json:"pay_time"` // 支付时间
	//[ 9] pay_type                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	PayType int32 `gorm:"column:pay_type;type:smallint;default:0;" json:"pay_type"` // 支付方式
	//[10] cps_type                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: []
	CpsType int32 `gorm:"column:cps_type;type:smallint;" json:"cps_type"` // 分销类型
	//[11] order_status                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	OrderStatus uint32 `gorm:"column:order_status;type:uint;default:0;" json:"order_status"` // 订单状态 1 待付款  2 待发货  3 待收货   4已完成
	//[12] express_fee                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExpressFee int32 `gorm:"column:express_fee;type:int;default:0;" json:"express_fee"` // 运费
	//[13] discount_fee                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DiscountFee int32 `gorm:"column:discount_fee;type:int;default:0;" json:"discount_fee"` // 优惠金额
	//[14] deliver_time                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DeliverTime int32 `gorm:"column:deliver_time;type:int;default:0;" json:"deliver_time"` // 发货时间
	//[15] express_no                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExpressNo int32 `gorm:"column:express_no;type:int;default:0;" json:"express_no"` // 快递单号
	//[16] express_code                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExpressCode int32 `gorm:"column:express_code;type:int;default:0;" json:"express_code"` // 快递公司 code
	//[17] mobile                                         char(11)             null: false  primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	Mobile string `gorm:"column:mobile;type:char;size:11;" json:"mobile"` // 手机号
	//[18] address                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Address string `gorm:"column:address;type:varchar;size:255;" json:"address"` // 收货地址
	//[19] remark                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Remark string `gorm:"column:remark;type:varchar;size:255;" json:"remark"` // 买家留言
	//[20] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 订单拉取时间
	//[21] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 订单更新时间

}

var order_infoTableInfo = &TableInfo{
	Name: "order_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "order_id",
			Comment:            `订单信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "OrderID",
			GoFieldType:        "int64",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "platform_id",
			Comment:            `所属平台 1：抖音 2：快手`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "uint32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "order_sn",
			Comment:            `订单编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "OrderSn",
			GoFieldType:        "string",
			JSONFieldName:      "order_sn",
			ProtobufFieldName:  "order_sn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "uint32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "merchant_id",
			Comment:            `商家id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "uint32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "shop_id",
			Comment:            `店铺id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "int32",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "pay_fee",
			Comment:            `支付金额`,
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
			Name:               "pay_time",
			Comment:            `支付时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PayTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "pay_time",
			ProtobufFieldName:  "pay_time",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "pay_type",
			Comment:            `支付方式`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "cps_type",
			Comment:            `分销类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "CpsType",
			GoFieldType:        "int32",
			JSONFieldName:      "cps_type",
			ProtobufFieldName:  "cps_type",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "order_status",
			Comment:            `订单状态 1 待付款  2 待发货  3 待收货   4已完成`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "OrderStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "order_status",
			ProtobufFieldName:  "order_status",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "express_fee",
			Comment:            `运费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExpressFee",
			GoFieldType:        "int32",
			JSONFieldName:      "express_fee",
			ProtobufFieldName:  "express_fee",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "discount_fee",
			Comment:            `优惠金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DiscountFee",
			GoFieldType:        "int32",
			JSONFieldName:      "discount_fee",
			ProtobufFieldName:  "discount_fee",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "deliver_time",
			Comment:            `发货时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DeliverTime",
			GoFieldType:        "int32",
			JSONFieldName:      "deliver_time",
			ProtobufFieldName:  "deliver_time",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "express_no",
			Comment:            `快递单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExpressNo",
			GoFieldType:        "int32",
			JSONFieldName:      "express_no",
			ProtobufFieldName:  "express_no",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "express_code",
			Comment:            `快递公司 code`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExpressCode",
			GoFieldType:        "int32",
			JSONFieldName:      "express_code",
			ProtobufFieldName:  "express_code",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "mobile",
			Comment:            `手机号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "Mobile",
			GoFieldType:        "string",
			JSONFieldName:      "mobile",
			ProtobufFieldName:  "mobile",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "address",
			Comment:            `收货地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Address",
			GoFieldType:        "string",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "remark",
			Comment:            `买家留言`,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "add_time",
			Comment:            `订单拉取时间`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "last_time",
			Comment:            `订单更新时间`,
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
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderInfo) TableName() string {
	return "order_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderInfo) TableInfo() *TableInfo {
	return order_infoTableInfo
}
