package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type CompanyDal struct{}

func (o *CompanyDal) Count(db *gorm.DB, Stext string, Outfit string) int64 {
	var Count int64
	TableName := companyTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("CompanyName LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *CompanyDal) Add(db *gorm.DB, Data mod.CompanyMod, Outfit string) (int64, error) {
	TableName := companyTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *CompanyDal) Update(db *gorm.DB, Data mod.CompanyMod, Outfit string) error {
	TableName := companyTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *CompanyDal) Data(db *gorm.DB, ID int64, Outfit string) mod.CompanyMod {
	TableName := companyTable + Outfit
	Data := mod.CompanyMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *CompanyDal) List(db *gorm.DB, Page, PageSize, Order int, Stext, Outfit string) (int, int, int, []mod.CompanyMod) {
	TableName := companyTable + Outfit
	Data := []mod.CompanyMod{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("CompanyName LIKE ?", "%"+Stext+"%")
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

	Count := o.Count(db, Stext, "")
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *CompanyDal) All(db *gorm.DB, Order int, Stext string, Outfit string) []mod.CompanyMod {
	TableName := companyTable + Outfit
	Data := []mod.CompanyMod{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("CompanyName LIKE ?", "%"+Stext+"%")
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

func (o *CompanyDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := companyTable + Outfit
	Data := mod.CompanyMod{}
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

func (o *CompanyDal) Check(db *gorm.DB, CompanyName, Outfit string) mod.CompanyMod {
	TableName := companyTable + Outfit
	Data := mod.CompanyMod{}
	db.Table(TableName).Where("CompanyName = ?", CompanyName).First(&Data)
	return Data
}
