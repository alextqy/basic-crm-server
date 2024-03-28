package dal

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var AdminTable = "Admin"
var CompanyTable = "Company"
var CustomerTable = "Customer"
var ManagerTable = "Manager"
var ManagerGroupTable = "ManagerGroup"
var SalesPlanTable = "SalesPlan"
var SalesTargetTable = "SalesTarget"

func InitDB() (bool, string, *xorm.Session, *xorm.EngineGroup) {
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
