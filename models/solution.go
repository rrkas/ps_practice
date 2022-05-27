package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/rrkas/ps_practice/data"
)

type Solution struct {
	ID          int64
	QuestionID  int64
	Language    string
	Code        string
	Description string
	DateTime    time.Time
}

func InitAnswerTable(db *sql.DB) (err error) {
	st := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %v (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	QuestionID INTEGER,
	Language TEXT,
	Code TEXT,
	Description TEXT,
	DateTime DATETIME
);`,
		data.SolutionsTableName,
	)
	query, err := db.Prepare(st)
	if err != nil {
		log.Println(err.Error())
		return
	}
	query.Exec()
	log.Printf("Table %v init success!", data.SolutionsTableName)
	return
}
