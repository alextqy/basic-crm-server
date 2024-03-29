package dal

import (
	mod "basic-crm-server/MOD"
	mtd "basic-crm-server/MTD"
	"math"
	"strings"

	"xorm.io/xorm"
)

func SalesTargetCount(db *xorm.Session, Stext string, CustomerID int64, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTarget{}
	engine := db.Table(TableName).Where("1=1")
	if Stext != "" {
		engine = engine.And("`TargetName` LIKE ?", "%"+Stext+"%")
	}
	if CustomerID > 0 {
		engine = engine.And("`CustomerID` = ?", CustomerID)
	}
	r, e := engine.Count(&Data)
	return r, e
}

func SalesTargetAdd(db *xorm.Session, Data mod.SalesTarget, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	r, e := db.Table(TableName).Insert(&Data)
	return r, e
}

func SalesTargetUpdate(db *xorm.Session, Data mod.SalesTarget, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	r, e := db.Table(TableName).ID(Data.ID).Update(&Data)
	return r, e
}

func SalesTargetData(db *xorm.Session, ID int64, Outfit string) (mod.SalesTarget, error) {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTarget{}
	_, err := db.Table(TableName).ID(ID).Get(&Data)
	return Data, err
}

func SalesTargetList(db *xorm.Session, Page int, PageSize int, Order int, Stext string, CustomerID int64, Outfit string) (int, int, int, []mod.SalesTarget) {
	TableName := salesTargetTable + Outfit
	Data := []mod.SalesTarget{}
	engine := db.Table(TableName).Where("1=1")
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
	Count, _ := SalesTargetCount(db, Stext, CustomerID, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	engine.OrderBy("`ID` "+OrderBy).Limit(int(PageSize), int((Page-1)*PageSize)).Find(&Data)
	return Page, PageSize, TotalPage, Data
}

func SalesTargetAll(db *xorm.Session, Order int, Stext string, CustomerID int64, Outfit string) []mod.SalesTarget {
	TableName := salesTargetTable + Outfit
	Data := []mod.SalesTarget{}
	engine := db.Table(TableName).Where("1=1")
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

func SalesTargetDel(db *xorm.Session, ID string, Outfit string) (int64, error) {
	TableName := salesTargetTable + Outfit
	Data := mod.SalesTarget{}
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
