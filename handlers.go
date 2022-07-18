package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pelletier/go-toml/v2"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func GetUserBody(id string) *User {

	resp, err := http.Get("https://gorest.co.in/public/v2/users/" + id)
	if err != nil {
		log.Print(http.NotFound)
	}
	out := make([]byte, 1024)

	body, err := resp.Body.Read(out)
	if err != io.EOF {
		fmt.Println(err.Error())
	}
	var k *User

	err = json.Unmarshal(out[:body], &k)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(k)
	return k
}

func ReadCSVintoFilm(path string) *Film {

	csvFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	reader := csv.NewReader(csvFile)
	reader.Comma = '|'
	var film *Film
	line, _ := reader.Read()

	userid, _ := strconv.Atoi(line[1])

	film = &Film{
		Title:   line[0],
		UserID:  userid,
		Content: line[2],
		Expires: line[3],
	}
	return film

}

func Create3(e echo.Context) error {

	id := e.QueryParam("id")

	k := GetUserBody(id)

	film := ReadCSVintoFilm("./cmd/web/csv/data.csv")

	db, err := openDB()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	if k.ID == film.UserID {
		if k.Status == "active" {
			stmt := `INSERT INTO film (titlename,userid, content, created, expires)
VALUES(?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

			result, err := db.Exec(stmt, &film.Title, &film.UserID, &film.Content, &film.Expires)
			if err != nil {
				e.Response().Write([]byte(err.Error() + ".\n" + "Film già esistente."))
				return err
			}
			id, err := result.LastInsertId()
			if err != nil {
				log.Print(http.NotFound)
			}
			return e.JSON(200, id)

		} else if k.Status == "inactive" {
			stmt := `UPDATE film set expires = current_time WHERE userid = ?`

			result, err := db.Exec(stmt, &film.UserID)
			if err != nil {
				e.Response().Write([]byte(err.Error() + ".\n" + "Errore."))
				return err
			}
			id, err := result.RowsAffected()
			if err != nil {
				log.Print(http.NotFound)
			}
			return e.JSON(200, id)

		}

	}
	return e.JSON(500, nil)
}

func Create2(e echo.Context) error {
	id := e.QueryParam("id")

	resp, err := http.Get("https://gorest.co.in/public/v2/users/" + id)
	if err != nil {
		log.Print(http.NotFound)
	}
	out := make([]byte, 1024)

	body, err := resp.Body.Read(out)
	if err != io.EOF {
		fmt.Println(err.Error())
	}
	var k *User

	err = json.Unmarshal(out[:body], &k)
	if err != nil {
		log.Print(err)
	}
	csvFile, err := os.Open("./cmd/web/csv/data.csv")

	if err != nil {
		fmt.Println(err.Error())
	}

	reader := csv.NewReader(csvFile)
	reader.Comma = '|'
	var film *Film
	line, _ := reader.Read()

	userid, _ := strconv.Atoi(line[1])

	film = &Film{
		Title:   line[0],
		UserID:  userid,
		Content: line[2],
		Expires: line[3],
	}

	db, err := openDB()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	if k.ID == film.UserID {
		if k.Status == "active" {

			stmt := `INSERT INTO film (titlename,userid, content, created, expires) VALUES(?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

			result, err := db.Exec(stmt, &film.Title, &film.UserID, &film.Content, &film.Expires)
			if err != nil {
				e.Response().Write([]byte(err.Error() + ".\n" + "Film già esistente."))
				return err
			}
			id, err := result.LastInsertId()
			if err != nil {
				log.Print(http.NotFound)
			}
			return e.JSON(200, id)

		} else if k.Status == "inactive" {
			stmt := `UPDATE film set expires = current_time WHERE userid = ?`

			result, err := db.Exec(stmt, &film.UserID)
			if err != nil {
				e.Response().Write([]byte(err.Error() + ".\n" + "Errore."))
				return err
			}
			id, err := result.RowsAffected()
			if err != nil {
				log.Print(http.NotFound)
			}
			return e.JSON(200, id)

		}

	}
	return e.JSON(500, nil)
}

func Get(e echo.Context) error {

	db, err := openDB()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	s := &Film{}

	id := e.QueryParam("id")

	stmt := `SELECT id, titlename, content, created, expires FROM film WHERE expires > UTC_TIMESTAMP() AND id = ?`

	_ = db.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	filmObj := &Film{ID: s.ID, Title: s.Title, Content: s.Content, Created: s.Created, Expires: s.Expires}

	return e.JSON(200, filmObj)
}

func Create(e echo.Context) error {

	//csvFile, err := os.Open("./cmd/web/csv/data.csv")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//reader := csv.NewReader(csvFile)
	//reader.Comma = '|'
	//var body *models.Film
	//
	//line, _ := reader.Read()
	//
	//body = &models.Film{
	//	Title:   line[0],
	//	Content: line[1],
	//	Expires: line[2],
	//}

	//for {
	//	line, error := reader.Read()
	//	if error == io.EOF {
	//		break
	//	} else if error != nil {
	//		log.Fatal(error)
	//	}
	//	body = &models.Film{
	//		Title:   line[0],
	//		Content: line[1],
	//		Expires: line[2],
	//	}
	//}

	//b, err := toml.Marshal(body)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var cfg models.Film
	//
	//err = toml.Unmarshal(b, &cfg)
	//if err != nil {
	//	panic(err)
	//}

	db, err := openDB()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	out := make([]byte, 1024)

	bodyLen, err := e.Request().Body.Read(out)
	if err != io.EOF {
		log.Print(http.StatusInternalServerError)
	}
	var k *Film

	err = toml.Unmarshal(out[:bodyLen], &k)
	fmt.Println(k)

	if err != nil {
		log.Print(http.StatusInternalServerError)
	}

	stmt := `INSERT INTO film (titlename, content, created, expires)
VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := db.Exec(stmt, &k.Title, &k.Content, &k.Expires)
	if err != nil {
		e.Response().Write([]byte(err.Error() + ".\n" + "Film già esistente."))
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Print(http.NotFound)

	}

	return e.JSON(200, id)
}

func Delete(e echo.Context) error {

	db, err := openDB()
	if err != nil {
		log.Fatal(http.StatusInternalServerError)
	}

	id := e.QueryParam("id")

	result, err := db.Exec("delete from film where id = ?", id)
	if err != nil {
		log.Print(http.NotFound)
		return err
	}
	id2, err := result.RowsAffected()
	if err != nil {
		log.Print(http.NoBody)
	}
	return e.JSON(200, id2)

}
