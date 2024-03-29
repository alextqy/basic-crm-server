package dal

import (
	mtd "basic-crm-server/MTD"
	"fmt"
	"log"
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

func initDB() (bool, string, *xorm.Session, *xorm.EngineGroup) {
	conf := mtd.CheckConf()
	fmt.Println(conf)
	conns := []string{
		"postgres://postgres:123456@localhost:5432/BasicCrm?sslmode=disable",
	}
	engine, err := xorm.NewEngineGroup("postgres", conns)
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	engine.Ping()
	engine.ShowSQL()
	if err != nil {
		engine.Close()
		log.Fatal(err.Error())
		return false, err.Error(), nil, nil
	} else {
		session := engine.NewSession()
		defer session.Close()
		return true, "", session, engine
	}
}

// func ConnDB() *xorm.Session {
// 	if xSession == nil {
// 		mu.Lock()
// 		defer mu.Unlock()
// 		if xSession == nil {
// 			_, _, xSession, _ = initDB()
// 		}
// 	}
// 	return xSessions
// }

func ConnDB() *xorm.Session {
	xOnce.Do(func() {
		_, _, xSession, _ = initDB()
	})
	return xSession
}
