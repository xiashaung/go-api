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


CREATE TABLE `coupon_dy` (
  `dy_coupon_id` bigint(20) NOT NULL COMMENT '抖音优惠券信息表',
  `dy_coupon_name` varchar(60) NOT NULL COMMENT '抖音优惠券名称',
  `dy_coupon_desc` text COMMENT '抖音优惠券描述信息',
  `dy_coupon_credit` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '满减/直减券金额',
  `dy_coupon_type` int(255) DEFAULT NULL COMMENT '优惠券类型',
  `dy_coupon_discount` int(255) DEFAULT NULL COMMENT '折扣券折扣',
  `dy_pay_type` int(255) DEFAULT NULL COMMENT '支付类型',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '优惠券领取时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`dy_coupon_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "dy_coupon_id": 88,    "dy_coupon_name": "hvluQZdwMIcdqnfmLHyHINDIf",    "dy_coupon_desc": "ibBvSdgnRkdIYyDHbuKiNiKTD",    "dy_coupon_credit": 67,    "dy_coupon_type": 63,    "dy_coupon_discount": 45,    "dy_pay_type": 5,    "add_time": 50,    "last_time": 37}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned



*/

// CouponDy struct is a row record of the coupon_dy table in the yx database
type CouponDy struct {
	//[ 0] dy_coupon_id                                   bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	DyCouponID int64 `gorm:"primary_key;column:dy_coupon_id;type:bigint;" json:"dy_coupon_id"` // 抖音优惠券信息表
	//[ 1] dy_coupon_name                                 varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	DyCouponName string `gorm:"column:dy_coupon_name;type:varchar;size:60;" json:"dy_coupon_name"` // 抖音优惠券名称
	//[ 2] dy_coupon_desc                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	DyCouponDesc sql.NullString `gorm:"column:dy_coupon_desc;type:text;size:65535;" json:"dy_coupon_desc"` // 抖音优惠券描述信息
	//[ 3] dy_coupon_credit                               uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	DyCouponCredit uint32 `gorm:"column:dy_coupon_credit;type:uint;default:0;" json:"dy_coupon_credit"` // 满减/直减券金额
	//[ 4] dy_coupon_type                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DyCouponType sql.NullInt64 `gorm:"column:dy_coupon_type;type:int;" json:"dy_coupon_type"` // 优惠券类型
	//[ 5] dy_coupon_discount                             int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DyCouponDiscount sql.NullInt64 `gorm:"column:dy_coupon_discount;type:int;" json:"dy_coupon_discount"` // 折扣券折扣
	//[ 6] dy_pay_type                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DyPayType sql.NullInt64 `gorm:"column:dy_pay_type;type:int;" json:"dy_pay_type"` // 支付类型
	//[ 7] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 优惠券领取时间
	//[ 8] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var coupon_dyTableInfo = &TableInfo{
	Name: "coupon_dy",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "dy_coupon_id",
			Comment:            `抖音优惠券信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "DyCouponID",
			GoFieldType:        "int64",
			JSONFieldName:      "dy_coupon_id",
			ProtobufFieldName:  "dy_coupon_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "dy_coupon_name",
			Comment:            `抖音优惠券名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "DyCouponName",
			GoFieldType:        "string",
			JSONFieldName:      "dy_coupon_name",
			ProtobufFieldName:  "dy_coupon_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "dy_coupon_desc",
			Comment:            `抖音优惠券描述信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "DyCouponDesc",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dy_coupon_desc",
			ProtobufFieldName:  "dy_coupon_desc",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "dy_coupon_credit",
			Comment:            `满减/直减券金额`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DyCouponCredit",
			GoFieldType:        "uint32",
			JSONFieldName:      "dy_coupon_credit",
			ProtobufFieldName:  "dy_coupon_credit",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "dy_coupon_type",
			Comment:            `优惠券类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DyCouponType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dy_coupon_type",
			ProtobufFieldName:  "dy_coupon_type",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "dy_coupon_discount",
			Comment:            `折扣券折扣`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DyCouponDiscount",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dy_coupon_discount",
			ProtobufFieldName:  "dy_coupon_discount",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "dy_pay_type",
			Comment:            `支付类型`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DyPayType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dy_pay_type",
			ProtobufFieldName:  "dy_pay_type",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "add_time",
			Comment:            `优惠券领取时间`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CouponDy) TableName() string {
	return "coupon_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CouponDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CouponDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CouponDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CouponDy) TableInfo() *TableInfo {
	return coupon_dyTableInfo
}
