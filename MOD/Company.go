package mod

type Company struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	CompanyName  string `gorm:"column:CompanyName"`
	CreationTime int64  `gorm:"column:CreationTime"`
	Remark       string `gorm:"column:Remark"`
}
