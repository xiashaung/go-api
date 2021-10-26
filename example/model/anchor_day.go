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


CREATE TABLE `anchor_day` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主播日销量信息表',
  `talent_id` int(11) DEFAULT NULL COMMENT '达人id',
  `anchor_id` int(11) DEFAULT NULL COMMENT '主播id',
  `product_num` int(11) DEFAULT NULL COMMENT '商品销售数量',
  `order_num` int(11) DEFAULT NULL COMMENT '订单数量',
  `fee_total` bigint(20) DEFAULT NULL COMMENT '日销售总金额，单位：分',
  `date` date DEFAULT NULL COMMENT '时间',
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_anchor_id_date` (`anchor_id`,`date`) USING BTREE,
  KEY `idx_talent_id` (`talent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=152 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 17,    "talent_id": 1,    "anchor_id": 43,    "product_num": 48,    "order_num": 83,    "fee_total": 62,    "date": "2298-10-10T20:42:08.717954385+08:00",    "add_time": 72,    "last_time": 27}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// AnchorDay struct is a row record of the anchor_day table in the yx database
type AnchorDay struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 主播日销量信息表
	//[ 1] talent_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TalentID sql.NullInt64 `gorm:"column:talent_id;type:int;" json:"talent_id"` // 达人id
	//[ 2] anchor_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AnchorID sql.NullInt64 `gorm:"column:anchor_id;type:int;" json:"anchor_id"` // 主播id
	//[ 3] product_num                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ProductNum sql.NullInt64 `gorm:"column:product_num;type:int;" json:"product_num"` // 商品销售数量
	//[ 4] order_num                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	OrderNum sql.NullInt64 `gorm:"column:order_num;type:int;" json:"order_num"` // 订单数量
	//[ 5] fee_total                                      bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	FeeTotal sql.NullInt64 `gorm:"column:fee_total;type:bigint;" json:"fee_total"` // 日销售总金额，单位：分
	//[ 6] date                                           date                 null: true   primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	Date time.Time `gorm:"column:date;type:date;" json:"date"` // 时间
	//[ 7] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 8] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var anchor_dayTableInfo = &TableInfo{
	Name: "anchor_day",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `主播日销量信息表`,
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
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_num",
			Comment:            `商品销售数量`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductNum",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "product_num",
			ProtobufFieldName:  "product_num",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "order_num",
			Comment:            `订单数量`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OrderNum",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "order_num",
			ProtobufFieldName:  "order_num",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "fee_total",
			Comment:            `日销售总金额，单位：分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "FeeTotal",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "fee_total",
			ProtobufFieldName:  "fee_total",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "date",
			Comment:            `时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "Date",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "add_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "last_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AnchorDay) TableName() string {
	return "anchor_day"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AnchorDay) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AnchorDay) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AnchorDay) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AnchorDay) TableInfo() *TableInfo {
	return anchor_dayTableInfo
}
