package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"openvpn/models"
)

type LogsController struct {
	beego.Controller
}

func (this *LogsController) Get() {
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	var err error
	this.Data["Logs"], err = models.GetLogsAll()
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "logs.html"

}
