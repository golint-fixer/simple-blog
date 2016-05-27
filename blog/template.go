package blog

import (
	"html/template"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/nvhbk16k53/simple-blog/db"
)

var templatePath = "template/"

var templates = template.Must(template.ParseGlob(templatePath + "*.tmpl"))

// RenderIndex render index page
func RenderIndex(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "index.tmpl", nil)
	if err != nil {
		log.Errorf("Cannot execute template index.tmpl: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderArticle render article page
func RenderArticle(w http.ResponseWriter, tmpl string, article *db.Article) {
	err := templates.ExecuteTemplate(w, tmpl + ".tmpl", article)
	if err != nil {
		log.Errorf("Cannot execute template %v.tmpl: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
