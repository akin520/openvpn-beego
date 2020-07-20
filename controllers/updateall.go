package controllers

import (
	//	"fmt"
	"openvpn/models"

	"github.com/astaxie/beego"
)

type UpdateallController struct {
	beego.Controller
}

func (this *UpdateallController) Get() {
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	var err error
	err = models.UpdateAllUser()
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 302)
	return
	//	this.Data["User"], err = models.GetAllUser()
	//	if err != nil {
	//		beego.Error(err)
	//	}
	//	this.TplName = "index.html"
}
