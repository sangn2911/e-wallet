package dbhandler

import (
	"database/sql"
	db "e-wallet/api/db"
	"e-wallet/api/objects"
	customStatus "e-wallet/api/utils"
	"errors"
)

func RegisterUser(user objects.User) (objects.User, error) {

	temp, status := IsUserExist(user.Username)

	if errors.Is(status, customStatus.UserNotFound) {
		var result sql.Result
		var err error

		result, err = db.DBconn.Exec(
			"INSERT INTO user (username, email, password) VALUES (?, ?, ?)",
			user.Username, user.Email, user.Passwd,
		)

		if err != nil {
			println("DB Error:", err.Error())
			return user, err
		}

		var id int64
		id, err = result.LastInsertId()
		if err != nil {
			println("DB Error:", err.Error())
			return user, err
		}

		temp.Id = int(id)
		temp.Clone(user)

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
	row := db.DBconn.QueryRow("SELECT * FROM user WHERE username = ?", username)

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
	row := db.DBconn.QueryRow("SELECT id, username, email FROM user WHERE id = ?", id)

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
	rows, err := db.DBconn.Query("SELECT id, username, email FROM user")

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
