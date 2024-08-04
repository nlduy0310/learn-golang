package db

import (
	"database/sql"
)

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

type DB struct {
	db *sql.DB
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed() error {
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, number := range data {
		if _, err := insertPhoneNumber(db.db, number); err != nil {
			return err
		}
	}
	return nil
}

func insertPhoneNumber(db *sql.DB, number string) (int, error) {
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, number).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func Migrate(driverName, dataSource string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createPhoneNumbersTable(db)
	if err != nil {
		return err
	}
	return db.Close()
}

func createPhoneNumbersTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS phone_numbers (
        id SERIAL,
        value VARCHAR(255)
    )`
	_, err := db.Exec(query)
	return err
}

func Reset(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

type Phone struct {
	Id     int
	Number string
}

// we dont need this right now
func getPhone(db *sql.DB, number string) (*Phone, error) {
	var p Phone
	err := db.QueryRow("SELECT id, value FROM phone_numbers WHERE value=$1", number).Scan(&p.Id, &p.Number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func updatePhone(db *sql.DB, p Phone) error {
	statement := "UPDATE phone_numbers SET value=$2 WHERE id=$1"
	_, err := db.Exec(statement, p.Id, p.Number)
	return err
}

func getAllPhones(db *sql.DB) ([]Phone, error) {
	// query
	rows, err := db.Query("SELECT id, value FROM phone_numbers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// iterate
	var res []Phone
	for rows.Next() {
		var p Phone
		if err := rows.Scan(&p.Id, &p.Number); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
