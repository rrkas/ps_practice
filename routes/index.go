package routes

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rrkas/ps_practice/utils"
)

func IndexPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryParams := r.URL.Query()
	var sortBy string
	var level, pageNum, pageSize int
	var reverse bool

	if queryParams.Has("sortBy") {
		sortBy = queryParams.Get("sortBy")
	} else {
		sortBy = "ID"
	}
	if queryParams.Has("level") {
		var err error
		level, err = strconv.Atoi(queryParams.Get("level"))
		if err != nil {
			log.Println(err.Error())
			level = -1
		}
	} else {
		level = -1
	}
	if queryParams.Has("pageNum") {
		var err error
		pageNum, err = strconv.Atoi(queryParams.Get("pageNum"))
		if err != nil {
			log.Println(err.Error())
			pageNum = 1
		}
		if pageNum < 1 {
			pageNum = 1
		}
	} else {
		pageNum = 1
	}
	if queryParams.Has("pageSize") {
		var err error
		pageSize, err = strconv.Atoi(queryParams.Get("pageSize"))
		if err != nil {
			log.Println(err.Error())
			pageSize = 20
		}
		if pageSize < 10 {
			pageSize = 10
		}
	} else {
		pageSize = 20
	}
	if queryParams.Has("reverse") {
		var err error
		reverse, err = strconv.ParseBool(queryParams.Get("reverse"))
		if err != nil {
			log.Println(err.Error())
			reverse = false
		}
	} else {
		reverse = false
	}

	q := utils.FetchQuestions(
		sortBy,
		level,
		pageNum,
		pageSize,
		reverse,
	)
	GenerateHTML(
		w,
		map[string]interface{}{
			"Questions": q,
			"pageNum":   pageNum,
			"pageSize":  pageSize,
			"reverse":   reverse,
			"level":     level,
			"sortBy":    sortBy,
			"lenq":      len(q),
		},
		[]string{
			"index",
			"layout",
			"navbar.home",
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
