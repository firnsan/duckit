package routers

import (
	"github.com/firnsan/duckit/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
    beego.Router("/editor", &controllers.EditorController{})
    beego.AutoRouter(&controllers.EvernoteController{})
}
