package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "docker:docker@tcp(db:3305)/test_db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
