package obj
import(
	"replayanaly/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"replayanaly/models" 
	"fmt"
	"strconv"
)
type HeroModel struct {
	HeroName string `json:"hero_name,omitempty"`
	Id uint32 `json:"HeroID,omitempty" bson:"HeroID,omitempty"`
	Abilities []*Ability `json:"Abilities,omitempty"`
	Talent[] string `json:"Talent"`
	Role string `json:"Role" bson:"Role"`
	ArmorPhysical string `json:"ArmorPhysical" bson:"ArmorPhysical"`
	AttackDamageMin int `json:"AttackDamageMin" bson:"AttackDamageMin"`
	AttackDamageMax int `json:"AttackDamageMax" bson:"AttackDamageMax"`
	AttackRate int `json:"AttackRate" bson:"AttackRate"`
	AttackRange string `json:"AttackRange" bson:"AttackRange"`
	AttributePrimary string `json:"AttributePrimary" bson:"AttributePrimary"`
	AttributeBaseStrength int `json:"AttributeBaseStrength" bson:"AttributeBaseStrength"`
	AttributeStrengthGain float64 `json:"AttributeStrengthGain" bson:"AttributeStrengthGain"`
	AttributeBaseIntelligence int `json:"AttributeBaseIntelligence" bson:"AttributeBaseIntelligence"`
	AttributeIntelligenceGain float64 `json:"AttributeIntelligenceGain" bson:"AttributeIntelligenceGain"`
	AttributeBaseAgility int `json:"AttributeBaseAgility" bson:"AttributeBaseAgility"`
	AttributeAgilityGain float64 `json:"AttributeAgilityGain" bson:"AttributeAgilityGain"`
	MovementSpeed int `json:"MovementSpeed" bson:"MovementSpeed"`
	MovementTurnRate float64 `json:"MovementTurnRate" bson:"MovementTurnRate"`
}
type NpcHeroes struct{
	Role string `json:"Role" bson:"Role"`
	//HeroName string `json:"hero_name"`
	Id uint32 `json:"HeroID" bson:"HeroID"`
	Abilities1 string `json:"Abilities1,omitempty" bson:"Abilities1"`
	Abilities2 string `json:"Abilities2,omitempty" bson:"Abilities2"`
	Abilities3 string `json:"Abilities3,omitempty" bson:"Abilities3"`
	Abilities4 string `json:"Abilities4,omitempty" bson:"Abilities4"`
	Abilities5 string `json:"Abilities5,omitempty" bson:"Abilities5"`
	Abilities6 string `json:"Abilities6,omitempty" bson:"Abilities6"`
	Abilities7 string `json:"Abilities7,omitempty" bson:"Abilities7"`
	Abilities8 string `json:"Abilities8,omitempty" bson:"Abilities8"`
	Abilities9 string `json:"Abilities9,omitempty" bson:"Abilities9"`
	Abilities10 string `json:"Abilities10,omitempty" bson:"Abilities10"`
	Abilities11 string `json:"Abilities11,omitempty" bson:"Abilities11"`
	Abilities12 string `json:"Abilities12,omitempty" bson:"Abilities12"`
	Abilities13 string `json:"Abilities13,omitempty" bson:"Abilities13"`
	Abilities14 string `json:"Abilities14,omitempty" bson:"Abilities14"`
	Abilities15 string `json:"Abilities15,omitempty" bson:"Abilities15"`
	Abilities16 string `json:"Abilities16,omitempty" bson:"Abilities16"`
	Abilities17 string `json:"Abilities17,omitempty" bson:"Abilities17"`
	Abilities18 string `json:"Abilities18,omitempty" bson:"Abilities18"`
	Abilities19 string `json:"Abilities19,omitempty" bson:"Abilities19"`
	Abilities20 string `json:"Abilities20,omitempty" bson:"Abilities20"`
	Abilities21 string `json:"Abilities21,omitempty" bson:"Abilities21"`
	Abilities22 string `json:"Abilities22,omitempty" bson:"Abilities22"`
	Abilities23 string `json:"Abilities23,omitempty" bson:"Abilities23"`
	Abilities24 string `json:"Abilities24,omitempty" bson:"Abilities24"`
	ArmorPhysical string `json:"ArmorPhysical" bson:"ArmorPhysical"`
	AttackDamageMin int `json:"AttackDamageMin" bson:"AttackDamageMin"`
	AttackDamageMax int `json:"AttackDamageMax" bson:"AttackDamageMax"`
	AttackRate int `json:"AttackRate" bson:"AttackRate"`
	AttackRange string `json:"AttackRange" bson:"AttackRange"`
	AttributePrimary string `json:"AttributePrimary" bson:"AttributePrimary"`
	AttributeBaseStrength int `json:"AttributeBaseStrength" bson:"AttributeBaseStrength"`
	AttributeStrengthGain float64 `json:"AttributeStrengthGain" bson:"AttributeStrengthGain"`
	AttributeBaseIntelligence int `json:"AttributeBaseIntelligence" bson:"AttributeBaseIntelligence"`
	AttributeIntelligenceGain float64 `json:"AttributeIntelligenceGain" bson:"AttributeIntelligenceGain"`
	AttributeBaseAgility int `json:"AttributeBaseAgility" bson:"AttributeBaseAgility"`
	AttributeAgilityGain float64 `json:"AttributeAgilityGain" bson:"AttributeAgilityGain"`
	MovementSpeed int `json:"MovementSpeed" bson:"MovementSpeed"`
	MovementTurnRate float64 `json:"MovementTurnRate" bson:"MovementTurnRate"`
}
type Ability struct {
	AbilityName string
	Id uint32
	Icon string
	AbilityBehavior string
	DamageType string
	Buff []DeBuff
	SpellImmunityType bool
	CD string
	HealthCost float32
	PerDamage int
	DamageTimes int
	BaseDamage int
	ManaCost string
	MoveSpeed int
	AttackSpeed int
	Attack int
	MagicResistance int
	HpHeal int
	AttributeChange float32
}
type DeBuff struct{
	Type int
	Duration float64
}
type NpcAbilitySpecial struct{
	Damage string `bson:"damage,omitempty" json:"damage,omitempty"`
	BaseDamage string `bson:"base_damage,omitempty" json:"base_damage,omitempty"`
	DamagePerTick string `bson:"damage_per_tick,omitempty" json:"damage_per_tick,omitempty"`
	Interval string `bson:"interval,omitempty" json:"interval,omitempty"`
	Duration string `bson:"duration,omitempty" json:"duration,omitempty"`
	StunDuration string `bson:"stun_duration,omitempty" json:"stun_duration,omitempty"`
	DamageCount string `bson:"damage_count,omitempty" json:"damage_count,omitempty"`
	Armor string `bson:"armor,omitempty" json:"armor,omitempty"`
	Coefficient string `bson:"coefficient,omitempty" json:"coefficient,omitempty"`
	AttackSpeed string `bson:"attack_speed,omitempty" json:"attack_speed,omitempty"`
	MovementSpeed string `bson:"movement_speed,omitempty" json:"movement_speed,omitempty"`
	MagicResistance string `bson:"magic_resistance,omitempty" json:"magic_resistance,omitempty"` 
	LinkedSpecialBonus string `bson:"LinkedSpecialBonus,omitempty" json:"LinkedSpecialBonus,omitempty"`

	AttackPec string `bson:"attack_pec,omitempty" json:"attack_pec,omitempty"`//增加百分比的（基础）（龙骑士全攻击力）攻击力
	Instances string `bson:"instances,omitempty" json:"instances,omitempty"`//圣堂刺客折光次数
	//承受物理伤害百分比 (美杜莎大招和巫妖冰甲)正数就是承受的物理伤害增加，负数就是承受的物理伤害降低
	PhysicalDamagePec string `bson:"physical_damage_pec,omitempty" json:"physical_damage_pec,omitempty"`
	AllDamagePec string `bson:"all_damage_pec,omitempty" json:"all_damage_pec,omitempty"`
	//BonusDamage string `bson:"bonus_damage,omitempty" json:"bonus_damage,omitempty"`
	//BounsAttack string `bson:"BounsAttack,omitempty" json:"BounsAttack,omitempty"`
	//承受魔法伤害百分比 正数就是承受的魔法伤害增加，负数就是承受的魔法伤害降低
	//MagicDamagePec string `bson:"magic_damage_pec,omitempty" json:"magic_damage_pec,omitempty"`
	//所有类型伤害加减 小精灵
	
}
/*
type AbilitySpecial1 struct{
}*/
type NpcAbilities struct{
	Id string `"json:ID" "bson:ID"`
	AbilityBehavior string `"json:AbilityBehavior" "bson:AbilityBehavior,omitempty"`
	DamageType string `json:"AbilityUnitDamageType,omitempty" json:"AbilityUnitDamageType,omitempty"`
	SpellImmunityType string `json:"SpellImmunityType,omitempty" json:"SpellImmunityType,omitempty"`
	AbilityCastRange string `json:"AbilityCastRange,omitempty" json:"AbilityCastRange,omitempty"`
	CD string `json:"AbilityCooldown,omitempty" bson:"AbilityCooldown,omitempty"`
	ManaCost string `json:"AbilityManaCost,omitempty" bson:"AbilityManaCost,omitempty"`
	//BaseDamage string `json:"AbilityDamage,omitempty" bson:"AbilityDamage,omitempty"`
	AbilitySpecial NpcAbilitySpecial `json:"AbilitySpecial,omitempty" bson:"AbilitySpecial,omitempty"`
	AbilityDamage string `bson:"AbilityDamage,omitempty" json:"AbilityDamage,omitempty"`
	AbilityDuration string`bson:"AbilityDuration,omitempty" json:"AbilityDuration,omitempty"`
}
func GetAbility (name string) *Ability{
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    var result NpcAbilities
    ability:=Ability{}
    c:=session.DB("dota2_big_data").C("abilities")
    err = d.Find(bson.M{}).Select(bson.M{name: 1}).One(&result)
	if err!=nil{
		fmt.Printf("err:%v\n", err)
	}
	ability.AbilityName=name
	ability.AbilityBehavior=result.AbilityBehavior
	ability.CD=result.CD
	ability.DamageType=result.DamageType
	ability.ManaCost=result.ManaCost
	ability.Id=result.Id
	ability.SpellImmunityType=result.SpellImmunityType
	return &ability
}

func GetHeroData(name string) *HeroModel{
	var result NpcHeroes
	herodata:=&HeroModel{}

	var ability1,ability2,ability3,ability4,ability5,ability6,ability7,ability8,ability9,ability10,ability11,
	ability12,ability13,ability14,ability15,ability16,ability17,ability18,ability19,ability20,ability21,ability22,ability23,ability24 Ability
	herodata.Abilities=[]*Ability{&ability1,&ability2,&ability3,&ability4,&ability5,&ability6,&ability7,&ability8,&ability9,&ability10,&ability11,
	&ability12,&ability13,&ability14,&ability15,&ability16,&ability17,&ability18,&ability19,&ability20,&ability21,&ability22,&ability23,&ability24}
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("herodata")
	err = d.Find(bson.M{}).Select(bson.M{name: 1}).One(&result)
	if err!=nil{
	fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("result:%v\n", result)
	abilitiesStr:=[]string{result.Abilities1,result.Abilities2,result.Abilities3,result.Abilities4,result.Abilities5,result.Abilities6,result.Abilities7,result.Abilities8,result.Abilities9,result.Abilities10,result.Abilities11,result.Abilities12,result.Abilities13,result.Abilities14,result.Abilities15,result.Abilities16,result.Abilities17,result.Abilities18,result.Abilities19,result.Abilities20,result.Abilities22,result.Abilities23,result.Abilities24}
	herodata.Id=result.Id
	herodata.HeroName=models.HeroIDName[herodata.Id]
	herodata.Role=result.Role
	herodata.ArmorPhysical,_=strconv.Atoi(result.ArmorPhysical)
	herodata.AttackDamageMax,_=strconv.Atoi(result.AttackDamageMax)
	herodata.AttackDamageMin,_=strconv.Atoi(result.AttackDamageMin)
	herodata.AttackRange,_=strconv.Atoi(result.AttackRange)
	herodata.AttackRate,_=strconv.Atoi(result.AttackRate)
	herodata.AttributeBaseAgility=strconv.Atoi(result.AttributeBaseAgility)
	herodata.AttributeAgilityGain,_=strconv.ParseFloat(result.AttributeAgilityGain, 64)
	herodata.AttributeBaseIntelligence,_=strconv.Atoi(result.AttributeBaseIntelligence)
	herodata.AttributeIntelligenceGain,_=strconv.ParseFloat(result.AttributeIntelligenceGain, 64)
	herodata.AttributeBaseStrength,_=strconv.Atoi(result.AttributeBaseStrength)
	herodata.AttributeStrengthGain,_=strconv.ParseFloat(result.AttributeStrengthGain, 64)
	herodata.AttributePrimary=result.AttributePrimary
	herodata.MovementSpeed=result.MovementSpeed
	herodata.MovementTurnRate=result.MovementTurnRate
	for key,value:=range herodata.Abilities{
		//value.AbilityName=abilitiesStr[key]
		GetAbility(value.AbilityName)
		if value!=""{

		}
	}
	//ability1.AbilityName=result.Abilities1
	return herodata
	//err1:=c.Find(bson.M{"matchid": 2271813303}).Select(bson.M{"last_hit.pre_5_count":1}).One(&NewCDotaGameInfo)
}