package mod

type Manager struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	Account      string `gorm:"column:Account"`
	Password     string `gorm:"column:Password"`
	Name         string `gorm:"column:Name"`
	Level        int64  `gorm:"column:Level"`
	Status       int64  `gorm:"column:Status"`
	Remark       string `gorm:"column:Remark"`
	Token        string `gorm:"column:Token"`
	CreationTime int64  `gorm:"column:CreationTime"`
	GroupID      int64  `gorm:"column:GroupID"`
}
