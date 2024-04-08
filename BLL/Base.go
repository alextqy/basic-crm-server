package bll

import (
	dal "basic-crm-server/DAL"
	mod "basic-crm-server/MOD"
	mtd "basic-crm-server/MTD"
	"strings"
)

var lang = mtd.SysLang()

var cacheHelper = mtd.CacheHelper{}
var fileHelper = mtd.FileHelper{}
var httpHelper = mtd.HttpHelper{}
var sysHelper = mtd.SysHelper{}
var tcpHelper = mtd.TcpHelper{}
var udpHelper = mtd.UdpHelper{}

var adminDal = dal.AdminDal{}
var companyDal = dal.CompanyDal{}
var customerDal = dal.CustomerDal{}
var managerDal = dal.ManagerDal{}
var managerGroupDal = dal.ManagerGroupDal{}
var salesPlanDal = dal.SalesPlanDal{}
var salesTargetDal = dal.SalesTargetDal{}

func PwdMD5(Password string) string {
	return sysHelper.MD5(sysHelper.EnBase64(Password))
}

func EnToken(Account string, Type int) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	tokenType := ""
	if Type == 1 {
		tokenType = "admin"
	} else if Type == 2 {
		tokenType = "manager"
	} else {
		result.Message = ""
		return result
	}

	k := fileHelper.CheckConf().EncKey
	if k == "" {
		result.Message = lang.The16bitKeyIsNotSet
		return result
	}
	t := sysHelper.Int64ToString(sysHelper.TimeStampMS())
	keyParam := sysHelper.RandStr(10) + " " + tokenType + " " + Account + " " + t
	b, e, r := sysHelper.AesEncrypterCBC(keyParam, k, k)
	result.State = b
	result.Message = e
	result.Data = sysHelper.EnBase64(r)
	return result
}

func DeToken(Token string) mod.Result {
	result := mod.Result{
		State:   false,
		Message: "",
		Code:    200,
		Data:    nil,
	}

	if Token == "" {
		return result
	}

	b, e, r := sysHelper.DeBase64(Token)
	if !b {
		result.Message = e
	} else {
		k := fileHelper.CheckConf().EncKey
		b, e, r := sysHelper.AesDecrypterCBC(r, k, k)
		if !b {
			result.Message = e
		} else {
			db := dal.ConnDB()
			t := strings.Split(r, " ")
			if t[1] == "admin" {
				admin, e := adminDal.Token(db, Token, "")
				if e != nil {
					result.Message = e.Error()
				} else if admin.ID == 0 {
					result.Message = lang.IncorrectToken
				} else {
					result.State = true
					result.Message = "admin"
					result.Data = admin
				}
			} else if t[1] == "manager" {
				manager, e := managerDal.Token(db, Token, "")
				if e != nil {
					result.Message = e.Error()
				} else if manager.ID == 0 {
					result.Message = lang.IncorrectToken
				} else {
					result.State = true
					result.Message = "manager"
					result.Data = manager
				}
			} else {
				result.Message = lang.IncorrectToken
			}
		}
	}
	return result
}
