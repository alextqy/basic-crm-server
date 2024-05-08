package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func CustomerQANew(Token, Title, Content string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 3 {
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
			checkData := customerQADal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.QADataDoesNotExist
			} else {
				checkData.Title = Title
				checkData.Content = Content
				e := customerQADal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			data := mod.CustomerQAMod{
				Title:        Title,
				Content:      Content,
				Display:      1,
				CreationTime: sysHelper.TimeStamp(),
			}
			_, e := customerQADal.Add(db, data, "")
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

func CustomerQAList(Token string, Page, PageSize, Order int, Stext string, Display int64) mod.ResultList {
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
	} else if CheckPerm(t) > 3 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = customerQADal.List(db, Page, PageSize, Order, Stext, Display, "")
	}
	return result
}

func CustomerQAAll(Token string, Order int, Stext string, Display int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 3 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = customerQADal.All(db, Order, Stext, Display, "")
	}
	return result
}

func CustomerQAData(Token string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 3 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = customerQADal.Data(db, ID, "")
	}
	return result
}

func CustomerQADel(Token, ID string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 3 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := customerQADal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.QADataDoesNotExist
		} else {
			e := customerQADal.Del(db, ID, "")
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

func CustomerQADisplay(Token string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) > 3 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		checkData := customerQADal.Data(db, ID, "")
		if checkData.ID == 0 {
			result.Message = lang.QADataDoesNotExist
		} else {
			if checkData.Display == 1 {
				checkData.Display = 2
			} else {
				checkData.Display = 1
			}

			e := customerQADal.Update(db, checkData, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				result.State = true
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
			}
		}
	}
	return result
}
