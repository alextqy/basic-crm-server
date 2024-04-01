package mod

type ManagerGroup struct {
	ID           int64  `xorm:"pk ID"`
	GroupName    string `xorm:"GroupName"`
	CreationTime int64  `xorm:"CreationTime"`
	Remark       string `xorm:"Remark"`
}
