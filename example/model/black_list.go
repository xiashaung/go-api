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


CREATE TABLE `black_list` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` int(11) DEFAULT '0' COMMENT '类型 1 查看联系方式',
  `target_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '关联id 可以是达人id,商家id,也可以是ip地址',
  `status` int(11) DEFAULT '1' COMMENT '状态  1 启动 0 已关闭',
  `remark` varchar(255) DEFAULT '' COMMENT '备注',
  `start_time` int(11) DEFAULT NULL COMMENT '开始时间',
  `end_time` int(11) DEFAULT NULL COMMENT '结束时间',
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_target_id` (`target_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='黑名单'

JSON Sample
-------------------------------------
{    "id": 71,    "type": 91,    "target_id": "rNoNJHBLfagShheNSWgZydbqM",    "status": 16,    "remark": "XYEVOyNnwmGDPmCadvBrOghMA",    "start_time": 82,    "end_time": 71,    "add_time": 85,    "last_time": 30}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// BlackList struct is a row record of the black_list table in the yx database
type BlackList struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] type                                           int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Type sql.NullInt64 `gorm:"column:type;type:int;default:0;" json:"type"` // 类型 1 查看联系方式
	//[ 2] target_id                                      varchar(30)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	TargetID sql.NullString `gorm:"column:target_id;type:varchar;size:30;" json:"target_id"` // 关联id 可以是达人id,商家id,也可以是ip地址
	//[ 3] status                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	Status sql.NullInt64 `gorm:"column:status;type:int;default:1;" json:"status"` // 状态  1 启动 0 已关闭
	//[ 4] remark                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Remark sql.NullString `gorm:"column:remark;type:varchar;size:255;" json:"remark"` // 备注
	//[ 5] start_time                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	StartTime sql.NullInt64 `gorm:"column:start_time;type:int;" json:"start_time"` // 开始时间
	//[ 6] end_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	EndTime sql.NullInt64 `gorm:"column:end_time;type:int;" json:"end_time"` // 结束时间
	//[ 7] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 8] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var black_listTableInfo = &TableInfo{
	Name: "black_list",
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
			Name:               "type",
			Comment:            `类型 1 查看联系方式`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Type",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "target_id",
			Comment:            `关联id 可以是达人id,商家id,也可以是ip地址`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "TargetID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "target_id",
			ProtobufFieldName:  "target_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "status",
			Comment:            `状态  1 启动 0 已关闭`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "remark",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Remark",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "remark",
			ProtobufFieldName:  "remark",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "start_time",
			Comment:            `开始时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "StartTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "start_time",
			ProtobufFieldName:  "start_time",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "end_time",
			Comment:            `结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "EndTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "end_time",
			ProtobufFieldName:  "end_time",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (b *BlackList) TableName() string {
	return "black_list"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BlackList) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BlackList) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BlackList) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BlackList) TableInfo() *TableInfo {
	return black_listTableInfo
}
