package main

import (
	"authentication_backend/filter"
	_ "authentication_backend/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*/hello", beego.BeforeRouter, filter.AuthenticationFilter)
	beego.Run()
}
