package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type ManagerGroupDal struct{}

func (m *ManagerGroupDal) Count(db *gorm.DB, Stext string, Outfit string) int64 {
	var Count int64
	TableName := managerGroupTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("GroupName LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (m *ManagerGroupDal) Add(db *gorm.DB, Data mod.ManagerGroup, Outfit string) (int64, error) {
	TableName := managerGroupTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (m *ManagerGroupDal) Update(db *gorm.DB, Data mod.ManagerGroup, Outfit string) error {
	TableName := managerGroupTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (m *ManagerGroupDal) Data(db *gorm.DB, ID int64, Outfit string) mod.ManagerGroup {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroup{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (m *ManagerGroupDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, Outfit string) (int, int, int, []mod.ManagerGroup) {
	TableName := managerGroupTable + Outfit
	Data := []mod.ManagerGroup{}
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

	Count := m.Count(db, Stext, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (m *ManagerGroupDal) All(db *gorm.DB, Order int, Stext string, Outfit string) []mod.ManagerGroup {
	TableName := managerGroupTable + Outfit
	Data := []mod.ManagerGroup{}
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

func (m *ManagerGroupDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroup{}
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
