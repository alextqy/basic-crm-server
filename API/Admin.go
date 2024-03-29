package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	Account := httpHelper.Post(r, "account")
	Password := httpHelper.Post(r, "password")
	httpHelper.HttpWrite(w, bll.SignIn(Account, Password))
}
