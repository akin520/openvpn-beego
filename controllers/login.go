package controllers

import (
	//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

}

func (this *LoginController) Post() {
	//输入内容
	//this.Ctx.WriteString(fmt.Sprint(this.Input()))
	uname := this.Input().Get("username")
	pwd := this.Input().Get("password")

	if beego.AppConfig.String("username") == uname &&
		beego.AppConfig.String("password") == pwd {
		maxAge := 3600
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		sess := this.StartSession()
		sess.Set("username", uname)
		this.Redirect("/", 301)
	}
	this.Redirect("/login", 301)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value
	beego.Debug(uname)
	beego.Debug(ctx.Input.Session("username"))
	return uname == beego.AppConfig.String("username") && uname == ctx.Input.Session("username")
}
