package user

import "github.com/geremde/genuine/orm/model"

type user struct {
	model.Model
	username string
	password string
}
