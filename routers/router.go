package routers

import (
	"beego-food/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	ns := beego.NewNamespace("/api",

		beego.NSNamespace("/v1",

			beego.NSRouter("/login", &controllers.UserControllers{}, "Post:Login"),
			beego.NSRouter("/signup", &controllers.UserControllers{}, "Post:Signup"),
			beego.NSRouter("/fd.items.list", &controllers.FoodControllers{}, "Get:ItemsList"),
			beego.NSRouter("/fd.item.detail.set", &controllers.FoodControllers{}, "Post:ItemDetailSet"),
			//beego.NSRouter("/fd.item.detail.get"),
			//beego.NSRouter("/fd.shop"),
		),
	)

	beego.AddNamespace(ns)
}
