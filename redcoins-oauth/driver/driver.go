package driver

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/oauth")
	//db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
