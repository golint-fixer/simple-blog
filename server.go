package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"./blog"
)

var log = logrus.New()

func main() {
	log.Info("Running server on *:8080")
	http.HandleFunc("/", blog.IndexHandler)
	http.HandleFunc("/show/", blog.ShowArticleHandler)
	http.HandleFunc("/edit/", blog.EditArticleHandler)
	http.HandleFunc("/save/", blog.SaveArticleHandler)
	http.HandleFunc("/delete/", blog.DeleteArticleHandler)
	http.ListenAndServe(":8080", nil)
}
