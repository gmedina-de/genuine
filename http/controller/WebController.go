package controller

import (
	"github.com/geremde/genuine/http/router"
)

type WebController interface {
	Routing(router router.Router)
}
