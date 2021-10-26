package example

import (
	"database/sql"
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `ad_info` (
  `ad_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '广告信息表',
  `site_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '对应的广告位置 1:官网首页头部轮播图 2:高佣专区 3：每日带货',
  `platform_id` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '所属平台 1:央选 2：抖音 3：快手',
  `product_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `ad_name` varchar(60) NOT NULL COMMENT '广告标题',
  `ad_pic` varchar(300) NOT NULL COMMENT '广告图片',
  `ad_url` varchar(300) NOT NULL COMMENT '跳转链接地址',
  `ad_weight` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '权重',
  `ad_desc` text COMMENT '广告描述',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加广告时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`ad_id`),
  KEY `idx_weight_siteid` (`ad_weight`,`site_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='广告信息表'

JSON Sample
-------------------------------------
{    "ad_id": 41,    "site_id": 14,    "platform_id": 5,    "product_id": 79,    "ad_name": "tqbowpgdlkaZcKCMSkeqvndEO",    "ad_pic": "VVDaSqfifkMBPARHPuyFKPknR",    "ad_url": "rRKJjjyilVjPIGRUyRCqGpFbP",    "ad_weight": 93,    "ad_desc": "frXAeHJCjnTFPJgtHnfMoXtcA",    "add_time": 69,    "last_time": 5}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 3] column is set for unsigned
[ 7] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned



*/

// AdInfo struct is a row record of the ad_info table in the yx database
type AdInfo struct {
	//[ 0] ad_id                                          uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	AdID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:ad_id;type:uint;" json:"ad_id"` // 广告信息表
	//[ 1] site_id                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	SiteID uint32 `gorm:"column:site_id;type:uint;default:0;" json:"site_id"` // 对应的广告位置 1:官网首页头部轮播图 2:高佣专区 3：每日带货
	//[ 2] platform_id                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	PlatformID uint32 `gorm:"column:platform_id;type:uint;default:1;" json:"platform_id"` // 所属平台 1:央选 2：抖音 3：快手
	//[ 3] product_id                                     ubigint              null: false  primary: false  isArray: false  auto: false  col: ubigint         len: -1      default: [0]
	ProductID uint64 `gorm:"column:product_id;type:ubigint;default:0;" json:"product_id"` // 商品id
	//[ 4] ad_name                                        varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	AdName string `gorm:"column:ad_name;type:varchar;size:60;" json:"ad_name"` // 广告标题
	//[ 5] ad_pic                                         varchar(300)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 300     default: []
	AdPic string `gorm:"column:ad_pic;type:varchar;size:300;" json:"ad_pic"` // 广告图片
	//[ 6] ad_url                                         varchar(300)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 300     default: []
	AdURL string `gorm:"column:ad_url;type:varchar;size:300;" json:"ad_url"` // 跳转链接地址
	//[ 7] ad_weight                                      utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	AdWeight uint32 `gorm:"column:ad_weight;type:utinyint;default:0;" json:"ad_weight"` // 权重
	//[ 8] ad_desc                                        text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	AdDesc sql.NullString `gorm:"column:ad_desc;type:text;size:65535;" json:"ad_desc"` // 广告描述
	//[ 9] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 添加广告时间
	//[10] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var ad_infoTableInfo = &TableInfo{
	Name: "ad_info",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ad_id",
			Comment:            `广告信息表`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AdID",
			GoFieldType:        "uint32",
			JSONFieldName:      "ad_id",
			ProtobufFieldName:  "ad_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "site_id",
			Comment:            `对应的广告位置 1:官网首页头部轮播图 2:高佣专区 3：每日带货`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "SiteID",
			GoFieldType:        "uint32",
			JSONFieldName:      "site_id",
			ProtobufFieldName:  "site_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "platform_id",
			Comment:            `所属平台 1:央选 2：抖音 3：快手`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PlatformID",
			GoFieldType:        "uint32",
			JSONFieldName:      "platform_id",
			ProtobufFieldName:  "platform_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "product_id",
			Comment:            `商品id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ProductID",
			GoFieldType:        "uint64",
			JSONFieldName:      "product_id",
			ProtobufFieldName:  "product_id",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "ad_name",
			Comment:            `广告标题`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "AdName",
			GoFieldType:        "string",
			JSONFieldName:      "ad_name",
			ProtobufFieldName:  "ad_name",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "ad_pic",
			Comment:            `广告图片`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       300,
			GoFieldName:        "AdPic",
			GoFieldType:        "string",
			JSONFieldName:      "ad_pic",
			ProtobufFieldName:  "ad_pic",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "ad_url",
			Comment:            `跳转链接地址`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       300,
			GoFieldName:        "AdURL",
			GoFieldType:        "string",
			JSONFieldName:      "ad_url",
			ProtobufFieldName:  "ad_url",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "ad_weight",
			Comment:            `权重`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "AdWeight",
			GoFieldType:        "uint32",
			JSONFieldName:      "ad_weight",
			ProtobufFieldName:  "ad_weight",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "ad_desc",
			Comment:            `广告描述`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "AdDesc",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ad_desc",
			ProtobufFieldName:  "ad_desc",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "add_time",
			Comment:            `添加广告时间`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AdInfo) TableName() string {
	return "ad_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AdInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AdInfo) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AdInfo) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AdInfo) TableInfo() *TableInfo {
	return ad_infoTableInfo
}
