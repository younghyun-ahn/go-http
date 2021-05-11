package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// curl -x http://localhost:18888 http://github.com

	proxyUrl, err := url.Parse("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	res, err := client.Get("http://gitub.com")
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
