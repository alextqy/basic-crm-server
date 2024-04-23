package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type CustomerDal struct{}

func (o *CustomerDal) Count(db *gorm.DB, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64, AfterServiceID int64, Outfit string) int64 {
	var Count int64
	TableName := customerTable + Outfit
	engine := db.Table(TableName)
	if Gender > 0 {
		engine = engine.Where("Gender = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.Where("Priority = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.Where("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if AfterServiceID > 0 {
		engine = engine.Where("AfterServiceID = ?", AfterServiceID)
	}
	if Stext != "" {
		engine = engine.Where("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *CustomerDal) Add(db *gorm.DB, Data mod.Customer, Outfit string) (int64, error) {
	TableName := customerTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *CustomerDal) Update(db *gorm.DB, Data mod.Customer, Outfit string) error {
	TableName := customerTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *CustomerDal) Data(db *gorm.DB, ID int64, Outfit string) mod.Customer {
	TableName := customerTable + Outfit
	Data := mod.Customer{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *CustomerDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64, AfterServiceID int64, Outfit string) (int, int, int, []mod.Customer) {
	TableName := customerTable + Outfit
	Data := []mod.Customer{}
	engine := db.Table(TableName)
	if Gender > 0 {
		engine = engine.Where("Gender = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.Where("Priority = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.Where("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if AfterServiceID > 0 {
		engine = engine.Where("AfterServiceID = ?", AfterServiceID)
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

	Count := o.Count(db, Stext, Gender, Priority, CompanyID, ManagerID, AfterServiceID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *CustomerDal) All(db *gorm.DB, Order int, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64, AfterServiceID int64, Outfit string) []mod.Customer {
	TableName := customerTable + Outfit
	Data := []mod.Customer{}
	engine := db.Table(TableName)
	if Gender > 0 {
		engine = engine.Where("Gender = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.Where("Priority = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.Where("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if AfterServiceID > 0 {
		engine = engine.Where("AfterServiceID = ?", AfterServiceID)
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

func (o *CustomerDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := customerTable + Outfit
	Data := mod.Customer{}
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
