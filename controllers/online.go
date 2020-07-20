package controllers

import (
	//	"fmt"
	"openvpn/models"

	"github.com/astaxie/beego"
)

type OnlineController struct {
	beego.Controller
}

func (this *OnlineController) Get() {
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	var err error
	this.Data["Online"], err = models.GetOnline()
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "online.html"
}
