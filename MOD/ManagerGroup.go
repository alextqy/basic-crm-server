package mod

type ManagerGroup struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	GroupName    string `gorm:"column:GroupName"`    // 名称
	CreationTime int64  `gorm:"column:CreationTime"` // 创建时间
	Remark       string `gorm:"column:Remark"`       // 备注
}
