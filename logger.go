package logging

import (
	"fmt"
	"runtime"
)

type Logger interface {
	Info(message string)
	Error(message string)
}

func NewLogger(config Config) (Logger, error) {
	switch runtime.GOOS {
	case "linux":
		return newLinuxLogger(config)
	case "windows":
		return newWindowsLogger(config)
	case "darwin":
		return newDarwinLogger(config)
	default:
		return nil, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
}
