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


CREATE TABLE `open_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `third_type` smallint(6) NOT NULL DEFAULT '1' COMMENT ' 微信类型 1 小程序 2 公众号',
  `app_id` varchar(100) NOT NULL DEFAULT '' COMMENT '小程序/公众号app_id',
  `merchant_id` int(11) NOT NULL DEFAULT '0' COMMENT '商家账号',
  `talent_id` int(11) NOT NULL DEFAULT '0' COMMENT '账号id',
  `openid` varchar(100) NOT NULL DEFAULT '' COMMENT '用户openid',
  `nick_name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `gender` smallint(6) NOT NULL DEFAULT '3' COMMENT '性别 1 男 2 女 3 未知',
  `province` varchar(10) NOT NULL DEFAULT '' COMMENT '省',
  `city` varchar(10) NOT NULL DEFAULT '' COMMENT '市',
  `country` varchar(20) NOT NULL DEFAULT '' COMMENT '国家',
  `avatar_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `union_id` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '联合id',
  `add_time` int(11) DEFAULT '0',
  `last_time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8 COMMENT='微信授权/公众号信息'

JSON Sample
-------------------------------------
{    "id": 71,    "third_type": 57,    "app_id": "tKrxUncinjQDTjsakkqEMmRIC",    "merchant_id": 12,    "talent_id": 3,    "openid": "XWHGvNPrypCyUMwdSDqcJVBCO",    "nick_name": "CdhdZxayJxKNjpLrEfyqwlTPE",    "gender": 10,    "province": "BRICnLNntCJBDAROicJKevLaA",    "city": "LAYLlwfLdrUqNVgawEfEuavWS",    "country": "nBnUxGfCjtxkItAbbVMYgdfFJ",    "avatar_url": "JgJSNniknHYiAYLPcqfGTVXmY",    "union_id": "EuaPMeeodKwtMqVJaVlcOBjDN",    "add_time": 89,    "last_time": 39}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// OpenInfo struct is a row record of the open_info table in the yx database
type OpenInfo struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"`
	//[ 1] third_type                                     smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [1]
	ThirdType int32 `gorm:"column:third_type;type:smallint;default:1;" json:"third_type"` //  微信类型 1 小程序 2 公众号
	//[ 2] app_id                                         varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	AppID string `gorm:"column:app_id;type:varchar;size:100;" json:"app_id"` // 小程序/公众号app_id
	//[ 3] merchant_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	MerchantID int32 `gorm:"column:merchant_id;type:int;default:0;" json:"merchant_id"` // 商家账号
	//[ 4] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"` // 账号id
	//[ 5] openid                                         varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Openid string `gorm:"column:openid;type:varchar;size:100;" json:"openid"` // 用户openid
	//[ 6] nick_name                                      varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	NickName string `gorm:"column:nick_name;type:varchar;size:20;" json:"nick_name"` // 用户昵称
	//[ 7] gender                                         smallint             null: false  primary: false  isArray: false  auto: false  col: smallint        len: -1      default: [3]
	Gender int32 `gorm:"column:gender;type:smallint;default:3;" json:"gender"` // 性别 1 男 2 女 3 未知
	//[ 8] province                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	Province string `gorm:"column:province;type:varchar;size:10;" json:"province"` // 省
	//[ 9] city                                           varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	City string `gorm:"column:city;type:varchar;size:10;" json:"city"` // 市
	//[10] country                                        varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Country string `gorm:"column:country;type:varchar;size:20;" json:"country"` // 国家
	//[11] avatar_url                                     varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AvatarURL string `gorm:"column:avatar_url;type:varchar;size:255;" json:"avatar_url"` // 头像
	//[12] union_id                                       varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	UnionID string `gorm:"column:union_id;type:varchar;size:200;" json:"union_id"` // 联合id
	//[13] add_time                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AddTime sql.NullInt64 `gorm:"column:add_time;type:int;default:0;" json:"add_time"`
	//[14] last_time                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LastTime sql.NullInt64 `gorm:"column:last_time;type:int;default:0;" json:"last_time"`
}

var open_infoTableInfo = &TableInfo{
	Name: "open_info",
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
			Name:               "third_type",
			Comment:            ` 微信类型 1 小程序 2 公众号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "ThirdType",
			GoFieldType:        "int32",
			JSONFieldName:      "third_type",
			ProtobufFieldName:  "third_type",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "app_id",
			Comment:            `小程序/公众号app_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "AppID",
			GoFieldType:        "string",
			JSONFieldName:      "app_id",
			ProtobufFieldName:  "app_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "merchant_id",
			Comment:            `商家账号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MerchantID",
			GoFieldType:        "int32",
			JSONFieldName:      "merchant_id",
			ProtobufFieldName:  "merchant_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "talent_id",
			Comment:            `账号id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TalentID",
			GoFieldType:        "int32",
			JSONFieldName:      "talent_id",
			ProtobufFieldName:  "talent_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "openid",
			Comment:            `用户openid`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Openid",
			GoFieldType:        "string",
			JSONFieldName:      "openid",
			ProtobufFieldName:  "openid",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "nick_name",
			Comment:            `用户昵称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "NickName",
			GoFieldType:        "string",
			JSONFieldName:      "nick_name",
			ProtobufFieldName:  "nick_name",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "gender",
			Comment:            `性别 1 男 2 女 3 未知`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "smallint",
			DatabaseTypePretty: "smallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "smallint",
			ColumnLength:       -1,
			GoFieldName:        "Gender",
			GoFieldType:        "int32",
			JSONFieldName:      "gender",
			ProtobufFieldName:  "gender",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "province",
			Comment:            `省`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "Province",
			GoFieldType:        "string",
			JSONFieldName:      "province",
			ProtobufFieldName:  "province",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "city",
			Comment:            `市`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "City",
			GoFieldType:        "string",
			JSONFieldName:      "city",
			ProtobufFieldName:  "city",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "country",
			Comment:            `国家`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Country",
			GoFieldType:        "string",
			JSONFieldName:      "country",
			ProtobufFieldName:  "country",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "avatar_url",
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
			GoFieldName:        "AvatarURL",
			GoFieldType:        "string",
			JSONFieldName:      "avatar_url",
			ProtobufFieldName:  "avatar_url",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "union_id",
			Comment:            `联合id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "UnionID",
			GoFieldType:        "string",
			JSONFieldName:      "union_id",
			ProtobufFieldName:  "union_id",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (o *OpenInfo) TableName() string {
	return "open_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OpenInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OpenInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OpenInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OpenInfo) TableInfo() *TableInfo {
	return open_infoTableInfo
}
