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
// 录像上传路由
type ReplayUploadController struct {
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
	//fmt.Printf("this.Ctx.Request :%v\n",this.Ctx.Request )
	obj := &Logindata{};
	json.Unmarshal(this.Ctx.Input.RequestBody, obj)
	Name:=obj.Username
	Password:=obj.Password
	Captchakey:=obj.Captchakey
	Captcha:=obj.Captcha
	//fmt.Printf("name:%v\n",Name)
	errcode,loginstate,userinfo:=models.LoginValidate(Name,Password,Captchakey,Captcha)
	fmt.Printf("errcode:%v",errcode)
	userName:=userinfo.Name
	privilege:=userinfo.Privilege
	//fmt.Printf("userName:%v\n",userName)
	this.SetSession("userName",userName)
	this.SetSession("privilege",privilege)
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
		//fmt.Printf("key:%v\n",key)
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
func (this *ReplayUploadController) Get() {
	v:=this.GetSession("userName")
	if v!=nil{
		//loginstate="true"
		username:=v.(string)
		this.Data["Name"]=username
		this.TplName = "replay.html"
	}else{
		this.Redirect("/adminlogin", 302)	
	}
	
}
func(this *ReplayUploadController) Post() {
	f, h, err := this.GetFile("myfile") //获取上传的文件
	version:=this.GetString("version")
	path := "./replays/" + h.Filename    //文件目录
	if err != nil {
		fmt.Println(err)
	}
	f.Close()                       //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("myfile", path) //存文件
	models.Parse(version,path)
	this.Redirect("/repupload/", 302)
}

func (this *TestController) Get(){
	//mysql 测试
	/*a:=models.Test("zhengzhihui","zhengzhihui123#")
	fmt.Printf("a:%v\n",a)
	this.TplName = "test.tpl"*/
	models.SelectTest()
	this.TplName = "test.tpl"
	//mongodb测试

}

func (this *TestController) Post(){
	/*fmt.Printf("this.Ctx.Request :%v\n",this.Ctx.Request )
	ob := &Test{};
	json.Unmarshal(this.Ctx.Input.RequestBody, ob)
	this.Data["json"] = ob
	this.ServeJSON()*/
}
type HeroDataController struct{
	beego.Controller
}
func (this *HeroDataController) Get(){
	//heroobj:=
}
type SetPlayerInfoController struct{
	beego.Controller
}

type AdminLoginController struct{
	beego.Controller
}

func (this *AdminLoginController) Get(){
	v:=this.GetSession("userName")
	v2:=this.GetSession("privilege")
	var privilege int
	if v!=nil{
		//username:=v.(string)
		if v2!=nil{
			privilege=v2.(int)
		}
		if privilege!=1{
			this.Redirect("/repupload/", 302)
		}
	}else{
		this.TplName="adminlogin.html"	
		this.Data["Name"]="请先登录"	
	}
	
}
func (this *AdminLoginController) Post(){
	adminName:=this.GetString("name")
	pwd:=this.GetString("pwd")
	bl,privilege:=models.AdminLogin(adminName,pwd)	
	if bl==true{
		this.SetSession("userName",adminName)
		this.SetSession("privilege",privilege)
		this.Redirect("/repupload/", 302)
	}
}