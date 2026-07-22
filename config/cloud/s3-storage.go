package cloud

import (
	"context"
	"log"
	config2 "ticketing-system/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client
var S3Bucket string

func LoadS3Config() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(config2.GetEnvOrDefault("AWS_REGION", "us-east-1")))
	if err != nil {
		log.Fatal("unable to load S3 config", err)
	}

	S3Client = s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.UsePathStyle = true
		options.BaseEndpoint = aws.String(config2.GetEnvOrDefault("AWS_ENDPOINT", "http://localstack:4566"))
	})

	S3Bucket = config2.GetEnvOrDefault("AWS_BUCKET", "go-ticketing")

	_, err = S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(S3Bucket),
	})

	if err != nil {
		log.Fatal("unable to create S3 bucket", err)
	}

	log.Println("S3 client loaded")
}
