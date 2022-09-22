package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
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

type FilesResponse struct {
	Name string `db:"file_name" json:"name"`
	Key  string `db:"s3_key" json:"key"`
}

func getAllFiles(c echo.Context) error {
	var err error
	var stmt *sql.Stmt

	if stmt, err = DB.Prepare("SELECT file_name, s3_key FROM files"); err != nil {
		return fmt.Errorf("failed to prepare statement for file upload query, %v", err)
	}
	defer stmt.Close()

	var results *sql.Rows
	results, err = stmt.Query()
	var response []FilesResponse

	for results.Next() {
		var file_name string
		var s3_key string
		if err := results.Scan(&file_name, &s3_key); err != nil {
			log.Fatal(err)
		}
		response = append(response, FilesResponse{Name: file_name, Key: s3_key})
	}

	return c.JSON(http.StatusOK, response)
}

func getFileByID(c echo.Context) error {
	var err error
	var id string
	id = c.Param("id")

	var result *s3.GetObjectOutput
	if result, err = S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("bucket"),
		Key:    aws.String(id),
	}); err != nil {
		return fmt.Errorf("failed to upload file to %s, %v", result, err)
	}

	size := result.ContentLength
	buffer := make([]byte, *size)

	var bbuffer bytes.Buffer
	for true {
		num, rerr := result.Body.Read(buffer)
		if num > 0 {
			bbuffer.Write(buffer[:num])
		} else if rerr == io.EOF || rerr != nil {
			break
		}
	}

	return c.JSON(http.StatusOK, bbuffer.String())
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
