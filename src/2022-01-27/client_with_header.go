package main

import (
  "bytes"
  "io"
  "log"
  "net/http"
  "net/http/httputil"
  "mime/multipart"
	"os"
)

func main() {
  client := &http.Client{}
  var buffer bytes.Buffer
  writer := multipart.NewWriter(&buffer)
  fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
  if err != nil {
    panic(err)
  }
  readFile, err := os.Open("photo.jpg")
  if err != nil {
    // ファイル読み込み失敗
    panic(err)
  }
  defer readFile.Close()
  io.Copy(fileWriter, readFile)
  writer.Close()

  request, err := http.NewRequest("POST", "http://localhost:18888", &buffer)
  if err != nil {
    panic(err)
  }
	request.Header.Add("Content-Type", "image/jpeg")
  resp, err := client.Do(request)
  if err != nil {
    panic(err)
  }
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  log.Println(string(dump))
}
