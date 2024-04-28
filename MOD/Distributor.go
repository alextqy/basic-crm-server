package mod

type DistributorMod struct {
	ID              int64  `gorm:"column:ID;primarykey"`
	Name            string `gorm:"column:Name"`            // 名称
	Email           string `gorm:"column:Email"`           // 电子邮件
	Tel             string `gorm:"column:Tel"`             // 电话号码
	DistributorInfo string `gorm:"column:DistributorInfo"` // 客户信息
	CreationTime    int64  `gorm:"column:CreationTime"`    // 创建时间
	CompanyID       int64  `gorm:"column:CompanyID"`       // 归属公司
	ManagerID       int64  `gorm:"column:ManagerID"`       // 归属销售人员
	AfterServiceID  int64  `gorm:"column:AfterServiceID"`  // 归属售后人员
	Level           int64  `gorm:"column:Level"`           // 0无 1~5
}
