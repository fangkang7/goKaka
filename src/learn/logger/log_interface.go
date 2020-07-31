package logger

type LogInterface interface {
	// 定义日志等级
	SetLevel(level int)

	// 定义不同日志实现方式
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}
