package routes

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rrkas/ps_practice/models"
	"github.com/rrkas/ps_practice/utils"
)

func QuestionPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	qidStr := p.ByName("ID")
	qid, err := strconv.Atoi(qidStr)
	if err != nil {
		log.Println(err.Error())
		QuestionErrorPage(w, r, qidStr, "Question doesn't exist!")
		return
	}
	q := utils.FetchQuestionByID(int64(qid))
	GenerateHTML(
		w,
		map[string]interface{}{
			"question": q,
		},
		[]string{
			"question.detail",
			"layout",
			"navbar.question",
			"question.row",
		},
		template.FuncMap{
			"prev": func(i int) int {
				return i - 1
			},
			"next": func(i int) int {
				return i + 1
			},
		},
	)
}

func QuestionErrorPage(w http.ResponseWriter, r *http.Request, ID interface{}, err string) {
	GenerateHTML(
		w,
		map[string]interface{}{
			"id":    ID,
			"error": err,
		},
		[]string{
			"question.detail",
			"layout",
			"navbar.home",
			"question.row",
		},
		template.FuncMap{},
	)
}

func QuestionNewPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	GenerateHTML(
		w,
		nil,
		[]string{
			"layout",
			"question.form",
			"navbar.home",
		},
		template.FuncMap{},
	)
}

func QuestionSavePage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	formData := r.Form
	q := models.ParseQuestionForm(formData)
	str, err := json.Marshal(q)
	log.Println(string(str))
	q.AddInDB(utils.DB)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
