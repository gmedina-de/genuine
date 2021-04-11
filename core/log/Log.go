package log

import "github.com/geremde/genuine/core/configuration"

type Log interface {
	Debug(message string, parameters ...interface{})
	Info(message string, parameters ...interface{})
	Warning(message string, parameters ...interface{})
	Error(message string, parameters ...interface{})
}

const (
	LogLevelSetting configuration.Setting = iota
)
