package routers

import (
	"replayanaly/controllers"
	"github.com/astaxie/beego"
	//"replayanaly/controllers/admin"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/adminlogin/", &controllers.AdminLoginController{})
    beego.Router("/getcaptcha", &controllers.CaptchaController{})
    beego.Router("/test", &controllers.TestController{})
    beego.Router("/getAuth", &controllers.AuthController{})
    beego.Router("/getherodata", &controllers.GetHeroDataController{})
    beego.Router("/getheroesicon", &controllers.GetHeroIconController{})
    beego.Router("/setplayerinfo/", &controllers.SetPlayerInfoController{})
    //replay 上传录像
    beego.Router("/repupload/", &controllers.ReplayUploadController{})
    beego.Router("/leagueinfo/", &controllers.LeagueInofController{})
    beego.Router("/getleagueinfo/", &controllers.GetLeagueInofController{})
    beego.Router("/teaminfo/",&controllers.TeamInofController{})

    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins:  true,
        //AllowOrigins:     []string{"http://192.168.0.107:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
    }))
}
