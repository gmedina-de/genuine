package router

import (
	"fmt"
	"github.com/geremde/genuine/core/log"
	"net/http"
	"strings"
)

type defaultRouter struct {
	log    log.Log
	routes map[string]Handler
}

func DefaultRouter(log log.Log) Router {
	return &defaultRouter{log: log, routes: make(map[string]Handler)}
}

func (r *defaultRouter) Get(route string, handler Handler) {
	r.routes["GET "+route] = handler
}

func (r *defaultRouter) Post(route string, handler Handler) {
	r.routes["POST "+route] = handler
}

func (r *defaultRouter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.log.Debug("%s %s", request.Method, request.URL)
	handler, found := r.routes[strings.ToUpper(request.Method)+" "+request.URL.Path]
	if found {
		handler(writer, request)
	} else {
		fmt.Fprintf(writer, "No route for %s found!", request.URL.Path)
	}
}
