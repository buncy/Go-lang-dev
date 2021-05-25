package main

import (
	"context"
	"io"
	"log"

	// "fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
)

var (
	LanguageCodeEnUs = "en-US"
	MediaEncodingPcm = "pcm"
)

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
	sess, err := session.NewSession()

	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := transcribestreamingservice.New(sess)
	var StartStreamTranscriptionInput = transcribestreamingservice.StartStreamTranscriptionInput{
		LanguageCode:         aws.String(LanguageCodeEnUs),
		MediaEncoding:        aws.String(MediaEncodingPcm),
		MediaSampleRateHertz: aws.Int64(16000),
	}
	resp, err := client.StartStreamTranscription(&StartStreamTranscriptionInput)
	if err != nil {
		log.Fatalf("failed to start streaming, %v", err)
	}
	stream := resp.GetStream()
	defer stream.Close()

	var audio io.Reader
	// TODO Set audio to an io.Reader to stream audio bytes from.
	//audio.Read([]byte(""))
	go transcribestreamingservice.StreamAudioFromReader(context.Background(), stream.Writer, 10*1024, audio)

	for event := range stream.Events() {
		switch e := event.(type) {
		case *TranscriptEvent:
			log.Printf("got event, %v results", len(e.Transcript.Results))
			for _, res := range e.Transcript.Results {
				for _, alt := range res.Alternatives {
					log.Printf("* %s", aws.StringValue(alt.Transcript))
				}
			}
		default:
			log.Fatalf("unexpected event, %T", event)
		}
	}

	if err := stream.Err(); err != nil {
		log.Fatalf("expect no error from stream, got %v", err)
	}
}
