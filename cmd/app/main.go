package main

import (

	// "log"
	// "net/http"

	"github.com/getground/tech-tasks/backend/config"
	"github.com/getground/tech-tasks/backend/database"
	"github.com/getground/tech-tasks/backend/logger"
	"github.com/getground/tech-tasks/backend/router"
	_ "github.com/go-sql-driver/mysql"
)

func Setup() {
	logger.Setup()
	config.Setup()
	database.Setup()
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run(":3000")
	// init mysql.
	// db, err := sql.Open("mysql", "user:password@/getground")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// // ping
	// http.HandleFunc("/ping", handlerPing)
	// http.ListenAndServe(":3000", nil)
}

// func handlerPing(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "pong\n")
// }
