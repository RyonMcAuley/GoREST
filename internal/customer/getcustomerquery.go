package customer

import "database/sql"

func GetCustomerByID(db *sql.DB, id int) (*Customer, error) {
	var c Customer
	err := db.QueryRow(`
		SELECT id, first_name, last_name, email
		FROM customers
		WHERE id = $1
	`, id).Scan(&c.ID, &c.FirstName, &c.LastName, &c.Email)

	if err != nil {
		return nil, err
	}

	return &c, err
}
