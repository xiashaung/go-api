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


CREATE TABLE `tender_survey` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `tender_id` int(11) NOT NULL DEFAULT '0' COMMENT '招商计划id',
  `cate_id` int(11) NOT NULL DEFAULT '0' COMMENT '栏目id',
  `cate_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '栏目名称',
  `tender_num` int(11) NOT NULL DEFAULT '0' COMMENT '招商商品数量',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '栏目配置佣金比例',
  `reg_product_num` int(11) NOT NULL DEFAULT '0' COMMENT '报名商品数量',
  `reg_shop_num` int(11) NOT NULL DEFAULT '0' COMMENT '报名商家数量',
  `yx_product_num` int(11) NOT NULL DEFAULT '0' COMMENT '意向商品数',
  `xp_product_num` int(11) NOT NULL DEFAULT '0' COMMENT '选品商品数量',
  `add_time` int(10) NOT NULL DEFAULT '0',
  `last_time` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_tender` (`tender_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 43,    "tender_id": 66,    "cate_id": 95,    "cate_name": "bESalmPPvVUPpwsDTXISEmpZU",    "tender_num": 58,    "commission_rate": 21,    "reg_product_num": 56,    "reg_shop_num": 94,    "yx_product_num": 74,    "xp_product_num": 66,    "add_time": 88,    "last_time": 11}



*/

// TenderSurvey struct is a row record of the tender_survey table in the yx database
type TenderSurvey struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"` // 自增主键
	//[ 1] tender_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderID int32 `gorm:"column:tender_id;type:int;default:0;" json:"tender_id"` // 招商计划id
	//[ 2] cate_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateID int32 `gorm:"column:cate_id;type:int;default:0;" json:"cate_id"` // 栏目id
	//[ 3] cate_name                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:255;" json:"cate_name"` // 栏目名称
	//[ 4] tender_num                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TenderNum int32 `gorm:"column:tender_num;type:int;default:0;" json:"tender_num"` // 招商商品数量
	//[ 5] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 栏目配置佣金比例
	//[ 6] reg_product_num                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RegProductNum int32 `gorm:"column:reg_product_num;type:int;default:0;" json:"reg_product_num"` // 报名商品数量
	//[ 7] reg_shop_num                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RegShopNum int32 `gorm:"column:reg_shop_num;type:int;default:0;" json:"reg_shop_num"` // 报名商家数量
	//[ 8] yx_product_num                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	YxProductNum int32 `gorm:"column:yx_product_num;type:int;default:0;" json:"yx_product_num"` // 意向商品数
	//[ 9] xp_product_num                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	XpProductNum int32 `gorm:"column:xp_product_num;type:int;default:0;" json:"xp_product_num"` // 选品商品数量
	//[10] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[11] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var tender_surveyTableInfo = &TableInfo{
	Name: "tender_survey",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `自增主键`,
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
			Name:               "tender_id",
			Comment:            `招商计划id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderID",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_id",
			ProtobufFieldName:  "tender_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_id",
			Comment:            `栏目id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_name",
			Comment:            `栏目名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "CateName",
			GoFieldType:        "string",
			JSONFieldName:      "cate_name",
			ProtobufFieldName:  "cate_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "tender_num",
			Comment:            `招商商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TenderNum",
			GoFieldType:        "int32",
			JSONFieldName:      "tender_num",
			ProtobufFieldName:  "tender_num",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "commission_rate",
			Comment:            `栏目配置佣金比例`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "reg_product_num",
			Comment:            `报名商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RegProductNum",
			GoFieldType:        "int32",
			JSONFieldName:      "reg_product_num",
			ProtobufFieldName:  "reg_product_num",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "reg_shop_num",
			Comment:            `报名商家数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RegShopNum",
			GoFieldType:        "int32",
			JSONFieldName:      "reg_shop_num",
			ProtobufFieldName:  "reg_shop_num",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "yx_product_num",
			Comment:            `意向商品数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "YxProductNum",
			GoFieldType:        "int32",
			JSONFieldName:      "yx_product_num",
			ProtobufFieldName:  "yx_product_num",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "xp_product_num",
			Comment:            `选品商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "XpProductNum",
			GoFieldType:        "int32",
			JSONFieldName:      "xp_product_num",
			ProtobufFieldName:  "xp_product_num",
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
func (t *TenderSurvey) TableName() string {
	return "tender_survey"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TenderSurvey) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TenderSurvey) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TenderSurvey) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TenderSurvey) TableInfo() *TableInfo {
	return tender_surveyTableInfo
}
