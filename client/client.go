package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalf("Invalid args. Usage:\n\t%v URL\n", filepath.Base(os.Args[0]))
	}

	url := os.Args[1]
	log.Printf("Sending GET request to %v...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println("Response status: " + resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response body: \"%v\"\n", string(body))
}
