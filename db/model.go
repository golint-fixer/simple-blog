package db

import (
	"errors"
	"database/sql"

	log "github.com/Sirupsen/logrus"
	_ "github.com/lib/pq"
)

type Article struct {
	Id int
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
	return &Article{Id: id, Title: title, Body: body}, nil
}

func getNextArticle(rows *sql.Rows) (*Article, error) {
	if !rows.Next() {
		log.Warn("db.getNextArticle: no more row to get")
		return nil, errors.New("End of stream")
	}
	return getArticle(rows)
}

func GetArticle(id int) (*Article, error) {
	rows, err := conn.Query("SELECT id, title, body FROM article WHERE id=$1", id)
	if err != nil {
		log.Error("db.GetArticle: ", err)
		return nil, err
	}
	return getNextArticle(rows)
}

func UpdateArticle(article *Article) error {
	_, err := conn.Exec("UPDATE article SET (title, body) = ($2, $3) WHERE id=$1", article.Id, article.Title, article.Body)
	if err != nil {
		log.Error("db.UpdateArticle: ", err)
	}
	return err
}

func InsertArticle(article *Article) error {
	_, err := conn.Exec("INSERT INTO article (id, title, body) VALUES ($1, $2, $3)", article.Id, article.Title, article.Body)
	if err != nil {
		log.Error("db.InsertArticle: ", err)
	}
	return err
}

func DeleteArticle(id int) error {
	_, err := conn.Exec("DELETE article WHERE id=$1", id)
	if err != nil {
		log.Error("db.DeleteArticle: ", err)
	}
	return err
}

func ListArticle() ([]Article, error) {
	rows, err := conn.Query("SELECT id, title, body FROM article ORDER BY id DECS")
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

func GetFirstArticle() (*Article, error) {
	rows, err := conn.Query("SELECT id, title, body FROM article ORDER BY id DECS")
	if err != nil {
		log.Error("db.GetFistArticle: ", err)
		return nil, err
	}
	return getNextArticle(rows)
}
