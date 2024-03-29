package api

import (
	bll "basic-crm-server/BLL"
	mtd "basic-crm-server/MTD"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	Account := mtd.Post(r, "account")
	Password := mtd.Post(r, "password")
	mtd.HttpWrite(w, bll.SignIn(Account, Password))
}
