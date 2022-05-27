package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/rrkas/ps_practice/routes"
	"github.com/rrkas/ps_practice/utils"
)

func main() {
	utils.InitDB()

	q := utils.FetchSolutionsOfQuestion(1)
	log.Println(len(q))
	// demo.InsertDemoQuestions(len(q))

	port := 5000

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	mux := httprouter.New()

	mux.GET("/", routes.IndexPage)
	mux.GET("/questions/:ID", routes.QuestionPage)

	addr := fmt.Sprintf("127.0.0.1:%v", port)

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Print("Serving at http://", addr)
	server.ListenAndServe()
}
