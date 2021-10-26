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


CREATE TABLE `order_product_ks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ks_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'ks订单商品id',
  `oid` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单id',
  `talent_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `anchor_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '主播id',
  `shop_id` int(11) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `merchant_id` int(11) NOT NULL DEFAULT '0' COMMENT '商家ID',
  `seller_id` int(11) NOT NULL DEFAULT '0' COMMENT '第三方商家ID',
  `order_status` smallint(6) NOT NULL DEFAULT '0' COMMENT '子订单状态',
  `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快手商品skuid',
  `rel_sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务商商品skuid',
  `sku_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'sku描述',
  `sku_nick` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'sku编码',
  `item_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快手商品id',
  `yx_product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'yx商品id',
  `rel_item_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '服务商商品id',
  `item_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `item_link_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品链接',
  `item_pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品图片地址',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '成交数量',
  `original_price` int(11) NOT NULL DEFAULT '0' COMMENT '商品促销前单价快照 以分为单位',
  `discount_fee` int(11) NOT NULL DEFAULT '0' COMMENT '折扣金额',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '商品单价快照，以分为单位',
  `pay_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付时间',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间-毫秒时间戳',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间-毫秒时间戳',
  `refund_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '退款订单id',
  `refund_status` int(11) NOT NULL DEFAULT '0' COMMENT '退款状态',
  `item_type` int(11) NOT NULL DEFAULT '0' COMMENT '1自建商品 2 闪电购商品',
  `refund_list` json DEFAULT NULL COMMENT '该skuId下的退款列表',
  `prev_info` json DEFAULT NULL COMMENT '订单修改之前的原订单信息',
  `distributor_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '推广者id',
  `distributor_name` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '推广者昵称',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '分销率',
  `estimated_income` int(11) NOT NULL DEFAULT '0' COMMENT '分销金额',
  `settlement_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结算时间',
  `settlement_success_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '结算成功时间',
  `platform_dp_rate` int(11) NOT NULL DEFAULT '0' COMMENT '平台分销推广佣金分成比例。千分比',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT 'db写入时间',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT 'db更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_anchor_id_pay_time` (`anchor_id`,`pay_time`) USING BTREE,
  KEY `oid_ksid_idx` (`oid`,`ks_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='快手子订单信息表'

JSON Sample
-------------------------------------
{    "id": 23,    "ks_id": 19,    "oid": 15,    "talent_id": 35,    "anchor_id": 59,    "shop_id": 72,    "merchant_id": 27,    "seller_id": 64,    "order_status": 70,    "sku_id": 84,    "rel_sku_id": 81,    "sku_desc": "ACHPpRbtnVCPDtuwGfDRatjOo",    "sku_nick": "eDUSgdZQIJCXdhOSGaEnKioZT",    "item_id": 13,    "yx_product_id": 6,    "rel_item_id": 38,    "item_title": "kfjMsaxDirCbjmkqUuAegTKaC",    "item_link_url": "wwPXimtIYamFynZBMUuJgCbiV",    "item_pic_url": "iobCXEAIcTyfRRKeAFWljcrgm",    "num": 61,    "original_price": 48,    "discount_fee": 10,    "price": 26,    "pay_time": 63,    "create_time": 49,    "update_time": 1,    "refund_id": 97,    "refund_status": 72,    "item_type": 92,    "refund_list": "TkyxnXoZlmIpBVhfsmpeYFqwD",    "prev_info": "PIKwKBfUQxAXIKfRHJKRbMcAd",    "distributor_id": 63,    "distributor_name": "VXBAgCfHyrqZHYpovxFrywwvF",    "commission_rate": 6,    "estimated_income": 22,    "settlement_time": 66,    "settlement_success_time": 91,    "platform_dp_rate": 45,    "add_time": 95,    "last_time": 50}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 4] column is set for unsigned



*/

// OrderProductKs struct is a row record of the order_product_ks table in the yx database
type OrderProductKs struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 主键
	//[ 1] ks_id                                          bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	KsID int64 `gorm:"column:ks_id;type:bigint;default:0;" json:"ks_id"` // ks订单商品id
	//[ 2] oid                                            bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	Oid int64 `gorm:"column:oid;type:bigint;default:0;" json:"oid"` // 订单id
	//[ 3] talent_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	TalentID uint64 `gorm:"column:talent_id;type:ubigint;default:0;" json:"talent_id"` // 达人id
	//[ 4] anchor_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	AnchorID uint64 `gorm:"column:anchor_id;type:ubigint;default:0;" json:"anchor_id"` // 主播id
	//[ 5] shop_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShopID int32 `gorm:"column:shop_id;type:int;default:0;" json:"shop_id"` // 店铺ID
	//[ 6] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"` // 商家ID
	//[ 7] seller_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SellerID int32 `gorm:"column:seller_id;type:int;default:0;" json:"seller_id"` // 第三方商家ID
	//[ 8] order_status                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	OrderStatus int32 `gorm:"column:order_status;type:smallint;default:0;" json:"order_status"` // 子订单状态
	//[ 9] sku_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SkuID int64 `gorm:"column:sku_id;type:bigint;default:0;" json:"sku_id"` // 快手商品skuid
	//[10] rel_sku_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	RelSkuID int64 `gorm:"column:rel_sku_id;type:bigint;default:0;" json:"rel_sku_id"` // 服务商商品skuid
	//[11] sku_desc                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SkuDesc string `gorm:"column:sku_desc;type:varchar;size:255;" json:"sku_desc"` // sku描述
	//[12] sku_nick                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SkuNick string `gorm:"column:sku_nick;type:varchar;size:255;" json:"sku_nick"` // sku编码
	//[13] item_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ItemID int64 `gorm:"column:item_id;type:bigint;default:0;" json:"item_id"` // 快手商品id
	//[14] yx_product_id                                  bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	YxProductID int64 `gorm:"column:yx_product_id;type:bigint;default:0;" json:"yx_product_id"` // yx商品id
	//[15] rel_item_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	RelItemID int64 `gorm:"column:rel_item_id;type:bigint;default:0;" json:"rel_item_id"` // 服务商商品id
	//[16] item_title                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ItemTitle string `gorm:"column:item_title;type:varchar;size:255;" json:"item_title"` // 商品名称
	//[17] item_link_url                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ItemLinkURL string `gorm:"column:item_link_url;type:varchar;size:255;" json:"item_link_url"` // 商品链接
	//[18] item_pic_url                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ItemPicURL string `gorm:"column:item_pic_url;type:varchar;size:255;" json:"item_pic_url"` // 商品图片地址
	//[19] num                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Num int32 `gorm:"column:num;type:int;default:0;" json:"num"` // 成交数量
	//[20] original_price                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OriginalPrice int32 `gorm:"column:original_price;type:int;default:0;" json:"original_price"` // 商品促销前单价快照 以分为单位
	//[21] discount_fee                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DiscountFee int32 `gorm:"column:discount_fee;type:int;default:0;" json:"discount_fee"` // 折扣金额
	//[22] price                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Price int32 `gorm:"column:price;type:int;default:0;" json:"price"` // 商品单价快照，以分为单位
	//[23] pay_time                                       bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PayTime int64 `gorm:"column:pay_time;type:bigint;default:0;" json:"pay_time"` // 支付时间
	//[24] create_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	CreateTime int64 `gorm:"column:create_time;type:bigint;default:0;" json:"create_time"` // 创建时间-毫秒时间戳
	//[25] update_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	UpdateTime int64 `gorm:"column:update_time;type:bigint;default:0;" json:"update_time"` // 更新时间-毫秒时间戳
	//[26] refund_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	RefundID int64 `gorm:"column:refund_id;type:bigint;default:0;" json:"refund_id"` // 退款订单id
	//[27] refund_status                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RefundStatus int32 `gorm:"column:refund_status;type:int;default:0;" json:"refund_status"` // 退款状态
	//[28] item_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemType int32 `gorm:"column:item_type;type:int;default:0;" json:"item_type"` // 1自建商品 2 闪电购商品
	//[29] refund_list                                    json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	RefundList sql.NullString `gorm:"column:refund_list;type:json;" json:"refund_list"` // 该skuId下的退款列表
	//[30] prev_info                                      json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	PrevInfo sql.NullString `gorm:"column:prev_info;type:json;" json:"prev_info"` // 订单修改之前的原订单信息
	//[31] distributor_id                                 bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	DistributorID int64 `gorm:"column:distributor_id;type:bigint;default:0;" json:"distributor_id"` // 推广者id
	//[32] distributor_name                               varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	DistributorName string `gorm:"column:distributor_name;type:varchar;size:125;" json:"distributor_name"` // 推广者昵称
	//[33] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 分销率
	//[34] estimated_income                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedIncome int32 `gorm:"column:estimated_income;type:int;default:0;" json:"estimated_income"` // 分销金额
	//[35] settlement_time                                bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SettlementTime int64 `gorm:"column:settlement_time;type:bigint;default:0;" json:"settlement_time"` // 结算时间
	//[36] settlement_success_time                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SettlementSuccessTime int64 `gorm:"column:settlement_success_time;type:bigint;default:0;" json:"settlement_success_time"` // 结算成功时间
	//[37] platform_dp_rate                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformDpRate int32 `gorm:"column:platform_dp_rate;type:int;default:0;" json:"platform_dp_rate"` // 平台分销推广佣金分成比例。千分比
	//[38] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // db写入时间
	//[39] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // db更新时间

}

var order_product_ksTableInfo = &TableInfo{
	Name: "order_product_ks",
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
			Name:               "ks_id",
			Comment:            `ks订单商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "KsID",
			GoFieldType:        "int64",
			JSONFieldName:      "ks_id",
			ProtobufFieldName:  "ks_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "oid",
			Comment:            `订单id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "merchant_id",
			Comment:            `商家ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "seller_id",
			Comment:            `第三方商家ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SellerID",
			GoFieldType:        "int32",
			JSONFieldName:      "seller_id",
			ProtobufFieldName:  "seller_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "order_status",
			Comment:            `子订单状态`,
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
			Name:               "sku_id",
			Comment:            `快手商品skuid`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SkuID",
			GoFieldType:        "int64",
			JSONFieldName:      "sku_id",
			ProtobufFieldName:  "sku_id",
			ProtobufType:       "int64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "rel_sku_id",
			Comment:            `服务商商品skuid`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RelSkuID",
			GoFieldType:        "int64",
			JSONFieldName:      "rel_sku_id",
			ProtobufFieldName:  "rel_sku_id",
			ProtobufType:       "int64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "sku_desc",
			Comment:            `sku描述`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SkuDesc",
			GoFieldType:        "string",
			JSONFieldName:      "sku_desc",
			ProtobufFieldName:  "sku_desc",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "sku_nick",
			Comment:            `sku编码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SkuNick",
			GoFieldType:        "string",
			JSONFieldName:      "sku_nick",
			ProtobufFieldName:  "sku_nick",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "item_id",
			Comment:            `快手商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ItemID",
			GoFieldType:        "int64",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int64",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "yx_product_id",
			Comment:            `yx商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "YxProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "yx_product_id",
			ProtobufFieldName:  "yx_product_id",
			ProtobufType:       "int64",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "rel_item_id",
			Comment:            `服务商商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RelItemID",
			GoFieldType:        "int64",
			JSONFieldName:      "rel_item_id",
			ProtobufFieldName:  "rel_item_id",
			ProtobufType:       "int64",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "item_title",
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
			GoFieldName:        "ItemTitle",
			GoFieldType:        "string",
			JSONFieldName:      "item_title",
			ProtobufFieldName:  "item_title",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "item_link_url",
			Comment:            `商品链接`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ItemLinkURL",
			GoFieldType:        "string",
			JSONFieldName:      "item_link_url",
			ProtobufFieldName:  "item_link_url",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "item_pic_url",
			Comment:            `商品图片地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ItemPicURL",
			GoFieldType:        "string",
			JSONFieldName:      "item_pic_url",
			ProtobufFieldName:  "item_pic_url",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "num",
			Comment:            `成交数量`,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "original_price",
			Comment:            `商品促销前单价快照 以分为单位`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OriginalPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "original_price",
			ProtobufFieldName:  "original_price",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "discount_fee",
			Comment:            `折扣金额`,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "price",
			Comment:            `商品单价快照，以分为单位`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Price",
			GoFieldType:        "int32",
			JSONFieldName:      "price",
			ProtobufFieldName:  "price",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "pay_time",
			Comment:            `支付时间`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "create_time",
			Comment:            `创建时间-毫秒时间戳`,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "update_time",
			Comment:            `更新时间-毫秒时间戳`,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "refund_id",
			Comment:            `退款订单id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "RefundID",
			GoFieldType:        "int64",
			JSONFieldName:      "refund_id",
			ProtobufFieldName:  "refund_id",
			ProtobufType:       "int64",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "refund_status",
			Comment:            `退款状态`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RefundStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "refund_status",
			ProtobufFieldName:  "refund_status",
			ProtobufType:       "int32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "item_type",
			Comment:            `1自建商品 2 闪电购商品`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ItemType",
			GoFieldType:        "int32",
			JSONFieldName:      "item_type",
			ProtobufFieldName:  "item_type",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "refund_list",
			Comment:            `该skuId下的退款列表`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "RefundList",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "refund_list",
			ProtobufFieldName:  "refund_list",
			ProtobufType:       "string",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "prev_info",
			Comment:            `订单修改之前的原订单信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "PrevInfo",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "prev_info",
			ProtobufFieldName:  "prev_info",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "distributor_id",
			Comment:            `推广者id`,
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
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "distributor_name",
			Comment:            `推广者昵称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "DistributorName",
			GoFieldType:        "string",
			JSONFieldName:      "distributor_name",
			ProtobufFieldName:  "distributor_name",
			ProtobufType:       "string",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "commission_rate",
			Comment:            `分销率`,
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
			Comment:            `分销金额`,
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
			Name:               "platform_dp_rate",
			Comment:            `平台分销推广佣金分成比例。千分比`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformDpRate",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_dp_rate",
			ProtobufFieldName:  "platform_dp_rate",
			ProtobufType:       "int32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
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
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
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
			ProtobufPos:        40,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderProductKs) TableName() string {
	return "order_product_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderProductKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderProductKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderProductKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderProductKs) TableInfo() *TableInfo {
	return order_product_ksTableInfo
}
