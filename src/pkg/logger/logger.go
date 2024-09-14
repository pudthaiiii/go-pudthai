package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Log zerolog.Logger
var Write zerolog.Logger
var file *os.File

func NewInitializeLogger() {
	initLog()

	initWrite()
}

func initLog() {
	var err error
	fileLog := "storage/logs.log"

	file, err = os.OpenFile(fileLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		msg := fmt.Sprintf("Unable to create log file: %s", fileLog)
		log.Fatal().Err(err).Msg(msg)
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	enabledCw := os.Getenv("AWS_CLOUDWATCH_ENABLED")
	if enabledCw == "true" {
		cw, err := NewCloudWatchLogger()
		if err != nil {
			log.Printf("failed to create CloudWatchLogger: %v", err)
		} else {
			multi = zerolog.MultiLevelWriter(consoleWriter, file, io.Writer(cw))
		}
	}

	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	zerolog.CallerFieldName = "path"

	Log = zerolog.New(multi).
		With().
		Str("userInfo", "go-ibooking").
		Caller().
		Timestamp().
		Logger()
}

func initWrite() {
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter)

	zerolog.CallerFieldName = "path"

	Write = zerolog.New(multi).
		With().
		Timestamp().
		Logger()
}

func CloseLogger() {
	if file != nil {
		file.Close()
	}
}
