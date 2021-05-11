package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	// curl -X DELETE http://localhost:18888

	// http.Client 는 GET, HEAD, POST 만 사용 가능. 이외 메서드는 http.Request 로 사용
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
