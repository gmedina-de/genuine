package main

import (
	"github.com/geremde/genuine/core"
	"github.com/geremde/genuine/core/configuration"
	"github.com/geremde/genuine/core/log"
	"github.com/geremde/genuine/http"
	"github.com/geremde/genuine/http/router"
	"github.com/geremde/genuine/http/server"
	"github.com/geremde/genuine/orm/repository"
	"github.com/geremde/genuine/protoapp/user"
)

func main() {
	core.Constructor(log.ConsoleLog)
	core.Constructor(configuration.MapConfiguration)
	core.Constructor(repository.GormRepository)
	core.Constructor(server.DefaultServer)
	core.Constructor(router.DefaultRouter)
	// todo allow arrays
	// todo allow only interface
	core.Constructor(user.UserController)
	core.Genuine(http.WebApplication).Run()
}
