package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func AdminSignIn(w http.ResponseWriter, r *http.Request) {
	Account := httpHelper.Post(r, "account")
	Password := httpHelper.Post(r, "password")
	httpHelper.HttpWrite(w, bll.AdminSignIn(Account, Password))
}

func AdminSignOut(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "token")
	httpHelper.HttpWrite(w, bll.AdminSignOut(Token))
}
