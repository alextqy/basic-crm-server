package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func SupplierNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Name := httpHelper.Post(r, "Name")
	Email := httpHelper.Post(r, "Email")
	Tel := httpHelper.Post(r, "Tel")
	Address := httpHelper.Post(r, "Address")
	SupplierInfo := httpHelper.Post(r, "SupplierInfo")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.SupplierNew(Token, Name, Email, Tel, Address, SupplierInfo, ID))
}

func SupplierList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	httpHelper.HttpWrite(w, bll.SupplierList(Token, Page, PageSize, Order, Stext))
}

func SupplierAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	httpHelper.HttpWrite(w, bll.SupplierAll(Token, Order, Stext))
}

func SupplierData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.SupplierData(Token, ID))
}

func SupplierDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.SupplierDel(Token, ID))
}
