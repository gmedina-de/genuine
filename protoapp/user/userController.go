package user

import (
	"fmt"
	"github.com/geremde/genuine/http/controller"
	"github.com/geremde/genuine/http/router"
	"github.com/geremde/genuine/orm/repository"
	"net/http"
)

type userController struct {
	userRepository repository.Repository
}

func UserController(repository repository.Repository) controller.WebController {
	repository.Migrate(&user{})
	return &userController{userRepository: repository}
}

func (this *userController) Routing(router router.Router) {
	router.Get("/Hello", this.SayHello)
}

func (this *userController) SayHello(writer http.ResponseWriter, request *http.Request) {
	newUser := user{username: "Admin", password: "asdf"}
	this.userRepository.Create(&newUser)
	var foundUser user
	this.userRepository.RetrieveFirst(&user{}, 1)
	fmt.Fprintf(writer, "%s", foundUser)
}
