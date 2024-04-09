package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

type ManagerGroupDal struct{}

func (m *ManagerGroupDal) Count(db *xorm.Session, Stext string, Outfit string) (int64, error) {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroup{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`GroupName` LIKE ?", "%"+Stext+"%")
	}
	r, e := engine.Count(&Data)
	return r, e
}

func (m *ManagerGroupDal) Add(db *xorm.Session, Data mod.ManagerGroup, Outfit string) (int64, error) {
	TableName := managerGroupTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func (m *ManagerGroupDal) Update(db *xorm.Session, Data mod.ManagerGroup, Outfit string) (int64, error) {
	TableName := managerGroupTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).AllCols().Update(&Data)
	return r, e
}

func (m *ManagerGroupDal) Data(db *xorm.Session, ID int64, Outfit string) (mod.ManagerGroup, error) {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroup{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func (m *ManagerGroupDal) List(db *xorm.Session, Page int, PageSize int, Order int, Stext string, Outfit string) (int, int, int, []mod.ManagerGroup) {
	TableName := managerGroupTable + Outfit
	Data := []mod.ManagerGroup{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`GroupName` LIKE ?", "%"+Stext+"%")
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

	Count, _ := m.Count(db, Stext, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (m *ManagerGroupDal) All(db *xorm.Session, Order int, Stext string, Outfit string) []mod.ManagerGroup {
	TableName := managerGroupTable + Outfit
	Data := []mod.ManagerGroup{}
	engine := db.Table(TableName)
	if Stext != "" {
		engine = engine.And("`GroupName` LIKE ?", "%"+Stext+"%")
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

func (m *ManagerGroupDal) Del(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := managerGroupTable + Outfit
	Data := mod.ManagerGroup{}
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
