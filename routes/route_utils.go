package routes

import (
	"fmt"
	"html/template"
	"net/http"
)

func GenerateHTML(w http.ResponseWriter, data interface{}, tmplFileNames []string, funcMap template.FuncMap) {
	var files []string
	for _, file := range tmplFileNames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.New("").Funcs(funcMap).ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
