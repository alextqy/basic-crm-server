package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

type ManagerDal struct{}

func (m *ManagerDal) Count(db *xorm.Session, Stext string, Level int64, Status int64, GroupID int64, Outfit string) (int64, error) {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`Account` LIKE ?", "%"+Stext+"%").Or("`Name` LIKE ?", "%"+Stext+"%")
	}
	if Level > 0 {
		engine = engine.And("`Level` = ?", Level)
	}
	if Status > 0 {
		engine = engine.And("`Status` = ?", Status)
	}
	if GroupID > 0 {
		engine = engine.And("`GroupID` = ?", GroupID)
	}
	r, e := engine.Count(&Data)
	return r, e
}

func (m *ManagerDal) Add(db *xorm.Session, Data mod.Manager, Outfit string) (int64, error) {
	TableName := managerTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func (m *ManagerDal) Update(db *xorm.Session, Data mod.Manager, Outfit string) (int64, error) {
	TableName := managerTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).AllCols().Update(&Data)
	return r, e
}

func (m *ManagerDal) Data(db *xorm.Session, ID int64, Outfit string) (mod.Manager, error) {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func (m *ManagerDal) List(db *xorm.Session, Page int, PageSize int, Order int, Stext string, Level int64, Status int64, GroupID int64, Outfit string) (int, int, int, []mod.Manager) {
	TableName := managerTable + Outfit
	Data := []mod.Manager{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`Account` LIKE ?", "%"+Stext+"%").Or("`Name` LIKE ?", "%"+Stext+"%")
	}
	if Level > 0 {
		engine = engine.And("`Level` = ?", Level)
	}
	if Status > 0 {
		engine = engine.And("`Status` = ?", Status)
	}
	if GroupID > 0 {
		engine = engine.And("`GroupID` = ?", GroupID)
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
	Count, _ := m.Count(db, Stext, Level, Status, GroupID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("`ID` "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func (m *ManagerDal) All(db *xorm.Session, Order int, Stext string, Level int64, Status int64, GroupID int64, Outfit string) []mod.Manager {
	TableName := managerTable + Outfit
	Data := []mod.Manager{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`Account` LIKE ?", "%"+Stext+"%").Or("`Name` LIKE ?", "%"+Stext+"%")
	}
	if Level > 0 {
		engine = engine.And("`Level` = ?", Level)
	}
	if Status > 0 {
		engine = engine.And("`Status` = ?", Status)
	}
	if GroupID > 0 {
		engine = engine.And("`GroupID` = ?", GroupID)
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

func (m *ManagerDal) Del(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
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

func (a *ManagerDal) Check(db *xorm.Session, Account, Outfit string) (mod.Manager, error) {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	_, err := db.Table(TableName).Where("`Account` = ?", Account).Get(&Data)
	return Data, err
}

func (a *ManagerDal) Token(db *xorm.Session, Token, Outfit string) (mod.Manager, error) {
	TableName := managerTable + Outfit
	Data := mod.Manager{}
	_, err := db.Table(TableName).Where("`Token` = ?", Token).Get(&Data)
	return Data, err
}
