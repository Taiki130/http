package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("Status:", resp.Status)
	log.Println("Headers", resp.Header)
	log.Println("Content-Length", resp.Header["Content-Length"])
}
