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
    beego.Router("/getCaptcha", &controllers.CaptchaController{})
    beego.Router("/test", &controllers.TestController{})
    beego.Router("/getAuth", &controllers.AuthController{})
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins:  true,
        //AllowOrigins:      []string{"https://192.168.0.102"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
    }))
}
