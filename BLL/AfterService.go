package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func AfterServiceNew(Token, Account, Password, Name, Remark string, ID int64) mod.Result {
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
	} else if Account == "" {
		result.Message = lang.IncorrectAccount
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else if len(Account) < 6 {
		result.Message = lang.TheAccountIsTooShort
	} else {
		db := dal.ConnDB()

		if ID > 0 {
			checkData := afterServiceDal.Data(db, ID, "")
			if checkData.ID == 0 {
				result.Message = lang.TheAccountDoesNotExist
			} else {
				if Password != "" && len(Password) < 6 {
					result.Message = lang.ThePasswordIsTooShort
					return result
				}

				newPwd := ""
				if Password == "" {
					newPwd = checkData.Password
				} else {
					newPwd = PwdMD5(Password)
				}
				checkData.Password = newPwd
				checkData.Name = Name
				checkData.Remark = Remark
				e := afterServiceDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
					result.State = true
				}
			}
		} else {
			if Account == "" {
				result.Message = lang.IncorrectAccount
			} else if len(Account) < 6 {
				result.Message = lang.TheAccountIsTooShort
			} else if Password == "" {
				result.Message = lang.IncorrectPassword
			} else if len(Password) < 6 {
				result.Message = lang.ThePasswordIsTooShort
			} else {
				data := mod.AfterService{
					Account:      Account,
					Password:     PwdMD5(Password),
					Name:         Name,
					Level:        1,
					Status:       1,
					Remark:       Remark,
					CreationTime: sysHelper.TimeStamp(),
				}
				checkData := afterServiceDal.Check(db, Account, "")
				if checkData.ID > 0 {
					result.Message = lang.TheAccountAlreadyExists
				} else {
					_, e := afterServiceDal.Add(db, data, "")
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
	}
	return result
}

func AfterServiceList(Token string, Page, PageSize, Order int, Stext string, Level, Status int64) mod.ResultList {
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
		result.Page, result.PageSize, result.TotalPage, result.Data = afterServiceDal.List(db, Page, PageSize, Order, Stext, Level, Status, "")
		for i := 0; i < len(result.Data.([]mod.AfterService)); i++ {
			result.Data.([]mod.AfterService)[i].Password = ""
		}
	}
	return result
}

func AfterServiceAll(Token string, Order int, Stext string, Level, Status int64) mod.Result {
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
		result.Data = afterServiceDal.All(db, Order, Stext, Level, Status, "")
		for i := 0; i < len(result.Data.([]mod.AfterService)); i++ {
			result.Data.([]mod.AfterService)[i].Password = ""
		}
	}
	return result
}

func AfterServiceData(Token string, ID int64) mod.Result {
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
		Data := afterServiceDal.Data(db, ID, "")
		Data.Password = ""
		result.Data = Data
	}
	return result
}

func AfterServiceDel(Token, ID string) mod.Result {
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
		checkData := afterServiceDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else {
			e := afterServiceDal.Del(db, ID, "")
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

func AfterServiceStatus(Token string, ID int64) mod.Result {
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
		checkData := afterServiceDal.Data(db, ID, "")
		if checkData.ID == 0 {
			result.Message = lang.NoData
		} else {
			if checkData.Status == 1 {
				checkData.Status = 2
			} else {
				checkData.Status = 1
			}

			e := afterServiceDal.Update(db, checkData, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				result.State = true
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
			}
		}
	}
	return result
}
