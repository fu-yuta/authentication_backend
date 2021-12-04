package routers

import (
	"authentication_backend/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthenticationController{},
			),
		),
		beego.NSNamespace("/hello",
			beego.NSInclude(
				&controllers.HelloController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
