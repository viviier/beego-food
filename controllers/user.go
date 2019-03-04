package controllers

import (
	"beego-food/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserControllers struct {
	beego.Controller
}

func (this *UserControllers) Login() {
	var user models.User
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		this.Data["json"] = models.Login(user.UserName, user.Password)
		this.SetSession("uid", user.UserName)
		this.Ctx.SetCookie("uid", user.UserName)
		this.ServeJSON()
	}
}

func (this *UserControllers) Signup() {
	var user models.User
	fmt.Println("ok")
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		this.Data["json"] = models.Signup(user)
	}
	this.Data["ok"] = "yes"
	this.ServeJSON()
}