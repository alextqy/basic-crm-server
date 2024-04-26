package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type ProductDal struct{}

func (o *ProductDal) Count(db *gorm.DB, Stext string, Status int64, Outfit string) int64 {
	var Count int64
	TableName := productTable + Outfit
	engine := db.Table(TableName)
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if Stext != "" {
		engine = engine.Where("ProductName LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *ProductDal) Add(db *gorm.DB, Data mod.ProductMod, Outfit string) (int64, error) {
	TableName := productTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *ProductDal) Update(db *gorm.DB, Data mod.ProductMod, Outfit string) error {
	TableName := productTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *ProductDal) Data(db *gorm.DB, ID int64, Outfit string) mod.ProductMod {
	TableName := productTable + Outfit
	Data := mod.ProductMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *ProductDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, Status int64, Outfit string) (int, int, int, []mod.ProductMod) {
	TableName := productTable + Outfit
	Data := []mod.ProductMod{}
	engine := db.Table(TableName)
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if Stext != "" {
		engine = engine.Where("ProductName LIKE ?", "%"+Stext+"%")
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

	Count := o.Count(db, Stext, Status, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *ProductDal) All(db *gorm.DB, Order int, Stext string, Status int64, Outfit string) []mod.ProductMod {
	TableName := productTable + Outfit
	Data := []mod.ProductMod{}
	engine := db.Table(TableName)
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if Stext != "" {
		engine = engine.Where("ProductName LIKE ?", "%"+Stext+"%")
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

func (o *ProductDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := productTable + Outfit
	Data := mod.ProductMod{}
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
