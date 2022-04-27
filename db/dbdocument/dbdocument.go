package dbdocument

import (
	db "e-wallet/api/db"
	"e-wallet/api/objects"
)

func GetDocumentsOfUser(userid int) ([]objects.Document, error) {
	docs := []objects.Document{}
	rows, err := db.DBconn.Query("SELECT * FROM document WHERE userid = ?", userid)
	if err != nil {
		return docs, err
	}
	defer rows.Close()

	for rows.Next() {
		var doc objects.Document
		if err := rows.Scan(
			&doc.Id,
			&doc.DocType,
			&doc.DocNumber,
			&doc.IssuingAuthority,
			&doc.ExpiryDate,
			&doc.Img,
		); err != nil {
			return docs, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}

func AddDocument(doc objects.Document) (objects.Document, error) {
	var temp objects.Document

	result, err := db.DBconn.Exec(
		"INSERT INTO document (docType,docNumber,issuingAuthority,expiryDate,img,userid) VALUES (?,?,?,?,?)",
		doc.Id,
		doc.DocType,
		doc.DocNumber,
		doc.IssuingAuthority,
		doc.ExpiryDate,
		doc.Img, doc.UserId,
	)
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
