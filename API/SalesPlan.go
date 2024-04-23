package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func SalesPlanNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	PlanName := httpHelper.Post(r, "PlanName")
	TargetID := httpHelper.PostInt64(r, "TargetID")
	PlanContent := httpHelper.Post(r, "PlanContent")
	Budget := httpHelper.PostFloat32(r, "Budget")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.SalesPlanNew(Token, PlanName, TargetID, PlanContent, Budget, ID))
}

func SalesPlanList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	TargetID := httpHelper.PostInt64(r, "TargetID")
	Status := httpHelper.PostInt64(r, "Status")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	httpHelper.HttpWrite(w, bll.SalesPlanList(Token, Page, PageSize, Order, Stext, TargetID, Status, ManagerID))
}

func SalesPlanAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	TargetID := httpHelper.PostInt64(r, "TargetID")
	Status := httpHelper.PostInt64(r, "Status")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	httpHelper.HttpWrite(w, bll.SalesPlanAll(Token, Order, Stext, TargetID, Status, ManagerID))
}

func SalesPlanData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.SalesPlanData(Token, ID))
}

func SalesPlanDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.SalesPlanDel(Token, ID))
}
