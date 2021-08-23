package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbuser     = "root"
	dbpassword = "mypassword"
	dbname     = "test"
)

type user struct {
	id   int
	name string
}

type users struct {
	users []user
}

func connection() (db *sql.DB) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@/%s", dbuser, dbpassword, dbname))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {

	fmt.Println("Starting server...")

	http.HandleFunc("/", hello)

	http.HandleFunc("/user", getUserByID)

	http.HandleFunc("/users", getUsers)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Let's Go!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db := connection()

	defer db.Close()

	row, err := db.Query("select id, name from users")

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	result := users{}

	for row.Next() {
		user := user{}
		err := row.Scan(&user.id, &user.name)
		if err != nil {
			log.Fatal(err)
		}
		result.users = append(result.users, user)
	}

	if err = row.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "users: %v", result.users)

}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	db := connection()

	defer db.Close()

	row, err := db.Query("select id, name from users where id = ?", 1)

	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()

	user := user{}
	for row.Next() {
		err := row.Scan(&user.id, &user.name)

		if err != nil {
			log.Fatal(err)
		}
	}

	if err = row.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "user: %v", user)
}
