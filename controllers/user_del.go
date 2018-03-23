package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"openvpn/models"
)

type UserdelController struct {
	beego.Controller
}

func (this *UserdelController) Get() {
	//添加用户验证，暂没实现
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	uid := this.Ctx.Input.Param(":id")
	err := models.DelUser(uid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
	return
}
