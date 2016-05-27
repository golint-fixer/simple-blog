package main

import (
	"testing"

	"github.com/nvhbk16k53/simple-blog/db"
)

func TestDBInsertArticle(t *testing.T) {
	// Success insert article into database
	article := db.Article{Id: 20, Title: "test", Body: "test"}
	err := db.InsertArticle(&article)
	if err != nil {
		t.Error("TestDBInsertArticle: ", err)
	}

	// Fail insert article into database
	err = db.InsertArticle(&article)
	if err == nil {
		t.Error("TestDBInsertArticle: nil")
	}

	// Delete article from database
	err = db.DeleteArticle(article.Id)
	if err != nil {
		log.Fatal("TestDBInsertArticle: Delete article error: ", err)
	}
}

func TestDBDeleteArticle(t *testing.T) {
	// Insert article into database
	article := db.Article{Id: 20, Title: "test", Body: "test"}
	err := db.InsertArticle(&article)
	if err != nil {
		log.Fatal("TestDBDeleteArticle: Create article error: ", err)
	}

	// Success delete article from database
	err = db.DeleteArticle(article.Id)
	if err != nil {
		t.Error("TestDBDeleteArticle: ", err)
	}

	// Fail delete article from database
	err = db.DeleteArticle(article.Id)
	if err != nil {
		t.Error("TestDBDeleteArticle: nil")
	}
}

func TestDBGetArticle(t *testing.T) {
	// Insert article into database
	article := db.Article{Id: 20, Title: "test", Body: "test"}
	err := db.InsertArticle(&article)
	if err != nil {
		log.Fatal("TestDBGetArticle: Create article error: ", err)
	}

	// Success get article from database
	newArticle, err := db.GetArticle(20)
	if err != nil {
		t.Error("TestDBGetArticle: ", err)
	}
	if article != *newArticle {
		t.Error("TestDBGetArticle: article != newArticle")
	}

	// Delete article from database
	err = db.DeleteArticle(article.Id)
	if err != nil {
		log.Fatal("TestDBGetArticle: Delete article error: ", err)
	}

	// Fail get article from database
	_, err = db.GetArticle(article.Id)
	if err == nil {
		t.Error("TestDBGetArticle: nil")
	}
}

func TestDBUpdateArticle(t *testing.T) {
	// Insert article into database
	article := db.Article{Id: 20, Title: "test", Body: "test"}
	err := db.InsertArticle(&article)
	if err != nil {
		log.Fatal("TestDBUpdateArticle: Create article error: ", err)
	}

	// Success update article in database
	err = db.UpdateArticle(&article)
	if err != nil {
		t.Error("TestDBUpdateArticle: ", err)
	}

	// Delete article from database
	err = db.DeleteArticle(article.Id)
	if err != nil {
		log.Fatal("TestDBUpdateArticle: Delete article error: ", err)
	}

	// Fail update article in database
	err = db.UpdateArticle(&article)
	if err != nil {
		t.Error("TestDBUpdateArticle: nil")
	}
}
