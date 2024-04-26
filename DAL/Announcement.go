package dal

import (
	mod "basic-crm-server/MOD"
	"math"
	"strings"

	"gorm.io/gorm"
)

type AnnouncementDal struct{}

func (o *AnnouncementDal) Count(db *gorm.DB, Stext string, AuthorID, Display int64, Outfit string) int64 {
	var Count int64
	TableName := AnnouncementTable + Outfit
	engine := db.Table(TableName)
	if AuthorID > 0 {
		engine = engine.Where("AuthorID = ?", AuthorID)
	}
	if Display > 0 {
		engine = engine.Where("Display = ?", Display)
	}
	if Stext != "" {
		engine = engine.Where("Title LIKE ?", "%"+Stext+"%")
	}
	engine.Count(&Count)
	return Count
}

func (o *AnnouncementDal) Add(db *gorm.DB, Data mod.AnnouncementMod, Outfit string) (int64, error) {
	TableName := AnnouncementTable + Outfit
	e := db.Table(TableName).Create(&Data).Error
	return Data.ID, e
}

func (o *AnnouncementDal) Update(db *gorm.DB, Data mod.AnnouncementMod, Outfit string) error {
	TableName := AnnouncementTable + Outfit
	return db.Table(TableName).Save(&Data).Error
}

func (o *AnnouncementDal) Data(db *gorm.DB, ID int64, Outfit string) mod.AnnouncementMod {
	TableName := AnnouncementTable + Outfit
	Data := mod.AnnouncementMod{}
	db.Table(TableName).First(&Data, ID)
	return Data
}

func (o *AnnouncementDal) List(db *gorm.DB, Page int, PageSize int, Order int, Stext string, AuthorID, Display int64, Outfit string) (int, int, int, []mod.AnnouncementMod) {
	TableName := AnnouncementTable + Outfit
	Data := []mod.AnnouncementMod{}
	engine := db.Table(TableName)
	if AuthorID > 0 {
		engine = engine.Where("AuthorID = ?", AuthorID)
	}
	if Display > 0 {
		engine = engine.Where("Display = ?", Display)
	}
	if Stext != "" {
		engine = engine.Where("Title LIKE ?", "%"+Stext+"%")
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

	Count := o.Count(db, Stext, AuthorID, Display, "")
	TotalPage := int(math.Ceil(float64(Count) / float64(PageSize)))
	if TotalPage > 0 && Page > TotalPage {
		Page = TotalPage
	}
	return Page, PageSize, TotalPage, Data
}

func (o *AnnouncementDal) All(db *gorm.DB, Order int, Stext string, AuthorID, Display int64, Outfit string) []mod.AnnouncementMod {
	TableName := AnnouncementTable + Outfit
	Data := []mod.AnnouncementMod{}
	engine := db.Table(TableName)
	if AuthorID > 0 {
		engine = engine.Where("AuthorID = ?", AuthorID)
	}
	if Display > 0 {
		engine = engine.Where("Display = ?", Display)
	}
	if Stext != "" {
		engine = engine.Where("Title LIKE ?", "%"+Stext+"%")
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

func (o *AnnouncementDal) Del(db *gorm.DB, ID string, Outfit string) error {
	TableName := AnnouncementTable + Outfit
	Data := mod.AnnouncementMod{}
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
