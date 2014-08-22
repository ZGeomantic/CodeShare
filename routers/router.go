package routers

import (
	"CodeShare/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
	beego.Router("/", &controllers.ForeignController{})
}
