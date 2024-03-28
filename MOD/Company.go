package mod

type Company struct {
	ID           int64  `xorm:"ID"`
	CompanyName  string `xorm:"CompanyName"`
	CreationTime int64  `xorm:"CreationTime"`
	Remark       string `xorm:"Remark"`
}
