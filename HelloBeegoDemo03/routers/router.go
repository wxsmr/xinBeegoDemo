package routers

import (
	"HelloBeegoDemo03/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/", &controllers.MainController{})
	beego.Router("",&controllers.QueryController{})
}
