package blog

import (
	"net/http"
	"errors"
	"regexp"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/nvhbk16k53/simple-blog/db"
)

func getId(w http.ResponseWriter, r *http.Request) (int, error) {
	validPath, err := regexp.Compile("^/(show|edit|save|delete)/([0-9]+)$")
	if err != nil {
		log.Error("Compile regexp error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 0, nil
	}
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		log.Warn("getId: Invalid Page Title")
		return 0, errors.New("Invalid Page Title")
	}
	return strconv.Atoi(m[2])
}

// IndexHandler handler request to index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("IndexHander: %s", r.URL.Path)
	validPath, err := regexp.Compile("^/$")
	if err != nil {
		log.Error("Compile regexp error: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !validPath.MatchString(r.URL.Path) {
		log.Warn("Invalid URL Path: ", r.URL.Path)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	article, err := GetFirstArticle()
	if err != nil {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	RenderArticle(w, "article", article)
}

// ShowArticleHandler handler request to show article page
func ShowArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("ShowArticleHandler: %s", r.URL.Path)
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid Title Name", http.StatusNotFound)
		return
	}
	if id == 0 {
		return
	}
	article, err := GetArticle(id)
	if err != nil {
		log.Warn("Page Not Found: ", r.URL.Path)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	RenderArticle(w, "article", article)
}

// EditArticleHandler handle request to edit page
func EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("EditArticleHandler: %s", r.URL.Path)
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid Title Name", http.StatusNotFound)
		return
	}
	if id == 0 {
		return
	}
	article, err := GetArticle(id)
	if err != nil {
		article = &db.Article{ID: id, Title: "Untitle"}
	}
	RenderArticle(w, "edit", article)
}

// DeleteArticleHandler handle request to delete page
func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("DeleteArticleHandler: %s", r.URL.Path)
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid Title Name", http.StatusNotFound)
		return
	}
	if id == 0 {
		return
	}
	article, err := GetArticle(id)
	if err != nil {
		log.Warn("DeleteArticle: ", err)
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	err = DeleteArticle(article)
	if err != nil {
		log.Error("DeleteArticle: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	RenderArticle(w, "delete", nil)
}

// SaveArticleHandler handle request to save page
func SaveArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("SaveArticleHandler: %s", r.URL.Path)
	id, err := getId(w, r)
	if err != nil {
		http.Error(w, "Invalid Title Name", http.StatusNotFound)
		return
	}
	if id == 0 {
		return
	}
	title := r.FormValue("titletext")
	body := r.FormValue("bodytext")
	article := &db.Article{ID: id, Title: title, Body: body}
	err = SaveArticle(article)
	if err != nil {
		log.Error("Save article: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/show/" + strconv.Itoa(id), http.StatusFound)
}
