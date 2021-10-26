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


CREATE TABLE `activity_product` (
  `activity_product_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int(11) NOT NULL DEFAULT '0' COMMENT '商家id',
  `activity_id` int(11) NOT NULL DEFAULT '0' COMMENT '活动id',
  `platform_id` int(3) unsigned NOT NULL DEFAULT '2' COMMENT '第三方平台id',
  `product_id` int(11) NOT NULL DEFAULT '0' COMMENT '商品id',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '商品佣金率',
  `start_time` int(11) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `exposure_talent_num` int(11) NOT NULL DEFAULT '0' COMMENT '曝光达人数',
  `commerce_talent_num` int(11) NOT NULL DEFAULT '0' COMMENT '带货达人数',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 1 开启 0 关闭',
  `commission_fee` int(10) NOT NULL DEFAULT '0' COMMENT '佣金金额',
  `detail_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品团长链接',
  `activity_price` int(10) NOT NULL DEFAULT '0' COMMENT '活动价格',
  `activity_commission_fee` int(10) NOT NULL DEFAULT '0' COMMENT '预估活动佣金金额',
  `reason` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '审核原因',
  `product_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `sale_price` int(10) NOT NULL DEFAULT '0' COMMENT '售价',
  `third_product_id` bigint(10) NOT NULL DEFAULT '0' COMMENT '第三方产品id',
  `first_cate` int(10) NOT NULL DEFAULT '0' COMMENT '一级分类',
  `second_cate` int(10) NOT NULL DEFAULT '0' COMMENT '二级分类',
  `third_cate` int(10) NOT NULL DEFAULT '0' COMMENT '三级分类',
  `product_img` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品首图',
  `shop_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `shop_contact` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '店铺联系方式',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`activity_product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=833 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='活动商品表'

JSON Sample
-------------------------------------
{    "activity_product_id": 61,    "merchant_id": 97,    "activity_id": 67,    "platform_id": 36,    "product_id": 22,    "commission_rate": 23,    "start_time": 67,    "end_time": 55,    "exposure_talent_num": 43,    "commerce_talent_num": 92,    "status": 22,    "commission_fee": 17,    "detail_url": "nDpgRqxpeNNxhLHWAQDDDhDrb",    "activity_price": 54,    "activity_commission_fee": 13,    "reason": "rjPqKwxuAyuWOfifnVDfImVdw",    "product_name": "QBqgCOvOAyAmlOYwMDKVpBmqd",    "sale_price": 8,    "third_product_id": 67,    "first_cate": 8,    "second_cate": 38,    "third_cate": 57,    "product_img": "CVaUvGAcRxxqCgBOXWNxutRel",    "shop_name": "NubTiDFdKdaBEOLUgJAhbigtB",    "shop_contact": "dTtAFbgrqiAwmhIvtgGZQPqhm",    "last_time": 2,    "add_time": 35}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 3] column is set for unsigned
[25] column is set for unsigned
[26] column is set for unsigned



*/

// ActivityProduct struct is a row record of the activity_product table in the yx database
type ActivityProduct struct {
	//[ 0] activity_product_id                            uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ActivityProductID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:activity_product_id;type:uint;" json:"activity_product_id"`
	//[ 1] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"` // 商家id
	//[ 2] activity_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ActivityID int32 `gorm:"column:activity_id;type:int;default:0;" json:"activity_id"` // 活动id
	//[ 3] platform_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [2]
	PlatformID uint32 `gorm:"column:platform_id;type:uint;default:2;" json:"platform_id"` // 第三方平台id
	//[ 4] product_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductID int32 `gorm:"column:product_id;type:int;default:0;" json:"product_id"` // 商品id
	//[ 5] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 商品佣金率
	//[ 6] start_time                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	StartTime int32 `gorm:"column:start_time;type:int;default:0;" json:"start_time"` // 开始时间
	//[ 7] end_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EndTime int32 `gorm:"column:end_time;type:int;default:0;" json:"end_time"` // 结束时间
	//[ 8] exposure_talent_num                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExposureTalentNum int32 `gorm:"column:exposure_talent_num;type:int;default:0;" json:"exposure_talent_num"` // 曝光达人数
	//[ 9] commerce_talent_num                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommerceTalentNum int32 `gorm:"column:commerce_talent_num;type:int;default:0;" json:"commerce_talent_num"` // 带货达人数
	//[10] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Status int32 `gorm:"column:status;type:int;default:1;" json:"status"` // 状态 1 开启 0 关闭
	//[11] commission_fee                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionFee int32 `gorm:"column:commission_fee;type:int;default:0;" json:"commission_fee"` // 佣金金额
	//[12] detail_url                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	DetailURL string `gorm:"column:detail_url;type:varchar;size:255;" json:"detail_url"` // 商品团长链接
	//[13] activity_price                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ActivityPrice int32 `gorm:"column:activity_price;type:int;default:0;" json:"activity_price"` // 活动价格
	//[14] activity_commission_fee                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ActivityCommissionFee int32 `gorm:"column:activity_commission_fee;type:int;default:0;" json:"activity_commission_fee"` // 预估活动佣金金额
	//[15] reason                                         varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Reason string `gorm:"column:reason;type:varchar;size:255;" json:"reason"` // 审核原因
	//[16] product_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:255;" json:"product_name"` // 商品名称
	//[17] sale_price                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SalePrice int32 `gorm:"column:sale_price;type:int;default:0;" json:"sale_price"` // 售价
	//[18] third_product_id                               bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ThirdProductID int64 `gorm:"column:third_product_id;type:bigint;default:0;" json:"third_product_id"` // 第三方产品id
	//[19] first_cate                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	FirstCate int32 `gorm:"column:first_cate;type:int;default:0;" json:"first_cate"` // 一级分类
	//[20] second_cate                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SecondCate int32 `gorm:"column:second_cate;type:int;default:0;" json:"second_cate"` // 二级分类
	//[21] third_cate                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ThirdCate int32 `gorm:"column:third_cate;type:int;default:0;" json:"third_cate"` // 三级分类
	//[22] product_img                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ProductImg string `gorm:"column:product_img;type:varchar;size:255;" json:"product_img"` // 商品首图
	//[23] shop_name                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ShopName string `gorm:"column:shop_name;type:varchar;size:100;" json:"shop_name"` // 店铺名称
	//[24] shop_contact                                   varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ShopContact string `gorm:"column:shop_contact;type:varchar;size:100;" json:"shop_contact"` // 店铺联系方式
	//[25] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
	//[26] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"`
}

var activity_productTableInfo = &TableInfo{
	Name: "activity_product",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "activity_product_id",
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
			GoFieldName:        "ActivityProductID",
			GoFieldType:        "uint32",
			JSONFieldName:      "activity_product_id",
			ProtobufFieldName:  "activity_product_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "merchant_id",
			Comment:            `商家id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "activity_id",
			Comment:            `活动id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "platform_id",
			Comment:            `第三方平台id`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "product_id",
			Comment:            `商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "int32",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "commission_rate",
			Comment:            `商品佣金率`,
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
			Name:               "start_time",
			Comment:            `开始时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "int32",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "end_time",
			Comment:            `结束时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "int32",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "exposure_talent_num",
			Comment:            `曝光达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExposureTalentNum",
			GoFieldType:        "int32",
			JSONFieldName:      "exposure_talent_num",
			ProtobufFieldName:  "exposure_talent_num",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "commerce_talent_num",
			Comment:            `带货达人数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommerceTalentNum",
			GoFieldType:        "int32",
			JSONFieldName:      "commerce_talent_num",
			ProtobufFieldName:  "commerce_talent_num",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "status",
			Comment:            `状态 1 开启 0 关闭`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "commission_fee",
			Comment:            `佣金金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CommissionFee",
			GoFieldType:        "int32",
			JSONFieldName:      "commission_fee",
			ProtobufFieldName:  "commission_fee",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "detail_url",
			Comment:            `商品团长链接`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "DetailURL",
			GoFieldType:        "string",
			JSONFieldName:      "detail_url",
			ProtobufFieldName:  "detail_url",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "activity_price",
			Comment:            `活动价格`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ActivityPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "activity_price",
			ProtobufFieldName:  "activity_price",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "activity_commission_fee",
			Comment:            `预估活动佣金金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ActivityCommissionFee",
			GoFieldType:        "int32",
			JSONFieldName:      "activity_commission_fee",
			ProtobufFieldName:  "activity_commission_fee",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "reason",
			Comment:            `审核原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Reason",
			GoFieldType:        "string",
			JSONFieldName:      "reason",
			ProtobufFieldName:  "reason",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "sale_price",
			Comment:            `售价`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SalePrice",
			GoFieldType:        "int32",
			JSONFieldName:      "sale_price",
			ProtobufFieldName:  "sale_price",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "third_product_id",
			Comment:            `第三方产品id`,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "first_cate",
			Comment:            `一级分类`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FirstCate",
			GoFieldType:        "int32",
			JSONFieldName:      "first_cate",
			ProtobufFieldName:  "first_cate",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "second_cate",
			Comment:            `二级分类`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SecondCate",
			GoFieldType:        "int32",
			JSONFieldName:      "second_cate",
			ProtobufFieldName:  "second_cate",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "third_cate",
			Comment:            `三级分类`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ThirdCate",
			GoFieldType:        "int32",
			JSONFieldName:      "third_cate",
			ProtobufFieldName:  "third_cate",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "product_img",
			Comment:            `商品首图`,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "shop_contact",
			Comment:            `店铺联系方式`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ShopContact",
			GoFieldType:        "string",
			JSONFieldName:      "shop_contact",
			ProtobufFieldName:  "shop_contact",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "add_time",
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
			GoFieldName:        "AddTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "uint32",
			ProtobufPos:        27,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ActivityProduct) TableName() string {
	return "activity_product"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ActivityProduct) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ActivityProduct) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ActivityProduct) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ActivityProduct) TableInfo() *TableInfo {
	return activity_productTableInfo
}
