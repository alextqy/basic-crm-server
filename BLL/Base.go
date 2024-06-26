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
var afterServiceDal = dal.AfterServiceDal{}
var announcementDal = dal.AnnouncementDal{}
var companyDal = dal.CompanyDal{}
var customerDal = dal.CustomerDal{}
var customerQADal = dal.CustomerQADal{}
var distributorDal = dal.DistributorDal{}
var managerDal = dal.ManagerDal{}
var managerGroupDal = dal.ManagerGroupDal{}
var orderDal = dal.OrderDal{}
var productDal = dal.ProductDal{}
var salesPlanDal = dal.SalesPlanDal{}
var salesTargetDal = dal.SalesTargetDal{}
var supplierDal = dal.SupplierDal{}

func CheckPerm(t mod.Result) int {
	var p int
	if t.Message == "afterService" {
		p = 3
	} else if t.Message == "manager" {
		p = 2
	} else if t.Message == "admin" {
		p = 1
	} else {
		p = 99
	}
	return p
}

func CheckID(t mod.Result) int64 {
	var ID int64
	if t.Message == "admin" {
		ID = t.Data.(mod.AdminMod).ID
	} else if t.Message == "manager" {
		ID = t.Data.(mod.ManagerMod).ID
	} else if t.Message == "afterService" {
		ID = t.Data.(mod.AfterServiceMod).ID
	} else {
		ID = 0
	}
	return ID
}

func CheckAccount(t mod.Result) string {
	account := ""
	if t.Message == "admin" {
		account = t.Data.(mod.AdminMod).Account
	} else if t.Message == "manager" {
		account = t.Data.(mod.ManagerMod).Account
	} else if t.Message == "afterService" {
		account = t.Data.(mod.AfterServiceMod).Account
	} else {
		account = ""
	}
	return account
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
	} else if Type == 3 {
		tokenType = "afterService"
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
				admin := adminDal.Token(db, Token, "")
				if admin.ID == 0 {
					result.Message = lang.IncorrectToken
				} else {
					result.State = true
					result.Message = "admin"
					result.Data = admin
				}
			} else if t[1] == "manager" {
				manager := managerDal.Token(db, Token, "")
				if manager.ID == 0 {
					result.Message = lang.IncorrectToken
				} else {
					result.State = true
					result.Message = "manager"
					result.Data = manager
				}
			} else if t[1] == "afterService" {
				manager := afterServiceDal.Token(db, Token, "")
				if manager.ID == 0 {
					result.Message = lang.IncorrectToken
				} else {
					result.State = true
					result.Message = "afterService"
					result.Data = manager
				}
			} else {
				result.Message = lang.IncorrectToken
			}
		}
	}
	return result
}

func PwdMD5(Password string) string {
	return sysHelper.MD5(sysHelper.EnBase64(Password))
}
