package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func SalesTargetNew(Token, TargetName string, ExpirationDate, CustomerID int64, Remark string, ID int64) mod.Result {
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
	} else if TargetName == "" {
		result.Message = lang.IncorrectName
	} else if ExpirationDate == 0 {
		result.Message = lang.IncorrectExpirationDate
	} else if CustomerID == 0 {
		result.Message = lang.IncorrectCustomer
	} else {
		db := dal.ConnDB()

		if ID > 0 {
			checkData := salesTargetDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.SalesTargetDataDoesNotExist
			} else {
				checkData.TargetName = TargetName
				checkData.ExpirationDate = ExpirationDate
				checkData.Remark = Remark
				e := salesTargetDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			checkData := customerDal.Data(db, CustomerID, "")
			if checkData.ID == 0 {
				result.Message = lang.CustomerDataDoesNotExist
			} else {
				data := mod.SalesTarget{
					TargetName:      TargetName,
					ExpirationDate:  ExpirationDate,
					CreationTime:    sysHelper.TimeStamp(),
					AchievementRate: 0,
					CustomerID:      CustomerID,
					Remark:          Remark,
				}
				_, e := salesTargetDal.Add(db, data, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(data)
					go fileHelper.WriteLog(CheckAccount(t), "Add data: "+string(jData), t.Message)
					result.State = true
				}
			}
		}
	}
	return result
}

func SalesTargetList(Token string, Page, PageSize, Order int, Stext string, CustomerID int64) mod.ResultList {
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
		result.Page, result.PageSize, result.TotalPage, result.Data = salesTargetDal.List(db, Page, PageSize, Order, Stext, CustomerID, "")
	}
	return result
}

func SalesTargetAll(Token string, Order int, Stext string, CustomerID int64) mod.Result {
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
		result.Data = salesTargetDal.All(db, Order, Stext, CustomerID, "")
	}
	return result
}

func SalesTargetData(Token string, ID int64) mod.Result {
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
		result.Data = salesTargetDal.Data(db, ID, "")
	}
	return result
}

func SalesTargetDel(Token, ID string) mod.Result {
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
		checkData := salesTargetDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.SalesTargetDataDoesNotExist
		} else {
			db.Begin()

			planList := salesPlanDal.All(db, 0, "", ID64, 0, "")
			if len(planList) > 0 {
				for i := 0; i < len(planList); i++ {
					ID := sysHelper.Int64ToString(planList[i].ID)
					e := salesPlanDal.Del(db, ID, "")
					if e != nil {
						db.Rollback()
						result.Message = e.Error()
						return result
					}
				}
			}

			e := salesTargetDal.Del(db, ID, "")
			if e != nil {
				result.Message = e.Error()
				db.Rollback()
			} else {
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Remove data: "+string(jData), t.Message)
				result.State = true
				db.Commit()
			}
		}
	}
	return result
}
