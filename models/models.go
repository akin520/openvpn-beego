package models

// package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var Cfg = beego.AppConfig

type User struct {
	Id           int64
	Name         string `orm:"unique"`
	Password     string
	Expired_time time.Time `orm:"index;null"`
	Active       int
}

type Login_log struct {
	Id             int64
	Username       string    `orm:"null"`
	Login_time     time.Time `orm:"index;null"`
	Trusted_ip     string    `orm:"null"`
	Trusted_port   string    `orm:"null"`
	Protocol       string    `orm:"null"`
	Remote_ip      string    `orm:"null"`
	End_time       time.Time `orm:"null"`
	Bytes_received string    `orm:"null"`
	Bytes_sent     string    `orm:"null"`
}

func (*User) TableEngine() string {
	return engine()
}

func (*Login_log) TableEngine() string {
	return engine()
}

func engine() string {
	return "INNODB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci"
}

func RegisterDB() {
	orm.RegisterModel(new(User), new(Login_log))
	// database
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	maxIdleConn, _ := Cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("db_max_open_conn")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FChongqing"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	orm.Debug = true
	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Shanghai")

}

func AddUser(name string, password string) error {
	o := orm.NewOrm()
	srcData := []byte(password)
	hash := md5.New()
	hash.Write(srcData)
	cipherText2 := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText2)
	beego.Debug(hexText)
	year, _ := Cfg.Int("expired_year")
	month, _ := Cfg.Int("expired_month")
	day, _ := Cfg.Int("expired_day")
	m := time.Now().AddDate(year, month, day)
	u := &User{Name: name, Password: string(hexText), Expired_time: m, Active: 1}
	_, err := o.Insert(u)
	if err != nil {
		return err
	}
	return nil
}

func ModifyUser(uid string, name string, password string) error {
	id, err := strconv.ParseInt(uid, 10, 64)
	srcData := []byte(password)
	hash := md5.New()
	hash.Write(srcData)
	cipherText2 := hash.Sum(nil)
	hexText := make([]byte, 32)
	hex.Encode(hexText, cipherText2)
	year, _ := Cfg.Int("expired_year")
	month, _ := Cfg.Int("expired_month")
	day, _ := Cfg.Int("expired_day")
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	u := &User{Id: id}
	if o.Read(u) == nil {
		u.Password = string(hexText)
		u.Expired_time = time.Now().AddDate(year, month, day)
		u.Active = 1
		_, err = o.Update(u)
		if err != nil {
			return err
		}
	}
	return nil
}

func DelUser(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	u := &User{Id: cid}
	_, err = o.Delete(u)
	return err
}

func GetUser(uid string) (*User, error) {
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable("user")
	err = qs.Filter("id", id).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser() ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}

func GetLogsAll() ([]*Login_log, error) {
	o := orm.NewOrm()
	logs := make([]*Login_log, 0)
	qs := o.QueryTable("login_log")
	//https://github.com/joiggama/beego-example/blob/master/Godeps/_workspace/src/github.com/astaxie/beego/orm/docs/zh/Query.md
	_, err := qs.OrderBy("-id").Limit(10).All(&logs)
	return logs, err
}
