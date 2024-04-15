package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func CompanyNew(Token, CompanyName, Remark string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" && t.Message != "manager" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if CompanyName == "" {
		result.Message = lang.IncorrectName
	} else {
		db := dal.ConnDB()
		if ID > 0 {
			checkData := companyDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.CompanyDataDoesNotExist
			} else {
				checkData.CompanyName = CompanyName
				checkData.Remark = Remark
				e := companyDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					result.State = true
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), "admin")
				}
			}
		} else {
			data := mod.Company{
				CompanyName:  CompanyName,
				CreationTime: sysHelper.TimeStamp(),
				Remark:       Remark,
			}
			checkData := companyDal.Check(db, CompanyName, "")
			if checkData.ID > 0 {
				result.Message = lang.DataWithTheSameNameExists
			} else {
				_, e := companyDal.Add(db, data, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(data)
					go fileHelper.WriteLog(CheckAccount(t), "Add data: "+string(jData), "admin")
					result.State = true
				}
			}
		}
	}
	return result
}

func CompanyList(Token string, Page, PageSize, Order int, Stext string) mod.ResultList {
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
	} else if t.Message != "admin" && t.Message != "manager" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = companyDal.List(db, Page, PageSize, Order, Stext, "")
	}
	return result
}

func CompanyAll(Token string, Order int, Stext string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" && t.Message != "manager" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = companyDal.All(db, Order, Stext, "")
	}
	return result
}

func CompanyData(Token string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" && t.Message != "manager" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = companyDal.Data(db, ID, "")
	}
	return result
}

func CompanyDel(Token, ID string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" && t.Message != "manager" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := companyDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.CompanyDataDoesNotExist
		} else {
			e := companyDal.Del(db, ID, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Remove data: "+string(jData), "admin")
				result.State = true
			}
		}
	}
	return result
}
