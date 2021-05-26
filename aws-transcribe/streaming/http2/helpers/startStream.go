package helpers

import (
    "crypto/tls"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/aws/external"
    "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
    "golang.org/x/net/http2"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

const (
    HeaderKeyLanguageCode   = "x-amzn-transcribe-language-code"  // en-US
    HeaderKeyMediaEncoding  = "x-amzn-transcribe-media-encoding" // pcm only
    HeaderKeySampleRate     = "x-amzn-transcribe-sample-rate"    // 8000, 16000 ... 48000
    HeaderKeySessionId      = "x-amzn-transcribe-session-id"     // For retrying a session. Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
    HeaderKeyVocabularyName = "x-amzn-transcribe-vocabulary-name"
    HeaderKeyRequestId = "x-amzn-request-id"
)

...

region := "eu-west-1"

cfg, err := external.LoadDefaultAWSConfig(aws.Config{
    Region: region,
})
if err != nil {
    log.Printf("could not load default AWS config: %v", err)
    return
}

signer := v4.NewSigner(cfg.Credentials)

transport := &http2.Transport{
    TLSClientConfig: &tls.Config{
        // allow insecure just for debugging
        InsecureSkipVerify: true,
    },
}
client := &http.Client{
    Transport: transport,
}

signTime := time.Now()

header := http.Header{}
header.Set(HeaderKeyLanguageCode, "en-US")
header.Set(HeaderKeyMediaEncoding, "pcm")
header.Set(HeaderKeySampleRate, "16000")
header.Set("Content-type", "application/json")

// Bi-directional streaming via a pipe.
pr, pw := io.Pipe()

req, err := http.NewRequest(http.MethodPost, "https://transcribestreaming.eu-west-1.amazonaws.com/stream-transcription", ioutil.NopCloser(pr))
if err != nil {
    log.Printf("err: %+v", err)
    return
}
req.Header = header

_, err = signer.Sign(req, nil, "transcribe", region, signTime)
if err != nil {
    log.Printf("problem signing headers: %+v", err)
    return
}

// This freezes and ends after 5 minutes with "unexpected EOF".
res, err := client.Do(req)
...