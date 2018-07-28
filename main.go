package main

import (
    "fmt"
    "flag"
    "uploader"
    "io"
    "os"
    "time"
    "log"
    "net/http"
    "encoding/json"
)

var (
  // flagPort is the open port the application listens on
  flagPort = flag.String("port", "8000", "Port to listen on")
)

var results []string

// GetHandler handles the index route
func GetHandler(w http.ResponseWriter, r *http.Request) {
  results = append(results, time.Now().Format(time.RFC3339))
  jsonBody, err := json.Marshal(results)

  if err != nil {
    http.Error(w, "Error converting results to json",
      http.StatusInternalServerError)
  }

  // fmt.Fprintf(w, "GoUploader Service is up, %q", time.Now())
  w.Write(jsonBody)
}

func init() {
  log.SetFlags(log.Lmicroseconds | log.Lshortfile)
  flag.Parse()
}


func main() {
  route := http.NewServeMux()
  route.HandleFunc("/", GetHandler)
  route.HandleFunc("/upload", PostHandler)

  log.Printf("listening on port %s", *flagPort)

  log.Fatal(http.ListenAndServe(":"+*flagPort, route))
}
