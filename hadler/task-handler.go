package hadler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/alnszzx/simple-crud--go/model"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (taskHandler TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {
	tasks := make([]model.Task, 0)

	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	writer.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(writer).Encode(tasks)
}

func (taskHandler TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task model.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = taskHandler.DB.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3);",
		task.Title,
		task.Description,
		task.Status)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}
