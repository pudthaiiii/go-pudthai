package logger

import (
	"context"
	"fmt"
	"go-ibooking/config"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/rs/zerolog/log"
)

type cloudWatchLogger struct {
	client        *cloudwatchlogs.Client
	LogGroupName  string
	LogStreamName string
}

func NewCloudWatchLogger(cfg *config.Config) (*cloudWatchLogger, error) {
	accessKeyID := cfg.Get("CloudWatch")["AccessKeyID"].(string)
	secretAccessKey := cfg.Get("CloudWatch")["SecretAccessKey"].(string)
	region := cfg.Get("CloudWatch")["Region"].(string)
	groupName := cfg.Get("CloudWatch")["LogGroupName"].(string)
	streamName := cfg.Get("CloudWatch")["LogStreamName"].(string)

	creds := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	config, err := awsConfig.LoadDefaultConfig(context.Background(),
		awsConfig.WithRegion(region),
		awsConfig.WithCredentialsProvider(creds),
	)

	if err != nil {
		log.Printf("unable to load SDK config, %v", err)
		return nil, err
	}

	client := cloudwatchlogs.NewFromConfig(config)

	return &cloudWatchLogger{
		client:        client,
		LogGroupName:  groupName,
		LogStreamName: streamName,
	}, nil
}

func (cw *cloudWatchLogger) Write(p []byte) (n int, err error) {
	message := string(p)

	cw.putLogEvent(message)
	return len(p), nil
}

func (cw *cloudWatchLogger) putLogEvent(message string) {
	describeStreamsOutput, err := cw.client.DescribeLogStreams(context.Background(), &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &cw.LogGroupName,
		LogStreamNamePrefix: &cw.LogStreamName,
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to send log to CloudWatch")
	}

	if len(describeStreamsOutput.LogStreams) == 0 {
		msg := fmt.Sprintf("log stream %s not found", cw.LogStreamName)
		log.Error().Msg(msg)
	}

	sequenceToken := describeStreamsOutput.LogStreams[0].UploadSequenceToken
	timestamp := time.Now().UnixMilli()

	_, err = cw.client.PutLogEvents(context.Background(), &cloudwatchlogs.PutLogEventsInput{
		LogEvents: []types.InputLogEvent{
			{
				Message:   aws.String(message),
				Timestamp: aws.Int64(timestamp),
			},
		},
		LogGroupName:  &cw.LogGroupName,
		LogStreamName: &cw.LogStreamName,
		SequenceToken: sequenceToken,
	})

	if err != nil {
		msg := fmt.Sprintf("failed to put log event: %s", err)
		log.Error().Err(err).Msg(msg)
	}
}
