package controllers

import (
	"github.com/astaxie/beego"
	"replayanaly/models"
	"replayanaly/models/obj"
	//_ "replayanaly/models/mongodb"
	"encoding/json"
	//"fmt"
)
type GetHeroDataController struct {
	beego.Controller
}
type GetHeroJson struct{
	IsGetHeroIcon int
	Version string
	HeroName string
}
type PostHeroJson struct{
	HeroIcons map[string]string `json:"hero_icons"`
	HeroData  *obj.HeroModel `json:"hero_data"`
}
func(this *GetHeroDataController) Get(){
	heroobj:=&GetHeroJson{} 
	json.Unmarshal(this.Ctx.Input.RequestBody, heroobj)
	var heroicons map[string]string
	heroname:=heroobj.HeroName
	version:=heroobj.Version
	herodata:=models.GetHeroData(heroname,version)
	if heroobj.IsGetHeroIcon==0{
		heroicons=models.HeroNameIcon
	}
	var resp PostHeroJson
	resp.HeroIcons=heroicons
	resp.HeroData=herodata
	this.Data["json"]=resp
}