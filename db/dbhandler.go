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
		Addr:   os.Getenv("SQLADDR") + ":3306",
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
	} else if errors.Is(status, customStatus.UserNotFound) {
		return user, customStatus.UserNotFound
	} else {
		return objects.User{}, status
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
	row := db.QueryRow("SELECT id, username, email FROM user WHERE id = ?", id)

	if err := row.Scan(&temp.Id, &temp.Username, &temp.Email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return objects.User{}, customStatus.UserNotFound
		} else {
			return objects.User{}, err
		}
	}

	return temp, nil
}

func GetAllUsers() ([]objects.User, error) {
	objectLst := make([]objects.User, 0)
	rows, err := db.Query("SELECT id, username, email FROM user")

	if err != nil {
		return objectLst, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp objects.User
		if err := rows.Scan(&temp.Id, &temp.Username, &temp.Email); err != nil {
			return objectLst, err
		}
		objectLst = append(objectLst, temp)
	}

	return objectLst, nil
}

func GetCustomerWithID(id string) (objects.Customer, error) {
	var temp objects.Customer
	row := db.QueryRow("SELECT * FROM customer WHERE id = ?", id)

	if err := row.Scan(
		&temp.Id,
		&temp.FirstName,
		&temp.LastName,
		&temp.DateOfBirth,
		&temp.Email,
		&temp.Nationality,
		&temp.Address,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return objects.Customer{}, customStatus.CustomerNotFound
		} else {
			return objects.Customer{}, err
		}
	}

	return temp, nil
}

func GetAllCustomers() ([]objects.Customer, error) {
	objectLst := make([]objects.Customer, 0)
	rows, err := db.Query("SELECT * FROM customer")

	if err != nil {
		return objectLst, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp objects.Customer
		if err := rows.Scan(
			&temp.Id,
			&temp.FirstName,
			&temp.LastName,
			&temp.DateOfBirth,
			&temp.Email,
			&temp.Nationality,
			&temp.Address,
		); err != nil {
			return objectLst, err
		}
		objectLst = append(objectLst, temp)
	}

	return objectLst, nil
}

func InsertCustomer(customer objects.Customer) (objects.Customer, error) {

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM customer WHERE email = ?", customer.Email)

	if err := row.Scan(&count); err != nil {
		return objects.Customer{}, err

	} else {
		if count > 0 {
			return objects.Customer{}, customStatus.ExistCustomer
		}
	}

	var temp objects.Customer

	result, err := db.Exec(
		"INSERT INTO customer (firstName,lastName,dateOfBirth,email,nationality,address) VALUES (?,?,?,?,?,?)", customer.LastName,
		customer.FirstName,
		customer.DateOfBirth,
		customer.Email,
		customer.Nationality,
		customer.Address,
	)

	if err != nil {
		return objects.Customer{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return objects.Customer{}, err
	}

	temp.Id = int(id)
	temp.LastName = customer.LastName
	temp.FirstName = customer.FirstName
	temp.DateOfBirth = customer.DateOfBirth
	temp.Email = customer.Email
	temp.Nationality = customer.Nationality
	temp.Address = customer.Address

	return temp, nil

}
