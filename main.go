// main.go

package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"adv_prog_5_6/auth"
	"adv_prog_5_6/filtering"
)

func main() {
	// Database initialization
	db, err := sql.Open("postgres", "postgres://adv_prog_user:SqD3b8CjbxsWAm6v7zr9bOb5Chqqh2bp@dpg-cni09m821fec73cohibg-a/adv_prog")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize authentication package
	auth.Init(db)

	// Initialize filtering package
	filtering.Init(db)

	createUsersTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE,
			email VARCHAR(255) UNIQUE,
			password VARCHAR(255),
			firstname VARCHAR(255),
			lastname VARCHAR(255),
			age INTEGER,
			role VARCHAR(50)
		);
	`

	_, err = db.Exec(createUsersTableSQL)
	if err != nil {
		panic(err)
	}

	createBarbersTableSQL := `
		CREATE TABLE IF NOT EXISTS barbers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			basic_info TEXT,
			price INTEGER,
			experience INTEGER,
			status TEXT,
			image_path TEXT
		);
	`

	_, err = db.Exec(createBarbersTableSQL)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", auth.IndexHandler).Methods("GET")
	r.HandleFunc("/register", auth.RegisterHandler).Methods("GET", "POST")
	r.HandleFunc("/login", auth.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/index", auth.IndexHandler).Methods("GET")
	r.HandleFunc("/barbers", filtering.FilteredBarbersHandler).Methods("GET")
	r.HandleFunc("/admin", adminHandler).Methods("GET")
	r.HandleFunc("/user", userHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)

}
func adminHandler(w http.ResponseWriter, r *http.Request) {
	// Serve user.html here
	http.ServeFile(w, r, "admin.html")
}
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Serve user.html here
	http.ServeFile(w, r, "user.html")
}
