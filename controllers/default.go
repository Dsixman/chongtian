package controllers

import (
	"github.com/astaxie/beego"
	"replayanaly/models"
	"replayanaly/models/sql"
	"encoding/json"
	"fmt"
)
type logindata struct{
		Loginstate string
		Userinfo *sql.User
	}
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type LoginController struct {
    beego.Controller
}
func (c *LoginController) Post() {
	
	var login logindata
	name:=c.GetString("username")
	password:=c.GetString("password")
	captcha:=c.GetString("captcha")
	key:=c.GetString("key")
	loginstate,userinfo:=models.LoginValidate(name,password,key,captcha)
	sess:=c.StartSession()
	sess.Set("beegosessionID",userinfo)
	login.Loginstate=loginstate
	login.Userinfo=userinfo
	loginjson,jsonerr:=json.Marshal(login)
	if jsonerr!=nil{
		fmt.Printf("jsonerr: %v",jsonerr)
	}
	c.Data["json"] = loginjson
	c.ServeJSON()

}
func (c *LoginController) Get() {
	var loginstate string
	v:=c.GetSession("beegosessionID")
	if v!=nil{
		loginstate="true"
	}else{
		loginstate="false"
	}
	c.Data["json"] = loginstate
	c.ServeJSON()
}
