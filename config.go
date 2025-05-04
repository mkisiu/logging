package logging

// level of logging
type LogLevel int

// Logging levels
const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

// Config defines logging options
type Config struct {
	LogToFile bool
	FilePath  string
	LogLevel  LogLevel
}
