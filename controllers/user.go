package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"openvpn/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	//输入内容
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	var err error
	id := this.Input().Get("uid")
	beego.Debug(id)
	user_name := this.Input().Get("username")
	user_pwd := this.Input().Get("password")
	if len(id) == 0 {
		err = models.AddUser(user_name, user_pwd)
	} else {
		err = models.ModifyUser(id, user_name, user_pwd)
	}

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
	return
}

func (this *UserController) Modify() {
	//检测登录
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.TplName = "user_modify.html"

	uid := this.Input().Get("uid")
	beego.Debug(uid)
	user, err := models.GetUser(uid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["User"] = user
	this.Data["Uid"] = uid
}
