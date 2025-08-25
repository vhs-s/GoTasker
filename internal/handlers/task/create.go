package task

import (
	"log"
	"net/http"
	"text/template"
)

func (th *TaskHandler) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("web/template/create.html")
		if err != nil {
			log.Println("Error parse template", err)
		}
		t.Execute(w, nil)
	}
}
