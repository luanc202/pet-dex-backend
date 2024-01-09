package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"pet-dex-backend/v2/api/routes"
)

func main() {
	dbc, err := sql.Open("mysql", "root:root@tcp")
	if err != nil {
		panic(err)
	}
	defer dbc.Close()
	// or := db.NewOngRepository()
	port := "3000"
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes.Routes(),
	}
	srv.ListenAndServe()
}
