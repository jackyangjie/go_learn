package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	//db, err = sql.Open("postgres", "postgres://postgres:postgres@192.168.1.155:5432/fmb2020?sslmode=disable")
	db, err = sql.Open("postgres", "host=192.168.1.155 port=5432 user=postgres password=postgres dbname=fmb2020 sslmode=disable")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func SqlSelect() {

	rows, err := db.Query("select id,username from acl_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		var id string
		var username string
		rows.Scan(&id, &username)
		fmt.Println("id=", id, "userName=", username)
	}
}
