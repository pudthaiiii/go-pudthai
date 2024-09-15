package pkg

import (
	"bytes"
	"context"
	"fmt"
	throw "go-ibooking/src/app/exception"
	"go-ibooking/src/pkg/logger"
	"go-ibooking/src/utils"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Datastore struct {
	client            *s3.Client
	Bucket            string
	AllowedExtensions []string
	MaxFileSize       string
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
		logger.Log.Err(err).Msg("unable to load SDK config")
	}

	client := s3.NewFromConfig(cfg)

	return &S3Datastore{
		client:            client,
		Bucket:            bucket,
		AllowedExtensions: []string{".jpg", ".jpeg", ".png", ".gif"},
		MaxFileSize:       "5mb",
	}
}

func (s *S3Datastore) CheckConnection() error {
	input := &s3.HeadBucketInput{
		Bucket: &s.Bucket,
	}

	_, err := s.client.HeadBucket(context.Background(), input)
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
		logger.Log.Err(err).Msg("failed to sign request")
		return "", fmt.Errorf("failed to sign request: %w", err)
	}

	return signedURL.URL, nil
}

func (s *S3Datastore) ValidateAndUpload(file *multipart.FileHeader, fileName string) (*s3.PutObjectOutput, error) {
	if file == nil {
		return nil, throw.Error(910003, fmt.Errorf("file is required"))
	}

	fileExtension := filepath.Ext(file.Filename)

	if !contains(s.AllowedExtensions, fileExtension) {
		return nil, throw.Error(910003, fmt.Errorf("invalid file type. %s", s.AllowedExtensions))
	}

	if s.MaxFileSize != "" {
		limit, _ := utils.CalFileSize(s.MaxFileSize)
		maxFileSize := limit

		if file.Size > maxFileSize {
			return nil, throw.Error(910003, fmt.Errorf("file size exceeds the limit of %s", s.MaxFileSize))
		}
	}

	src, err := file.Open()
	if err != nil {
		return nil, throw.Error(910003, err)
	}
	defer src.Close()

	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(src); err != nil {
		return nil, throw.Error(910003, err)
	}

	return s.UploadFile(context.Background(), fileName, buffer.Bytes())
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
		return nil, throw.Error(910003, fmt.Errorf("failed to upload file"))
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
		logger.Log.Err(err).Msg("failed to get file")
		return nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		logger.Log.Err(err).Msg("failed to read file")
		return nil, err
	}

	return body, nil
}

func (s *S3Datastore) DeleteFile(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: &s.Bucket,
		Key:    &key,
	}

	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		logger.Log.Err(err).Msg("failed to delete file")
		return err
	}

	return nil
}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
