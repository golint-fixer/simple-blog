package blog

import (
	"github.com/nvhbk16k53/simple-blog/db"
)

// SaveArticle save an article information
func SaveArticle(article *db.Article) error {
	_, err := db.GetArticle(article.ID)
	if err != nil {
		return db.InsertArticle(article)
	}
	return db.UpdateArticle(article)
}

// DeleteArticle delete an article record
func DeleteArticle(article *db.Article) error {
	return db.DeleteArticle(article.ID)
}

// GetArticle get an article record
func GetArticle(id int) (*db.Article, error) {
	return db.GetArticle(id)
}

// GetFirstArticle get first article record
func GetFirstArticle() (*db.Article, error) {
	return db.GetFirstArticle()
}
