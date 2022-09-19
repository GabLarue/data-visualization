package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func uploadFile(c echo.Context) error {
	f, err := c.FormFile("file")
	check(err)

	src, err := f.Open()
	check(err)
	defer src.Close()

	dst, err := os.Create("data/" + f.Filename)
	check(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	check(err)

	return c.JSON(http.StatusOK, f.Filename)
}

// func getFileByID(c echo.Context) error {
// 	f, err := os.ReadFile("data")
// 	check(err)

// 	return c.JSON(http.StatusOK, f)
// }

func getAllFiles(c echo.Context) error {
	files, _ := ioutil.ReadDir("./data")

	filesNames := []string{}

	for _, f := range files {
		filesNames = append(filesNames, f.Name())
	}

	return c.JSON(http.StatusOK, filesNames)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	//to start db container => docker-compose -f docker-compose-postgres.yml up
	e := echo.New()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	check(err)
	defer db.Close()

	db.Ping()
	fmt.Println("Successfully ping db")

	row := db.QueryRow("SELECT (file_url) from files where id = 1;")

	fmt.Println(row)

	sqlStmt := `
	INSERT INTO files (file_url)
	VALUES ('www.go.com')`

	_, err = db.Exec(sqlStmt)
	check(err)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.POST("/upload", uploadFile)
	e.GET("/files", getAllFiles)

	e.Logger.Fatal(e.Start(":8080"))
}
