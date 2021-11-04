package database

import (
	"database/sql"
	"fmt"
	"time"

	// "log"

	"github.com/getground/tech-tasks/backend/config"
	"github.com/getground/tech-tasks/backend/logger"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConn() (db *sql.DB) {
	source := fmt.Sprintf("%s:%s@tcp(godockerDB)/%s", config.Env.DBUser, config.Env.DBPassword, config.Env.DBName)
	db, err := sql.Open("mysql", source)
	if err != nil {
		logger.Error.Fatalf("Setup MySQL connect error %+v\n", err)
	}
	return db
}
func Setup() {
	// init mysql.

	// db, err := sql.Open("mysql", "user:password@/database")
	db := DbConn()

	defer db.Close()

	// // Open doesn't open a connection. Validate DSN data:
	// err := db.Ping()
	// if err != nil {
	// 	logger.Error.Fatalf("Setup MySQL ping error %+v\n", err)
	// }
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("MySQL connection successful")

	// sql := `CREATE TABLE guest_table (
	// 	table_number INT NOT NULL,
	// 	space int,
	// 	PRIMARY KEY (table_number)
	//   );`

	// if _, err := DB.Exec(sql); err != nil {
	// 	fmt.Println("create table failed:", err)
	// 	return
	// }
	// fmt.Println("create table successd")
	// log.Println("MySQL connection successful")

	// insert, err := DB.Query("INSERT INTO guest_table VALUES ( 1, 10 )")
	// // insert, err := db.Query("INSERT INTO guest_table VALUES ( 2, 10 )")
	// // insert, err := db.Query("INSERT INTO guest_table VALUES ( 3, 10 )")

	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // be careful deferring Queries if you are using transactions
	// defer insert.Close()
}

type Tag struct {
	Table_number int `json:"table_number"`
	Space        int `json:"space"`
}

// func DeleteData() {
// 	result, err := db.Exec("delete from guest_table where table_number=?", 1)
// 	if err != nil {
// 		fmt.Printf("delete failed,err:%v\n", err)
// 		return
// 	}
// 	fmt.Println("delete data successd:", result)

// 	rowsaffected, err := result.RowsAffected()
// 	if err != nil {
// 		fmt.Printf("delete RowsAffected failed,err:%v\n", err)
// 		return
// 	}
// 	fmt.Println("delete Affected rows:", rowsaffected)
// }

// func SelectData() {
// 	var tag Tag
// 	// Execute the query
// 	db.QueryRow("SELECT table_number, Space FROM guest_table", 2).Scan(&tag.Table_number, &tag.Space)
// 	// if err != nil {
// 	// 	panic(err.Error()) // proper error handling instead of panic in your app
// 	// }

// 	log.Println(tag.Table_number)
// 	log.Println(tag.Space)

// 	// for results.Next() {
// 	// 	var tag Tag
// 	// 	// for each row, scan the result into our tag composite object
// 	// 	err = results.Scan(&tag.Table_number, &tag.Space)
// 	// 	if err != nil {
// 	// 		panic(err.Error()) // proper error handling instead of panic in your app
// 	// 	}
// 	// and then print out the tag's Name attribute
// 	// log.Println(tag.Table_number)
// 	// log.Println(tag.Space)
// 	// }
// }
func InsertData() {
	db := DbConn()
	insDB, err := db.Prepare("INSERT INTO guest_table(table_number, space) VALUES(?,?)")
	if err != nil {
		fmt.Printf("Insert data failed,err:%v", err)
		return
	}
	insDB.Exec(3, 10)
	defer db.Close()
}
