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


CREATE TABLE `prop_ks` (
  `prop_id` bigint(20) NOT NULL COMMENT '快手prop信息表',
  `prop_value_pid` bigint(20) NOT NULL COMMENT '依赖的属性值id',
  `prop_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'prop名称',
  `input_type` int(20) unsigned NOT NULL DEFAULT '0' COMMENT '输入类型',
  `prop_type` int(255) unsigned NOT NULL DEFAULT '1' COMMENT '属性类型 1：sku 2：类目',
  `prop_required` int(255) NOT NULL COMMENT '是否必需 0：否 1：是 ',
  `prop_value_maximum` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '最大可选数量',
  `sort_num` int(11) DEFAULT NULL COMMENT '排序顺序',
  `is_leaf` int(255) unsigned NOT NULL DEFAULT '0' COMMENT '是否叶子节点',
  `status` int(255) unsigned NOT NULL DEFAULT '0' COMMENT 'prop状态',
  `date_type` varchar(20) DEFAULT NULL COMMENT '日期类型，决定range_list里面的日期格式',
  `date_rangelist` varchar(60) DEFAULT NULL COMMENT '日期字符串 下标0代表起始时间 1代表结束时间',
  `text_patternlist` varchar(60) DEFAULT NULL COMMENT '字符串匹配模式，正则语法',
  `text_uselimit` int(255) DEFAULT '0' COMMENT '是否有长度限制 0：没有 1：有',
  `text_max` int(255) unsigned DEFAULT NULL COMMENT '字符串最大长度',
  `text_min` int(255) unsigned DEFAULT NULL COMMENT '字符串最小长度',
  `text_message` varchar(100) DEFAULT NULL COMMENT '错误信息',
  `image_mincount` int(255) DEFAULT NULL COMMENT '图片最小数量',
  `image_maxcount` int(255) DEFAULT NULL COMMENT '图片最大数量',
  `image_minwidth` int(11) DEFAULT NULL COMMENT '最小宽度',
  `image_height` int(11) DEFAULT NULL COMMENT '最大宽度',
  `prop_inputtype` varchar(20) DEFAULT NULL COMMENT '输入类型。TEXT：文本，CHECKBOX：多选，NUMBER：数字，EMAIL：email，DATETIME：时间，URL：链接，DATETIMERANGE：时间段，RADIO：单选，IMAGE：图片，INVALID_PROP_INPUT_TYPE：无效值',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`prop_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

JSON Sample
-------------------------------------
{    "prop_id": 7,    "prop_value_pid": 85,    "prop_name": "TXSCGJtLcTVwWyUneFfBXSEVh",    "input_type": 62,    "prop_type": 73,    "prop_required": 62,    "prop_value_maximum": 59,    "sort_num": 19,    "is_leaf": 84,    "status": 84,    "date_type": "CityPqZdWXYFyQOLciYmWtXjG",    "date_rangelist": "FTIiHhJfBUBnuEhpcGRNMIdLY",    "text_patternlist": "goFbCJhTnJpmFIILQnGiAfVLq",    "text_uselimit": 23,    "text_max": 68,    "text_min": 48,    "text_message": "SRXkxAdYZmEqqkHbjOmpTAGqp",    "image_mincount": 18,    "image_maxcount": 77,    "image_minwidth": 39,    "image_height": 49,    "prop_inputtype": "jRHHEZPclpVtgTPSrpjNDwGRI",    "add_time": 38,    "last_time": 75}


Comments
-------------------------------------
[ 3] column is set for unsigned
[ 4] column is set for unsigned
[ 6] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned
[22] column is set for unsigned
[23] column is set for unsigned



*/

// PropKs struct is a row record of the prop_ks table in the yx database
type PropKs struct {
	//[ 0] prop_id                                        bigint               null: false  primary: true   isArray: false  auto: false  col: bigint          len: -1      default: []
	PropID int64 `gorm:"primary_key;column:prop_id;type:bigint;" json:"prop_id"` // 快手prop信息表
	//[ 1] prop_value_pid                                 bigint               null: false  primary: false  isArray: false  auto: false  col: bigint          len: -1      default: []
	PropValuePid int64 `gorm:"column:prop_value_pid;type:bigint;" json:"prop_value_pid"` // 依赖的属性值id
	//[ 2] prop_name                                      varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	PropName string `gorm:"column:prop_name;type:varchar;size:20;" json:"prop_name"` // prop名称
	//[ 3] input_type                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	InputType uint32 `gorm:"column:input_type;type:uint;default:0;" json:"input_type"` // 输入类型
	//[ 4] prop_type                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [1]
	PropType uint32 `gorm:"column:prop_type;type:uint;default:1;" json:"prop_type"` // 属性类型 1：sku 2：类目
	//[ 5] prop_required                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PropRequired int32 `gorm:"column:prop_required;type:int;" json:"prop_required"` // 是否必需 0：否 1：是
	//[ 6] prop_value_maximum                             uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	PropValueMaximum uint32 `gorm:"column:prop_value_maximum;type:uint;default:0;" json:"prop_value_maximum"` // 最大可选数量
	//[ 7] sort_num                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SortNum sql.NullInt64 `gorm:"column:sort_num;type:int;" json:"sort_num"` // 排序顺序
	//[ 8] is_leaf                                        uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	IsLeaf uint32 `gorm:"column:is_leaf;type:uint;default:0;" json:"is_leaf"` // 是否叶子节点
	//[ 9] status                                         uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	Status uint32 `gorm:"column:status;type:uint;default:0;" json:"status"` // prop状态
	//[10] date_type                                      varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	DateType sql.NullString `gorm:"column:date_type;type:varchar;size:20;" json:"date_type"` // 日期类型，决定range_list里面的日期格式
	//[11] date_rangelist                                 varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	DateRangelist sql.NullString `gorm:"column:date_rangelist;type:varchar;size:60;" json:"date_rangelist"` // 日期字符串 下标0代表起始时间 1代表结束时间
	//[12] text_patternlist                               varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	TextPatternlist sql.NullString `gorm:"column:text_patternlist;type:varchar;size:60;" json:"text_patternlist"` // 字符串匹配模式，正则语法
	//[13] text_uselimit                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TextUselimit sql.NullInt64 `gorm:"column:text_uselimit;type:int;default:0;" json:"text_uselimit"` // 是否有长度限制 0：没有 1：有
	//[14] text_max                                       uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TextMax sql.NullInt64 `gorm:"column:text_max;type:uint;" json:"text_max"` // 字符串最大长度
	//[15] text_min                                       uint                 null: true   primary: false  isArray: false  auto: false  col: uint            len: -1      default: []
	TextMin sql.NullInt64 `gorm:"column:text_min;type:uint;" json:"text_min"` // 字符串最小长度
	//[16] text_message                                   varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	TextMessage sql.NullString `gorm:"column:text_message;type:varchar;size:100;" json:"text_message"` // 错误信息
	//[17] image_mincount                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ImageMincount sql.NullInt64 `gorm:"column:image_mincount;type:int;" json:"image_mincount"` // 图片最小数量
	//[18] image_maxcount                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ImageMaxcount sql.NullInt64 `gorm:"column:image_maxcount;type:int;" json:"image_maxcount"` // 图片最大数量
	//[19] image_minwidth                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ImageMinwidth sql.NullInt64 `gorm:"column:image_minwidth;type:int;" json:"image_minwidth"` // 最小宽度
	//[20] image_height                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ImageHeight sql.NullInt64 `gorm:"column:image_height;type:int;" json:"image_height"` // 最大宽度
	//[21] prop_inputtype                                 varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	PropInputtype sql.NullString `gorm:"column:prop_inputtype;type:varchar;size:20;" json:"prop_inputtype"` // 输入类型。TEXT：文本，CHECKBOX：多选，NUMBER：数字，EMAIL：email，DATETIME：时间，URL：链接，DATETIMERANGE：时间段，RADIO：单选，IMAGE：图片，INVALID_PROP_INPUT_TYPE：无效值
	//[22] add_time                                       uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AddTime uint32 `gorm:"column:add_time;type:uint;default:0;" json:"add_time"` // 创建时间
	//[23] last_time                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LastTime uint32 `gorm:"column:last_time;type:uint;default:0;" json:"last_time"`
}

var prop_ksTableInfo = &TableInfo{
	Name: "prop_ks",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "prop_id",
			Comment:            `快手prop信息表`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PropID",
			GoFieldType:        "int64",
			JSONFieldName:      "prop_id",
			ProtobufFieldName:  "prop_id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "prop_value_pid",
			Comment:            `依赖的属性值id`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "PropValuePid",
			GoFieldType:        "int64",
			JSONFieldName:      "prop_value_pid",
			ProtobufFieldName:  "prop_value_pid",
			ProtobufType:       "int64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "prop_name",
			Comment:            `prop名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "PropName",
			GoFieldType:        "string",
			JSONFieldName:      "prop_name",
			ProtobufFieldName:  "prop_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "input_type",
			Comment:            `输入类型`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "InputType",
			GoFieldType:        "uint32",
			JSONFieldName:      "input_type",
			ProtobufFieldName:  "input_type",
			ProtobufType:       "uint32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "prop_type",
			Comment:            `属性类型 1：sku 2：类目`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PropType",
			GoFieldType:        "uint32",
			JSONFieldName:      "prop_type",
			ProtobufFieldName:  "prop_type",
			ProtobufType:       "uint32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "prop_required",
			Comment:            `是否必需 0：否 1：是 `,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PropRequired",
			GoFieldType:        "int32",
			JSONFieldName:      "prop_required",
			ProtobufFieldName:  "prop_required",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "prop_value_maximum",
			Comment:            `最大可选数量`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PropValueMaximum",
			GoFieldType:        "uint32",
			JSONFieldName:      "prop_value_maximum",
			ProtobufFieldName:  "prop_value_maximum",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "sort_num",
			Comment:            `排序顺序`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SortNum",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "sort_num",
			ProtobufFieldName:  "sort_num",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "is_leaf",
			Comment:            `是否叶子节点`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "IsLeaf",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_leaf",
			ProtobufFieldName:  "is_leaf",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "status",
			Comment:            `prop状态`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Status",
			GoFieldType:        "uint32",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "date_type",
			Comment:            `日期类型，决定range_list里面的日期格式`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "DateType",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "date_type",
			ProtobufFieldName:  "date_type",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "date_rangelist",
			Comment:            `日期字符串 下标0代表起始时间 1代表结束时间`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "DateRangelist",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "date_rangelist",
			ProtobufFieldName:  "date_rangelist",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "text_patternlist",
			Comment:            `字符串匹配模式，正则语法`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "TextPatternlist",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "text_patternlist",
			ProtobufFieldName:  "text_patternlist",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "text_uselimit",
			Comment:            `是否有长度限制 0：没有 1：有`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TextUselimit",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "text_uselimit",
			ProtobufFieldName:  "text_uselimit",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "text_max",
			Comment:            `字符串最大长度`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TextMax",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "text_max",
			ProtobufFieldName:  "text_max",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "text_min",
			Comment:            `字符串最小长度`,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "TextMin",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "text_min",
			ProtobufFieldName:  "text_min",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "text_message",
			Comment:            `错误信息`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "TextMessage",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "text_message",
			ProtobufFieldName:  "text_message",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "image_mincount",
			Comment:            `图片最小数量`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ImageMincount",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "image_mincount",
			ProtobufFieldName:  "image_mincount",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "image_maxcount",
			Comment:            `图片最大数量`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ImageMaxcount",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "image_maxcount",
			ProtobufFieldName:  "image_maxcount",
			ProtobufType:       "int32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "image_minwidth",
			Comment:            `最小宽度`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ImageMinwidth",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "image_minwidth",
			ProtobufFieldName:  "image_minwidth",
			ProtobufType:       "int32",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "image_height",
			Comment:            `最大宽度`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ImageHeight",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "image_height",
			ProtobufFieldName:  "image_height",
			ProtobufType:       "int32",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "prop_inputtype",
			Comment:            `输入类型。TEXT：文本，CHECKBOX：多选，NUMBER：数字，EMAIL：email，DATETIME：时间，URL：链接，DATETIMERANGE：时间段，RADIO：单选，IMAGE：图片，INVALID_PROP_INPUT_TYPE：无效值`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "PropInputtype",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "prop_inputtype",
			ProtobufFieldName:  "prop_inputtype",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "add_time",
			Comment:            `创建时间`,
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
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},
	},
}

// TableName sets the insert table name for this struct type
func (p *PropKs) TableName() string {
	return "prop_ks"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PropKs) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PropKs) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PropKs) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (p *PropKs) TableInfo() *TableInfo {
	return prop_ksTableInfo
}
