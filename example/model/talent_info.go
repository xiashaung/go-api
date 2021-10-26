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


CREATE TABLE `talent_info` (
  `talent_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '达人账号信息表',
  `talent_mobile` char(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '达人账号信息',
  `talent_passwd` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '达人密码',
  `talent_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '达人名称',
  `talent_level` int(10) NOT NULL DEFAULT '0' COMMENT '达人级别',
  `city_id` int(11) NOT NULL DEFAULT '0' COMMENT '达人所在城市',
  `talent_head` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '达人头像',
  `is_public_contact` tinyint(1) NOT NULL DEFAULT '0' COMMENT '公开联系方式 1是 0否',
  `contact_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人',
  `job_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人职位',
  `contact_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人电话',
  `wechat` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人微信',
  `pe_level` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '开业活动 产品体验官等级 ',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '达人入驻时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_access_time` int(10) NOT NULL DEFAULT '0' COMMENT '最后访问时间',
  `from_type` int(10) NOT NULL DEFAULT '1' COMMENT '来源类型 0 未知 1 pc 2 小程序',
  `register_type` int(10) NOT NULL DEFAULT '0' COMMENT '注册类型 0 未知 1商家 2 达人',
  `talent_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`talent_id`) USING BTREE,
  UNIQUE KEY `idx_mobile` (`talent_mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=286 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='达人账号信息表'

JSON Sample
-------------------------------------
{    "talent_id": 14,    "talent_mobile": "DiwVPogIWNtOUNsybOOflXSJR",    "talent_passwd": "tkiZcTSrvQsqBFBASTMgrrPqZ",    "talent_name": "YHEiiufotkRwPvjrSQWQgtrZi",    "talent_level": 31,    "city_id": 84,    "talent_head": "ZObNgNmCKjUmUkRDSbkicoBuA",    "is_public_contact": 93,    "contact_name": "afiipjYHEVVcUInRcJDybYfcn",    "job_title": "PyPCFjyjQrXcfTiSWOlWbMCga",    "contact_mobile": "eArTZVlIayGcqZoAbnFAyrMtA",    "wechat": "PMgrqSnJCdGHvRbZGBWFFxNfC",    "pe_level": 16,    "add_time": 51,    "last_time": 62,    "last_access_time": 11,    "from_type": 31,    "register_type": 74,    "talent_desc": "QYSPjkkKULbcyPfnuKLhgjERG"}


Comments
-------------------------------------
[12] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned



*/

// TalentInfo struct is a row record of the talent_info table in the yx database
type TalentInfo struct {
	//[ 0] talent_id                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	TalentID int32 `gorm:"primary_key;AUTO_INCREMENT;column:talent_id;type:int;" json:"talent_id"` // 达人账号信息表
	//[ 1] talent_mobile                                  char(11)             null: false  primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	TalentMobile string `gorm:"column:talent_mobile;type:char;size:11;" json:"talent_mobile"` // 达人账号信息
	//[ 2] talent_passwd                                  char(32)             null: false  primary: false  isArray: false  auto: false  col: char            len: 32      default: []
	TalentPasswd string `gorm:"column:talent_passwd;type:char;size:32;" json:"talent_passwd"` // 达人密码
	//[ 3] talent_name                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TalentName string `gorm:"column:talent_name;type:varchar;size:255;" json:"talent_name"` // 达人名称
	//[ 4] talent_level                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentLevel int32 `gorm:"column:talent_level;type:int;default:0;" json:"talent_level"` // 达人级别
	//[ 5] city_id                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CityID int32 `gorm:"column:city_id;type:int;default:0;" json:"city_id"` // 达人所在城市
	//[ 6] talent_head                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TalentHead string `gorm:"column:talent_head;type:varchar;size:255;" json:"talent_head"` // 达人头像
	//[ 7] is_public_contact                              tinyint              null: false  primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	IsPublicContact int32 `gorm:"column:is_public_contact;type:tinyint;default:0;" json:"is_public_contact"` // 公开联系方式 1是 0否
	//[ 8] contact_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactName string `gorm:"column:contact_name;type:varchar;size:255;" json:"contact_name"` // 联系人
	//[ 9] job_title                                      varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	JobTitle string `gorm:"column:job_title;type:varchar;size:255;" json:"job_title"` // 联系人职位
	//[10] contact_mobile                                 varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	ContactMobile string `gorm:"column:contact_mobile;type:varchar;size:20;" json:"contact_mobile"` // 联系人电话
	//[11] wechat                                         varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	Wechat string `gorm:"column:wechat;type:varchar;size:30;" json:"wechat"` // 联系人微信
	//[12] pe_level                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PeLevel uint32 `gorm:"column:pe_level;type:uint;default:0;" json:"pe_level"` // 开业活动 产品体验官等级
	//[13] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 达人入驻时间
	//[14] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
	//[15] last_access_time                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastAccessTime int32 `gorm:"column:last_access_time;type:int;default:0;" json:"last_access_time"` // 最后访问时间
	//[16] from_type                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	FromType int32 `gorm:"column:from_type;type:int;default:1;" json:"from_type"` // 来源类型 0 未知 1 pc 2 小程序
	//[17] register_type                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RegisterType int32 `gorm:"column:register_type;type:int;default:0;" json:"register_type"` // 注册类型 0 未知 1商家 2 达人
	//[18] talent_desc                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TalentDesc string `gorm:"column:talent_desc;type:varchar;size:255;" json:"talent_desc"` // 备注

}

var talent_infoTableInfo = &TableInfo{
	Name: "talent_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "talent_id",
			Comment:            `达人账号信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "talent_mobile",
			Comment:            `达人账号信息`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "TalentMobile",
			GoFieldType:        "string",
			JSONFieldName:      "talent_mobile",
			ProtobufFieldName:  "talent_mobile",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "talent_passwd",
			Comment:            `达人密码`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       32,
			GoFieldName:        "TalentPasswd",
			GoFieldType:        "string",
			JSONFieldName:      "talent_passwd",
			ProtobufFieldName:  "talent_passwd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "talent_name",
			Comment:            `达人名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "TalentName",
			GoFieldType:        "string",
			JSONFieldName:      "talent_name",
			ProtobufFieldName:  "talent_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "talent_level",
			Comment:            `达人级别`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentLevel",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_level",
			ProtobufFieldName:  "talent_level",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "city_id",
			Comment:            `达人所在城市`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CityID",
			GoFieldType:        "int32",
			JSONFieldName:      "city_id",
			ProtobufFieldName:  "city_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "talent_head",
			Comment:            `达人头像`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "TalentHead",
			GoFieldType:        "string",
			JSONFieldName:      "talent_head",
			ProtobufFieldName:  "talent_head",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "is_public_contact",
			Comment:            `公开联系方式 1是 0否`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "IsPublicContact",
			GoFieldType:        "int32",
			JSONFieldName:      "is_public_contact",
			ProtobufFieldName:  "is_public_contact",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "contact_name",
			Comment:            `联系人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ContactName",
			GoFieldType:        "string",
			JSONFieldName:      "contact_name",
			ProtobufFieldName:  "contact_name",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "job_title",
			Comment:            `联系人职位`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "JobTitle",
			GoFieldType:        "string",
			JSONFieldName:      "job_title",
			ProtobufFieldName:  "job_title",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "contact_mobile",
			Comment:            `联系人电话`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "ContactMobile",
			GoFieldType:        "string",
			JSONFieldName:      "contact_mobile",
			ProtobufFieldName:  "contact_mobile",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "wechat",
			Comment:            `联系人微信`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "Wechat",
			GoFieldType:        "string",
			JSONFieldName:      "wechat",
			ProtobufFieldName:  "wechat",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "pe_level",
			Comment:            `开业活动 产品体验官等级 `,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PeLevel",
			GoFieldType:        "uint32",
			JSONFieldName:      "pe_level",
			ProtobufFieldName:  "pe_level",
			ProtobufType:       "uint32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "add_time",
			Comment:            `达人入驻时间`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "last_access_time",
			Comment:            `最后访问时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LastAccessTime",
			GoFieldType:        "int32",
			JSONFieldName:      "last_access_time",
			ProtobufFieldName:  "last_access_time",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "from_type",
			Comment:            `来源类型 0 未知 1 pc 2 小程序`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FromType",
			GoFieldType:        "int32",
			JSONFieldName:      "from_type",
			ProtobufFieldName:  "from_type",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "register_type",
			Comment:            `注册类型 0 未知 1商家 2 达人`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RegisterType",
			GoFieldType:        "int32",
			JSONFieldName:      "register_type",
			ProtobufFieldName:  "register_type",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "talent_desc",
			Comment:            `备注`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "TalentDesc",
			GoFieldType:        "string",
			JSONFieldName:      "talent_desc",
			ProtobufFieldName:  "talent_desc",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TalentInfo) TableName() string {
	return "talent_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TalentInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TalentInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TalentInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TalentInfo) TableInfo() *TableInfo {
	return talent_infoTableInfo
}
