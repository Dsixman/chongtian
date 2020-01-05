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
type LeagueInofController struct{
	beego.Controller
}

type TeamInofController struct{
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
	var loginstate bool
	var resp = make(map[string]interface{})
	loginstate=false
	key,capstr:=models.CreateCaptcha()
		//fmt.Printf("key:%v\n",key)
	resp["loginState"] = loginstate
	resp["capKey"] =key
	resp["capStr"] =capstr
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
	var loginstate bool
	var username string
	if v!=nil{
		loginstate=true
		username=v.(string)	
	}else{
		loginstate=false
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
	fname:=h.Filename
	models.Parse(version,path,fname)
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

type GetLeagueInofController struct{
	beego.Controller
}

func (this *SetPlayerInfoController) Get(){
	v:=this.GetSession("userName")
	v2:=this.GetSession("privilege")
	var privilege int
	if v!=nil{
		username:=v.(string)
		if v2!=nil{
			privilege=v2.(int)
			if privilege==1{
				this.TplName="player.html"	
				this.Data["Name"]=username
			}else{
				this.TplName="adminlogin.html"	
				this.Data["Name"]="您没有权限访问player info页面,请与管理员联系"		
			}
		}else{
			this.TplName="adminlogin.html"
			this.Data["Name"]="请先登录"
		}
		
	}else{
		this.TplName="adminlogin.html"
		this.Data["Name"]="请先登录"
	}
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
		if privilege==1{
			this.Redirect("/repupload/", 302)
		}else{
			this.TplName="adminlogin.html"	
			this.Data["Name"]="请先登录"		
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
func (this *LeagueInofController) Get(){
	v:=this.GetSession("userName")
	v2:=this.GetSession("privilege")
	var privilege int
	if v!=nil{
		username:=v.(string)
		if v2!=nil{
			privilege=v2.(int)
			if privilege==1{
				this.TplName="league.html"	
				this.Data["Name"]=username
			}else{
				this.TplName="adminlogin.html"	
				this.Data["Name"]="您没有权限访问这个页面,请与管理员联系"		
			}
		}else{
			this.TplName="adminlogin.html"
		}
		
	}else{
		this.TplName="adminlogin.html"
		this.Data["Name"]="请先登录"
	}
}

func (this *LeagueInofController) Post(){
	f, h, err := this.GetFile("myfile") //获取上传的文件
	leaguename:=this.GetString("leaguename")
	leagueid,_:=this.GetInt("leagueid")
	seriesid,_:=this.GetInt("seriesid")
	path := "./views/src/static/img/league/" + h.Filename    //文件目录
	if err != nil {
		fmt.Println(err)
	}
	f.Close()                       //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("myfile", path) //存文件
	models.SaveLeagueInfo(h.Filename,leaguename,uint32(leagueid),uint32(seriesid))
	this.Redirect("/leagueinfo/", 302)
}

func (this *GetLeagueInofController) Get(){
	leageuname:=this.GetString("league_name")
	leagueinfo:=models.GetLeagueInfo(leageuname)
	this.Data["json"] =leagueinfo
	this.ServeJSON()
}

func (this *TeamInofController) Get(){
	v:=this.GetSession("userName")
	if v!=nil{
		v2:=this.GetSession("privilege")
		privilege:=v2.(int)
		if privilege==1{
			this.TplName="club.html"
			this.Data["name"]=v.(string)
		}else{
			this.TplName="club.html"
			this.Data["name"]="你没有权限访问此页面,请与管理员联系"
		}
	}else{
		this.TplName="club.html"
		this.Data["name"]="你还没有登录,请先登录"
	}
}

func (this *TeamInofController) Post(){
	f, h, err := this.GetFile("myfile") //获取上传的文件
	teamname:=this.GetString("teamname")
	teamtagname:=this.GetString("teamtagname")
	teamid,_:=this.GetInt("teamid")
	carry:=this.GetString("carry")
	mid:=this.GetString("mid")
	hard:=this.GetString("hard")
	softsupport:=this.GetString("softsupport")
	hardsupport:=this.GetString("hardsupport")
	path := "./views/src/static/img/team/" + h.Filename    //文件目录
	if err != nil {
		fmt.Println(err)
	}
	f.Close()                       //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	this.SaveToFile("myfile", path) //存文件
	models.SaveTeamInfo(h.Filename,teamname,teamtagname,uint32(teamid),carry,mid,hard,softsupport,hardsupport)
	this.Redirect("/teaminfo/", 302)
}
type AllTeamInofController struct{
	beego.Controller
}
func (this *AllTeamInofController) Get(){
	allinfo:=models.GetAllTeamInfo()
	this.Data["json"]=allinfo
	this.ServeJSON()
}
type TeamHeroPoolController struct{
	beego.Controller
}
func (this *TeamHeroPoolController) Get(){
	teamtagname:=this.GetString("teamtagname")
	version:=this.GetString("version")
	result:=models.GetTeamHeroPool(teamtagname,version)
	this.Data["json"]=result
	this.ServeJSON()
}