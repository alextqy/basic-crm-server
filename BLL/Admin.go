package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
)

func SignIn(Account, Password string) mod.Result {
	result := mod.Result{}
	db := dal.ConnDB()
	r, e := dal.AdminCheck(db, Account, Password, "")
	if e != nil {
		result.Message = e.Error()
	} else {
		result.Data = r
		result.State = true
	}
	return result
}
