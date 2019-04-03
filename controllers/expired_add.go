package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"openvpn/models"
)

type AddtimeController struct {
	beego.Controller
}

func (this *AddtimeController) Get() {
	//添加用户验证，暂没实现
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	uid := this.Ctx.Input.Param(":id")
	beego.Debug(uid)
	err := models.AddTime(uid)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
	//this.Ctx.WriteString("{'code':200}")
	return
}
