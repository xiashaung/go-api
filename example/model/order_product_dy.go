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


CREATE TABLE `order_product_dy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '抖音子订单信息表',
  `order_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '主订单ID',
  `suborder_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '子订单ID',
  `shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `suborder_status` int(11) NOT NULL DEFAULT '0' COMMENT '子订单状态',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '佣金率',
  `estimated_commission` int(11) NOT NULL DEFAULT '0' COMMENT '预估佣金',
  `real_commission` int(11) NOT NULL DEFAULT '0' COMMENT '真实佣金',
  `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '在抖音小程序下单时，买家的抖音小程序ID',
  `biz` smallint(6) NOT NULL DEFAULT '0' COMMENT '业务来源：1-鲁班 2-小店 3-好好学习等',
  `order_type` int(11) NOT NULL DEFAULT '0' COMMENT '订单类型 (0实物，2普通虚拟，4poi核销，5三方核销，6服务市场)',
  `trade_type` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '交易类型：1-拼团 2-定金预售',
  `pre_sale_type` smallint(6) unsigned NOT NULL DEFAULT '0' COMMENT '预售类型 ，0 现货类型，1 全款预售 2 阶梯发货',
  `process_type` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '子流程类型：0-无流程, 1-退款子流程, 2-售后退货子流程, 3-整单退款, 4-玉石质检流程, 5-换货流程, 6-拼团流程, 7-风控子流程, 8-敏捷退款子流程, 9-冷静期, 10-清关证件照上传, 11-支付人一致性验证',
  `pay_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型 (0：货到付款，1：微信，2：支付宝）',
  `pay_time` int(11) NOT NULL DEFAULT '0' COMMENT '支付时间 (pay_type为0货到付款时, 此字段为空)',
  `finish_time` int(11) NOT NULL DEFAULT '0' COMMENT '订单完成时间\n',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '下单时间',
  `cancel_reason` varchar(255) NOT NULL DEFAULT '' COMMENT '取消原因',
  `send_pay` smallint(6) NOT NULL DEFAULT '0' COMMENT '流量来源：1-鲁班广告 2-联盟 3-商城 4-自主经营 5-线索通支付表单 6-抖音门店 7-抖+ 8-穿山甲',
  `talent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `author_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '直播主播id（达人)',
  `author_name` varchar(32) NOT NULL DEFAULT '' COMMENT '直播主播名称',
  `theme_type` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '下单来源：1-直播间 2-短视频 3-文章',
  `app_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '小程序id',
  `room_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '直播间id',
  `channel_payment_no` varchar(32) NOT NULL DEFAULT '' COMMENT '支付流水号',
  `order_amount` int(11) NOT NULL DEFAULT '0' COMMENT '订单金额（分）',
  `pay_amount` int(11) NOT NULL DEFAULT '0' COMMENT '支付金额（分）',
  `post_amount` int(11) NOT NULL DEFAULT '0' COMMENT '运费，单位分',
  `post_insurance_amount` int(11) NOT NULL DEFAULT '0' COMMENT '运费险金额\n',
  `ship_time` int(11) NOT NULL DEFAULT '0' COMMENT '发货时间',
  `logistics_receipt_time` int(11) NOT NULL DEFAULT '0' COMMENT '物流收货时间',
  `confirm_receipt_time` int(11) NOT NULL DEFAULT '0' COMMENT '用户确认收货时间',
  `promotion_amount` int(11) NOT NULL DEFAULT '0' COMMENT '单优惠总金额= 店铺优惠金额+ 平台优惠金额+ 达人优惠金额+ 支付优惠金额',
  `promotion_detail` json NOT NULL COMMENT '优惠信息',
  `logistics_info` json NOT NULL COMMENT '物流信息',
  `yx_product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'yx商品id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品ID',
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `product_pic` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头图，商品主图第一张',
  `sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品skuId\n',
  `spec` json NOT NULL COMMENT '规格信息',
  `category_list` json NOT NULL COMMENT '商品类目',
  `origin_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品现价',
  `item_num` int(11) NOT NULL DEFAULT '0' COMMENT '商品件数',
  `is_comment` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否评价 :1已评价，0未评价',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `oid_unique` (`order_id`) USING BTREE,
  KEY `talentid` (`talent_id`,`order_id`) USING BTREE,
  KEY `talentid_time` (`talent_id`,`finish_time`,`pay_time`) USING BTREE,
  KEY `idx_order_id` (`order_id`),
  KEY `idx_yx_product_id` (`yx_product_id`),
  KEY `idx_openid` (`open_id`,`pay_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='抖音子订单信息表'

JSON Sample
-------------------------------------
{    "id": 52,    "order_id": 16,    "suborder_id": 51,    "shop_id": 30,    "suborder_status": 0,    "commission_rate": 83,    "estimated_commission": 54,    "real_commission": 74,    "open_id": "ywoKnfnvurcdjrZqwaYgcGYcs",    "biz": 94,    "order_type": 95,    "trade_type": 22,    "pre_sale_type": 14,    "process_type": 76,    "pay_type": 87,    "pay_time": 16,    "finish_time": 43,    "create_time": 8,    "cancel_reason": "OryqOWuVhDRtiarIZuvLwbTLI",    "send_pay": 9,    "talent_id": 14,    "author_id": 50,    "author_name": "QYKEtcRfbLVqnyeAfsghcmLcK",    "theme_type": "ROBFxisbBHjHsDUYGjGqsfWKo",    "app_id": 42,    "room_id": 73,    "channel_payment_no": "DQBRYMSyoSbJPVETnfjNRfplD",    "order_amount": 77,    "pay_amount": 82,    "post_amount": 83,    "post_insurance_amount": 38,    "ship_time": 8,    "logistics_receipt_time": 89,    "confirm_receipt_time": 77,    "promotion_amount": 77,    "promotion_detail": "eFqopgRMIYrPpcYJpPuxmenYM",    "logistics_info": "tdBKslAekaovlyYkGFNvXJYOd",    "yx_product_id": 66,    "product_id": 66,    "product_name": "kpPveugClpecISQFUyOwYXqDv",    "product_pic": "aCqTgEnRVVBMOEiWlEgkhyLgP",    "sku_id": 9,    "spec": "EICONXXAvFTbaYBePwkLTUZSa",    "category_list": "YExhVWTjhvOjinYrYQjSSNaCh",    "origin_amount": 53,    "item_num": 9,    "is_comment": 96,    "add_time": 97,    "last_time": 90}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned
[20] column is set for unsigned
[37] column is set for unsigned
[38] column is set for unsigned
[41] column is set for unsigned
[44] column is set for unsigned
[47] column is set for unsigned
[48] column is set for unsigned



*/

// OrderProductDy struct is a row record of the order_product_dy table in the yx database
type OrderProductDy struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 抖音子订单信息表
	//[ 1] order_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	OrderID uint64 `gorm:"column:order_id;type:ubigint;default:0;" json:"order_id"` // 主订单ID
	//[ 2] suborder_id                                    ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SuborderID uint64 `gorm:"column:suborder_id;type:ubigint;default:0;" json:"suborder_id"` // 子订单ID
	//[ 3] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ShopID int64 `gorm:"column:shop_id;type:bigint;default:0;" json:"shop_id"` // 店铺ID
	//[ 4] suborder_status                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SuborderStatus int32 `gorm:"column:suborder_status;type:int;default:0;" json:"suborder_status"` // 子订单状态
	//[ 5] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 佣金率
	//[ 6] estimated_commission                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedCommission int32 `gorm:"column:estimated_commission;type:int;default:0;" json:"estimated_commission"` // 预估佣金
	//[ 7] real_commission                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RealCommission int32 `gorm:"column:real_commission;type:int;default:0;" json:"real_commission"` // 真实佣金
	//[ 8] open_id                                        varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	OpenID string `gorm:"column:open_id;type:varchar;size:64;" json:"open_id"` // 在抖音小程序下单时，买家的抖音小程序ID
	//[ 9] biz                                            smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	Biz int32 `gorm:"column:biz;type:smallint;default:0;" json:"biz"` // 业务来源：1-鲁班 2-小店 3-好好学习等
	//[10] order_type                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OrderType int32 `gorm:"column:order_type;type:int;default:0;" json:"order_type"` // 订单类型 (0实物，2普通虚拟，4poi核销，5三方核销，6服务市场)
	//[11] trade_type                                     usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	TradeType uint32 `gorm:"column:trade_type;type:usmallint;default:0;" json:"trade_type"` // 交易类型：1-拼团 2-定金预售
	//[12] pre_sale_type                                  usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	PreSaleType uint32 `gorm:"column:pre_sale_type;type:usmallint;default:0;" json:"pre_sale_type"` // 预售类型 ，0 现货类型，1 全款预售 2 阶梯发货
	//[13] process_type                                   usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	ProcessType uint32 `gorm:"column:process_type;type:usmallint;default:0;" json:"process_type"` // 子流程类型：0-无流程, 1-退款子流程, 2-售后退货子流程, 3-整单退款, 4-玉石质检流程, 5-换货流程, 6-拼团流程, 7-风控子流程, 8-敏捷退款子流程, 9-冷静期, 10-清关证件照上传, 11-支付人一致性验证
	//[14] pay_type                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayType uint32 `gorm:"column:pay_type;type:uint;default:0;" json:"pay_type"` // 支付类型 (0：货到付款，1：微信，2：支付宝）
	//[15] pay_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PayTime int32 `gorm:"column:pay_time;type:int;default:0;" json:"pay_time"` // 支付时间 (pay_type为0货到付款时, 此字段为空)
	//[16] finish_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FinishTime int32 `gorm:"column:finish_time;type:int;default:0;" json:"finish_time"` // 订单完成时间\n
	//[17] create_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CreateTime int32 `gorm:"column:create_time;type:int;default:0;" json:"create_time"` // 下单时间
	//[18] cancel_reason                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CancelReason string `gorm:"column:cancel_reason;type:varchar;size:255;" json:"cancel_reason"` // 取消原因
	//[19] send_pay                                       smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	SendPay int32 `gorm:"column:send_pay;type:smallint;default:0;" json:"send_pay"` // 流量来源：1-鲁班广告 2-联盟 3-商城 4-自主经营 5-线索通支付表单 6-抖音门店 7-抖+ 8-穿山甲
	//[20] talent_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TalentID uint64 `gorm:"column:talent_id;type:ubigint;default:0;" json:"talent_id"` // 达人id
	//[21] author_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AuthorID int64 `gorm:"column:author_id;type:bigint;default:0;" json:"author_id"` // 直播主播id（达人)
	//[22] author_name                                    varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	AuthorName string `gorm:"column:author_name;type:varchar;size:32;" json:"author_name"` // 直播主播名称
	//[23] theme_type                                     varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	ThemeType string `gorm:"column:theme_type;type:varchar;size:10;" json:"theme_type"` // 下单来源：1-直播间 2-短视频 3-文章
	//[24] app_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AppID int64 `gorm:"column:app_id;type:bigint;default:0;" json:"app_id"` // 小程序id
	//[25] room_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	RoomID int64 `gorm:"column:room_id;type:bigint;default:0;" json:"room_id"` // 直播间id
	//[26] channel_payment_no                             varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	ChannelPaymentNo string `gorm:"column:channel_payment_no;type:varchar;size:32;" json:"channel_payment_no"` // 支付流水号
	//[27] order_amount                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OrderAmount int32 `gorm:"column:order_amount;type:int;default:0;" json:"order_amount"` // 订单金额（分）
	//[28] pay_amount                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PayAmount int32 `gorm:"column:pay_amount;type:int;default:0;" json:"pay_amount"` // 支付金额（分）
	//[29] post_amount                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PostAmount int32 `gorm:"column:post_amount;type:int;default:0;" json:"post_amount"` // 运费，单位分
	//[30] post_insurance_amount                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PostInsuranceAmount int32 `gorm:"column:post_insurance_amount;type:int;default:0;" json:"post_insurance_amount"` // 运费险金额\n
	//[31] ship_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShipTime int32 `gorm:"column:ship_time;type:int;default:0;" json:"ship_time"` // 发货时间
	//[32] logistics_receipt_time                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LogisticsReceiptTime int32 `gorm:"column:logistics_receipt_time;type:int;default:0;" json:"logistics_receipt_time"` // 物流收货时间
	//[33] confirm_receipt_time                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ConfirmReceiptTime int32 `gorm:"column:confirm_receipt_time;type:int;default:0;" json:"confirm_receipt_time"` // 用户确认收货时间
	//[34] promotion_amount                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PromotionAmount int32 `gorm:"column:promotion_amount;type:int;default:0;" json:"promotion_amount"` // 单优惠总金额= 店铺优惠金额+ 平台优惠金额+ 达人优惠金额+ 支付优惠金额
	//[35] promotion_detail                               json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	PromotionDetail string `gorm:"column:promotion_detail;type:json;" json:"promotion_detail"` // 优惠信息
	//[36] logistics_info                                 json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	LogisticsInfo string `gorm:"column:logistics_info;type:json;" json:"logistics_info"` // 物流信息
	//[37] yx_product_id                                  ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	YxProductID uint64 `gorm:"column:yx_product_id;type:ubigint;default:0;" json:"yx_product_id"` // yx商品id
	//[38] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品ID
	//[39] product_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:255;" json:"product_name"` // 商品名称
	//[40] product_pic                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductPic string `gorm:"column:product_pic;type:varchar;size:255;" json:"product_pic"` // 头图，商品主图第一张
	//[41] sku_id                                         ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SkuID uint64 `gorm:"column:sku_id;type:ubigint;default:0;" json:"sku_id"` // 商品skuId\n
	//[42] spec                                           json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	Spec string `gorm:"column:spec;type:json;" json:"spec"` // 规格信息
	//[43] category_list                                  json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	CategoryList string `gorm:"column:category_list;type:json;" json:"category_list"` // 商品类目
	//[44] origin_amount                                  uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	OriginAmount uint32 `gorm:"column:origin_amount;type:uint;default:0;" json:"origin_amount"` // 商品现价
	//[45] item_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemNum int32 `gorm:"column:item_num;type:int;default:0;" json:"item_num"` // 商品件数
	//[46] is_comment                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsComment int32 `gorm:"column:is_comment;type:smallint;default:0;" json:"is_comment"` // 是否评价 :1已评价，0未评价
	//[47] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[48] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 修改时间

}

var order_product_dyTableInfo = &TableInfo{
	Name: "order_product_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `抖音子订单信息表`,
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
			Name:               "order_id",
			Comment:            `主订单ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "OrderID",
			GoFieldType:        "uint64",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "suborder_id",
			Comment:            `子订单ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SuborderID",
			GoFieldType:        "uint64",
			JSONFieldName:      "suborder_id",
			ProtobufFieldName:  "suborder_id",
			ProtobufType:       "uint64",
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
			Name:               "suborder_status",
			Comment:            `子订单状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SuborderStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "suborder_status",
			ProtobufFieldName:  "suborder_status",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "estimated_commission",
			Comment:            `预估佣金`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "estimated_commission",
			ProtobufFieldName:  "estimated_commission",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "real_commission",
			Comment:            `真实佣金`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RealCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "real_commission",
			ProtobufFieldName:  "real_commission",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "open_id",
			Comment:            `在抖音小程序下单时，买家的抖音小程序ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "OpenID",
			GoFieldType:        "string",
			JSONFieldName:      "open_id",
			ProtobufFieldName:  "open_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "biz",
			Comment:            `业务来源：1-鲁班 2-小店 3-好好学习等`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Biz",
			GoFieldType:        "int32",
			JSONFieldName:      "biz",
			ProtobufFieldName:  "biz",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "order_type",
			Comment:            `订单类型 (0实物，2普通虚拟，4poi核销，5三方核销，6服务市场)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrderType",
			GoFieldType:        "int32",
			JSONFieldName:      "order_type",
			ProtobufFieldName:  "order_type",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "trade_type",
			Comment:            `交易类型：1-拼团 2-定金预售`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "TradeType",
			GoFieldType:        "uint32",
			JSONFieldName:      "trade_type",
			ProtobufFieldName:  "trade_type",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "pre_sale_type",
			Comment:            `预售类型 ，0 现货类型，1 全款预售 2 阶梯发货`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "PreSaleType",
			GoFieldType:        "uint32",
			JSONFieldName:      "pre_sale_type",
			ProtobufFieldName:  "pre_sale_type",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "process_type",
			Comment:            `子流程类型：0-无流程, 1-退款子流程, 2-售后退货子流程, 3-整单退款, 4-玉石质检流程, 5-换货流程, 6-拼团流程, 7-风控子流程, 8-敏捷退款子流程, 9-冷静期, 10-清关证件照上传, 11-支付人一致性验证`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "ProcessType",
			GoFieldType:        "uint32",
			JSONFieldName:      "process_type",
			ProtobufFieldName:  "process_type",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "pay_type",
			Comment:            `支付类型 (0：货到付款，1：微信，2：支付宝）`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PayType",
			GoFieldType:        "uint32",
			JSONFieldName:      "pay_type",
			ProtobufFieldName:  "pay_type",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "pay_time",
			Comment:            `支付时间 (pay_type为0货到付款时, 此字段为空)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PayTime",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_time",
			ProtobufFieldName:  "pay_time",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "finish_time",
			Comment:            `订单完成时间\n`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FinishTime",
			GoFieldType:        "int32",
			JSONFieldName:      "finish_time",
			ProtobufFieldName:  "finish_time",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "create_time",
			Comment:            `下单时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CreateTime",
			GoFieldType:        "int32",
			JSONFieldName:      "create_time",
			ProtobufFieldName:  "create_time",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "cancel_reason",
			Comment:            `取消原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "CancelReason",
			GoFieldType:        "string",
			JSONFieldName:      "cancel_reason",
			ProtobufFieldName:  "cancel_reason",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "send_pay",
			Comment:            `流量来源：1-鲁班广告 2-联盟 3-商城 4-自主经营 5-线索通支付表单 6-抖音门店 7-抖+ 8-穿山甲`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "SendPay",
			GoFieldType:        "int32",
			JSONFieldName:      "send_pay",
			ProtobufFieldName:  "send_pay",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "author_id",
			Comment:            `直播主播id（达人)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AuthorID",
			GoFieldType:        "int64",
			JSONFieldName:      "author_id",
			ProtobufFieldName:  "author_id",
			ProtobufType:       "int64",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "author_name",
			Comment:            `直播主播名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "AuthorName",
			GoFieldType:        "string",
			JSONFieldName:      "author_name",
			ProtobufFieldName:  "author_name",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "theme_type",
			Comment:            `下单来源：1-直播间 2-短视频 3-文章`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "ThemeType",
			GoFieldType:        "string",
			JSONFieldName:      "theme_type",
			ProtobufFieldName:  "theme_type",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "app_id",
			Comment:            `小程序id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AppID",
			GoFieldType:        "int64",
			JSONFieldName:      "app_id",
			ProtobufFieldName:  "app_id",
			ProtobufType:       "int64",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "room_id",
			Comment:            `直播间id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RoomID",
			GoFieldType:        "int64",
			JSONFieldName:      "room_id",
			ProtobufFieldName:  "room_id",
			ProtobufType:       "int64",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "channel_payment_no",
			Comment:            `支付流水号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "ChannelPaymentNo",
			GoFieldType:        "string",
			JSONFieldName:      "channel_payment_no",
			ProtobufFieldName:  "channel_payment_no",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "order_amount",
			Comment:            `订单金额（分）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrderAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "order_amount",
			ProtobufFieldName:  "order_amount",
			ProtobufType:       "int32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "pay_amount",
			Comment:            `支付金额（分）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PayAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_amount",
			ProtobufFieldName:  "pay_amount",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "post_amount",
			Comment:            `运费，单位分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PostAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "post_amount",
			ProtobufFieldName:  "post_amount",
			ProtobufType:       "int32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "post_insurance_amount",
			Comment:            `运费险金额\n`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PostInsuranceAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "post_insurance_amount",
			ProtobufFieldName:  "post_insurance_amount",
			ProtobufType:       "int32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "ship_time",
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
			GoFieldName:        "ShipTime",
			GoFieldType:        "int32",
			JSONFieldName:      "ship_time",
			ProtobufFieldName:  "ship_time",
			ProtobufType:       "int32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "logistics_receipt_time",
			Comment:            `物流收货时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LogisticsReceiptTime",
			GoFieldType:        "int32",
			JSONFieldName:      "logistics_receipt_time",
			ProtobufFieldName:  "logistics_receipt_time",
			ProtobufType:       "int32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "confirm_receipt_time",
			Comment:            `用户确认收货时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ConfirmReceiptTime",
			GoFieldType:        "int32",
			JSONFieldName:      "confirm_receipt_time",
			ProtobufFieldName:  "confirm_receipt_time",
			ProtobufType:       "int32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "promotion_amount",
			Comment:            `单优惠总金额= 店铺优惠金额+ 平台优惠金额+ 达人优惠金额+ 支付优惠金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PromotionAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "promotion_amount",
			ProtobufFieldName:  "promotion_amount",
			ProtobufType:       "int32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "promotion_detail",
			Comment:            `优惠信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "PromotionDetail",
			GoFieldType:        "string",
			JSONFieldName:      "promotion_detail",
			ProtobufFieldName:  "promotion_detail",
			ProtobufType:       "string",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "logistics_info",
			Comment:            `物流信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "LogisticsInfo",
			GoFieldType:        "string",
			JSONFieldName:      "logistics_info",
			ProtobufFieldName:  "logistics_info",
			ProtobufType:       "string",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "yx_product_id",
			Comment:            `yx商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "YxProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "yx_product_id",
			ProtobufFieldName:  "yx_product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "product_id",
			Comment:            `商品ID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "product_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ProductName",
			GoFieldType:        "string",
			JSONFieldName:      "product_name",
			ProtobufFieldName:  "product_name",
			ProtobufType:       "string",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "product_pic",
			Comment:            `头图，商品主图第一张`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ProductPic",
			GoFieldType:        "string",
			JSONFieldName:      "product_pic",
			ProtobufFieldName:  "product_pic",
			ProtobufType:       "string",
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "sku_id",
			Comment:            `商品skuId\n`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SkuID",
			GoFieldType:        "uint64",
			JSONFieldName:      "sku_id",
			ProtobufFieldName:  "sku_id",
			ProtobufType:       "uint64",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "spec",
			Comment:            `规格信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "Spec",
			GoFieldType:        "string",
			JSONFieldName:      "spec",
			ProtobufFieldName:  "spec",
			ProtobufType:       "string",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
			Name:               "category_list",
			Comment:            `商品类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "CategoryList",
			GoFieldType:        "string",
			JSONFieldName:      "category_list",
			ProtobufFieldName:  "category_list",
			ProtobufType:       "string",
			ProtobufPos:        44,
		},

		&ColumnInfo{
			Index:              44,
			Name:               "origin_amount",
			Comment:            `商品现价`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "OriginAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "origin_amount",
			ProtobufFieldName:  "origin_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        45,
		},

		&ColumnInfo{
			Index:              45,
			Name:               "item_num",
			Comment:            `商品件数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ItemNum",
			GoFieldType:        "int32",
			JSONFieldName:      "item_num",
			ProtobufFieldName:  "item_num",
			ProtobufType:       "int32",
			ProtobufPos:        46,
		},

		&ColumnInfo{
			Index:              46,
			Name:               "is_comment",
			Comment:            `是否评价 :1已评价，0未评价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsComment",
			GoFieldType:        "int32",
			JSONFieldName:      "is_comment",
			ProtobufFieldName:  "is_comment",
			ProtobufType:       "int32",
			ProtobufPos:        47,
		},

		&ColumnInfo{
			Index:              47,
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
			ProtobufPos:        48,
		},

		&ColumnInfo{
			Index:              48,
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
			ProtobufPos:        49,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderProductDy) TableName() string {
	return "order_product_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderProductDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderProductDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderProductDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderProductDy) TableInfo() *TableInfo {
	return order_product_dyTableInfo
}
