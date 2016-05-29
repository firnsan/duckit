package main

import (
	_ "github.com/firnsan/duckit/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.DelStaticPath("/static")
	beego.SetStaticPath("//", "static")
	beego.Run()
}

