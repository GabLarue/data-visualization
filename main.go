package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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
	if result, err = S3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("bucket"),
		Key:    aws.String(nf.Filename),
		Body:   f,
	}); err != nil {
		return fmt.Errorf("failed to upload file to %s, %v", result, err)
	}

	// ADD FILE NAME + BUCKET KEY TO DB

	return c.JSON(http.StatusOK, result)
}

func getAllFiles(c echo.Context) error {
	files, _ := ioutil.ReadDir("./data")

	filesNames := []string{}

	for _, f := range files {
		filesNames = append(filesNames, f.Name())
	}

	output, err := S3.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("bucket"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%p size=%d", object.Key, object.Size)
	}

	return c.JSON(http.StatusOK, filesNames)
}

func main() {
	initAws()
	initDB()
	e := echo.New()

	row := DB.QueryRow("SELECT (file_url) from files where id = 1;")

	fmt.Println(row)

	sqlStmt := `
	INSERT INTO files (file_url)
	VALUES ('www.go.com')`

	_, err := DB.Exec(sqlStmt)
	check(err)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	e.POST("/upload", uploadFile)
	e.GET("/files", getAllFiles)

	e.Logger.Fatal(e.Start(":8080"))
}
