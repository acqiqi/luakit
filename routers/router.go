package routers

import (
	"github.com/astaxie/beego"
	"luakit/controllers"
	"luakit/controllers/v1/face"
	"luakit/controllers/v1/marketing"
	"luakit/controllers/v1/message"
	"luakit/controllers/v1/platform"
	"luakit/controllers/v1/ucenter"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/marketing",
			beego.NSAutoRouter(
				&marketing.SeckillController{},
			),
			beego.NSAutoRouter(
				&marketing.CouponController{},
			),
		),
		beego.NSNamespace("/face",
			beego.NSAutoRouter(
				&face.FaceController{},
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
			beego.NSAutoRouter(
				&ucenter.OrdersController{},
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
