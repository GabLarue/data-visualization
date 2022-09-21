package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	nf, err := c.FormFile("file")
	check(err)

	// Initialize a session
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})
	if err != nil {
		return fmt.Errorf("failed to initialize session %v", err)
	}

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := nf.Open()
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", nf.Filename, err)
	}

	println(f)

	// // Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("bucket"),
		Key:    aws.String(nf.Filename),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to %s, %v", result.Location, err)
	}

	return c.JSON(http.StatusOK, result)
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

	// Initialize a session
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      credentials.NewStaticCredentials("test", "test", ""),
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4566"),
	})

	// Create S3 service client
	client := s3.New(sess)

	output, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("bucket"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", object.Key, object.Size)
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
