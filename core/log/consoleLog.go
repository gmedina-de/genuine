package log

import (
	"fmt"
	"github.com/geremde/genuine/core/configuration"
	"time"
)

type consoleLog struct {
	configuration configuration.Configuration
}

func ConsoleLog(configuration configuration.Configuration) Log {
	return &consoleLog{configuration: configuration}
	// TODO use built in go logger?
}

func (this *consoleLog) Debug(message string, parameters ...interface{}) {
	this.log(Debug, message, parameters...)
}

func (this *consoleLog) Info(message string, parameters ...interface{}) {
	this.log(Info, message, parameters...)
}

func (this *consoleLog) Warning(message string, parameters ...interface{}) {
	this.log(Warning, message, parameters...)
}

func (this *consoleLog) Error(message string, parameters ...interface{}) {
	this.log(Error, message, parameters...)
}

func (this *consoleLog) log(level Level, message string, parameters ...interface{}) {
	fmt.Printf("%s %s %s%s %s%s %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level.toBgColor(),
		Reset,
		level.toFgColor(),
		level,
		Reset,
		fmt.Sprintf(message, parameters...),
	)
}
