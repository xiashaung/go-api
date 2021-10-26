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


CREATE TABLE `product_window` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '橱窗商品关联表',
  `talent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `anchor_id` int(11) NOT NULL DEFAULT '0' COMMENT '主播id',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `third_product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '第三方商品id',
  `platform_id` int(11) NOT NULL DEFAULT '0' COMMENT '平台id',
  `activity_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '活动id',
  `third_window_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '第三方平台橱窗id，抖音橱窗删除时需要用到',
  `window_status` int(1) NOT NULL DEFAULT '0' COMMENT '橱窗状态 0 橱窗下  1 橱窗上',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_pid_talend_plat` (`product_id`,`talent_id`,`platform_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3452 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 44,    "talent_id": 16,    "anchor_id": 80,    "product_id": 39,    "third_product_id": 98,    "platform_id": 26,    "activity_id": 87,    "third_window_id": 24,    "window_status": 3,    "add_time": 72,    "last_time": 75}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// ProductWindow struct is a row record of the product_window table in the yx database
type ProductWindow struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" json:"id"` // 橱窗商品关联表
	//[ 1] talent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TalentID uint32 `gorm:"column:talent_id;type:uint;default:0;" json:"talent_id"` // 达人id
	//[ 2] anchor_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AnchorID int32 `gorm:"column:anchor_id;type:int;default:0;" json:"anchor_id"` // 主播id
	//[ 3] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品id
	//[ 4] third_product_id                               ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ThirdProductID uint64 `gorm:"column:third_product_id;type:ubigint;default:0;" json:"third_product_id"` // 第三方商品id
	//[ 5] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"` // 平台id
	//[ 6] activity_id                                    utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	ActivityID uint32 `gorm:"column:activity_id;type:utinyint;default:0;" json:"activity_id"` // 活动id
	//[ 7] third_window_id                                ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ThirdWindowID uint64 `gorm:"column:third_window_id;type:ubigint;default:0;" json:"third_window_id"` // 第三方平台橱窗id，抖音橱窗删除时需要用到
	//[ 8] window_status                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	WindowStatus int32 `gorm:"column:window_status;type:int;default:0;" json:"window_status"` // 橱窗状态 0 橱窗下  1 橱窗上
	//[ 9] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[10] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var product_windowTableInfo = &TableInfo{
	Name: "product_window",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `橱窗商品关联表`,
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
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "anchor_id",
			Comment:            `主播id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "int32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_id",
			Comment:            `商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "third_product_id",
			Comment:            `第三方商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "third_product_id",
			ProtobufFieldName:  "third_product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "platform_id",
			Comment:            `平台id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
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
			Name:               "activity_id",
			Comment:            `活动id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ActivityID",
			GoFieldType:        "uint32",
			JSONFieldName:      "activity_id",
			ProtobufFieldName:  "activity_id",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "third_window_id",
			Comment:            `第三方平台橱窗id，抖音橱窗删除时需要用到`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdWindowID",
			GoFieldType:        "uint64",
			JSONFieldName:      "third_window_id",
			ProtobufFieldName:  "third_window_id",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "window_status",
			Comment:            `橱窗状态 0 橱窗下  1 橱窗上`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "WindowStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "window_status",
			ProtobufFieldName:  "window_status",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "add_time",
			Comment:            `添加时间`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *ProductWindow) TableName() string {
	return "product_window"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *ProductWindow) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *ProductWindow) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *ProductWindow) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *ProductWindow) TableInfo() *TableInfo {
	return product_windowTableInfo
}
