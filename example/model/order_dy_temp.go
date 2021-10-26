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


CREATE TABLE `order_dy_temp` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '抖音订单信息表',
  `order_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主订单ID',
  `shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺id',
  `anchor_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播id',
  `anchor_open_id` varchar(64) NOT NULL DEFAULT '0' COMMENT '主播openid',
  `talent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `seller_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '央选卖家id',
  `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '在抖音小程序下单时，买家的抖音小程序ID',
  `order_status` smallint(6) NOT NULL DEFAULT '0' COMMENT '订单状态',
  `order_type` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '订单类型',
  `trade_type` smallint(6) NOT NULL DEFAULT '0' COMMENT '交易类型：1-拼团 2-定金预售',
  `main_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '主流程状态\n',
  `pay_time` int(11) NOT NULL DEFAULT '0' COMMENT '支付时间 (pay_type为0货到付款时, 此字段为空)',
  `finish_time` int(11) NOT NULL DEFAULT '0' COMMENT '订单完成时间',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '下单时间',
  `cancel_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '订单取消原因',
  `buyer_words` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '买家备注',
  `seller_words` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '卖家备注',
  `b_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '下单端：0-站外 1-火山 2-抖音等',
  `sub_b_type` int(11) NOT NULL DEFAULT '0' COMMENT '下单场景：0 未知 1 app 2 小程序 3 H5',
  `app_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '小程序id',
  `pay_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '支付类型 (0：货到付款，1：微信，2：支付宝）',
  `order_amount` int(11) NOT NULL DEFAULT '0' COMMENT '订单金额（分）',
  `pay_amount` int(11) NOT NULL DEFAULT '0' COMMENT '支付金额（分）',
  `post_amount` int(11) DEFAULT '0' COMMENT '快递费（分）',
  `post_insurance_amount` int(11) NOT NULL DEFAULT '0' COMMENT '运费险金额',
  `post_addr` json DEFAULT NULL COMMENT '收件人地址',
  `post_tel` varchar(18) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收件人电话',
  `post_receiver` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收件人姓名',
  `ship_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '发货时间。未发货时为"0"，已发货返回秒级时间戳',
  `exp_ship_time` int(11) NOT NULL DEFAULT '0' COMMENT '预计发货时间',
  `promotion_detail` json DEFAULT NULL COMMENT '优惠信息',
  `promotion_platform_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '平台优惠金额',
  `promotion_shop_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '店铺优惠金额',
  `shop_cost_amount` int(11) NOT NULL DEFAULT '0' COMMENT '平台优惠金额卖家承担部分',
  `platform_cost_amount` int(11) NOT NULL DEFAULT '0' COMMENT '平台优惠金额平台承担部分',
  `promotion_talent_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人优惠金额',
  `promotion_pay_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '支付优惠金额',
  `estimated_commission` int(11) NOT NULL DEFAULT '0' COMMENT '预估佣金\n',
  `seller_remark_stars` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '卖家订单标记 小旗子star取值0～5，分别表示 灰紫青绿橙红',
  `is_colonel` int(10) NOT NULL DEFAULT '0' COMMENT '是否团长订单 0 否 1是',
  `order_phase_list` json DEFAULT NULL COMMENT '定金预售阶段单',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `talentid_oid` (`order_id`,`talent_id`) USING BTREE,
  KEY `talentid_time` (`talent_id`,`finish_time`,`pay_time`) USING BTREE,
  KEY `idx_anchor_id` (`anchor_id`) USING BTREE,
  KEY `idx_openid_paytime` (`open_id`,`pay_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=248535 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='抖音订单信息表'

JSON Sample
-------------------------------------
{    "id": 86,    "order_id": "uuUvKcKHlmyDRKgaZGyQPfocS",    "shop_id": 81,    "anchor_id": 8,    "anchor_open_id": "jpaHfMOYaJpYjZFSRYuiVhJJh",    "talent_id": 58,    "seller_id": 84,    "open_id": "MMAPFMpdpGMGKCxwyVUWIyKAR",    "order_status": 63,    "order_type": 72,    "trade_type": 85,    "main_status": 52,    "pay_time": 0,    "finish_time": 0,    "create_time": 18,    "cancel_reason": "yNYMrJYduEyBMrCZLrlKHNQYM",    "buyer_words": "RLmmpFlqwMVeeXDoahqhYdegs",    "seller_words": "hRCUvnrcoNrxsyLDUGLvilhRe",    "b_type": 36,    "sub_b_type": 91,    "app_id": 28,    "pay_type": 56,    "order_amount": 32,    "pay_amount": 3,    "post_amount": 35,    "post_insurance_amount": 80,    "post_addr": "laLexAXFmRSPydoRtiMVvnCMe",    "post_tel": "obAjtNNySBAbDmexpIABHqBkR",    "post_receiver": "juDCFNaqPdcNgCmVPDLJsThMW",    "ship_time": 23,    "exp_ship_time": 30,    "promotion_detail": "UIYajxWXmiNddINqjTmvcJddp",    "promotion_platform_amount": 27,    "promotion_shop_amount": 88,    "shop_cost_amount": 80,    "platform_cost_amount": 44,    "promotion_talent_amount": 13,    "promotion_pay_amount": 74,    "estimated_commission": 87,    "seller_remark_stars": 5,    "is_colonel": 79,    "order_phase_list": "rEeyNxQCZeitsicJLyurTTjkL",    "add_time": 44,    "last_time": 99}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 9] column is set for unsigned
[11] column is set for unsigned
[18] column is set for unsigned
[21] column is set for unsigned
[29] column is set for unsigned
[32] column is set for unsigned
[33] column is set for unsigned
[36] column is set for unsigned
[37] column is set for unsigned
[39] column is set for unsigned
[42] column is set for unsigned
[43] column is set for unsigned



*/

// OrderDyTemp struct is a row record of the order_dy_temp table in the yx database
type OrderDyTemp struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 抖音订单信息表
	//[ 1] order_id                                       varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	OrderID string `gorm:"column:order_id;type:varchar;size:30;" json:"order_id"` // 主订单ID
	//[ 2] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ShopID int64 `gorm:"column:shop_id;type:bigint;default:0;" json:"shop_id"` // 店铺id
	//[ 3] anchor_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AnchorID int64 `gorm:"column:anchor_id;type:bigint;default:0;" json:"anchor_id"` // 主播id
	//[ 4] anchor_open_id                                 varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: [0]
	AnchorOpenID string `gorm:"column:anchor_open_id;type:varchar;size:64;default:0;" json:"anchor_open_id"` // 主播openid
	//[ 5] talent_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TalentID uint64 `gorm:"column:talent_id;type:ubigint;default:0;" json:"talent_id"` // 达人id
	//[ 6] seller_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SellerID uint64 `gorm:"column:seller_id;type:ubigint;default:0;" json:"seller_id"` // 央选卖家id
	//[ 7] open_id                                        varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	OpenID string `gorm:"column:open_id;type:varchar;size:64;" json:"open_id"` // 在抖音小程序下单时，买家的抖音小程序ID
	//[ 8] order_status                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	OrderStatus int32 `gorm:"column:order_status;type:smallint;default:0;" json:"order_status"` // 订单状态
	//[ 9] order_type                                     usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	OrderType uint32 `gorm:"column:order_type;type:usmallint;default:0;" json:"order_type"` // 订单类型
	//[10] trade_type                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	TradeType int32 `gorm:"column:trade_type;type:smallint;default:0;" json:"trade_type"` // 交易类型：1-拼团 2-定金预售
	//[11] main_status                                    usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	MainStatus uint32 `gorm:"column:main_status;type:usmallint;default:0;" json:"main_status"` // 主流程状态\n
	//[12] pay_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PayTime int32 `gorm:"column:pay_time;type:int;default:0;" json:"pay_time"` // 支付时间 (pay_type为0货到付款时, 此字段为空)
	//[13] finish_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FinishTime int32 `gorm:"column:finish_time;type:int;default:0;" json:"finish_time"` // 订单完成时间
	//[14] create_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CreateTime int32 `gorm:"column:create_time;type:int;default:0;" json:"create_time"` // 下单时间
	//[15] cancel_reason                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CancelReason string `gorm:"column:cancel_reason;type:varchar;size:255;" json:"cancel_reason"` // 订单取消原因
	//[16] buyer_words                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BuyerWords string `gorm:"column:buyer_words;type:varchar;size:255;" json:"buyer_words"` // 买家备注
	//[17] seller_words                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SellerWords string `gorm:"column:seller_words;type:varchar;size:255;" json:"seller_words"` // 卖家备注
	//[18] b_type                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	BType uint32 `gorm:"column:b_type;type:uint;default:0;" json:"b_type"` // 下单端：0-站外 1-火山 2-抖音等
	//[19] sub_b_type                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SubBType int32 `gorm:"column:sub_b_type;type:int;default:0;" json:"sub_b_type"` // 下单场景：0 未知 1 app 2 小程序 3 H5
	//[20] app_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AppID int64 `gorm:"column:app_id;type:bigint;default:0;" json:"app_id"` // 小程序id
	//[21] pay_type                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayType uint32 `gorm:"column:pay_type;type:uint;default:0;" json:"pay_type"` // 支付类型 (0：货到付款，1：微信，2：支付宝）
	//[22] order_amount                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OrderAmount int32 `gorm:"column:order_amount;type:int;default:0;" json:"order_amount"` // 订单金额（分）
	//[23] pay_amount                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PayAmount int32 `gorm:"column:pay_amount;type:int;default:0;" json:"pay_amount"` // 支付金额（分）
	//[24] post_amount                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PostAmount sql.NullInt64 `gorm:"column:post_amount;type:int;default:0;" json:"post_amount"` // 快递费（分）
	//[25] post_insurance_amount                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PostInsuranceAmount int32 `gorm:"column:post_insurance_amount;type:int;default:0;" json:"post_insurance_amount"` // 运费险金额
	//[26] post_addr                                      json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	PostAddr sql.NullString `gorm:"column:post_addr;type:json;" json:"post_addr"` // 收件人地址
	//[27] post_tel                                       varchar(18)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 18      default: []
	PostTel string `gorm:"column:post_tel;type:varchar;size:18;" json:"post_tel"` // 收件人电话
	//[28] post_receiver                                  varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	PostReceiver string `gorm:"column:post_receiver;type:varchar;size:32;" json:"post_receiver"` // 收件人姓名
	//[29] ship_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ShipTime uint32 `gorm:"column:ship_time;type:uint;default:0;" json:"ship_time"` // 发货时间。未发货时为"0"，已发货返回秒级时间戳
	//[30] exp_ship_time                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExpShipTime int32 `gorm:"column:exp_ship_time;type:int;default:0;" json:"exp_ship_time"` // 预计发货时间
	//[31] promotion_detail                               json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	PromotionDetail sql.NullString `gorm:"column:promotion_detail;type:json;" json:"promotion_detail"` // 优惠信息
	//[32] promotion_platform_amount                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PromotionPlatformAmount uint32 `gorm:"column:promotion_platform_amount;type:uint;default:0;" json:"promotion_platform_amount"` // 平台优惠金额
	//[33] promotion_shop_amount                          uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PromotionShopAmount uint32 `gorm:"column:promotion_shop_amount;type:uint;default:0;" json:"promotion_shop_amount"` // 店铺优惠金额
	//[34] shop_cost_amount                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShopCostAmount int32 `gorm:"column:shop_cost_amount;type:int;default:0;" json:"shop_cost_amount"` // 平台优惠金额卖家承担部分
	//[35] platform_cost_amount                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformCostAmount int32 `gorm:"column:platform_cost_amount;type:int;default:0;" json:"platform_cost_amount"` // 平台优惠金额平台承担部分
	//[36] promotion_talent_amount                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PromotionTalentAmount uint32 `gorm:"column:promotion_talent_amount;type:uint;default:0;" json:"promotion_talent_amount"` // 达人优惠金额
	//[37] promotion_pay_amount                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PromotionPayAmount uint32 `gorm:"column:promotion_pay_amount;type:uint;default:0;" json:"promotion_pay_amount"` // 支付优惠金额
	//[38] estimated_commission                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedCommission int32 `gorm:"column:estimated_commission;type:int;default:0;" json:"estimated_commission"` // 预估佣金\n
	//[39] seller_remark_stars                            usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	SellerRemarkStars uint32 `gorm:"column:seller_remark_stars;type:usmallint;default:0;" json:"seller_remark_stars"` // 卖家订单标记 小旗子star取值0～5，分别表示 灰紫青绿橙红
	//[40] is_colonel                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsColonel int32 `gorm:"column:is_colonel;type:int;default:0;" json:"is_colonel"` // 是否团长订单 0 否 1是
	//[41] order_phase_list                               json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	OrderPhaseList sql.NullString `gorm:"column:order_phase_list;type:json;" json:"order_phase_list"` // 定金预售阶段单
	//[42] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[43] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 修改时间

}

var order_dy_tempTableInfo = &TableInfo{
	Name: "order_dy_temp",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `抖音订单信息表`,
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
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "OrderID",
			GoFieldType:        "string",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "shop_id",
			Comment:            `店铺id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "anchor_id",
			Comment:            `主播id`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "anchor_open_id",
			Comment:            `主播openid`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "AnchorOpenID",
			GoFieldType:        "string",
			JSONFieldName:      "anchor_open_id",
			ProtobufFieldName:  "anchor_open_id",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "seller_id",
			Comment:            `央选卖家id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SellerID",
			GoFieldType:        "uint64",
			JSONFieldName:      "seller_id",
			ProtobufFieldName:  "seller_id",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "order_status",
			Comment:            `订单状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "OrderStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "order_status",
			ProtobufFieldName:  "order_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "order_type",
			Comment:            `订单类型`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "OrderType",
			GoFieldType:        "uint32",
			JSONFieldName:      "order_type",
			ProtobufFieldName:  "order_type",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "trade_type",
			Comment:            `交易类型：1-拼团 2-定金预售`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "TradeType",
			GoFieldType:        "int32",
			JSONFieldName:      "trade_type",
			ProtobufFieldName:  "trade_type",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "main_status",
			Comment:            `主流程状态\n`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "MainStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "main_status",
			ProtobufFieldName:  "main_status",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "finish_time",
			Comment:            `订单完成时间`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "cancel_reason",
			Comment:            `订单取消原因`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "buyer_words",
			Comment:            `买家备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "BuyerWords",
			GoFieldType:        "string",
			JSONFieldName:      "buyer_words",
			ProtobufFieldName:  "buyer_words",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "seller_words",
			Comment:            `卖家备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SellerWords",
			GoFieldType:        "string",
			JSONFieldName:      "seller_words",
			ProtobufFieldName:  "seller_words",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "b_type",
			Comment:            `下单端：0-站外 1-火山 2-抖音等`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "BType",
			GoFieldType:        "uint32",
			JSONFieldName:      "b_type",
			ProtobufFieldName:  "b_type",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "sub_b_type",
			Comment:            `下单场景：0 未知 1 app 2 小程序 3 H5`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SubBType",
			GoFieldType:        "int32",
			JSONFieldName:      "sub_b_type",
			ProtobufFieldName:  "sub_b_type",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "post_amount",
			Comment:            `快递费（分）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PostAmount",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "post_amount",
			ProtobufFieldName:  "post_amount",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "post_insurance_amount",
			Comment:            `运费险金额`,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "post_addr",
			Comment:            `收件人地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "PostAddr",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "post_addr",
			ProtobufFieldName:  "post_addr",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "post_tel",
			Comment:            `收件人电话`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(18)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       18,
			GoFieldName:        "PostTel",
			GoFieldType:        "string",
			JSONFieldName:      "post_tel",
			ProtobufFieldName:  "post_tel",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "post_receiver",
			Comment:            `收件人姓名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "PostReceiver",
			GoFieldType:        "string",
			JSONFieldName:      "post_receiver",
			ProtobufFieldName:  "post_receiver",
			ProtobufType:       "string",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "ship_time",
			Comment:            `发货时间。未发货时为"0"，已发货返回秒级时间戳`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ShipTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "ship_time",
			ProtobufFieldName:  "ship_time",
			ProtobufType:       "uint32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "exp_ship_time",
			Comment:            `预计发货时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExpShipTime",
			GoFieldType:        "int32",
			JSONFieldName:      "exp_ship_time",
			ProtobufFieldName:  "exp_ship_time",
			ProtobufType:       "int32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "promotion_detail",
			Comment:            `优惠信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "PromotionDetail",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "promotion_detail",
			ProtobufFieldName:  "promotion_detail",
			ProtobufType:       "string",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "promotion_platform_amount",
			Comment:            `平台优惠金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PromotionPlatformAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "promotion_platform_amount",
			ProtobufFieldName:  "promotion_platform_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "promotion_shop_amount",
			Comment:            `店铺优惠金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PromotionShopAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "promotion_shop_amount",
			ProtobufFieldName:  "promotion_shop_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "shop_cost_amount",
			Comment:            `平台优惠金额卖家承担部分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ShopCostAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "shop_cost_amount",
			ProtobufFieldName:  "shop_cost_amount",
			ProtobufType:       "int32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "platform_cost_amount",
			Comment:            `平台优惠金额平台承担部分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformCostAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_cost_amount",
			ProtobufFieldName:  "platform_cost_amount",
			ProtobufType:       "int32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "promotion_talent_amount",
			Comment:            `达人优惠金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PromotionTalentAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "promotion_talent_amount",
			ProtobufFieldName:  "promotion_talent_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "promotion_pay_amount",
			Comment:            `支付优惠金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PromotionPayAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "promotion_pay_amount",
			ProtobufFieldName:  "promotion_pay_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "estimated_commission",
			Comment:            `预估佣金\n`,
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
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "seller_remark_stars",
			Comment:            `卖家订单标记 小旗子star取值0～5，分别表示 灰紫青绿橙红`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "SellerRemarkStars",
			GoFieldType:        "uint32",
			JSONFieldName:      "seller_remark_stars",
			ProtobufFieldName:  "seller_remark_stars",
			ProtobufType:       "uint32",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "is_colonel",
			Comment:            `是否团长订单 0 否 1是`,
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
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "order_phase_list",
			Comment:            `定金预售阶段单`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "OrderPhaseList",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "order_phase_list",
			ProtobufFieldName:  "order_phase_list",
			ProtobufType:       "string",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
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
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
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
			ProtobufPos:        44,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderDyTemp) TableName() string {
	return "order_dy_temp"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderDyTemp) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderDyTemp) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderDyTemp) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderDyTemp) TableInfo() *TableInfo {
	return order_dy_tempTableInfo
}