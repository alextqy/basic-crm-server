package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type SupplierDal struct{}

func (o *SupplierDal) Count(db *gorm.DB, Stext string, Outfit string) int64 {
	var Count int64
	TableName := supplierTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *SupplierDal) Add(db *gorm.DB, Data mod.SupplierMod, Outfit string) (int64, error) {
	TableName := supplierTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *SupplierDal) Update(db *gorm.DB, Data mod.SupplierMod, Outfit string) error {
	TableName := supplierTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *SupplierDal) Data(db *gorm.DB, ID int64, Outfit string) mod.SupplierMod {
	TableName := supplierTable + Outfit
	Data := mod.SupplierMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *SupplierDal) List(db *gorm.DB, Page, PageSize, Order int, Stext, Outfit string) (int, int, int, []mod.SupplierMod) {
	TableName := supplierTable + Outfit
	Data := []mod.SupplierMod{}
	engine := db.Table(TableName)
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

	Count := o.Count(db, Stext, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *SupplierDal) All(db *gorm.DB, Order int, Stext, Outfit string) []mod.SupplierMod {
	TableName := supplierTable + Outfit
	Data := []mod.SupplierMod{}
	engine := db.Table(TableName)
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

func (o *SupplierDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := supplierTable + Outfit
	Data := mod.SupplierMod{}
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
