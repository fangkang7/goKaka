package logger

import (
	"testing"
)

func TestNewConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] id come form china", 19961004)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()
}

func TestNewFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "D:/logs", "kaka")
	logger.Debug("user id[%d] id come form china", 19961004)
	logger.Warn("test warn log")
	logger.Close()
}
