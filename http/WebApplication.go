package http

import (
	"github.com/geremde/genuine/core"
	"github.com/geremde/genuine/core/log"
	"github.com/geremde/genuine/http/controller"
	"github.com/geremde/genuine/http/router"
	"github.com/geremde/genuine/http/server"
)

type webApplication struct {
	controller controller.WebController
	server     server.Server
	router     router.Router
}

func WebApplication(log log.Log, controller controller.WebController, server server.Server, router router.Router) core.Application {
	log.Info("Hello")
	log.Info("Hello")
	log.Debug("Hello")
	log.Debug("Hello")
	log.Info("Hello")
	log.Error("Hello")
	log.Error("Hello")
	log.Info("Hello")
	log.Warning("Hello")
	log.Info("Hello")
	return &webApplication{controller, server, router}
}

func (this *webApplication) Run() {
	this.controller.Routing(this.router)
	this.server.Start()
}
