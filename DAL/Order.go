package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type OrderDal struct{}

func (o *OrderDal) Count(db *gorm.DB, Stext string, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review int64, Outfit string) int64 {
	var Count int64
	TableName := orderTable + Outfit
	engine := db.Table(TableName)
	if ProductID > 0 {
		engine = engine.Where("ProductID = ?", ProductID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if CustomerID > 0 {
		engine = engine.Where("CustomerID = ?", CustomerID)
	}
	if DistributorID > 0 {
		engine = engine.Where("DistributorID = ?", DistributorID)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if OrderType > 0 {
		engine = engine.Where("OrderType = ?", OrderType)
	}
	if Payment > 0 {
		engine = engine.Where("Payment = ?", Payment)
	}
	if Review > 0 {
		engine = engine.Where("Review = ?", Review)
	}
	if Stext != "" {
		engine = engine.Where("OrderNo LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *OrderDal) Add(db *gorm.DB, Data mod.OrderMod, Outfit string) (int64, error) {
	TableName := orderTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *OrderDal) Update(db *gorm.DB, Data mod.OrderMod, Outfit string) error {
	TableName := orderTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *OrderDal) Data(db *gorm.DB, ID int64, Outfit string) mod.OrderMod {
	TableName := orderTable + Outfit
	Data := mod.OrderMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *OrderDal) List(db *gorm.DB, Page, PageSize, Order int, Stext string, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review int64, Outfit string) (int, int, int, []mod.OrderMod) {
	TableName := orderTable + Outfit
	Data := []mod.OrderMod{}
	engine := db.Table(TableName)
	if ProductID > 0 {
		engine = engine.Where("ProductID = ?", ProductID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if CustomerID > 0 {
		engine = engine.Where("CustomerID = ?", CustomerID)
	}
	if DistributorID > 0 {
		engine = engine.Where("DistributorID = ?", DistributorID)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if OrderType > 0 {
		engine = engine.Where("OrderType = ?", OrderType)
	}
	if Payment > 0 {
		engine = engine.Where("Payment = ?", Payment)
	}
	if Review > 0 {
		engine = engine.Where("Review = ?", Review)
	}
	if Stext != "" {
		engine = engine.Where("OrderNo LIKE ?", "%"+Stext+"%")
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
	engine.Order("ID " + OrderBy).Limit(int(PageSize)).Offset(int((Page - 1) * PageSize)).Find(&Data)

	Count := o.Count(db, Stext, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review, Outfit)
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *OrderDal) All(db *gorm.DB, Order int, Stext string, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review int64, Outfit string) []mod.OrderMod {
	TableName := orderTable + Outfit
	Data := []mod.OrderMod{}
	engine := db.Table(TableName)
	if ProductID > 0 {
		engine = engine.Where("ProductID = ?", ProductID)
	}
	if ManagerID > 0 {
		engine = engine.Where("ManagerID = ?", ManagerID)
	}
	if CustomerID > 0 {
		engine = engine.Where("CustomerID = ?", CustomerID)
	}
	if DistributorID > 0 {
		engine = engine.Where("DistributorID = ?", DistributorID)
	}
	if Status > 0 {
		engine = engine.Where("Status = ?", Status)
	}
	if OrderType > 0 {
		engine = engine.Where("OrderType = ?", OrderType)
	}
	if Payment > 0 {
		engine = engine.Where("Payment = ?", Payment)
	}
	if Review > 0 {
		engine = engine.Where("Review = ?", Review)
	}
	if Stext != "" {
		engine = engine.Where("OrderNo LIKE ?", "%"+Stext+"%")
	}
	OrderBy := ""
	if Order == -1 {
		OrderBy = "DESC"
	} else {
		OrderBy = "ASC"
	}
	engine.Order("ID " + OrderBy).Find(&Data)
	return Data
}

func (o *OrderDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := orderTable + Outfit
	Data := mod.OrderMod{}
	var e error
	if sysHelper.StringContains(ID, ",") {
		ids := strings.Split(ID, ",")
		intArr := []int{}
		for i := 0; i < len(ids); i++ {
			_, _, n := sysHelper.StringToInt(ids[i])
			intArr = append(intArr, n)
		}
		e = db.Table(TableName).Delete(Data, intArr).Error
	} else {
		e = db.Table(TableName).Delete(Data, ID).Error
	}
	return e
}

func (o *OrderDal) Check(db *gorm.DB, OrderNo, Outfit string) mod.OrderMod {
	TableName := orderTable + Outfit
	Data := mod.OrderMod{}
	db.Table(TableName).Where("OrderNo = ?", OrderNo).First(&Data)
	return Data
}
