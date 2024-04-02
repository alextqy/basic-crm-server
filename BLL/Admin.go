package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
)

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
			if r.Password != sysHelper.MD5(sysHelper.EnBase64(Password)) {
				result.Message = lang.IncorrectPassword
			} else {
				r.Token = sysHelper.MD5(sysHelper.EnBase64(sysHelper.RandStr(10)))
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
		result.Message = Token
	} else {
		db := dal.ConnDB()
		r, e := adminDal.Token(db, Token, "")
		if e != nil {
			result.Message = e.Error()
		} else if r.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else {
			r.Token = ""
			_, e := adminDal.Update(db, r, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				result.State = true
				result.Message = ""
			}
		}
	}

	return result
}
