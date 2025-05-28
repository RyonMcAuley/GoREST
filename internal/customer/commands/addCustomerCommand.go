package customer

import "database/sql"

func CreateCustomer(db *sql.DB, c Customer) (int, error) {
	var id int
	err := db.QueryRow(`
        INSERT INTO customers (first_name, last_name, email)
        VALUES ($1, $2, $3)
        RETURNING id
    `, c.FirstName, c.LastName, c.Email).Scan(&id)

	return id, err
}
