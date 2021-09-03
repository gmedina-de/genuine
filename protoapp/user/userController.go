package user

import (
	"github.com/geremde/genuine/http/controller"
	"github.com/geremde/genuine/http/router"
	"github.com/joncalhoun/form"
	"html/template"
	"net/http"
)

var inputTpl = `
<label {{with .ID}}for="{{.}}"{{end}}>
	{{.Label}}
</label>
<input {{with .ID}}id="{{.}}"{{end}} type="{{.Type}}" name="{{.Name}}" placeholder="{{.Placeholder}}" {{with .Value}}value="{{.}}"{{end}}>
{{with .Footer}}
  <p>{{.}}</p>
{{end}}
`

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
	newUser := User{Username: "Admin", Password: "asdf"}
	this.userRepository.Create(&newUser)
	var users []User
	this.userRepository.RetrieveAll(&users)

	tpl := template.Must(template.New("").Parse(inputTpl))
	fb := form.Builder{
		InputTemplate: tpl,
	}

	tmpl, err := template.New("userView.html").Funcs(fb.FuncMap()).ParseFiles("user/userView.html")

	err = tmpl.Execute(writer,
		struct {
			Users   []User
			NewUser User
		}{users, newUser},
	)
	if err != nil {
		panic(err)
	}
}
