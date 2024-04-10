package mod

type SalesPlan struct {
	ID           int64   `gorm:"column:ID;primarykey"`
	PlanName     string  `gorm:"column:PlanName"`
	TargetID     int64   `gorm:"column:TargetID"`
	PlanContent  string  `gorm:"column:PlanContent"`
	CreationTime int64   `gorm:"column:CreationTime"`
	Status       int64   `gorm:"column:Status"`
	Budget       float32 `gorm:"column:Budget"`
}
