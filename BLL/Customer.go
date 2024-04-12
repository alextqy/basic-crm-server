package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func CustomerNew(Token, Name string, Birthday, Gender int64, Email, Tel, CustomerInfo string, Priority, CompanyID, ID int64) mod.Result {
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
			checkCompany := companyDal.Data(db, CompanyID, "")
			if checkCompany.ID == 0 {
				result.Message = lang.CompanyDataDoesNotExist
				return result
			}
		}

		var ManagerID int64
		if t.Message == "manager" {
			ManagerID = customerDal.Data(db, t.Data.(mod.Manager).ID, "").ID
		} else {
			ManagerID = 0
		}

		if ID > 0 {
			checkCustomer := customerDal.Data(db, ID, "")
			if checkCustomer.ID == 0 {
				result.Message = lang.CustomerDataDoesNotExist
			} else {
				checkCustomer.Name = Name
				checkCustomer.Birthday = Birthday
				checkCustomer.Gender = Gender
				checkCustomer.Email = Email
				checkCustomer.Tel = Tel
				checkCustomer.CustomerInfo = CustomerInfo
				checkCustomer.Priority = Priority
				checkCustomer.CompanyID = CompanyID
				e := customerDal.Update(db, checkCustomer, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkCustomer)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData))
					result.State = true
				}
			}
		} else {
			data := mod.Customer{
				Name:         Name,
				Birthday:     Birthday,
				Gender:       Gender,
				Email:        Email,
				Tel:          Tel,
				CustomerInfo: CustomerInfo,
				Priority:     Priority,
				CreationTime: sysHelper.TimeStamp(),
				CompanyID:    CompanyID,
				ManagerID:    ManagerID,
			}
			_, e := customerDal.Add(db, data, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(data)
				go fileHelper.WriteLog(CheckAccount(t), "Add data: "+string(jData))
				result.State = true
			}
		}
	}
	return result
}

func CustomerList(Token string, Page, PageSize, Order int, Stext string, Gender, Priority, CompanyID, ManagerID int64) mod.ResultList {
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
		result.Page, result.PageSize, result.TotalPage, result.Data = customerDal.List(db, Page, PageSize, Order, Stext, Gender, Priority, CompanyID, ManagerID, "")
	}
	return result
}

func CustomerAll(Token string, Order int, Stext string, Gender, Priority, CompanyID, ManagerID int64) mod.Result {
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
		result.Data = customerDal.All(db, Order, Stext, Gender, Priority, CompanyID, ManagerID, "")
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
	} else if t.Message != "admin" && t.Message != "manager" {
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
	} else if t.Message != "admin" && t.Message != "manager" {
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
				go fileHelper.WriteLog(CheckAccount(t), "Remove data: "+string(jData))
				result.State = true
			}
		}
	}
	return result
}
