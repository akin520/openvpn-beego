package routers

import (
	"openvpn/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logs", &controllers.LogsController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/updateall", &controllers.UpdateallController{})
	beego.Router("/online", &controllers.OnlineController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.Router("/del/:id", &controllers.UserdelController{})
	beego.Router("/addtime/:id", &controllers.AddtimeController{})
}
