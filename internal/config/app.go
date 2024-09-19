package config

import (
	"go-ibooking/internal/utils"
)

var (
	bodyLimit, _  = utils.CalFileSize("10mb")
	bufferSize, _ = utils.CalFileSize("10mb")
)

func (c *Config) loadAppConfig() {
	c.cfg = map[string]interface{}{
		// logger config
		"LoggerConfig": map[string]interface{}{
			"FileLog": "storage/logs.log",
			"Enabled": getEnv("AWS_CLOUDWATCH_ENABLED", "false"),
		},

		// fiber config
		"FiberConfig": map[string]interface{}{
			"BodyLimit":       int(bodyLimit),
			"ReadBufferSize":  int(bufferSize),
			"WriteBufferSize": int(bufferSize),
			"Port":            getEnv("PORT", "3000"),
		},

		// Postgresql config
		"Postgresql": map[string]interface{}{
			"Host":     getEnv("DB_HOST"),
			"Port":     getEnv("DB_PORT"),
			"User":     getEnv("DB_USERNAME"),
			"Password": getEnv("DB_PASSWORD"),
			"DBName":   getEnv("DB_DATABASE"),
			"SSL":      getEnv("DB_SSL"),
		},

		// Redis config
		"Redis": map[string]interface{}{
			"ClusterEnabled": utils.StringToBool(getEnv("REDIS_CLUSTER_ENABLED")),
			"Host":           getEnv("REDIS_HOST"),
			"Port":           getEnv("REDIS_PORT"),
			"Password":       getEnv("REDIS_PASSWORD"),
			"DB":             getEnv("REDIS_DB"),
			"ClusterNodes":   getEnv("REDIS_CLUSTER_NODES"),
		},

		// S3 config
		"S3": map[string]interface{}{
			"AccessKeyID":     getEnv("AWS_S3_ACCESS_KEY_ID"),
			"SecretAccessKey": getEnv("AWS_S3_SECRET_ACCESS_KEY"),
			"Region":          getEnv("AWS_S3_REGION"),
			"Bucket":          getEnv("AWS_S3_BUCKET"),
		},

		// CloudWatch config
		"CloudWatch": map[string]interface{}{
			"AccessKeyID":     getEnv("AWS_CLOUDWATCH_ACCESS_KEY_ID"),
			"SecretAccessKey": getEnv("AWS_CLOUDWATCH_SECRET_ACCESS_KEY"),
			"Region":          getEnv("AWS_CLOUDWATCH_REGION"),
			"LogGroupName":    getEnv("AWS_CLOUDWATCH_LOG_GROUP_NAME"),
			"LogStreamName":   getEnv("AWS_CLOUDWATCH_LOG_STREAM_NAME"),
		},
	}
}
