package controllers

import (
	"github.com/astaxie/beego"
	"replayanaly/models"
	//"replayanaly/models/obj"
	//_ "replayanaly/models/mongodb"
	"encoding/json"
	//"fmt"
)
type GetHeroIconController struct{
	beego.Controller
}
type GetHeroDataController struct {
	beego.Controller
}
type GetHeroJson struct{
	Version string `json:"version"`
	HeroName string `json:"heroname`
}
/*type PostHeroJson struct{
	HeroData  *obj.HeroModel `json:"hero_data"`
}*/
func(this *GetHeroDataController) Post(){
	heroobj:=&GetHeroJson{} 
	json.Unmarshal(this.Ctx.Input.RequestBody, heroobj)
	//fmt.Printf("heroobj:%v\n",heroobj)
	heroname:=heroobj.HeroName
	version:=heroobj.Version
	//fmt.Printf("heroname1:%v\n",heroname)
	herodata:=models.GetHeroData(heroname,version)
	//fmt.Printf("herodata:%v\n",herodata)
	//resp:=&obj.HeroModel{}
	resp:=herodata
	this.Data["json"]=resp
	this.ServeJSON()
}

func(this *GetHeroIconController) Get(){
	var resp=models.HeroNameIcon
	this.Data["json"]=resp
}