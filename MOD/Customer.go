package mod

type Customer struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	Name         string `gorm:"column:Name"`
	Birthday     int64  `gorm:"column:Birthday"`
	Gender       int64  `gorm:"column:Gender"`
	Email        string `gorm:"column:Email"`
	Tel          string `gorm:"column:Tel"`
	CustomerInfo string `gorm:"column:CustomerInfo"`
	Priority     int64  `gorm:"column:Priority"`
	CreationTime int64  `gorm:"column:CreationTime"`
	CompanyID    int64  `gorm:"column:CompanyID"`
	ManagerID    int64  `gorm:"column:ManagerID"`
}
