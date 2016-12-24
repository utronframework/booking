package controllers

import (
	"github.com/gernest/utron/controller"
	"github.com/gorilla/schema"
	"github.com/kr/pretty"
	"github.com/utronframework/booking/models"
)

var decoder = schema.NewDecoder()

type Account struct {
	controller.BaseController
	Routes []string
}

func (a *Account) Index() {
	a.Ctx.Template = "application/index"
	a.Ctx.Data["title"] = "Home"
}

func (a *Account) Register() {
	r := a.Ctx.Request()
	r.ParseForm()
	if r.Method == "GET" {
		a.Ctx.Template = "application/register"
		a.Ctx.Data["title"] = "Register"
		return
	}
	u := &models.Account{}
	err := decoder.Decode(u, r.PostForm)
	if err != nil {
		// set flash messages
		a.Ctx.Log.Errors(err)
		return
	}
	err = u.Validate()
	if err != nil {
		// set flash messages
		a.Ctx.Log.Errors(err)
		return
	}
	pretty.Println(u)
}

func (a *Account) Login() {
}

func (a *Account) Logout() {
}

func NewAccount() controller.Controller {
	return &Account{
		Routes: []string{
			"get;/;Index",
			"get,post;/register;Register",
			"get,post;/login;/Login",
			"get;/logout;Logout",
		},
	}
}
