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


CREATE TABLE `merchant_cate` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '商家类目表',
  `merchant_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '商家ID',
  `shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '店铺ID',
  `cate_id2` int(11) NOT NULL DEFAULT '0' COMMENT '2级类目',
  `cate_id3` int(11) NOT NULL DEFAULT '0' COMMENT '3级类目',
  `platform_id` tinyint(2) NOT NULL DEFAULT '0' COMMENT '平台1：抖音  2：快手',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_merchant_cate2` (`merchant_id`,`cate_id2`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 26,    "merchant_id": 91,    "shop_id": 31,    "cate_id_2": 48,    "cate_id_3": 81,    "platform_id": 27,    "add_time": 13,    "last_time": 14}



*/

// MerchantCate struct is a row record of the merchant_cate table in the yx database
type MerchantCate struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 商家类目表
	//[ 1] merchant_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	MerchantID int64 `gorm:"column:merchant_id;type:bigint;default:0;" json:"merchant_id"` // 商家ID
	//[ 2] shop_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ShopID int64 `gorm:"column:shop_id;type:bigint;default:0;" json:"shop_id"` // 店铺ID
	//[ 3] cate_id2                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateID2 int32 `gorm:"column:cate_id2;type:int;default:0;" json:"cate_id_2"` // 2级类目
	//[ 4] cate_id3                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateID3 int32 `gorm:"column:cate_id3;type:int;default:0;" json:"cate_id_3"` // 3级类目
	//[ 5] platform_id                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:tinyint;default:0;" json:"platform_id"` // 平台1：抖音  2：快手
	//[ 6] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 7] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var merchant_cateTableInfo = &TableInfo{
	Name: "merchant_cate",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `商家类目表`,
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
			Name:               "merchant_id",
			Comment:            `商家ID`,
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
			Name:               "shop_id",
			Comment:            `店铺ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ShopID",
			GoFieldType:        "int64",
			JSONFieldName:      "shop_id",
			ProtobufFieldName:  "shop_id",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_id2",
			Comment:            `2级类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateID2",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_id_2",
			ProtobufFieldName:  "cate_id_2",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "cate_id3",
			Comment:            `3级类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateID3",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_id_3",
			ProtobufFieldName:  "cate_id_3",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "platform_id",
			Comment:            `平台1：抖音  2：快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MerchantCate) TableName() string {
	return "merchant_cate"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MerchantCate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MerchantCate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MerchantCate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MerchantCate) TableInfo() *TableInfo {
	return merchant_cateTableInfo
}
