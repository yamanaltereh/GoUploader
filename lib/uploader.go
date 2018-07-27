package uploader

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func newSession() {
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")
	awsID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	sess, err := session.NewSession(&aws.Config{
	 Region:      aws.String(awsRegion),
	 Credentials: credentials.NewStaticCredentials(awsID, awsSecretKey, ""),
	})
}
