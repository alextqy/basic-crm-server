package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

type SalesTargetDal struct{}

func (s *SalesTargetDal) Count(db *xorm.Session, Stext string, CustomerID int64, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTarget{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`TargetName` LIKE ?", "%"+Stext+"%")
	}
	if CustomerID > 0 {
		engine = engine.And("`CustomerID` = ?", CustomerID)
	}
	r, e := engine.Count(&Data)
	return r, e
}

func (s *SalesTargetDal) Add(db *xorm.Session, Data mod.SalesTarget, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func (s *SalesTargetDal) Update(db *xorm.Session, Data mod.SalesTarget, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).AllCols().Update(&Data)
	return r, e
}

func (s *SalesTargetDal) Data(db *xorm.Session, ID int64, Outfit string) (mod.SalesTarget, error) {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTarget{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func (s *SalesTargetDal) List(db *xorm.Session, Page int, PageSize int, Order int, Stext string, CustomerID int64, Outfit string) (int, int, int, []mod.SalesTarget) {
	TableName := salesTargetTable + Outfit
	Data := []mod.SalesTarget{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`TargetName` LIKE ?", "%"+Stext+"%")
	}
	if CustomerID > 0 {
		engine = engine.And("`CustomerID` = ?", CustomerID)
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

	Count, _ := s.Count(db, Stext, CustomerID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (s *SalesTargetDal) All(db *xorm.Session, Order int, Stext string, CustomerID int64, Outfit string) []mod.SalesTarget {
	TableName := salesTargetTable + Outfit
	Data := []mod.SalesTarget{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`TargetName` LIKE ?", "%"+Stext+"%")
	}
	if CustomerID > 0 {
		engine = engine.And("`CustomerID` = ?", CustomerID)
	}
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.OrderBy("ID " + OrderBy).Find(&Data)
	return Data
}

func (s *SalesTargetDal) Del(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTarget{}
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
