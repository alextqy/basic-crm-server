package mod

type ManagerMod struct {
	ID           int64  `gorm:"column:ID;primarykey"`
	Account      string `gorm:"column:Account"`      // 账号
	Password     string `gorm:"column:Password"`     // 密码
	Name         string `gorm:"column:Name"`         // 名称
	Level        int64  `gorm:"column:Level"`        // 1普通 2高级
	Status       int64  `gorm:"column:Status"`       // 1正常 2禁用
	Remark       string `gorm:"column:Remark"`       // 备注
	Token        string `gorm:"column:Token"`        // 令牌
	CreationTime int64  `gorm:"column:CreationTime"` // 创建时间
	GroupID      int64  `gorm:"column:GroupID"`      // 归属小组
}
