package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	var httpServer = http.Server{
		Addr: ":9191",
	}
	var http2Server = http2.Server{}
	_ = http2.ConfigureServer(&httpServer, &http2Server)
	//http.HandleFunc("/hello/sayHello", echoPayload)
	log.Printf("Go Backend: { HTTPVersion = 2 }; serving on https://localhost:9191/hello/sayHello")
	log.Fatal(httpServer.ListenAndServeTLS("./cert/localhost.crt", "./cert/localhost.key"))
	http.HandleFunc("/", greet)
	//http.ListenAndServe(":8080", nil)
}
