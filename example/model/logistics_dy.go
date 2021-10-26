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


CREATE TABLE `logistics_dy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '抖音物流信息表',
  `dy_pack_id` int(11) NOT NULL COMMENT '包裹ID',
  `dy_shipped_num` int(11) NOT NULL COMMENT '包裹中已发货商品数量',
  `dy_logistics_company` int(11) NOT NULL COMMENT '包裹的物流公司code码，如需要对应的数字ID',
  `dy_logistics_code` varchar(30) NOT NULL COMMENT '包裹的物流单号',
  `dy_logistics_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '包裹的发货时间',
  `dy_order_list` text COMMENT '包裹中已发货的子订单列表',
  `dy_order_id` varchar(100) DEFAULT NULL COMMENT '已发货子订单号',
  `dy_item_ids` varchar(100) DEFAULT NULL COMMENT '已发货子订单中，已发货的品ID列表',
  `dy_sku_num` int(11) DEFAULT NULL COMMENT '已发货子订单的sku数量',
  `add_time` int(11) NOT NULL COMMENT '抖音物流信息',
  `last_time` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 6,    "dy_pack_id": 87,    "dy_shipped_num": 95,    "dy_logistics_company": 68,    "dy_logistics_code": "mKETglhRaToymMJAcWsWuLbXc",    "dy_logistics_time": 38,    "dy_order_list": "oHPQjjUNOhKHwBBLnAwbDllSL",    "dy_order_id": "uImHGRIPNDZOUBtrwdjWbNOIw",    "dy_item_ids": "tDxhXXRqaOmFVCHmnkyqAoeXK",    "dy_sku_num": 25,    "add_time": 65,    "last_time": 22}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 5] column is set for unsigned



*/

// LogisticsDy struct is a row record of the logistics_dy table in the yx database
type LogisticsDy struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 抖音物流信息表
	//[ 1] dy_pack_id                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DyPackID int32 `gorm:"column:dy_pack_id;type:int;" json:"dy_pack_id"` // 包裹ID
	//[ 2] dy_shipped_num                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DyShippedNum int32 `gorm:"column:dy_shipped_num;type:int;" json:"dy_shipped_num"` // 包裹中已发货商品数量
	//[ 3] dy_logistics_company                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DyLogisticsCompany int32 `gorm:"column:dy_logistics_company;type:int;" json:"dy_logistics_company"` // 包裹的物流公司code码，如需要对应的数字ID
	//[ 4] dy_logistics_code                              varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	DyLogisticsCode string `gorm:"column:dy_logistics_code;type:varchar;size:30;" json:"dy_logistics_code"` // 包裹的物流单号
	//[ 5] dy_logistics_time                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DyLogisticsTime uint32 `gorm:"column:dy_logistics_time;type:uint;default:0;" json:"dy_logistics_time"` // 包裹的发货时间
	//[ 6] dy_order_list                                  text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	DyOrderList sql.NullString `gorm:"column:dy_order_list;type:text;size:65535;" json:"dy_order_list"` // 包裹中已发货的子订单列表
	//[ 7] dy_order_id                                    varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	DyOrderID sql.NullString `gorm:"column:dy_order_id;type:varchar;size:100;" json:"dy_order_id"` // 已发货子订单号
	//[ 8] dy_item_ids                                    varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	DyItemIds sql.NullString `gorm:"column:dy_item_ids;type:varchar;size:100;" json:"dy_item_ids"` // 已发货子订单中，已发货的品ID列表
	//[ 9] dy_sku_num                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DySkuNum sql.NullInt64 `gorm:"column:dy_sku_num;type:int;" json:"dy_sku_num"` // 已发货子订单的sku数量
	//[10] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime int32 `gorm:"column:add_time;type:int;" json:"add_time"` // 抖音物流信息
	//[11] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime int32 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var logistics_dyTableInfo = &TableInfo{
	Name: "logistics_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `抖音物流信息表`,
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
			Name:               "dy_pack_id",
			Comment:            `包裹ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DyPackID",
			GoFieldType:        "int32",
			JSONFieldName:      "dy_pack_id",
			ProtobufFieldName:  "dy_pack_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "dy_shipped_num",
			Comment:            `包裹中已发货商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DyShippedNum",
			GoFieldType:        "int32",
			JSONFieldName:      "dy_shipped_num",
			ProtobufFieldName:  "dy_shipped_num",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "dy_logistics_company",
			Comment:            `包裹的物流公司code码，如需要对应的数字ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DyLogisticsCompany",
			GoFieldType:        "int32",
			JSONFieldName:      "dy_logistics_company",
			ProtobufFieldName:  "dy_logistics_company",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "dy_logistics_code",
			Comment:            `包裹的物流单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "DyLogisticsCode",
			GoFieldType:        "string",
			JSONFieldName:      "dy_logistics_code",
			ProtobufFieldName:  "dy_logistics_code",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "dy_logistics_time",
			Comment:            `包裹的发货时间`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DyLogisticsTime",
			GoFieldType:        "uint32",
			JSONFieldName:      "dy_logistics_time",
			ProtobufFieldName:  "dy_logistics_time",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "dy_order_list",
			Comment:            `包裹中已发货的子订单列表`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "DyOrderList",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dy_order_list",
			ProtobufFieldName:  "dy_order_list",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "dy_order_id",
			Comment:            `已发货子订单号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "DyOrderID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dy_order_id",
			ProtobufFieldName:  "dy_order_id",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "dy_item_ids",
			Comment:            `已发货子订单中，已发货的品ID列表`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "DyItemIds",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dy_item_ids",
			ProtobufFieldName:  "dy_item_ids",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "dy_sku_num",
			Comment:            `已发货子订单的sku数量`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DySkuNum",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dy_sku_num",
			ProtobufFieldName:  "dy_sku_num",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "add_time",
			Comment:            `抖音物流信息`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LogisticsDy) TableName() string {
	return "logistics_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LogisticsDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LogisticsDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LogisticsDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LogisticsDy) TableInfo() *TableInfo {
	return logistics_dyTableInfo
}
