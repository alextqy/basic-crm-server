package entity

type Admin struct {
	ID           int64
	Account      string `xorm:"'Account'"`
	Password     string `xorm:"'Password'"`
	Name         string `xorm:"'Name'"`
	Level        int64  `xorm:"'Level'"`
	Status       int64  `xorm:"'Status'"`
	Remark       string `xorm:"'Remark'"`
	Token        string `xorm:"'Token'"`
	CreationTime int64  `xorm:"'CreationTime'"`
}
