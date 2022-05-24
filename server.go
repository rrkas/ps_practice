package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 5000

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: mux,
	}

	fmt.Println("Serving at port:", port)
	server.ListenAndServe()
}
