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
func GetAbility (name string) *obj.Ability{
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    ability:=obj.Ability{}
    c:=session.DB("dota2_big_data").C("abilities")
    err := c.Find(bson.M{"Name":name}).One(&ability)
	if err!=nil{
		fmt.Printf("err:%v\n", err)
	}
	return &ability
}

func GetHeroData(heroname string,version string) *obj.HeroModel{
	var allresult obj.AllNpcHeroes
	var result *obj.NpcHeroes
	herodata:=&obj.HeroModel{}

	var ability1,ability2,ability3,ability4,ability5,ability6,ability7,ability8,ability9,ability10,ability11,
	ability12,ability13,ability14,ability15,ability16,ability17,ability18,ability19,ability20,ability21,ability22,ability23,ability24 obj.Ability
	
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("herodata")
	c2:=session.DB("dota2_big_data").C("abilities")
	//id:=HeroNameID[heroname]
	err := c.Find(bson.M{"Version":version}).One(&allresult)
	if err!=nil{
		fmt.Printf("err:%v\n", err)
	}
	heroid:=HeroNameID["heroname"]
	heroenname:=HeroIDENName[heroid]
	Loop:
	for _,v:=range allresult.Heroes{
		if v.HeroName==heroenname{
			result=v
			break Loop
		}
	}
	//fmt.Printf("result:%v\n", result)
	//abilitiesStr:=[]string{result.Abilities1,result.Abilities2,result.Abilities3,result.Abilities4,result.Abilities5,result.Abilities6,result.Abilities7,result.Abilities8,result.Abilities9,result.Abilities10,result.Abilities11,result.Abilities12,result.Abilities13,result.Abilities14,result.Abilities15,result.Abilities16,result.Abilities17,result.Abilities18,result.Abilities19,result.Abilities20,result.Abilities22,result.Abilities23,result.Abilities24}
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
	err1 := c2.Find(bson.M{"abilities.Name":result.Abilities1}).One(&ability1)
	if err1!=nil{
			fmt.Printf("err:%v\n", err1)
	}
	ability1.IconUrl=AbilityIcon[ability1.Name]
	herodata.Abilities=append(herodata.Abilities,&ability1)
	err2 := c2.Find(bson.M{"abilities.Name":result.Abilities2}).One(&ability2)
	if err2!=nil{
			fmt.Printf("err:%v\n", err2)
	}
	ability2.IconUrl=AbilityIcon[ability2.Name]
	herodata.Abilities=append(herodata.Abilities,&ability2)
	err3 := c2.Find(bson.M{"abilities.Name":result.Abilities3}).One(&ability3)
	if err3!=nil{
			fmt.Printf("err:%v\n", err3)
	}
	ability3.IconUrl=AbilityIcon[ability3.Name]
	herodata.Abilities=append(herodata.Abilities,&ability3)
	if result.Abilities4!=""{
		err4 := c2.Find(bson.M{"abilities.Name":result.Abilities4}).One(&ability4)
		if err4!=nil{
			fmt.Printf("err:%v\n", err4)
		}
		ability4.IconUrl=AbilityIcon[ability4.Name]
		herodata.Abilities=append(herodata.Abilities,&ability4)
	}
	if result.Abilities5!=""{
		err4 := c2.Find(bson.M{"abilities.Name":result.Abilities5}).One(&ability5)
		if err4!=nil{
			fmt.Printf("err:%v\n", err4)
		}
		ability5.IconUrl=AbilityIcon[ability5.Name]
		herodata.Abilities=append(herodata.Abilities,&ability5)
	}
	if result.Abilities6!=""{
		err6 := c2.Find(bson.M{"abilities.Name":result.Abilities6}).One(&ability6)
		if err6!=nil{
				fmt.Printf("err:%v\n", err6)
		}
		ability6.IconUrl=AbilityIcon[ability6.Name]
		herodata.Abilities=append(herodata.Abilities,&ability6)
	}
	if result.Abilities7!=""{
		err7 := c2.Find(bson.M{"abilities.Name":result.Abilities7}).One(&ability7)
		if err7!=nil{
				fmt.Printf("err:%v\n", err7)
		}
		ability7.IconUrl=AbilityIcon[ability7.Name]
		herodata.Abilities=append(herodata.Abilities,&ability7)
	}
	if result.Abilities8!=""{
		err8 := c2.Find(bson.M{"abilities.Name":result.Abilities8}).One(&ability8)
		if err8!=nil{
				fmt.Printf("err:%v\n", err8)
		}
		ability8.IconUrl=AbilityIcon[ability8.Name]
		herodata.Abilities=append(herodata.Abilities,&ability8)
	}
	
	if result.Abilities9!=""{
		err9 := c2.Find(bson.M{"abilities.Name":result.Abilities9}).One(&ability9)
		if err9!=nil{
				fmt.Printf("err:%v\n", err9)
		}
		ability9.IconUrl=AbilityIcon[ability9.Name]
		herodata.Abilities=append(herodata.Abilities,&ability9)
	}
	err10 := c2.Find(bson.M{"abilities.Name":result.Abilities10}).One(&ability10)
	if err10!=nil{
			fmt.Printf("err:%v\n", err10)
	}
	
	if result.Abilities10!=""{
		isexist:=strings.Index(result.Abilities10,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability10)
		}else{
			ability10.IconUrl=AbilityIcon[ability10.Name]
			herodata.Abilities=append(herodata.Talent,&ability10)
		}
	}
	err11 := c2.Find(bson.M{"abilities.Name":result.Abilities11}).One(&ability11)
	if err11!=nil{
			fmt.Printf("err:%v\n", err11)
	}
	if result.Abilities11!=""{
		isexist:=strings.Index(result.Abilities11,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability11)
		}else{
				ability11.IconUrl=AbilityIcon[ability11.Name]
			herodata.Abilities=append(herodata.Talent,&ability11)
		}
	}
	err12 := c2.Find(bson.M{"abilities.Name":result.Abilities12}).One(&ability12)
	if err12!=nil{
			fmt.Printf("err:%v\n", err12)
	}

	if result.Abilities12!=""{
		isexist:=strings.Index(result.Abilities12,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability12)
		}else{
				ability12.IconUrl=AbilityIcon[ability12.Name]
			herodata.Abilities=append(herodata.Talent,&ability12)
		}
	}

	err13 := c2.Find(bson.M{"abilities.Name":result.Abilities13}).One(&ability13)
	if err13!=nil{
			fmt.Printf("err:%v\n", err13)
	}
	
	if result.Abilities13!=""{
		isexist:=strings.Index(result.Abilities13,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability13)
		}else{
				ability13.IconUrl=AbilityIcon[ability13.Name]
			herodata.Abilities=append(herodata.Talent,&ability13)
		}
	}

	err14 := c2.Find(bson.M{"abilities.Name":result.Abilities14}).One(&ability14)
	if err14!=nil{
			fmt.Printf("err:%v\n", err14)
	}

	if result.Abilities14!=""{
		isexist:=strings.Index(result.Abilities14,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability14)
		}else{
				ability14.IconUrl=AbilityIcon[ability14.Name]
			herodata.Abilities=append(herodata.Talent,&ability14)
		}
	}

	err15 := c2.Find(bson.M{"abilities.Name":result.Abilities15}).One(&ability15)
	if err15!=nil{
			fmt.Printf("err:%v\n", err15)
	}

	if result.Abilities15!=""{
		isexist:=strings.Index(result.Abilities15,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability15)
		}else{
				ability15.IconUrl=AbilityIcon[ability15.Name]
			herodata.Abilities=append(herodata.Talent,&ability15)
		}
	}

	err16 := c2.Find(bson.M{"abilities.Name":result.Abilities16}).One(&ability16)
	if err16!=nil{
			fmt.Printf("err:%v\n", err16)
	}

	if result.Abilities16!=""{
		isexist:=strings.Index(result.Abilities16,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability16)
		}else{
				ability16.IconUrl=AbilityIcon[ability16.Name]
			herodata.Abilities=append(herodata.Talent,&ability16)
		}
	}

	err17 := c2.Find(bson.M{"abilities.Name":result.Abilities17}).One(&ability17)
	if err17!=nil{
			fmt.Printf("err:%v\n", err17)
	}

	if result.Abilities17!=""{
		isexist:=strings.Index(result.Abilities17,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability17)
		}else{
			herodata.Abilities=append(herodata.Talent,&ability17)
		}
	}

	if result.Abilities18!=""{
		err18 := c2.Find(bson.M{"abilities.Name":result.Abilities18}).One(&ability18)
		if err18!=nil{
			fmt.Printf("err:%v\n", err18)
		}
		isexist:=strings.Index(result.Abilities18,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability18)
		}else{
				
			herodata.Abilities=append(herodata.Talent,&ability18)
		}
	}



	if result.Abilities19!=""{
		err19 := c2.Find(bson.M{"abilities.Name":result.Abilities19}).One(&ability19)
		if err19!=nil{
			fmt.Printf("err:%v\n", err19)
		}	
		isexist:=strings.Index(result.Abilities19,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability19)
		}else{
				
			herodata.Abilities=append(herodata.Talent,&ability19)
		}
	}

	if result.Abilities20!=""{
		err20 := c2.Find(bson.M{"abilities.Name":result.Abilities20}).One(&ability20)
		if err20!=nil{
			fmt.Printf("err:%v\n", err20)
		}
		isexist:=strings.Index(result.Abilities20,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability20)
		}else{
				
			herodata.Abilities=append(herodata.Talent,&ability20)
		}
	}
	if result.Abilities21!=""{
		err21 := c2.Find(bson.M{"abilities.Name":result.Abilities21}).One(&ability21)
		if err21!=nil{
			fmt.Printf("err:%v\n", err21)
		}
		isexist:=strings.Index(result.Abilities21,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability21)
		}else{
				
			herodata.Abilities=append(herodata.Talent,&ability21)
		}
	}
	if result.Abilities22!=""{
		err22 := c2.Find(bson.M{"abilities.Name":result.Abilities22}).One(&ability22)
		if err22!=nil{
			fmt.Printf("err:%v\n", err22)
		}
		isexist:=strings.Index(result.Abilities22,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability22)
		}else{
				
			herodata.Abilities=append(herodata.Talent,&ability22)
		}
	}
	if result.Abilities23!=""{
		err23 := c2.Find(bson.M{"abilities.Name":result.Abilities23}).One(&ability23)
		if err23!=nil{
			fmt.Printf("err:%v\n", err23)
		}
		isexist:=strings.Index(result.Abilities23,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability23)
		}else{
				
			herodata.Abilities=append(herodata.Talent,&ability23)
		}
	}
	if result.Abilities24!=""{
		err24 := c2.Find(bson.M{"abilities.Name":result.Abilities24}).One(&ability24)
		if err24!=nil{
			fmt.Printf("err:%v\n", err24)
		}
		isexist:=strings.Index(result.Abilities24,"special_bonus")
		if isexist!=-1{
			herodata.Talent=append(herodata.Abilities,&ability24)
		}else{

			herodata.Abilities=append(herodata.Talent,&ability24)
		}
	}
	return herodata
}

