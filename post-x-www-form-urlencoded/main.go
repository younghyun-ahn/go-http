package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	// curl -d test=value http://localhost:18888

	values := url.Values{
		"test": {"value"},
	}

	res, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", res.Status)
}
