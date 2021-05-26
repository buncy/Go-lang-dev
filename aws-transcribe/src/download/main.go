package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// download will download a file from aws s3 bucket and save it locally
func DownloadFromAWS(fileKey, localOutputPath, bucketName string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create an S3 client using the loaded configuration

	client := s3.NewFromConfig(cfg)

	downloader := manager.NewDownloader(client)

	// download file into the memory

	f, err := os.Create(localOutputPath)
	if err != nil {
		fmt.Printf("error creating file %v", err)
		return err
	}

	numBytesDownloaded, err := downloader.Download(context.TODO(), f, &s3.GetObjectInput{

		Bucket: aws.String(bucketName),

		Key: aws.String(fileKey),
	})
	f.Close()

	fmt.Printf("file downloaded of size %v bytes", numBytesDownloaded)

	if err != nil {
		os.Remove(localOutputPath)
		return err
	}
	return nil
}

func main() {

	bucket := flag.String("bucket", "testbucketfortranscribing", "*required field, The s3 bucket to get file from")
	object := flag.String("objectKey", "coronaNews.mp4", "*required field, The name of the file in s3 bucket")
	localPath := flag.String("localPath", "/home/karthik/Downloads/aws-transcribe/test.json", "path to save the file")
	flag.Parse()

	DownloadFromAWS(*object, *localPath, *bucket)
}
