package routers

import (
	"blink2/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/fibonacci", &controllers.MainController{})
}
