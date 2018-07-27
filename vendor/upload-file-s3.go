/*
  https://www.youtube.com/watch?v=iOGIKG3EptI
  https://github.com/awslabs/aws-go-wordfreq-sample/blob/master/cmd/uploads3/main.go
  
  https://docs.aws.amazon.com/sdk-for-go/api/aws/
  - first configure your aws credentials run: aws configure
  - go get -u github.com/aws/aws-sdk-go/aws
  - login to UI web aws s3 interface
  - go to S3 service
  - create a Bucket called com.example in the desired region (I used Oregon us-west-2)
  - run:   go run main.go com.example fileToUpload
*/

package main

import (
  "fmt"
  "os"
  "path/filepath"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
  if len(os.Args) != 3 {
    fmt.Printf("usage: %s <bucket> <filename>\n", filepath.Base(os.Args[0]))
    os.Exit(1)
  }

  bucket := os.Args[1]
  filename := os.Args[2]

  file, err := os.Open(filename)
  if err != nil {
    fmt.Println("Failed to open file", filename, err)
    os.Exit(1)
  }
  defer file.Close()

  //select Region to use.
  conf := aws.Config{Region: aws.String("us-west-2")}
  sess := session.New(&conf)
  svc := s3manager.NewUploader(sess)

  fmt.Println("Uploading file to S3...")
  result, err := svc.Upload(&s3manager.UploadInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(filepath.Base(filename)),
    Body:   file,
  })
  if err != nil {
    fmt.Println("error", err)
    os.Exit(1)
  }

  fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}