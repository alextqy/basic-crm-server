package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type CompanyDal struct{}

func (c *CompanyDal) Count(db *gorm.DB, Stext string, Outfit string) int64 {
	var Count int64
	TableName := companyTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("CompanyName LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (c *CompanyDal) Add(db *gorm.DB, Data mod.Company, Outfit string) (int64, error) {
	TableName := companyTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (c *CompanyDal) Update(db *gorm.DB, Data mod.Company, Outfit string) error {
	TableName := companyTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (c *CompanyDal) Data(db *gorm.DB, ID int64, Outfit string) mod.Company {
	TableName := companyTable + Outfit
	Data := mod.Company{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (c *CompanyDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, Outfit string) (int, int, int, []mod.Company) {
	TableName := companyTable + Outfit
	Data := []mod.Company{}
	engine := db.Table(TableName)
	if Page <= 1 {
		Page = 1
	}
	if PageSize <= 0 {
		PageSize = 10
	}
	if Stext != "" {
		engine = engine.Where("CompanyName LIKE ?", "%"+Stext+"%")
	}
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.Order("ID " + OrderBy).Limit(int(PageSize)).Offset(int((Page - 1) * PageSize)).Find(&Data)

	Count := c.Count(db, Stext, TableName)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (c *CompanyDal) All(db *gorm.DB, Order int, Stext string, Outfit string) []mod.Company {
	TableName := companyTable + Outfit
	Data := []mod.Company{}
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

func (c *CompanyDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := companyTable + Outfit
	Data := mod.Company{}
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

func (c *CompanyDal) Check(db *gorm.DB, CompanyName, Outfit string) mod.Company {
	TableName := companyTable + Outfit
	Data := mod.Company{}
	db.Table(TableName).Where("CompanyName = ?", CompanyName).First(&Data)
	return Data
}
