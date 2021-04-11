package controller

import (
	"github.com/geremde/genuine/core/configuration"
	log2 "github.com/geremde/genuine/core/log"
)

type consoleController struct {
	log           log2.Log
	configuration configuration.Configuration
}

func ConsoleController(log log2.Log, configuration configuration.Configuration) *consoleController {
	return &consoleController{log: log, configuration: configuration}
}

func (this *consoleController) LogSomething() {
	this.log.Debug("Hello world")
	this.log.Info("Hello world")
	this.log.Error("Hello world")
	this.log.Warning("Hello world")
}
