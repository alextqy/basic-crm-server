package dal

import (
	mtd "basic-crm-server/MTD"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var adminTable = "Admin"
var companyTable = "Company"
var customerTable = "Customer"
var managerTable = "Manager"
var managerGroupTable = "ManagerGroup"
var salesPlanTable = "SalesPlan"
var salesTargetTable = "SalesTarget"

// var mu sync.Mutex
var xOnce sync.Once
var xSession *xorm.Session

var fileHelper = mtd.FileHelper{}
var sysHelper = mtd.SysHelper{}

func initDB() (bool, *xorm.Session, *xorm.EngineGroup) {
	conf := fileHelper.CheckConf()
	conns := []string{
		"postgres://" + conf.DbUser + ":" + conf.DbPwd + "@" + conf.DbHost + ":" + conf.DbPort + "/BasicCrm?sslmode=disable",
	}
	engine, err := xorm.NewEngineGroup("postgres", conns)
	if err != nil {
		log.Panic(err.Error())
		engine.Close()
		return false, nil, nil
	} else {
		engine.SetMaxOpenConns(100) // 连接池中最大连接数
		engine.SetMaxIdleConns(5)   // 连接池中最大空闲连接数
		engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
		engine.Ping()
		engine.ShowSQL(conf.DbDebug)

		session := engine.NewSession()
		defer session.Close()
		return true, session, engine
	}
}

// func ConnDB() *xorm.Session {
// 	if xSession == nil {
// 		mu.Lock()
// 		defer mu.Unlock()
// 		if xSession == nil {
// 			_, xSession, _ = initDB()
// 		}
// 	}
// 	return xSessions
// }

func ConnDB() *xorm.Session {
	var b bool
	xOnce.Do(func() {
		b, xSession, _ = initDB()
		if !b {
			os.Exit(0)
		}
	})
	return xSession
}
