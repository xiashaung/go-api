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


CREATE TABLE `talent_cate` (
  `item_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `talent_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '达人id',
  `one_level_cate_id` int(11) NOT NULL DEFAULT '0' COMMENT '一级意向类目',
  `cate_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '二级意向类目id',
  `cate_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '二级意向类目名称',
  `add_time` int(10) NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) NOT NULL DEFAULT '0' COMMENT '最近修改时间',
  PRIMARY KEY (`item_id`),
  KEY `idx_talent` (`talent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=740 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='达人意向类目表'

JSON Sample
-------------------------------------
{    "item_id": 10,    "talent_id": 75,    "one_level_cate_id": 28,    "cate_id": 60,    "cate_name": "YyMbwLfKjZUxrWXOHnraQuUcZ",    "add_time": 98,    "last_time": 51}


Comments
-------------------------------------
[ 1] column is set for unsigned
[ 3] column is set for unsigned



*/

// TalentCate struct is a row record of the talent_cate table in the yx database
type TalentCate struct {
	//[ 0] item_id                                        bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ItemID int64 `gorm:"primary_key;AUTO_INCREMENT;column:item_id;type:bigint;" json:"item_id"` // 自增主键
	//[ 1] talent_id                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	TalentID uint32 `gorm:"column:talent_id;type:uint;default:0;" json:"talent_id"` // 达人id
	//[ 2] one_level_cate_id                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	OneLevelCateID int32 `gorm:"column:one_level_cate_id;type:int;default:0;" json:"one_level_cate_id"` // 一级意向类目
	//[ 3] cate_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CateID uint32 `gorm:"column:cate_id;type:uint;default:0;" json:"cate_id"` // 二级意向类目id
	//[ 4] cate_name                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	CateName string `gorm:"column:cate_name;type:varchar;size:255;" json:"cate_name"` // 二级意向类目名称
	//[ 5] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"` // 添加时间
	//[ 6] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime int32 `gorm:"column:last_time;type:int;default:0;" json:"last_time"` // 最近修改时间

}

var talent_cateTableInfo = &TableInfo{
	Name: "talent_cate",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "item_id",
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
			GoFieldName:        "ItemID",
			GoFieldType:        "int64",
			JSONFieldName:      "item_id",
			ProtobufFieldName:  "item_id",
			ProtobufType:       "int64",
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
			Name:               "one_level_cate_id",
			Comment:            `一级意向类目`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OneLevelCateID",
			GoFieldType:        "int32",
			JSONFieldName:      "one_level_cate_id",
			ProtobufFieldName:  "one_level_cate_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "cate_id",
			Comment:            `二级意向类目id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CateID",
			GoFieldType:        "uint32",
			JSONFieldName:      "cate_id",
			ProtobufFieldName:  "cate_id",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "cate_name",
			Comment:            `二级意向类目名称`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "add_time",
			Comment:            `添加时间`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_time",
			Comment:            `最近修改时间`,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TalentCate) TableName() string {
	return "talent_cate"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TalentCate) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TalentCate) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TalentCate) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TalentCate) TableInfo() *TableInfo {
	return talent_cateTableInfo
}
