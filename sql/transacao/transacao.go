package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/go2")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id,nome) values(?,?)")
	stmt.Exec(22, "Maria")
	stmt.Exec(23, "Joao")

	_, err = stmt.Exec(4, "Pedro")

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
