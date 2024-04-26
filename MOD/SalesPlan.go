package mod

type SalesPlanMod struct {
	ID           int64   `gorm:"column:ID;primarykey"`
	PlanName     string  `gorm:"column:PlanName"`     // 名称
	TargetID     int64   `gorm:"column:TargetID"`     // 归属目标
	PlanContent  string  `gorm:"column:PlanContent"`  // 计划内容
	CreationTime int64   `gorm:"column:CreationTime"` // 创建时间
	Status       int64   `gorm:"column:Status"`       // 状态 1正常 2作废
	Budget       float32 `gorm:"column:Budget"`       // 预算
	ManagerID    int64   `gorm:"column:ManagerID"`    // 归属销售人员
}
