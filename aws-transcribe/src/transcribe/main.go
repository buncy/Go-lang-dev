package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

func main() {

	bucket := flag.String("bucket", "testbucketfortranscribing", "*required field, The s3 bucket to get file from")
	object := flag.String("objectKey", "coronaNews.mp4", "*required field, The name of the file in s3 bucket")
	jobname := flag.String("jobName", "autodetect", "*required field, name of the job")
	outputBucketName := flag.String("outputBucket", "testbucketfortranscribing", "The s3 bucket to get file from")
	flag.Parse()
	cfg, err := config.LoadDefaultConfig(context.TODO())

	transcriber := transcribe.NewFromConfig(cfg)

	// exit if unable to create a Transcribe session
	if transcriber == nil {
		fmt.Println("Unable to create Transcribe session\n")
		return
	} else {
		fmt.Println("Transcribe session successfully created\n")
	}

	// create a random id for the jobname

	mediafileuri := fmt.Sprintf("https://s3.amazonaws.com/%s/%s", *bucket, *object)

	identifyLanguage := true
	var mediaformat types.MediaFormat = "mp4"
	media := &types.Media{
		MediaFileUri: &mediafileuri,
	}

	transcriber.StartTranscriptionJob(context.TODO(), &transcribe.StartTranscriptionJobInput{
		TranscriptionJobName: jobname,
		Media:                media,
		MediaFormat:          mediaformat,
		OutputBucketName:     outputBucketName,
		IdentifyLanguage:     &identifyLanguage,
	})

	if err != nil {
		fmt.Println(err)
	}

}
