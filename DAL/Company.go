package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"xorm.io/xorm"
)

type CompanyDal struct{}

func (c *CompanyDal) Count(db *xorm.Session, Stext string, Outfit string) (int64, error) {
	TableName := companyTable + Outfit
	Data := mod.Company{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`CompanyName` LIKE ?", "%"+Stext+"%")
	}
	r, e := engine.Count(&Data)
	return r, e
}

func (c *CompanyDal) Add(db *xorm.Session, Data mod.Company, Outfit string) (int64, error) {
	TableName := companyTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func (c *CompanyDal) Update(db *xorm.Session, Data mod.Company, Outfit string) (int64, error) {
	TableName := companyTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).AllCols().Update(&Data)
	return r, e
}

func (c *CompanyDal) Data(db *xorm.Session, ID int64, Outfit string) (mod.Company, error) {
	TableName := companyTable + Outfit
	Data := mod.Company{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func (c *CompanyDal) List(db *xorm.Session, Page int, PageSize int, Order int, Stext string, Outfit string) (int, int, int, []mod.Company) {
	TableName := companyTable + Outfit
	Data := []mod.Company{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`CompanyName` LIKE ?", "%"+Stext+"%")
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
	Count, _ := c.Count(db, Stext, TableName)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("`ID` "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func (c *CompanyDal) All(db *xorm.Session, Order int, Stext string, Outfit string) []mod.Company {
	TableName := companyTable + Outfit
	Data := []mod.Company{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`CompanyName` LIKE ?", "%"+Stext+"%")
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

func (c *CompanyDal) Del(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := companyTable + Outfit
	Data := mod.Company{}
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
