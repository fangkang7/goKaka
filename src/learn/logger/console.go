package logger

import "os"

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger(level int) LogInterface {
	return &ConsoleLogger{
		level: level,
	}
}

/**
设置等级  容错非法输出
*/
func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > LogLevelDebug {
		return
	}
	waiteLog(os.Stdout, LogLevelDebug, format, args...)
}
func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > LogLevelTrace {
		return
	}
	waiteLog(os.Stdout, LogLevelTrace, format, args...)
}
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > LogLevelInfo {
		return
	}
	waiteLog(os.Stdout, LogLevelInfo, format, args...)
}
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > LogLevelWarn {
		return
	}
	waiteLog(os.Stdout, LogLevelWarn, format, args...)
}
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > LogLevelError {
		return
	}
	waiteLog(os.Stdout, LogLevelError, format, args...)
}
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > LogLevelFatal {
		return
	}
	waiteLog(os.Stdout, LogLevelFatal, format, args...)
}

func (c *ConsoleLogger) Close() {

}
