package logger

import (
	"fmt"
	"os"
)

type FileLogger struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
}

func NewFileLogger(level int, LogPath, LogName string) LogInterface {
	logger := &FileLogger{
		level:   level,
		logPath: LogPath,
		logName: LogName,
	}
	logger.init()
	return logger
}

func (f *FileLogger) init() {
	fileName := fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,err:%v", fileName, err))
	}
	f.file = file

	fileName = fmt.Sprintf("%s%s.log.ws", f.logPath, f.logName)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,err:%v", fileName, err))
	}
	f.warnFile = file
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	waiteLog(f.file, LogLevelDebug, format, args...)
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	waiteLog(f.file, LogLevelTrace, format, args...)
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	waiteLog(f.file, LogLevelInfo, format, args...)
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	waiteLog(f.warnFile, LogLevelWarn, format, args...)
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	waiteLog(f.warnFile, LogLevelError, format, args...)
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	waiteLog(f.warnFile, LogLevelFatal, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
