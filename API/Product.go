package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func ProductNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ProductName := httpHelper.Post(r, "ProductName")
	Price := httpHelper.PostFloat32(r, "Price")
	Cost := httpHelper.PostFloat32(r, "Cost")
	Remark := httpHelper.Post(r, "Remark")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.ProductNew(Token, ProductName, Price, Cost, Remark, ID))
}

func ProductList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.ProductList(Token, Page, PageSize, Order, Stext, Status))
}

func ProductAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.ProductAll(Token, Order, Stext, Status))
}

func ProductData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.ProductData(Token, ID))
}

func ProductDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.ProductDel(Token, ID))
}
