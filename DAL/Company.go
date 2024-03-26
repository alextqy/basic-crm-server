package dal

import (
	entity "basic-crm-server/Entity"
	lib "basic-crm-server/LIB"
	"math"
	"strings"

	"xorm.io/xorm"
)

func CompanyCount(db *xorm.Session, Stext string) (int64, error) {
	Data := entity.Company{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("CompanyName LIKE ?", "%"+Stext+"%")
	}
	r, e := engine.Count(&Data)
	return r, e
}

func CompanyAdd(db *xorm.Session, Data entity.Company) (int64, error) {
	r, e := db.Insert(&Data)
	return r, e
}

func CompanyUpdate(db *xorm.Session, Data entity.Company) (int64, error) {
	r, e := db.ID(Data.ID).Update(&Data)
	return r, e
}

func CompanyData(db *xorm.Session, ID int64) (entity.Company, error) {
	Data := entity.Company{}
	_, err := db.ID(ID).Get(&Data)
	return Data, err
}

func CompanyList(db *xorm.Session, Page int, PageSize int, Order int, Stext string) (int, int, int, []entity.Company) {
	Data := []entity.Company{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("CompanyName LIKE ?", "%"+Stext+"%")
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
	Count, _ := CompanyCount(db, Stext)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("ID "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func CompanyAll(db *xorm.Session, Page int, PageSize int, Order int, Stext string) []entity.Company {
	Data := []entity.Company{}
	engine := db.Where("1=1")
	if Stext != "" {
		engine = engine.And("CompanyName LIKE ?", "%"+Stext+"%")
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

func CompanyDel(db *xorm.Session, ID string) (int64, error) {
	Data := entity.Company{}
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
