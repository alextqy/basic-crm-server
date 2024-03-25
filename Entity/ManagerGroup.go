package entity

type ManagerGroup struct {
	ID           int64
	GroupName    string `xorm:"'GroupName'"`
	CreationTime int64  `xorm:"'CreationTime'"`
	Remark       string `xorm:"'Remark'"`
}
