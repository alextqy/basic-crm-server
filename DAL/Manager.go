package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type ManagerDal struct{}

func (o *ManagerDal) Count(db *gorm.DB, Stext string, Level int64, Status int64, GroupID int64, Outfit string) int64 {
	var Count int64
	TableName := managerTable + Outfit
	engine := db.Table(TableName)
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if GroupID > 0 {
		engine = engine.Where("GroupID = ?", GroupID)
	}
	if Stext != "" {
		engine = engine.Where("Account LIKE ?", "%"+Stext+"%").Or("Name LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *ManagerDal) Add(db *gorm.DB, Data mod.Manager, Outfit string) (int64, error) {
	TableName := managerTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *ManagerDal) Update(db *gorm.DB, Data mod.Manager, Outfit string) error {
	TableName := managerTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *ManagerDal) Data(db *gorm.DB, ID int64, Outfit string) mod.Manager {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *ManagerDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, Level int64, Status int64, GroupID int64, Outfit string) (int, int, int, []mod.Manager) {
	TableName := managerTable + Outfit
	Data := []mod.Manager{}
	engine := db.Table(TableName)
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if GroupID > 0 {
		engine = engine.Where("GroupID = ?", GroupID)
	}
	if Stext != "" {
		engine = engine.Where("Account LIKE ?", "%"+Stext+"%").Or("Name LIKE ?", "%"+Stext+"%")
	}
	if Page <= 1 {
		Page = 1
	}
	if PageSize <= 0 {
		PageSize = 10
	}
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.Order("ID " + OrderBy).Limit(int(PageSize)).Offset(int((Page - 1) * PageSize)).Find(&Data)

	Count := o.Count(db, Stext, Level, Status, GroupID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *ManagerDal) All(db *gorm.DB, Order int, Stext string, Level int64, Status int64, GroupID int64, Outfit string) []mod.Manager {
	TableName := managerTable + Outfit
	Data := []mod.Manager{}
	engine := db.Table(TableName)
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if GroupID > 0 {
		engine = engine.Where("GroupID = ?", GroupID)
	}
	if Stext != "" {
		engine = engine.Where("Account LIKE ?", "%"+Stext+"%").Or("Name LIKE ?", "%"+Stext+"%")
	}
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.Order("ID " + OrderBy).Find(&Data)
	return Data
}

func (o *ManagerDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	var e error
	if sysHelper.StringContains(ID, ",") {
		ids := strings.Split(ID, ",")
		intArr := []int{}
		for i := 0; i < len(ids); i++ {
			_, _, n := sysHelper.StringToInt(ids[i])
			intArr = append(intArr, n)
		}
		e = db.Table(TableName).Delete(Data, intArr).Error
	} else {
		e = db.Table(TableName).Delete(Data, ID).Error
	}
	return e
}

func (o *ManagerDal) Check(db *gorm.DB, Account, Outfit string) mod.Manager {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	db.Table(TableName).Where("Account = ?", Account).First(&Data)
	return Data
}

func (o *ManagerDal) Token(db *gorm.DB, Token, Outfit string) mod.Manager {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	db.Table(TableName).Where("Token = ?", Token).First(&Data)
	return Data
}
