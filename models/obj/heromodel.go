package obj
import(
	
)

type HeroModel struct {
	HeroName string `json:"hero_name,omitempty"`
	Id uint32 `json:"HeroID,omitempty" bson:"HeroID,omitempty"`
	ImgUrl string `json:"ImgUrl,omitempty" bson:"ImgUrl,omitempty"`
	Abilities []*Ability `json:"Abilities,omitempty"`
	Talent []*Ability `json:"Talent"`
	Role string `json:"Role" bson:"Role"`
	ArmorPhysical int `json:"ArmorPhysical" bson:"ArmorPhysical"`
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
type AllNpcHeroes struct{
	Version string `bson:"Version" json:"Version"`
	Heroes []*NpcHeroes `bson:"Heroes" json:"Heroes"`
}

type NpcHeroes struct{
	Role string `json:"Role" bson:"Role"`
	HeroName string `json:"Name" bson:"Name"`
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
	Abilities18 string `json:"Abilities18,omitempty" bson:"Abilities18,omitempty"`
	Abilities19 string `json:"Abilities19,omitempty" bson:"Abilities19,omitempty"`
	Abilities20 string `json:"Abilities20,omitempty" bson:"Abilities20,omitempty"`
	Abilities21 string `json:"Abilities21,omitempty" bson:"Abilities21,omitempty"`
	Abilities22 string `json:"Abilities22,omitempty" bson:"Abilities22,omitempty"`
	Abilities23 string `json:"Abilities23,omitempty" bson:"Abilities23,omitempty"`
	Abilities24 string `json:"Abilities24,omitempty" bson:"Abilities24,omitempty "`
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
}

type NpcAbilitySpecial struct{
	Damage string `bson:"damage,omitempty" json:"damage,omitempty"`
	BaseDamage string `bson:"base_damage,omitempty" json:"base_damage,omitempty"`
	DamagePerTick string `bson:"damage_per_tick,omitempty" json:"damage_per_tick,omitempty"`
	Interval string `bson:"interval,omitempty" json:"interval,omitempty"`
	Duration string `bson:"duration,omitempty" json:"duration,omitempty"`
	StunDuration string `bson:"stun_duration,omitempty" json:"stun_duration,omitempty"`
	MaxStun string `bson:"max_stun,omitempty" json:"max_stun,omitempty"`
	MinStun string `bson:"min_stun,omitempty" json:"min_stun,omitempty"`
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
	//所有类型伤害加减 小精灵
	AllDamagePec string `bson:"all_damage_pec,omitempty" json:"all_damage_pec,omitempty"`
	HpRegen int `bson:"hp_regen" json:"hp_regen"`//生命恢复量
	//HpRegenPec float32 `bson:"hp_regen_pec" json:"hp_regen_pec"`
	//BonusDamage string `bson:"bonus_damage,omitempty" json:"bonus_damage,omitempty"`
	//BounsAttack string `bson:"BounsAttack,omitempty" json:"BounsAttack,omitempty"`
	//承受魔法伤害百分比 正数就是承受的魔法伤害增加，负数就是承受的魔法伤害降低
	//MagicDamagePec string `bson:"magic_damage_pec,omitempty" json:"magic_damage_pec,omitempty"`
	//所有类型伤害加减 小精灵
	
}
type Ability struct{
	Id string `json:"ID" bson:"ID"`
	Name string `json:"Name" bson:"Name"`
	IconUrl string `json:"IconUrl,omitempty"" Bson:"IconUrl",omitempty"`
	AbilityBehavior string `json:"AbilityBehavior" bson:"AbilityBehavior"`
	DamageType string `json:"AbilityUnitDamageType" bson:"AbilityUnitDamageType"`
	SpellImmunityType string `json:"SpellImmunityType,omitempty" bson:"SpellImmunityType,omitempty"` //技能免疫
	AbilityCastRange string `json:"AbilityCastRange,omitempty" bson:"AbilityCastRange,omitempty"`
	CD string `json:"AbilityCooldown,omitempty" bson:"AbilityCooldown,omitempty"`
	ManaCost string `json:"AbilityManaCost,omitempty" bson:"AbilityManaCost,omitempty"`
	//BaseDamage string `json:"AbilityDamage,omitempty" bson:"AbilityDamage,omitempty"`
	AbilitySpecial NpcAbilitySpecial `json:"AbilitySpecial,omitempty" bson:"AbilitySpecial,omitempty"`
	AbilityDamage string `bson:"AbilityDamage,omitempty" json:"AbilityDamage,omitempty"`
	AbilityDuration string`bson:"AbilityDuration,omitempty" json:"AbilityDuration,omitempty"`
}
type AllNpcAbilities struct{
	Version string `bson:"Version" json:"Version"`
	Abilities []*Ability `bson:"abilities" json:"abilities"`	
}
