package bll

import (
	dal "basic-crm-server/DAL"
	mtd "basic-crm-server/MTD"
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

func AdminVerification() int {
	return 0
}

func ManagerVerification() int {
	return 0
}
