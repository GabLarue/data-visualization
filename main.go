package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func uploadFile(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	f, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := f.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(f.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", f.Filename))
}

func main() {
	e := echo.New()

	e.POST("/upload", uploadFile)

	e.Logger.Fatal(e.Start(":8080"))
}
