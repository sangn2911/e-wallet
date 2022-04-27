package dbtransaction

import (
	db "e-wallet/api/db"
	"e-wallet/api/objects"
	"fmt"
)

func GetAllTransactions() ([]objects.Transaction, error) {
	var transacLst []objects.Transaction

	rows, err := db.DBconn.Query("SELECT * FROM transaction")
	if err != nil {
		return nil, fmt.Errorf("GetTransaction: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var transaction objects.Transaction
		if err := rows.Scan(
			&transaction.Id,
			&transaction.SenderName,
			&transaction.ReceiverName,
			&transaction.Date,
			&transaction.Money,
			&transaction.Message,
		); err != nil {
			return nil, err
		}
		transacLst = append(transacLst, transaction)
	}

	return transacLst, nil
}

func AddTransaction(transaction objects.Transaction) (objects.Transaction, error) {

	var temp objects.Transaction
	result, err := db.DBconn.Exec(
		"INSERT INTO transaction (senderName,receiverName,date,money,message) VALUES (?,?,?,?,?)",
		transaction.Id,
		transaction.SenderName,
		transaction.ReceiverName,
		transaction.Date,
		transaction.Money,
		transaction.Message,
	)
	if err != nil {
		return temp, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return temp, err
	}

	temp.Id = int(id)
	temp.Clone(transaction)

	return temp, nil
}

func DeleteTransaction(id int) error {
	_, err := db.DBconn.Exec("DELETE  FROM transaction WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("DeleteAffiliate: %v", err)
	}

	return nil
}
