package dal

import (
	lib "basic-crm-server/LIB"
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

func ManagerGroupCount(db *xorm.Session, Stext string) (int64, error) {
	Data := mod.ManagerGroup{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("GroupName LIKE ?", "%"+Stext+"%")
	}
	r, e := engine.Count(&Data)
	return r, e
}

func ManagerGroupAdd(db *xorm.Session, Data mod.ManagerGroup) (int64, error) {
	r, e := db.Insert(&Data)
	return r, e
}

func ManagerGroupUpdate(db *xorm.Session, Data mod.ManagerGroup) (int64, error) {
	r, e := db.ID(Data.ID).Update(&Data)
	return r, e
}

func ManagerGroupData(db *xorm.Session, ID int64) (mod.ManagerGroup, error) {
	Data := mod.ManagerGroup{}
	_, err := db.ID(ID).Get(&Data)
	return Data, err
}

func ManagerGroupList(db *xorm.Session, Page int, PageSize int, Order int, Stext string) (int, int, int, []mod.ManagerGroup) {
	Data := []mod.ManagerGroup{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("GroupName LIKE ?", "%"+Stext+"%")
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
	Count, _ := ManagerGroupCount(db, Stext)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("ID "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func ManagerGroupAll(db *xorm.Session, Order int, Stext string) []mod.ManagerGroup {
	Data := []mod.ManagerGroup{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("GroupName LIKE ?", "%"+Stext+"%")
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

func ManagerGroupDel(db *xorm.Session, ID string) (int64, error) {
	Data := mod.ManagerGroup{}
	if lib.StringContains(ID, ",") {
		ids := strings.Split(ID, ",")
		intArr := []int{}
		for i := 0; i < len(ids); i++ {
			_, _, n := lib.StringToInt(ids[i])
			intArr = append(intArr, n)
		}
		r, e := db.In("ID", intArr).Delete(Data)
		return r, e
	} else {
		r, e := db.ID(ID).Delete(Data)
		return r, e
	}
}
