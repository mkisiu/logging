//go:build windows

package logging

import (
    "log"
    "os"
)

type windowsLogger struct {
    logger *log.Logger
}

func newWindowsLogger(config Config) (Logger, error) {
    if !config.LogToFile {
        return &windowsLogger{logger: log.New(os.Stdout, "", log.LstdFlags)}, nil
    }

    f, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }

    return &windowsLogger{logger: log.New(f, "", log.LstdFlags)}, nil
}

func (l *windowsLogger) Info(message string) {
    l.logger.Println("[INFO]", message)
}

func (l *windowsLogger) Error(message string) {
    l.logger.Println("[ERROR]", message)
}
