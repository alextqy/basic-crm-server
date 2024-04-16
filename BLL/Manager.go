package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	"encoding/json"
)

func ManagerNew(Token, Account, Password, Name, Remark string, GroupID, ID int64) mod.Result {
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
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else {
		db := dal.ConnDB()

		if GroupID > 0 {
			checkData := managerGroupDal.Data(db, GroupID, "")
			if checkData.ID == 0 {
				result.Message = lang.IncorrectGroup
				return result
			}
		}

		if ID > 0 {
			checkData := managerDal.Data(db, ID, "")
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
				checkData.GroupID = GroupID
				e := managerDal.Update(db, checkData, "")
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
				data := mod.Manager{
					Account:      Account,
					Password:     PwdMD5(Password),
					Name:         Name,
					Level:        1,
					Status:       1,
					Remark:       Remark,
					CreationTime: sysHelper.TimeStamp(),
					GroupID:      GroupID,
				}
				checkData := managerDal.Check(db, Account, "")
				if checkData.ID > 0 {
					result.Message = lang.TheAccountAlreadyExists
				} else {
					_, e := managerDal.Add(db, data, "")
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

func ManagerList(Token string, Page, PageSize, Order int, Stext string, Level, Status, GroupID int64) mod.ResultList {
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
		result.Page, result.PageSize, result.TotalPage, result.Data = managerDal.List(db, Page, PageSize, Order, Stext, Level, Status, GroupID, "")
		for i := 0; i < len(result.Data.([]mod.Manager)); i++ {
			result.Data.([]mod.Manager)[i].Password = ""
		}
	}
	return result
}

func ManagerAll(Token string, Order int, Stext string, Level, Status, GroupID int64) mod.Result {
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
		result.Data = managerDal.All(db, Order, Stext, Level, Status, GroupID, "")
		for i := 0; i < len(result.Data.([]mod.Manager)); i++ {
			result.Data.([]mod.Manager)[i].Password = ""
		}
	}
	return result
}

func ManagerData(Token string, ID int64) mod.Result {
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
		data := managerDal.Data(db, ID, "")
		data.Password = ""
		result.Data = data
	}
	return result
}

func ManagerDel(Token, ID string) mod.Result {
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
		_, _, ID64 := sysHelper.StringToInt64(ID)
		checkData := managerDal.Data(db, ID64, "")
		if checkData.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else {
			e := managerDal.Del(db, ID, "")
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

func ManagerStatus(Token string, ID int64) mod.Result {
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
		checkData := managerDal.Data(db, ID, "")
		if checkData.ID == 0 {
			result.Message = lang.NoData
		} else {
			if checkData.Status == 1 {
				checkData.Status = 2
			} else {
				checkData.Status = 1
			}

			e := managerDal.Update(db, checkData, "")
			if e != nil {
				result.Message = e.Error()
			} else {
				jData, _ := json.Marshal(checkData)
				go fileHelper.WriteLog(CheckAccount(t), "Modify the data: "+string(jData), t.Message)
				result.State = true
			}
		}
	}
	return result
}

func ManagerSignIn(Account, Password string) mod.Result {
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
		checkData := managerDal.Check(db, Account, "")
		if checkData.ID == 0 {
			result.Message = lang.TheAccountDoesNotExist
		} else if checkData.Status != 1 {
			result.Message = lang.AccountDisabled
		} else {
			if checkData.Password != PwdMD5(Password) {
				result.Message = lang.IncorrectPassword
			} else {
				t := EnToken(Account, 2)
				if !t.State {
					result.Message = t.Message
				} else {
					result.Data = t.Data.(string)
					checkData.Token = t.Data.(string)
					e := managerDal.Update(db, checkData, "")
					if e != nil {
						result.Message = e.Error()
					} else {
						go fileHelper.WriteLog(checkData.Account, checkData.Account+" login", "manager")
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

func ManagerSignOut(Token string) mod.Result {
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
			if r.Message == "manager" {
				userData := r.Data.(mod.Manager)
				if userData.ID == 0 {
					result.Message = lang.TheAccountDoesNotExist
				} else {
					userData.Token = ""
					e := managerDal.Update(db, userData, "")
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

func ManagerUpdate(Token, Password, Name, Remark string, GroupID int64) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	t := DeToken(Token)
	if !t.State {
		result.Message = t.Message
	} else if t.Message != "manager" {
		result.Message = lang.PermissionDenied
	} else if CheckID(t) == 0 {
		result.Message = lang.TheAccountDoesNotExist
	} else if Name == "" {
		result.Message = lang.IncorrectName
	} else {
		db := dal.ConnDB()

		if GroupID > 0 {
			checkData := managerGroupDal.Data(db, GroupID, "")
			if checkData.ID == 0 {
				result.Message = lang.IncorrectGroup
				return result
			}
		}

		if Password != "" && len(Password) < 6 {
			result.Message = lang.ThePasswordIsTooShort
			return result
		}

		data := t.Data.(mod.Manager)
		newPwd := ""
		if Password == "" {
			newPwd = data.Password
		} else {
			newPwd = PwdMD5(Password)
		}
		data.Password = newPwd
		data.Name = Name
		data.Remark = Remark
		data.GroupID = GroupID
		e := managerDal.Update(db, data, "")
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
