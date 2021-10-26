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


CREATE TABLE `sku_dy` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'skuid',
  `out_sku_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '外部的skuId\n外部的skuId\n外部的skuId',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '抖音商品id',
  `yx_sku_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'sku id',
  `spec_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '抖音商品规格id',
  `spec_detail_id1` bigint(255) unsigned NOT NULL DEFAULT '0' COMMENT '一级规格id',
  `spec_detail_id2` bigint(255) NOT NULL DEFAULT '0' COMMENT '二级规格id，缺失为0',
  `spec_detail_id3` bigint(255) NOT NULL DEFAULT '0' COMMENT '三级规格id，缺失为0，最多三级',
  `spec_detail_name1` varchar(200) NOT NULL DEFAULT '' COMMENT '第一级子规格名\n第一级子规格名\n',
  `spec_detail_name2` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第二级子规格名\n第二级子规格名\n',
  `spec_detail_name3` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三级子规格名\n第二级子规格名\n',
  `stock_num` int(255) NOT NULL DEFAULT '0' COMMENT '库存数量',
  `step_stock_num` int(255) NOT NULL DEFAULT '0' COMMENT '阶梯库存数量，只有预售有效',
  `price` int(10) NOT NULL DEFAULT '0' COMMENT '抖音原价 单位分',
  `settlement_price` int(10) DEFAULT NULL COMMENT '结算价格 (单位 分)',
  `code` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '商品code',
  `out_warehouse_id` int(11) DEFAULT NULL COMMENT '外部仓库ID，需要注意，如果设置了此参数，该sku类型会变成区域库存类，原来的全国库存数据会被覆盖',
  `supplier_id` varchar(50) NOT NULL DEFAULT '' COMMENT '供应商ID',
  `sku_type` int(255) unsigned DEFAULT '0' COMMENT '库存类型：0普通库存，1区域库存',
  `stock_map` varchar(255) DEFAULT NULL COMMENT '区域仓库存信息，out_warehouse_id与stock_num对应关系',
  `add_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建抖音sku时间',
  `last_time` int(11) unsigned NOT NULL DEFAULT '0',
  `prom_stock_num` int(11) NOT NULL DEFAULT '0' COMMENT '活动库存',
  `prom_step_stock_num` int(11) NOT NULL DEFAULT '0' COMMENT '活动阶梯库存',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_sku_id` (`sku_id`),
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=121 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 98,    "sku_id": 32,    "out_sku_id": 56,    "product_id": 26,    "yx_sku_id": 38,    "spec_id": 8,    "spec_detail_id_1": 1,    "spec_detail_id_2": 37,    "spec_detail_id_3": 26,    "spec_detail_name_1": "cGUDTrnLnNCUkeIYOxfDfuTCD",    "spec_detail_name_2": "issBYSesQMCKJTStkjgRATlBo",    "spec_detail_name_3": "acuKMoYJZFpxdginOiLLKseXL",    "stock_num": 70,    "step_stock_num": 10,    "price": 23,    "settlement_price": 77,    "code": "onvQAZlFWOInjvywYCvCvnBgx",    "out_warehouse_id": 24,    "supplier_id": "JOfKLOoyGPvjtBcZSPebsYJEM",    "sku_type": 23,    "stock_map": "yXJPSRkFTPoPjjKUsCsTEAHfa",    "add_time": 94,    "last_time": 59,    "prom_stock_num": 79,    "prom_step_stock_num": 6}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 4] column is set for unsigned
[ 6] column is set for unsigned
[19] column is set for unsigned
[21] column is set for unsigned
[22] column is set for unsigned



*/

// SkuDy struct is a row record of the sku_dy table in the yx database
type SkuDy struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 主键
	//[ 1] sku_id                                         bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SkuID int64 `gorm:"column:sku_id;type:bigint;default:0;" json:"sku_id"` // skuid
	//[ 2] out_sku_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	OutSkuID int64 `gorm:"column:out_sku_id;type:bigint;default:0;" json:"out_sku_id"` // 外部的skuId\n外部的skuId\n外部的skuId
	//[ 3] product_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ProductID int64 `gorm:"column:product_id;type:bigint;default:0;" json:"product_id"` // 抖音商品id
	//[ 4] yx_sku_id                                      ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	YxSkuID uint64 `gorm:"column:yx_sku_id;type:ubigint;default:0;" json:"yx_sku_id"` // sku id
	//[ 5] spec_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SpecID int64 `gorm:"column:spec_id;type:bigint;default:0;" json:"spec_id"` // 抖音商品规格id
	//[ 6] spec_detail_id1                                ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	SpecDetailID1 uint64 `gorm:"column:spec_detail_id1;type:ubigint;default:0;" json:"spec_detail_id_1"` // 一级规格id
	//[ 7] spec_detail_id2                                bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SpecDetailID2 int64 `gorm:"column:spec_detail_id2;type:bigint;default:0;" json:"spec_detail_id_2"` // 二级规格id，缺失为0
	//[ 8] spec_detail_id3                                bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	SpecDetailID3 int64 `gorm:"column:spec_detail_id3;type:bigint;default:0;" json:"spec_detail_id_3"` // 三级规格id，缺失为0，最多三级
	//[ 9] spec_detail_name1                              varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	SpecDetailName1 string `gorm:"column:spec_detail_name1;type:varchar;size:200;" json:"spec_detail_name_1"` // 第一级子规格名\n第一级子规格名\n
	//[10] spec_detail_name2                              varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	SpecDetailName2 string `gorm:"column:spec_detail_name2;type:varchar;size:200;" json:"spec_detail_name_2"` // 第二级子规格名\n第二级子规格名\n
	//[11] spec_detail_name3                              varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	SpecDetailName3 string `gorm:"column:spec_detail_name3;type:varchar;size:200;" json:"spec_detail_name_3"` // 第三级子规格名\n第二级子规格名\n
	//[12] stock_num                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	StockNum int32 `gorm:"column:stock_num;type:int;default:0;" json:"stock_num"` // 库存数量
	//[13] step_stock_num                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	StepStockNum int32 `gorm:"column:step_stock_num;type:int;default:0;" json:"step_stock_num"` // 阶梯库存数量，只有预售有效
	//[14] price                                          int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Price int32 `gorm:"column:price;type:int;default:0;" json:"price"` // 抖音原价 单位分
	//[15] settlement_price                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SettlementPrice sql.NullInt64 `gorm:"column:settlement_price;type:int;" json:"settlement_price"` // 结算价格 (单位 分)
	//[16] code                                           varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	Code sql.NullString `gorm:"column:code;type:varchar;size:60;" json:"code"` // 商品code
	//[17] out_warehouse_id                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	OutWarehouseID sql.NullInt64 `gorm:"column:out_warehouse_id;type:int;" json:"out_warehouse_id"` // 外部仓库ID，需要注意，如果设置了此参数，该sku类型会变成区域库存类，原来的全国库存数据会被覆盖
	//[18] supplier_id                                    varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	SupplierID string `gorm:"column:supplier_id;type:varchar;size:50;" json:"supplier_id"` // 供应商ID
	//[19] sku_type                                       uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SkuType sql.NullInt64 `gorm:"column:sku_type;type:uint;default:0;" json:"sku_type"` // 库存类型：0普通库存，1区域库存
	//[20] stock_map                                      varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	StockMap sql.NullString `gorm:"column:stock_map;type:varchar;size:255;" json:"stock_map"` // 区域仓库存信息，out_warehouse_id与stock_num对应关系
	//[21] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建抖音sku时间
	//[22] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
	//[23] prom_stock_num                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PromStockNum int32 `gorm:"column:prom_stock_num;type:int;default:0;" json:"prom_stock_num"` // 活动库存
	//[24] prom_step_stock_num                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PromStepStockNum int32 `gorm:"column:prom_step_stock_num;type:int;default:0;" json:"prom_step_stock_num"` // 活动阶梯库存

}

var sku_dyTableInfo = &TableInfo{
	Name: "sku_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主键`,
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
			Name:               "sku_id",
			Comment:            `skuid`,
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
			Name:               "out_sku_id",
			Comment:            `外部的skuId\n外部的skuId\n外部的skuId`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "OutSkuID",
			GoFieldType:        "int64",
			JSONFieldName:      "out_sku_id",
			ProtobufFieldName:  "out_sku_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_id",
			Comment:            `抖音商品id`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "yx_sku_id",
			Comment:            `sku id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "YxSkuID",
			GoFieldType:        "uint64",
			JSONFieldName:      "yx_sku_id",
			ProtobufFieldName:  "yx_sku_id",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "spec_id",
			Comment:            `抖音商品规格id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SpecID",
			GoFieldType:        "int64",
			JSONFieldName:      "spec_id",
			ProtobufFieldName:  "spec_id",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "spec_detail_id1",
			Comment:            `一级规格id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "SpecDetailID1",
			GoFieldType:        "uint64",
			JSONFieldName:      "spec_detail_id_1",
			ProtobufFieldName:  "spec_detail_id_1",
			ProtobufType:       "uint64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "spec_detail_id2",
			Comment:            `二级规格id，缺失为0`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SpecDetailID2",
			GoFieldType:        "int64",
			JSONFieldName:      "spec_detail_id_2",
			ProtobufFieldName:  "spec_detail_id_2",
			ProtobufType:       "int64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "spec_detail_id3",
			Comment:            `三级规格id，缺失为0，最多三级`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "SpecDetailID3",
			GoFieldType:        "int64",
			JSONFieldName:      "spec_detail_id_3",
			ProtobufFieldName:  "spec_detail_id_3",
			ProtobufType:       "int64",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "spec_detail_name1",
			Comment:            `第一级子规格名\n第一级子规格名\n`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "SpecDetailName1",
			GoFieldType:        "string",
			JSONFieldName:      "spec_detail_name_1",
			ProtobufFieldName:  "spec_detail_name_1",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "spec_detail_name2",
			Comment:            `第二级子规格名\n第二级子规格名\n`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "SpecDetailName2",
			GoFieldType:        "string",
			JSONFieldName:      "spec_detail_name_2",
			ProtobufFieldName:  "spec_detail_name_2",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "spec_detail_name3",
			Comment:            `第三级子规格名\n第二级子规格名\n`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "SpecDetailName3",
			GoFieldType:        "string",
			JSONFieldName:      "spec_detail_name_3",
			ProtobufFieldName:  "spec_detail_name_3",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "stock_num",
			Comment:            `库存数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StockNum",
			GoFieldType:        "int32",
			JSONFieldName:      "stock_num",
			ProtobufFieldName:  "stock_num",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "step_stock_num",
			Comment:            `阶梯库存数量，只有预售有效`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StepStockNum",
			GoFieldType:        "int32",
			JSONFieldName:      "step_stock_num",
			ProtobufFieldName:  "step_stock_num",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "price",
			Comment:            `抖音原价 单位分`,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "settlement_price",
			Comment:            `结算价格 (单位 分)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SettlementPrice",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "settlement_price",
			ProtobufFieldName:  "settlement_price",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "code",
			Comment:            `商品code`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "Code",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "code",
			ProtobufFieldName:  "code",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "out_warehouse_id",
			Comment:            `外部仓库ID，需要注意，如果设置了此参数，该sku类型会变成区域库存类，原来的全国库存数据会被覆盖`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OutWarehouseID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "out_warehouse_id",
			ProtobufFieldName:  "out_warehouse_id",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "supplier_id",
			Comment:            `供应商ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "SupplierID",
			GoFieldType:        "string",
			JSONFieldName:      "supplier_id",
			ProtobufFieldName:  "supplier_id",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "sku_type",
			Comment:            `库存类型：0普通库存，1区域库存`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SkuType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "sku_type",
			ProtobufFieldName:  "sku_type",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "stock_map",
			Comment:            `区域仓库存信息，out_warehouse_id与stock_num对应关系`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "StockMap",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "stock_map",
			ProtobufFieldName:  "stock_map",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "add_time",
			Comment:            `创建抖音sku时间`,
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
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "prom_stock_num",
			Comment:            `活动库存`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PromStockNum",
			GoFieldType:        "int32",
			JSONFieldName:      "prom_stock_num",
			ProtobufFieldName:  "prom_stock_num",
			ProtobufType:       "int32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "prom_step_stock_num",
			Comment:            `活动阶梯库存`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PromStepStockNum",
			GoFieldType:        "int32",
			JSONFieldName:      "prom_step_stock_num",
			ProtobufFieldName:  "prom_step_stock_num",
			ProtobufType:       "int32",
			ProtobufPos:        25,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SkuDy) TableName() string {
	return "sku_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SkuDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SkuDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SkuDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SkuDy) TableInfo() *TableInfo {
	return sku_dyTableInfo
}
