package controllers

import (
	"github.com/astaxie/beego"
)

type ViewerController struct {
	beego.Controller
}

func (c *ViewerController) Get() {
	c.TplNames = "viewer.tpl"
}
