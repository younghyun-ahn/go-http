package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	postContent()
	postString()
}

func postContent() {
	// curl -T main.go -H "Content-Type: text/plain" http://localhost:18888

	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}

	res, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", res.Status)
}

func postString() {
	reader := strings.NewReader("텍스트")
	res, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", res.Status)
}