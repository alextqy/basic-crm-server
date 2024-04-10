package mod

type SalesTarget struct {
	ID              int64   `gorm:"column:ID;primarykey"`
	TargetName      string  `gorm:"column:TargetName"`
	ExpirationDate  int64   `gorm:"column:ExpirationDate"`
	CreationTime    int64   `gorm:"column:CreationTime"`
	AchievementRate float32 `gorm:"column:AchievementRate"`
	CustomerID      int64   `gorm:"column:CustomerID"`
	Remark          string  `gorm:"column:Remark"`
}
