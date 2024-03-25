package entity

type Company struct {
	ID           int64
	CompanyName  string `xorm:"'CompanyName'"`
	CreationTime int64  `xorm:"'CreationTime'"`
	Remark       string `xorm:"'Remark'"`
}
