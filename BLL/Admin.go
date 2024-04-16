package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func Test(Test string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	result.State = true
	result.Message = PwdMD5(Test)
	return result
}

func AdminSignIn(Account, Password string) mod.Result {
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
		checkData := adminDal.Check(db, Account, "")
		if checkData.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else if checkData.Status != 1 {
			result.Message = lang.AccountDisabled
		} else {
			if checkData.Password != PwdMD5(Password) {
				result.Message = lang.IncorrectPassword
			} else {
				Token := EnToken(Account, 1)
				if !Token.State {
					result.Message = Token.Message
				} else {
					result.Data = Token.Data.(string)
					checkData.Token = Token.Data.(string)
					e := adminDal.Update(db, checkData, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						go fileHelper.WriteLog(checkData.Account, checkData.Account+" login", "admin")
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

func AdminSignOut(Token string) mod.Result {
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
			if r.Message == "admin" {
				userData := r.Data.(mod.Admin)
				if userData.ID == 0 {
					result.Message = lang.TheAccountDoesNotExist
				} else {
					userData.Token = ""
					e := adminDal.Update(db, userData, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						result.State = true
						go fileHelper.WriteLog(userData.Account, userData.Account+" logout", "admin")
					}
				}
			} else {
				result.Message = lang.TheAccountDoesNotExist
			}
		}
	}
	return result
}

func AdminNew(Token, Account, Password, Name, Remark string, ID int64) mod.Result {
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
	} else if Account == "" {
		result.Message = lang.IncorrectAccount
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else if len(Account) < 6 {
		result.Message = lang.TheAccountIsTooShort
	} else {
		db := dal.ConnDB()

		if ID > 0 {
			checkData := adminDal.Data(db, ID, "")
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
				e := adminDal.Update(db, checkData, "")
				if e != nil {
					result.Message = e.Error()
				} else {
					result.State = true
					jData, _ := json.Marshal(checkData)
					go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), "admin")
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
				data := mod.Admin{
					Account:      Account,
					Password:     PwdMD5(Password),
					Name:         Name,
					Level:        1,
					Status:       1,
					Remark:       Remark,
					CreationTime: sysHelper.TimeStamp(),
				}
				checkData := adminDal.Check(db, Account, "")
				if checkData.ID > 0 {
					result.Message = lang.TheAccountAlreadyExists
				} else {
					_, e := adminDal.Add(db, data, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						jData, _ := json.Marshal(data)
						go fileHelper.WriteLog(CheckAccount(t), "Add data: "+string(jData), "admin")
						result.State = true
					}
				}
			}
		}
	}
	return result
}

func AdminList(Token string, Page, PageSize, Order int, Stext string, Level, Status int64) mod.ResultList {
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
	} else if t.Message != "admin" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = adminDal.List(db, Page, PageSize, Order, Stext, Level, Status, "")
		for i := 0; i < len(result.Data.([]mod.Admin)); i++ {
			result.Data.([]mod.Admin)[i].Password = ""
		}
	}
	return result
}

func AdminAll(Token string, Order int, Stext string, Level, Status int64) mod.Result {
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
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = adminDal.All(db, Order, Stext, Level, Status, "")
		for i := 0; i < len(result.Data.([]mod.Admin)); i++ {
			result.Data.([]mod.Admin)[i].Password = ""
		}
	}
	return result
}

func AdminData(Token string, ID int64) mod.Result {
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
	} else {
		db := dal.ConnDB()
		result.State = true
		Data := adminDal.Data(db, ID, "")
		Data.Password = ""
		result.Data = Data
	}
	return result
}

func AdminDel(Token, ID string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	if ID == "1" {
		return result
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := adminDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else {
			e := adminDal.Del(db, ID, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Remove data: "+string(jData), "admin")
				result.State = true
			}
		}
	}
	return result
}

func AdminStatus(Token string, ID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	if ID == 1 {
		return result
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "admin" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		checkData := adminDal.Data(db, ID, "")
		if checkData.ID == 0 {
			result.Message = lang.NoData
		} else {
			if checkData.Status == 1 {
				checkData.Status = 2
			} else {
				checkData.Status = 1
			}

			e := adminDal.Update(db, checkData, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				result.State = true
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), "admin")
			}
		}
	}
	return result
}
