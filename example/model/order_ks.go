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


CREATE TABLE `order_ks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `oid` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单号',
  `pay_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '付款时间',
  `buyer_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '买家id',
  `seller_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '卖家id',
  `merchant_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '商家id',
  `shop_id` int(11) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `talent_id` int(11) NOT NULL DEFAULT '0' COMMENT '达人ID',
  `anchor_id` int(11) NOT NULL DEFAULT '0' COMMENT '主播ID',
  `distributor_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '推广者ID',
  `seller_nick` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '卖家昵称',
  `express_fee` int(11) NOT NULL DEFAULT '0' COMMENT '运费（单位分）',
  `discount_fee` int(11) NOT NULL DEFAULT '0' COMMENT '折扣价格（单位分）',
  `total_fee` int(11) NOT NULL DEFAULT '0' COMMENT '子订单商品总价（单位分）',
  `platform_allowance` int(11) NOT NULL DEFAULT '0' COMMENT '优惠补贴，platformBearAmount + merchantBearAmount',
  `seller_recv` int(11) NOT NULL DEFAULT '0' COMMENT '商家实收，货款 + 运费',
  `buyer_pay` int(11) NOT NULL DEFAULT '0' COMMENT '买家实付，货款 + 运费 - 平台承担补贴（platformBearAmount）',
  `order_status` int(11) NOT NULL DEFAULT '0' COMMENT '订单状态：[0, "未知状态"], [10, "待付款"], [30, "已付款"], [40, "已发货"], [50, "已签收"], [70, "订单成功"], [80, "订单失败"];',
  `send_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '发货时间',
  `refund_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '退款时间(只要买家有申请退款行为)',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT 'sku商品种类数',
  `mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收件人手机号（此字段与address 中mobile 意义相同，修改address 中的mobile，此字段值也会被修改）',
  `address` json DEFAULT NULL COMMENT '{"consignee": "不帅", //收货人姓名 "mobile": 15000000000, //收货人手机 "provinceCode": 125620, //收货人省份编号 "province": "北京", //收货人省份 "cityCode": 100000, //收货人城市编码 "city": "北京", //收货人城市 "districtCode": 123456, //收货人区\\县编码 "district": "", //收货人区\\县 "address": "" //收货人详细地址 }',
  `refund` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否申请退款 0未申请 1已申请',
  `remark` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '买家留言',
  `seller_note_list` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '卖家备注;["请给好评","请给好评"]',
  `receive_time` int(11) NOT NULL DEFAULT '0' COMMENT '签收时间',
  `deliver_goods_time` smallint(11) NOT NULL DEFAULT '0' COMMENT '发货间隔时间，单位：天',
  `promise_deliver_time` bigint(11) NOT NULL DEFAULT '0' COMMENT '非预售商品承诺发货时间点,即订单支付时间+承诺发货时间间隔',
  `activity_type` int(11) NOT NULL DEFAULT '1' COMMENT '活动类型',
  `cps_type` int(11) NOT NULL DEFAULT '1' COMMENT '分销类型 0-全部 1-普通订单 2-分销订单',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '佣金率',
  `estimated_income` int(11) NOT NULL DEFAULT '0' COMMENT '分销总金额',
  `settlement_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结算时间',
  `settlement_success_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结算成功时间',
  `pay_type` int(11) NOT NULL DEFAULT '1' COMMENT '支付方式 1微信 2支付宝',
  `express_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递单号',
  `express_code` bigint(20) NOT NULL DEFAULT '0' COMMENT '快递公司code',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT 'db写入时间',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT 'db更新时间',
  `is_colonel` int(10) NOT NULL DEFAULT '0' COMMENT '是否为团长订单 0 否 1 是',
  `is_platform` int(10) NOT NULL DEFAULT '0' COMMENT '是否为平台订单 0 否 1是',
  PRIMARY KEY (`id`),
  KEY `idx_seller_ptime` (`seller_id`,`pay_time`) USING BTREE,
  KEY `idx_anchor_time` (`pay_time`) USING BTREE,
  KEY `idx_oid` (`oid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='快手订单表'

JSON Sample
-------------------------------------
{    "id": 45,    "oid": 97,    "pay_time": 88,    "buyer_id": 7,    "seller_id": 95,    "merchant_id": 26,    "shop_id": 40,    "talent_id": 12,    "anchor_id": 23,    "distributor_id": 47,    "seller_nick": "wKKxBCciOHhopogtnVxwpVILy",    "express_fee": 9,    "discount_fee": 5,    "total_fee": 37,    "platform_allowance": 92,    "seller_recv": 85,    "buyer_pay": 68,    "order_status": 78,    "send_time": 26,    "refund_time": 24,    "create_time": 90,    "update_time": 44,    "num": 96,    "mobile": "hRTNyxXnvwqermFFXMHIQPMpb",    "address": "KbtxhucVtTHCpxJxrGSCAFfCu",    "refund": 52,    "remark": "nZcfFgDAQqbuNGEHGuoZPBoSK",    "seller_note_list": "uOvqAMUIsqfhcBlJyAvtvepsr",    "receive_time": 54,    "deliver_goods_time": 74,    "promise_deliver_time": 86,    "activity_type": 31,    "cps_type": 55,    "commission_rate": 95,    "estimated_income": 6,    "settlement_time": 74,    "settlement_success_time": 96,    "pay_type": 20,    "express_no": "ZBUvyLXnoPesLTgZytOMLSnke",    "express_code": 78,    "add_time": 19,    "last_time": 14,    "is_colonel": 64,    "is_platform": 79}



*/

// OrderKs struct is a row record of the order_ks table in the yx database
type OrderKs struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 主键
	//[ 1] oid                                            bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	Oid int64 `gorm:"column:oid;type:bigint;default:0;" json:"oid"` // 订单号
	//[ 2] pay_time                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PayTime int64 `gorm:"column:pay_time;type:bigint;default:0;" json:"pay_time"` // 付款时间
	//[ 3] buyer_id                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	BuyerID int64 `gorm:"column:buyer_id;type:bigint;default:0;" json:"buyer_id"` // 买家id
	//[ 4] seller_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SellerID int64 `gorm:"column:seller_id;type:bigint;default:0;" json:"seller_id"` // 卖家id
	//[ 5] merchant_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	MerchantID int64 `gorm:"column:merchant_id;type:bigint;default:0;" json:"merchant_id"` // 商家id
	//[ 6] shop_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShopID int32 `gorm:"column:shop_id;type:int;default:0;" json:"shop_id"` // 店铺ID
	//[ 7] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"` // 达人ID
	//[ 8] anchor_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AnchorID int32 `gorm:"column:anchor_id;type:int;default:0;" json:"anchor_id"` // 主播ID
	//[ 9] distributor_id                                 bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	DistributorID int64 `gorm:"column:distributor_id;type:bigint;default:0;" json:"distributor_id"` // 推广者ID
	//[10] seller_nick                                    varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	SellerNick string `gorm:"column:seller_nick;type:varchar;size:125;" json:"seller_nick"` // 卖家昵称
	//[11] express_fee                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExpressFee int32 `gorm:"column:express_fee;type:int;default:0;" json:"express_fee"` // 运费（单位分）
	//[12] discount_fee                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DiscountFee int32 `gorm:"column:discount_fee;type:int;default:0;" json:"discount_fee"` // 折扣价格（单位分）
	//[13] total_fee                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TotalFee int32 `gorm:"column:total_fee;type:int;default:0;" json:"total_fee"` // 子订单商品总价（单位分）
	//[14] platform_allowance                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformAllowance int32 `gorm:"column:platform_allowance;type:int;default:0;" json:"platform_allowance"` // 优惠补贴，platformBearAmount + merchantBearAmount
	//[15] seller_recv                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SellerRecv int32 `gorm:"column:seller_recv;type:int;default:0;" json:"seller_recv"` // 商家实收，货款 + 运费
	//[16] buyer_pay                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	BuyerPay int32 `gorm:"column:buyer_pay;type:int;default:0;" json:"buyer_pay"` // 买家实付，货款 + 运费 - 平台承担补贴（platformBearAmount）
	//[17] order_status                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OrderStatus int32 `gorm:"column:order_status;type:int;default:0;" json:"order_status"` // 订单状态：[0, "未知状态"], [10, "待付款"], [30, "已付款"], [40, "已发货"], [50, "已签收"], [70, "订单成功"], [80, "订单失败"];
	//[18] send_time                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SendTime int64 `gorm:"column:send_time;type:bigint;default:0;" json:"send_time"` // 发货时间
	//[19] refund_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	RefundTime int64 `gorm:"column:refund_time;type:bigint;default:0;" json:"refund_time"` // 退款时间(只要买家有申请退款行为)
	//[20] create_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	CreateTime int64 `gorm:"column:create_time;type:bigint;default:0;" json:"create_time"` // 创建时间
	//[21] update_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	UpdateTime int64 `gorm:"column:update_time;type:bigint;default:0;" json:"update_time"` // 更新时间
	//[22] num                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Num int32 `gorm:"column:num;type:int;default:0;" json:"num"` // sku商品种类数
	//[23] mobile                                         char(11)             null: false  primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	Mobile string `gorm:"column:mobile;type:char;size:11;" json:"mobile"` // 收件人手机号（此字段与address 中mobile 意义相同，修改address 中的mobile，此字段值也会被修改）
	//[24] address                                        json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Address sql.NullString `gorm:"column:address;type:json;" json:"address"` // {"consignee": "不帅", //收货人姓名 "mobile": 15000000000, //收货人手机 "provinceCode": 125620, //收货人省份编号 "province": "北京", //收货人省份 "cityCode": 100000, //收货人城市编码 "city": "北京", //收货人城市 "districtCode": 123456, //收货人区\\县编码 "district": "", //收货人区\\县 "address": "" //收货人详细地址 }
	//[25] refund                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Refund int32 `gorm:"column:refund;type:smallint;default:0;" json:"refund"` // 是否申请退款 0未申请 1已申请
	//[26] remark                                         text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Remark string `gorm:"column:remark;type:text;size:65535;" json:"remark"` // 买家留言
	//[27] seller_note_list                               text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	SellerNoteList string `gorm:"column:seller_note_list;type:text;size:65535;" json:"seller_note_list"` // 卖家备注;["请给好评","请给好评"]
	//[28] receive_time                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ReceiveTime int32 `gorm:"column:receive_time;type:int;default:0;" json:"receive_time"` // 签收时间
	//[29] deliver_goods_time                             smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	DeliverGoodsTime int32 `gorm:"column:deliver_goods_time;type:smallint;default:0;" json:"deliver_goods_time"` // 发货间隔时间，单位：天
	//[30] promise_deliver_time                           bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PromiseDeliverTime int64 `gorm:"column:promise_deliver_time;type:bigint;default:0;" json:"promise_deliver_time"` // 非预售商品承诺发货时间点,即订单支付时间+承诺发货时间间隔
	//[31] activity_type                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	ActivityType int32 `gorm:"column:activity_type;type:int;default:1;" json:"activity_type"` // 活动类型
	//[32] cps_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	CpsType int32 `gorm:"column:cps_type;type:int;default:1;" json:"cps_type"` // 分销类型 0-全部 1-普通订单 2-分销订单
	//[33] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 佣金率
	//[34] estimated_income                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedIncome int32 `gorm:"column:estimated_income;type:int;default:0;" json:"estimated_income"` // 分销总金额
	//[35] settlement_time                                bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SettlementTime int64 `gorm:"column:settlement_time;type:bigint;default:0;" json:"settlement_time"` // 结算时间
	//[36] settlement_success_time                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SettlementSuccessTime int64 `gorm:"column:settlement_success_time;type:bigint;default:0;" json:"settlement_success_time"` // 结算成功时间
	//[37] pay_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	PayType int32 `gorm:"column:pay_type;type:int;default:1;" json:"pay_type"` // 支付方式 1微信 2支付宝
	//[38] express_no                                     varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	ExpressNo string `gorm:"column:express_no;type:varchar;size:50;" json:"express_no"` // 快递单号
	//[39] express_code                                   bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ExpressCode int64 `gorm:"column:express_code;type:bigint;default:0;" json:"express_code"` // 快递公司code
	//[40] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // db写入时间
	//[41] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // db更新时间
	//[42] is_colonel                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsColonel int32 `gorm:"column:is_colonel;type:int;default:0;" json:"is_colonel"` // 是否为团长订单 0 否 1 是
	//[43] is_platform                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsPlatform int32 `gorm:"column:is_platform;type:int;default:0;" json:"is_platform"` // 是否为平台订单 0 否 1是

}

var order_ksTableInfo = &TableInfo{
	Name: "order_ks",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
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
			Name:               "oid",
			Comment:            `订单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "Oid",
			GoFieldType:        "int64",
			JSONFieldName:      "oid",
			ProtobufFieldName:  "oid",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "pay_time",
			Comment:            `付款时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PayTime",
			GoFieldType:        "int64",
			JSONFieldName:      "pay_time",
			ProtobufFieldName:  "pay_time",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "buyer_id",
			Comment:            `买家id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "BuyerID",
			GoFieldType:        "int64",
			JSONFieldName:      "buyer_id",
			ProtobufFieldName:  "buyer_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "seller_id",
			Comment:            `卖家id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SellerID",
			GoFieldType:        "int64",
			JSONFieldName:      "seller_id",
			ProtobufFieldName:  "seller_id",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "merchant_id",
			Comment:            `商家id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int64",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "shop_id",
			Comment:            `店铺ID`,
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
			Name:               "talent_id",
			Comment:            `达人ID`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "anchor_id",
			Comment:            `主播ID`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "distributor_id",
			Comment:            `推广者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "DistributorID",
			GoFieldType:        "int64",
			JSONFieldName:      "distributor_id",
			ProtobufFieldName:  "distributor_id",
			ProtobufType:       "int64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "seller_nick",
			Comment:            `卖家昵称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "SellerNick",
			GoFieldType:        "string",
			JSONFieldName:      "seller_nick",
			ProtobufFieldName:  "seller_nick",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "express_fee",
			Comment:            `运费（单位分）`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "discount_fee",
			Comment:            `折扣价格（单位分）`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "total_fee",
			Comment:            `子订单商品总价（单位分）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TotalFee",
			GoFieldType:        "int32",
			JSONFieldName:      "total_fee",
			ProtobufFieldName:  "total_fee",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "platform_allowance",
			Comment:            `优惠补贴，platformBearAmount + merchantBearAmount`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformAllowance",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_allowance",
			ProtobufFieldName:  "platform_allowance",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "seller_recv",
			Comment:            `商家实收，货款 + 运费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SellerRecv",
			GoFieldType:        "int32",
			JSONFieldName:      "seller_recv",
			ProtobufFieldName:  "seller_recv",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "buyer_pay",
			Comment:            `买家实付，货款 + 运费 - 平台承担补贴（platformBearAmount）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BuyerPay",
			GoFieldType:        "int32",
			JSONFieldName:      "buyer_pay",
			ProtobufFieldName:  "buyer_pay",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "order_status",
			Comment:            `订单状态：[0, "未知状态"], [10, "待付款"], [30, "已付款"], [40, "已发货"], [50, "已签收"], [70, "订单成功"], [80, "订单失败"];`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrderStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "order_status",
			ProtobufFieldName:  "order_status",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "send_time",
			Comment:            `发货时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SendTime",
			GoFieldType:        "int64",
			JSONFieldName:      "send_time",
			ProtobufFieldName:  "send_time",
			ProtobufType:       "int64",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "refund_time",
			Comment:            `退款时间(只要买家有申请退款行为)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RefundTime",
			GoFieldType:        "int64",
			JSONFieldName:      "refund_time",
			ProtobufFieldName:  "refund_time",
			ProtobufType:       "int64",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "create_time",
			Comment:            `创建时间`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "update_time",
			Comment:            `更新时间`,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "num",
			Comment:            `sku商品种类数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Num",
			GoFieldType:        "int32",
			JSONFieldName:      "num",
			ProtobufFieldName:  "num",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "mobile",
			Comment:            `收件人手机号（此字段与address 中mobile 意义相同，修改address 中的mobile，此字段值也会被修改）`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "address",
			Comment:            `{"consignee": "不帅", //收货人姓名 "mobile": 15000000000, //收货人手机 "provinceCode": 125620, //收货人省份编号 "province": "北京", //收货人省份 "cityCode": 100000, //收货人城市编码 "city": "北京", //收货人城市 "districtCode": 123456, //收货人区\\县编码 "district": "", //收货人区\\县 "address": "" //收货人详细地址 }`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "Address",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "refund",
			Comment:            `是否申请退款 0未申请 1已申请`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Refund",
			GoFieldType:        "int32",
			JSONFieldName:      "refund",
			ProtobufFieldName:  "refund",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "remark",
			Comment:            `买家留言`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Remark",
			GoFieldType:        "string",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "seller_note_list",
			Comment:            `卖家备注;["请给好评","请给好评"]`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "SellerNoteList",
			GoFieldType:        "string",
			JSONFieldName:      "seller_note_list",
			ProtobufFieldName:  "seller_note_list",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "receive_time",
			Comment:            `签收时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ReceiveTime",
			GoFieldType:        "int32",
			JSONFieldName:      "receive_time",
			ProtobufFieldName:  "receive_time",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "deliver_goods_time",
			Comment:            `发货间隔时间，单位：天`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "DeliverGoodsTime",
			GoFieldType:        "int32",
			JSONFieldName:      "deliver_goods_time",
			ProtobufFieldName:  "deliver_goods_time",
			ProtobufType:       "int32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "promise_deliver_time",
			Comment:            `非预售商品承诺发货时间点,即订单支付时间+承诺发货时间间隔`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PromiseDeliverTime",
			GoFieldType:        "int64",
			JSONFieldName:      "promise_deliver_time",
			ProtobufFieldName:  "promise_deliver_time",
			ProtobufType:       "int64",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "activity_type",
			Comment:            `活动类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ActivityType",
			GoFieldType:        "int32",
			JSONFieldName:      "activity_type",
			ProtobufFieldName:  "activity_type",
			ProtobufType:       "int32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "cps_type",
			Comment:            `分销类型 0-全部 1-普通订单 2-分销订单`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CpsType",
			GoFieldType:        "int32",
			JSONFieldName:      "cps_type",
			ProtobufFieldName:  "cps_type",
			ProtobufType:       "int32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "commission_rate",
			Comment:            `佣金率`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommissionRate",
			GoFieldType:        "int32",
			JSONFieldName:      "commission_rate",
			ProtobufFieldName:  "commission_rate",
			ProtobufType:       "int32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "estimated_income",
			Comment:            `分销总金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedIncome",
			GoFieldType:        "int32",
			JSONFieldName:      "estimated_income",
			ProtobufFieldName:  "estimated_income",
			ProtobufType:       "int32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "settlement_time",
			Comment:            `结算时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SettlementTime",
			GoFieldType:        "int64",
			JSONFieldName:      "settlement_time",
			ProtobufFieldName:  "settlement_time",
			ProtobufType:       "int64",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "settlement_success_time",
			Comment:            `结算成功时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SettlementSuccessTime",
			GoFieldType:        "int64",
			JSONFieldName:      "settlement_success_time",
			ProtobufFieldName:  "settlement_success_time",
			ProtobufType:       "int64",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "pay_type",
			Comment:            `支付方式 1微信 2支付宝`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PayType",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_type",
			ProtobufFieldName:  "pay_type",
			ProtobufType:       "int32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "express_no",
			Comment:            `快递单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "ExpressNo",
			GoFieldType:        "string",
			JSONFieldName:      "express_no",
			ProtobufFieldName:  "express_no",
			ProtobufType:       "string",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "express_code",
			Comment:            `快递公司code`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressCode",
			GoFieldType:        "int64",
			JSONFieldName:      "express_code",
			ProtobufFieldName:  "express_code",
			ProtobufType:       "int64",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "add_time",
			Comment:            `db写入时间`,
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
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "last_time",
			Comment:            `db更新时间`,
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
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "is_colonel",
			Comment:            `是否为团长订单 0 否 1 是`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsColonel",
			GoFieldType:        "int32",
			JSONFieldName:      "is_colonel",
			ProtobufFieldName:  "is_colonel",
			ProtobufType:       "int32",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
			Name:               "is_platform",
			Comment:            `是否为平台订单 0 否 1是`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsPlatform",
			GoFieldType:        "int32",
			JSONFieldName:      "is_platform",
			ProtobufFieldName:  "is_platform",
			ProtobufType:       "int32",
			ProtobufPos:        44,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderKs) TableName() string {
	return "order_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderKs) TableInfo() *TableInfo {
	return order_ksTableInfo
}
