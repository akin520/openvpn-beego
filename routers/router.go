package routers

import (
	"github.com/astaxie/beego"
	"openvpn/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logs", &controllers.LogsController{})
	beego.Router("/user", &controllers.UserController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/del/:id", &controllers.UserdelController{})
	beego.Router("/addtime/:id", &controllers.AddtimeController{})
}
