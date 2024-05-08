package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func CustomerQANew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Title := httpHelper.Post(r, "Title")
	Content := httpHelper.Post(r, "Content")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerQANew(Token, Title, Content, ID))
}

func CustomerQAList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Display := httpHelper.PostInt64(r, "Display")
	httpHelper.HttpWrite(w, bll.CustomerQAList(Token, Page, PageSize, Order, Stext, Display))
}

func CustomerQAAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	Display := httpHelper.PostInt64(r, "Display")
	httpHelper.HttpWrite(w, bll.CustomerQAAll(Token, Order, Stext, Display))
}

func CustomerQAData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerQAData(Token, ID))
}

func CustomerQADel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerQADel(Token, ID))
}

func CustomerQADisplay(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.CustomerQADisplay(Token, ID))
}
