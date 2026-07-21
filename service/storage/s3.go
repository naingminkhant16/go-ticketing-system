package storage

import (
	"context"
	"io"
	"log"
	"ticketing-system/config/cloud"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 struct {
	client *s3.Client
}

func NewS3(client *s3.Client) *S3 {
	return &S3{client: client}
}

func (s S3) Upload(ctx context.Context, key string, body io.Reader) (string, error) {
	log.Printf("Uploading %s to S3", key)
	_, err := s.client.PutObject(
		ctx,
		&s3.PutObjectInput{Bucket: aws.String(cloud.S3Bucket), Key: aws.String(key), Body: body},
	)
	if err != nil {
		log.Println("failed to upload to S3 bucket!", err)
		return "", err
	}
	return key, nil
}

func (s S3) Delete(ctx context.Context, key string) error {
	//TODO implement me
	panic("implement me")
}

func (s S3) GetPresignedURL(ctx context.Context, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s S3) Download(ctx context.Context, key string) (io.ReadCloser, error) {
	//TODO implement me
	panic("implement me")
}
