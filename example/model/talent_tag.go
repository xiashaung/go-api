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


CREATE TABLE `talent_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `talent_id` int(11) DEFAULT NULL COMMENT '达人id',
  `tag_type` int(11) NOT NULL COMMENT '达人标签类型id',
  `first_id` int(11) DEFAULT NULL COMMENT '达人id',
  `tag_id` int(11) DEFAULT NULL COMMENT '标签id',
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `talent_id_idx` (`talent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='达人标签表'

JSON Sample
-------------------------------------
{    "id": 27,    "talent_id": 41,    "tag_type": 39,    "first_id": 10,    "tag_id": 20,    "add_time": 83,    "last_time": 39}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// TalentTag struct is a row record of the talent_tag table in the yx database
type TalentTag struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] talent_id                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TalentID sql.NullInt64 `gorm:"column:talent_id;type:int;" json:"talent_id"` // 达人id
	//[ 2] tag_type                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TagType int32 `gorm:"column:tag_type;type:int;" json:"tag_type"` // 达人标签类型id
	//[ 3] first_id                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FirstID sql.NullInt64 `gorm:"column:first_id;type:int;" json:"first_id"` // 达人id
	//[ 4] tag_id                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TagID sql.NullInt64 `gorm:"column:tag_id;type:int;" json:"tag_id"` // 标签id
	//[ 5] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 6] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var talent_tagTableInfo = &TableInfo{
	Name: "talent_tag",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "talent_id",
			Comment:            `达人id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "tag_type",
			Comment:            `达人标签类型id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TagType",
			GoFieldType:        "int32",
			JSONFieldName:      "tag_type",
			ProtobufFieldName:  "tag_type",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "first_id",
			Comment:            `达人id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FirstID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "first_id",
			ProtobufFieldName:  "first_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "tag_id",
			Comment:            `标签id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TagID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "tag_id",
			ProtobufFieldName:  "tag_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "add_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AddTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "add_time",
			ProtobufFieldName:  "add_time",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "last_time",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TalentTag) TableName() string {
	return "talent_tag"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TalentTag) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TalentTag) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TalentTag) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TalentTag) TableInfo() *TableInfo {
	return talent_tagTableInfo
}
