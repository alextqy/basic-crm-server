package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func SalesTargetNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	TargetName := httpHelper.Post(r, "TargetName")
	ExpirationDate := httpHelper.PostInt64(r, "ExpirationDate")
	CustomerID := httpHelper.PostInt64(r, "CustomerID")
	Remark := httpHelper.Post(r, "Remark")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.SalesTargetNew(Token, TargetName, ExpirationDate, CustomerID, Remark, ID))
}

func SalesTargetList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	CustomerID := httpHelper.PostInt64(r, "CustomerID")
	httpHelper.HttpWrite(w, bll.SalesTargetList(Token, Page, PageSize, Order, Stext, CustomerID))
}

func SalesTargetAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	CustomerID := httpHelper.PostInt64(r, "CustomerID")
	httpHelper.HttpWrite(w, bll.SalesTargetAll(Token, Order, Stext, CustomerID))
}

func SalesTargetData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.SalesTargetData(Token, ID))
}

func SalesTargetDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.SalesTargetDel(Token, ID))
}
