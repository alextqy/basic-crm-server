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
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Gender := httpHelper.PostInt64(r, "Gender")
	Priority := httpHelper.PostInt64(r, "Priority")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	httpHelper.HttpWrite(w, bll.CustomerList(Token, Page, PageSize, Order, Stext, Gender, Priority, CompanyID, ManagerID))
}

func CustomerAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Gender := httpHelper.PostInt64(r, "Gender")
	Priority := httpHelper.PostInt64(r, "Priority")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	httpHelper.HttpWrite(w, bll.CustomerAll(Token, Order, Stext, Gender, Priority, CompanyID, ManagerID))
}

func CustomerData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerData(Token, ID))
}

func CustomerDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerDel(Token, ID))
}
