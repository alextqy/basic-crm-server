package mod

// 产品
type ProductMod struct {
	ID           int64   `gorm:"column:ID;primarykey"`
	ProductName  string  `gorm:"column:ProductName"`  // 名称
	Price        float32 `gorm:"column:Price"`        // 价格
	Cost         float32 `gorm:"column:Cost"`         // 产品成本
	Status       int64   `gorm:"column:Status"`       // 状态 1正常 2作废
	Remark       string  `gorm:"column:Remark"`       // 备注
	CreationTime int64   `gorm:"column:CreationTime"` // 创建时间
}
