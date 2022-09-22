package main

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func uploadFile(c echo.Context) error {
	var nf *multipart.FileHeader
	var err error

	if nf, err = c.FormFile("file"); err != nil {
		return fmt.Errorf("failed to return the multipart form file for the provider name")
	}

	var f multipart.File
	if f, err = nf.Open(); err != nil {
		return fmt.Errorf("failed to open file %q, %v", nf.Filename, err)
	}

	var result *s3.PutObjectOutput
	key := uuid.New()
	if result, err = S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("bucket"),
		Key:    aws.String(key.String()),
		Body:   f,
	}); err != nil {
		return fmt.Errorf("failed to upload file to %s, %v", result, err)
	}

	if _, err = DB.Exec("INSERT INTO files(file_name, s3_key) VALUES($1,$2)", nf.Filename, key.String()); err != nil {
		return fmt.Errorf("failed to execute file upload query, %v", err)
	}

	return c.JSON(http.StatusOK, nf.Filename)
}

func getAllFiles(c echo.Context) error {
	var err error
	var stmt *sql.Stmt

	if stmt, err = DB.Prepare("SELECT file_name FROM files"); err != nil {
		return fmt.Errorf("failed to prepare statement for file upload query, %v", err)
	}
	defer stmt.Close()

	var results *sql.Rows
	results, err = stmt.Query()
	var response []string

	for results.Next() {
		var file_name string
		if err := results.Scan(&file_name); err != nil {
			log.Fatal(err)
		}
		response = append(response, file_name)
	}

	return c.JSON(http.StatusOK, response)
}

func getFileByID(c echo.Context) error {
	var id string
	id = c.Param("id")

	return c.JSON(http.StatusOK, id)
}

func main() {
	initAws()
	initDB()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.POST("/upload", uploadFile)
	e.GET("/files", getAllFiles)
	e.GET("/files/:id", getFileByID)

	e.Logger.Fatal(e.Start(":8080"))
}
