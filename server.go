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

	mux := httprouter.New()
	mux.GET("/", routes.IndexPage)
	mux.GET("/questions/:ID", routes.QuestionPage)
	mux.GET("/question/new", routes.QuestionNewPage)
	mux.POST("/question/save", routes.QuestionSavePage)
	mux.ServeFiles("/static/*filepath", http.Dir("./static"))

	port := 5000
	addr := fmt.Sprintf("127.0.0.1:%v", port)

	server := http.Server{Addr: addr, Handler: mux}

	log.Print("Serving at http://", addr)
	server.ListenAndServe()
}
