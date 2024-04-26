package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func AnnouncementNew(Token, Title, Content string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 1 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if Title == "" {
		result.Message = lang.IncorrectTitle
	} else if Content == "" {
		result.Message = lang.IncorrectContent
	} else {
		db := dal.ConnDB()
		if ID > 0 {
			checkData := announcementDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.AnnouncementDataDoesNotExist
			} else {
				checkData.Title = Title
				checkData.Content = Content
				e := announcementDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			data := mod.AnnouncementMod{
				Title:        Title,
				Content:      Content,
				AuthorID:     CheckID(t),
				Display:      1,
				CreationTime: sysHelper.TimeStamp(),
			}
			_, e := announcementDal.Add(db, data, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(data)
				go fileHelper.WriteLog(CheckAccount(t), "Add data: "+string(jData), t.Message)
				result.State = true
			}
		}
	}
	return result
}

func AnnouncementList(Token string, Page, PageSize, Order int, Stext string, AuthorID, Display int64) mod.ResultList {
	result := mod.ResultList{
		State:     false,
		Code:      200,
		Message:   "",
		Page:      0,
		PageSize:  0,
		TotalPage: 0,
		Data:      nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 2 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = announcementDal.List(db, Page, PageSize, Order, Stext, AuthorID, Display, "")
	}
	return result
}

func AnnouncementAll(Token string, Order int, Stext string, AuthorID, Display int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 2 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = announcementDal.All(db, Order, Stext, AuthorID, Display, "")
	}
	return result
}

func AnnouncementData(Token string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 2 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = announcementDal.Data(db, ID, "")
	}
	return result
}

func AnnouncementDel(Token, ID string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 2 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := announcementDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.AnnouncementDataDoesNotExist
		} else {
			e := announcementDal.Del(db, ID, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Remove data: "+string(jData), t.Message)
				result.State = true
			}
		}
	}
	return result
}
