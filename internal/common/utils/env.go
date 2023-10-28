package utils

import "os"

func GetLogLevel() string {
	return os.Getenv("LogLevel")
}
