package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

func Render(w http.ResponseWriter, tmplName string, context map[string]interface{}) {
	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, context)
	return
}
func Home(w http.ResponseWriter, r *http.Request) {
	Data := make(map[string]interface{})
	switch r.Method {
	case "POST":
	case "GET":
		Render(w, "./view/index.html", Data)

	default:
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method), http.StatusMethodNotAllowed)
	}
}
