package uploader

import (
  "fmt"
  "os"
  dotenv "github.com/joho/godotenv"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "log"
)

func loadAwsCredential() (access_key_id string, secret_access_key string, region string, bucket string) {
  err := dotenv.Load()
  if err != nil {
    log.Fatal("Error laoding .env file")
  }
  access_key_id = os.Getenv("S3_BUCKET")
  secret_access_key = os.Getenv("AWS_ACCESS_KEY_ID")
  region = os.Getenv("AWS_SECRET_ACCESS_KEY")
  bucket = os.Getenv("ap-southeast-1")

  return
}

func Upload(filepath string) {
  fmt.Println("Uploader")

  access_key_id, secret_access_key, region, bucket := loadAwsCredential()

  // awsRegion := os.Getenv(region)
  // s3Bucket := os.Getenv(bucket)
  // awsID := os.Getenv(access_key_id)
  // awsSecretKey := os.Getenv(secret_access_key)
  awsRegion := region
  s3Bucket := bucket
  awsID := access_key_id
  awsSecretKey := secret_access_key
  sess, err := session.NewSession(&aws.Config{
   Region:      aws.String(awsRegion),
   Credentials: credentials.NewStaticCredentials(awsID, awsSecretKey, ""),
  })

  file, err := os.Open(filepath)
  if err != nil {
    fmt.Println("Failed to open file", filepath, err)
    os.Exit(1)
  }
  defer file.Close()

  svc := s3manager.NewUploader(sess)
  fmt.Println("Uploading file to S3...")
  result, err := svc.Upload(&s3manager.UploadInput{
    Bucket: aws.String(s3Bucket),
    Key:    aws.String("filepath"),
    Body:   file,
  })
  if err != nil {
    fmt.Println("error", err)
    os.Exit(1)
  }

  fmt.Printf("Successfully uploaded %s to %s\n", "go_uploader_filename", result.Location)
}