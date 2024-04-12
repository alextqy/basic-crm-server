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
	CreationTime := httpHelper.PostInt64(r, "CreationTime")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	httpHelper.HttpWrite(w, bll.CustomerNew(Token, Name, Birthday, Gender, Email, Tel, CustomerInfo, Priority, CreationTime, CompanyID))
}

func CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.CustomerUpdate(Token))
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
