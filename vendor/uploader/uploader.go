package uploader

import (
  "fmt"
  "os"
  dotenv "github.com/joho/godotenv"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/credentials"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
  "log"
)

func loadAwsCredential() (access_key_id string, secret_access_key string, region string, bucket string) {
  err := dotenv.Load()
  if err != nil {
    log.Fatal("Error laoding .env file")
  }

  bucket = os.Getenv("S3_BUCKET")
  access_key_id = os.Getenv("AWS_ACCESS_KEY_ID")
  secret_access_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
  region = os.Getenv("ap-southeast-1")

  return
}

func UploadMultipart(file multipart.File) {
  fmt.Println("Multipart Uploader")
  awsID, awsSecretKey, awsRegion, s3Bucket := loadAwsCredential()
  token := ""

  creds := credentials.NewStaticCredentials(awsID, awsSecretKey, token)

  // check credential validity
  _, err := creds.Get()
  if err != nil {
    // handle error
  }

  cfg := aws.NewConfig().WithRegion(awsRegion).WithCredentials(creds)

  svc := s3.New(session.New(), cfg)

  fileInfo, _ := file.Stat()

  size := fileInfo.Size()

  buffer := make([]byte, size) // read file content to buffer

  file.Read(buffer)

  fileBytes := bytes.NewReader(buffer)
  fileType := http.DetectContentType(buffer)
  path := "/poc/" + file.Name()

  params := &s3.PutObjectInput{
    Bucket: s3Bucket,
    Key: awsSecretKey,
    Body: fileBytes,
    ContentLength: aws.Int64(size),
    ContentType: aws.String(fileType),
  }

  result, err := svc.PutObject(params)
  if err != nil {
    // handle error
  }

  fmt.Printf("response %s", awsutil.StringValue(result))
}

func Upload(filepath string) {
  fmt.Println("Uploader")

  awsID, awsSecretKey, awsRegion, s3Bucket := loadAwsCredential()

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
