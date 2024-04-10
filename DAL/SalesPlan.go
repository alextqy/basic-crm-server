package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type SalesPlanDal struct{}

func (s *SalesPlanDal) Count(db *gorm.DB, Stext string, TargetID int64, Status int64, Outfit string) int64 {
	var Count int64
	TableName := salesPlanTable + Outfit
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("PlanName LIKE ?", "%"+Stext+"%")
	}
	if TargetID > 0 {
		engine = engine.Where("TargetID = ?", TargetID)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	engine.Count(&Count)
	return Count
}

func (s *SalesPlanDal) Add(db *gorm.DB, Data mod.SalesPlan, Outfit string) (int64, error) {
	TableName := salesPlanTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (s *SalesPlanDal) Update(db *gorm.DB, Data mod.SalesPlan, Outfit string) error {
	TableName := salesPlanTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (s *SalesPlanDal) Data(db *gorm.DB, ID int64, Outfit string) mod.SalesPlan {
	TableName := salesPlanTable + Outfit
	Data := mod.SalesPlan{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (s *SalesPlanDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, TargetID int64, Status int64, Outfit string) (int, int, int, []mod.SalesPlan) {
	TableName := salesPlanTable + Outfit
	Data := []mod.SalesPlan{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("PlanName LIKE ?", "%"+Stext+"%")
	}
	if TargetID > 0 {
		engine = engine.Where("TargetID = ?", TargetID)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
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

	Count := s.Count(db, Stext, TargetID, Status, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (s *SalesPlanDal) All(db *gorm.DB, Order int, Stext string, TargetID int64, Status int64, Outfit string) []mod.SalesPlan {
	TableName := salesPlanTable + Outfit
	Data := []mod.SalesPlan{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.Where("PlanName LIKE ?", "%"+Stext+"%")
	}
	if TargetID > 0 {
		engine = engine.Where("TargetID = ?", TargetID)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
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

func (s *SalesPlanDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := salesPlanTable + Outfit
	Data := mod.SalesPlan{}
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
