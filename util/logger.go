package util

import (
	"fmt"
	"os"
	"time"
)

const (
	// LevelError 错误
	LevelError = iota
	// LevelWarning 警告
	LevelWarning
	// LevelInformational 提示
	LevelInformational
	// LevelDebug 查错
	LevelDebug
)

var logger *Logger

type Logger struct {
	level int
}

// Println 打印
func (log *Logger)Println(msg string)  {
	fmt.Printf("%s %s", time.Now().Format("2006-01-02 15:01:03"), msg)
}

// Panic 极端错误
func (log *Logger)Panic(format string, v ...interface{})  {
	if LevelError > log.level {
		return
	}

	msg := fmt.Sprintf("[Panic] " + format, v...)
	log.Println(msg)
	os.Exit(0)
}

// Error 错误
func (log *Logger) Error(format string, v ...interface{}) {
	if LevelError > log.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	log.Println(msg)
}

// Warning 警告
func (log *Logger) Warning(format string, v ...interface{}) {
	if LevelWarning > log.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	log.Println(msg)
}

// Info 信息
func (log *Logger) Info(format string, v ...interface{}) {
	if LevelInformational > log.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	log.Println(msg)
}

// Debug 校验
func (log *Logger) Debug(format string, v ...interface{}) {
	if LevelDebug > log.level {
		return
	}
	msg := fmt.Sprintf("[D] "+format, v...)
	log.Println(msg)
}

// BuildLogger 构建logger
func BuildLogger(level string) {
	intLevel := LevelError
	switch level {
	case "error":
		intLevel = LevelError
	case "warning":
		intLevel = LevelWarning
	case "info":
		intLevel = LevelInformational
	case "debug":
		intLevel = LevelDebug
	}
	l := Logger{
		level: intLevel,
	}
	logger = &l
}

// Log 返回日志对象
func Log() *Logger {
	if logger == nil {
		l := Logger{
			level: LevelDebug,
		}
		logger = &l
	}
	return logger
}
