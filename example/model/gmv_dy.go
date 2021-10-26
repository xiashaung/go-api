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


CREATE TABLE `gmv_dy` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `order_gmv` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单gmv',
  `colonel_gmv` bigint(20) NOT NULL DEFAULT '0' COMMENT '团长gmv',
  `date` date NOT NULL COMMENT '日期',
  `add_time` int(11) DEFAULT '0',
  `last_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 44,    "order_gmv": 74,    "colonel_gmv": 69,    "date": "2287-04-07T23:15:12.693150013+08:00",    "add_time": 29,    "last_time": 20}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// GmvDy struct is a row record of the gmv_dy table in the yx database
type GmvDy struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] order_gmv                                      bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	OrderGmv int64 `gorm:"column:order_gmv;type:bigint;default:0;" json:"order_gmv"` // 订单gmv
	//[ 2] colonel_gmv                                    bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: [0]
	ColonelGmv int64 `gorm:"column:colonel_gmv;type:bigint;default:0;" json:"colonel_gmv"` // 团长gmv
	//[ 3] date                                           date                 null: false  primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	Date time.Time `gorm:"column:date;type:date;" json:"date"` // 日期
	//[ 4] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 5] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var gmv_dyTableInfo = &TableInfo{
	Name: "gmv_dy",
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
			Name:               "order_gmv",
			Comment:            `订单gmv`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "OrderGmv",
			GoFieldType:        "int64",
			JSONFieldName:      "order_gmv",
			ProtobufFieldName:  "order_gmv",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "colonel_gmv",
			Comment:            `团长gmv`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ColonelGmv",
			GoFieldType:        "int64",
			JSONFieldName:      "colonel_gmv",
			ProtobufFieldName:  "colonel_gmv",
			ProtobufType:       "int64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "date",
			Comment:            `日期`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "Date",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GmvDy) TableName() string {
	return "gmv_dy"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GmvDy) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GmvDy) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GmvDy) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GmvDy) TableInfo() *TableInfo {
	return gmv_dyTableInfo
}
