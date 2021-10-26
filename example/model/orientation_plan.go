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


CREATE TABLE `orientation_plan` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '达人定向计划表',
  `plan_id` bigint(20) NOT NULL COMMENT '快手佣金计划id',
  `merchant_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '央选商家id',
  `anchor_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主播ID',
  `commission_id` bigint(20) NOT NULL COMMENT '快手佣金ID',
  `promoter_id` bigint(20) NOT NULL COMMENT '快手推广人id',
  `third_product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '快手商品id',
  `product_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '央选商品id',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '快手佣金率 1000倍',
  `plan_type` tinyint(2) NOT NULL DEFAULT '0' COMMENT '计划类型。1-普通计划；2-商品定向佣金;5-专属计划\n',
  `plan_status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '定向计划状态。1-开启; 3-关闭 2:关闭中 投放状态',
  `platform_id` tinyint(2) NOT NULL DEFAULT '2' COMMENT '1:抖音 2快手',
  `activity_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '活动id',
  `add_time` int(11) NOT NULL DEFAULT '0',
  `last_time` int(11) NOT NULL DEFAULT '0',
  `start_time` int(11) NOT NULL DEFAULT '0' COMMENT '推广开始时间',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '推广结束时间',
  `new_cos_status` tinyint(2) NOT NULL DEFAULT '1' COMMENT '修改后的佣金状态1:已生效 2:待生效',
  `new_cos_rate` int(11) NOT NULL DEFAULT '0' COMMENT '修改后的佣金',
  `new_cos_rate_start_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改后的生效时间',
  `talent_delete_time` int(11) NOT NULL DEFAULT '0' COMMENT '达人移除时间',
  `is_create` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否重建计划&达人 0:不需要 1:需要',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_commission_id` (`commission_id`,`plan_id`,`platform_id`) USING BTREE,
  KEY `idx_plan_id` (`plan_id`),
  KEY `idx_merchant_id` (`merchant_id`,`anchor_id`) USING BTREE,
  KEY `idx_product_id_platform` (`product_id`,`platform_id`) USING BTREE,
  KEY `idx_talent_delete_time` (`talent_delete_time`),
  KEY `idx_new_cos_rate_time_status` (`new_cos_rate_start_time`,`new_cos_status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=704 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 5,    "plan_id": 15,    "merchant_id": 49,    "anchor_id": 55,    "commission_id": 35,    "promoter_id": 25,    "third_product_id": 87,    "product_id": 44,    "commission_rate": 65,    "plan_type": 18,    "plan_status": 13,    "platform_id": 47,    "activity_id": 79,    "add_time": 59,    "last_time": 79,    "start_time": 31,    "end_time": 83,    "new_cos_status": 98,    "new_cos_rate": 12,    "new_cos_rate_start_time": 42,    "talent_delete_time": 48,    "is_create": 16}


Comments
-------------------------------------
[12] column is set for unsigned



*/

// OrientationPlan struct is a row record of the orientation_plan table in the yx database
type OrientationPlan struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 达人定向计划表
	//[ 1] plan_id                                        bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PlanID int64 `gorm:"column:plan_id;type:bigint;" json:"plan_id"` // 快手佣金计划id
	//[ 2] merchant_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	MerchantID int64 `gorm:"column:merchant_id;type:bigint;default:0;" json:"merchant_id"` // 央选商家id
	//[ 3] anchor_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	AnchorID int64 `gorm:"column:anchor_id;type:bigint;default:0;" json:"anchor_id"` // 主播ID
	//[ 4] commission_id                                  bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	CommissionID int64 `gorm:"column:commission_id;type:bigint;" json:"commission_id"` // 快手佣金ID
	//[ 5] promoter_id                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PromoterID int64 `gorm:"column:promoter_id;type:bigint;" json:"promoter_id"` // 快手推广人id
	//[ 6] third_product_id                               bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ThirdProductID int64 `gorm:"column:third_product_id;type:bigint;default:0;" json:"third_product_id"` // 快手商品id
	//[ 7] product_id                                     bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ProductID int64 `gorm:"column:product_id;type:bigint;default:0;" json:"product_id"` // 央选商品id
	//[ 8] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 快手佣金率 1000倍
	//[ 9] plan_type                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	PlanType int32 `gorm:"column:plan_type;type:tinyint;default:0;" json:"plan_type"` // 计划类型。1-普通计划；2-商品定向佣金;5-专属计划\n
	//[10] plan_status                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	PlanStatus int32 `gorm:"column:plan_status;type:tinyint;default:1;" json:"plan_status"` // 定向计划状态。1-开启; 3-关闭 2:关闭中 投放状态
	//[11] platform_id                                    tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [2]
	PlatformID int32 `gorm:"column:platform_id;type:tinyint;default:2;" json:"platform_id"` // 1:抖音 2快手
	//[12] activity_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	ActivityID uint32 `gorm:"column:activity_id;type:uint;default:0;" json:"activity_id"` // 活动id
	//[13] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[14] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
	//[15] start_time                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	StartTime int32 `gorm:"column:start_time;type:int;default:0;" json:"start_time"` // 推广开始时间
	//[16] end_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EndTime int32 `gorm:"column:end_time;type:int;default:0;" json:"end_time"` // 推广结束时间
	//[17] new_cos_status                                 tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	NewCosStatus int32 `gorm:"column:new_cos_status;type:tinyint;default:1;" json:"new_cos_status"` // 修改后的佣金状态1:已生效 2:待生效
	//[18] new_cos_rate                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NewCosRate int32 `gorm:"column:new_cos_rate;type:int;default:0;" json:"new_cos_rate"` // 修改后的佣金
	//[19] new_cos_rate_start_time                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	NewCosRateStartTime int32 `gorm:"column:new_cos_rate_start_time;type:int;default:0;" json:"new_cos_rate_start_time"` // 修改后的生效时间
	//[20] talent_delete_time                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentDeleteTime int32 `gorm:"column:talent_delete_time;type:int;default:0;" json:"talent_delete_time"` // 达人移除时间
	//[21] is_create                                      tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsCreate int32 `gorm:"column:is_create;type:tinyint;default:0;" json:"is_create"` // 是否重建计划&达人 0:不需要 1:需要

}

var orientation_planTableInfo = &TableInfo{
	Name: "orientation_plan",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `达人定向计划表`,
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
			Comment:            `快手佣金计划id`,
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
			Name:               "merchant_id",
			Comment:            `央选商家id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "anchor_id",
			Comment:            `主播ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "int64",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "commission_id",
			Comment:            `快手佣金ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "CommissionID",
			GoFieldType:        "int64",
			JSONFieldName:      "commission_id",
			ProtobufFieldName:  "commission_id",
			ProtobufType:       "int64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "promoter_id",
			Comment:            `快手推广人id`,
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
			Name:               "third_product_id",
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
			GoFieldName:        "ThirdProductID",
			GoFieldType:        "int64",
			JSONFieldName:      "third_product_id",
			ProtobufFieldName:  "third_product_id",
			ProtobufType:       "int64",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "product_id",
			Comment:            `央选商品id`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "commission_rate",
			Comment:            `快手佣金率 1000倍`,
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
			Name:               "plan_status",
			Comment:            `定向计划状态。1-开启; 3-关闭 2:关闭中 投放状态`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "platform_id",
			Comment:            `1:抖音 2快手`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "activity_id",
			Comment:            `活动id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ActivityID",
			GoFieldType:        "uint32",
			JSONFieldName:      "activity_id",
			ProtobufFieldName:  "activity_id",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "start_time",
			Comment:            `推广开始时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "int32",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "end_time",
			Comment:            `推广结束时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "int32",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "new_cos_status",
			Comment:            `修改后的佣金状态1:已生效 2:待生效`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "NewCosStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "new_cos_status",
			ProtobufFieldName:  "new_cos_status",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "new_cos_rate",
			Comment:            `修改后的佣金`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "NewCosRate",
			GoFieldType:        "int32",
			JSONFieldName:      "new_cos_rate",
			ProtobufFieldName:  "new_cos_rate",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "new_cos_rate_start_time",
			Comment:            `修改后的生效时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "NewCosRateStartTime",
			GoFieldType:        "int32",
			JSONFieldName:      "new_cos_rate_start_time",
			ProtobufFieldName:  "new_cos_rate_start_time",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "talent_delete_time",
			Comment:            `达人移除时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentDeleteTime",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_delete_time",
			ProtobufFieldName:  "talent_delete_time",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "is_create",
			Comment:            `是否重建计划&达人 0:不需要 1:需要`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsCreate",
			GoFieldType:        "int32",
			JSONFieldName:      "is_create",
			ProtobufFieldName:  "is_create",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OrientationPlan) TableName() string {
	return "orientation_plan"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrientationPlan) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrientationPlan) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrientationPlan) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OrientationPlan) TableInfo() *TableInfo {
	return orientation_planTableInfo
}
