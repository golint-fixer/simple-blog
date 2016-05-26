package blog

import (
	"github.com/nvhbk16k53/simple-blog/db"
)

func SaveArticle(article *db.Article) error {
	_, err := db.GetArticle(article.Id)
	if err != nil {
		return db.InsertArticle(article)
	} else {
		return db.UpdateArticle(article)
	}
}

func DeleteArticle(article *db.Article) error {
	return db.DeleteArticle(article.Id)
}

func GetArticle(id int) (*db.Article, error) {
	return db.GetArticle(id)
}

func GetFirstArticle() (*db.Article, error) {
	return db.GetFirstArticle()
}
