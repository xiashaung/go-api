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


CREATE TABLE `delivery_info` (
  `delivery_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '达人收货地址信息表',
  `talent_id` int(11) NOT NULL DEFAULT '0' COMMENT '达人id',
  `delivery_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收货人姓名',
  `delivery_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收货人手机号',
  `province` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '省名称',
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '城市名称',
  `area` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '区县名称',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货地址code',
  `delivery_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收货人详细地址',
  `delivery_status` smallint(6) unsigned NOT NULL DEFAULT '1' COMMENT '状态 1：正常 0：删除',
  `is_default` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否默认地址 0 否 1 是',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建收货地址时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`delivery_id`),
  KEY `idx_talent` (`talent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='达人收货地址信息表'

JSON Sample
-------------------------------------
{    "delivery_id": 34,    "talent_id": 54,    "delivery_name": "ewgixHIfqSkRSpkuVyUjoBboc",    "delivery_mobile": "mwnPnuNUrveSBpNQvXXYlhvQM",    "province": "qwuKQFhkDkOifFNsVmIdOvGvh",    "city": "TLXIcDfgiZVHFXiOQeOTvFRbP",    "area": "GpogYyAFECZLfGqeytDtPtlPF",    "code": "DOaDPAYkqbwRekbcIsZttHqOH",    "delivery_addr": "hScbrOPOlxyQuWZdHUbkuvhDQ",    "delivery_status": 23,    "is_default": 95,    "add_time": 27,    "last_time": 90}


Comments
-------------------------------------
[ 9] column is set for unsigned
[11] column is set for unsigned
[12] column is set for unsigned



*/

// DeliveryInfo struct is a row record of the delivery_info table in the yx database
type DeliveryInfo struct {
	//[ 0] delivery_id                                    int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	DeliveryID int32 `gorm:"primary_key;AUTO_INCREMENT;column:delivery_id;type:int;" json:"delivery_id"` // 达人收货地址信息表
	//[ 1] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"` // 达人id
	//[ 2] delivery_name                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	DeliveryName string `gorm:"column:delivery_name;type:varchar;size:255;" json:"delivery_name"` // 收货人姓名
	//[ 3] delivery_mobile                                char(11)             null: false  primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	DeliveryMobile string `gorm:"column:delivery_mobile;type:char;size:11;" json:"delivery_mobile"` // 收货人手机号
	//[ 4] province                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Province string `gorm:"column:province;type:varchar;size:255;" json:"province"` // 省名称
	//[ 5] city                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	City string `gorm:"column:city;type:varchar;size:255;" json:"city"` // 城市名称
	//[ 6] area                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Area string `gorm:"column:area;type:varchar;size:255;" json:"area"` // 区县名称
	//[ 7] code                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Code string `gorm:"column:code;type:varchar;size:255;" json:"code"` // 收货地址code
	//[ 8] delivery_addr                                  varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	DeliveryAddr string `gorm:"column:delivery_addr;type:varchar;size:255;" json:"delivery_addr"` // 收货人详细地址
	//[ 9] delivery_status                                usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [1]
	DeliveryStatus uint32 `gorm:"column:delivery_status;type:usmallint;default:1;" json:"delivery_status"` // 状态 1：正常 0：删除
	//[10] is_default                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [0]
	IsDefault int32 `gorm:"column:is_default;type:smallint;default:0;" json:"is_default"` // 是否默认地址 0 否 1 是
	//[11] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建收货地址时间
	//[12] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var delivery_infoTableInfo = &TableInfo{
	Name: "delivery_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "delivery_id",
			Comment:            `达人收货地址信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DeliveryID",
			GoFieldType:        "int32",
			JSONFieldName:      "delivery_id",
			ProtobufFieldName:  "delivery_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "delivery_name",
			Comment:            `收货人姓名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "DeliveryName",
			GoFieldType:        "string",
			JSONFieldName:      "delivery_name",
			ProtobufFieldName:  "delivery_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "delivery_mobile",
			Comment:            `收货人手机号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "DeliveryMobile",
			GoFieldType:        "string",
			JSONFieldName:      "delivery_mobile",
			ProtobufFieldName:  "delivery_mobile",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "province",
			Comment:            `省名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Province",
			GoFieldType:        "string",
			JSONFieldName:      "province",
			ProtobufFieldName:  "province",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "city",
			Comment:            `城市名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "City",
			GoFieldType:        "string",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "area",
			Comment:            `区县名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Area",
			GoFieldType:        "string",
			JSONFieldName:      "area",
			ProtobufFieldName:  "area",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "code",
			Comment:            `收货地址code`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Code",
			GoFieldType:        "string",
			JSONFieldName:      "code",
			ProtobufFieldName:  "code",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "delivery_addr",
			Comment:            `收货人详细地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "DeliveryAddr",
			GoFieldType:        "string",
			JSONFieldName:      "delivery_addr",
			ProtobufFieldName:  "delivery_addr",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "delivery_status",
			Comment:            `状态 1：正常 0：删除`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "DeliveryStatus",
			GoFieldType:        "uint32",
			JSONFieldName:      "delivery_status",
			ProtobufFieldName:  "delivery_status",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "is_default",
			Comment:            `是否默认地址 0 否 1 是`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "IsDefault",
			GoFieldType:        "int32",
			JSONFieldName:      "is_default",
			ProtobufFieldName:  "is_default",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "add_time",
			Comment:            `创建收货地址时间`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DeliveryInfo) TableName() string {
	return "delivery_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DeliveryInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DeliveryInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DeliveryInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DeliveryInfo) TableInfo() *TableInfo {
	return delivery_infoTableInfo
}
