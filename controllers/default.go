package controllers

import (
	"github.com/astaxie/beego"
	"replayanaly/models"
	//"replayanaly/models/sql"
	"encoding/json"
	"fmt"
)
type Logindata struct{
		//Loginstate string
		Username string
		Password string
		Captcha string
		Captchakey string
	}
	type Test struct{
		//Loginstate string
		T string
	}
type Captcha struct{
	loginstate string
	key string
	capstr string
}
type CaptchaController struct{
	beego.Controller
}
type MainController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}
type AuthController struct {
	beego.Controller
}
func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "index.tpl"
}

type LoginController struct {
    beego.Controller
}
func (this *LoginController) Post() {
	//var login Logindata
	fmt.Printf("this.Ctx.Request :%v\n",this.Ctx.Request )
	ob := &Logindata{};
	json.Unmarshal(this.Ctx.Input.RequestBody, ob)
	Name:=ob.Username
	Password:=ob.Password
	Captchakey:=ob.Captchakey
	Captcha:=ob.Captcha
	//fmt.Printf("name:%v\n",Name)
	errcode,loginstate,userinfo:=models.LoginValidate(Name,Password,Captchakey,Captcha)
	fmt.Printf("errcode:%v",errcode)
	userName:=userinfo.Name
	fmt.Printf("userName:%v\n",userName)
	this.SetSession("userName",userName)
	//fmt.Printf("uinfo:%v",userinfo)
	resp := make(map[string]interface{})
	resp["loginState"] = loginstate
	resp["userName"] = userinfo.Name
	this.Data["json"] = resp
	this.ServeJSON()

}
func (this *LoginController) Get() {
	var loginstate string
	var resp = make(map[string]interface{})
	v:=this.GetSession("userName")
	if v!=nil{
		loginstate="true"
		resp["loginState"] = loginstate
		resp["userName"] = v.(string) 
	}else{
		loginstate="false"
		key,capstr:=models.CreateCaptcha()
		fmt.Printf("key:%v\n",key)
		resp["loginState"] = loginstate
		resp["capKey"] =key
		resp["capStr"] =capstr
	}
	this.Data["json"] = resp
	this.ServeJSON()
}
func (this *CaptchaController) Get(){
	key,capstr:=models.CreateCaptcha()
	resp := make(map[string]interface{})
	resp["capKey"] =key
	resp["capStr"] =capstr
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *AuthController) Get(){
	v:=this.GetSession("userName")
	var loginstate string
	var username string
	if v!=nil{
		loginstate="true"
		username=v.(string)	
	}else{
		loginstate="false"
	}
	resp := make(map[string]interface{})
	resp["loginState"] = loginstate
	resp["userName"] = username
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *TestController) Get(){
	a:=models.Test("zhengzhihui","zhengzhihui123#")
	fmt.Printf("a:%v\n",a)
	this.TplName = "test.tpl"
}

func (this *TestController) Post(){
	fmt.Printf("this.Ctx.Request :%v\n",this.Ctx.Request )
	ob := &T{};
	json.Unmarshal(this.Ctx.Input.RequestBody, ob)
	this.Data["json"] = ob
	this.ServeJSON()
}