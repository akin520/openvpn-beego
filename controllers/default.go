package controllers

import (
	"github.com/astaxie/beego"
	"openvpn/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	sess := this.StartSession()
	username := sess.Get("username")
	beego.Debug(username)
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	var err error
	this.Data["User"], err = models.GetAllUser()
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "index.html"
}
