package dbconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DBconn *sql.DB

func StartSqlConnection() {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("SQLADDR") + ":3306",
		DBName: "kaasi",
	}

	var err error
	DBconn, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Cannot connect to mysql server")
		log.Fatalln("Open:", err)
	}

	isSuccess := DBconn.Ping()
	if isSuccess != nil {
		fmt.Println("No response from mysql server")
		log.Fatalln("Ping:", isSuccess)
	}

	fmt.Println("Success!")
}
