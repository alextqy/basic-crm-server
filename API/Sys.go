package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	Test := httpHelper.Post(r, "test")
	httpHelper.HttpWrite(w, bll.Test(Test))
}

func CheckTheLogs(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Date := httpHelper.Post(r, "Date")
	Type := httpHelper.Post(r, "Type")
	Account := httpHelper.Post(r, "Account")
	httpHelper.HttpWrite(w, bll.CheckTheLogs(Token, Date, Type, Account))
}
