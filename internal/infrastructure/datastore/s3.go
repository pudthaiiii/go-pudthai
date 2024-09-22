package datastore

import (
	"bytes"
	"context"
	"fmt"
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/logger"
	"go-ibooking/internal/throw"
	"go-ibooking/internal/utils"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
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

func NewS3Datastore(cfg *config.Config) *S3Datastore {
	accessKeyID := cfg.Get("S3")["AccessKeyID"].(string)
	secretAccessKey := cfg.Get("S3")["SecretAccessKey"].(string)
	region := cfg.Get("S3")["Region"].(string)
	bucket := cfg.Get("S3")["Bucket"].(string)

	creds := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	config, err := awsConfig.LoadDefaultConfig(context.Background(),
		awsConfig.WithRegion(region),
		awsConfig.WithCredentialsProvider(creds),
	)

	if err != nil {
		logger.Log.Err(err).Msg("unable to load SDK config")
	}

	client := s3.NewFromConfig(config)

	logger.Write.Info().Msg("Successfully connected to S3")

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
		return "", fmt.Errorf("failed to sign request: %w", err)
	}

	return signedURL.URL, nil
}

func (s *S3Datastore) ValidateAndUpload(ctx context.Context, file *multipart.FileHeader, fileName string) (*s3.PutObjectOutput, error) {
	if file == nil {
		return nil, fmt.Errorf("file is required")
	}

	fileExtension := filepath.Ext(file.Filename)

	if !contains(s.AllowedExtensions, fileExtension) {
		return nil, fmt.Errorf("invalid file type. %s", s.AllowedExtensions)
	}

	if s.MaxFileSize != "" {
		limit, _ := utils.CalFileSize(s.MaxFileSize)
		maxFileSize := limit

		if file.Size > maxFileSize {
			return nil, fmt.Errorf("file size exceeds the limit of %s", s.MaxFileSize)
		}
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buffer := new(bytes.Buffer)
	if _, err := buffer.ReadFrom(src); err != nil {
		return nil, err
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
		return nil, err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
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
