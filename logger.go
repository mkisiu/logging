package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

type Config struct {
	LogToFile bool
	FilePath  string
	LogLevel  LogLevel
}

type Logger struct {
	logger *log.Logger
	level  LogLevel
}

func NewLogger(config Config) (*Logger, error) {
	var writers []io.Writer

	// Linux-specific: attach syslog
	if runtime.GOOS == "linux" {
		if w, err := getSyslogWriter(); err == nil {
			writers = append(writers, w)
		}
	}

	// Optional file logger
	if config.LogToFile {
		fileWriter, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open log file: %w", err)
		}
		writers = append(writers, fileWriter)
	}

	if len(writers) == 0 {
		// Default to stdout
		writers = append(writers, os.Stdout)
	}

	mw := io.MultiWriter(writers...)
	return &Logger{
		logger: log.New(mw, "", log.LstdFlags),
		level:  config.LogLevel,
	}, nil
}

func (l *Logger) shouldLog(level LogLevel) bool {
	return level >= l.level
}

func (l *Logger) Debug(msg string) {
	if l.shouldLog(DEBUG) {
		l.logger.Println("[DEBUG]", msg)
	}
}

func (l *Logger) Info(msg string) {
	if l.shouldLog(INFO) {
		l.logger.Println("[INFO]", msg)
	}
}

func (l *Logger) Warning(msg string) {
	if l.shouldLog(WARNING) {
		l.logger.Println("[WARNING]", msg)
	}
}

func (l *Logger) Error(msg string) {
	if l.shouldLog(ERROR) {
		l.logger.Println("[ERROR]", msg)
	}
}
