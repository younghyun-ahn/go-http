package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	// curl -H "Content-Type=image/jpeg" -d "@image.jpeg" http://localhost:18888

	values := url.Values{
		"test": {"value"},
	}
	reader := strings.NewReader(values.Encode())

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:18888", reader)
	if err != nil {
		panic(err)
	}

	// 헤더 직접 설정
	req.Header.Add("Content-Type", "image/jpeg")
	// 헬퍼메서드 이용
	req.SetBasicAuth("유저명", "패스워드")                           // --basic -u 유저명:패스워드
	req.AddCookie(&http.Cookie{Name: "test", Value: "value"}) // -c

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
