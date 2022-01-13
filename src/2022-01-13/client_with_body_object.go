package main

import (
  "log"
  "net/http"
  "net/http/httputil"
  "strings"
	"net/url"
)

func main() {
  client := &http.Client{}

	values := url.Values{
		"message": {"hello world."},
	}
  reader := strings.NewReader(values.Encode())
	// 第3引数はio.Reader インターフェイス型
  request, err := http.NewRequest("POST", "http://localhost:18888", reader)
  if err != nil {
    panic(err)
  }
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