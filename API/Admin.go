package api

import (
	bll "basic-crm-server/BLL"
	lib "basic-crm-server/LIB"
	"net/http"
	"strings"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	Account := strings.TrimSpace(lib.Post(r, "account"))
	Password := strings.TrimSpace(lib.Post(r, "password"))
	lib.HttpWrite(w, bll.SignIn(Account, Password))
}
