//go:build linux

package logging

import (
	"log/syslog"
)

func getSyslogWriter() (*syslog.Writer, error) {
	return syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "")
}
