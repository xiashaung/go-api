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


CREATE TABLE `admin_user_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `role_id` int(11) DEFAULT NULL,
  `add_time` int(11) DEFAULT NULL,
  `last_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "id": 27,    "user_id": 33,    "role_id": 22,    "add_time": 76,    "last_time": 24}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// AdminUserRole struct is a row record of the admin_user_role table in the yx database
type AdminUserRole struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] user_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UserID sql.NullInt64 `gorm:"column:user_id;type:int;" json:"user_id"`
	//[ 2] role_id                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RoleID sql.NullInt64 `gorm:"column:role_id;type:int;" json:"role_id"`
	//[ 3] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"`
	//[ 4] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var admin_user_roleTableInfo = &TableInfo{
	Name: "admin_user_role",
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
			Name:               "user_id",
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
			GoFieldName:        "UserID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "role_id",
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
			GoFieldName:        "RoleID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "role_id",
			ProtobufFieldName:  "role_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdminUserRole) TableName() string {
	return "admin_user_role"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdminUserRole) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdminUserRole) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdminUserRole) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdminUserRole) TableInfo() *TableInfo {
	return admin_user_roleTableInfo
}
