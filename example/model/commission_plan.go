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


CREATE TABLE `commission_plan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '佣金计划表',
  `plan_id` bigint(20) NOT NULL COMMENT '计划id',
  `product_id` bigint(20) NOT NULL COMMENT '商品ID',
  `shop_id` int(11) NOT NULL COMMENT '店铺ID',
  `merchant_id` bigint(20) NOT NULL COMMENT '商家ID',
  `promoter_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '推广者ID',
  `commission_rate` int(11) NOT NULL COMMENT '佣金率 * 1000',
  `platform_id` tinyint(2) NOT NULL DEFAULT '2' COMMENT '1:抖音2：快手',
  `plan_status` tinyint(2) NOT NULL DEFAULT '0' COMMENT '计划状态。1-开启；3-关闭',
  `plan_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '计划类型。1-普通计划；2-商品定向佣金;5-专属计划\n',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_mid_type` (`merchant_id`,`plan_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 88,    "plan_id": 45,    "product_id": 4,    "shop_id": 21,    "merchant_id": 40,    "promoter_id": 86,    "commission_rate": 24,    "platform_id": 9,    "plan_status": 55,    "plan_type": 37,    "add_time": 93,    "last_time": 52}



*/

// CommissionPlan struct is a row record of the commission_plan table in the yx database
type CommissionPlan struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 佣金计划表
	//[ 1] plan_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PlanID int64 `gorm:"column:plan_id;type:bigint;" json:"plan_id"` // 计划id
	//[ 2] product_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	ProductID int64 `gorm:"column:product_id;type:bigint;" json:"product_id"` // 商品ID
	//[ 3] shop_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ShopID int32 `gorm:"column:shop_id;type:int;" json:"shop_id"` // 店铺ID
	//[ 4] merchant_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	MerchantID int64 `gorm:"column:merchant_id;type:bigint;" json:"merchant_id"` // 商家ID
	//[ 5] promoter_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	PromoterID int64 `gorm:"column:promoter_id;type:bigint;default:0;" json:"promoter_id"` // 推广者ID
	//[ 6] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CommissionRate int32 `gorm:"column:commission_rate;type:int;" json:"commission_rate"` // 佣金率 * 1000
	//[ 7] platform_id                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [2]
	PlatformID int32 `gorm:"column:platform_id;type:tinyint;default:2;" json:"platform_id"` // 1:抖音2：快手
	//[ 8] plan_status                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	PlanStatus int32 `gorm:"column:plan_status;type:tinyint;default:0;" json:"plan_status"` // 计划状态。1-开启；3-关闭
	//[ 9] plan_type                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	PlanType int32 `gorm:"column:plan_type;type:tinyint;default:0;" json:"plan_type"` // 计划类型。1-普通计划；2-商品定向佣金;5-专属计划\n
	//[10] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[11] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var commission_planTableInfo = &TableInfo{
	Name: "commission_plan",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `佣金计划表`,
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
			Name:               "plan_id",
			Comment:            `计划id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PlanID",
			GoFieldType:        "int64",
			JSONFieldName:      "plan_id",
			ProtobufFieldName:  "plan_id",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "product_id",
			Comment:            `商品ID`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "shop_id",
			Comment:            `店铺ID`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "promoter_id",
			Comment:            `推广者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PromoterID",
			GoFieldType:        "int64",
			JSONFieldName:      "promoter_id",
			ProtobufFieldName:  "promoter_id",
			ProtobufType:       "int64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "commission_rate",
			Comment:            `佣金率 * 1000`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "platform_id",
			Comment:            `1:抖音2：快手`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "plan_status",
			Comment:            `计划状态。1-开启；3-关闭`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "PlanStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "plan_status",
			ProtobufFieldName:  "plan_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "plan_type",
			Comment:            `计划类型。1-普通计划；2-商品定向佣金;5-专属计划\n`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "PlanType",
			GoFieldType:        "int32",
			JSONFieldName:      "plan_type",
			ProtobufFieldName:  "plan_type",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
func (c *CommissionPlan) TableName() string {
	return "commission_plan"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CommissionPlan) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CommissionPlan) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CommissionPlan) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CommissionPlan) TableInfo() *TableInfo {
	return commission_planTableInfo
}
