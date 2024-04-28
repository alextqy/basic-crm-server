package dal

import (
	mtd "basic-crm-server/MTD"
	"os"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var adminTable = "Admin"
var afterServiceTable = "AfterService"
var announcementTable = "Announcement"
var companyTable = "Company"
var customerTable = "Customer"
var distributorTable = "Distributor"
var managerTable = "Manager"
var managerGroupTable = "ManagerGroup"
var orderTable = "Order"
var productTable = "Product"
var salesPlanTable = "SalesPlan"
var salesTargetTable = "SalesTarget"

var cacheHelper = mtd.CacheHelper{}
var fileHelper = mtd.FileHelper{}
var httpHelper = mtd.HttpHelper{}
var sysHelper = mtd.SysHelper{}
var tcpHelper = mtd.TcpHelper{}
var udpHelper = mtd.UdpHelper{}

var gOnce sync.Once
var gDB *gorm.DB

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("Dao.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	return db, err
}

func ConnDB() *gorm.DB {
	var e error
	gOnce.Do(func() {
		gDB, e = initDB()
		if e != nil {
			os.Exit(0)
		}
	})
	return gDB
}
