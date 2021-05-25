package main

import (
	"crypto/tls"
	"fmt"

	//"io"

	"net/http"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	server := http.Server{
		Addr:    ":3001",
		Handler: mux,
		TLSConfig: &tls.Config{
			NextProtos: []string{"h2", "http/1.1"},
		},
	}

	fmt.Printf("Server listening on %s", server.Addr)
	go http.ListenAndServe(":3000", mux)
	if err := server.ListenAndServeTLS("certs/localhost.crt", "certs/localhost.key"); err != nil {
		fmt.Println("\n connection error:", err.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello HTTP/2")

	transcribeHTTPsigner := v4.NewSigner()

}
