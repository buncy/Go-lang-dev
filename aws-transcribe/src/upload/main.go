package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {

	bucket := flag.String("bucket", "testbucketfortranscribing", "The s3 bucket to upload to")
	filename := flag.String("filename", "", "The file to be uploaded")
	flag.Parse()

	// sess, _ := session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// })

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	client := s3.NewFromConfig(cfg)

	//uploader := s3manager.NewUploader(sess)
	uploader := manager.NewUploader(client)
	file, err := os.Open(*filename)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	key := filepath.Base(file.Name())

	// _, err = uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(*bucket),
	// 	Key:    aws.String(key),
	// 	Body:   file,
	// })

	response, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(*bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

}
