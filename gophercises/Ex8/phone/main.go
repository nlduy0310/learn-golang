package main

import (
	"database/sql"
	"fmt"
	"regexp"

	phonedb "phone-normalizer/db"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "password"
	db_name  = "dev"
)

func main() {
	// first connection to reset db
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	must(phonedb.Reset("postgres", connStr, db_name))

	connStr = fmt.Sprintf("%s dbname=%s", connStr, db_name)
	must(phonedb.Migrate("postgres", connStr))

	db, err := phonedb.Open("postgres", connStr)
	must(err)
	defer db.Close()

	must(db.Seed())

	// number, err := getPhoneNumber(db, id)
	// must(err)
	// fmt.Println("Number is: ", number)

	// phones, err := getAllPhones(db)
	// must(err)
	// for _, p := range phones {
	// 	normalizedNumber := normalize(p.number)
	// 	fmt.Printf("Working on %+v\n", p)
	// 	if normalizedNumber != p.number {
	// 		existing, err := getPhone(db, normalizedNumber)
	// 		must(err)
	// 		if existing != nil {
	// 			// delete this number
	// 			must(deletePhone(db, p.id))
	// 		} else {
	// 			// update this number
	// 			p.number = normalizedNumber
	// 			must(updatePhone(db, p))
	// 		}
	// 	} else {
	// 		fmt.Println("No changes required")
	// 	}
	// }
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func deletePhone(db *sql.DB, id int) error {
	statement := "DELETE FROM phone_numbers WHERE id=$1"
	_, err := db.Exec(statement, id)
	return err
}

func getPhoneNumber(db *sql.DB, id int) (string, error) {
	var number string
	err := db.QueryRow(`SELECT value FROM phone_numbers WHERE id=$1`, id).Scan(&number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")
}
