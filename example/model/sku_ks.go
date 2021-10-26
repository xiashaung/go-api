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


CREATE TABLE `sku_ks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快手sku信息表',
  `rel_sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外部商品 Sku id，仅供记录外部商品sku和快手商品sku对应关系',
  `yx_sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '央选sku id',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快手商品id',
  `image_url` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'Sku 图片url',
  `sku_stock` int(11) NOT NULL DEFAULT '0' COMMENT 'Sku 库存',
  `sku_sale_price` int(10) NOT NULL DEFAULT '0' COMMENT 'Sku 价格，单位为分',
  `volume` int(11) NOT NULL DEFAULT '0' COMMENT '已售',
  `is_valid` int(11) NOT NULL DEFAULT '0' COMMENT '是否有效',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间 时间戳 （毫秒）',
  `update_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '更新时间 时间戳 （毫秒）',
  `specification` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '规格',
  `sku_nick` varchar(255) NOT NULL DEFAULT '' COMMENT 'skuNick',
  `sku_prop` json NOT NULL COMMENT 'sku属性信息',
  `add_time` int(11) NOT NULL DEFAULT '0' COMMENT 'sku创建时间',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_sku_id` (`sku_id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 65,    "sku_id": 36,    "rel_sku_id": 92,    "yx_sku_id": 16,    "product_id": 58,    "image_url": "IynIkgBSfgdKyIDNWqsHgZdyq",    "sku_stock": 81,    "sku_sale_price": 96,    "volume": 11,    "is_valid": 11,    "create_time": 90,    "update_time": 67,    "specification": "RxJmDTYqNaocRyLjLwagXYvxK",    "sku_nick": "lOdNxDsSKRBTPSMwAeWATRAbB",    "sku_prop": "CpoGMFAcJcjuyvhncXFVQmwni",    "add_time": 20,    "last_time": 54}



*/

// SkuKs struct is a row record of the sku_ks table in the yx database
type SkuKs struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 主键
	//[ 1] sku_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SkuID int64 `gorm:"column:sku_id;type:bigint;default:0;" json:"sku_id"` // 快手sku信息表
	//[ 2] rel_sku_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	RelSkuID int64 `gorm:"column:rel_sku_id;type:bigint;default:0;" json:"rel_sku_id"` // 外部商品 Sku id，仅供记录外部商品sku和快手商品sku对应关系
	//[ 3] yx_sku_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	YxSkuID int64 `gorm:"column:yx_sku_id;type:bigint;default:0;" json:"yx_sku_id"` // 央选sku id
	//[ 4] product_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ProductID int64 `gorm:"column:product_id;type:bigint;default:0;" json:"product_id"` // 快手商品id
	//[ 5] image_url                                      varchar(125)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 125     default: []
	ImageURL string `gorm:"column:image_url;type:varchar;size:125;" json:"image_url"` // Sku 图片url
	//[ 6] sku_stock                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SkuStock int32 `gorm:"column:sku_stock;type:int;default:0;" json:"sku_stock"` // Sku 库存
	//[ 7] sku_sale_price                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SkuSalePrice int32 `gorm:"column:sku_sale_price;type:int;default:0;" json:"sku_sale_price"` // Sku 价格，单位为分
	//[ 8] volume                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Volume int32 `gorm:"column:volume;type:int;default:0;" json:"volume"` // 已售
	//[ 9] is_valid                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsValid int32 `gorm:"column:is_valid;type:int;default:0;" json:"is_valid"` // 是否有效
	//[10] create_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	CreateTime int64 `gorm:"column:create_time;type:bigint;default:0;" json:"create_time"` // 创建时间 时间戳 （毫秒）
	//[11] update_time                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	UpdateTime int64 `gorm:"column:update_time;type:bigint;default:0;" json:"update_time"` // 更新时间 时间戳 （毫秒）
	//[12] specification                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Specification string `gorm:"column:specification;type:varchar;size:255;" json:"specification"` // 规格
	//[13] sku_nick                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SkuNick string `gorm:"column:sku_nick;type:varchar;size:255;" json:"sku_nick"` // skuNick
	//[14] sku_prop                                       json                 null: false  primary: false  isArray: false  auto: false  col: json            len: -1      default: []
	SkuProp string `gorm:"column:sku_prop;type:json;" json:"sku_prop"` // sku属性信息
	//[15] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // sku创建时间
	//[16] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 修改时间

}

var sku_ksTableInfo = &TableInfo{
	Name: "sku_ks",
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
			Name:               "sku_id",
			Comment:            `快手sku信息表`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "rel_sku_id",
			Comment:            `外部商品 Sku id，仅供记录外部商品sku和快手商品sku对应关系`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "yx_sku_id",
			Comment:            `央选sku id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "YxSkuID",
			GoFieldType:        "int64",
			JSONFieldName:      "yx_sku_id",
			ProtobufFieldName:  "yx_sku_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "product_id",
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
			GoFieldName:        "ProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "image_url",
			Comment:            `Sku 图片url`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(125)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       125,
			GoFieldName:        "ImageURL",
			GoFieldType:        "string",
			JSONFieldName:      "image_url",
			ProtobufFieldName:  "image_url",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "sku_stock",
			Comment:            `Sku 库存`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SkuStock",
			GoFieldType:        "int32",
			JSONFieldName:      "sku_stock",
			ProtobufFieldName:  "sku_stock",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sku_sale_price",
			Comment:            `Sku 价格，单位为分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SkuSalePrice",
			GoFieldType:        "int32",
			JSONFieldName:      "sku_sale_price",
			ProtobufFieldName:  "sku_sale_price",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "volume",
			Comment:            `已售`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Volume",
			GoFieldType:        "int32",
			JSONFieldName:      "volume",
			ProtobufFieldName:  "volume",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "is_valid",
			Comment:            `是否有效`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsValid",
			GoFieldType:        "int32",
			JSONFieldName:      "is_valid",
			ProtobufFieldName:  "is_valid",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "create_time",
			Comment:            `创建时间 时间戳 （毫秒）`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "update_time",
			Comment:            `更新时间 时间戳 （毫秒）`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "specification",
			Comment:            `规格`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Specification",
			GoFieldType:        "string",
			JSONFieldName:      "specification",
			ProtobufFieldName:  "specification",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "sku_nick",
			Comment:            `skuNick`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "sku_prop",
			Comment:            `sku属性信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "json",
			DatabaseTypePretty: "json",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "json",
			ColumnLength:       -1,
			GoFieldName:        "SkuProp",
			GoFieldType:        "string",
			JSONFieldName:      "sku_prop",
			ProtobufFieldName:  "sku_prop",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "add_time",
			Comment:            `sku创建时间`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SkuKs) TableName() string {
	return "sku_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SkuKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SkuKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SkuKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SkuKs) TableInfo() *TableInfo {
	return sku_ksTableInfo
}
