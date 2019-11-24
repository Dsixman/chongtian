package models

import(
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"strconv"
)
func GetAbility (name string) *obj.Ability{
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    //c := session.DB("dota2_big_data").C("CDotaGameInfo")
    //var result NpcAbilities
    ability:=obj.Ability{}
    c:=session.DB("dota2_big_data").C("abilities")
    err := c.Find(bson.M{}).Select(bson.M{name: 1}).One(&ability)
	if err!=nil{
		fmt.Printf("err:%v\n", err)
	}
	return &ability
}

func GetHeroData(name string) *obj.HeroModel{
	var result obj.NpcHeroes
	herodata:=&obj.HeroModel{}

	var ability1,ability2,ability3,ability4,ability5,ability6,ability7,ability8,ability9,ability10,ability11,
	ability12,ability13,ability14,ability15,ability16,ability17,ability18,ability19,ability20,ability21,ability22,ability23,ability24 obj.Ability
	herodata.Abilities=[]*obj.Ability{&ability1,&ability2,&ability3,&ability4,&ability5,&ability6,&ability7,&ability8,&ability9,&ability10,&ability11,
	&ability12,&ability13,&ability14,&ability15,&ability16,&ability17,&ability18,&ability19,&ability20,&ability21,&ability22,&ability23,&ability24}
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("herodata")
	err := c.Find(bson.M{}).Select(bson.M{name: 1}).One(&result)
	if err!=nil{
	fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("result:%v\n", result)
	abilitiesStr:=[]string{result.Abilities1,result.Abilities2,result.Abilities3,result.Abilities4,result.Abilities5,result.Abilities6,result.Abilities7,result.Abilities8,result.Abilities9,result.Abilities10,result.Abilities11,result.Abilities12,result.Abilities13,result.Abilities14,result.Abilities15,result.Abilities16,result.Abilities17,result.Abilities18,result.Abilities19,result.Abilities20,result.Abilities22,result.Abilities23,result.Abilities24}
	herodata.Id=result.Id
	herodata.HeroName=HeroIDName[herodata.Id]
	herodata.Role=result.Role
	herodata.ArmorPhysical,_=strconv.Atoi(result.ArmorPhysical)
	herodata.AttackDamageMax,_=strconv.Atoi(result.AttackDamageMax)
	herodata.AttackDamageMin,_=strconv.Atoi(result.AttackDamageMin)
	herodata.AttackRange=result.AttackRange
	herodata.AttackRate,_=strconv.Atoi(result.AttackRate)
	herodata.AttributeBaseAgility,_=strconv.Atoi(result.AttributeBaseAgility)
	herodata.AttributeAgilityGain,_=strconv.ParseFloat(result.AttributeAgilityGain, 64)
	herodata.AttributeBaseIntelligence,_=strconv.Atoi(result.AttributeBaseIntelligence)
	herodata.AttributeIntelligenceGain,_=strconv.ParseFloat(result.AttributeIntelligenceGain, 64)
	herodata.AttributeBaseStrength,_=strconv.Atoi(result.AttributeBaseStrength)
	herodata.AttributeStrengthGain,_=strconv.ParseFloat(result.AttributeStrengthGain, 64)
	herodata.AttributePrimary=result.AttributePrimary
	herodata.MovementSpeed,_=strconv.Atoi(result.MovementSpeed)
	herodata.MovementTurnRate,_=strconv.ParseFloat(result.MovementTurnRate,64)
	fmt.Printf("abilitiesStr", abilitiesStr)
	/*for key,value:=range herodata.Abilities{
		//value.AbilityName=abilitiesStr[key]
		GetAbility(value.AbilityName)
		if value!=""{

		}
	}*/
	//ability1.AbilityName=result.Abilities1
	return herodata
	//err1:=c.Find(bson.M{"matchid": 2271813303}).Select(bson.M{"last_hit.pre_5_count":1}).One(&NewCDotaGameInfo)
}