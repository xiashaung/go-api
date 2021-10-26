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


CREATE TABLE `product_dy` (
  `product_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `yx_product_id` bigint(20) NOT NULL COMMENT '央选商品id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '抖音用户ID',
  `cate_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '类目id，类目类型为抖音',
  `brand_id` int(11) NOT NULL DEFAULT '0' COMMENT '品牌id',
  `spec_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '规格id',
  `name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品名称',
  `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '主图 第一张',
  `pic` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品主图',
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '商品详情',
  `market_price` int(10) NOT NULL COMMENT '划线价，单位分',
  `discount_price` int(10) NOT NULL COMMENT '售价，单位分',
  `settlement_price` int(10) NOT NULL DEFAULT '0' COMMENT '结算价 分',
  `pay_type` int(255) NOT NULL DEFAULT '1' COMMENT '支持的支付方式：0货到付款 1在线支付 2两者都支持',
  `product_format` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci COMMENT '属性名称|属性值 json\n之间用|分隔, 多组之间用^分开',
  `recommend_remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商家推荐语',
  `extra` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '扩展字段',
  `presell_type` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '预售类型，1-全款预售，0-非预售，2-阶梯库存，默认0',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '电话号',
  `check_status` int(255) unsigned NOT NULL DEFAULT '1' COMMENT '商品审核状态：1未提审 2审核中 3审核通过 4审核驳回 5封禁',
  `status` int(255) NOT NULL DEFAULT '0' COMMENT '抖音商品状态 0上架 1下架 2已删除',
  `commission_rate` int(255) NOT NULL DEFAULT '0' COMMENT '商品佣金率 扩大1000倍存储',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建抖音商品的时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后更新时间',
  `category_detail` json DEFAULT NULL COMMENT '新类目的详情',
  `sales` int(11) NOT NULL DEFAULT '0' COMMENT '销量',
  `specs` text COMMENT '规格',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3498854384895429776 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "product_id": 48,    "yx_product_id": 97,    "user_id": 19,    "cate_id": 65,    "brand_id": 32,    "spec_id": 87,    "name": "LZDVEaWfRvXPbIdFexBjKmdMO",    "img": "bXUtSXNlveRRILfixjquSuDuh",    "pic": "jfqdYZyNYPQXNISIYHRoGPmit",    "description": "PCiBBHqDokhTPSQjLrlUBnCuL",    "market_price": 53,    "discount_price": 80,    "settlement_price": 95,    "pay_type": 62,    "product_format": "hTbtrdwFRKOIslEFhelmHRQdh",    "recommend_remark": "fNwHPXVWLFoFSDnSMXLsBfLMr",    "extra": "COOOgKcTcSaKeJsVFZZrKtDgo",    "presell_type": 65,    "mobile": "sMtqYpqDWJQnRiZpXnrKWkbvC",    "check_status": 93,    "status": 33,    "commission_rate": 95,    "add_time": 31,    "last_time": 2,    "category_detail": "dmZWHKDLbIBbJoHBINqBjRdIQ",    "sales": 76,    "specs": "wpUZaoDfGSVsrZrqEhGLFtHLv"}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 5] column is set for unsigned
[17] column is set for unsigned
[19] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned



*/

// ProductDy struct is a row record of the product_dy table in the yx database
type ProductDy struct {
	//[ 0] product_id                                     bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ProductID int64 `gorm:"primary_key;AUTO_INCREMENT;column:product_id;type:bigint;" json:"product_id"` // 商品id
	//[ 1] yx_product_id                                  bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	YxProductID int64 `gorm:"column:yx_product_id;type:bigint;" json:"yx_product_id"` // 央选商品id
	//[ 2] user_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UserID int32 `gorm:"column:user_id;type:int;default:0;" json:"user_id"` // 抖音用户ID
	//[ 3] cate_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateID uint32 `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"` // 类目id，类目类型为抖音
	//[ 4] brand_id                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	BrandID int32 `gorm:"column:brand_id;type:int;default:0;" json:"brand_id"` // 品牌id
	//[ 5] spec_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SpecID uint64 `gorm:"column:spec_id;type:ubigint;default:0;" json:"spec_id"` // 规格id
	//[ 6] name                                           varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	Name string `gorm:"column:name;type:varchar;size:60;" json:"name"` // 商品名称
	//[ 7] img                                            varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Img string `gorm:"column:img;type:varchar;size:255;" json:"img"` // 主图 第一张
	//[ 8] pic                                            text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Pic string `gorm:"column:pic;type:text;size:65535;" json:"pic"` // 商品主图
	//[ 9] description                                    text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Description sql.NullString `gorm:"column:description;type:text;size:65535;" json:"description"` // 商品详情
	//[10] market_price                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MarketPrice int32 `gorm:"column:market_price;type:int;" json:"market_price"` // 划线价，单位分
	//[11] discount_price                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DiscountPrice int32 `gorm:"column:discount_price;type:int;" json:"discount_price"` // 售价，单位分
	//[12] settlement_price                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SettlementPrice int32 `gorm:"column:settlement_price;type:int;default:0;" json:"settlement_price"` // 结算价 分
	//[13] pay_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	PayType int32 `gorm:"column:pay_type;type:int;default:1;" json:"pay_type"` // 支持的支付方式：0货到付款 1在线支付 2两者都支持
	//[14] product_format                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	ProductFormat sql.NullString `gorm:"column:product_format;type:text;size:65535;" json:"product_format"` // 属性名称|属性值 json\n之间用|分隔, 多组之间用^分开
	//[15] recommend_remark                               varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	RecommendRemark string `gorm:"column:recommend_remark;type:varchar;size:200;" json:"recommend_remark"` // 商家推荐语
	//[16] extra                                          text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Extra string `gorm:"column:extra;type:text;size:65535;" json:"extra"` // 扩展字段
	//[17] presell_type                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PresellType uint32 `gorm:"column:presell_type;type:uint;default:0;" json:"presell_type"` // 预售类型，1-全款预售，0-非预售，2-阶梯库存，默认0
	//[18] mobile                                         varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Mobile sql.NullString `gorm:"column:mobile;type:varchar;size:20;" json:"mobile"` // 电话号
	//[19] check_status                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	CheckStatus uint32 `gorm:"column:check_status;type:uint;default:1;" json:"check_status"` // 商品审核状态：1未提审 2审核中 3审核通过 4审核驳回 5封禁
	//[20] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Status int32 `gorm:"column:status;type:int;default:0;" json:"status"` // 抖音商品状态 0上架 1下架 2已删除
	//[21] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 商品佣金率 扩大1000倍存储
	//[22] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建抖音商品的时间
	//[23] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"` // 最后更新时间
	//[24] category_detail                                json                 null: true   primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	CategoryDetail sql.NullString `gorm:"column:category_detail;type:json;" json:"category_detail"` // 新类目的详情
	//[25] sales                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Sales int32 `gorm:"column:sales;type:int;default:0;" json:"sales"` // 销量
	//[26] specs                                          text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Specs sql.NullString `gorm:"column:specs;type:text;size:65535;" json:"specs"` // 规格

}

var product_dyTableInfo = &TableInfo{
	Name: "product_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "product_id",
			Comment:            `商品id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "yx_product_id",
			Comment:            `央选商品id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "user_id",
			Comment:            `抖音用户ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_id",
			Comment:            `类目id，类目类型为抖音`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "brand_id",
			Comment:            `品牌id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BrandID",
			GoFieldType:        "int32",
			JSONFieldName:      "brand_id",
			ProtobufFieldName:  "brand_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "spec_id",
			Comment:            `规格id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SpecID",
			GoFieldType:        "uint64",
			JSONFieldName:      "spec_id",
			ProtobufFieldName:  "spec_id",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "img",
			Comment:            `主图 第一张`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Img",
			GoFieldType:        "string",
			JSONFieldName:      "img",
			ProtobufFieldName:  "img",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "pic",
			Comment:            `商品主图`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Pic",
			GoFieldType:        "string",
			JSONFieldName:      "pic",
			ProtobufFieldName:  "pic",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "description",
			Comment:            `商品详情`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Description",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "market_price",
			Comment:            `划线价，单位分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MarketPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "market_price",
			ProtobufFieldName:  "market_price",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "discount_price",
			Comment:            `售价，单位分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DiscountPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "discount_price",
			ProtobufFieldName:  "discount_price",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "settlement_price",
			Comment:            `结算价 分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SettlementPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "settlement_price",
			ProtobufFieldName:  "settlement_price",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "pay_type",
			Comment:            `支持的支付方式：0货到付款 1在线支付 2两者都支持`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "product_format",
			Comment:            `属性名称|属性值 json\n之间用|分隔, 多组之间用^分开`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "ProductFormat",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "product_format",
			ProtobufFieldName:  "product_format",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "recommend_remark",
			Comment:            `商家推荐语`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "RecommendRemark",
			GoFieldType:        "string",
			JSONFieldName:      "recommend_remark",
			ProtobufFieldName:  "recommend_remark",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "extra",
			Comment:            `扩展字段`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Extra",
			GoFieldType:        "string",
			JSONFieldName:      "extra",
			ProtobufFieldName:  "extra",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "presell_type",
			Comment:            `预售类型，1-全款预售，0-非预售，2-阶梯库存，默认0`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PresellType",
			GoFieldType:        "uint32",
			JSONFieldName:      "presell_type",
			ProtobufFieldName:  "presell_type",
			ProtobufType:       "uint32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "mobile",
			Comment:            `电话号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Mobile",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mobile",
			ProtobufFieldName:  "mobile",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "check_status",
			Comment:            `商品审核状态：1未提审 2审核中 3审核通过 4审核驳回 5封禁`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CheckStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "check_status",
			ProtobufFieldName:  "check_status",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "status",
			Comment:            `抖音商品状态 0上架 1下架 2已删除`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "commission_rate",
			Comment:            `商品佣金率 扩大1000倍存储`,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "add_time",
			Comment:            `创建抖音商品的时间`,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "last_time",
			Comment:            `最后更新时间`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "category_detail",
			Comment:            `新类目的详情`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "CategoryDetail",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "category_detail",
			ProtobufFieldName:  "category_detail",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "sales",
			Comment:            `销量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Sales",
			GoFieldType:        "int32",
			JSONFieldName:      "sales",
			ProtobufFieldName:  "sales",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "specs",
			Comment:            `规格`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Specs",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "specs",
			ProtobufFieldName:  "specs",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProductDy) TableName() string {
	return "product_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductDy) TableInfo() *TableInfo {
	return product_dyTableInfo
}
