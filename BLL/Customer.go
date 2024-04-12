package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func CustomerNew(Token, Name string, Birthday, Gender int64, Email, Tel, CustomerInfo string, Priority, CreationTime, CompanyID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else if Birthday <= 0 {
		result.Message = lang.IncorrectBirthday
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
	return result
}

func CustomerUpdate(Token string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	return result
}

func CustomerList(Token string) mod.ResultList {
	result := mod.ResultList{
		State:     false,
		Code:      200,
		Message:   "",
		Page:      0,
		PageSize:  0,
		TotalPage: 0,
		Data:      nil,
	}

	return result
}

func CustomerAll(Token string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	return result
}

func CustomerDel(Token string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	return result
}
