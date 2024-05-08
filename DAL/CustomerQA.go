package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type CustomerQADal struct{}

func (o *CustomerQADal) Count(db *gorm.DB, Stext string, Display int64, Outfit string) int64 {
	var Count int64
	TableName := customerQATable + Outfit
	engine := db.Table(TableName)
	if Display > 0 {
		engine = engine.Where("Display = ?", Display)
	}
	if Stext != "" {
		engine = engine.Where("Title LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *CustomerQADal) Add(db *gorm.DB, Data mod.CustomerQAMod, Outfit string) (int64, error) {
	TableName := customerQATable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *CustomerQADal) Update(db *gorm.DB, Data mod.CustomerQAMod, Outfit string) error {
	TableName := customerQATable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *CustomerQADal) Data(db *gorm.DB, ID int64, Outfit string) mod.CustomerQAMod {
	TableName := customerQATable + Outfit
	Data := mod.CustomerQAMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *CustomerQADal) List(db *gorm.DB, Page, PageSize, Order int, Stext string, Display int64, Outfit string) (int, int, int, []mod.CustomerQAMod) {
	TableName := customerQATable + Outfit
	Data := []mod.CustomerQAMod{}
	engine := db.Table(TableName)
	if Display > 0 {
		engine = engine.Where("Display = ?", Display)
	}
	if Stext != "" {
		engine = engine.Where("Title LIKE ?", "%"+Stext+"%")
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

	Count := o.Count(db, Stext, Display, "")
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *CustomerQADal) All(db *gorm.DB, Order int, Stext string, Display int64, Outfit string) []mod.CustomerQAMod {
	TableName := customerQATable + Outfit
	Data := []mod.CustomerQAMod{}
	engine := db.Table(TableName)
	if Display > 0 {
		engine = engine.Where("Display = ?", Display)
	}
	if Stext != "" {
		engine = engine.Where("Title LIKE ?", "%"+Stext+"%")
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

func (o *CustomerQADal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := customerQATable + Outfit
	Data := mod.CustomerQAMod{}
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
