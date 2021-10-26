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


CREATE TABLE `merchant_enroll_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商家报名Id',
  `merchant_id` bigint(20) NOT NULL COMMENT '商家id',
  `tender_id` bigint(20) NOT NULL COMMENT '招商计划ID',
  `product_num` int(11) NOT NULL COMMENT '报名商品数',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mt_id` (`merchant_id`,`tender_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 57,    "merchant_id": 54,    "tender_id": 37,    "product_num": 77,    "add_time": 73,    "last_time": 50}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// MerchantEnrollLog struct is a row record of the merchant_enroll_log table in the yx database
type MerchantEnrollLog struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 商家报名Id
	//[ 1] merchant_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	MerchantID int64 `gorm:"column:merchant_id;type:bigint;" json:"merchant_id"` // 商家id
	//[ 2] tender_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	TenderID int64 `gorm:"column:tender_id;type:bigint;" json:"tender_id"` // 招商计划ID
	//[ 3] product_num                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ProductNum int32 `gorm:"column:product_num;type:int;" json:"product_num"` // 报名商品数
	//[ 4] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 5] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var merchant_enroll_logTableInfo = &TableInfo{
	Name: "merchant_enroll_log",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `商家报名Id`,
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
			Name:               "merchant_id",
			Comment:            `商家id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int64",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "tender_id",
			Comment:            `招商计划ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "TenderID",
			GoFieldType:        "int64",
			JSONFieldName:      "tender_id",
			ProtobufFieldName:  "tender_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_num",
			Comment:            `报名商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ProductNum",
			GoFieldType:        "int32",
			JSONFieldName:      "product_num",
			ProtobufFieldName:  "product_num",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "add_time",
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
			GoFieldName:        "AddTime",
			GoFieldType:        "int32",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MerchantEnrollLog) TableName() string {
	return "merchant_enroll_log"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MerchantEnrollLog) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MerchantEnrollLog) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MerchantEnrollLog) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MerchantEnrollLog) TableInfo() *TableInfo {
	return merchant_enroll_logTableInfo
}
