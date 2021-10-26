package model

// ActivityInfo struct is a row record of the activity_info table in the yx database
type ActivityInfo struct {
	ActivityID          uint32 `gorm:"primary_key;AUTO_INCREMENT;column:activity_id;type:uint;" json:"activity_id"`
	ActivityName        string `gorm:"column:activity_name;type:varchar;size:255;" json:"activity_name"`                 // 活动名称
	ExposureTalentNum   int32  `gorm:"column:exposure_talent_num;type:int;default:0;" json:"exposure_talent_num"`        // 曝光达人数
	CommerceTalentNum   int32  `gorm:"column:commerce_talent_num;type:int;default:0;" json:"commerce_talent_num"`        // 带货达人数
	PlatformID          int32  `gorm:"column:platform_id;type:int;default:0;" json:"platform_id"`                        // 平台id 1 抖音 2 快手
	ActivityDesc        string `gorm:"column:activity_desc;type:varchar;size:255;" json:"activity_desc"`                 // 活动描述
	ThirdActivityID     int32  `gorm:"column:third_activity_id;type:int;default:0;" json:"third_activity_id"`            // 第三方活动id
	ServiceRate         int32  `gorm:"column:service_rate;type:int;default:0;" json:"service_rate"`                      // 最低团长服务费率
	CommissionRate      int32  `gorm:"column:commission_rate;type:int;default:0;" json:"commission_rate"`                // 最低商品佣金率
	EstimatedSingleSale int64  `gorm:"column:estimated_single_sale;type:bigint;default:0;" json:"estimated_single_sale"` // 预估销售额
	StartTime           int32  `gorm:"column:start_time;type:int;default:0;" json:"start_time"`                          // 开始时间
	EndTime             int32  `gorm:"column:end_time;type:int;default:0;" json:"end_time"`                              // 活动结束时间
	PromoteStartTime    int32  `gorm:"column:promote_start_time;type:int;default:0;" json:"promote_start_time"`          // 推广开始时间
	PromoteEndTime      int32  `gorm:"column:promote_end_time;type:int;default:0;" json:"promote_end_time"`              // 推广结束时间
	Status              int32  `gorm:"column:status;type:int;default:0;" json:"status"`                                  // 状态 目前只有团长有用
	StatusDesc          string `gorm:"column:status_desc;type:varchar;size:30;" json:"status_desc"`                      // 状态描述
	ActivityType        int32  `gorm:"column:activity_type;type:int;default:1;" json:"activity_type"`
	ModelTime
}

// TableName sets the insert table name for this struct type
func (a *ActivityInfo) TableName() string {
	return "activity_info"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ActivityInfo) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ActivityInfo) Prepare() {
}
