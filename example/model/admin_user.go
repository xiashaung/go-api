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


CREATE TABLE `admin_user` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `head_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `loginfailure` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '失败次数',
  `login_time` bigint(20) DEFAULT NULL COMMENT '登录时间',
  `token` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'token标识',
  `status` tinyint(30) NOT NULL DEFAULT '1' COMMENT '状态 1正常 0 关闭',
  `login_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录IP',
  `add_time` int(10) DEFAULT NULL COMMENT '创建时间',
  `last_time` int(10) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `username` (`user_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表'

JSON Sample
-------------------------------------
{    "user_id": 32,    "user_name": "OwDdAmLGVriKJytpeRxhHCVmX",    "password": "pyuGFZatQsjHkXfFqWTxvxIuv",    "salt": "ZdEsoquSOgtAvOtGmUnxXbSrk",    "head_img": "xEtTiyxKevWhyYIdQkASQUrJd",    "user_email": "xNAhRyhyCEcAmMGDvdkLbGspK",    "loginfailure": 63,    "login_time": 31,    "token": "RgSDdXKIuUUQwpvjPigFaVRuL",    "status": 64,    "login_ip": "skUqSfvThDvKPxkoXIHifJOFA",    "add_time": 6,    "last_time": 49}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 6] column is set for unsigned



*/

// AdminUser struct is a row record of the admin_user table in the yx database
type AdminUser struct {
	//[ 0] user_id                                        uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	UserID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:user_id;type:uint;" json:"user_id"` // 用户id
	//[ 1] user_name                                      varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	UserName string `gorm:"column:user_name;type:varchar;size:20;" json:"user_name"` // 用户名
	//[ 2] password                                       varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Password string `gorm:"column:password;type:varchar;size:32;" json:"password"` // 密码
	//[ 3] salt                                           char(6)              null: false  primary: false  isArray: false  auto: false  col: char            len: 6       default: []
	Salt string `gorm:"column:salt;type:char;size:6;" json:"salt"` // 密码盐
	//[ 4] head_img                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	HeadImg string `gorm:"column:head_img;type:varchar;size:255;" json:"head_img"` // 头像
	//[ 5] user_email                                     varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	UserEmail string `gorm:"column:user_email;type:varchar;size:100;" json:"user_email"` // 电子邮箱
	//[ 6] loginfailure                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	Loginfailure uint32 `gorm:"column:loginfailure;type:utinyint;default:0;" json:"loginfailure"` // 失败次数
	//[ 7] login_time                                     bigint               null: true   primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	LoginTime sql.NullInt64 `gorm:"column:login_time;type:bigint;" json:"login_time"` // 登录时间
	//[ 8] token                                          char(32)             null: false  primary: false  isArray: false  auto: false  col: char            len: 32      default: []
	Token string `gorm:"column:token;type:char;size:32;" json:"token"` // token标识
	//[ 9] status                                         tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [1]
	Status int32 `gorm:"column:status;type:tinyint;default:1;" json:"status"` // 状态 1正常 0 关闭
	//[10] login_ip                                       varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LoginIP sql.NullString `gorm:"column:login_ip;type:varchar;size:50;" json:"login_ip"` // 登录IP
	//[11] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;" json:"add_time"` // 创建时间
	//[12] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;" json:"last_time"` // 更新时间

}

var admin_userTableInfo = &TableInfo{
	Name: "admin_user",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "user_id",
			Comment:            `用户id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "UserID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "user_name",
			Comment:            `用户名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "UserName",
			GoFieldType:        "string",
			JSONFieldName:      "user_name",
			ProtobufFieldName:  "user_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "password",
			Comment:            `密码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Password",
			GoFieldType:        "string",
			JSONFieldName:      "password",
			ProtobufFieldName:  "password",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "salt",
			Comment:            `密码盐`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(6)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       6,
			GoFieldName:        "Salt",
			GoFieldType:        "string",
			JSONFieldName:      "salt",
			ProtobufFieldName:  "salt",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "head_img",
			Comment:            `头像`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "HeadImg",
			GoFieldType:        "string",
			JSONFieldName:      "head_img",
			ProtobufFieldName:  "head_img",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "user_email",
			Comment:            `电子邮箱`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "UserEmail",
			GoFieldType:        "string",
			JSONFieldName:      "user_email",
			ProtobufFieldName:  "user_email",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "loginfailure",
			Comment:            `失败次数`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "Loginfailure",
			GoFieldType:        "uint32",
			JSONFieldName:      "loginfailure",
			ProtobufFieldName:  "loginfailure",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "login_time",
			Comment:            `登录时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "LoginTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "login_time",
			ProtobufFieldName:  "login_time",
			ProtobufType:       "int64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "token",
			Comment:            `token标识`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       32,
			GoFieldName:        "Token",
			GoFieldType:        "string",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status",
			Comment:            `状态 1正常 0 关闭`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "int32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "login_ip",
			Comment:            `登录IP`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "LoginIP",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "login_ip",
			ProtobufFieldName:  "login_ip",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "add_time",
			Comment:            `创建时间`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "last_time",
			Comment:            `更新时间`,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdminUser) TableName() string {
	return "admin_user"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdminUser) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdminUser) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdminUser) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdminUser) TableInfo() *TableInfo {
	return admin_userTableInfo
}
