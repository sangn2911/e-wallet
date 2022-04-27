package dbcustomer

import (
	"database/sql"
	db "e-wallet/api/db"
	"e-wallet/api/objects"
	customStatus "e-wallet/api/utils"
	"errors"
)

func GetAllCustomers() ([]objects.Customer, error) {
	objectLst := make([]objects.Customer, 0)
	rows, err := db.DBconn.Query("SELECT * FROM customer")

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

func GetCustomerWithID(id string) (objects.Customer, error) {
	var temp objects.Customer
	row := db.DBconn.QueryRow("SELECT * FROM customer WHERE id = ?", id)

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

func InsertCustomer(customer objects.Customer) (objects.Customer, error) {

	var count int
	row := db.DBconn.QueryRow("SELECT COUNT(*) FROM customer WHERE email = ?", customer.Email)

	if err := row.Scan(&count); err != nil {
		return objects.Customer{}, err

	} else {
		if count > 0 {
			return objects.Customer{}, customStatus.ExistCustomer
		}
	}

	var temp objects.Customer

	result, err := db.DBconn.Exec(
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
	temp.Clone(customer)

	return temp, nil

}

func EditCustomerInfo(customer objects.Customer) (objects.Customer, error) {
	var temp objects.Customer

	_, err := db.DBconn.Exec(
		"UPDATE customer SET firstName=?,lastName=?,dateOfBirth=?,email=?,nationality=?,address=? WHERE id = ?",
		temp.FirstName,
		temp.LastName,
		temp.DateOfBirth,
		temp.Email,
		temp.Nationality,
		temp.Address,
		temp.Id,
	)
	if err != nil {
		return customer, err
	}

	temp.Id = customer.Id
	temp.Clone(customer)
	return temp, nil
}

func DeleteCustomer(id int) error {
	_, err := db.DBconn.Exec("DELETE FROM customer WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
