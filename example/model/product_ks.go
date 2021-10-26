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


CREATE TABLE `product_ks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '快手商品信息表',
  `yx_product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '央选商品id',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '快手用户id',
  `rel_item_id` int(11) NOT NULL DEFAULT '0' COMMENT '外部商品id，仅供记录外部商品和快手商品对应关系',
  `title` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品标题',
  `details` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '商品描述',
  `main_image_urls` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品主图链接',
  `link_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品链接',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '叶子类目id',
  `category_list` json NOT NULL COMMENT '商品所属类目',
  `price` int(10) NOT NULL DEFAULT '0' COMMENT '商品价格 单位:分;',
  `volume` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '销量',
  `detail_image_urls` varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品图片URL',
  `show_image_urls` varchar(5000) NOT NULL DEFAULT '' COMMENT '详情展示图',
  `express_fee` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '邮费（单位：分）',
  `service_rule` json NOT NULL COMMENT '服务规则',
  `audit_status` smallint(6) NOT NULL DEFAULT '0' COMMENT '审核状态 0待审核 1审核待修改 2审核通过 3审核拒绝',
  `audit_reason` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '审核原因',
  `shelf_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '上下架状态 0-下架 1-上架',
  `shelf_reason` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '上下架原因',
  `app_key` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'appkey',
  `from_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '自建小店商品源标识 0 未知 1 快手自营 2 服务商通过OpenApi添加',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '佣金率 1000倍数',
  `item_status` int(10) NOT NULL DEFAULT '0' COMMENT '商品状态 0未知 1正常 2主播删除 3系统删除',
  `express_template_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '运费模板',
  `item_prop_values` json NOT NULL,
  `purchase_limit` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否限购',
  `limit_count` int(11) NOT NULL DEFAULT '0' COMMENT '限购数量',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间(毫秒)',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间(毫秒)',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(10) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `up_commission_rate_time` int(10) NOT NULL DEFAULT '0' COMMENT '佣金最近修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_yx_pid` (`yx_product_id`) USING BTREE,
  KEY `idx_pid` (`product_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='快手商品信息表'

JSON Sample
-------------------------------------
{    "id": 46,    "product_id": 69,    "yx_product_id": 29,    "user_id": 86,    "rel_item_id": 88,    "title": "XXbKbDdMFZhITUxXCBsHTfdye",    "details": "tEoVQFQGaEMSrbdmmmcnvDpBc",    "main_image_urls": "SWFiHZJXEavNgswASomsQljHD",    "link_url": "PqpQUnZUUNIrvtyHTEDJCUEPn",    "category_id": 95,    "category_list": "FBNSpQqcHXQoVOZmxWkJqliOg",    "price": 64,    "volume": 38,    "detail_image_urls": "QugPkYtBXbvRAqSZXLPxXrFih",    "show_image_urls": "joZyFhwcpefLHErIqYgKRhYeK",    "express_fee": 86,    "service_rule": "OWuXjKsZlacaNuGKXvLMpbijl",    "audit_status": 10,    "audit_reason": "wmSoqqtCgFvNXMEjtrxFcwfTR",    "shelf_status": 65,    "shelf_reason": "HNHQtiuSfDIMCOpKIhPFflBPA",    "app_key": "VuecGAJqHLbZLhWKdJVXDdoBV",    "from_type": 25,    "commission_rate": 10,    "item_status": 58,    "express_template_id": 38,    "item_prop_values": "ofxcFPaPjRjBULWOXNBxBHCIL",    "purchase_limit": 11,    "limit_count": 59,    "create_time": 12,    "update_time": 36,    "add_time": 94,    "last_time": 12,    "up_commission_rate_time": 41}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[12] column is set for unsigned
[15] column is set for unsigned
[19] column is set for unsigned
[22] column is set for unsigned



*/

// ProductKs struct is a row record of the product_ks table in the yx database
type ProductKs struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 主键
	//[ 1] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 快手商品信息表
	//[ 2] yx_product_id                                  ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	YxProductID uint64 `gorm:"column:yx_product_id;type:ubigint;default:0;" json:"yx_product_id"` // 央选商品id
	//[ 3] user_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	UserID uint32 `gorm:"column:user_id;type:uint;default:0;" json:"user_id"` // 快手用户id
	//[ 4] rel_item_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RelItemID int32 `gorm:"column:rel_item_id;type:int;default:0;" json:"rel_item_id"` // 外部商品id，仅供记录外部商品和快手商品对应关系
	//[ 5] title                                          varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	Title string `gorm:"column:title;type:varchar;size:125;" json:"title"` // 商品标题
	//[ 6] details                                        text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	Details string `gorm:"column:details;type:text;size:65535;" json:"details"` // 商品描述
	//[ 7] main_image_urls                                varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	MainImageUrls string `gorm:"column:main_image_urls;type:varchar;size:255;" json:"main_image_urls"` // 商品主图链接
	//[ 8] link_url                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	LinkURL string `gorm:"column:link_url;type:varchar;size:255;" json:"link_url"` // 商品链接
	//[ 9] category_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CategoryID int32 `gorm:"column:category_id;type:int;default:0;" json:"category_id"` // 叶子类目id
	//[10] category_list                                  json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	CategoryList string `gorm:"column:category_list;type:json;" json:"category_list"` // 商品所属类目
	//[11] price                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Price int32 `gorm:"column:price;type:int;default:0;" json:"price"` // 商品价格 单位:分;
	//[12] volume                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Volume uint32 `gorm:"column:volume;type:uint;default:0;" json:"volume"` // 销量
	//[13] detail_image_urls                              varchar(5000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 5000    default: []
	DetailImageUrls string `gorm:"column:detail_image_urls;type:varchar;size:5000;" json:"detail_image_urls"` // 商品图片URL
	//[14] show_image_urls                                varchar(5000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 5000    default: []
	ShowImageUrls string `gorm:"column:show_image_urls;type:varchar;size:5000;" json:"show_image_urls"` // 详情展示图
	//[15] express_fee                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ExpressFee uint32 `gorm:"column:express_fee;type:uint;default:0;" json:"express_fee"` // 邮费（单位：分）
	//[16] service_rule                                   json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	ServiceRule string `gorm:"column:service_rule;type:json;" json:"service_rule"` // 服务规则
	//[17] audit_status                                   smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	AuditStatus int32 `gorm:"column:audit_status;type:smallint;default:0;" json:"audit_status"` // 审核状态 0待审核 1审核待修改 2审核通过 3审核拒绝
	//[18] audit_reason                                   varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	AuditReason string `gorm:"column:audit_reason;type:varchar;size:200;" json:"audit_reason"` // 审核原因
	//[19] shelf_status                                   usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	ShelfStatus uint32 `gorm:"column:shelf_status;type:usmallint;default:0;" json:"shelf_status"` // 上下架状态 0-下架 1-上架
	//[20] shelf_reason                                   varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	ShelfReason string `gorm:"column:shelf_reason;type:varchar;size:200;" json:"shelf_reason"` // 上下架原因
	//[21] app_key                                        varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	AppKey string `gorm:"column:app_key;type:varchar;size:125;" json:"app_key"` // appkey
	//[22] from_type                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	FromType uint32 `gorm:"column:from_type;type:uint;default:0;" json:"from_type"` // 自建小店商品源标识 0 未知 1 快手自营 2 服务商通过OpenApi添加
	//[23] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 佣金率 1000倍数
	//[24] item_status                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ItemStatus int32 `gorm:"column:item_status;type:int;default:0;" json:"item_status"` // 商品状态 0未知 1正常 2主播删除 3系统删除
	//[25] express_template_id                            bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ExpressTemplateID int64 `gorm:"column:express_template_id;type:bigint;default:0;" json:"express_template_id"` // 运费模板
	//[26] item_prop_values                               json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	ItemPropValues string `gorm:"column:item_prop_values;type:json;" json:"item_prop_values"`
	//[27] purchase_limit                                 smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	PurchaseLimit int32 `gorm:"column:purchase_limit;type:smallint;default:0;" json:"purchase_limit"` // 是否限购
	//[28] limit_count                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LimitCount int32 `gorm:"column:limit_count;type:int;default:0;" json:"limit_count"` // 限购数量
	//[29] create_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	CreateTime int64 `gorm:"column:create_time;type:bigint;default:0;" json:"create_time"` // 创建时间(毫秒)
	//[30] update_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	UpdateTime int64 `gorm:"column:update_time;type:bigint;default:0;" json:"update_time"` // 更新时间(毫秒)
	//[31] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 创建时间
	//[32] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 修改时间
	//[33] up_commission_rate_time                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UpCommissionRateTime int32 `gorm:"column:up_commission_rate_time;type:int;default:0;" json:"up_commission_rate_time"` // 佣金最近修改时间

}

var product_ksTableInfo = &TableInfo{
	Name: "product_ks",
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
			Name:               "product_id",
			Comment:            `快手商品信息表`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "yx_product_id",
			Comment:            `央选商品id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "user_id",
			Comment:            `快手用户id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "rel_item_id",
			Comment:            `外部商品id，仅供记录外部商品和快手商品对应关系`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RelItemID",
			GoFieldType:        "int32",
			JSONFieldName:      "rel_item_id",
			ProtobufFieldName:  "rel_item_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "title",
			Comment:            `商品标题`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "Title",
			GoFieldType:        "string",
			JSONFieldName:      "title",
			ProtobufFieldName:  "title",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "details",
			Comment:            `商品描述`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "Details",
			GoFieldType:        "string",
			JSONFieldName:      "details",
			ProtobufFieldName:  "details",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "main_image_urls",
			Comment:            `商品主图链接`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "MainImageUrls",
			GoFieldType:        "string",
			JSONFieldName:      "main_image_urls",
			ProtobufFieldName:  "main_image_urls",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "link_url",
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
			GoFieldName:        "LinkURL",
			GoFieldType:        "string",
			JSONFieldName:      "link_url",
			ProtobufFieldName:  "link_url",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "category_id",
			Comment:            `叶子类目id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CategoryID",
			GoFieldType:        "int32",
			JSONFieldName:      "category_id",
			ProtobufFieldName:  "category_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "category_list",
			Comment:            `商品所属类目`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "price",
			Comment:            `商品价格 单位:分;`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "volume",
			Comment:            `销量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Volume",
			GoFieldType:        "uint32",
			JSONFieldName:      "volume",
			ProtobufFieldName:  "volume",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "detail_image_urls",
			Comment:            `商品图片URL`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(5000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       5000,
			GoFieldName:        "DetailImageUrls",
			GoFieldType:        "string",
			JSONFieldName:      "detail_image_urls",
			ProtobufFieldName:  "detail_image_urls",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "show_image_urls",
			Comment:            `详情展示图`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(5000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       5000,
			GoFieldName:        "ShowImageUrls",
			GoFieldType:        "string",
			JSONFieldName:      "show_image_urls",
			ProtobufFieldName:  "show_image_urls",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "express_fee",
			Comment:            `邮费（单位：分）`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "express_fee",
			ProtobufFieldName:  "express_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "service_rule",
			Comment:            `服务规则`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "ServiceRule",
			GoFieldType:        "string",
			JSONFieldName:      "service_rule",
			ProtobufFieldName:  "service_rule",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "audit_status",
			Comment:            `审核状态 0待审核 1审核待修改 2审核通过 3审核拒绝`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "AuditStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "audit_status",
			ProtobufFieldName:  "audit_status",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "audit_reason",
			Comment:            `审核原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "AuditReason",
			GoFieldType:        "string",
			JSONFieldName:      "audit_reason",
			ProtobufFieldName:  "audit_reason",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "shelf_status",
			Comment:            `上下架状态 0-下架 1-上架`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "ShelfStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "shelf_status",
			ProtobufFieldName:  "shelf_status",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "shelf_reason",
			Comment:            `上下架原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "ShelfReason",
			GoFieldType:        "string",
			JSONFieldName:      "shelf_reason",
			ProtobufFieldName:  "shelf_reason",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "app_key",
			Comment:            `appkey`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "AppKey",
			GoFieldType:        "string",
			JSONFieldName:      "app_key",
			ProtobufFieldName:  "app_key",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "from_type",
			Comment:            `自建小店商品源标识 0 未知 1 快手自营 2 服务商通过OpenApi添加`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "FromType",
			GoFieldType:        "uint32",
			JSONFieldName:      "from_type",
			ProtobufFieldName:  "from_type",
			ProtobufType:       "uint32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "commission_rate",
			Comment:            `佣金率 1000倍数`,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "item_status",
			Comment:            `商品状态 0未知 1正常 2主播删除 3系统删除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ItemStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "item_status",
			ProtobufFieldName:  "item_status",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "express_template_id",
			Comment:            `运费模板`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ExpressTemplateID",
			GoFieldType:        "int64",
			JSONFieldName:      "express_template_id",
			ProtobufFieldName:  "express_template_id",
			ProtobufType:       "int64",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "item_prop_values",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "ItemPropValues",
			GoFieldType:        "string",
			JSONFieldName:      "item_prop_values",
			ProtobufFieldName:  "item_prop_values",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "purchase_limit",
			Comment:            `是否限购`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "PurchaseLimit",
			GoFieldType:        "int32",
			JSONFieldName:      "purchase_limit",
			ProtobufFieldName:  "purchase_limit",
			ProtobufType:       "int32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "limit_count",
			Comment:            `限购数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LimitCount",
			GoFieldType:        "int32",
			JSONFieldName:      "limit_count",
			ProtobufFieldName:  "limit_count",
			ProtobufType:       "int32",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "create_time",
			Comment:            `创建时间(毫秒)`,
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
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "update_time",
			Comment:            `更新时间(毫秒)`,
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
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "add_time",
			Comment:            `创建时间`,
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
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "last_time",
			Comment:            `修改时间`,
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
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "up_commission_rate_time",
			Comment:            `佣金最近修改时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UpCommissionRateTime",
			GoFieldType:        "int32",
			JSONFieldName:      "up_commission_rate_time",
			ProtobufFieldName:  "up_commission_rate_time",
			ProtobufType:       "int32",
			ProtobufPos:        34,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProductKs) TableName() string {
	return "product_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductKs) TableInfo() *TableInfo {
	return product_ksTableInfo
}
