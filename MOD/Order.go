package mod

// 订单
type OrderMod struct {
	ID            int64   `gorm:"column:ID;primarykey"`
	OrderNo       string  `gorm:"column:OrderNo"`       // 订单号
	ProductID     int64   `gorm:"column:ProductID"`     // 产品ID
	ManagerID     int64   `gorm:"column:ManagerID"`     // 销售ID
	CustomerID    int64   `gorm:"column:CustomerID"`    // 客户ID
	DistributorID int64   `gorm:"column:DistributorID"` // 渠道商ID
	OrderPrice    float32 `gorm:"column:OrderPrice"`    // 订单价格
	ProductPrice  float32 `gorm:"column:ProductPrice"`  // 产品价格
	ProductCost   float32 `gorm:"column:ProductCost"`   // 产品成本
	Status        int64   `gorm:"column:Status"`        // 状态 1正常 2作废
	Remark        string  `gorm:"column:Remark"`        // 备注
	CreationTime  int64   `gorm:"column:CreationTime"`  // 创建时间
	OrderType     int64   `gorm:"column:OrderType"`     // 订单类型 1客户 2渠道商
	Payment       int64   `gorm:"column:Payment"`       // 付款状态 1未支付 2已支付
	Review        int64   `gorm:"column:Review"`        // 合同审核状态 1未审核 2已审核
}
