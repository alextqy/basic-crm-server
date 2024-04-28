package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func DistributorNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Name := httpHelper.Post(r, "Name")
	Email := httpHelper.Post(r, "Email")
	Tel := httpHelper.Post(r, "Tel")
	DistributorInfo := httpHelper.Post(r, "DistributorInfo")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	AfterServiceID := httpHelper.PostInt64(r, "AfterServiceID")
	Level := httpHelper.PostInt64(r, "Level")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.DistributorNew(Token, Name, Email, Tel, DistributorInfo, CompanyID, ManagerID, AfterServiceID, Level, ID))
}

func DistributorList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	AfterServiceID := httpHelper.PostInt64(r, "AfterServiceID")
	Level := httpHelper.PostInt64(r, "Level")
	httpHelper.HttpWrite(w, bll.DistributorList(Token, Page, PageSize, Order, Stext, CompanyID, ManagerID, AfterServiceID, Level))
}

func DistributorAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	CompanyID := httpHelper.PostInt64(r, "CompanyID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	AfterServiceID := httpHelper.PostInt64(r, "AfterServiceID")
	Level := httpHelper.PostInt64(r, "Level")
	httpHelper.HttpWrite(w, bll.DistributorAll(Token, Order, Stext, CompanyID, ManagerID, AfterServiceID, Level))
}

func DistributorData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.DistributorData(Token, ID))
}

func DistributorDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.DistributorDel(Token, ID))
}
