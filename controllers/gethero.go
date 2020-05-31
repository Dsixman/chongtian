package controllers

import (
	"github.com/astaxie/beego"
	"replayanaly/models"
	//"replayanaly/models/obj"
	//_ "replayanaly/models/mongodb"
	"encoding/json"
	"fmt"
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
	Position string `json:"position"`
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
	position:=heroobj.Position
	fmt.Printf("heroname1:%v\n",heroname)
	fmt.Printf("version:%v\n",version)
	herobasedata:=models.GetHeroData(heroname,version)
	heromatchdata:=models.GetMatchHeroData(heroname,version,position)
	//resp:=&obj.HeroModel{}
	resp := make(map[string]interface{})
	resp["base_data"]=herobasedata
	resp["match_data"]=heromatchdata
	this.Data["json"]=resp
	this.ServeJSON()
}

func(this *GetHeroIconController) Get(){
	var resp=models.HeroNameIcon
	this.Data["json"]=resp
}