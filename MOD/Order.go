package mod

type OrderMod struct {
	ID           int64   `gorm:"column:ID;primarykey"`
	OrderNo      string  `gorm:"column:OrderNo"`      // 订单号
	ProductID    int64   `gorm:"column:ProductID"`    // 产品ID
	ManagerID    int64   `gorm:"column:ManagerID"`    // 销售ID
	OrderPrice   float32 `gorm:"column:OrderPrice"`   // 订单价格
	ProductPrice float32 `gorm:"column:ProductPrice"` // 产品价格
	ProductCost  float32 `gorm:"column:ProductCost"`  // 产品成本
	Status       int64   `gorm:"column:Status"`       // 状态 1正常 2作废
	Remark       string  `gorm:"column:Remark"`       // 备注
	CreationTime int64   `gorm:"column:CreationTime"` // 创建时间
}
