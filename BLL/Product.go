package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func ProductNew(Token, ProductName string, Price, Cost float32, Remark string, ID int64) mod.Result {
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
	} else if ProductName == "" {
		result.Message = lang.IncorrectName
	} else {
		db := dal.ConnDB()

		if ID > 0 {
			checkData := productDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheProductDataDoesNotExist
			} else {
				checkData.ProductName = ProductName
				checkData.Price = Price
				checkData.Cost = Cost
				checkData.Remark = Remark
				e := productDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			data := mod.ProductMod{
				ProductName:  ProductName,
				Price:        Price,
				Cost:         Cost,
				Status:       1,
				Remark:       Remark,
				CreationTime: sysHelper.TimeStamp(),
			}
			_, e := productDal.Add(db, data, "")
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

func ProductList(Token string, Page, PageSize, Order int, Stext string, Status int64) mod.ResultList {
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
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = productDal.List(db, Page, PageSize, Order, Stext, Status, "")
	}
	return result
}

func ProductAll(Token string, Order int, Stext string, Status int64) mod.Result {
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
		result.State = true
		result.Data = productDal.All(db, Order, Stext, Status, "")
	}
	return result
}

func ProductData(Token string, ID int64) mod.Result {
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
		result.State = true
		result.Data = productDal.Data(db, ID, "")
	}
	return result
}

func ProductDel(Token, ID string) mod.Result {
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
		checkData := productDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.TheProductDataDoesNotExist
		} else {
			e := productDal.Del(db, ID, "")
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
