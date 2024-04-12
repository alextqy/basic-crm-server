package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	Test := httpHelper.Post(r, "test")
	httpHelper.HttpWrite(w, bll.Test(Test))
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

func AdminNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Account := httpHelper.Post(r, "Account")
	Password := httpHelper.Post(r, "Password")
	Name := httpHelper.Post(r, "Name")
	Remark := httpHelper.Post(r, "Remark")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AdminNew(Token, Account, Password, Name, Remark, ID))
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

func AdminAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Level := httpHelper.PostInt64(r, "Level")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.AdminAll(Token, Order, Stext, Level, Status))
}

func AdminData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AdminData(Token, ID))
}

func AdminDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.AdminDel(Token, ID))
}
