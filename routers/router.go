package routers

import (
	"CodeShare/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*", &controllers.MainController{})
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/", &controllers.ForeignController{})

}
