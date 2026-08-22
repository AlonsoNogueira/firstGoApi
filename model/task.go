package model

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Status      bool   `json:"status"`
}

const (
	TableName      = "tasks"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT FALSE)`
)
