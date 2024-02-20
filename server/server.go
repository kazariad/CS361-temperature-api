package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	port := 8080
	if len(os.Args) == 2 {
		var err error
		port, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Invalid args. Usage:\n\t%v [port]\n", filepath.Base(os.Args[0]))
		}
	} else if len(os.Args) > 2 {
		log.Fatalf("Invalid args. Usage:\n\t%v [port]\n", filepath.Base(os.Args[0]))
	}

	log.Println("Starting server...")
	http.HandleFunc("/", all)
	http.HandleFunc("/insight", insight)
	http.HandleFunc("/version", version)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}

func all(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.RemoteAddr + " -> " + request.URL.String())
	writer.WriteHeader(404)
}

var messages = []string{
	"Remember: Consistent temperature monitoring can prevent equipment failures.",
	"Coming soon: Analytics for temperature trends.",
	"Tip: Regularly calibrate your sensors to ensure accurate readings.",
}

func insight(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.RemoteAddr + " -> " + request.URL.String())
	if request.Method != http.MethodGet {
		writer.WriteHeader(404)
	} else {
		_, err := writer.Write([]byte(messages[rand.Intn(len(messages))]))
		if err != nil {
			log.Printf("/insight error: %v\n", err)
		}
	}
}

func version(writer http.ResponseWriter, request *http.Request) {
	log.Println(request.RemoteAddr + " -> " + request.URL.String())
	if request.Method != http.MethodGet {
		writer.WriteHeader(404)
	} else {
		_, err := writer.Write([]byte("v1.0.0"))
		if err != nil {
			log.Printf("/version error: %v\n", err)
		}
	}
}
