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


CREATE TABLE `anchor_info` (
  `anchor_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主播信息表',
  `talent_id` int(11) NOT NULL DEFAULT '0' COMMENT '达人id',
  `platform_id` int(11) NOT NULL DEFAULT '0' COMMENT '所属平台 1：抖音 2：快手',
  `third_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三方平台id',
  `anchor_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '主播账号名',
  `sex` int(1) NOT NULL DEFAULT '0' COMMENT '性别 1 男 2 女',
  `head` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `access_token` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'access_token',
  `expires_in` int(10) NOT NULL DEFAULT '0' COMMENT 'access_token 过期时间，需要定时更新',
  `refresh_token` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'refresh_token',
  `refresh_expires_in` int(10) NOT NULL DEFAULT '0' COMMENT 'refresh_token 过期时间，需要定时刷新',
  `open_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三方 open_id',
  `auth_status` int(1) NOT NULL DEFAULT '1' COMMENT '授权状态，1 正常  0 已过期',
  `anchor_uid` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主播uid信息 对应抖音的作者id',
  `anchor_unionid` varchar(60) DEFAULT NULL COMMENT '主播的联盟平台id，抖音为百应id',
  `anchor_account` varchar(60) DEFAULT NULL COMMENT '主播第三方平台账号 抖音号或者快手号',
  `window_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '主播橱窗id',
  `contact_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '橱窗联系人姓名',
  `contact_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '橱窗联系人电话',
  `city_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '主播所在城市',
  `contact_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `fans_num` int(10) NOT NULL DEFAULT '-1' COMMENT '粉丝数量',
  `follow` int(10) NOT NULL DEFAULT '0' COMMENT '关注数',
  `praise_score` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '带货口碑分 扩大100倍存储',
  `sale_num` varchar(60) NOT NULL DEFAULT '' COMMENT '总体销量',
  `sync_order_time` int(10) NOT NULL DEFAULT '0' COMMENT '同步订单时间(日期时间戳)',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '主播授权时间',
  `last_time` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`anchor_id`) USING BTREE,
  KEY `idx_open_id` (`open_id`) USING BTREE,
  KEY `idx_talent_id` (`talent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=729 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='主播信息表'

JSON Sample
-------------------------------------
{    "anchor_id": 83,    "talent_id": 34,    "platform_id": 74,    "third_id": "tRJwWTIEwovtZaBcScHvScWvB",    "anchor_name": "DccRmtPyGNEupeAKoPIkIyIeq",    "sex": 8,    "head": "YXSLPMOSZpmsYgbVtkTYaYjoh",    "access_token": "KXSGxqvjsGaSpcxyymWTZgidM",    "expires_in": 10,    "refresh_token": "uEPLSFaeTWmFusvqcsRFlCZZt",    "refresh_expires_in": 8,    "open_id": "huphufZFdiqXdPykcbgfkGERA",    "auth_status": 35,    "anchor_uid": "JkeGhXtGYSGBFqecUjSNoXues",    "anchor_unionid": "BuUbghijMfNWRvbxcvSmomSXD",    "anchor_account": "rRtEqKvwiIeBwXYYOctRnpliK",    "window_id": "RSPeimVOEMypNQLUqDsvnavyu",    "contact_name": "IhhygAYVhkxesJZRYFKjxBaKn",    "contact_mobile": "sHRJEOQrkSGUAdvMqwthTOhhr",    "city_id": 36,    "contact_addr": "BcBEIfneqeyjorvHKROSjvIdn",    "fans_num": 66,    "follow": 54,    "praise_score": 88,    "sale_num": "ksbOZmlpOVGcygCEXXnWdRpkZ",    "sync_order_time": 66,    "add_time": 42,    "last_time": 17}


Comments
-------------------------------------
[19] column is set for unsigned
[23] column is set for unsigned
[26] column is set for unsigned
[27] column is set for unsigned



*/

// AnchorInfo struct is a row record of the anchor_info table in the yx database
type AnchorInfo struct {
	//[ 0] anchor_id                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	AnchorID int32 `gorm:"primary_key;AUTO_INCREMENT;column:anchor_id;type:int;" json:"anchor_id"` // 主播信息表
	//[ 1] talent_id                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TalentID int32 `gorm:"column:talent_id;type:int;default:0;" json:"talent_id"` // 达人id
	//[ 2] platform_id                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	PlatformID int32 `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"` // 所属平台 1：抖音 2：快手
	//[ 3] third_id                                       varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ThirdID string `gorm:"column:third_id;type:varchar;size:255;" json:"third_id"` // 第三方平台id
	//[ 4] anchor_name                                    varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	AnchorName string `gorm:"column:anchor_name;type:varchar;size:255;" json:"anchor_name"` // 主播账号名
	//[ 5] sex                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Sex int32 `gorm:"column:sex;type:int;default:0;" json:"sex"` // 性别 1 男 2 女
	//[ 6] head                                           varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Head string `gorm:"column:head;type:varchar;size:255;" json:"head"` // 头像
	//[ 7] access_token                                   varchar(1000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1000    default: []
	AccessToken string `gorm:"column:access_token;type:varchar;size:1000;" json:"access_token"` // access_token
	//[ 8] expires_in                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ExpiresIn int32 `gorm:"column:expires_in;type:int;default:0;" json:"expires_in"` // access_token 过期时间，需要定时更新
	//[ 9] refresh_token                                  varchar(1000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 1000    default: []
	RefreshToken string `gorm:"column:refresh_token;type:varchar;size:1000;" json:"refresh_token"` // refresh_token
	//[10] refresh_expires_in                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RefreshExpiresIn int32 `gorm:"column:refresh_expires_in;type:int;default:0;" json:"refresh_expires_in"` // refresh_token 过期时间，需要定时刷新
	//[11] open_id                                        varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	OpenID string `gorm:"column:open_id;type:varchar;size:255;" json:"open_id"` // 第三方 open_id
	//[12] auth_status                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	AuthStatus int32 `gorm:"column:auth_status;type:int;default:1;" json:"auth_status"` // 授权状态，1 正常  0 已过期
	//[13] anchor_uid                                     varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	AnchorUID string `gorm:"column:anchor_uid;type:varchar;size:60;" json:"anchor_uid"` // 主播uid信息 对应抖音的作者id
	//[14] anchor_unionid                                 varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	AnchorUnionid sql.NullString `gorm:"column:anchor_unionid;type:varchar;size:60;" json:"anchor_unionid"` // 主播的联盟平台id，抖音为百应id
	//[15] anchor_account                                 varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	AnchorAccount sql.NullString `gorm:"column:anchor_account;type:varchar;size:60;" json:"anchor_account"` // 主播第三方平台账号 抖音号或者快手号
	//[16] window_id                                      varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	WindowID string `gorm:"column:window_id;type:varchar;size:30;" json:"window_id"` // 主播橱窗id
	//[17] contact_name                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactName string `gorm:"column:contact_name;type:varchar;size:255;" json:"contact_name"` // 橱窗联系人姓名
	//[18] contact_mobile                                 varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	ContactMobile string `gorm:"column:contact_mobile;type:varchar;size:20;" json:"contact_mobile"` // 橱窗联系人电话
	//[19] city_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CityID uint32 `gorm:"column:city_id;type:uint;default:0;" json:"city_id"` // 主播所在城市
	//[20] contact_addr                                   varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	ContactAddr string `gorm:"column:contact_addr;type:varchar;size:255;" json:"contact_addr"` // 联系地址
	//[21] fans_num                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [-1]
	FansNum int32 `gorm:"column:fans_num;type:int;default:-1;" json:"fans_num"` // 粉丝数量
	//[22] follow                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	Follow int32 `gorm:"column:follow;type:int;default:0;" json:"follow"` // 关注数
	//[23] praise_score                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PraiseScore uint32 `gorm:"column:praise_score;type:uint;default:0;" json:"praise_score"` // 带货口碑分 扩大100倍存储
	//[24] sale_num                                       varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	SaleNum string `gorm:"column:sale_num;type:varchar;size:60;" json:"sale_num"` // 总体销量
	//[25] sync_order_time                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SyncOrderTime int32 `gorm:"column:sync_order_time;type:int;default:0;" json:"sync_order_time"` // 同步订单时间(日期时间戳)
	//[26] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 主播授权时间
	//[27] last_time                                      uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime sql.NullInt64 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var anchor_infoTableInfo = &TableInfo{
	Name: "anchor_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "anchor_id",
			Comment:            `主播信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AnchorID",
			GoFieldType:        "int32",
			JSONFieldName:      "anchor_id",
			ProtobufFieldName:  "anchor_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "talent_id",
			Comment:            `达人id`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "platform_id",
			Comment:            `所属平台 1：抖音 2：快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "int32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "third_id",
			Comment:            `第三方平台id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ThirdID",
			GoFieldType:        "string",
			JSONFieldName:      "third_id",
			ProtobufFieldName:  "third_id",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "anchor_name",
			Comment:            `主播账号名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "AnchorName",
			GoFieldType:        "string",
			JSONFieldName:      "anchor_name",
			ProtobufFieldName:  "anchor_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "sex",
			Comment:            `性别 1 男 2 女`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Sex",
			GoFieldType:        "int32",
			JSONFieldName:      "sex",
			ProtobufFieldName:  "sex",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "head",
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
			GoFieldName:        "Head",
			GoFieldType:        "string",
			JSONFieldName:      "head",
			ProtobufFieldName:  "head",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "access_token",
			Comment:            `access_token`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1000,
			GoFieldName:        "AccessToken",
			GoFieldType:        "string",
			JSONFieldName:      "access_token",
			ProtobufFieldName:  "access_token",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "expires_in",
			Comment:            `access_token 过期时间，需要定时更新`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ExpiresIn",
			GoFieldType:        "int32",
			JSONFieldName:      "expires_in",
			ProtobufFieldName:  "expires_in",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "refresh_token",
			Comment:            `refresh_token`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1000,
			GoFieldName:        "RefreshToken",
			GoFieldType:        "string",
			JSONFieldName:      "refresh_token",
			ProtobufFieldName:  "refresh_token",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "refresh_expires_in",
			Comment:            `refresh_token 过期时间，需要定时刷新`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RefreshExpiresIn",
			GoFieldType:        "int32",
			JSONFieldName:      "refresh_expires_in",
			ProtobufFieldName:  "refresh_expires_in",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "open_id",
			Comment:            `第三方 open_id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "OpenID",
			GoFieldType:        "string",
			JSONFieldName:      "open_id",
			ProtobufFieldName:  "open_id",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "auth_status",
			Comment:            `授权状态，1 正常  0 已过期`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AuthStatus",
			GoFieldType:        "int32",
			JSONFieldName:      "auth_status",
			ProtobufFieldName:  "auth_status",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "anchor_uid",
			Comment:            `主播uid信息 对应抖音的作者id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "AnchorUID",
			GoFieldType:        "string",
			JSONFieldName:      "anchor_uid",
			ProtobufFieldName:  "anchor_uid",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "anchor_unionid",
			Comment:            `主播的联盟平台id，抖音为百应id`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "AnchorUnionid",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anchor_unionid",
			ProtobufFieldName:  "anchor_unionid",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "anchor_account",
			Comment:            `主播第三方平台账号 抖音号或者快手号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "AnchorAccount",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anchor_account",
			ProtobufFieldName:  "anchor_account",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "window_id",
			Comment:            `主播橱窗id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "WindowID",
			GoFieldType:        "string",
			JSONFieldName:      "window_id",
			ProtobufFieldName:  "window_id",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "contact_name",
			Comment:            `橱窗联系人姓名`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "contact_mobile",
			Comment:            `橱窗联系人电话`,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "city_id",
			Comment:            `主播所在城市`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CityID",
			GoFieldType:        "uint32",
			JSONFieldName:      "city_id",
			ProtobufFieldName:  "city_id",
			ProtobufType:       "uint32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "contact_addr",
			Comment:            `联系地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "ContactAddr",
			GoFieldType:        "string",
			JSONFieldName:      "contact_addr",
			ProtobufFieldName:  "contact_addr",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "fans_num",
			Comment:            `粉丝数量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FansNum",
			GoFieldType:        "int32",
			JSONFieldName:      "fans_num",
			ProtobufFieldName:  "fans_num",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "follow",
			Comment:            `关注数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "Follow",
			GoFieldType:        "int32",
			JSONFieldName:      "follow",
			ProtobufFieldName:  "follow",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "praise_score",
			Comment:            `带货口碑分 扩大100倍存储`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PraiseScore",
			GoFieldType:        "uint32",
			JSONFieldName:      "praise_score",
			ProtobufFieldName:  "praise_score",
			ProtobufType:       "uint32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "sale_num",
			Comment:            `总体销量`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "SaleNum",
			GoFieldType:        "string",
			JSONFieldName:      "sale_num",
			ProtobufFieldName:  "sale_num",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "sync_order_time",
			Comment:            `同步订单时间(日期时间戳)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SyncOrderTime",
			GoFieldType:        "int32",
			JSONFieldName:      "sync_order_time",
			ProtobufFieldName:  "sync_order_time",
			ProtobufType:       "int32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "add_time",
			Comment:            `主播授权时间`,
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
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "last_time",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "LastTime",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "last_time",
			ProtobufFieldName:  "last_time",
			ProtobufType:       "uint32",
			ProtobufPos:        28,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AnchorInfo) TableName() string {
	return "anchor_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AnchorInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AnchorInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AnchorInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AnchorInfo) TableInfo() *TableInfo {
	return anchor_infoTableInfo
}
