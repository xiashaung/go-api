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


CREATE TABLE `dy_order_history` (
  `item_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `order_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '订单号',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `product_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `product_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品图片',
  `author_account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '达人昵称',
  `author_openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '达人open_id',
  `shop_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家店铺名称',
  `total_pay_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '订单支付金额，单位分',
  `commission_rate` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人佣金率，此处保存为真实数据x1万之后，如真实是0.35，这里是3500',
  `flow_point` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)',
  `app` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'App名称（抖音，火山）',
  `buyer_openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '下单用户open_id',
  `update_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新时间 [联盟侧订单更新时间]',
  `pay_success_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '付款时间',
  `settle_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '结算时间，结算前为空字符串',
  `pay_goods_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '预估参与结算金额',
  `settled_goods_amount` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '实际参与结算金额',
  `real_commission` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人实际佣金',
  `estimated_commission` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人预估佣金',
  `item_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品数目',
  `shop_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '店铺id',
  `refund_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '退款订单退款时间',
  `pid_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '分销订单相关参数',
  `estimated_total_commission` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '总佣金（预估），对应百应订单明细中的总佣金',
  `estimated_tech_service_fee` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '预估平台技术服务费',
  `pick_source_client_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选品App client_key',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`item_id`),
  UNIQUE KEY `uni_order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "item_id": 24,    "order_id": 71,    "product_id": 45,    "product_name": "FmVeZQGaZCZwqKxbyvrswVwOm",    "product_img": "oaYZbGiiyHdMNrTvkqaGuouTf",    "author_account": "YDIaGGeOEBcLjRdBKEuLhOfoe",    "author_openid": "tAnApBrZlyWaKojyYibOwbFUr",    "shop_name": "NygJMpXdLediAfMomvOJXvkbC",    "total_pay_amount": 14,    "commission_rate": 3,    "flow_point": "NGZyQoqxnuiLjIbdkwwaOxlTy",    "app": "HCyjhjxvaeocgGqvimYkCZNOM",    "buyer_openid": "ehUTHSaeOVnNFhSMIWkZNrwgE",    "update_time": "xykEdTDfIaTSlGcfWawQvwWNb",    "pay_success_time": "UWUKUSqvfpKRVOMXZGcvNcdsk",    "settle_time": "kUboBDPmRhTFkbfvqRjUEbuEc",    "pay_goods_amount": 16,    "settled_goods_amount": 35,    "real_commission": 24,    "estimated_commission": 53,    "item_num": 19,    "shop_id": 63,    "refund_time": "TvgywMimmtETTrxbblajKMrqu",    "pid_info": "DgdABtZeNSilSiaTACNZSrIWk",    "estimated_total_commission": 94,    "estimated_tech_service_fee": 89,    "pick_source_client_key": "xYIPQHFMHVdlTaFLyOfHebTXn",    "add_time": 78,    "last_time": 49}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[16] column is set for unsigned
[17] column is set for unsigned
[18] column is set for unsigned
[19] column is set for unsigned
[20] column is set for unsigned
[21] column is set for unsigned
[24] column is set for unsigned
[25] column is set for unsigned



*/

// DyOrderHistory struct is a row record of the dy_order_history table in the yx database
type DyOrderHistory struct {
	//[ 0] item_id                                        ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ItemID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:item_id;type:ubigint;" json:"item_id"` // 自增主键
	//[ 1] order_id                                       ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	OrderID uint64 `gorm:"column:order_id;type:ubigint;default:0;" json:"order_id"` // 订单号
	//[ 2] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品id
	//[ 3] product_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:255;" json:"product_name"` // 商品名称
	//[ 4] product_img                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductImg string `gorm:"column:product_img;type:varchar;size:255;" json:"product_img"` // 商品图片
	//[ 5] author_account                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AuthorAccount string `gorm:"column:author_account;type:varchar;size:255;" json:"author_account"` // 达人昵称
	//[ 6] author_openid                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AuthorOpenid string `gorm:"column:author_openid;type:varchar;size:255;" json:"author_openid"` // 达人open_id
	//[ 7] shop_name                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ShopName string `gorm:"column:shop_name;type:varchar;size:255;" json:"shop_name"` // 商家店铺名称
	//[ 8] total_pay_amount                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TotalPayAmount uint32 `gorm:"column:total_pay_amount;type:uint;default:0;" json:"total_pay_amount"` // 订单支付金额，单位分
	//[ 9] commission_rate                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CommissionRate uint32 `gorm:"column:commission_rate;type:uint;default:0;" json:"commission_rate"` // 达人佣金率，此处保存为真实数据x1万之后，如真实是0.35，这里是3500
	//[10] flow_point                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FlowPoint string `gorm:"column:flow_point;type:varchar;size:255;" json:"flow_point"` // 订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)
	//[11] app                                            varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	App string `gorm:"column:app;type:varchar;size:255;" json:"app"` // App名称（抖音，火山）
	//[12] buyer_openid                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BuyerOpenid string `gorm:"column:buyer_openid;type:varchar;size:255;" json:"buyer_openid"` // 下单用户open_id
	//[13] update_time                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	UpdateTime string `gorm:"column:update_time;type:varchar;size:255;" json:"update_time"` // 更新时间 [联盟侧订单更新时间]
	//[14] pay_success_time                               varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	PaySuccessTime string `gorm:"column:pay_success_time;type:varchar;size:255;" json:"pay_success_time"` // 付款时间
	//[15] settle_time                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SettleTime string `gorm:"column:settle_time;type:varchar;size:255;" json:"settle_time"` // 结算时间，结算前为空字符串
	//[16] pay_goods_amount                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayGoodsAmount uint32 `gorm:"column:pay_goods_amount;type:uint;default:0;" json:"pay_goods_amount"` // 预估参与结算金额
	//[17] settled_goods_amount                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SettledGoodsAmount uint32 `gorm:"column:settled_goods_amount;type:uint;default:0;" json:"settled_goods_amount"` // 实际参与结算金额
	//[18] real_commission                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	RealCommission uint32 `gorm:"column:real_commission;type:uint;default:0;" json:"real_commission"` // 达人实际佣金
	//[19] estimated_commission                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EstimatedCommission uint32 `gorm:"column:estimated_commission;type:uint;default:0;" json:"estimated_commission"` // 达人预估佣金
	//[20] item_num                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ItemNum uint32 `gorm:"column:item_num;type:uint;default:0;" json:"item_num"` // 商品数目
	//[21] shop_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ShopID uint32 `gorm:"column:shop_id;type:uint;default:0;" json:"shop_id"` // 店铺id
	//[22] refund_time                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	RefundTime string `gorm:"column:refund_time;type:varchar;size:255;" json:"refund_time"` // 退款订单退款时间
	//[23] pid_info                                       text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	PidInfo string `gorm:"column:pid_info;type:text;size:65535;" json:"pid_info"` // 分销订单相关参数
	//[24] estimated_total_commission                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EstimatedTotalCommission uint32 `gorm:"column:estimated_total_commission;type:uint;default:0;" json:"estimated_total_commission"` // 总佣金（预估），对应百应订单明细中的总佣金
	//[25] estimated_tech_service_fee                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EstimatedTechServiceFee uint32 `gorm:"column:estimated_tech_service_fee;type:uint;default:0;" json:"estimated_tech_service_fee"` // 预估平台技术服务费
	//[26] pick_source_client_key                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	PickSourceClientKey string `gorm:"column:pick_source_client_key;type:varchar;size:255;" json:"pick_source_client_key"` // 选品App client_key
	//[27] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加时间
	//[28] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 更新时间

}

var dy_order_historyTableInfo = &TableInfo{
	Name: "dy_order_history",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "item_id",
			Comment:            `自增主键`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ItemID",
			GoFieldType:        "uint64",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "order_id",
			Comment:            `订单号`,
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
			Name:               "product_id",
			Comment:            `商品id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "product_img",
			Comment:            `商品图片`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ProductImg",
			GoFieldType:        "string",
			JSONFieldName:      "product_img",
			ProtobufFieldName:  "product_img",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "author_account",
			Comment:            `达人昵称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "AuthorAccount",
			GoFieldType:        "string",
			JSONFieldName:      "author_account",
			ProtobufFieldName:  "author_account",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "author_openid",
			Comment:            `达人open_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "AuthorOpenid",
			GoFieldType:        "string",
			JSONFieldName:      "author_openid",
			ProtobufFieldName:  "author_openid",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "shop_name",
			Comment:            `商家店铺名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ShopName",
			GoFieldType:        "string",
			JSONFieldName:      "shop_name",
			ProtobufFieldName:  "shop_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "total_pay_amount",
			Comment:            `订单支付金额，单位分`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TotalPayAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "total_pay_amount",
			ProtobufFieldName:  "total_pay_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "commission_rate",
			Comment:            `达人佣金率，此处保存为真实数据x1万之后，如真实是0.35，这里是3500`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CommissionRate",
			GoFieldType:        "uint32",
			JSONFieldName:      "commission_rate",
			ProtobufFieldName:  "commission_rate",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "flow_point",
			Comment:            `订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "FlowPoint",
			GoFieldType:        "string",
			JSONFieldName:      "flow_point",
			ProtobufFieldName:  "flow_point",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "app",
			Comment:            `App名称（抖音，火山）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "App",
			GoFieldType:        "string",
			JSONFieldName:      "app",
			ProtobufFieldName:  "app",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "buyer_openid",
			Comment:            `下单用户open_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "BuyerOpenid",
			GoFieldType:        "string",
			JSONFieldName:      "buyer_openid",
			ProtobufFieldName:  "buyer_openid",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "update_time",
			Comment:            `更新时间 [联盟侧订单更新时间]`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "string",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "pay_success_time",
			Comment:            `付款时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "PaySuccessTime",
			GoFieldType:        "string",
			JSONFieldName:      "pay_success_time",
			ProtobufFieldName:  "pay_success_time",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "settle_time",
			Comment:            `结算时间，结算前为空字符串`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SettleTime",
			GoFieldType:        "string",
			JSONFieldName:      "settle_time",
			ProtobufFieldName:  "settle_time",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "pay_goods_amount",
			Comment:            `预估参与结算金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PayGoodsAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "pay_goods_amount",
			ProtobufFieldName:  "pay_goods_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "settled_goods_amount",
			Comment:            `实际参与结算金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SettledGoodsAmount",
			GoFieldType:        "uint32",
			JSONFieldName:      "settled_goods_amount",
			ProtobufFieldName:  "settled_goods_amount",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "real_commission",
			Comment:            `达人实际佣金`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "RealCommission",
			GoFieldType:        "uint32",
			JSONFieldName:      "real_commission",
			ProtobufFieldName:  "real_commission",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "estimated_commission",
			Comment:            `达人预估佣金`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedCommission",
			GoFieldType:        "uint32",
			JSONFieldName:      "estimated_commission",
			ProtobufFieldName:  "estimated_commission",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "item_num",
			Comment:            `商品数目`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ItemNum",
			GoFieldType:        "uint32",
			JSONFieldName:      "item_num",
			ProtobufFieldName:  "item_num",
			ProtobufType:       "uint32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "shop_id",
			Comment:            `店铺id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "uint32",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "uint32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "refund_time",
			Comment:            `退款订单退款时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "RefundTime",
			GoFieldType:        "string",
			JSONFieldName:      "refund_time",
			ProtobufFieldName:  "refund_time",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "pid_info",
			Comment:            `分销订单相关参数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "PidInfo",
			GoFieldType:        "string",
			JSONFieldName:      "pid_info",
			ProtobufFieldName:  "pid_info",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "estimated_total_commission",
			Comment:            `总佣金（预估），对应百应订单明细中的总佣金`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedTotalCommission",
			GoFieldType:        "uint32",
			JSONFieldName:      "estimated_total_commission",
			ProtobufFieldName:  "estimated_total_commission",
			ProtobufType:       "uint32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "estimated_tech_service_fee",
			Comment:            `预估平台技术服务费`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedTechServiceFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "estimated_tech_service_fee",
			ProtobufFieldName:  "estimated_tech_service_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "pick_source_client_key",
			Comment:            `选品App client_key`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "PickSourceClientKey",
			GoFieldType:        "string",
			JSONFieldName:      "pick_source_client_key",
			ProtobufFieldName:  "pick_source_client_key",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
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
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "last_time",
			Comment:            `更新时间`,
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
			ProtobufPos:        29,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DyOrderHistory) TableName() string {
	return "dy_order_history"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DyOrderHistory) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DyOrderHistory) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DyOrderHistory) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DyOrderHistory) TableInfo() *TableInfo {
	return dy_order_historyTableInfo
}
