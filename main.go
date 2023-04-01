package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"

	"goSqlRaw/configs"
	"goSqlRaw/connection"
	"goSqlRaw/handlers"
	"goSqlRaw/utils"
)

const port = 8005

var initDBTableOnce sync.Once

func init() {
	dbConfig := configs.NewDBConfig()
	dbInstance := connection.ConnectDB(dbConfig)

	initDBTableOnce.Do(func() {
		err := dbInstance.Migration()
		if err != nil {
			log.Println("failed to create table ", "error: ", err)
		}

		err = dbInstance.InsertInitialDataIntoTable()
		if err != nil {
			log.Println("failed to insert initial batch data", "error: ", err)
		}
		err = utils.InitFileWriter()
		if err != nil {
			log.Println("failed to initialize file writer")
		}
	})
}

func main() {
	router := chi.NewRouter()
	handler := handlers.NewHandler()
	router.Route("/api", handler.Handle)
	fmt.Println("server running on port 8005")
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
