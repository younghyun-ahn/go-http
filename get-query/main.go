package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// curl -G --data-urlencode "query=hello world" http://localhost:18888
	// -G <-> --get, --data <-> --data-urlencode

	values := url.Values{
		"query": {"hello world"},
	}

	res, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
