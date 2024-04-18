package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type AfterServiceDal struct{}

func (a *AfterServiceDal) Count(db *gorm.DB, Stext string, Level int64, Status int64, Outfit string) int64 {
	var Count int64
	TableName := AfterServiceTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("Account LIKE ?", "%"+Stext+"%").Or("Name LIKE ?", "%"+Stext+"%")
	}
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	engine.Count(&Count)
	return Count
}

func (a *AfterServiceDal) Add(db *gorm.DB, Data mod.AfterService, Outfit string) (int64, error) {
	TableName := AfterServiceTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (a *AfterServiceDal) Update(db *gorm.DB, Data mod.AfterService, Outfit string) error {
	TableName := AfterServiceTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (a *AfterServiceDal) Data(db *gorm.DB, ID int64, Outfit string) mod.AfterService {
	TableName := AfterServiceTable + Outfit
	Data := mod.AfterService{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (a *AfterServiceDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, Level int64, Status int64, Outfit string) (int, int, int, []mod.AfterService) {
	TableName := AfterServiceTable + Outfit
	Data := []mod.AfterService{}
	engine := db.Table(TableName)
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
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

	Count := a.Count(db, Stext, Level, Status, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (a *AfterServiceDal) All(db *gorm.DB, Order int, Stext string, Level int64, Status int64, Outfit string) []mod.AfterService {
	TableName := AfterServiceTable + Outfit
	Data := []mod.AfterService{}
	engine := db.Table(TableName)
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
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

func (a *AfterServiceDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := AfterServiceTable + Outfit
	Data := mod.AfterService{}
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

func (a *AfterServiceDal) Check(db *gorm.DB, Account, Outfit string) mod.AfterService {
	TableName := AfterServiceTable + Outfit
	Data := mod.AfterService{}
	db.Table(TableName).Where("Account = ?", Account).First(&Data)
	return Data
}

func (a *AfterServiceDal) Token(db *gorm.DB, Token, Outfit string) mod.AfterService {
	TableName := AfterServiceTable + Outfit
	Data := mod.AfterService{}
	db.Table(TableName).Where("Token = ?", Token).First(&Data)
	return Data
}
