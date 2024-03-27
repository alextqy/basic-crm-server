package dal

import (
	lib "basic-crm-server/LIB"
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

func CustomerCount(db *xorm.Session, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64) (int64, error) {
	Data := mod.Customer{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
	}
	if Gender > 0 {
		engine = engine.And("Gender = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.And("Priority = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.And("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.And("ManagerID = ?", ManagerID)
	}
	r, e := engine.Count(&Data)
	return r, e
}

func CustomerAdd(db *xorm.Session, Data mod.Customer) (int64, error) {
	r, e := db.Insert(&Data)
	return r, e
}

func CustomerUpdate(db *xorm.Session, Data mod.Customer) (int64, error) {
	r, e := db.ID(Data.ID).Update(&Data)
	return r, e
}

func CustomerData(db *xorm.Session, ID int64) (mod.Customer, error) {
	Data := mod.Customer{}
	_, err := db.ID(ID).Get(&Data)
	return Data, err
}

func CustomerList(db *xorm.Session, Page int, PageSize int, Order int, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64) (int, int, int, []mod.Customer) {
	Data := []mod.Customer{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
	}
	if Gender > 0 {
		engine = engine.And("Gender = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.And("Priority = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.And("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.And("ManagerID = ?", ManagerID)
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
	Count, _ := CustomerCount(db, Stext, Gender, Priority, CompanyID, ManagerID)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("ID "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func CustomerAll(db *xorm.Session, Order int, Stext string, Gender int64, Priority int64, CompanyID int64, ManagerID int64) []mod.Customer {
	Data := []mod.Customer{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("Name LIKE ?", "%"+Stext+"%").Or("Email LIKE ?", "%"+Stext+"%").Or("Tel LIKE ?", "%"+Stext+"%")
	}
	if Gender > 0 {
		engine = engine.And("Gender = ?", Gender)
	}
	if Priority > 0 {
		engine = engine.And("Priority = ?", Priority)
	}
	if CompanyID > 0 {
		engine = engine.And("CompanyID = ?", CompanyID)
	}
	if ManagerID > 0 {
		engine = engine.And("ManagerID = ?", ManagerID)
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

func CustomerDel(db *xorm.Session, ID string) (int64, error) {
	Data := mod.Customer{}
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
