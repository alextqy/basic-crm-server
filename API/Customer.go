package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func CustomerNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Name := httpHelper.Post(r, "Name")
	Birthday := httpHelper.PostInt64(r, "Birthday")
	Gender := httpHelper.PostInt64(r, "Gender")
	Email := httpHelper.Post(r, "Email")
	Tel := httpHelper.Post(r, "Tel")
	CustomerInfo := httpHelper.Post(r, "CustomerInfo")
	Priority := httpHelper.PostInt64(r, "Priority")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerNew(Token, Name, Birthday, Gender, Email, Tel, CustomerInfo, Priority, CompanyID, ID))
}

func CustomerList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.CustomerList(Token))
}

func CustomerAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.CustomerAll(Token))
}

func CustomerDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.CustomerDel(Token))
}
