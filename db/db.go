package db

import (
	"api/e-wallet/entities"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"strconv"
)

var db *sql.DB

func Connection() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "Anhtripro@113doma",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "kaasi",
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("an lz cho nay roi")
		//log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("awn lz khuc nay luon")
		//log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
func GetUsers() ([]entities.User, error) {
	Connection()
	var users []entities.User

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		return nil, fmt.Errorf("GetUsers: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
			return nil, fmt.Errorf("GetUsers: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetUsers: %v", err)
	}
	return users, nil
}
func GetUserById(id int) (entities.User, error) {
	Connection()
	var user entities.User
	row := db.QueryRow("SELECT * from user WHERE id = ?", id)
	if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("GetUserByID %d: No such user", id)
		}
		return user, fmt.Errorf("GetUserByID %d, %v", id, err)
	}
	return user, nil
}
func GetCustomers() ([]entities.Customer, error) {
	Connection()
	var customers []entities.Customer
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		return nil, fmt.Errorf("GetCustomers: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var customer entities.Customer
		if err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.DateOfBirth, &customer.Email, &customer.Nationality, &customer.Address); err != nil {
			return nil, fmt.Errorf("GetCustomers: %v", err)
		}
		customers = append(customers, customer)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetCustomers: %v", err)
	}
	return customers, nil

}
func GetCustomerByID(id int) (entities.Customer, error) {
	Connection()
	var customer entities.Customer
	row := db.QueryRow("SELECT * from customer WHERE id = ?", id)
	if err := row.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.DateOfBirth, &customer.Email, &customer.Nationality, &customer.Address); err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("GetCustomerByID %d: No such user", id)
		}
		return customer, fmt.Errorf("GetCustomerByID %d, %v", id, err)
	}
	return customer, nil
}

func AddCustomer(fname, lname, dob, email, nation, address string) (entities.Customer, error) {
	Connection()
	var customer entities.Customer
	result, err := db.Exec("INSERT INTO customer (firstName, lastName,dateOfBirth,  email, nationality, address) VALUES (?, ?, ?,?,?,?)", fname, lname, dob, email, nation, address)
	if err != nil {
		return customer, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return customer, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Println(id)
	customer.Id = strconv.Itoa(int(id))
	customer.LastName = lname
	customer.FirstName = fname
	customer.DateOfBirth = dob
	customer.Email = email
	customer.Nationality = nation
	customer.Address = address
	return customer, nil

}
func EditCustomer(id int, fname, lname, dob, email, nation, address string) (entities.Customer, error) {
	Connection()
	var customer entities.Customer
	_, err := db.Exec("UPDATE customer SET firstName=?,lastName =?, dateOfBirth=?, email=? ,nationality=?, address=? WHERE  id=?", fname, lname, dob, email, nation, address, id)
	if err != nil {
		return customer, fmt.Errorf("addAlbum: %v", err)
	}

	if err != nil {
		return customer, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Println(id)

	customer.LastName = lname
	customer.FirstName = fname
	customer.DateOfBirth = dob
	customer.Email = email
	customer.Nationality = nation
	customer.Address = address
	return customer, nil
}
func DeleteCustomer(id string) error {
	Connection()
	tmp, _ := strconv.Atoi(id)
	_, err := db.Exec("DELETE  FROM customer WHERE id=?", tmp)
	if err != nil {
		return fmt.Errorf("DeleteCustomer: %v", err)
	}
	return nil
}
func GetTransactions() ([]entities.Transaction, error) {
	Connection()
	var transactions []entities.Transaction
	rows, err := db.Query("SELECT * FROM transaction")
	if err != nil {
		return nil, fmt.Errorf("GetTransaction: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var transaction entities.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.SenderName, &transaction.ReceiverName, &transaction.Date, &transaction.Money, &transaction.Message); err != nil {
			return nil, fmt.Errorf("GetTransaction: %v", err)
		}
		transactions = append(transactions, transaction)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetTransactions: %v", err)
	}
	return transactions, nil
}
func GetAffiliates() ([]entities.Affiliate, error) {
	Connection()
	var affiliates []entities.Affiliate
	rows, err := db.Query("SELECT * FROM affiliate")
	if err != nil {
		return nil, fmt.Errorf("GetAffiliates: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var affiliate entities.Affiliate
		if err := rows.Scan(&affiliate.Id, &affiliate.AffiliateName, &affiliate.District, &affiliate.Address, &affiliate.PhoneNumber, &affiliate.Fax, &affiliate.Email); err != nil {
			return nil, fmt.Errorf("GetTransaction: %v", err)
		}
		affiliates = append(affiliates, affiliate)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAffliates: %v", err)
	}
	return affiliates, nil
}
func GetAffiliateByID(id int) (entities.Affiliate, error) {
	Connection()
	var affiliate entities.Affiliate
	row := db.QueryRow("SELECT * from affiliate WHERE id = ?", id)
	if err := row.Scan(&affiliate.Id, &affiliate.AffiliateName, &affiliate.District, &affiliate.Address, &affiliate.PhoneNumber, &affiliate.Fax, &affiliate.Email); err != nil {
		if err == sql.ErrNoRows {
			return affiliate, fmt.Errorf("GetUserByID %d: No such user", id)
		}
		return affiliate, fmt.Errorf("GetUserByID %d, %v", id, err)
	}
	return affiliate, nil
}
func AddAffiliate(name, district, address, phoneNumber, fax, email string) (entities.Affiliate, error) {
	Connection()
	var affiliate entities.Affiliate
	result, err := db.Exec("INSERT INTO affiliate ( affiliateName, district, address, phoneNumber, fax, email) VALUES (?, ?, ?,?,?,?)", name, district, address, phoneNumber, fax, email)
	if err != nil {
		return affiliate, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return affiliate, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Println(id)
	affiliate.Id = strconv.Itoa(int(id))
	affiliate.AffiliateName = name
	affiliate.District = district
	affiliate.Address = address
	affiliate.PhoneNumber = phoneNumber
	affiliate.Fax = fax
	affiliate.Email = email
	return affiliate, nil

}
func DeleteAffiliate(id string) error {
	Connection()
	tmp, _ := strconv.Atoi(id)
	_, err := db.Exec("DELETE  FROM affiliate WHERE id=?", tmp)
	if err != nil {
		return fmt.Errorf("DeleteAffiliate: %v", err)
	}
	return nil
}
func AddTransaction(send, receive, date, money, message string) (entities.Transaction, error) {
	Connection()
	var tran entities.Transaction
	result, err := db.Exec("INSERT INTO transaction ( senderName, receiverName, date, money, message) VALUES (?, ?, ?,?,?)", send, receive, date, money, message)
	if err != nil {
		return tran, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return tran, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Println(id)
	tran.Id = strconv.Itoa(int(id))
	tran.SenderName = send
	tran.ReceiverName = receive
	tran.Date = date
	tran.Money = money
	tran.Message = message
	return tran, nil
}
func DeleteTransaction(id string) error {
	Connection()
	tmp, _ := strconv.Atoi(id)
	_, err := db.Exec("DELETE  FROM transaction WHERE id=?", tmp)
	if err != nil {
		return fmt.Errorf("DeleteAffiliate: %v", err)
	}
	return nil
}
func AddUser(username, email, password string) (entities.User, error) {
	Connection()
	var user entities.User
	result, err := db.Exec("INSERT INTO user ( username, email, password) VALUES (?, ?, ?)", username, email, password)
	if err != nil {
		return user, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Println(id)
	user.Id = strconv.Itoa(int(id))
	user.Username = username
	user.Email = email
	user.Password = password
	return user, nil
}
func GetDocumentByID(id int) (entities.Document, error) {
	Connection()
	var doc entities.Document
	row := db.QueryRow("SELECT * FROM document WHERE id = ?", id)
	if err := row.Scan(&doc.Id, &doc.DocType, &doc.DocNumber, &doc.IssuingAuthority, &doc.ExpiryDate, &doc.Img); err != nil {
		if err == sql.ErrNoRows {
			return doc, fmt.Errorf("GetDocByID %d: No such user", id)
		}
		return doc, fmt.Errorf("GetDocByID %d, %v", id, err)
	}
	return doc, nil

}
func AddDocument(doctype, docnum, issuing, expiry, img string) (entities.Document, error) {
	Connection()
	var doc entities.Document
	result, err := db.Exec("INSERT INTO document ( docType, docNumber, issuingAuthority,expiryDate, img) VALUES (?, ?, ?,?,?)", doctype, docnum, issuing, expiry, img)
	if err != nil {
		return doc, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return doc, fmt.Errorf("addAlbum: %v", err)
	}
	fmt.Println(id)
	doc.Id = strconv.Itoa(int(id))
	doc.DocType = doctype
	doc.DocNumber = docnum
	doc.IssuingAuthority = issuing
	doc.ExpiryDate = expiry
	doc.Img = img
	return doc, nil
}
