package dal

import (
	mod "basic-crm-server/MOD"
	mtd "basic-crm-server/MTD"
	"math"
	"strings"

	"xorm.io/xorm"
)

func CustomerCount(db *xorm.Session, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64, Outfit string) (int64, error) {
	TableName := customerTable + Outfit
	Data := mod.Customer{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`Name` LIKE ?", "%"+Stext+"%").Or("`Email` LIKE ?", "%"+Stext+"%").Or("`Tel` LIKE ?", "%"+Stext+"%")
	}
	if Gender > 0 {
		engine = engine.And("`Gender` = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.And("`Priority` = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.And("`CompanyID` = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.And("`ManagerID` = ?", ManagerID)
	}
	r, e := engine.Count(&Data)
	return r, e
}

func CustomerAdd(db *xorm.Session, Data mod.Customer, Outfit string) (int64, error) {
	TableName := customerTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func CustomerUpdate(db *xorm.Session, Data mod.Customer, Outfit string) (int64, error) {
	TableName := customerTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).Update(&Data)
	return r, e
}

func CustomerData(db *xorm.Session, ID int64, Outfit string) (mod.Customer, error) {
	TableName := customerTable + Outfit
	Data := mod.Customer{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func CustomerList(db *xorm.Session, Page int, PageSize int, Order int, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64, Outfit string) (int, int, int, []mod.Customer) {
	TableName := customerTable + Outfit
	Data := []mod.Customer{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`Name` LIKE ?", "%"+Stext+"%").Or("`Email` LIKE ?", "%"+Stext+"%").Or("`Tel` LIKE ?", "%"+Stext+"%")
	}
	if Gender > 0 {
		engine = engine.And("`Gender` = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.And("`Priority` = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.And("`CompanyID` = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.And("`ManagerID` = ?", ManagerID)
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
	Count, _ := CustomerCount(db, Stext, Gender, Priority, CompanyID, ManagerID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("`ID` "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func CustomerAll(db *xorm.Session, Order int, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64, Outfit string) []mod.Customer {
	TableName := customerTable + Outfit
	Data := []mod.Customer{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`Name` LIKE ?", "%"+Stext+"%").Or("`Email` LIKE ?", "%"+Stext+"%").Or("`Tel` LIKE ?", "%"+Stext+"%")
	}
	if Gender > 0 {
		engine = engine.And("`Gender` = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.And("`Priority` = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.And("`CompanyID` = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.And("`ManagerID` = ?", ManagerID)
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

func CustomerDel(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := customerTable + Outfit
	Data := mod.Customer{}
	if mtd.StringContains(ID, ",") {
		ids := strings.Split(ID, ",")
		intArr := []int{}
		for i := 0; i < len(ids); i++ {
			_, _, n := mtd.StringToInt(ids[i])
			intArr = append(intArr, n)
		}
		r, e := db.Table(TableName).In("`ID`", intArr).Delete(Data)
		return r, e
	} else {
		r, e := db.Table(TableName).ID(ID).Delete(Data)
		return r, e
	}
}
