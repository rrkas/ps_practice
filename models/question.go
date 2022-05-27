package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/rrkas/ps_practice/data"
	"github.com/xeonx/timeago"
)

const (
	QuestionEasy   = iota // 0
	QuestionMedium = iota // 1
	QuestionHard   = iota // 2
)

type Question struct {
	ID           int64
	Title        string
	Statement    string
	InputFormat  string
	OutputFormat string
	Level        int64 // easy, medium or hard
	SampleIOs    []IO  // examples
	DateTime     time.Time
}

func InitQuestionTable(db *sql.DB) (err error) {
	st := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	Title TEXT,
	Statement TEXT,
	InputFormat TEXT,
	OutputFormat TEXT,
	Level INTEGER,
	SampleIOs TEXT,
	DateTime DATETIME
);`, data.QuestionsTableName)
	query, err := db.Prepare(st)
	if err != nil {
		log.Println(err.Error())
		return
	}
	query.Exec()
	log.Printf("Table %v init success!", data.QuestionsTableName)
	return
}

func (q *Question) AddInDB(db *sql.DB) {
	query, err := db.Prepare(
		fmt.Sprintf(`INSERT INTO
%v(Title, Statement, InputFormat, OutputFormat, Level, SampleIOs, DateTime)
VALUES (?, ?, ?, ?, ?, ?, ?);`, data.QuestionsTableName),
	)
	if err != nil {
		log.Println(err.Error())
		return
	}
	samples, err := json.Marshal(q.SampleIOs)
	if err != nil {
		log.Println(err.Error())
		return
	}
	res, err := query.Exec(
		q.Title,
		q.Statement,
		q.InputFormat,
		q.OutputFormat,
		q.Level,
		string(samples),
		q.DateTime,
	)
	if err != nil {
		log.Println(err.Error())
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return
	}
	q.ID = int64(id)
}

func (q *Question) UpdateInDB(db *sql.DB) {
	query, err := db.Prepare(
		fmt.Sprintf(`UPDATE TABLE
%v(Statement, InputFormat, OutputFormat, Level, SampleIOs)
VALUES (?, ?, ?, ?, ?)
WHERE ID=?;`, data.QuestionsTableName),
	)
	if err != nil {
		log.Println(err.Error())
		return
	}
	samples, err := json.Marshal(q.SampleIOs)
	if err != nil {
		log.Println(err.Error())
		return
	}
	res, err := query.Exec(q.Statement, q.InputFormat, q.OutputFormat, q.Level, string(samples), q.ID)
	if err != nil {
		log.Println(err.Error())
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return
	}
	q.ID = int64(id)
}

func (q *Question) LevelString() (s string) {
	switch q.Level {
	case QuestionEasy:
		{
			s = "Easy"
		}
	case QuestionMedium:
		{
			s = "Medium"
		}
	case QuestionHard:
		{
			s = "Hard"
		}
	}
	return
}

func (q *Question) DateTimeAgo() (s string) {
	return timeago.English.Format(q.DateTime)
}

func (q *Question) DateTimeString() string {
	year, month, day := q.DateTime.Date()
	hr, min, sec := q.DateTime.Clock()
	s := fmt.Sprintf(
		"%v %v %v - %v:%v:%v",
		day,
		month.String(),
		year,
		hr,
		min,
		sec,
	)
	// fmt.Println(s)
	return s
}
