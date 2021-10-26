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


CREATE TABLE `tender_cate` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '招商计划类目关联表',
  `tender_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '招商计划id',
  `cate_id` int(11) NOT NULL DEFAULT '0' COMMENT '类目id',
  `cate_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类目名称',
  `cate_level` int(10) unsigned NOT NULL DEFAULT '2' COMMENT '分类等级 0 行业大类 1 一级类目 2 二级类目',
  `commission_rate` int(11) NOT NULL DEFAULT '0' COMMENT '佣金率，扩大一千倍',
  `product_num` int(11) NOT NULL DEFAULT '0' COMMENT '招商数量',
  `enroll_merchant` int(11) NOT NULL DEFAULT '0' COMMENT '已经报名的商家数量',
  `enroll_goods` int(11) NOT NULL DEFAULT '0' COMMENT '已经报名的商品数量',
  `purpose_num` int(11) NOT NULL DEFAULT '0' COMMENT '意向商品数量',
  `selected_num` int(11) NOT NULL DEFAULT '0' COMMENT '入选商品数量',
  `add_time` int(10) NOT NULL DEFAULT '0',
  `last_time` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `idx_tender_id` (`tender_id`),
  KEY `idx_cate_id` (`cate_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='招商计划类目关联表'

JSON Sample
-------------------------------------
{    "id": 98,    "tender_id": 20,    "cate_id": 33,    "cate_name": "wSskOGOAeKosdSvrhtOtZyjaa",    "cate_level": 48,    "commission_rate": 84,    "product_num": 58,    "enroll_merchant": 64,    "enroll_goods": 78,    "purpose_num": 82,    "selected_num": 98,    "add_time": 12,    "last_time": 43}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 4] column is set for unsigned



*/

// TenderCate struct is a row record of the tender_cate table in the yx database
type TenderCate struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 招商计划类目关联表
	//[ 1] tender_id                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	TenderID int64 `gorm:"column:tender_id;type:bigint;default:0;" json:"tender_id"` // 招商计划id
	//[ 2] cate_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CateID int32 `gorm:"column:cate_id;type:int;default:0;" json:"cate_id"` // 类目id
	//[ 3] cate_name                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:255;" json:"cate_name"` // 类目名称
	//[ 4] cate_level                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [2]
	CateLevel uint32 `gorm:"column:cate_level;type:uint;default:2;" json:"cate_level"` // 分类等级 0 行业大类 1 一级类目 2 二级类目
	//[ 5] commission_rate                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CommissionRate int32 `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"` // 佣金率，扩大一千倍
	//[ 6] product_num                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ProductNum int32 `gorm:"column:product_num;type:int;default:0;" json:"product_num"` // 招商数量
	//[ 7] enroll_merchant                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EnrollMerchant int32 `gorm:"column:enroll_merchant;type:int;default:0;" json:"enroll_merchant"` // 已经报名的商家数量
	//[ 8] enroll_goods                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	EnrollGoods int32 `gorm:"column:enroll_goods;type:int;default:0;" json:"enroll_goods"` // 已经报名的商品数量
	//[ 9] purpose_num                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PurposeNum int32 `gorm:"column:purpose_num;type:int;default:0;" json:"purpose_num"` // 意向商品数量
	//[10] selected_num                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SelectedNum int32 `gorm:"column:selected_num;type:int;default:0;" json:"selected_num"` // 入选商品数量
	//[11] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[12] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var tender_cateTableInfo = &TableInfo{
	Name: "tender_cate",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `招商计划类目关联表`,
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
			Name:               "tender_id",
			Comment:            `招商计划id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "cate_id",
			Comment:            `类目id`,
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
			Comment:            `类目名称`,
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
			Name:               "cate_level",
			Comment:            `分类等级 0 行业大类 1 一级类目 2 二级类目`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_level",
			ProtobufFieldName:  "cate_level",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "commission_rate",
			Comment:            `佣金率，扩大一千倍`,
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
			Name:               "product_num",
			Comment:            `招商数量`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "enroll_merchant",
			Comment:            `已经报名的商家数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EnrollMerchant",
			GoFieldType:        "int32",
			JSONFieldName:      "enroll_merchant",
			ProtobufFieldName:  "enroll_merchant",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "enroll_goods",
			Comment:            `已经报名的商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EnrollGoods",
			GoFieldType:        "int32",
			JSONFieldName:      "enroll_goods",
			ProtobufFieldName:  "enroll_goods",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "purpose_num",
			Comment:            `意向商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PurposeNum",
			GoFieldType:        "int32",
			JSONFieldName:      "purpose_num",
			ProtobufFieldName:  "purpose_num",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "selected_num",
			Comment:            `入选商品数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SelectedNum",
			GoFieldType:        "int32",
			JSONFieldName:      "selected_num",
			ProtobufFieldName:  "selected_num",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TenderCate) TableName() string {
	return "tender_cate"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TenderCate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TenderCate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TenderCate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TenderCate) TableInfo() *TableInfo {
	return tender_cateTableInfo
}
