package main

import (
    "fmt"
    "uploader"
    "time"
    "log"
    "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "GoUploader Service is up, %q", time.Now())
  })

  log.Fatal(http.ListenAndServe(":8080", nil))

  file_path := ""
  uploader.Upload(file_path)
}
