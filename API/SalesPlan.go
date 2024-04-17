package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func SalesPlanNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.SalesPlanNew(Token))
}

func SalesPlanList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.SalesPlanList(Token))
}

func SalesPlanAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.SalesPlanAll(Token))
}

func SalesPlanData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.SalesPlanData(Token))
}

func SalesPlanDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	httpHelper.HttpWrite(w, bll.SalesPlanDel(Token))
}
