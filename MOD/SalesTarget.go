package mod

type SalesTarget struct {
	ID              int64   `xorm:"ID"`
	TargetName      string  `xorm:"TargetName"`
	ExpirationDate  int64   `xorm:"ExpirationDate"`
	CreationTime    int64   `xorm:"CreationTime"`
	AchievementRate float32 `xorm:"AchievementRate"`
	CustomerID      int64   `xorm:"CustomerID"`
	Remark          string  `xorm:"Remark"`
}
