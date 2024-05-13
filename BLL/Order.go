package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func OrderNew(Token, OrderNo string, ProductID, ManagerID, CustomerID, DistributorID int64, OrderPrice float32, Remark string, OrderType, Payment, Review, ID int64) mod.Result {
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
	} else if OrderPrice == 0 {
		result.Message = lang.IncorrectOrderPrice
	} else if OrderType == 0 {
		result.Message = lang.TypeError
	} else if OrderType == 1 && CustomerID == 0 {
		result.Message = lang.TypeError
	} else if OrderType == 2 && DistributorID == 0 {
		result.Message = lang.TypeError
	} else {
		if OrderType == 1 {
			DistributorID = 0
		}
		if OrderType == 2 {
			CustomerID = 0
		}

		db := dal.ConnDB()

		productData := productDal.Data(db, ProductID, "")
		if productData.ID == 0 {
			result.Message = lang.TheProductDataDoesNotExist
			return result
		}

		if CheckPerm(t) == 2 {
			ManagerID = CheckID(t)
		} else if ManagerID > 0 {
			checkData := managerDal.Data(db, ManagerID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheSalesManagerDoesNotExist
				return result
			}
		} else {
			ManagerID = 0
		}

		if CustomerID > 0 {
			checkData := customerDal.Data(db, CustomerID, "")
			if checkData.ID == 0 {
				result.Message = lang.CustomerDataDoesNotExist
				return result
			}
		}

		if DistributorID > 0 {
			checkData := distributorDal.Data(db, DistributorID, "")
			if checkData.ID == 0 {
				result.Message = lang.DistributorDataDoesNotExist
				return result
			}
		}

		if ID > 0 {
			checkData := orderDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheOrderDataDoesNotExist
			} else {
				checkData.ManagerID = ManagerID
				checkData.CustomerID = CustomerID
				checkData.DistributorID = DistributorID
				checkData.OrderPrice = OrderPrice
				checkData.Remark = Remark
				checkData.Payment = Payment
				checkData.Review = Review
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
				data := mod.OrderMod{
					OrderNo:       OrderNo,
					ProductID:     ProductID,
					ManagerID:     ManagerID,
					CustomerID:    CustomerID,
					DistributorID: DistributorID,
					OrderPrice:    OrderPrice,
					ProductPrice:  productData.Price,
					ProductCost:   productData.Cost,
					Status:        1,
					Remark:        Remark,
					CreationTime:  sysHelper.TimeStamp(),
					OrderType:     OrderType,
					Payment:       Payment,
					Review:        Review,
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

func OrderList(Token string, Page, PageSize, Order int, Stext string, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review int64) mod.ResultList {
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
		result.Page, result.PageSize, result.TotalPage, result.Data = orderDal.List(db, Page, PageSize, Order, Stext, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review, "")
	}
	return result
}

func OrderAll(Token string, Order int, Stext string, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review int64) mod.Result {
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
		result.Data = orderDal.All(db, Order, Stext, ProductID, ManagerID, CustomerID, DistributorID, Status, OrderType, Payment, Review, "")
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
		data := orderDal.Data(db, ID, "")
		if CheckPerm(t) == 2 && data.ManagerID != CheckID(t) {
			result.Data = mod.OrderMod{}
		} else {
			result.Data = data
		}
		result.State = true
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
			if CheckPerm(t) == 2 {
				if checkData.ManagerID != CheckID(t) {
					result.Message = lang.PermissionDenied
					return result
				}
			}
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
