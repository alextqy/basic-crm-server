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
	} else if Name == "" {
		result.Message = lang.IncorrectName
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
			} else if !sysHelper.RegEnNum(Account) {
				result.Message = lang.IncorrectAccountFormat
			} else if Password == "" {
				result.Message = lang.IncorrectPassword
			} else if len(Password) < 6 {
				result.Message = lang.ThePasswordIsTooShort
			} else {
				data := mod.AfterServiceMod{
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
		for i := 0; i < len(result.Data.([]mod.AfterServiceMod)); i++ {
			result.Data.([]mod.AfterServiceMod)[i].Password = ""
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
		for i := 0; i < len(result.Data.([]mod.AfterServiceMod)); i++ {
			result.Data.([]mod.AfterServiceMod)[i].Password = ""
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
			result.Message = lang.AfterSalesPersonnelDoNot
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

func AfterServiceSignIn(Account, Password string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	if Account == "" {
		result.Message = lang.IncorrectAccount
	} else if Password == "" {
		result.Message = lang.IncorrectPassword
	} else {
		db := dal.ConnDB()
		checkData := afterServiceDal.Check(db, Account, "")
		if checkData.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else if checkData.Status != 1 {
			result.Message = lang.AccountDisabled
		} else {
			if checkData.Password != PwdMD5(Password) {
				result.Message = lang.IncorrectPassword
			} else {
				t := EnToken(Account, 3)
				if !t.State {
					result.Message = t.Message
				} else {
					result.Data = t.Data.(string)
					checkData.Token = t.Data.(string)
					e := afterServiceDal.Update(db, checkData, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						go fileHelper.WriteLog(checkData.Account, checkData.Account+" login", "afterService")
						checkData.Password = ""
						result.State = true
						result.Data = checkData
					}
				}
			}
		}
	}
	return result
}

func AfterServiceSignOut(Token string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	if Token == "" {
		result.Message = lang.IncorrectToken
	} else {
		r := DeToken(Token)
		if !r.State {
			result.Message = r.Message
		} else {
			db := dal.ConnDB()
			if r.Message == "afterService" {
				userData := r.Data.(mod.AfterServiceMod)
				if userData.ID == 0 {
					result.Message = lang.TheAccountDoesNotExist
				} else {
					userData.Token = ""
					e := afterServiceDal.Update(db, userData, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						go fileHelper.WriteLog(userData.Account, userData.Account+" logout", r.Message)
						result.State = true
					}
				}
			} else {
				result.Message = lang.TheAccountDoesNotExist
			}
		}
	}
	return result
}

func AfterServiceUpdate(Token, Password, Name, Remark string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if CheckPerm(t) != 3 {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else {
		db := dal.ConnDB()

		if Password != "" && len(Password) < 6 {
			result.Message = lang.ThePasswordIsTooShort
			return result
		}

		data := t.Data.(mod.AfterServiceMod)
		newPwd := ""
		if Password == "" {
			newPwd = data.Password
		} else {
			newPwd = PwdMD5(Password)
		}
		data.Password = newPwd
		data.Name = Name
		data.Remark = Remark
		e := afterServiceDal.Update(db, data, "")
		if e != nil {
			result.Message = e.Error()
		} else {
			jData, _ := json.Marshal(data)
			go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
			result.State = true
		}
	}
	return result
}
