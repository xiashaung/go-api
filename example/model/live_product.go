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


CREATE TABLE `live_product` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '直播商品关联表',
  `live_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '直播场次id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `status` int(1) NOT NULL DEFAULT '1' COMMENT '状态 0 取消  1 待确定 2 待直播 3 直播中 4 已经结束',
  `is_push_window` int(1) NOT NULL DEFAULT '0' COMMENT '是否推送上橱窗 0 未推送  1 推送成功  2 推送失败',
  `pool_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联的商品池主键id',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加选品池时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=120 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='直播场次商品关联表'

JSON Sample
-------------------------------------
{    "id": 70,    "live_id": 32,    "product_id": 18,    "status": 64,    "is_push_window": 62,    "pool_id": 25,    "add_time": 71,    "last_time": 32}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned



*/

// LiveProduct struct is a row record of the live_product table in the yx database
type LiveProduct struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 直播商品关联表
	//[ 1] live_id                                        ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	LiveID uint64 `gorm:"column:live_id;type:ubigint;default:0;" json:"live_id"` // 直播场次id
	//[ 2] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品id
	//[ 3] status                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Status int32 `gorm:"column:status;type:int;default:1;" json:"status"` // 状态 0 取消  1 待确定 2 待直播 3 直播中 4 已经结束
	//[ 4] is_push_window                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	IsPushWindow int32 `gorm:"column:is_push_window;type:int;default:0;" json:"is_push_window"` // 是否推送上橱窗 0 未推送  1 推送成功  2 推送失败
	//[ 5] pool_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PoolID int64 `gorm:"column:pool_id;type:bigint;default:0;" json:"pool_id"` // 关联的商品池主键id
	//[ 6] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加选品池时间
	//[ 7] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var live_productTableInfo = &TableInfo{
	Name: "live_product",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `直播商品关联表`,
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
			Name:               "live_id",
			Comment:            `直播场次id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "LiveID",
			GoFieldType:        "uint64",
			JSONFieldName:      "live_id",
			ProtobufFieldName:  "live_id",
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
			Name:               "status",
			Comment:            `状态 0 取消  1 待确定 2 待直播 3 直播中 4 已经结束`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "is_push_window",
			Comment:            `是否推送上橱窗 0 未推送  1 推送成功  2 推送失败`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "IsPushWindow",
			GoFieldType:        "int32",
			JSONFieldName:      "is_push_window",
			ProtobufFieldName:  "is_push_window",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "pool_id",
			Comment:            `关联的商品池主键id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PoolID",
			GoFieldType:        "int64",
			JSONFieldName:      "pool_id",
			ProtobufFieldName:  "pool_id",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "add_time",
			Comment:            `添加选品池时间`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LiveProduct) TableName() string {
	return "live_product"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LiveProduct) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LiveProduct) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LiveProduct) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LiveProduct) TableInfo() *TableInfo {
	return live_productTableInfo
}
