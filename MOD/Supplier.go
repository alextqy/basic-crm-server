package mod

// 供应商
type SupplierMod struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	Name         string `gorm:"column:Name"`         // 名称
	Email        string `gorm:"column:Email"`        // 电子邮件
	Tel          string `gorm:"column:Tel"`          // 电话号码
	Address      string `gorm:"column:Address"`      // 地址
	SupplierInfo string `gorm:"column:SupplierInfo"` // 客户信息
	CreationTime int64  `gorm:"column:CreationTime"` // 创建时间
}
