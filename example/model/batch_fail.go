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


CREATE TABLE `batch_fail` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `merchant_id` int(11) NOT NULL DEFAULT '0',
  `batch_sn` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '处理批次号',
  `sample_sn` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `shipping_conpany` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递公司',
  `shipping_sn` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '快递单号',
  `fail_reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '失败原因',
  `add_time` int(11) DEFAULT '0',
  `last_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=448 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='批量发货错误信息表'

JSON Sample
-------------------------------------
{    "id": 56,    "merchant_id": 11,    "batch_sn": "DGjMOvogvLvEhZJeTxywGJXjd",    "sample_sn": "qnnZvMbLNVZiBCZLPBogDPQKW",    "shipping_conpany": "puigTAVXwneJcjwSFUCFcCRiS",    "shipping_sn": "ehCMWaFNNkySjcquHrgyfdbBm",    "fail_reason": "yMWhjdTKGhpUAJFRrrQSghNgG",    "add_time": 43,    "last_time": 61}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// BatchFail struct is a row record of the batch_fail table in the yx database
type BatchFail struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"`
	//[ 2] batch_sn                                       varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	BatchSn string `gorm:"column:batch_sn;type:varchar;size:60;" json:"batch_sn"` // 处理批次号
	//[ 3] sample_sn                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SampleSn string `gorm:"column:sample_sn;type:varchar;size:255;" json:"sample_sn"` // 订单号
	//[ 4] shipping_conpany                               varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	ShippingConpany string `gorm:"column:shipping_conpany;type:varchar;size:50;" json:"shipping_conpany"` // 快递公司
	//[ 5] shipping_sn                                    varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ShippingSn string `gorm:"column:shipping_sn;type:varchar;size:100;" json:"shipping_sn"` // 快递单号
	//[ 6] fail_reason                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	FailReason string `gorm:"column:fail_reason;type:varchar;size:255;" json:"fail_reason"` // 失败原因
	//[ 7] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 8] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var batch_failTableInfo = &TableInfo{
	Name: "batch_fail",
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
			Name:               "merchant_id",
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
			GoFieldName:        "MerchantID",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "batch_sn",
			Comment:            `处理批次号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "BatchSn",
			GoFieldType:        "string",
			JSONFieldName:      "batch_sn",
			ProtobufFieldName:  "batch_sn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "sample_sn",
			Comment:            `订单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SampleSn",
			GoFieldType:        "string",
			JSONFieldName:      "sample_sn",
			ProtobufFieldName:  "sample_sn",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "shipping_conpany",
			Comment:            `快递公司`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "ShippingConpany",
			GoFieldType:        "string",
			JSONFieldName:      "shipping_conpany",
			ProtobufFieldName:  "shipping_conpany",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "shipping_sn",
			Comment:            `快递单号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ShippingSn",
			GoFieldType:        "string",
			JSONFieldName:      "shipping_sn",
			ProtobufFieldName:  "shipping_sn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "fail_reason",
			Comment:            `失败原因`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "FailReason",
			GoFieldType:        "string",
			JSONFieldName:      "fail_reason",
			ProtobufFieldName:  "fail_reason",
			ProtobufType:       "string",
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
func (b *BatchFail) TableName() string {
	return "batch_fail"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BatchFail) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BatchFail) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BatchFail) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BatchFail) TableInfo() *TableInfo {
	return batch_failTableInfo
}
