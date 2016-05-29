package controllers

import (
	"github.com/astaxie/beego"
)

type EditorController struct {
	beego.Controller
}

func (c *EditorController) Get() {
	c.TplNames = "editor.tpl"
}
