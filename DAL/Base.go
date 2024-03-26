package dal

import (
	"log"

	"xorm.io/xorm"
)

func InitDB() (bool, string, *xorm.Session, *xorm.Engine) {
	engine, err := xorm.NewEngine("postgres", "postgres://postgres:123456@localhost:5432/BasicCrm?sslmode=disable")

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
