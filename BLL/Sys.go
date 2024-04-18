package bll

import (
	mod "basic-crm-server/MOD"
	"fmt"
)

func Test(Test string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	result.State = true
	result.Message = sysHelper.TimeNowStr()
	result.Data = Test
	return result
}

func CheckTheLogs(Token, Date, Type, Account string) mod.Result {
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
	} else if Type == "" {
		result.Message = lang.TypeError
	} else {
		if Type != "admin" && Type != "manager" {
			result.Message = lang.NoData
		} else {
			logFile := "./Log/" + Date + "/" + Type + "/" + Account + ".log"
			fmt.Println(logFile)
			if !fileHelper.FileExist(logFile) {
				result.Message = lang.NoData
			} else {
				b, r := fileHelper.FileRead(logFile)
				if !b {
					result.Message = r
				} else {
					result.State = true
					result.Data = r
				}
			}
		}
	}
	return result
}

func CheckEnv(Token string) mod.Result {
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
		result.State = true
		result.Data = sysHelper.SysEnvs()
	}
	return result
}
