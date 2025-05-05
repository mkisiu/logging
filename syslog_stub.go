//go:build !linux

package logging

import (
	"errors"
	"io"
)

func getSyslogWriter() (io.Writer, error) {
	return nil, errors.New("syslog not available on this platform")
}
