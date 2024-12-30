package logger

import (
	"fmt"
	"github.com/fatih/color"
)

var (
	InfoLogger    = color.New(color.FgGreen).PrintfFunc()
	ErrorLogger   = color.New(color.FgRed).PrintfFunc()
	WarnLogger    = color.New(color.FgYellow).PrintfFunc()
	DatabaseLogger = color.New(color.FgCyan).PrintfFunc()
)

func Info(format string, args ...interface{}) {
	InfoLogger(format+"\n", args...)
}

func Error(format string, args ...interface{}) {
	ErrorLogger(format+"\n", args...)
}

func Warn(format string, args ...interface{}) {
	WarnLogger(format+"\n", args...)
}

func DatabaseConnection(dbName string, stats map[string]string) {
	DatabaseLogger("=== %s Connection Established ===\n", dbName)
	if stats != nil {
		for key, value := range stats {
			DatabaseLogger("  %s: %s\n", key, value)
		}
		DatabaseLogger("===============================\n")
	}
}