package main

import (
   "fmt"
   "log"
   "net/http"
   "net/http/httputil"
   "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    dump, err := httputil.DumpRequest(r, true)
    if err != nil {
        http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
        return
    }
    fmt.Println(string(dump))
    fmt.Println("*--*--*--*--*--*--*--*")

    // 重い処理シミュレート
    time.Sleep(5 * time.Second)
    fmt.Fprintf(w, "<html><body>Good evening!</body></html>\n")
}

func main() {
   var httpServer http.Server
   http.HandleFunc("/", handler)
   log.Println("start http listening :18888")
   log.Println("http://127.0.0.1:18888")
   fmt.Println("+-+-+-+-+-+-+-+-+-+-")
   httpServer.Addr = ":18888"
   log.Println(httpServer.ListenAndServe())
}
