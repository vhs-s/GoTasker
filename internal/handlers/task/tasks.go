package task

import (
	"fmt"
	"go-tasker/internal/entities"
	"net/http"
	"text/template"
)

func (th *TaskHandler) TasksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		switch r.Method {
		case http.MethodPost:
			task := entities.New(r.FormValue("TaskName"), entities.DelayConv(r.FormValue("TaskDelay")), r.FormValue("TaskMessage"))
			th.TaskRepo.Create(task)
			fmt.Println(task)
			http.Redirect(w, r, "/tasks/", http.StatusSeeOther)
			return
		case http.MethodGet:
			t, err := template.ParseFiles("web/template/tasks.html")
			if err != nil {
				fmt.Println("ERROR to parse tasks template")
			}
			data := th.TaskRepo.Get()
			fmt.Printf("%+v\n", data)
			t.Execute(w, data)
			return
		}

	}
}
