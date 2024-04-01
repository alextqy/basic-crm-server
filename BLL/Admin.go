package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
)

func SignIn(Account, Password string) mod.Result {
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
					result.Data = r
					result.State = true
				}
			}
		}
	}
	return result
}
