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
	Account := httpHelper.Post(r, "Account")
	Password := httpHelper.Post(r, "Password")
	httpHelper.HttpWrite(w, bll.AdminSignIn(Account, Password))
}

func AdminSignOut(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.AdminSignOut(Token))
}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Password := httpHelper.Post(r, "Password")
	Name := httpHelper.Post(r, "Name")
	Remark := httpHelper.Post(r, "Remark")
	httpHelper.HttpWrite(w, bll.AdminUpdate(Token, Password, Name, Remark))
}

func AdminList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Level := httpHelper.PostInt64(r, "Level")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.AdminList(Token, Page, PageSize, Order, Stext, Level, Status))
}
