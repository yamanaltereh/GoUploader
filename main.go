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

// PostHandler converts post request body to string
func PostHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    r.ParseMultipartForm(32 << 20)

    file, handler, err := r.FormFile("file")
    file_name := handler.Filename

    if err != nil {
      fmt.Println(err)
      return
    }

    defer file.Close()

    fmt.Fprintf(w, "%v", handler.Header)
    file_path := "./tmp/" + file_name
    f, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer f.Close()
    io.Copy(f, file)

    // response := map[string]string{"filename": file_name}
    // json_response, _ := json.Marshal(response)

    // w.Write(json_response)

    var file_url string
    uploader.Upload(file_path, file_name, file_url)

    jsonBody, err := json.Marshal(file_url)
    w.Write(jsonBody)
  } else {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
  }
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
