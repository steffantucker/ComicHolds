package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db struct to interact with db
type Db struct {
	*sql.DB
}

// New creates a new db connection
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// check connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a formated connection string
func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf(
		"host= %s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	)
}

// User struct
type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

// GetUsersByName gets users
func (d *Db) GetUsersByName(name string) []User {
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUsersByName preperation err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName query err: ", err)
	}

	var r User
	users := []User{}
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}
	return users
}
