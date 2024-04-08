package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func Test(Test string) {}

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
		r, e := adminDal.Check(db, Account, "")
		if e != nil {
			result.Message = e.Error()
		} else if r.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else {
			if r.Password != PwdMD5(Password) {
				result.Message = lang.IncorrectPassword
			} else {
				Token := EnToken(Account, 1)
				if !Token.State {
					result.Message = Token.Message
				} else {
					result.Data = Token.Data.(string)
					r.Token = Token.Data.(string)
					_, e := adminDal.Update(db, r, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						go fileHelper.WriteLog(r.Account, r.Account+" login")
						r.Password = ""
						result.State = true
						result.Data = r
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
					_, e := adminDal.Update(db, userData, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						go fileHelper.WriteLog(userData.Account, userData.Account+" logout")
						result.State = true
						result.Message = ""
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
		_, e := adminDal.Update(db, userData, "")
		if e != nil {
			result.Message = e.Error()
			return result
		} else {
			jData, _ := json.Marshal(userData)
			go fileHelper.WriteLog(userData.Account, "The data has been updated. Old data: "+string(jData))
			result.State = true
			result.Message = ""
		}

	}
	return result
}
