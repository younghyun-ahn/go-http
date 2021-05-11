package main

import (
	"log"
	"net/http"
)

func main() {
	// curl --head http://localhost:18888

	res, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	log.Println("Status:", res.Status)
	log.Println("Headers:", res.Status)

}
