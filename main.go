package main

import (
    "fmt"
    "uploader"
    "time"
    "log"
    "net/http"
)

func main() {
  fmt.Println("start main")
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "GoUploader Service is up, %q", time.Now())
  })

  log.Fatal(http.ListenAndServe(":8080", nil))
  uploader.Upload("/Users/yaman/Desktop/P_20170919_094831.jpg")
}
