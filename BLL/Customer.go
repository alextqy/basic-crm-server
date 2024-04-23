package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func CustomerNew(Token, Name string, Birthday, Gender int64, Email, Tel, CustomerInfo string, Priority, CompanyID, AfterServiceID, ID int64) mod.Result {
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
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else if Gender <= 0 {
		result.Message = lang.IncorrectGender
	} else if Tel == "" {
		result.Message = lang.IncorrectPhoneNumber
	} else if Priority <= 0 {
		result.Message = lang.IncorrectPriority
	} else {
		db := dal.ConnDB()

		if CompanyID > 0 {
			checkData := companyDal.Data(db, CompanyID, "")
			if checkData.ID == 0 {
				result.Message = lang.CompanyDataDoesNotExist
				return result
			}
		}

		if AfterServiceID > 0 {
			checkData := afterServiceDal.Data(db, AfterServiceID, "")
			if checkData.ID == 0 {
				result.Message = lang.AfterSalesPersonnelDoNot
			}
		}

		var ManagerID int64
		if t.Message == "manager" {
			ManagerID = customerDal.Data(db, t.Data.(mod.Manager).ID, "").ID
		} else {
			ManagerID = 0
		}

		if ID > 0 {
			checkData := customerDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.CustomerDataDoesNotExist
			} else {
				checkData.Name = Name
				checkData.Birthday = Birthday
				checkData.Gender = Gender
				checkData.Email = Email
				checkData.Tel = Tel
				checkData.CustomerInfo = CustomerInfo
				checkData.Priority = Priority
				checkData.CompanyID = CompanyID
				checkData.AfterServiceID = AfterServiceID
				e := customerDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			data := mod.Customer{
				Name:           Name,
				Birthday:       Birthday,
				Gender:         Gender,
				Email:          Email,
				Tel:            Tel,
				CustomerInfo:   CustomerInfo,
				Priority:       Priority,
				CreationTime:   sysHelper.TimeStamp(),
				CompanyID:      CompanyID,
				ManagerID:      ManagerID,
				AfterServiceID: AfterServiceID,
			}
			_, e := customerDal.Add(db, data, "")
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

func CustomerList(Token string, Page, PageSize, Order int, Stext string, Gender, Priority, CompanyID, ManagerID, AfterServiceID int64) mod.ResultList {
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
		if t.Message == "manager" {
			ManagerID = CheckID(t)
		}
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = customerDal.List(db, Page, PageSize, Order, Stext, Gender, Priority, CompanyID, ManagerID, AfterServiceID, "")
	}
	return result
}

func CustomerAll(Token string, Order int, Stext string, Gender, Priority, CompanyID, ManagerID, AfterServiceID int64) mod.Result {
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
		if t.Message == "manager" {
			ManagerID = CheckID(t)
		}
		db := dal.ConnDB()
		result.State = true
		result.Data = customerDal.All(db, Order, Stext, Gender, Priority, CompanyID, ManagerID, AfterServiceID, "")
	}
	return result
}

func CustomerData(Token string, ID int64) mod.Result {
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
		result.Data = customerDal.Data(db, ID, "")
	}
	return result
}

func CustomerDel(Token, ID string) mod.Result {
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
		checkData := customerDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.CustomerDataDoesNotExist
		} else {
			e := customerDal.Del(db, ID, "")
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
