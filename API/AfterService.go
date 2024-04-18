package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func AfterServiceNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Account := httpHelper.Post(r, "Account")
	Password := httpHelper.Post(r, "Password")
	Name := httpHelper.Post(r, "Name")
	Remark := httpHelper.Post(r, "Remark")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AfterServiceNew(Token, Account, Password, Name, Remark, ID))
}

func AfterServiceList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Level := httpHelper.PostInt64(r, "Level")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.AfterServiceList(Token, Page, PageSize, Order, Stext, Level, Status))
}

func AfterServiceAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Level := httpHelper.PostInt64(r, "Level")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.AfterServiceAll(Token, Order, Stext, Level, Status))
}

func AfterServiceData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AfterServiceData(Token, ID))
}

func AfterServiceDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.AfterServiceDel(Token, ID))
}

func AfterServiceStatus(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AfterServiceStatus(Token, ID))
}
