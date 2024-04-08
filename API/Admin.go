package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	Test := httpHelper.Post(r, "test")
	bll.Test(Test)
}

func AdminSignIn(w http.ResponseWriter, r *http.Request) {
	Account := httpHelper.Post(r, "account")
	Password := httpHelper.Post(r, "password")
	httpHelper.HttpWrite(w, bll.AdminSignIn(Account, Password))
}

func AdminSignOut(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "token")
	httpHelper.HttpWrite(w, bll.AdminSignOut(Token))
}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "token")
	Password := httpHelper.Post(r, "password")
	Name := httpHelper.Post(r, "name")
	Remark := httpHelper.Post(r, "remark")
	httpHelper.HttpWrite(w, bll.AdminUpdate(Token, Password, Name, Remark))
}
