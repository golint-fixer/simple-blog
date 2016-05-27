package db

import (
	"errors"
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
)

// Article data type of article record
type Article struct {
	ID int
	Title string
	Body string
}

var conn *sql.DB

func init() {
	var err error
	conn, err = sql.Open("postgres", "user=hiepnv dbname=blog password=dunghoianh")
	if err != nil {
		log.Fatal("db.init:", err)
	}
}

func getArticle(rows *sql.Rows) (*Article, error) {
	var id int
	var title, body string
	err := rows.Scan(&id, &title, &body)
	if err != nil {
		log.Error("db.getArticle: ", err)
		return nil, err
	}
	return &Article{ID: id, Title: title, Body: body}, nil
}

func getNextArticle(rows *sql.Rows) (*Article, error) {
	if !rows.Next() {
		log.Warn("db.getNextArticle: no more row to get")
		return nil, errors.New("End of stream")
	}
	return getArticle(rows)
}

// GetArticle get an article by id
func GetArticle(id int) (*Article, error) {
	rows, err := conn.Query("SELECT id, title, body FROM article WHERE id=$1", id)
	if err != nil {
		log.Error("db.GetArticle: ", err)
		return nil, err
	}
	return getNextArticle(rows)
}

// UpdateArticle update article information
func UpdateArticle(article *Article) error {
	_, err := conn.Exec("UPDATE article SET (title, body) = ($2, $3) WHERE id=$1", article.ID, article.Title, article.Body)
	if err != nil {
		log.Error("db.UpdateArticle: ", err)
	}
	return err
}

// InsertArticle insert article information
func InsertArticle(article *Article) error {
	_, err := conn.Exec("INSERT INTO article (id, title, body) VALUES ($1, $2, $3)", article.ID, article.Title, article.Body)
	if err != nil {
		log.Error("db.InsertArticle: ", err)
	}
	return err
}

// DeleteArticle delete an article record
func DeleteArticle(id int) error {
	_, err := conn.Exec("DELETE FROM article WHERE id=$1", id)
	if err != nil {
		log.Error("db.DeleteArticle: ", err)
	}
	return err
}

// ListArticle list all article
func ListArticle() ([]Article, error) {
	rows, err := conn.Query("SELECT id, title, body FROM article ORDER BY id DESC")
	if err != nil {
		log.Error("db.ListArticle: ", err)
		return make([]Article, 0), err
	}
	articles := make([]Article, 0)
	for rows.Next() {
		article, err := getArticle(rows)
		if err != nil {
			break;
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

// GetFirstArticle get first artilce record
func GetFirstArticle() (*Article, error) {
	rows, err := conn.Query("SELECT id, title, body FROM article ORDER BY id DESC")
	if err != nil {
		log.Error("db.GetFistArticle: ", err)
		return nil, err
	}
	return getNextArticle(rows)
}
