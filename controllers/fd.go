package controllers

import (
	"beego-food/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type FoodControllers struct {
	beego.Controller
}

func (this *FoodControllers) ItemsList() {
	this.Data["json"] = models.ItemsList()
	_ = this.GetSession("uid")
	fmt.Println(this.CruSession, "session")
	this.ServeJSON()
}

func (this *FoodControllers) ItemDetailSet() {
	var fd models.Fd
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &fd); err == nil {
		this.Data["json"] = models.ItemDetailSet(fd)
		this.ServeJSON()
	} else {
		fmt.Println(err)
	}
}