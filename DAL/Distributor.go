package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type DistributorDal struct{}

func (o *DistributorDal) Count(db *gorm.DB, Stext string, CompanyID, ManagerID, AfterServiceID, Level int64, Outfit string) int64 {
	var Count int64
	TableName := distributorTable + Outfit
	engine := db.Table(TableName)
	if CompanyID > 0 {
		engine = engine.Where("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if AfterServiceID > 0 {
		engine = engine.Where("AfterServiceID = ?", AfterServiceID)
	}
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Stext != "" {
		engine = engine.Where("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *DistributorDal) Add(db *gorm.DB, Data mod.DistributorMod, Outfit string) (int64, error) {
	TableName := distributorTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *DistributorDal) Update(db *gorm.DB, Data mod.DistributorMod, Outfit string) error {
	TableName := distributorTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *DistributorDal) Data(db *gorm.DB, ID int64, Outfit string) mod.DistributorMod {
	TableName := distributorTable + Outfit
	Data := mod.DistributorMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *DistributorDal) List(db *gorm.DB, Page, PageSize, Order int, Stext string, CompanyID, ManagerID, AfterServiceID, Level int64, Outfit string) (int, int, int, []mod.DistributorMod) {
	TableName := distributorTable + Outfit
	Data := []mod.DistributorMod{}
	engine := db.Table(TableName)
	if CompanyID > 0 {
		engine = engine.Where("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if AfterServiceID > 0 {
		engine = engine.Where("AfterServiceID = ?", AfterServiceID)
	}
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Stext != "" {
		engine = engine.Where("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
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

	Count := o.Count(db, Stext, CompanyID, ManagerID, AfterServiceID, Level, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *DistributorDal) All(db *gorm.DB, Order int, Stext string, CompanyID, ManagerID, AfterServiceID, Level int64, Outfit string) []mod.DistributorMod {
	TableName := distributorTable + Outfit
	Data := []mod.DistributorMod{}
	engine := db.Table(TableName)
	if CompanyID > 0 {
		engine = engine.Where("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if AfterServiceID > 0 {
		engine = engine.Where("AfterServiceID = ?", AfterServiceID)
	}
	if Level > 0 {
		engine = engine.Where("Level = ?", Level)
	}
	if Stext != "" {
		engine = engine.Where("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
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

func (o *DistributorDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := distributorTable + Outfit
	Data := mod.DistributorMod{}
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
