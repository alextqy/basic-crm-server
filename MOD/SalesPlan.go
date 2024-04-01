package mod

type SalesPlan struct {
	ID           int64   `xorm:"pk ID"`
	PlanName     string  `xorm:"PlanName"`
	TargetID     int64   `xorm:"TargetID"`
	PlanContent  string  `xorm:"PlanContent"`
	CreationTime int64   `xorm:"CreationTime"`
	Status       int64   `xorm:"Status"`
	Budget       float32 `xorm:"Budget"`
}
