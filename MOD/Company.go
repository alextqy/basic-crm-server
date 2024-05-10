package mod

// 公司
type CompanyMod struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	CompanyName  string `gorm:"column:CompanyName"`  // 名称
	CreationTime int64  `gorm:"column:CreationTime"` // 创建时间
	Remark       string `gorm:"column:Remark"`       // 备注
}
