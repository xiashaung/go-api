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


CREATE TABLE `prop_info` (
  `prop_id` bigint(20) NOT NULL COMMENT '售卖属性信息表',
  `cate_id` int(11) NOT NULL COMMENT '类目id',
  `shop_id` int(11) NOT NULL COMMENT '店铺id',
  `prop_name` varchar(20) NOT NULL COMMENT '售卖属性名称',
  `prop_type` int(255) unsigned NOT NULL DEFAULT '1' COMMENT '类目类型 1：系统标准	2：自定义属性',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '售卖属性添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`prop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='售卖属性信息表'

JSON Sample
-------------------------------------
{    "prop_id": 75,    "cate_id": 55,    "shop_id": 27,    "prop_name": "bZnjqRGsqVqjeyPlbAEkjwfnx",    "prop_type": 14,    "add_time": 73,    "last_time": 85}


Comments
-------------------------------------
[ 4] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned



*/

// PropInfo struct is a row record of the prop_info table in the yx database
type PropInfo struct {
	//[ 0] prop_id                                        bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	PropID int64 `gorm:"primary_key;column:prop_id;type:bigint;" json:"prop_id"` // 售卖属性信息表
	//[ 1] cate_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CateID int32 `gorm:"column:cate_id;type:int;" json:"cate_id"` // 类目id
	//[ 2] shop_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ShopID int32 `gorm:"column:shop_id;type:int;" json:"shop_id"` // 店铺id
	//[ 3] prop_name                                      varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	PropName string `gorm:"column:prop_name;type:varchar;size:20;" json:"prop_name"` // 售卖属性名称
	//[ 4] prop_type                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	PropType uint32 `gorm:"column:prop_type;type:uint;default:1;" json:"prop_type"` // 类目类型 1：系统标准	2：自定义属性
	//[ 5] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 售卖属性添加时间
	//[ 6] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var prop_infoTableInfo = &TableInfo{
	Name: "prop_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "prop_id",
			Comment:            `售卖属性信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PropID",
			GoFieldType:        "int64",
			JSONFieldName:      "prop_id",
			ProtobufFieldName:  "prop_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "shop_id",
			Comment:            `店铺id`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "prop_name",
			Comment:            `售卖属性名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "PropName",
			GoFieldType:        "string",
			JSONFieldName:      "prop_name",
			ProtobufFieldName:  "prop_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "prop_type",
			Comment:            `类目类型 1：系统标准	2：自定义属性`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PropType",
			GoFieldType:        "uint32",
			JSONFieldName:      "prop_type",
			ProtobufFieldName:  "prop_type",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "add_time",
			Comment:            `售卖属性添加时间`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PropInfo) TableName() string {
	return "prop_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PropInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PropInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PropInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PropInfo) TableInfo() *TableInfo {
	return prop_infoTableInfo
}
