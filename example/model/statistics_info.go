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


CREATE TABLE `statistics_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `add_merchant` int(6) NOT NULL DEFAULT '0' COMMENT '新增商家数',
  `add_merchant_auth` int(6) NOT NULL DEFAULT '0' COMMENT '新增授权商家数',
  `add_talent` int(6) NOT NULL DEFAULT '0' COMMENT '新增达人数',
  `add_talent_auth` int(6) NOT NULL DEFAULT '0' COMMENT '新增授权达人数',
  `uv` int(8) NOT NULL DEFAULT '0' COMMENT '平台uv',
  `pv` int(10) NOT NULL DEFAULT '0' COMMENT '平台pv',
  `putaway_product` int(10) NOT NULL DEFAULT '0' COMMENT '上架商品数',
  `putaway_product_merchant` int(10) NOT NULL DEFAULT '0' COMMENT '上架商品商家数',
  `promotion_plan_product` int(10) NOT NULL DEFAULT '0' COMMENT '推广计划商品数',
  `promotion_plan_merchant` int(10) NOT NULL DEFAULT '0' COMMENT '推广计划商家数',
  `super_cate_product` int(10) NOT NULL DEFAULT '0' COMMENT '超拥商品数',
  `super_cate_merchant` int(10) NOT NULL DEFAULT '0' COMMENT '超拥商家数',
  `plan` int(10) NOT NULL DEFAULT '0' COMMENT '招商计划数',
  `plan_talent` int(10) NOT NULL DEFAULT '0' COMMENT '招商计划达人数',
  `live` int(10) NOT NULL DEFAULT '0' COMMENT '直播场次数',
  `live_talent` int(10) NOT NULL DEFAULT '0' COMMENT '直播场次达人数',
  `xp_pool` int(10) NOT NULL DEFAULT '0' COMMENT '选品池商品数',
  `xp_pool_talent` int(10) NOT NULL DEFAULT '0' COMMENT '选品池达人数',
  `db_pool` int(10) NOT NULL DEFAULT '0' COMMENT '待播池商品数',
  `db_pool_talent` int(10) NOT NULL DEFAULT '0' COMMENT '待播池达人数',
  `login_merchant_auth` int(8) NOT NULL DEFAULT '0' COMMENT '登录商家数（授权）',
  `login_merchant_unauth` int(8) NOT NULL DEFAULT '0' COMMENT '登录商家数（未授权）',
  `login_talent_auth` int(8) NOT NULL DEFAULT '0' COMMENT '登录达人数（授权）',
  `login_talent_unauth` int(8) NOT NULL DEFAULT '0' COMMENT '登录达人数（未授权）',
  `free_order` int(10) NOT NULL DEFAULT '0' COMMENT '申请免费样品订单数',
  `pass_free_order` int(10) NOT NULL DEFAULT '0' COMMENT '审核通过免费样品订单数',
  `free_order_talent` int(8) NOT NULL DEFAULT '0' COMMENT '申请免费样品达人数',
  `pass_free_order_talent` int(8) NOT NULL DEFAULT '0' COMMENT '免费样品审核通达人数',
  `free_order_merchant` int(8) NOT NULL DEFAULT '0' COMMENT '被申请免费样品商家数',
  `pass_free_order_merchant` int(8) NOT NULL DEFAULT '0' COMMENT '免费样品审核通过商家数',
  `pay_order` int(8) NOT NULL DEFAULT '0' COMMENT '付费样品订单数',
  `sale_talent` int(8) NOT NULL DEFAULT '0' COMMENT '动销达人数',
  `sale_merchant` int(8) NOT NULL DEFAULT '0' COMMENT '动销商家数',
  `sale_product` int(8) NOT NULL DEFAULT '0' COMMENT '动销商品数',
  `sale` int(8) NOT NULL DEFAULT '0' COMMENT '成交销量',
  `gmv` int(11) NOT NULL DEFAULT '0' COMMENT '成交GMV（分）',
  `time` int(10) NOT NULL DEFAULT '0' COMMENT '统计时间 如统计的21年7月2号的数据',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `add_time` (`add_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='运营统计表'

JSON Sample
-------------------------------------
{    "id": 11,    "add_merchant": 16,    "add_merchant_auth": 28,    "add_talent": 7,    "add_talent_auth": 59,    "uv": 40,    "pv": 94,    "putaway_product": 57,    "putaway_product_merchant": 5,    "promotion_plan_product": 71,    "promotion_plan_merchant": 13,    "super_cate_product": 97,    "super_cate_merchant": 40,    "plan": 44,    "plan_talent": 39,    "live": 92,    "live_talent": 65,    "xp_pool": 17,    "xp_pool_talent": 82,    "db_pool": 32,    "db_pool_talent": 65,    "login_merchant_auth": 76,    "login_merchant_unauth": 22,    "login_talent_auth": 42,    "login_talent_unauth": 47,    "free_order": 77,    "pass_free_order": 62,    "free_order_talent": 15,    "pass_free_order_talent": 85,    "free_order_merchant": 94,    "pass_free_order_merchant": 42,    "pay_order": 81,    "sale_talent": 40,    "sale_merchant": 62,    "sale_product": 38,    "sale": 89,    "gmv": 99,    "time": 86,    "add_time": 24,    "last_time": 14}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// StatisticsInfo struct is a row record of the statistics_info table in the yx database
type StatisticsInfo struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] add_merchant                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddMerchant int32 `gorm:"column:add_merchant;type:int;default:0;" json:"add_merchant"` // 新增商家数
	//[ 2] add_merchant_auth                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddMerchantAuth int32 `gorm:"column:add_merchant_auth;type:int;default:0;" json:"add_merchant_auth"` // 新增授权商家数
	//[ 3] add_talent                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTalent int32 `gorm:"column:add_talent;type:int;default:0;" json:"add_talent"` // 新增达人数
	//[ 4] add_talent_auth                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTalentAuth int32 `gorm:"column:add_talent_auth;type:int;default:0;" json:"add_talent_auth"` // 新增授权达人数
	//[ 5] uv                                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Uv int32 `gorm:"column:uv;type:int;default:0;" json:"uv"` // 平台uv
	//[ 6] pv                                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Pv int32 `gorm:"column:pv;type:int;default:0;" json:"pv"` // 平台pv
	//[ 7] putaway_product                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PutawayProduct int32 `gorm:"column:putaway_product;type:int;default:0;" json:"putaway_product"` // 上架商品数
	//[ 8] putaway_product_merchant                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PutawayProductMerchant int32 `gorm:"column:putaway_product_merchant;type:int;default:0;" json:"putaway_product_merchant"` // 上架商品商家数
	//[ 9] promotion_plan_product                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PromotionPlanProduct int32 `gorm:"column:promotion_plan_product;type:int;default:0;" json:"promotion_plan_product"` // 推广计划商品数
	//[10] promotion_plan_merchant                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PromotionPlanMerchant int32 `gorm:"column:promotion_plan_merchant;type:int;default:0;" json:"promotion_plan_merchant"` // 推广计划商家数
	//[11] super_cate_product                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SuperCateProduct int32 `gorm:"column:super_cate_product;type:int;default:0;" json:"super_cate_product"` // 超拥商品数
	//[12] super_cate_merchant                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SuperCateMerchant int32 `gorm:"column:super_cate_merchant;type:int;default:0;" json:"super_cate_merchant"` // 超拥商家数
	//[13] plan                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Plan int32 `gorm:"column:plan;type:int;default:0;" json:"plan"` // 招商计划数
	//[14] plan_talent                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlanTalent int32 `gorm:"column:plan_talent;type:int;default:0;" json:"plan_talent"` // 招商计划达人数
	//[15] live                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Live int32 `gorm:"column:live;type:int;default:0;" json:"live"` // 直播场次数
	//[16] live_talent                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LiveTalent int32 `gorm:"column:live_talent;type:int;default:0;" json:"live_talent"` // 直播场次达人数
	//[17] xp_pool                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	XpPool int32 `gorm:"column:xp_pool;type:int;default:0;" json:"xp_pool"` // 选品池商品数
	//[18] xp_pool_talent                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	XpPoolTalent int32 `gorm:"column:xp_pool_talent;type:int;default:0;" json:"xp_pool_talent"` // 选品池达人数
	//[19] db_pool                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DbPool int32 `gorm:"column:db_pool;type:int;default:0;" json:"db_pool"` // 待播池商品数
	//[20] db_pool_talent                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DbPoolTalent int32 `gorm:"column:db_pool_talent;type:int;default:0;" json:"db_pool_talent"` // 待播池达人数
	//[21] login_merchant_auth                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LoginMerchantAuth int32 `gorm:"column:login_merchant_auth;type:int;default:0;" json:"login_merchant_auth"` // 登录商家数（授权）
	//[22] login_merchant_unauth                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LoginMerchantUnauth int32 `gorm:"column:login_merchant_unauth;type:int;default:0;" json:"login_merchant_unauth"` // 登录商家数（未授权）
	//[23] login_talent_auth                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LoginTalentAuth int32 `gorm:"column:login_talent_auth;type:int;default:0;" json:"login_talent_auth"` // 登录达人数（授权）
	//[24] login_talent_unauth                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LoginTalentUnauth int32 `gorm:"column:login_talent_unauth;type:int;default:0;" json:"login_talent_unauth"` // 登录达人数（未授权）
	//[25] free_order                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FreeOrder int32 `gorm:"column:free_order;type:int;default:0;" json:"free_order"` // 申请免费样品订单数
	//[26] pass_free_order                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PassFreeOrder int32 `gorm:"column:pass_free_order;type:int;default:0;" json:"pass_free_order"` // 审核通过免费样品订单数
	//[27] free_order_talent                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FreeOrderTalent int32 `gorm:"column:free_order_talent;type:int;default:0;" json:"free_order_talent"` // 申请免费样品达人数
	//[28] pass_free_order_talent                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PassFreeOrderTalent int32 `gorm:"column:pass_free_order_talent;type:int;default:0;" json:"pass_free_order_talent"` // 免费样品审核通达人数
	//[29] free_order_merchant                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FreeOrderMerchant int32 `gorm:"column:free_order_merchant;type:int;default:0;" json:"free_order_merchant"` // 被申请免费样品商家数
	//[30] pass_free_order_merchant                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PassFreeOrderMerchant int32 `gorm:"column:pass_free_order_merchant;type:int;default:0;" json:"pass_free_order_merchant"` // 免费样品审核通过商家数
	//[31] pay_order                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PayOrder int32 `gorm:"column:pay_order;type:int;default:0;" json:"pay_order"` // 付费样品订单数
	//[32] sale_talent                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SaleTalent int32 `gorm:"column:sale_talent;type:int;default:0;" json:"sale_talent"` // 动销达人数
	//[33] sale_merchant                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SaleMerchant int32 `gorm:"column:sale_merchant;type:int;default:0;" json:"sale_merchant"` // 动销商家数
	//[34] sale_product                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SaleProduct int32 `gorm:"column:sale_product;type:int;default:0;" json:"sale_product"` // 动销商品数
	//[35] sale                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Sale int32 `gorm:"column:sale;type:int;default:0;" json:"sale"` // 成交销量
	//[36] gmv                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Gmv int32 `gorm:"column:gmv;type:int;default:0;" json:"gmv"` // 成交GMV（分）
	//[37] time                                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Time int32 `gorm:"column:time;type:int;default:0;" json:"time"` // 统计时间 如统计的21年7月2号的数据
	//[38] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加时间
	//[39] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 更新时间

}

var statistics_infoTableInfo = &TableInfo{
	Name: "statistics_info",
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
			Name:               "add_merchant",
			Comment:            `新增商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "add_merchant",
			ProtobufFieldName:  "add_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "add_merchant_auth",
			Comment:            `新增授权商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddMerchantAuth",
			GoFieldType:        "int32",
			JSONFieldName:      "add_merchant_auth",
			ProtobufFieldName:  "add_merchant_auth",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "add_talent",
			Comment:            `新增达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "add_talent",
			ProtobufFieldName:  "add_talent",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_talent_auth",
			Comment:            `新增授权达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTalentAuth",
			GoFieldType:        "int32",
			JSONFieldName:      "add_talent_auth",
			ProtobufFieldName:  "add_talent_auth",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "uv",
			Comment:            `平台uv`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Uv",
			GoFieldType:        "int32",
			JSONFieldName:      "uv",
			ProtobufFieldName:  "uv",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "pv",
			Comment:            `平台pv`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Pv",
			GoFieldType:        "int32",
			JSONFieldName:      "pv",
			ProtobufFieldName:  "pv",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "putaway_product",
			Comment:            `上架商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PutawayProduct",
			GoFieldType:        "int32",
			JSONFieldName:      "putaway_product",
			ProtobufFieldName:  "putaway_product",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "putaway_product_merchant",
			Comment:            `上架商品商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PutawayProductMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "putaway_product_merchant",
			ProtobufFieldName:  "putaway_product_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "promotion_plan_product",
			Comment:            `推广计划商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PromotionPlanProduct",
			GoFieldType:        "int32",
			JSONFieldName:      "promotion_plan_product",
			ProtobufFieldName:  "promotion_plan_product",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "promotion_plan_merchant",
			Comment:            `推广计划商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PromotionPlanMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "promotion_plan_merchant",
			ProtobufFieldName:  "promotion_plan_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "super_cate_product",
			Comment:            `超拥商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SuperCateProduct",
			GoFieldType:        "int32",
			JSONFieldName:      "super_cate_product",
			ProtobufFieldName:  "super_cate_product",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "super_cate_merchant",
			Comment:            `超拥商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SuperCateMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "super_cate_merchant",
			ProtobufFieldName:  "super_cate_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "plan",
			Comment:            `招商计划数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Plan",
			GoFieldType:        "int32",
			JSONFieldName:      "plan",
			ProtobufFieldName:  "plan",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "plan_talent",
			Comment:            `招商计划达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlanTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "plan_talent",
			ProtobufFieldName:  "plan_talent",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "live",
			Comment:            `直播场次数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Live",
			GoFieldType:        "int32",
			JSONFieldName:      "live",
			ProtobufFieldName:  "live",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "live_talent",
			Comment:            `直播场次达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LiveTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "live_talent",
			ProtobufFieldName:  "live_talent",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "xp_pool",
			Comment:            `选品池商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "XpPool",
			GoFieldType:        "int32",
			JSONFieldName:      "xp_pool",
			ProtobufFieldName:  "xp_pool",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "xp_pool_talent",
			Comment:            `选品池达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "XpPoolTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "xp_pool_talent",
			ProtobufFieldName:  "xp_pool_talent",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "db_pool",
			Comment:            `待播池商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DbPool",
			GoFieldType:        "int32",
			JSONFieldName:      "db_pool",
			ProtobufFieldName:  "db_pool",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "db_pool_talent",
			Comment:            `待播池达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DbPoolTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "db_pool_talent",
			ProtobufFieldName:  "db_pool_talent",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "login_merchant_auth",
			Comment:            `登录商家数（授权）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LoginMerchantAuth",
			GoFieldType:        "int32",
			JSONFieldName:      "login_merchant_auth",
			ProtobufFieldName:  "login_merchant_auth",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "login_merchant_unauth",
			Comment:            `登录商家数（未授权）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LoginMerchantUnauth",
			GoFieldType:        "int32",
			JSONFieldName:      "login_merchant_unauth",
			ProtobufFieldName:  "login_merchant_unauth",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "login_talent_auth",
			Comment:            `登录达人数（授权）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LoginTalentAuth",
			GoFieldType:        "int32",
			JSONFieldName:      "login_talent_auth",
			ProtobufFieldName:  "login_talent_auth",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "login_talent_unauth",
			Comment:            `登录达人数（未授权）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LoginTalentUnauth",
			GoFieldType:        "int32",
			JSONFieldName:      "login_talent_unauth",
			ProtobufFieldName:  "login_talent_unauth",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "free_order",
			Comment:            `申请免费样品订单数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FreeOrder",
			GoFieldType:        "int32",
			JSONFieldName:      "free_order",
			ProtobufFieldName:  "free_order",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "pass_free_order",
			Comment:            `审核通过免费样品订单数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PassFreeOrder",
			GoFieldType:        "int32",
			JSONFieldName:      "pass_free_order",
			ProtobufFieldName:  "pass_free_order",
			ProtobufType:       "int32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "free_order_talent",
			Comment:            `申请免费样品达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FreeOrderTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "free_order_talent",
			ProtobufFieldName:  "free_order_talent",
			ProtobufType:       "int32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "pass_free_order_talent",
			Comment:            `免费样品审核通达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PassFreeOrderTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "pass_free_order_talent",
			ProtobufFieldName:  "pass_free_order_talent",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "free_order_merchant",
			Comment:            `被申请免费样品商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FreeOrderMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "free_order_merchant",
			ProtobufFieldName:  "free_order_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "pass_free_order_merchant",
			Comment:            `免费样品审核通过商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PassFreeOrderMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "pass_free_order_merchant",
			ProtobufFieldName:  "pass_free_order_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "pay_order",
			Comment:            `付费样品订单数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PayOrder",
			GoFieldType:        "int32",
			JSONFieldName:      "pay_order",
			ProtobufFieldName:  "pay_order",
			ProtobufType:       "int32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "sale_talent",
			Comment:            `动销达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleTalent",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_talent",
			ProtobufFieldName:  "sale_talent",
			ProtobufType:       "int32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "sale_merchant",
			Comment:            `动销商家数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_merchant",
			ProtobufFieldName:  "sale_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "sale_product",
			Comment:            `动销商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SaleProduct",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_product",
			ProtobufFieldName:  "sale_product",
			ProtobufType:       "int32",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "sale",
			Comment:            `成交销量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Sale",
			GoFieldType:        "int32",
			JSONFieldName:      "sale",
			ProtobufFieldName:  "sale",
			ProtobufType:       "int32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "gmv",
			Comment:            `成交GMV（分）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Gmv",
			GoFieldType:        "int32",
			JSONFieldName:      "gmv",
			ProtobufFieldName:  "gmv",
			ProtobufType:       "int32",
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
			Name:               "time",
			Comment:            `统计时间 如统计的21年7月2号的数据`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Time",
			GoFieldType:        "int32",
			JSONFieldName:      "time",
			ProtobufFieldName:  "time",
			ProtobufType:       "int32",
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
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
			ProtobufPos:        39,
		},

		&ColumnInfo{
			Index:              39,
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
			ProtobufPos:        40,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *StatisticsInfo) TableName() string {
	return "statistics_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *StatisticsInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *StatisticsInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *StatisticsInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *StatisticsInfo) TableInfo() *TableInfo {
	return statistics_infoTableInfo
}
