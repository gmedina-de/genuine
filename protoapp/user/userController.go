package user

import (
	"github.com/geremde/genuine/http/controller"
	"github.com/geremde/genuine/http/router"
	"html/template"
	"net/http"
)

type userController struct {
	userRepository userRepository
}

func UserController(userRepository userRepository) controller.WebController {
	return &userController{userRepository: userRepository}
}

func (this *userController) Routing(router router.Router) {
	router.Get("/Hello", this.SayHello)
}

func (this *userController) SayHello(writer http.ResponseWriter, request *http.Request) {
	newUser := user{Username: "Admin", password: "asdf"}
	this.userRepository.Create(&newUser)
	var users []user
	this.userRepository.RetrieveAll(&users)

	tmpl, _ := template.ParseFiles("user/userView.html")
	_ = tmpl.Execute(writer,
		struct{ Users []user }{users},
	)
}
