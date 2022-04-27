package dbdocument

import (
	"database/sql"
	db "e-wallet/api/db"
	"e-wallet/api/objects"
	CustomStatus "e-wallet/api/utils"
)

func GetDocumentByID(id int) (objects.Document, error) {
	var doc objects.Document
	row := db.DBconn.QueryRow("SELECT * FROM document WHERE id = ?", id)

	if err := row.Scan(
		&doc.Id,
		&doc.DocType,
		&doc.DocNumber,
		&doc.IssuingAuthority,
		&doc.ExpiryDate,
		&doc.Img,
	); err != nil {
		if err == sql.ErrNoRows {
			return doc, CustomStatus.DocumentNotFound
		}
		return doc, err
	}
	return doc, nil
}

func AddDocument(doc objects.Document) (objects.Document, error) {
	var temp objects.Document

	result, err := db.DBconn.Exec(
		"INSERT INTO document (docType,docNumber,issuingAuthority,expiryDate,img) VALUES (?, ?, ?,?,?)",
		doc.Id,
		doc.DocType,
		doc.DocNumber,
		doc.IssuingAuthority,
		doc.ExpiryDate,
		doc.Img)
	if err != nil {
		return doc, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return doc, err
	}

	doc.Id = int(id)
	temp.Clone(doc)
	return doc, nil
}
