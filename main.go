package main

import (
	"log"
	"net/http"

	"github.com/alnszzx/simple-crud--go/hadler"

	"github.com/alnszzx/simple-crud--go/config"
	"github.com/alnszzx/simple-crud--go/model"
	"github.com/gorilla/mux"
)

func main() {
	//Chama a função
	dbConnection := config.SetUpDataBase()

	//Encerra conexao com o banco
	defer dbConnection.Close()

	_, err := dbConnection.Exec(model.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	taskHandler := hadler.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
