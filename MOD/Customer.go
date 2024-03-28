package mod

type Customer struct {
	ID           int64  `xorm:"ID"`
	Name         string `xorm:"Name"`
	Birthday     int64  `xorm:"Birthday"`
	Gender       int64  `xorm:"Gender"`
	Email        string `xorm:"Email"`
	Tel          string `xorm:"Tel"`
	CustomerInfo string `xorm:"CustomerInfo"`
	Priority     int64  `xorm:"Priority"`
	CreationTime int64  `xorm:"CreationTime"`
	CompanyID    int64  `xorm:"CompanyID"`
	ManagerID    int64  `xorm:"ManagerID"`
}
