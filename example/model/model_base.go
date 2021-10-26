package example

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["activity_info"] = activity_infoTableInfo
	tables["activity_product"] = activity_productTableInfo
	tables["activity_sku"] = activity_skuTableInfo
	tables["ad_info"] = ad_infoTableInfo
	tables["admin_node"] = admin_nodeTableInfo
	tables["admin_role"] = admin_roleTableInfo
	tables["admin_role_node"] = admin_role_nodeTableInfo
	tables["admin_user"] = admin_userTableInfo
	tables["admin_user_role"] = admin_user_roleTableInfo
	tables["agreement_info"] = agreement_infoTableInfo
	tables["anchor_day"] = anchor_dayTableInfo
	tables["anchor_info"] = anchor_infoTableInfo
	tables["attr_info"] = attr_infoTableInfo
	tables["attr_value"] = attr_valueTableInfo
	tables["batch_fail"] = batch_failTableInfo
	tables["black_list"] = black_listTableInfo
	tables["brand_dy"] = brand_dyTableInfo
	tables["brand_info"] = brand_infoTableInfo
	tables["cate_dy"] = cate_dyTableInfo
	tables["cate_industry"] = cate_industryTableInfo
	tables["cate_industry_relate"] = cate_industry_relateTableInfo
	tables["cate_info"] = cate_infoTableInfo
	tables["cate_ks"] = cate_ksTableInfo
	tables["cate_tb"] = cate_tbTableInfo
	tables["catelm_info"] = catelm_infoTableInfo
	tables["catelm_relation"] = catelm_relationTableInfo
	tables["city_info"] = city_infoTableInfo
	tables["colonel_order_dy"] = colonel_order_dyTableInfo
	tables["commission_info"] = commission_infoTableInfo
	tables["commission_plan"] = commission_planTableInfo
	tables["coupon_dy"] = coupon_dyTableInfo
	tables["crontab_break"] = crontab_breakTableInfo
	tables["delivery_info"] = delivery_infoTableInfo
	tables["dict_info"] = dict_infoTableInfo
	tables["dy_order_history"] = dy_order_historyTableInfo
	tables["express_company"] = express_companyTableInfo
	tables["feedback_info"] = feedback_infoTableInfo
	tables["gmv_dy"] = gmv_dyTableInfo
	tables["gmv_ks"] = gmv_ksTableInfo
	tables["live_info"] = live_infoTableInfo
	tables["live_product"] = live_productTableInfo
	tables["logistics_company"] = logistics_companyTableInfo
	tables["logistics_dy"] = logistics_dyTableInfo
	tables["mcn_info"] = mcn_infoTableInfo
	tables["merchant_cate"] = merchant_cateTableInfo
	tables["merchant_enroll_log"] = merchant_enroll_logTableInfo
	tables["merchant_info"] = merchant_infoTableInfo
	tables["merchant_right"] = merchant_rightTableInfo
	tables["notify_order_dy"] = notify_order_dyTableInfo
	tables["notify_order_ks"] = notify_order_ksTableInfo
	tables["notify_product_ks"] = notify_product_ksTableInfo
	tables["oauth_task"] = oauth_taskTableInfo
	tables["open_info"] = open_infoTableInfo
	tables["order_dy"] = order_dyTableInfo
	tables["order_dy_temp"] = order_dy_tempTableInfo
	tables["order_info"] = order_infoTableInfo
	tables["order_ks"] = order_ksTableInfo
	tables["order_product"] = order_productTableInfo
	tables["order_product_dy"] = order_product_dyTableInfo
	tables["order_product_dy_temp"] = order_product_dy_tempTableInfo
	tables["order_product_ks"] = order_product_ksTableInfo
	tables["orientation_plan"] = orientation_planTableInfo
	tables["product_attr"] = product_attrTableInfo
	tables["product_dy"] = product_dyTableInfo
	tables["product_info"] = product_infoTableInfo
	tables["product_ks"] = product_ksTableInfo
	tables["product_pool"] = product_poolTableInfo
	tables["product_rule"] = product_ruleTableInfo
	tables["product_window"] = product_windowTableInfo
	tables["prop_dy"] = prop_dyTableInfo
	tables["prop_info"] = prop_infoTableInfo
	tables["prop_ks"] = prop_ksTableInfo
	tables["prop_value"] = prop_valueTableInfo
	tables["prop_value_ks"] = prop_value_ksTableInfo
	tables["rule_info"] = rule_infoTableInfo
	tables["sample_order"] = sample_orderTableInfo
	tables["sample_order_product"] = sample_order_productTableInfo
	tables["share_scene"] = share_sceneTableInfo
	tables["shop_info"] = shop_infoTableInfo
	tables["shop_rule"] = shop_ruleTableInfo
	tables["sku_dy"] = sku_dyTableInfo
	tables["sku_info"] = sku_infoTableInfo
	tables["sku_ks"] = sku_ksTableInfo
	tables["spec_detail_dy"] = spec_detail_dyTableInfo
	tables["spec_dy"] = spec_dyTableInfo
	tables["statistics_info"] = statistics_infoTableInfo
	tables["statistics_temp"] = statistics_tempTableInfo
	tables["tag_info"] = tag_infoTableInfo
	tables["talent_cate"] = talent_cateTableInfo
	tables["talent_cate_industry"] = talent_cate_industryTableInfo
	tables["talent_info"] = talent_infoTableInfo
	tables["talent_right"] = talent_rightTableInfo
	tables["talent_tag"] = talent_tagTableInfo
	tables["tender_cate"] = tender_cateTableInfo
	tables["tender_info"] = tender_infoTableInfo
	tables["tender_survey"] = tender_surveyTableInfo
	tables["trade_info"] = trade_infoTableInfo
	tables["white_list"] = white_listTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
