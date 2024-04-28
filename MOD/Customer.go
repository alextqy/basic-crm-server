package mod

type CustomerMod struct {
	ID             int64  `gorm:"column:ID;primarykey"`
	Name           string `gorm:"column:Name"`           // 名称
	Birthday       int64  `gorm:"column:Birthday"`       // 出生日期
	Gender         int64  `gorm:"column:Gender"`         // 性别 1女 2男 3其他
	Email          string `gorm:"column:Email"`          // 电子邮件
	Tel            string `gorm:"column:Tel"`            // 电话号码
	CustomerInfo   string `gorm:"column:CustomerInfo"`   // 客户信息
	Priority       int64  `gorm:"column:Priority"`       // 优先级
	CreationTime   int64  `gorm:"column:CreationTime"`   // 创建时间
	CompanyID      int64  `gorm:"column:CompanyID"`      // 归属公司
	ManagerID      int64  `gorm:"column:ManagerID"`      // 归属销售人员
	AfterServiceID int64  `gorm:"column:AfterServiceID"` // 归属售后人员
	Level          int64  `gorm:"column:Level"`          // 0无 1~5
}
