package routers

import (
	"github.com/astaxie/beego"
	"sdrms_test/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
}
