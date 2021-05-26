package main

import (
	"crypto/hmac"
	"crypto/sha256"

	// "encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	//	"net"
	"net/http"
	"time"
	//"golang.org/x/net/http2/h2c"
)

func main() {

	var httpServer = http.Server{
		Addr: ":3002",
	}
	http.HandleFunc("/", home)
	// var http2Server = http2.Server{}
	// _ = http2.ConfigureServer(&httpServer, &http2Server)

	// log.Printf("Go Backend: { HTTPVersion = 2 }; serving on https://localhost:3002/")

	// if err := httpServer.ListenAndServeTLS("certs/localhost.crt", "certs/localhost.key"); err != nil {
	// 	fmt.Println("\n connection error:", err.Error())
	// }
	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Println("\n connection error:", err.Error())
	}
}

const (
	HeaderKeyLanguageCode   = "x-amzn-transcribe-language-code:"  // en-US
	HeaderKeyMediaEncoding  = "x-amzn-transcribe-media-encoding:" // pcm only
	HeaderKeySampleRate     = "x-amzn-transcribe-sample-rate:"    // 8000, 16000 ... 48000
	HeaderKeySessionId      = "x-amzn-transcribe-session-id:"     // For retrying a session. Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	HeaderKeyVocabularyName = "x-amzn-transcribe-vocabulary-name:"
	HeaderKeyRequestId      = "x-amzn-request-id:"
	host                    = "transcribestreaming.us-east-1.amazonaws.com"
)

func home(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}

	// Create a pool with the server certificate since it is not signed
	// by a known CA
	// caCert, err := ioutil.ReadFile("./cert/localhost.crt")
	// if err != nil {
	// 	log.Fatalf("Reading server certificate: %s", err)
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// // Create TLS configuration with the certificate of the server
	// tlsConfig := &tls.Config{
	// 	RootCAs: caCertPool,
	// }

	// // Use the proper transport in the client
	// client.Transport = &http2.Transport{
	// 	TLSClientConfig: tlsConfig,
	// }

	//-------Echo the request --------------//

	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)

	// client := http.Client{
	// 	Transport: &http2.Transport{
	// 		// So http2.Transport doesn't complain the URL scheme isn't 'https'
	// 		AllowHTTP: true,
	// 		// Pretend we are dialing a TLS endpoint.
	// 		// Note, we ignore the passed tls.Config
	// 		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
	// 			return net.Dial(network, addr)
	// 		},
	// 	},
	// }

	scheduledTimeObject := time.Now()
	//loc, _ := time.LoadLocation("America/Los_Angeles")
	//scheduledOnInPT := scheduledTimeObject.In(loc)

	//date := scheduledOnInPT.Format("20060102T150405Z")
	date := scheduledTimeObject.Format("20060102T150405Z")
	date_stamp := time.Now().Format("20060102")
	region := "us-east-1"
	service := "transcribestreaming"

	AccessKeyID := "AKIA4SPD6NJGWJZTAPK2"
	SecretAccessKey := "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"

	canonical_headers := "host:" + host + "\n" + "x-amz-date:" + date + "\n" + HeaderKeyLanguageCode + "en-US" + "\n" + HeaderKeyMediaEncoding + "pcm" + HeaderKeySampleRate + "16000" + "\n"

	//-------signing headers ------------//

	signed_headers := "host;x-amz-date;x-amzn-transcribe-language-code;x-amzn-transcribe-media-encoding;x-amzn-transcribe-sample-rate"

	//--------Canonical request------------//

	canonicalRequest :=
		"POST" + "\n" +
			"/stream-transcription" + "\n" +
			canonical_headers + "\n" +
			signed_headers + "\n" +
			hexEncodedSHA256("")

	//-------Create string to sign-------------//
	algorithm := "AWS4-HMAC-SHA256"
	credential_scope := date_stamp + "/" + region + "/" + service + "/" + "aws4_request"
	string_to_sign := algorithm + "\n" + date + "\n" + credential_scope + "\n" + hexEncodedSHA256(canonicalRequest)

	//------signing key---------------------//

	signing_key := getSignatureKey(SecretAccessKey, date_stamp, region, service)

	//---------- Sign the string_to_sign using the signing_key-------------//

	signature := ComputeHmac256(signing_key, []byte(string_to_sign))

	//---------add signing information------------------//

	authorization_header := algorithm + " " + "Credential=" + AccessKeyID + "/" + credential_scope + ", " + "SignedHeaders=" + signed_headers + ", " + "Signature=" + signature

	//--------------add headers to request-------------//

	req, err := http.NewRequest("POST", "https://transcribestreaming.us-east-1.amazonaws.com/stream-transcription", nil)

	if err != nil {
		fmt.Println("this is the error for creating request :", err.Error())
	}
	req.Header.Add("x-amz-date", date)
	req.Header.Add("x-amzn-transcribe-language-code", "en-US")
	req.Header.Add("x-amzn-transcribe-media-encoding", "pcm")
	req.Header.Add("x-amzn-transcribe-sample-rate", "16000")
	req.Header.Add("host", "transcribestreaming.us-east-1.amazonaws.com")
	req.Header.Add("Authorization", authorization_header)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("this is the error for response :", err.Error())
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("this is the error for err :", err.Error())
	}
	//spew.Dump(string(content))
	fmt.Println("This is the response :=\n", string(content))

}

func ComputeHmac256(data string, secret []byte) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, secret)

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}

func hexEncodedSHA256(data string) string {

	//hex encoded payload hash of empty string

	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func getSignatureKey(key string, dateStamp string, region string, serviceName string) string {

	kdate := ComputeHmac256(("AWS4" + key), []byte(dateStamp))
	kRegion := ComputeHmac256(kdate, []byte(region))
	kService := ComputeHmac256(kRegion, []byte(serviceName))
	kSigning := ComputeHmac256(kService, []byte("aws4_request"))
	return kSigning
}

// func test(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "hello world")
// }
