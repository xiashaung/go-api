package example

/*
DB Table Details
-------------------------------------


CREATE TABLE `crontab_break` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `break_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '断点名称',
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '断点位置',
  `add_time` int(10) NOT NULL DEFAULT '0',
  `last_time` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='定时任务断点表'

JSON Sample
-------------------------------------
{    "id": 39,    "break_name": "uoMAuvjcNDcaGqJwxYLOAUycC",    "value": "fqZPJFgyyIxGBWQlWeFLgTnCi",    "add_time": 35,    "last_time": 51}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// CrontabBreak struct is a row record of the crontab_break table in the yx database
type CrontabBreak struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] break_name                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	BreakName string `gorm:"column:break_name;type:varchar;size:255;" json:"break_name"` // 断点名称
	//[ 2] value                                          varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Value string `gorm:"column:value;type:varchar;size:255;" json:"value"` // 断点位置
	//[ 3] add_time                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime int32 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[ 4] last_time                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime int32 `gorm:"column:last_time;type:int;" json:"last_time"`
}

var crontab_breakTableInfo = &TableInfo{
	Name: "crontab_break",
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
			Name:               "break_name",
			Comment:            `断点名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "BreakName",
			GoFieldType:        "string",
			JSONFieldName:      "break_name",
			ProtobufFieldName:  "break_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "value",
			Comment:            `断点位置`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Value",
			GoFieldType:        "string",
			JSONFieldName:      "value",
			ProtobufFieldName:  "value",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "add_time",
			Comment:            ``,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "last_time",
			Comment:            ``,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CrontabBreak) TableName() string {
	return "crontab_break"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CrontabBreak) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CrontabBreak) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CrontabBreak) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CrontabBreak) TableInfo() *TableInfo {
	return crontab_breakTableInfo
}
