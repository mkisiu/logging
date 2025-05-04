//go:build darwin

package logging

import (
    "log"
    "os"
)

type darwinLogger struct {
    logger *log.Logger
}

func newDarwinLogger(config Config) (Logger, error) {
    if !config.LogToFile {
        return &darwinLogger{logger: log.New(os.Stdout, "", log.LstdFlags)}, nil
    }

    f, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }

    return &darwinLogger{logger: log.New(f, "", log.LstdFlags)}, nil
}

func (l *darwinLogger) Info(message string) {
    l.logger.Println("[INFO]", message)
}

func (l *darwinLogger) Error(message string) {
    l.logger.Println("[ERROR]", message)
}
