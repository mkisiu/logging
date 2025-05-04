//go:build linux

package logging

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

type linuxLogger struct {
	logger *log.Logger
	level  LogLevel
}

func newLinuxLogger(config Config) (Logger, error) {
	sysWriter, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "")
	if err != nil {
		return nil, err
	}

	writers := []io.Writer{sysWriter}

	if config.LogToFile {
		fileWriter, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}
		writers = append(writers, fileWriter)
	}

	multi := io.MultiWriter(writers...)
	return &linuxLogger{
		logger: log.New(multi, "", log.LstdFlags),
		level:  config.LogLevel,
	}, nil
}

func (l *linuxLogger) Info(message string) {
	l.logger.Println("[INFO]", message)
}

func (l *linuxLogger) Error(message string) {
	l.logger.Println("[ERROR]", message)
}
