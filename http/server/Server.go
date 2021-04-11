package server

import (
	"github.com/geremde/genuine/core/configuration"
)

type Server interface {
	Start()
	Stop()
}

const (
	Port configuration.Setting = iota
)
