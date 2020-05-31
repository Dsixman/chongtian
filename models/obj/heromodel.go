package obj
import(
	
)

type HeroModel struct {
	HeroName string `json:"Name"`
	HeroID uint32 `json:"HeroID" bson:"HeroID"`
	HeroIcon string `json:"HeroIcon" bson:"HeroIcon"`
	StatusHealthRegen float64  `json:"StatusHealthRegen" bson:"StatusHealthRegen"`
	StatusManaRegen float64 `json:"StatusManaRegen" bson:"StatusManaRegen"`
	Ability []*Ability `json:"Ability"`
	Talent []*Ability `json:"Talent"`
	Role string `json:"Role" bson:"Role"`
	ArmorPhysical int `json:"ArmorPhysical" bson:"ArmorPhysical"`
	AttackDamageMin int `json:"AttackDamageMin" bson:"AttackDamageMin"`
	AttackDamageMax int `json:"AttackDamageMax" bson:"AttackDamageMax"`
	AttackRate float64 `json:"AttackRate" bson:"AttackRate"`
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
type AllNpcHeroes struct{
	Version string `bson:"Version" json:"Version"`
	Heroes []*NpcHeroes `bson:"Heroes" json:"Heroes"`
}

type NpcHeroes struct{
	Role string `json:"Role" bson:"Role"`
	HeroName string `json:"Name" bson:"Name"`
	HeroID string `json:"HeroID" bson:"HeroID"`
	Icon string `json:"Icon" bson:"Icon"`
	Ability1 string `json:"Ability1" bson:"Ability1"`
	Ability2 string `json:"Ability2" bson:"Ability2"`
	Ability3 string `json:"Ability3" bson:"Ability3"`
	Ability4 string `json:"Ability4" bson:"Ability4"`
	Ability5 string `json:"Ability5" bson:"Ability5"`
	Ability6 string `json:"Ability6" bson:"Ability6"`
	Ability7 string `json:"Ability7" bson:"Ability7"`
	Ability8 string `json:"Ability8" bson:"Ability8"`
	Ability9 string `json:"Ability9" bson:"Ability9"`
	Ability10 string `json:"Ability10" bson:"Ability10"`
	Ability11 string `json:"Ability11" bson:"Ability11"`
	Ability12 string `json:"Ability12" bson:"Ability12"`
	Ability13 string `json:"Ability13" bson:"Ability13"`
	Ability14 string `json:"Ability14" bson:"Ability14"`
	Ability15 string `json:"Ability15" bson:"Ability15"`
	Ability16 string `json:"Ability16" bson:"Ability16"`
	Ability17 string `json:"Ability17" bson:"Ability17"`
	Ability18 string `json:"Ability18" bson:"Ability18"`
	Ability19 string `json:"Ability19" bson:"Ability19"`
	Ability20 string `json:"Ability20" bson:"Ability20"`
	Ability21 string `json:"Ability21" bson:"Ability21"`
	Ability22 string `json:"Ability22" bson:"Ability22"`
	Ability23 string `json:"Ability23" bson:"Ability23"`
	Ability24 string `json:"Ability24" bson:"Ability24 "`
	ArmorPhysical string `json:"ArmorPhysical" bson:"ArmorPhysical"`
	AttackDamageMin string `json:"AttackDamageMin" bson:"AttackDamageMin"`
	AttackDamageMax string `json:"AttackDamageMax" bson:"AttackDamageMax"`
	AttackRate string `json:"AttackRate" bson:"AttackRate"`
	AttackRange string `json:"AttackRange" bson:"AttackRange"`
	AttributePrimary string `json:"AttributePrimary" bson:"AttributePrimary"`
	AttributeBaseStrength string `json:"AttributeBaseStrength" bson:"AttributeBaseStrength"`
	AttributeStrengthGain string `json:"AttributeStrengthGain" bson:"AttributeStrengthGain"`
	AttributeBaseIntelligence string `json:"AttributeBaseIntelligence" bson:"AttributeBaseIntelligence"`
	AttributeIntelligenceGain string `json:"AttributeIntelligenceGain" bson:"AttributeIntelligenceGain"`
	AttributeBaseAgility string `json:"AttributeBaseAgility" bson:"AttributeBaseAgility"`
	AttributeAgilityGain string `json:"AttributeAgilityGain" bson:"AttributeAgilityGain"`
	MovementSpeed string `json:"MovementSpeed" bson:"MovementSpeed"`
	MovementTurnRate string `json:"MovementTurnRate" bson:"MovementTurnRate"`
	StatusHealthRegen string  `json:"StatusHealthRegen" bson:"StatusHealthRegen"`
	StatusManaRegen string `json:"StatusManaRegen" bson:"StatusManaRegen"` 
}

type NpcAbilitySpecial struct{
	Damage string `bson:"damage" json:"damage"`
	BaseDamage string `bson:"base_damage" json:"base_damage"`
	DamagePerTick string `bson:"damage_per_tick" json:"damage_per_tick"`
	Interval string `bson:"interval" json:"interval"`
	Duration string `bson:"duration" json:"duration"`
	StunDuration string `bson:"stun_duration" json:"stun_duration"`
	MaxStun string `bson:"max_stun" json:"max_stun"`
	MinStun string `bson:"min_stun" json:"min_stun"`
	DamageCount string `bson:"damage_count" json:"damage_count"`
	Armor string `bson:"armor" json:"armor"`
	Coefficient string `bson:"coefficient" json:"coefficient"`
	AttackSpeed string `bson:"attack_speed" json:"attack_speed"`
	MovementSpeed string `bson:"movement_speed" json:"movement_speed"`
	MagicResistance string `bson:"magic_resistance" json:"magic_resistance"` 
	LinkedSpecialBonus string `bson:"LinkedSpecialBonus" json:"LinkedSpecialBonus"`
	AttackPec string `bson:"attack_pec" json:"attack_pec"`//增加百分比的（基础）（龙骑士全攻击力）攻击力
	Instances string `bson:"instances" json:"instances"`//圣堂刺客折光次数
	//承受物理伤害百分比 (美杜莎大招和巫妖冰甲)正数就是承受的物理伤害增加，负数就是承受的物理伤害降低
	PhysicalDamagePec string `bson:"physical_damage_pec" json:"physical_damage_pec"`
	//所有类型伤害加减 小精灵
	AllDamagePec string `bson:"all_damage_pec" json:"all_damage_pec"`
	HpRegen int `bson:"hp_regen" json:"hp_regen"`//生命恢复量
	//HpRegenPec float32 `bson:"hp_regen_pec" json:"hp_regen_pec"`
	//BonusDamage string `bson:"bonus_damage" json:"bonus_damage"`
	//BounsAttack string `bson:"BounsAttack" json:"BounsAttack"`
	//承受魔法伤害百分比 正数就是承受的魔法伤害增加，负数就是承受的魔法伤害降低
	//MagicDamagePec string `bson:"magic_damage_pec" json:"magic_damage_pec"`
	//所有类型伤害加减 小精灵
	
}
type Ability struct{
	Id string `json:"ID" bson:"ID"`
	Name string `json:"Name" bson:"Name"`
	IconUrl string `json:"Icon" bson:"Icon"`
	AbilityBehavior string `json:"AbilityBehavior" bson:"AbilityBehavior"`
	DamageType string `json:"AbilityUnitDamageType" bson:"AbilityUnitDamageType"`
	SpellImmunityType string `json:"SpellImmunityType" bson:"SpellImmunityType"` //技能免疫
	AbilityCastRange string `json:"AbilityCastRange" bson:"AbilityCastRange"`
	CD string `json:"AbilityCooldown" bson:"AbilityCooldown"`
	ManaCost string `json:"AbilityManaCost" bson:"AbilityManaCost"`
	//BaseDamage string `json:"AbilityDamage" bson:"AbilityDamage"`
	AbilitySpecial NpcAbilitySpecial `json:"AbilitySpecial" bson:"AbilitySpecial"`
	AbilityDamage string `bson:"AbilityDamage" json:"AbilityDamage"`
	AbilityDuration string`bson:"AbilityDuration" json:"AbilityDuration"`
}
type AllNpcAbility struct{
	Version string `bson:"Version" json:"Version"`
	Ability []*Ability `bson:"Abilities" json:"Abilities"`	
}
