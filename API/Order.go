package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func OrderNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	OrderNo := httpHelper.Post(r, "OrderNo")
	ProductID := httpHelper.PostInt64(r, "ProductID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	CustomerID := httpHelper.PostInt64(r, "CustomerID")
	DistributorID := httpHelper.PostInt64(r, "DistributorID")
	OrderPrice := httpHelper.PostFloat32(r, "OrderPrice")
	Remark := httpHelper.Post(r, "Remark")
	Type := httpHelper.PostInt64(r, "Type")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.OrderNew(Token, OrderNo, ProductID, ManagerID, CustomerID, DistributorID, OrderPrice, Remark, Type, ID))
}

func OrderList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	ProductID := httpHelper.PostInt64(r, "ProductID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	CustomerID := httpHelper.PostInt64(r, "CustomerID")
	DistributorID := httpHelper.PostInt64(r, "DistributorID")
	Type := httpHelper.PostInt64(r, "Type")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.OrderList(Token, Page, PageSize, Order, Stext, ProductID, ManagerID, CustomerID, DistributorID, Type, Status))
}

func OrderAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	ProductID := httpHelper.PostInt64(r, "ProductID")
	ManagerID := httpHelper.PostInt64(r, "ManagerID")
	CustomerID := httpHelper.PostInt64(r, "CustomerID")
	DistributorID := httpHelper.PostInt64(r, "DistributorID")
	Type := httpHelper.PostInt64(r, "Type")
	Status := httpHelper.PostInt64(r, "Status")
	httpHelper.HttpWrite(w, bll.OrderAll(Token, Order, Stext, ProductID, ManagerID, CustomerID, DistributorID, Type, Status))
}

func OrderData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.OrderData(Token, ID))
}

func OrderDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.OrderDel(Token, ID))
}
