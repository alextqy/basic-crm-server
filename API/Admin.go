package api

import (
	bll "basic-crm-server/BLL"
	mtd "basic-crm-server/MTD"
	"net/http"
	"strings"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	Account := strings.TrimSpace(mtd.Post(r, "account"))
	Password := strings.TrimSpace(mtd.Post(r, "password"))
	mtd.HttpWrite(w, bll.SignIn(Account, Password))
}
