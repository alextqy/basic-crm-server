package mod

type AnnouncementMod struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	Title        string `gorm:"column:Title"`        // 标题
	Content      string `gorm:"column:Content"`      // 内容
	AuthorID     int64  `gorm:"column:AuthorID"`     // 作者
	Display      int64  `gorm:"column:Display"`      // 展示 1是 2否
	CreationTime int64  `gorm:"column:CreationTime"` // 创建时间
}
