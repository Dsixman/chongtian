package models

import(
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"strconv"
	"strings"
)
/*func GetAbility (name string) *obj.Ability{
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    ability:=obj.Ability{}
    c:=session.DB("dota2_big_data").C("Ability")
    err := c.Find(bson.M{"Name":name}).One(&ability)
	if err!=nil{
		fmt.Printf("err:%v\n", err)
	}
	return &ability
}*/

func GetHeroData(heroname string,version string) *obj.HeroModel{
	var allresult obj.AllNpcHeroes
	var result *obj.NpcHeroes
	var allabilities obj.AllNpcAbility
	herodata:=&obj.HeroModel{}

	/*var ability1,ability2,ability3,ability4,ability5,ability6,ability7,ability8,ability9,ability10,ability11,
	ability12,ability13,ability14,ability15,ability16,ability17,ability18,ability19,ability20,ability21,ability22,ability23,ability24 obj.Ability*/
	
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("herodata")
	c2:=session.DB("dota2_big_data").C("abilities")
	//id:=HeroNameID[heroname]
	heroid:=HeroNameID[heroname]
	heroenname:=HeroIDENName[heroid]
	err := c.Find(bson.M{"Version":version}).One(&allresult)
	if err!=nil{
		fmt.Printf("find hero err:%v\n", err)
	}
	
	Loop:
	for _,v:=range allresult.Heroes{
		if v.HeroName==heroenname{
			result=v
			break Loop
		}
	}
	fmt.Printf("result:%v\n", result)
	herodata.HeroID=heroid
	herodata.HeroName=HeroIDName[herodata.HeroID]
	herodata.Role=result.Role
	herodata.HeroIcon=result.Icon
	herodata.ArmorPhysical,_=strconv.Atoi(result.ArmorPhysical)
	herodata.AttackDamageMax,_=strconv.Atoi(result.AttackDamageMax)
	if result.StatusHealthRegen!="" {
		herodata.StatusHealthRegen,_=strconv.ParseFloat(result.StatusHealthRegen,64)
	}else{
		herodata.StatusHealthRegen=0.25
	}
	if result.StatusManaRegen!="" {
		herodata.StatusManaRegen,_=strconv.ParseFloat(result.StatusHealthRegen,64)
	}else{
		herodata.StatusHealthRegen=0.0
	}
	herodata.AttackDamageMax,_=strconv.Atoi(result.AttackDamageMax)
	herodata.AttackDamageMin,_=strconv.Atoi(result.AttackDamageMin)
	herodata.AttackRange=result.AttackRange
	herodata.AttackRate,_=strconv.ParseFloat(result.AttackRate,64)
	herodata.AttributeBaseAgility,_=strconv.Atoi(result.AttributeBaseAgility)
	herodata.AttributeAgilityGain,_=strconv.ParseFloat(result.AttributeAgilityGain, 64)
	herodata.AttributeBaseIntelligence,_=strconv.Atoi(result.AttributeBaseIntelligence)
	herodata.AttributeIntelligenceGain,_=strconv.ParseFloat(result.AttributeIntelligenceGain, 64)
	herodata.AttributeBaseStrength,_=strconv.Atoi(result.AttributeBaseStrength)
	herodata.AttributeStrengthGain,_=strconv.ParseFloat(result.AttributeStrengthGain, 64)
	herodata.AttributePrimary=result.AttributePrimary
	herodata.MovementSpeed,_=strconv.Atoi(result.MovementSpeed)
	herodata.MovementTurnRate,_=strconv.ParseFloat(result.MovementTurnRate,64)
	var abilityname []string
	abilityname=append(abilityname,result.Ability1,result.Ability2,result.Ability3)
	if result.Ability4!="generic_hidden"{
		abilityname=append(abilityname,result.Ability4)
	}
	if result.Ability5!="generic_hidden"{
		abilityname=append(abilityname,result.Ability5)
	}
	abilityname=append(abilityname,result.Ability6)
	if result.Ability7!=""{
		abilityname=append(abilityname,result.Ability7)
	}
	if result.Ability8!=""{
		abilityname=append(abilityname,result.Ability8)
	}
	if result.Ability9!=""{
		abilityname=append(abilityname,result.Ability9)
	}
	if result.Ability10!=""{
		abilityname=append(abilityname,result.Ability10)
	}
	if result.Ability11!=""{
		abilityname=append(abilityname,result.Ability11)
	}
	if result.Ability12!=""{
		abilityname=append(abilityname,result.Ability12)
	}
	if result.Ability13!=""{
		abilityname=append(abilityname,result.Ability13)
	}
	if result.Ability14!=""{
		abilityname=append(abilityname,result.Ability14)
	}
	if result.Ability15!=""{
		abilityname=append(abilityname,result.Ability15)
	}
	if result.Ability16!=""{
		abilityname=append(abilityname,result.Ability16)
	}
	if result.Ability17!=""{
		abilityname=append(abilityname,result.Ability17)
	}
	if result.Ability18!=""{
		abilityname=append(abilityname,result.Ability18)
	}
	if result.Ability19!=""{
		abilityname=append(abilityname,result.Ability19)
	}
	if result.Ability20!=""{
		abilityname=append(abilityname,result.Ability20)
	}
	if result.Ability21!=""{
		abilityname=append(abilityname,result.Ability21)
	}
	if result.Ability22!=""{
		abilityname=append(abilityname,result.Ability22)
	}
	if result.Ability23!=""{
		abilityname=append(abilityname,result.Ability23)
	}
	if result.Ability24!=""{
		abilityname=append(abilityname,result.Ability24)
	}
	findabilitieserr := c2.Find(bson.M{"Version":version}).One(&allabilities)
	if findabilitieserr!=nil{
		fmt.Printf("find hero err:%v\n", findabilitieserr)
	}
	for _,abilityvalue:=range abilityname{
		loop2:
		for _,v:=range allabilities.Ability{
			if abilityvalue==v.Name{
				if strings.Index(v.Name,"special_bonus")!=-1{
					herodata.Ability=append(herodata.Ability,v)
					break loop2
				}else{
					herodata.Talent=append(herodata.Talent,v)
					break loop2
				}
			}
		}
	}
	return herodata
}

