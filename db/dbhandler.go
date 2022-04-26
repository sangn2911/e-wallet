package dbhandler

import (
	"database/sql"
	"e-wallet/api/objects"
	customStatus "e-wallet/api/utils"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StartSqlConnection() {

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("SQLADDR")+":3306",
		DBName: "kaasi",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Cannot connect to mysql server")
		log.Fatalln("Open:", err)
	}

	isSuccess := db.Ping()
	if isSuccess != nil {
		fmt.Println("No response from mysql server")
		log.Fatalln("Ping:", isSuccess)
	}

	fmt.Println("Success!")
}

func RegisterUser(user objects.User) (objects.User, error) {

	temp, status := IsUserExist(user.Username)

	if errors.Is(status, customStatus.UserNotFound) {
		var result sql.Result
		var err error

		result, err = db.Exec(
			"INSERT INTO user (username, email, password) VALUES (?, ?, ?)",
			user.Username, user.Email, user.Passwd,
		)

		if err != nil {
			println("DB Error:", err.Error())
			return user, fmt.Errorf("db error: %v", err)
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			println("DB Error:", err.Error())
			return user, fmt.Errorf("db error: %v", err)
		}

		temp.Id = int(id)
		temp.Username = user.Username
		temp.Email = user.Email
		temp.Passwd = user.Passwd

		return temp, nil
	} else {
		return objects.User{}, status
	}

}

func LoginUser(user objects.User) (objects.User, error) {

	temp, status := IsUserExist(user.Username)

	if errors.Is(status, customStatus.ExistUser) {
		if user.Passwd == temp.Passwd {
			return temp, nil
		} else {
			return objects.User{}, customStatus.WrongPasswd
		}
	} else {
		return user, customStatus.UserNotFound
	}

}

func IsUserExist(username string) (objects.User, error) {
	var temp objects.User
	row := db.QueryRow("SELECT * FROM user WHERE username = ?", username)

	if err := row.Scan(&temp.Id, &temp.Username, &temp.Email, &temp.Passwd); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			println("NotFound")
			return objects.User{}, customStatus.UserNotFound
		} else {
			println("SQL Error:", err.Error())
			return objects.User{}, err
		}
	}

	return temp, customStatus.ExistUser
}

func GetUserWithID(id string) (objects.User, error) {
	var temp objects.User
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)

	if err := row.Scan(&temp.Id, &temp.Username, &temp.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return objects.User{}, customStatus.UserNotFound
		} else {
			return objects.User{}, err
		}
	}

	return temp, customStatus.ExistUser
}
