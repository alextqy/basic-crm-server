package dal

import (
	lib "basic-crm-server/LIB"
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

func AdminCount(db *xorm.Session, Stext string, Level int64, Status int64, Outfit string) (int64, error) {
	TableName := AdminTable + Outfit
	Data := mod.Admin{}
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
	r, e := engine.Count(&Data)
	return r, e
}

func AdminAdd(db *xorm.Session, Data mod.Admin, Outfit string) (int64, error) {
	TableName := AdminTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func AdminUpdate(db *xorm.Session, Data mod.Admin, Outfit string) (int64, error) {
	TableName := AdminTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).Update(&Data)
	return r, e
}

func AdminData(db *xorm.Session, ID int64, Outfit string) (mod.Admin, error) {
	TableName := AdminTable + Outfit
	Data := mod.Admin{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func AdminList(db *xorm.Session, Page int, PageSize int, Order int, Stext string, Level int64, Status int64, Outfit string) (int, int, int, []mod.Admin) {
	TableName := AdminTable + Outfit
	Data := []mod.Admin{}
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
	Count, _ := AdminCount(db, Stext, Level, Status, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("ID "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func AdminAll(db *xorm.Session, Order int, Stext string, Level int64, Status int64, Outfit string) []mod.Admin {
	TableName := AdminTable + Outfit
	Data := []mod.Admin{}
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
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.OrderBy("`ID` " + OrderBy).Find(&Data)
	return Data
}

func AdminDel(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := AdminTable + Outfit
	Data := mod.Admin{}
	if lib.StringContains(ID, ",") {
		ids := strings.Split(ID, ",")
		intArr := []int{}
		for i := 0; i < len(ids); i++ {
			_, _, n := lib.StringToInt(ids[i])
			intArr = append(intArr, n)
		}
		r, e := db.Table(TableName).In("`ID`", intArr).Delete(Data)
		return r, e
	} else {
		r, e := db.Table(TableName).ID(ID).Delete(Data)
		return r, e
	}
}

func AdminCheck(db *xorm.Session, Account, Password string, Outfit string) (mod.Admin, error) {
	TableName := AdminTable + Outfit
	Data := mod.Admin{}
	_, err := db.Table(TableName).Where("`Account` = ?", Account).Where("`Password` = ?", Password).Get(&Data)
	return Data, err
}
