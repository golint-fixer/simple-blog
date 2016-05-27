package blog

import (
	"github.com/nvhbk16k53/simple-blog/db"
)

/* Save article */
func SaveArticle(article *db.Article) error {
	_, err := db.GetArticle(article.Id)
	if err != nil {
		return db.InsertArticle(article)
	} else {
		return db.UpdateArticle(article)
	}
}

/* Delete Article */
func DeleteArticle(article *db.Article) error {
	return db.DeleteArticle(article.Id)
}

/* Get Article */
func GetArticle(id int) (*db.Article, error) {
	return db.GetArticle(id)
}

/* Get First Article */
func GetFirstArticle() (*db.Article, error) {
	return db.GetFirstArticle()
}
