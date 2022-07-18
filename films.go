package main

import (
	"database/sql"
	_ "github.com/labstack/echo/v4"
)

type FilmModel struct {
	DB *sql.DB
}
