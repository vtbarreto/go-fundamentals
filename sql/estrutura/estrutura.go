package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:root@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "create database if not exists go")
	exec(db, "use go2")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id int auto_increment,
		nome varchar(80), 
		PRIMARY KEY (id))`)
}
