package database

import (
	"database/sql"
	"fmt"

	// "log"

	"github.com/getground/tech-tasks/backend/logger"
	_ "github.com/go-sql-driver/mysql"
)

func Setup() {
	// init mysql.
	db, err := sql.Open("mysql", "user:password@/getground")
	if err != nil {
		logger.Error.Fatalf("Setup MongoDB connect error %+v\n", err)
	}
	defer db.Close()
	fmt.Println("MySQL connection successful")

}
