package mod

type Company struct {
	ID           int64  `xorm:"pk ID"`
	CompanyName  string `xorm:"CompanyName"`
	CreationTime int64  `xorm:"CreationTime"`
	Remark       string `xorm:"Remark"`
}
