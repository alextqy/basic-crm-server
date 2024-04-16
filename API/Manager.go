package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func ManagerNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Account := httpHelper.Post(r, "Account")
	Password := httpHelper.Post(r, "Password")
	Name := httpHelper.Post(r, "Name")
	Remark := httpHelper.Post(r, "Remark")
	GroupID := httpHelper.PostInt64(r, "GroupID")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.ManagerNew(Token, Account, Password, Name, Remark, GroupID, ID))
}

func ManagerList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Level := httpHelper.PostInt64(r, "Level")
	Status := httpHelper.PostInt64(r, "Status")
	GroupID := httpHelper.PostInt64(r, "GroupID")
	httpHelper.HttpWrite(w, bll.ManagerList(Token, Page, PageSize, Order, Stext, Level, Status, GroupID))
}

func ManagerAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Level := httpHelper.PostInt64(r, "Level")
	Status := httpHelper.PostInt64(r, "Status")
	GroupID := httpHelper.PostInt64(r, "GroupID")
	httpHelper.HttpWrite(w, bll.ManagerAll(Token, Order, Stext, Level, Status, GroupID))
}

func ManagerData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.ManagerData(Token, ID))
}

func ManagerDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.ManagerDel(Token, ID))
}

func ManagerStatus(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.ManagerStatus(Token, ID))
}

func ManagerSignIn(w http.ResponseWriter, r *http.Request) {
	Account := httpHelper.Post(r, "Account")
	Password := httpHelper.Post(r, "Password")
	httpHelper.HttpWrite(w, bll.ManagerSignIn(Account, Password))
}

func ManagerSignOut(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.ManagerSignOut(Token))
}

func ManagerUpdate(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Password := httpHelper.Post(r, "Password")
	Name := httpHelper.Post(r, "Name")
	Remark := httpHelper.Post(r, "Remark")
	GroupID := httpHelper.PostInt64(r, "GroupID")
	httpHelper.HttpWrite(w, bll.ManagerUpdate(Token, Password, Name, Remark, GroupID))
}
