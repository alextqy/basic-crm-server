package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

type SalesPlanDal struct{}

func (s *SalesPlanDal) Count(db *xorm.Session, Stext string, TargetID int64, Status int64, Outfit string) (int64, error) {
	TableName := salesPlanTable + Outfit
	Data := mod.SalesPlan{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`PlanName` LIKE ?", "%"+Stext+"%")
	}
	if TargetID > 0 {
		engine = engine.And("`TargetID` = ?", TargetID)
	}
	if Status > 0 {
		engine = engine.And("`Status` = ?", Status)
	}
	r, e := engine.Count(&Data)
	return r, e
}

func (s *SalesPlanDal) Add(db *xorm.Session, Data mod.SalesPlan, Outfit string) (int64, error) {
	TableName := salesPlanTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func (s *SalesPlanDal) Update(db *xorm.Session, Data mod.SalesPlan, Outfit string) (int64, error) {
	TableName := salesPlanTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).AllCols().Update(&Data)
	return r, e
}

func (s *SalesPlanDal) Data(db *xorm.Session, ID int64, Outfit string) (mod.SalesPlan, error) {
	TableName := salesPlanTable + Outfit
	Data := mod.SalesPlan{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func (s *SalesPlanDal) List(db *xorm.Session, Page int, PageSize int, Order int, Stext string, TargetID int64, Status int64, Outfit string) (int, int, int, []mod.SalesPlan) {
	TableName := salesPlanTable + Outfit
	Data := []mod.SalesPlan{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`PlanName` LIKE ?", "%"+Stext+"%")
	}
	if TargetID > 0 {
		engine = engine.And("`TargetID` = ?", TargetID)
	}
	if Status > 0 {
		engine = engine.And("`Status` = ?", Status)
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
	engine.AllCols().OrderBy("`ID` "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)

	Count, _ := s.Count(db, Stext, TargetID, Status, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (s *SalesPlanDal) All(db *xorm.Session, Order int, Stext string, TargetID int64, Status int64, Outfit string) []mod.SalesPlan {
	TableName := salesPlanTable + Outfit
	Data := []mod.SalesPlan{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`PlanName` LIKE ?", "%"+Stext+"%")
	}
	if TargetID > 0 {
		engine = engine.And("`TargetID` = ?", TargetID)
	}
	if Status > 0 {
		engine = engine.And("`Status` = ?", Status)
	}
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.OrderBy("`ID` " + OrderBy).Find(&Data)
	return Data
}

func (s *SalesPlanDal) Del(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := salesPlanTable + Outfit
	Data := mod.SalesPlan{}
	if sysHelper.StringContains(ID, ",") {
		ids := strings.Split(ID, ",")
		intArr := []int{}
		for i := 0; i < len(ids); i++ {
			_, _, n := sysHelper.StringToInt(ids[i])
			intArr = append(intArr, n)
		}
		r, e := db.Table(TableName).In("`ID`", intArr).Delete(Data)
		return r, e
	} else {
		r, e := db.Table(TableName).ID(ID).Delete(Data)
		return r, e
	}
}
