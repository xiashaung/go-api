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


CREATE TABLE `share_scene` (
  `scene_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `scene_key` varchar(50) NOT NULL DEFAULT '' COMMENT 'scene对应的key',
  `scene_value` varchar(255) NOT NULL DEFAULT '' COMMENT 'scene对应的值',
  `scene_type` tinyint(3) DEFAULT '0' COMMENT '场景值类型：1、分享商品',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`scene_id`),
  UNIQUE KEY `idx_scene_key` (`scene_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='小程序分享scene表'

JSON Sample
-------------------------------------
{    "scene_id": 62,    "scene_key": "QsixXNOrpXxAsSdsnuITpoLhA",    "scene_value": "BgvOOTwyAQZDNCHAMjGyoHmhY",    "scene_type": 28,    "add_time": 95,    "last_time": 47}


Comments
-------------------------------------
[ 4] column is set for unsigned
[ 5] column is set for unsigned



*/

// ShareScene struct is a row record of the share_scene table in the yx database
type ShareScene struct {
	//[ 0] scene_id                                       int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	SceneID int32 `gorm:"primary_key;AUTO_INCREMENT;column:scene_id;type:int;" json:"scene_id"` // 自增id
	//[ 1] scene_key                                      varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	SceneKey string `gorm:"column:scene_key;type:varchar;size:50;" json:"scene_key"` // scene对应的key
	//[ 2] scene_value                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SceneValue string `gorm:"column:scene_value;type:varchar;size:255;" json:"scene_value"` // scene对应的值
	//[ 3] scene_type                                     tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	SceneType sql.NullInt64 `gorm:"column:scene_type;type:tinyint;default:0;" json:"scene_type"` // 场景值类型：1、分享商品
	//[ 4] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加时间
	//[ 5] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var share_sceneTableInfo = &TableInfo{
	Name: "share_scene",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "scene_id",
			Comment:            `自增id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SceneID",
			GoFieldType:        "int32",
			JSONFieldName:      "scene_id",
			ProtobufFieldName:  "scene_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "scene_key",
			Comment:            `scene对应的key`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "SceneKey",
			GoFieldType:        "string",
			JSONFieldName:      "scene_key",
			ProtobufFieldName:  "scene_key",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "scene_value",
			Comment:            `scene对应的值`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SceneValue",
			GoFieldType:        "string",
			JSONFieldName:      "scene_value",
			ProtobufFieldName:  "scene_value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "scene_type",
			Comment:            `场景值类型：1、分享商品`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "SceneType",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "scene_type",
			ProtobufFieldName:  "scene_type",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *ShareScene) TableName() string {
	return "share_scene"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ShareScene) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ShareScene) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ShareScene) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *ShareScene) TableInfo() *TableInfo {
	return share_sceneTableInfo
}
