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
						go fileHelper.WriteLog(checkData.Account, checkData.Account+" login")
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
						go fileHelper.WriteLog(userData.Account, userData.Account+" logout")
					}
				}
			} else {
				result.Message = lang.TheAccountDoesNotExist
			}
		}
	}
	return result
}

func AdminUpdate(Token, Password, Name, Remark string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Data.(mod.Admin).ID == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else {
		userData := t.Data.(mod.Admin)
		db := dal.ConnDB()
		newPwd := ""
		if Password == "" {
			newPwd = userData.Password
		} else {
			newPwd = PwdMD5(Password)
		}
		userData.Password = newPwd
		userData.Name = Name
		userData.Remark = Remark
		e := adminDal.Update(db, userData, "")
		if e != nil {
			result.Message = e.Error()
		} else {
			result.State = true
			jData, _ := json.Marshal(userData)
			go fileHelper.WriteLog(userData.Account, "The data has been updated. Old data: "+string(jData))
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
	} else if t.Data.(mod.Admin).ID == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Page, result.PageSize, result.TotalPage, result.Data = adminDal.List(db, Page, PageSize, Order, Stext, Level, Status, "")
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
	} else if t.Data.(mod.Admin).ID == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		db := dal.ConnDB()
		result.State = true
		result.Data = adminDal.All(db, Order, Stext, Level, Status, "")
	}
	return result
}

func AdminNew(Token string, Account, Password, Name, Remark string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Data.(mod.Admin).ID == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if Account == "" {
		result.Message = lang.IncorrectAccount
	} else if Password == "" {
		result.Message = lang.IncorrectPassword
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else if len(Account) < 6 {
		result.Message = lang.TheAccountIsTooShort
	} else if len(Password) < 6 {
		result.Message = lang.ThePasswordIsTooShort
	} else {
		userData := t.Data.(mod.Admin)

		data := mod.Admin{
			Account:      Account,
			Password:     PwdMD5(Password),
			Name:         Name,
			Level:        1,
			Status:       1,
			Remark:       Remark,
			CreationTime: sysHelper.TimeStamp(),
		}
		db := dal.ConnDB()
		checkData := adminDal.Check(db, Account, "")
		if checkData.ID > 0 {
			result.Message = lang.TheAccountAlreadyExists
		} else {
			_, e := adminDal.Add(db, data, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(data)
				go fileHelper.WriteLog(userData.Account, "Add data: "+string(jData))
				result.State = true
			}
		}
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

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Data.(mod.Admin).ID == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else {
		userData := t.Data.(mod.Admin)

		db := dal.ConnDB()
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := adminDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.NoData
		} else {
			e := adminDal.Del(db, ID, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(userData.Account, "Remove data: "+string(jData))
				result.State = true
			}
		}
	}
	return result
}
