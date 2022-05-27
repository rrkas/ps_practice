package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/rrkas/ps_practice/data"
	"github.com/rrkas/ps_practice/models"
)

var DB *sql.DB

func InitDB() {
	// create db
	dbPath := "data/ps.db"
	if _, err := os.Stat(dbPath); errors.Is(err, os.ErrNotExist) {
		os.Create(dbPath)
	}
	_db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	DB = _db

	// create tables
	err = models.InitQuestionTable(DB)
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = models.InitAnswerTable(DB)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("DB init successful!")
}

func FetchQuestions(sortBy string, level, pageNum, pageSize int, reverse bool) []models.Question {
	Start := int(math.Max((float64(pageNum)-1)*float64(pageSize), 0))
	var st, ascDsc string
	if reverse {
		ascDsc = "DESC"
	} else {
		ascDsc = "ASC"
	}
	if level > -1 {
		st = fmt.Sprintf(
			`SELECT * FROM %v
WHERE Level=%v
ORDER BY %v %v
LIMIT %v, %v;`,
			data.QuestionsTableName,
			level,
			sortBy,
			ascDsc,
			Start,
			pageSize,
		)
	} else {
		st = fmt.Sprintf(
			`SELECT * FROM %v
ORDER BY %v %v
LIMIT %v, %v;`,
			data.QuestionsTableName,
			sortBy,
			ascDsc,
			Start,
			pageSize,
		)
	}
	rows, err := DB.Query(st)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()
	var questions []models.Question
	for rows.Next() {
		q := models.Question{}
		var s string
		err = rows.Scan(&q.ID, &q.Title, &q.Statement, &q.InputFormat, &q.OutputFormat, &q.Level, &s, &q.DateTime)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		err = json.Unmarshal([]byte(s), &q.SampleIOs)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		questions = append(questions, q)
	}
	return questions
}

func FetchQuestionByID(ID int64) *models.Question {
	st := fmt.Sprintf(
		`SELECT * FROM %v WHERE ID=%v LIMIT 1;`,
		data.QuestionsTableName,
		ID,
	)
	rows, err := DB.Query(st)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		q := models.Question{}
		var s string
		err = rows.Scan(&q.ID, &q.Title, &q.Statement, &q.InputFormat, &q.OutputFormat, &q.Level, &s, &q.DateTime)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		err = json.Unmarshal([]byte(s), &q.SampleIOs)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		return &q
	}
	return nil
}

func FetchSolutionsOfQuestion(QuestionID int64) []models.Solution {
	st := fmt.Sprintf(
		`SELECT * from %v WHERE QuestionID=?;`,
		data.SolutionsTableName,
	)
	rows, err := DB.Query(st, QuestionID)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()
	var sols []models.Solution
	for rows.Next() {
		s := models.Solution{}
		err = rows.Scan(&s.ID, &s.QuestionID, &s.Language, &s.Code, &s.Description, &s.DateTime)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		sols = append(sols, s)
	}
	return sols
}
