package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	defaultMime()
}

// application/octet-stream
func defaultMime() {
	// curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:18888"

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	res, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer) // FormDataContentType() = "multipart/form-data; boundary=" + writer.Boundary()
	if err != nil {
		panic(err)
	}
	log.Println("Status:", res.Status)
}

func mime() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.jpg"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("photo.jpg")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	res, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer) // FormDataContentType() = "multipart/form-data; boundary=" + writer.Boundary()
	if err != nil {
		panic(err)
	}
	log.Println("Status:", res.Status)
}