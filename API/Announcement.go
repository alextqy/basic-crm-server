package api

import (
	bll "basic-crm-server/BLL"
	"net/http"
)

func AnnouncementNew(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Title := httpHelper.Post(r, "Title")
	Content := httpHelper.Post(r, "Content")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AnnouncementNew(Token, Title, Content, ID))
}

func AnnouncementList(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Page := httpHelper.PostInt(r, "Page")
	PageSize := httpHelper.PostInt(r, "PageSize")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	AuthorID := httpHelper.PostInt64(r, "AuthorID")
	Display := httpHelper.PostInt64(r, "Display")
	httpHelper.HttpWrite(w, bll.AnnouncementList(Token, Page, PageSize, Order, Stext, AuthorID, Display))
}

func AnnouncementAll(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	Order := httpHelper.PostInt(r, "Order")
	Stext := httpHelper.Post(r, "Stext")
	AuthorID := httpHelper.PostInt64(r, "AuthorID")
	Display := httpHelper.PostInt64(r, "Display")
	httpHelper.HttpWrite(w, bll.AnnouncementAll(Token, Order, Stext, AuthorID, Display))
}

func AnnouncementData(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AnnouncementData(Token, ID))
}

func AnnouncementDel(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.Post(r, "ID")
	httpHelper.HttpWrite(w, bll.AnnouncementDel(Token, ID))
}

func AnnouncementDisplay(w http.ResponseWriter, r *http.Request) {
	Token := httpHelper.Post(r, "Token")
	ID := httpHelper.PostInt64(r, "ID")
	httpHelper.HttpWrite(w, bll.AnnouncementDisplay(Token, ID))
}
