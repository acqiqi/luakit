package routers

import (
	"github.com/astaxie/beego"
	"luakit/controllers"
	"luakit/controllers/sm"
	"luakit/controllers/v1/message"
	"luakit/controllers/v1/platform"
	"luakit/controllers/v1/ucenter"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/sm",
			beego.NSAutoRouter(
				&sm.ApiController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSAutoRouter(
				&message.ApiController{},
			),
		),
		beego.NSNamespace("/ucenter",
			beego.NSAutoRouter(
				&ucenter.PublicController{},
			),
			beego.NSAutoRouter(
				&ucenter.ApiController{},
			),
		),
		beego.NSNamespace("/platform",
			beego.NSAutoRouter(
				&platform.PublicController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.AutoRouter(&sm.SmController{})
}
