package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	ID   int
	Nome string
}

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "fora")
	}

}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:root@/go2")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@/go2")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var u []Usuario
	rows, _ := db.Query("select id, nome from usuarios")
	defer rows.Close()

	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		u = append(u, usuario)
	}

	w.Header().Set("Content-Type", "application/json")
	json, _ := json.Marshal(u)
	fmt.Fprint(w, string(json))

}

func main() {
	http.HandleFunc("/usuarios/", UsuarioHandler)
	log.Println("exec")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
