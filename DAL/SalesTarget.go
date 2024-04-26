package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type SalesTargetDal struct{}

func (o *SalesTargetDal) Count(db *gorm.DB, Stext string, CustomerID int64, Outfit string) int64 {
	var Count int64
	TableName := salesTargetTable + Outfit
	engine := db.Table(TableName)
	if CustomerID > 0 {
		engine = engine.Where("CustomerID = ?", CustomerID)
	}
	if Stext != "" {
		engine = engine.Where("TargetName LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *SalesTargetDal) Add(db *gorm.DB, Data mod.SalesTargetMod, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *SalesTargetDal) Update(db *gorm.DB, Data mod.SalesTargetMod, Outfit string) error {
	TableName := salesTargetTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *SalesTargetDal) Data(db *gorm.DB, ID int64, Outfit string) mod.SalesTargetMod {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTargetMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *SalesTargetDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, CustomerID int64, ManagerID int64, Outfit string) (int, int, int, []mod.SalesTargetMod) {
	TableName := salesTargetTable + Outfit
	Data := []mod.SalesTargetMod{}
	engine := db.Table(TableName)
	if CustomerID > 0 {
		engine = engine.Where("CustomerID = ?", CustomerID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if Stext != "" {
		engine = engine.Where("TargetName LIKE ?", "%"+Stext+"%")
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

	Count := o.Count(db, Stext, CustomerID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *SalesTargetDal) All(db *gorm.DB, Order int, Stext string, CustomerID int64, ManagerID int64, Outfit string) []mod.SalesTargetMod {
	TableName := salesTargetTable + Outfit
	Data := []mod.SalesTargetMod{}
	engine := db.Table(TableName)
	if CustomerID > 0 {
		engine = engine.Where("CustomerID = ?", CustomerID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if Stext != "" {
		engine = engine.Where("TargetName LIKE ?", "%"+Stext+"%")
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

func (o *SalesTargetDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTargetMod{}
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
