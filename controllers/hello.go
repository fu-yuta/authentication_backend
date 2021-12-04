package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about HelloController
type HelloController struct {
	beego.Controller
}

// @Title Hello
// @Description Hello
// @Success 200
// @router / [get]
func (h *HelloController) Hello() {
	h.Data["json"] = map[string]string{"message": "Hello World"}
	h.ServeJSON()
}
