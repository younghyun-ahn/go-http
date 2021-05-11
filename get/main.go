package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// curl http://localhost:18888
	res, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
