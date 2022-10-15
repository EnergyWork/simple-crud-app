package logger

import (
	"fmt"
	"os"
	"time"
)

// handmade logger
// log message example => date | INFO | id | method |: text of log message

const (
	INFO       = "INFO"
	DEBUG      = "DEBUG"
	ERROR      = "ERROR"
	WARNING    = "WARNING"
	FATAL      = "FATAL"
	TimeFormat = "2006-01-02 15:04:05"
	Undefined  = "-"
)

type Logger struct {
	id     string
	method string
}

func NewLogger() *Logger {
	return &Logger{
		id:     Undefined,
		method: Undefined,
	}
}

func (log *Logger) SetID(id string) *Logger {
	log.id = id
	return log
}

func (log *Logger) SetMethod(method string) *Logger {
	log.method = method
	return log
}

func (log *Logger) template(level string, data ...interface{}) string {
	prefix := fmt.Sprintf("%v | %s | %s |", time.Now().Format(TimeFormat), level, log.method)
	msg := fmt.Sprint(data...)
	return fmt.Sprintln(prefix, msg)
}

func (log *Logger) templatef(level string, str string, data ...interface{}) string {
	prefix := fmt.Sprintf("%v | %s | %s |", time.Now().Format(TimeFormat), level, log.method)
	msg := fmt.Sprintf(str, data...)
	return fmt.Sprintln(prefix, msg)
}

func (log *Logger) Info(data ...interface{}) {
	fmt.Print(log.template(INFO, data...))
}

func (log *Logger) Infof(str string, data ...interface{}) {
	fmt.Print(log.templatef(INFO, str, data...))
}

func (log *Logger) Error(data ...interface{}) {
	fmt.Print(log.template(ERROR, data...))
}

func (log *Logger) Errorf(str string, data ...interface{}) {
	fmt.Print(log.templatef(ERROR, str, data...))
}

func (log *Logger) Warning(data ...interface{}) {
	fmt.Print(log.template(WARNING, data...))
}

func (log *Logger) Warningf(str string, data ...interface{}) {
	fmt.Print(log.templatef(WARNING, str, data...))
}

func (log *Logger) Debug(data ...interface{}) {
	fmt.Print(log.template(DEBUG, data...))
}

func (log *Logger) Debugf(str string, data ...interface{}) {
	fmt.Print(log.templatef(DEBUG, str, data...))
}

func (log *Logger) Fatal(data ...interface{}) {
	fmt.Print(log.template(FATAL, data...))
	os.Exit(1)
}

func (log *Logger) Fatalf(str string, data ...interface{}) {
	fmt.Print(log.templatef(FATAL, str, data...))
	os.Exit(1)
}
