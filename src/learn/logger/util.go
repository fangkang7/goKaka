package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

/**
获取文件名即行号
*/
func GetLineInfo() (fileName, funName string, lineNo int) {
	/**
	0:当前文件
	1:调用文件
	2:真实文件
	*/
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		funName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

/**
封装写日志
*/
func waiteLog(file *os.File, level int, format string, args ...interface{}) {
	// 格式化时间
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")

	// 日志级别字符串
	levelStr := getLevelText(level)
	// 获取文件名 方法名  以及行号
	fileName, funName, lineNo := GetLineInfo()
	// 只获取文件名和方法名
	fileName = path.Base(fileName)
	funName = path.Base(funName)
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s %s [%s:%s:%d] %s\n", nowStr, levelStr, fileName, funName, lineNo, msg)
}
