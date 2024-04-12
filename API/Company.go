package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func CompanyNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	CompanyName := httpHelper.Post(r, "CompanyName")
	Remark := httpHelper.Post(r, "Remark")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CompanyNew(Token, CompanyName, Remark, ID))
}

func CompanyList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	httpHelper.HttpWrite(w, bll.CompanyList(Token, Page, PageSize, Order, Stext))
}

func CompanyAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	httpHelper.HttpWrite(w, bll.CompanyAll(Token, Order, Stext))
}

func CompanyData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CompanyData(Token, ID))
}

func CompanyDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.CompanyDel(Token, ID))
}
