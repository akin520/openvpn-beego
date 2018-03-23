package main

//编译成LINUX下面的软件
//set GOOS=linux
//set GOARCH=amd64
//go install

//version
//bee v1.6.2
//beego v1.7.2
//go v1.6.2

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	"openvpn/models"
	_ "openvpn/routers"
)

var globalSessions *session.Manager

func init() {
	models.RegisterDB()
	// globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	// go globalSessions.GC()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
