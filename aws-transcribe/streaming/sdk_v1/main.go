package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	// "fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"golang.org/x/net/http2"
)

type HTTPClientSettings struct {
	Connect          time.Duration
	ConnKeepAlive    time.Duration
	ExpectContinue   time.Duration
	IdleConn         time.Duration
	MaxAllIdleConns  int
	MaxHostIdleConns int
	ResponseHeader   time.Duration
	TLSHandshake     time.Duration
}

func NewHTTPClientWithSettings(httpSettings HTTPClientSettings) (*http.Client, error) {

	var client http.Client
	tr := &http.Transport{
		ResponseHeaderTimeout: httpSettings.ResponseHeader,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: httpSettings.ConnKeepAlive,
			DualStack: true,
			Timeout:   httpSettings.Connect,
		}).DialContext,
		MaxIdleConns:          httpSettings.MaxAllIdleConns,
		IdleConnTimeout:       httpSettings.IdleConn,
		TLSHandshakeTimeout:   httpSettings.TLSHandshake,
		MaxIdleConnsPerHost:   httpSettings.MaxHostIdleConns,
		ExpectContinueTimeout: httpSettings.ExpectContinue,
	}

	// So client makes HTTP/2 requests
	err := http2.ConfigureTransport(tr)

	if err != nil {
		fmt.Println("error creating client", err.Error())
		return &client, err
	}

	return &http.Client{
		Transport: tr,
	}, nil
}

var (
	LanguageCodeEnUs = "en-US"
	MediaEncodingPcm = "pcm"
)

func main() {
	var httpServer = http.Server{
		Addr: ":9191",
	}

	var http2Server = http2.Server{}
	_ = http2.ConfigureServer(&httpServer, &http2Server)

	http.HandleFunc("/", home)
	log.Printf("Go Backend: { HTTPVersion = 2 }; serving on https://localhost:9191/")

	if err := httpServer.ListenAndServeTLS("certs/localhost.crt", "certs/localhost.key"); err != nil {
		fmt.Println("\n connection error:", err.Error())
	}

	// http.HandleFunc("/", hello)
	// http.ListenAndServe(":3001", nil)
	fmt.Println("server  strted ")
}

func home(w http.ResponseWriter, req *http.Request) {

	httpClient, err := NewHTTPClientWithSettings(HTTPClientSettings{
		Connect:          5 * time.Second,
		ExpectContinue:   1 * time.Second,
		IdleConn:         90 * time.Second,
		ConnKeepAlive:    30 * time.Second,
		MaxAllIdleConns:  100,
		MaxHostIdleConns: 10,
		ResponseHeader:   5 * time.Second,
		TLSHandshake:     5 * time.Second,
	})
	if err != nil {
		fmt.Println("Got an error creating custom HTTP client:")
		fmt.Println(err)
		return
	}

	fmt.Println("hello strted ")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:     aws.String("us-east-1"),
		HTTPClient: httpClient,
	}))

	if err != nil {
		log.Fatalf("failed to load SDK configuration, %v", err)
	}

	client := transcribestreamingservice.New(sess, aws.NewConfig().WithHTTPClient(httpClient))
	var StartStreamTranscriptionInput = transcribestreamingservice.StartStreamTranscriptionInput{
		LanguageCode:         aws.String(LanguageCodeEnUs),
		MediaEncoding:        aws.String(MediaEncodingPcm),
		MediaSampleRateHertz: aws.Int64(16000),
	}

	//var audio io.Reader
	// TODO Set audio to an io.Reader to stream audio bytes from.
	//audio.Read([]byte(""))
	var reader *bytes.Reader
	// var file io.Reader
	// var err error
	localPath := "./test.mp3"
	file, err := os.Open(localPath)
	if err == nil {
		defer file.Close()
		//defer os.Remove(localPath)
		file.Seek(0, 0)
		buffer := make([]byte, 10*1024)

		_, err := file.Read(buffer)
		if err != nil {
			fmt.Println(err.Error())
		}
		if err == nil {
			fmt.Println(buffer, "buffer")
		}

		reader = bytes.NewReader(buffer)
		if reader != nil {
			fmt.Println("Success")
		}
	}
	if err != nil {
		fmt.Println(err.Error(), "error")
	}

	resp, err := client.StartStreamTranscription(&StartStreamTranscriptionInput)
	if err != nil {
		log.Fatalf("failed to start streaming, %v", err)
	}
	stream := resp.GetStream()
	//defer stream.Close()

	transErr := transcribestreamingservice.StreamAudioFromReader(context.Background(), stream.Writer, 10*1024, reader)

	if transErr != nil {
		fmt.Println("Error in StreamAudioFromReader ", transErr.Error())
	}

	if len(stream.Events()) > 0 {
		for event := range stream.Events() {
			switch e := event.(type) {
			case *transcribestreamingservice.TranscriptEvent:
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
	} else {
		fmt.Println("No events found")
	}

	if err := stream.Err(); err != nil {
		log.Fatalf("expect no error from stream, got %v", err)
	}
}
