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


CREATE TABLE `colonel_order_dy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '订单id',
  `merchant_id` int(11) NOT NULL DEFAULT '0',
  `talent_id` int(11) NOT NULL DEFAULT '0',
  `activity_id` int(11) NOT NULL DEFAULT '0' COMMENT '团长活动id',
  `third_product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '第三方商品id',
  `product_name` varchar(200) NOT NULL DEFAULT '' COMMENT '商品名称',
  `product_img` varchar(255) NOT NULL DEFAULT '' COMMENT '商品图片',
  `author_account` varchar(255) NOT NULL DEFAULT '' COMMENT '作者账号昵称(抖音/火山作者)',
  `author_openid` varchar(100) NOT NULL DEFAULT '' COMMENT '作者抖音open_id',
  `shop_name` varchar(100) NOT NULL DEFAULT '' COMMENT '店铺名称',
  `total_pay_amount` int(10) NOT NULL DEFAULT '0' COMMENT '订单支付金额',
  `commission_rate` int(10) NOT NULL DEFAULT '0' COMMENT '达人佣金率',
  `estimated_commission` int(10) NOT NULL DEFAULT '0' COMMENT '达人预估佣金收入',
  `real_commission` int(10) NOT NULL DEFAULT '0' COMMENT '达人实际佣金收入',
  `flow_point` varchar(20) NOT NULL DEFAULT '' COMMENT '订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)',
  `app` varchar(20) NOT NULL DEFAULT '' COMMENT 'App名称（抖音，火山）',
  `buyer_openid` varchar(60) NOT NULL DEFAULT '' COMMENT '下单用户open_id',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '抖音更新时间',
  `pay_success_time` int(11) NOT NULL DEFAULT '0' COMMENT '支付成功时间',
  `settle_time` int(11) NOT NULL DEFAULT '0' COMMENT '结算时间',
  `pay_goods_amount` int(10) NOT NULL DEFAULT '0' COMMENT '预估参与结算金额',
  `settled_goods_amount` int(10) NOT NULL DEFAULT '0' COMMENT '实际参与结算金额',
  `third_shop_id` int(11) NOT NULL DEFAULT '0' COMMENT '第三方店铺ID',
  `shop_estimated_commission` int(10) NOT NULL DEFAULT '0' COMMENT '商家预估佣金支出',
  `shop_real_commission` int(10) NOT NULL DEFAULT '0' COMMENT '商家实际佣金支出',
  `alliance_split_rate` int(10) NOT NULL DEFAULT '0' COMMENT '分成比例，仅机构绑定的达人订单里才有该字段',
  `tech_service_fee_rate` int(10) NOT NULL DEFAULT '0' COMMENT '平台技术服务费率',
  `estimated_tech_service_fee` int(10) NOT NULL DEFAULT '0' COMMENT '平台预估技术服务费',
  `settled_tech_service_fee` int(10) NOT NULL DEFAULT '0' COMMENT '平台结算技术服务费',
  `estimated_institution_commission` int(10) NOT NULL DEFAULT '0' COMMENT '预估机构分成，机构和达人有绑定关系时返回',
  `extra` varchar(255) NOT NULL DEFAULT '' COMMENT '其他',
  `item_num` int(10) NOT NULL DEFAULT '0' COMMENT '商品数目',
  `refund_time` int(10) NOT NULL DEFAULT '0' COMMENT '退款订单退款时间',
  `estimated_total_commission` int(255) NOT NULL DEFAULT '0' COMMENT '总佣金（预估），对应百应订单明细中的总佣金',
  `colonel_commission_rate` int(10) NOT NULL DEFAULT '0' COMMENT '团长服务费率',
  `colonel_estimated_commission` int(10) NOT NULL DEFAULT '0' COMMENT '团长预估服务费',
  `colonel_real_commission` int(10) NOT NULL DEFAULT '0' COMMENT '团长结算服务费',
  `colonel_tech_service_fee_rate` int(10) NOT NULL DEFAULT '0' COMMENT '团长技术服务费率',
  `colonel_tech_service_fee` int(10) NOT NULL DEFAULT '0' COMMENT '团长预估技术服务费',
  `colonel_settled_tech_service_fee` int(10) NOT NULL DEFAULT '0' COMMENT '团长结算技术服务费',
  `colonel_institution_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '团长机构ID',
  `colonel_institution_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '团长机构名称',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=789 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 70,    "order_id": "YMqlGroWeTmLhSufFJGgybcUX",    "merchant_id": 12,    "talent_id": 52,    "activity_id": 18,    "third_product_id": 85,    "product_name": "aJWGywZMpNfngHvyqTkvEUunM",    "product_img": "BMbqpfSZenxEMpCXXFeadBZPY",    "author_account": "ZyPRVUZWfqDGRayHXtwUKLKYy",    "author_openid": "fvuZaXJwCBEAJvUQPrNDcoomb",    "shop_name": "tVpRhcmpoAtsduxUyjmvCZdOd",    "total_pay_amount": 28,    "commission_rate": 75,    "estimated_commission": 74,    "real_commission": 81,    "flow_point": "jrqwGAkOnjjToIBwNvayLBJgw",    "app": "jOLDyGtXbVJlKRUsKVutsVXbM",    "buyer_openid": "WCwFtrNsLIGYwxKpXfUZsknme",    "update_time": 96,    "pay_success_time": 73,    "settle_time": 5,    "pay_goods_amount": 8,    "settled_goods_amount": 93,    "third_shop_id": 78,    "shop_estimated_commission": 42,    "shop_real_commission": 30,    "alliance_split_rate": 70,    "tech_service_fee_rate": 33,    "estimated_tech_service_fee": 20,    "settled_tech_service_fee": 43,    "estimated_institution_commission": 3,    "extra": "OnlFYxBSjmXPwurrFDUkeDeDB",    "item_num": 54,    "refund_time": 14,    "estimated_total_commission": 47,    "colonel_commission_rate": 61,    "colonel_estimated_commission": 78,    "colonel_real_commission": 67,    "colonel_tech_service_fee_rate": 50,    "colonel_tech_service_fee": 85,    "colonel_settled_tech_service_fee": 72,    "colonel_institution_id": 89,    "colonel_institution_name": "OrYjyLCqCBxcCalZfebYCpvZm",    "add_time": 88,    "last_time": 59}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// ColonelOrderDy struct is a row record of the colonel_order_dy table in the yx database
type ColonelOrderDy struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] order_id                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	OrderID string `gorm:"column:order_id;type:varchar;size:64;" json:"order_id"` // 订单id
	//[ 2] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"`
	//[ 3] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"`
	//[ 4] activity_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ActivityID int32 `gorm:"column:activity_id;type:int;default:0;" json:"activity_id"` // 团长活动id
	//[ 5] third_product_id                               bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ThirdProductID int64 `gorm:"column:third_product_id;type:bigint;default:0;" json:"third_product_id"` // 第三方商品id
	//[ 6] product_name                                   varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:200;" json:"product_name"` // 商品名称
	//[ 7] product_img                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductImg string `gorm:"column:product_img;type:varchar;size:255;" json:"product_img"` // 商品图片
	//[ 8] author_account                                 varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AuthorAccount string `gorm:"column:author_account;type:varchar;size:255;" json:"author_account"` // 作者账号昵称(抖音/火山作者)
	//[ 9] author_openid                                  varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	AuthorOpenid string `gorm:"column:author_openid;type:varchar;size:100;" json:"author_openid"` // 作者抖音open_id
	//[10] shop_name                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ShopName string `gorm:"column:shop_name;type:varchar;size:100;" json:"shop_name"` // 店铺名称
	//[11] total_pay_amount                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TotalPayAmount int32 `gorm:"column:total_pay_amount;type:int;default:0;" json:"total_pay_amount"` // 订单支付金额
	//[12] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 达人佣金率
	//[13] estimated_commission                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedCommission int32 `gorm:"column:estimated_commission;type:int;default:0;" json:"estimated_commission"` // 达人预估佣金收入
	//[14] real_commission                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RealCommission int32 `gorm:"column:real_commission;type:int;default:0;" json:"real_commission"` // 达人实际佣金收入
	//[15] flow_point                                     varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	FlowPoint string `gorm:"column:flow_point;type:varchar;size:20;" json:"flow_point"` // 订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)
	//[16] app                                            varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	App string `gorm:"column:app;type:varchar;size:20;" json:"app"` // App名称（抖音，火山）
	//[17] buyer_openid                                   varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	BuyerOpenid string `gorm:"column:buyer_openid;type:varchar;size:60;" json:"buyer_openid"` // 下单用户open_id
	//[18] update_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UpdateTime int32 `gorm:"column:update_time;type:int;default:0;" json:"update_time"` // 抖音更新时间
	//[19] pay_success_time                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PaySuccessTime int32 `gorm:"column:pay_success_time;type:int;default:0;" json:"pay_success_time"` // 支付成功时间
	//[20] settle_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SettleTime int32 `gorm:"column:settle_time;type:int;default:0;" json:"settle_time"` // 结算时间
	//[21] pay_goods_amount                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PayGoodsAmount int32 `gorm:"column:pay_goods_amount;type:int;default:0;" json:"pay_goods_amount"` // 预估参与结算金额
	//[22] settled_goods_amount                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SettledGoodsAmount int32 `gorm:"column:settled_goods_amount;type:int;default:0;" json:"settled_goods_amount"` // 实际参与结算金额
	//[23] third_shop_id                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ThirdShopID int32 `gorm:"column:third_shop_id;type:int;default:0;" json:"third_shop_id"` // 第三方店铺ID
	//[24] shop_estimated_commission                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShopEstimatedCommission int32 `gorm:"column:shop_estimated_commission;type:int;default:0;" json:"shop_estimated_commission"` // 商家预估佣金支出
	//[25] shop_real_commission                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ShopRealCommission int32 `gorm:"column:shop_real_commission;type:int;default:0;" json:"shop_real_commission"` // 商家实际佣金支出
	//[26] alliance_split_rate                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AllianceSplitRate int32 `gorm:"column:alliance_split_rate;type:int;default:0;" json:"alliance_split_rate"` // 分成比例，仅机构绑定的达人订单里才有该字段
	//[27] tech_service_fee_rate                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TechServiceFeeRate int32 `gorm:"column:tech_service_fee_rate;type:int;default:0;" json:"tech_service_fee_rate"` // 平台技术服务费率
	//[28] estimated_tech_service_fee                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedTechServiceFee int32 `gorm:"column:estimated_tech_service_fee;type:int;default:0;" json:"estimated_tech_service_fee"` // 平台预估技术服务费
	//[29] settled_tech_service_fee                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SettledTechServiceFee int32 `gorm:"column:settled_tech_service_fee;type:int;default:0;" json:"settled_tech_service_fee"` // 平台结算技术服务费
	//[30] estimated_institution_commission               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedInstitutionCommission int32 `gorm:"column:estimated_institution_commission;type:int;default:0;" json:"estimated_institution_commission"` // 预估机构分成，机构和达人有绑定关系时返回
	//[31] extra                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Extra string `gorm:"column:extra;type:varchar;size:255;" json:"extra"` // 其他
	//[32] item_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemNum int32 `gorm:"column:item_num;type:int;default:0;" json:"item_num"` // 商品数目
	//[33] refund_time                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RefundTime int32 `gorm:"column:refund_time;type:int;default:0;" json:"refund_time"` // 退款订单退款时间
	//[34] estimated_total_commission                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EstimatedTotalCommission int32 `gorm:"column:estimated_total_commission;type:int;default:0;" json:"estimated_total_commission"` // 总佣金（预估），对应百应订单明细中的总佣金
	//[35] colonel_commission_rate                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ColonelCommissionRate int32 `gorm:"column:colonel_commission_rate;type:int;default:0;" json:"colonel_commission_rate"` // 团长服务费率
	//[36] colonel_estimated_commission                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ColonelEstimatedCommission int32 `gorm:"column:colonel_estimated_commission;type:int;default:0;" json:"colonel_estimated_commission"` // 团长预估服务费
	//[37] colonel_real_commission                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ColonelRealCommission int32 `gorm:"column:colonel_real_commission;type:int;default:0;" json:"colonel_real_commission"` // 团长结算服务费
	//[38] colonel_tech_service_fee_rate                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ColonelTechServiceFeeRate int32 `gorm:"column:colonel_tech_service_fee_rate;type:int;default:0;" json:"colonel_tech_service_fee_rate"` // 团长技术服务费率
	//[39] colonel_tech_service_fee                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ColonelTechServiceFee int32 `gorm:"column:colonel_tech_service_fee;type:int;default:0;" json:"colonel_tech_service_fee"` // 团长预估技术服务费
	//[40] colonel_settled_tech_service_fee               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ColonelSettledTechServiceFee int32 `gorm:"column:colonel_settled_tech_service_fee;type:int;default:0;" json:"colonel_settled_tech_service_fee"` // 团长结算技术服务费
	//[41] colonel_institution_id                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ColonelInstitutionID int64 `gorm:"column:colonel_institution_id;type:bigint;default:0;" json:"colonel_institution_id"` // 团长机构ID
	//[42] colonel_institution_name                       varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ColonelInstitutionName string `gorm:"column:colonel_institution_name;type:varchar;size:30;" json:"colonel_institution_name"` // 团长机构名称
	//[43] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[44] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var colonel_order_dyTableInfo = &TableInfo{
	Name: "colonel_order_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "order_id",
			Comment:            `订单id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "OrderID",
			GoFieldType:        "string",
			JSONFieldName:      "order_id",
			ProtobufFieldName:  "order_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "merchant_id",
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
			GoFieldName:        "MerchantID",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "talent_id",
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
			GoFieldName:        "TalentID",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "activity_id",
			Comment:            `团长活动id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ActivityID",
			GoFieldType:        "int32",
			JSONFieldName:      "activity_id",
			ProtobufFieldName:  "activity_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "third_product_id",
			Comment:            `第三方商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "third_product_id",
			ProtobufFieldName:  "third_product_id",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "product_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "ProductName",
			GoFieldType:        "string",
			JSONFieldName:      "product_name",
			ProtobufFieldName:  "product_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "author_account",
			Comment:            `作者账号昵称(抖音/火山作者)`,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "author_openid",
			Comment:            `作者抖音open_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "AuthorOpenid",
			GoFieldType:        "string",
			JSONFieldName:      "author_openid",
			ProtobufFieldName:  "author_openid",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "shop_name",
			Comment:            `店铺名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ShopName",
			GoFieldType:        "string",
			JSONFieldName:      "shop_name",
			ProtobufFieldName:  "shop_name",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "total_pay_amount",
			Comment:            `订单支付金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TotalPayAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "total_pay_amount",
			ProtobufFieldName:  "total_pay_amount",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "commission_rate",
			Comment:            `达人佣金率`,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "estimated_commission",
			Comment:            `达人预估佣金收入`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "real_commission",
			Comment:            `达人实际佣金收入`,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "flow_point",
			Comment:            `订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "FlowPoint",
			GoFieldType:        "string",
			JSONFieldName:      "flow_point",
			ProtobufFieldName:  "flow_point",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "app",
			Comment:            `App名称（抖音，火山）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "App",
			GoFieldType:        "string",
			JSONFieldName:      "app",
			ProtobufFieldName:  "app",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "buyer_openid",
			Comment:            `下单用户open_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "BuyerOpenid",
			GoFieldType:        "string",
			JSONFieldName:      "buyer_openid",
			ProtobufFieldName:  "buyer_openid",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "update_time",
			Comment:            `抖音更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UpdateTime",
			GoFieldType:        "int32",
			JSONFieldName:      "update_time",
			ProtobufFieldName:  "update_time",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "pay_success_time",
			Comment:            `支付成功时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PaySuccessTime",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_success_time",
			ProtobufFieldName:  "pay_success_time",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "settle_time",
			Comment:            `结算时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SettleTime",
			GoFieldType:        "int32",
			JSONFieldName:      "settle_time",
			ProtobufFieldName:  "settle_time",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "pay_goods_amount",
			Comment:            `预估参与结算金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PayGoodsAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_goods_amount",
			ProtobufFieldName:  "pay_goods_amount",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "settled_goods_amount",
			Comment:            `实际参与结算金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SettledGoodsAmount",
			GoFieldType:        "int32",
			JSONFieldName:      "settled_goods_amount",
			ProtobufFieldName:  "settled_goods_amount",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "third_shop_id",
			Comment:            `第三方店铺ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ThirdShopID",
			GoFieldType:        "int32",
			JSONFieldName:      "third_shop_id",
			ProtobufFieldName:  "third_shop_id",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "shop_estimated_commission",
			Comment:            `商家预估佣金支出`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ShopEstimatedCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "shop_estimated_commission",
			ProtobufFieldName:  "shop_estimated_commission",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "shop_real_commission",
			Comment:            `商家实际佣金支出`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ShopRealCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "shop_real_commission",
			ProtobufFieldName:  "shop_real_commission",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "alliance_split_rate",
			Comment:            `分成比例，仅机构绑定的达人订单里才有该字段`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AllianceSplitRate",
			GoFieldType:        "int32",
			JSONFieldName:      "alliance_split_rate",
			ProtobufFieldName:  "alliance_split_rate",
			ProtobufType:       "int32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "tech_service_fee_rate",
			Comment:            `平台技术服务费率`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TechServiceFeeRate",
			GoFieldType:        "int32",
			JSONFieldName:      "tech_service_fee_rate",
			ProtobufFieldName:  "tech_service_fee_rate",
			ProtobufType:       "int32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "estimated_tech_service_fee",
			Comment:            `平台预估技术服务费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedTechServiceFee",
			GoFieldType:        "int32",
			JSONFieldName:      "estimated_tech_service_fee",
			ProtobufFieldName:  "estimated_tech_service_fee",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "settled_tech_service_fee",
			Comment:            `平台结算技术服务费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SettledTechServiceFee",
			GoFieldType:        "int32",
			JSONFieldName:      "settled_tech_service_fee",
			ProtobufFieldName:  "settled_tech_service_fee",
			ProtobufType:       "int32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "estimated_institution_commission",
			Comment:            `预估机构分成，机构和达人有绑定关系时返回`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedInstitutionCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "estimated_institution_commission",
			ProtobufFieldName:  "estimated_institution_commission",
			ProtobufType:       "int32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "extra",
			Comment:            `其他`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Extra",
			GoFieldType:        "string",
			JSONFieldName:      "extra",
			ProtobufFieldName:  "extra",
			ProtobufType:       "string",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "item_num",
			Comment:            `商品数目`,
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
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "refund_time",
			Comment:            `退款订单退款时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RefundTime",
			GoFieldType:        "int32",
			JSONFieldName:      "refund_time",
			ProtobufFieldName:  "refund_time",
			ProtobufType:       "int32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "estimated_total_commission",
			Comment:            `总佣金（预估），对应百应订单明细中的总佣金`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EstimatedTotalCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "estimated_total_commission",
			ProtobufFieldName:  "estimated_total_commission",
			ProtobufType:       "int32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "colonel_commission_rate",
			Comment:            `团长服务费率`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColonelCommissionRate",
			GoFieldType:        "int32",
			JSONFieldName:      "colonel_commission_rate",
			ProtobufFieldName:  "colonel_commission_rate",
			ProtobufType:       "int32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "colonel_estimated_commission",
			Comment:            `团长预估服务费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColonelEstimatedCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "colonel_estimated_commission",
			ProtobufFieldName:  "colonel_estimated_commission",
			ProtobufType:       "int32",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "colonel_real_commission",
			Comment:            `团长结算服务费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColonelRealCommission",
			GoFieldType:        "int32",
			JSONFieldName:      "colonel_real_commission",
			ProtobufFieldName:  "colonel_real_commission",
			ProtobufType:       "int32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
			Name:               "colonel_tech_service_fee_rate",
			Comment:            `团长技术服务费率`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColonelTechServiceFeeRate",
			GoFieldType:        "int32",
			JSONFieldName:      "colonel_tech_service_fee_rate",
			ProtobufFieldName:  "colonel_tech_service_fee_rate",
			ProtobufType:       "int32",
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
			Name:               "colonel_tech_service_fee",
			Comment:            `团长预估技术服务费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColonelTechServiceFee",
			GoFieldType:        "int32",
			JSONFieldName:      "colonel_tech_service_fee",
			ProtobufFieldName:  "colonel_tech_service_fee",
			ProtobufType:       "int32",
			ProtobufPos:        40,
		},

		&ColumnInfo{
			Index:              40,
			Name:               "colonel_settled_tech_service_fee",
			Comment:            `团长结算技术服务费`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ColonelSettledTechServiceFee",
			GoFieldType:        "int32",
			JSONFieldName:      "colonel_settled_tech_service_fee",
			ProtobufFieldName:  "colonel_settled_tech_service_fee",
			ProtobufType:       "int32",
			ProtobufPos:        41,
		},

		&ColumnInfo{
			Index:              41,
			Name:               "colonel_institution_id",
			Comment:            `团长机构ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ColonelInstitutionID",
			GoFieldType:        "int64",
			JSONFieldName:      "colonel_institution_id",
			ProtobufFieldName:  "colonel_institution_id",
			ProtobufType:       "int64",
			ProtobufPos:        42,
		},

		&ColumnInfo{
			Index:              42,
			Name:               "colonel_institution_name",
			Comment:            `团长机构名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "ColonelInstitutionName",
			GoFieldType:        "string",
			JSONFieldName:      "colonel_institution_name",
			ProtobufFieldName:  "colonel_institution_name",
			ProtobufType:       "string",
			ProtobufPos:        43,
		},

		&ColumnInfo{
			Index:              43,
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
			ProtobufPos:        44,
		},

		&ColumnInfo{
			Index:              44,
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
			ProtobufPos:        45,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *ColonelOrderDy) TableName() string {
	return "colonel_order_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *ColonelOrderDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *ColonelOrderDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *ColonelOrderDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *ColonelOrderDy) TableInfo() *TableInfo {
	return colonel_order_dyTableInfo
}
