package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func GroupNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	GroupName := httpHelper.Post(r, "GroupName")
	Remark := httpHelper.Post(r, "Remark")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.GroupNew(Token, GroupName, Remark, ID))
}

func GroupList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	httpHelper.HttpWrite(w, bll.GroupList(Token, Page, PageSize, Order, Stext))
}

func GroupAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	httpHelper.HttpWrite(w, bll.GroupAll(Token, Order, Stext))
}

func GroupData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.GroupData(Token, ID))
}

func GroupDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.GroupDel(Token, ID))
}
