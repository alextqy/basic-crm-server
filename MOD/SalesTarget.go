package mod

type SalesTargetMod struct {
	ID              int64   `gorm:"column:ID;primarykey"`
	TargetName      string  `gorm:"column:TargetName"`      // 名称
	ExpirationDate  int64   `gorm:"column:ExpirationDate"`  // 截止日期
	CreationTime    int64   `gorm:"column:CreationTime"`    // 创建时间
	AchievementRate float32 `gorm:"column:AchievementRate"` // 完成率
	CustomerID      int64   `gorm:"column:CustomerID"`      // 归属客户
	ManagerID       int64   `gorm:"column:ManagerID"`       // 归属销售人员
	Remark          string  `gorm:"column:Remark"`          // 备注
}
