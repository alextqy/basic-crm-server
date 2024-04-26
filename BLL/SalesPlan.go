package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func SalesPlanNew(Token, PlanName string, TargetID int64, PlanContent string, Budget float32, ID int64) mod.Result {
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
	} else if PlanName == "" {
		result.Message = lang.IncorrectName
	} else if TargetID == 0 {
		result.Message = lang.IncorrectSalesTarget
	} else {
		db := dal.ConnDB()

		if ID > 0 {
			checkData := salesPlanDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheSalesPlanDataDoesNotExist
			} else {
				checkData.PlanName = PlanName
				checkData.PlanContent = PlanContent
				checkData.Budget = Budget
				e := salesPlanDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			checkData := salesTargetDal.Data(db, TargetID, "")
			if checkData.ID == 0 {
				result.Message = lang.SalesTargetDataDoesNotExist
			} else if CheckPerm(t) == 2 && checkData.ManagerID != CheckID(t) {
				result.Message = lang.PermissionDenied
			} else {
				data := mod.SalesPlanMod{
					PlanName:     PlanName,
					TargetID:     TargetID,
					PlanContent:  PlanContent,
					CreationTime: sysHelper.TimeStamp(),
					Status:       1,
					Budget:       Budget,
					ManagerID:    checkData.ManagerID,
				}
				_, e := salesPlanDal.Add(db, data, "")
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

func SalesPlanList(Token string, Page, PageSize, Order int, Stext string, TargetID, Status, ManagerID int64) mod.ResultList {
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
		if CheckPerm(t) == 2 {
			ManagerID = CheckID(t)
		}
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = salesPlanDal.List(db, Page, PageSize, Order, Stext, TargetID, Status, ManagerID, "")
	}
	return result
}

func SalesPlanAll(Token string, Order int, Stext string, TargetID, Status, ManagerID int64) mod.Result {
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
		if CheckPerm(t) == 2 {
			ManagerID = CheckID(t)
		}
		db := dal.ConnDB()
		result.State = true
		result.Data = salesPlanDal.All(db, Order, Stext, TargetID, Status, ManagerID, "")
	}
	return result
}

func SalesPlanData(Token string, ID int64) mod.Result {
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
		data := salesPlanDal.Data(db, ID, "")
		if CheckPerm(t) == 2 && data.ManagerID != CheckID(t) {
			result.Data = mod.SalesPlanMod{}
		} else {
			result.Data = data
		}
		result.State = true
	}
	return result
}

func SalesPlanDel(Token string, ID string) mod.Result {
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
		checkData := salesPlanDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.TheSalesPlanDataDoesNotExist
		} else {
			if CheckPerm(t) == 2 {
				if checkData.ManagerID != CheckID(t) {
					result.Message = lang.PermissionDenied
					return result
				}
			}
			e := salesPlanDal.Del(db, ID, "")
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
