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


CREATE TABLE `order_product` (
  `id` bigint(20) NOT NULL COMMENT '订单商品表',
  `order_sn` varchar(64) NOT NULL COMMENT '订单编号',
  `platform_id` smallint(6) NOT NULL DEFAULT '0' COMMENT '平台id',
  `talent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `anchor_id` int(11) NOT NULL COMMENT '商家id',
  `shop_id` int(11) NOT NULL COMMENT '店铺id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `product_name` varchar(100) NOT NULL COMMENT '商品名称',
  `product_price` int(11) NOT NULL COMMENT '商品价格',
  `product_num` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '商品数量',
  `commission_rate` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '佣金率',
  `pay_fee` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '支付金额',
  `status` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '状态',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '下单时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='订单商品表'

JSON Sample
-------------------------------------
{    "id": 2,    "order_sn": "XaIBWpmFasHBjLogGHErYquIp",    "platform_id": 57,    "talent_id": 9,    "anchor_id": 26,    "shop_id": 83,    "product_id": 50,    "product_name": "CrwsofXyUEHCuFARmrLsFlDXb",    "product_price": 72,    "product_num": 97,    "commission_rate": 26,    "pay_fee": 72,    "status": 31,    "add_time": 20,    "last_time": 54}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 6] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned



*/

// OrderProduct struct is a row record of the order_product table in the yx database
type OrderProduct struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;column:id;type:bigint;" json:"id"` // 订单商品表
	//[ 1] order_sn                                       varchar(64)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 64      default: []
	OrderSn string `gorm:"column:order_sn;type:varchar;size:64;" json:"order_sn"` // 订单编号
	//[ 2] platform_id                                    smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:smallint;default:0;" json:"platform_id"` // 平台id
	//[ 3] talent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TalentID uint32 `gorm:"column:talent_id;type:uint;default:0;" json:"talent_id"` // 达人id
	//[ 4] anchor_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AnchorID int32 `gorm:"column:anchor_id;type:int;" json:"anchor_id"` // 商家id
	//[ 5] shop_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ShopID int32 `gorm:"column:shop_id;type:int;" json:"shop_id"` // 店铺id
	//[ 6] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品id
	//[ 7] product_name                                   varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ProductName string `gorm:"column:product_name;type:varchar;size:100;" json:"product_name"` // 商品名称
	//[ 8] product_price                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ProductPrice int32 `gorm:"column:product_price;type:int;" json:"product_price"` // 商品价格
	//[ 9] product_num                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	ProductNum uint32 `gorm:"column:product_num;type:uint;default:1;" json:"product_num"` // 商品数量
	//[10] commission_rate                                uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CommissionRate uint32 `gorm:"column:commission_rate;type:uint;default:0;" json:"commission_rate"` // 佣金率
	//[11] pay_fee                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PayFee uint32 `gorm:"column:pay_fee;type:uint;default:0;" json:"pay_fee"` // 支付金额
	//[12] status                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Status uint32 `gorm:"column:status;type:uint;default:0;" json:"status"` // 状态
	//[13] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 下单时间
	//[14] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var order_productTableInfo = &TableInfo{
	Name: "order_product",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `订单商品表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
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
			Name:               "order_sn",
			Comment:            `订单编号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(64)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       64,
			GoFieldName:        "OrderSn",
			GoFieldType:        "string",
			JSONFieldName:      "order_sn",
			ProtobufFieldName:  "order_sn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "platform_id",
			Comment:            `平台id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "anchor_id",
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
			GoFieldName:        "AnchorID",
			GoFieldType:        "int32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "shop_id",
			Comment:            `店铺id`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "product_name",
			Comment:            `商品名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ProductName",
			GoFieldType:        "string",
			JSONFieldName:      "product_name",
			ProtobufFieldName:  "product_name",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "product_price",
			Comment:            `商品价格`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductPrice",
			GoFieldType:        "int32",
			JSONFieldName:      "product_price",
			ProtobufFieldName:  "product_price",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "product_num",
			Comment:            `商品数量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ProductNum",
			GoFieldType:        "uint32",
			JSONFieldName:      "product_num",
			ProtobufFieldName:  "product_num",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "commission_rate",
			Comment:            `佣金率`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "pay_fee",
			Comment:            `支付金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PayFee",
			GoFieldType:        "uint32",
			JSONFieldName:      "pay_fee",
			ProtobufFieldName:  "pay_fee",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "status",
			Comment:            `状态`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "uint32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "add_time",
			Comment:            `下单时间`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrderProduct) TableName() string {
	return "order_product"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderProduct) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderProduct) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderProduct) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrderProduct) TableInfo() *TableInfo {
	return order_productTableInfo
}
