package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", f.Filename))
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

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.POST("/upload", uploadFile)
	e.GET("/files", getAllFiles)

	e.Logger.Fatal(e.Start(":4000"))
}
