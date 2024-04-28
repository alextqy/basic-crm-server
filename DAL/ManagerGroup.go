package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type ManagerGroupDal struct{}

func (o *ManagerGroupDal) Count(db *gorm.DB, Stext, Outfit string) int64 {
	var Count int64
	TableName := managerGroupTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("GroupName LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *ManagerGroupDal) Add(db *gorm.DB, Data mod.ManagerGroupMod, Outfit string) (int64, error) {
	TableName := managerGroupTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *ManagerGroupDal) Update(db *gorm.DB, Data mod.ManagerGroupMod, Outfit string) error {
	TableName := managerGroupTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *ManagerGroupDal) Data(db *gorm.DB, ID int64, Outfit string) mod.ManagerGroupMod {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroupMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *ManagerGroupDal) List(db *gorm.DB, Page, PageSize, Order int, Stext, Outfit string) (int, int, int, []mod.ManagerGroupMod) {
	TableName := managerGroupTable + Outfit
	Data := []mod.ManagerGroupMod{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("GroupName LIKE ?", "%"+Stext+"%")
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

func (o *ManagerGroupDal) All(db *gorm.DB, Order int, Stext, Outfit string) []mod.ManagerGroupMod {
	TableName := managerGroupTable + Outfit
	Data := []mod.ManagerGroupMod{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("GroupName LIKE ?", "%"+Stext+"%")
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

func (o *ManagerGroupDal) Del(db *gorm.DB, ID, Outfit string) error {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroupMod{}
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
