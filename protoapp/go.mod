module github.com/geremde/genuine/protoapp

go 1.14

require (
	github.com/geremde/genuine/core v0.0.0
	github.com/geremde/genuine/http v0.0.0
	github.com/geremde/genuine/orm v0.0.0
	github.com/gotk3/gotk3 v0.5.1

	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.9
)

replace (
	github.com/geremde/genuine/core => ../core
	github.com/geremde/genuine/http => ../http
	github.com/geremde/genuine/orm => ../orm
)
