package main

import (
	"database/sql"
	"fmt"
	"my_packages/db"

	_ "github.com/lib/pq"
)

func main() {
	pgsql := fmt.Sprintf("host=localhost port=5432 user=postgres password=satheesh dbname=postgres sslmode=disable")
	db1, err := sql.Open("postgres", pgsql)
	if err != nil {
		fmt.Println("error while connecting db", err)
	}
	defer db1.Close()
	err = db1.Ping()
	if err != nil {
		fmt.Println("error while pinging db", err)
	}
	fmt.Println("connected...")
	_ = db.InsertQuery(db1)
	_ = db.SelectQuery(db1)
	_ = db.UpdateQuery(db1)
	_ = db.DeleteQuery(db1)

}
