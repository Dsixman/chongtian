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
    beego.Router("/setplayerinfo/?:playerid", &controllers.SetPlayerInfoController{})
    beego.Router("/gettenmindata/", &controllers.GetTenMinDataController{})
    //replay 上传录像
    beego.Router("/repupload/", &controllers.ReplayUploadController{})
    beego.Router("/leagueinfo/", &controllers.LeagueInofController{})
    beego.Router("/getleagueinfo/", &controllers.GetLeagueInofController{})
    beego.Router("/saveteaminfo/",&controllers.SaveTeamInofController{})
    beego.Router("/getallteaminfo/",&controllers.AllTeamInofController{})
    beego.Router("/getteamalldata/",&controllers.GetTeamInofController{})
    beego.Router("/getteamheropool",&controllers.TeamHeroPoolController{})
    beego.Router("/getmatchdetails",&controllers.MatchDetailsController{})
    beego.Router("/getmiduseheroes",&controllers.MidUseHeroesController{})
    beego.Router("/getsidelineup",&controllers.SideLineUpController{})
    beego.Router("/getplayersamehero",&controllers.PlayerSameHeroController{})


    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins:  true,
        //AllowOrigins:     []string{"http://192.168.0.107:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
    }))
}
