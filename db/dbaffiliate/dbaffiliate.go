package dbaffiliate

import (
	"database/sql"
	db "e-wallet/api/db"
	"e-wallet/api/objects"
	CustomStatus "e-wallet/api/utils"
)

func GetAffiliates() ([]objects.Affiliate, error) {

	var affiliates []objects.Affiliate
	rows, err := db.DBconn.Query("SELECT * FROM affiliate")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var affiliate objects.Affiliate
		if err := rows.Scan(
			&affiliate.Id,
			&affiliate.AffiliateName,
			&affiliate.District,
			&affiliate.Address,
			&affiliate.PhoneNumber,
			&affiliate.Fax,
			&affiliate.Email,
		); err != nil {
			return nil, err
		}
		affiliates = append(affiliates, affiliate)
	}

	return affiliates, nil
}
func GetAffiliateByID(id string) (objects.Affiliate, error) {

	var affi objects.Affiliate
	row := db.DBconn.QueryRow("SELECT * from affiliate WHERE id = ?", id)
	if err := row.Scan(
		&affi.Id,
		&affi.AffiliateName,
		&affi.District,
		&affi.Address,
		&affi.PhoneNumber,
		&affi.Fax,
		&affi.Email,
	); err != nil {
		if err == sql.ErrNoRows {
			return affi, CustomStatus.AffiliateNotFound
		}
		return affi, err
	}
	return affi, nil
}
func AddAffiliate(affi objects.Affiliate) (objects.Affiliate, error) {
	var affiliate objects.Affiliate
	result, err := db.DBconn.Exec(
		"INSERT INTO affiliate (affiname,district,address,phoneNumber,fax,email) VALUES (?,?,?,?,?,?)",
		affi.AffiliateName,
		affi.District,
		affi.Address,
		affi.PhoneNumber,
		affi.Fax,
		affi.Email,
	)
	if err != nil {
		return affiliate, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return affiliate, err
	}

	affiliate.Id = int(id)
	affiliate.Clone(affi)

	return affiliate, nil
}

func DeleteAffiliate(id int) error {
	_, err := db.DBconn.Exec("DELETE FROM affiliate WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
