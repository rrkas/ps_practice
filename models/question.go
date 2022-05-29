package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
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

func ParseQuestionForm(values url.Values) (q Question) {
	if values.Has("Title") {
		q.Title = values.Get("Title")
	}
	if values.Has("Statement") {
		q.Statement = values.Get("Statement")
	}
	if values.Has("InputFormat") {
		q.InputFormat = values.Get("InputFormat")
	}
	if values.Has("OutputFormat") {
		q.OutputFormat = values.Get("OutputFormat")
	}
	if values.Has("Level") {
		level, err := strconv.Atoi(values.Get("Level"))
		if err != nil {
			log.Println(err)
		} else {
			q.Level = int64(level)
		}
	}
	var ios []IO

	for i := 0; i < 3; i++ {
		var sampleInput, sampleOutput, description string
		siField := fmt.Sprintf("SampleInput%v", i)
		soField := fmt.Sprintf("SampleOutput%v", i)
		descField := fmt.Sprintf("Explanation%v", i)
		if values.Has(siField) {
			sampleInput = strings.Trim(values.Get(siField), " ")
		}
		if values.Has(soField) {
			sampleOutput = strings.Trim(values.Get(soField), " ")
		}
		if values.Has(descField) {
			description = strings.Trim(values.Get(descField), " ")
		}
		if len(sampleInput) > 0 && len(sampleOutput) > 0 {
			t := IO{sampleInput, sampleOutput, description}
			ios = append(ios, t)
		}
	}

	q.SampleIOs = ios
	q.DateTime = time.Now()
	return
}
