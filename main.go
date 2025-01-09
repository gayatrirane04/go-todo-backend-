package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

// Task structure
type Task struct {
	ID   int
	Name string
	Done bool
}

var (
	tasks    []Task
	taskID   int
	taskLock sync.Mutex
)

// Templates
var templates = template.Must(template.ParseFiles("index.html"))

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	taskLock.Lock()
	defer taskLock.Unlock()
	templates.Execute(w, tasks)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		taskLock.Lock()
		defer taskLock.Unlock()

		r.ParseForm()
		taskName := r.FormValue("task")
		taskID++
		tasks = append(tasks, Task{ID: taskID, Name: taskName, Done: false})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func toggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		taskLock.Lock()
		defer taskLock.Unlock()

		r.ParseForm()
		id := r.FormValue("id")
		for i, task := range tasks {
			if fmt.Sprint(task.ID) == id {
				tasks[i].Done = !task.Done
				break
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Main function
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addTaskHandler)
	http.HandleFunc("/toggle", toggleTaskHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":3000", nil)
}
