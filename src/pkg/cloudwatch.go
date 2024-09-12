package pkg

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type CloudWatchLogsDatastore struct {
	Client        *cloudwatchlogs.Client
	LogGroupName  string
	LogStreamName string
}

func NewCloudWatchLogsDatastore() *CloudWatchLogsDatastore {
	accessKeyID := os.Getenv("AWS_CLOUDWATCH_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_CLOUDWATCH_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_CLOUDWATCH_REGION")
	logGroupName := os.Getenv("AWS_CLOUDWATCH_LOG_GROUP_NAME")
	logStreamName := os.Getenv("AWS_CLOUDWATCH_LOG_STREAM_NAME")

	creds := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(creds),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := cloudwatchlogs.NewFromConfig(cfg)

	return &CloudWatchLogsDatastore{
		Client:        client,
		LogGroupName:  logGroupName,
		LogStreamName: logStreamName,
	}
}

func (cw *CloudWatchLogsDatastore) PutLogEvent(ctx context.Context, message string) error {
	describeStreamsOutput, err := cw.Client.DescribeLogStreams(ctx, &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &cw.LogGroupName,
		LogStreamNamePrefix: &cw.LogStreamName,
	})

	if err != nil {
		return fmt.Errorf("failed to describe log streams: %w", err)
	}

	if len(describeStreamsOutput.LogStreams) == 0 {
		return fmt.Errorf("log stream %s not found", cw.LogStreamName)
	}

	sequenceToken := describeStreamsOutput.LogStreams[0].UploadSequenceToken

	timestamp := time.Now().UnixMilli()
	_, err = cw.Client.PutLogEvents(ctx, &cloudwatchlogs.PutLogEventsInput{
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
		return fmt.Errorf("failed to put log event: %w", err)
	}

	return nil
}
