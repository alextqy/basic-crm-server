package mod

type ManagerGroup struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	GroupName    string `gorm:"column:GroupName"`
	CreationTime int64  `gorm:"column:CreationTime"`
	Remark       string `gorm:"column:Remark"`
}
