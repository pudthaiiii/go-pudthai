package pkg

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Datastore struct {
	client *s3.Client
	Bucket string
}

func NewS3Datastore() *S3Datastore {
	accessKeyID := os.Getenv("AWS_S3_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_S3_REGION")
	bucket := os.Getenv("AWS_S3_BUCKET")

	creds := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Datastore{
		client: client,
		Bucket: bucket,
	}
}

func (s *S3Datastore) CheckConnection(ctx context.Context) error {
	input := &s3.HeadBucketInput{
		Bucket: &s.Bucket,
	}

	_, err := s.client.HeadBucket(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Datastore) GenerateSignedURL(key string, expiresIn time.Duration) (string, error) {
	psClient := s3.NewPresignClient(s.client)

	signedURL, err := psClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expiresIn))

	if err != nil {
		return "", fmt.Errorf("failed to sign request: %w", err)
	}

	return signedURL.URL, nil
}

func (s *S3Datastore) UploadFile(ctx context.Context, key string, body []byte) (*s3.PutObjectOutput, error) {
	bodyReader := bytes.NewReader(body)

	input := &s3.PutObjectInput{
		Bucket: &s.Bucket,
		Key:    &key,
		Body:   bodyReader,
		ACL:    types.ObjectCannedACLPrivate,
	}

	result, err := s.client.PutObject(ctx, input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *S3Datastore) GetFile(ctx context.Context, key string) ([]byte, error) {
	input := &s3.GetObjectInput{
		Bucket: &s.Bucket,
		Key:    &key,
	}

	result, err := s.client.GetObject(ctx, input)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
