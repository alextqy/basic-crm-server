package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func DistributorNew(Token, Name, Email, Tel, DistributorInfo string, CompanyID, ManagerID, AfterServiceID, Level, ID int64) mod.Result {
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
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else if Tel == "" {
		result.Message = lang.IncorrectPhoneNumber
	} else {
		db := dal.ConnDB()

		if CompanyID > 0 {
			checkData := companyDal.Data(db, CompanyID, "")
			if checkData.ID == 0 {
				result.Message = lang.CompanyDataDoesNotExist
				return result
			}
		}

		if ManagerID > 0 {
			checkData := managerDal.Data(db, ManagerID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheSalesManagerDoesNotExist
				return result
			}
		}

		if AfterServiceID > 0 {
			checkData := afterServiceDal.Data(db, AfterServiceID, "")
			if checkData.ID == 0 {
				result.Message = lang.AfterSalesPersonnelDoNot
			}
		}

		if ID > 0 {
			checkData := distributorDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.DistributorDataDoesNotExist
			} else {
				checkData.Name = Name
				checkData.Email = Email
				checkData.Tel = Tel
				checkData.DistributorInfo = DistributorInfo
				checkData.CompanyID = CompanyID
				checkData.ManagerID = ManagerID
				checkData.AfterServiceID = AfterServiceID
				checkData.Level = Level
				e := distributorDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			data := mod.DistributorMod{
				Name:            Name,
				Email:           Email,
				Tel:             Tel,
				DistributorInfo: DistributorInfo,
				CreationTime:    sysHelper.TimeStamp(),
				CompanyID:       CompanyID,
				ManagerID:       ManagerID,
				AfterServiceID:  AfterServiceID,
				Level:           Level,
			}
			_, e := distributorDal.Add(db, data, "")
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

func DistributorList(Token string, Page, PageSize, Order int, Stext string, CompanyID, ManagerID, AfterServiceID, Level int64) mod.ResultList {
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
	} else if CheckPerm(t) > 1 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		if CheckPerm(t) == 2 {
			ManagerID = CheckID(t)
		}
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = distributorDal.List(db, Page, PageSize, Order, Stext, CompanyID, ManagerID, AfterServiceID, Level, "")
	}
	return result
}

func DistributorAll(Token string, Order int, Stext string, CompanyID, ManagerID, AfterServiceID, Level int64) mod.Result {
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
	} else {
		if CheckPerm(t) == 2 {
			ManagerID = CheckID(t)
		}
		db := dal.ConnDB()
		result.State = true
		result.Data = distributorDal.All(db, Order, Stext, CompanyID, ManagerID, AfterServiceID, Level, "")
	}
	return result
}

func DistributorData(Token string, ID int64) mod.Result {
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
	} else {
		db := dal.ConnDB()
		data := distributorDal.Data(db, ID, "")
		if CheckPerm(t) == 2 && data.ManagerID != CheckID(t) {
			result.Data = mod.DistributorMod{}
		} else {
			result.Data = data
		}
		result.State = true
	}
	return result
}

func DistributorDel(Token, ID string) mod.Result {
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
	} else {
		db := dal.ConnDB()
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := distributorDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.DistributorDataDoesNotExist
		} else {
			if CheckPerm(t) == 2 {
				if checkData.ManagerID != CheckID(t) {
					result.Message = lang.PermissionDenied
					return result
				}
			}
			e := distributorDal.Del(db, ID, "")
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
