package main

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"

	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"golang.org/x/net/http2"
)

type Data struct {
	AudioStream AudioEventStruct
}
type AudioEventStruct struct {
	AudioEvent Chunk
}

type Chunk struct {
	AudioChunk []byte
}

const (
	HeaderKeyLanguageCode   = "x-amzn-transcribe-language-code"  // en-US
	HeaderKeyMediaEncoding  = "x-amzn-transcribe-media-encoding" // pcm only
	HeaderKeySampleRate     = "x-amzn-transcribe-sample-rate"    // 8000, 16000 ... 48000
	HeaderKeySessionId      = "x-amzn-transcribe-session-id"     // For retrying a session. Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	HeaderKeyVocabularyName = "x-amzn-transcribe-vocabulary-name"
	HeaderKeyRequestId      = "x-amzn-request-id"
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
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)

	fmt.Fprint(w, "Hello HTTP/2")

	region := "us-east-1"

	cred := aws.Credentials{
		AccessKeyID:     "AKIA4SPD6NJGWJZTAPK2",
		SecretAccessKey: "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",
	}

	signer := v4.NewSigner()

	signTime := time.Now()

	header := http.Header{}
	header.Set(HeaderKeyLanguageCode, "en-US")
	header.Set(HeaderKeyMediaEncoding, "pcm")
	header.Set(HeaderKeySampleRate, "16000")
	header.Set("host", "transcribestreaming.region.amazonaws.com")
	//header.Set("authorization",)
	//header.Set("Content-type", "application/json")
	// header.Set("x-amz-target","com.amazonaws.transcribe.Transcribe.StartStreamTranscription")
	// header.Set("x-amz-content-sha256","STREAMING-AWS4-HMAC-SHA256-EVENTS")
	// header.Set("x-amz-date",time.Now().Format("2006-01-02T15:04:05Z"))
	// header.Set("Content-type", "application/vnd.amazon.eventstream")
	// header.Set("transfer-encoding","chunked")
	// Bi-directional streaming via a pipe.
	//pr, pw := io.Pipe()

	//Encode the data
	body := Data{
		AudioStream: AudioEventStruct{
			AudioEvent: Chunk{
				AudioChunk: []byte(""),
			},
		},
	}

	json, err := json.Marshal(body)

	requestBody := bytes.NewBuffer(json)

	if err != nil {
		fmt.Println("this is the error for err :", err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, "https://transcribestreaming.us-east-1.amazonaws.com/stream-transcription", requestBody)
	if err != nil {
		log.Printf("err: %+v", err)
		return
	}
	req.Header = header

	//payload hash of empty string

	s := ""
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	err = signer.SignHTTP(context.TODO(), cred, req, sha1_hash, "transcribe", region, signTime)
	if err != nil {
		log.Printf("problem signing headers: %+v", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("this is the error for err :", err.Error())
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("this is the error for err :", err.Error())
	}
	//spew.Dump(string(content))
	fmt.Println("\n", string(content))
}

// func signature()  {

// 	//-----Request values ----------//
// 	method := "POST"
// service := "transcribestreaming"
// host := "transcribestreaming.us-east-1.amazonaws.com"
// region := "us-east-1"
// endpoint := "https://transcribestreaming.us-east-1.amazonaws.com"
// request_parameters := ""

// //-----Create signing key---------//

// }

// func getSignatureKey(key string, dateStamp string, regionName string, serviceName string) error {

// 	kdate :=
// }

// func ComputeHmac256(data string,secret []byte) error  {
// 	key := []byte(secret)
// 	h := hmac.New(sha256.New, key)
// 	h.Write([]byte(data))
// 	return base64.StdEncoding.EncodeToString(h.Sum(nil))
// }