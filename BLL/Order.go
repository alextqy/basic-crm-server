package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func OrderNew(Token, OrderNo string, ProductID, ManagerID int64, OrderPrice float32, Remark string, ID int64) mod.Result {
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
	} else if OrderNo == "" {
		result.Message = lang.IncorrectOrderNo
	} else if ProductID == 0 {
		result.Message = lang.TheProductDataDoesNotExist
	} else if ManagerID == 0 {
		result.Message = lang.TheSalesManagerDoesNotExist
	} else if OrderPrice == 0 {
		result.Message = lang.IncorrectOrderPrice
	} else {
		db := dal.ConnDB()

		productData := productDal.Data(db, ProductID, "")
		if productData.ID == 0 {
			result.Message = lang.TheProductDataDoesNotExist
			return result
		}

		managerData := managerDal.Data(db, ManagerID, "")
		if managerData.ID == 0 {
			result.Message = lang.TheSalesManagerDoesNotExist
			return result
		}

		if ID > 0 {
			checkData := orderDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheOrderDataDoesNotExist
			} else {
				checkData.ManagerID = ManagerID
				checkData.OrderPrice = OrderPrice
				checkData.Remark = Remark
				e := orderDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			checkData := orderDal.Check(db, OrderNo, "")
			if checkData.ID > 0 {
				result.Message = lang.TheOrderNumberIsDuplicated
			} else {
				data := mod.Order{
					OrderNo:      OrderNo,
					ProductID:    ProductID,
					ManagerID:    ManagerID,
					OrderPrice:   OrderPrice,
					ProductPrice: productData.Price,
					ProductCost:  productData.Cost,
					Status:       1,
					Remark:       Remark,
					CreationTime: sysHelper.TimeStamp(),
				}
				_, e := orderDal.Add(db, data, "")
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

func OrderList(Token string, Page int, PageSize int, Order int, Stext string, ProductID int64, ManagerID int64, Status int64) mod.ResultList {
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
		result.Page, result.PageSize, result.TotalPage, result.Data = orderDal.List(db, Page, PageSize, Order, Stext, ProductID, ManagerID, Status, "")
	}
	return result
}

func OrderAll(Token string, Order int, Stext string, ProductID int64, ManagerID int64, Status int64) mod.Result {
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
		result.Data = orderDal.All(db, Order, Stext, ProductID, ManagerID, Status, "")
	}
	return result
}

func OrderData(Token string, ID int64) mod.Result {
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
		result.Data = orderDal.Data(db, ID, "")
	}
	return result
}

func OrderDel(Token, ID string) mod.Result {
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
		checkData := orderDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.TheOrderDataDoesNotExist
		} else {
			e := orderDal.Del(db, ID, "")
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
