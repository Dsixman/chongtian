package models

import (

)

var HeroIDName = map[uint32]string{
	1:   "敌法师",
	2:   "斧王",
	3:   "霍乱之源",
	4:   "血魔",
	5:   "冰女",
	6:   "卓尔游侠",
	7:   "撼地者",
	8:   "主宰",
	9:   "米拉娜",
	10:  "水人",
	11:  "影魔",
	12:  "幻影长矛手",
	13:  "帕克",
	14:  "屠夫",
	15:  "剃刀",
	16:  "沙王",
	17:  "风暴之灵",
	18:  "斯温",
	19:  "小小",
	20:  "复仇之魂",
	21:  "风行者",
	22:  "宙斯",
	23:  "昆卡",
	25:  "莉娜",
	26:  "莱恩",
	27:  "暗影萨满",
	28:  "斯拉达",
	29:  "潮汐猎人",
	30:  "巫医",
	31:  "巫妖",
	32:  "力丸",
	33:  "谜团",
	34:  "修补匠",
	35:  "狙击手",
	36:  "瘟疫法师",
	37:  "术士",
	38:  "兽王",
	39:  "痛苦女王",
	40:  "剧毒术士",
	41:  "虚空假面",
	42:  "冥魂大帝",
	43:  "死亡先知",
	44:  "幻影刺客",
	45:  "帕格纳",
	46:  "圣堂刺客",
	47:  "冥界亚龙",
	48:  "露娜",
	49:  "龙骑士",
	50:  "戴泽",
	51:  "发条技师",
	52:  "拉席克",
	53:  "先知",
	54:  "噬魂鬼",
	55:  "黑暗贤者",
	56:  "克林克兹",
	57:  "全能骑士",
	58:  "魅惑魔女",
	59:  "哈斯卡",
	60:  "暗夜魔王",
	61:  "育母蜘蛛",
	62:  "赏金猎人",
	63:  "编织者",
	64:  "杰奇洛",
	65:  "蝙蝠骑士",
	66:  "陈",
	67:  "幽鬼",
	68:  "远古冰魄",
	69:  "末日使者",
	70:  "熊战士",
	71:  "裂魂人",
	72:  "矮人直升机",
	73:  "炼金术士",
	74:  "祈求者",
	75:  "沉默术士",
	76:  "黑鸟",
	77:  "狼人",
	78:  "酒仙",
	79:  "暗影恶魔",
	80:  "德鲁伊",
	81:  "混沌骑士",
	82:  "米波",
	83:  "树精卫士",
	84:  "蓝胖",
	85:  "尸王",
	86:  "拉比克",
	87:  "干扰者",
	88:  "司夜刺客",
	89:  "娜迦海妖",
	90:  "光法",
	91:  "小精灵",
	92:  "维萨吉",
	93:  "斯拉克",
	94:  "美杜莎",
	95:  "巨魔战将",
	96:  "半人马战士",
	97:  "马格纳斯",
	98:  "伐木机",
	99:  "钢背兽",
	100: "巨牙海民",
	101: "天怒",
	102: "亚巴顿",
	103: "上古巨神",
	104: "军团指挥官",
	105: "炸弹人",
	106: "灰烬之灵",
	107: "大地之灵",
	108: "孽主",
	109: "恐怖利刃",
	110: "凤凰",
	111: "神谕者",
	112: "寒冬飞龙",
	113: "天穹守望者",
	114: "大圣",
	119: "邪影芳灵",
	120: "石鳞剑士",
	121:"天涯墨客",
	126:"虚空之灵",
	128:"电炎绝手",
	129:"马尔斯",
}
var HeroNameIcon = map[string]string{
	"敌法师":   "antimage.jpg",
	"斧王":   "axe.jpg",
	"霍乱之源":   "bane.jpg",
	"血魔":   "bloodseeker.jpg",
	"冰女":   "crystalmaiden.jpg",
	"卓尔游侠":  "drowranger.jpg",
	"撼地者":   "earthshaker.jpg",
	"主宰":   "juggernaut.jpg",
	"米拉娜":   "mirana.jpg",
	"水人": "morphling.jpg",
	"影魔": "shadowfiend.jpg",
	"幻影长矛手": "phantomlancer.jpg",
	"帕克": "puck.jpg",
	"屠夫": "pudge.jpg",
	"剃刀": "razor.jpg",
	"沙王": "sandking.jpg",
	"风暴之灵": "stormspirit.jpg",
	"斯温": "sven.jpg",
	"小小": "tiny.jpg",
	"复仇之魂": "vengeful-spirit.jpg",
	"风行者": "windranger.jpg",
	"宙斯": "zeus.jpg",
	"昆卡": "kunkka.jpg",
	"莉娜": "lina.jpg",
	"莱恩": "lion.jpg",
	 "暗影萨满": "shadowshaman.jpg",
	"斯拉达": "slardar.jpg",
	"潮汐猎人": "tidehunter.jpg",
	"巫医": "witchdoctor.jpg",
	 "巫妖": "lich.jpg",
	"力丸": "riki.jpg",
	 "谜团": "enigma.jpg",
	"修补匠": "tinker.jpg",
	"狙击手": "sniper.jpg",
	"瘟疫法师": "necrophos.jpg",
	"术士": "warlock.jpg",
	"兽王": "beastmaster.jpg",
	"痛苦女王": "queenofpain.jpg",
	 "剧毒术士": "venomancer.jpg",
	"虚空假面": "gyrocopter.jpg",
	"冥魂大帝": "wraithking.jpg",
	"死亡先知": "deathprophet.jpg",
	"幻影刺客": "phantomassassin.jpg",
	"帕格纳": "pugna.jpg",
	"圣堂刺客": "templarassassin.jpg",
	 "冥界亚龙": "viper.jpg",
	 "露娜": "luna.jpg",
	 "龙骑士": "dragonknight.jpg",
	"戴泽": "dazzle.jpg",
	"发条技师": "clockwerk.jpg",
	"拉席克": "leshrac.jpg",
	"先知": "naturesprophet.jpg",
	"噬魂鬼": "lifestealer.jpg",
	"黑暗贤者": "dark-seer.jpg",
	"克林克兹": "clinkz.jpg",
	"全能骑士": "omniknight.jpg",
	"魅惑魔女": "enchantress.jpg",
	"哈斯卡": "huskar.jpg",
	"暗夜魔王":"nightstalker.jpg",
	"育母蜘蛛": "broodmother.jpg",
	"赏金猎人": "bountyhunter.jpg",
	"编织者": "weaver.jpg",
	"杰奇洛": "jakiro.jpg",
	"蝙蝠骑士": "batrider.jpg",
	"陈": "chen.jpg",
	"幽鬼": "spectre.jpg",
	"远古冰魄": "ancient.jpg",
	"末日使者": "doom.jpg",
	"熊战士": "ursa.jpg",
	"裂魂人": "spiritbreaker.jpg",
	"矮人直升机": "gyrocopter.jpg",
	"炼金术士": "alchemist.jpg",
	"祈求者": "invoker.jpg",
	"沉默术士": "silencer.jpg",
	"黑鸟": "outworlddevourer.jpg",
	"狼人": "lycan.jpg",
	"酒仙": "brewmaster.jpg",
	"暗影恶魔": "shadowdemon.jpg",
	"德鲁伊": "lone-druid.jpg",
	"混沌骑士": "chaosknight.jpg",
	"米波": "meepo.jpg",
	"树精卫士": "treantprotector.jpg",
	"蓝胖": "ogremagi.jpg",
	"尸王": "undying.jpg",
	"拉比克": "rubick.jpg",
	"干扰者": "disruptor.jpg",
	"司夜刺客": "nyx-assassin.jpg",
	"娜迦海妖": "nagasiren.jpg",
	"光法": "keeperofthelight.jpg",
	"小精灵": "io.jpg",
	"维萨吉": "visage.jpg",
	"斯拉克": "slark.jpg",
	"美杜莎": "medusa.jpg",
	"巨魔战将": "trollwarlord.jpg",
	"半人马战士": "centaurwarrunner.jpg",
	"马格纳斯": "magnus.jpg",
	"伐木机": "timbersaw.jpg",
	"钢背兽": "bristleback.jpg",
	"巨牙海民": "tusk.jpg",
	 "天怒": "skywrathmage.jpg",
	"亚巴顿": "abaddon.jpg",
	"上古巨神": "eldertitan.jpg",
	"军团指挥官": "legioncommander.jpg",
	"炸弹人": "techies.jpg",
	"灰烬之灵": "emberspirit.jpg",
	"大地之灵": "earthspirit.jpg",
	"孽主": "underlord.jpg",
	"恐怖利刃": "terrorblade.jpg",
	 "凤凰": "phoenix.jpg",
	"神谕者": "oracle.jpg",
	"寒冬飞龙": "winterwyvern.jpg",
	"天穹守望者": "arcwarden.jpg",
	"大圣": "monkeyking.jpg",
	"邪影芳灵": "dark-willow.jpg",
	"石鳞剑士": "pangolier.jpg",
	"天涯墨客":"grimstroke.jpg",
	"电炎绝手":"snapfire.jpg",
	"虚空之灵":"void-spirit.jpg",
	"马尔斯":"mars.jpg",
}
var HeroNameID = map[string]uint32{
	"敌法师":1   ,
	"斧王":2,
	"霍乱之源":3,
	"血魔":4,
	"冰女":5,
	"卓尔游侠":6,
	"撼地者": 7,
	"主宰": 8,
	"米拉娜": 9,
	"水人": 10,
	"影魔": 11,
	"幻影长矛手": 12,
	"帕克": 13,
	"屠夫": 14,
	"剃刀": 15,
	"沙王": 16,
	"风暴之灵": 17,
	"斯温": 18,
	"小小": 19,
	"复仇之魂": 20,
	"风行者": 21,
	"宙斯": 22,
	"昆卡": 23,
	"莉娜": 25,
	"莱恩": 26,
	"暗影萨满": 27,
	"斯拉达": 28,
	"潮汐猎人": 29,
	"巫医": 30,
	"巫妖": 31,
	"力丸": 32,
	"谜团": 33,
	"修补匠": 34,
	"狙击手": 35,
	"瘟疫法师": 36,
	"术士": 37,
	"兽王": 38,
	"痛苦女王": 39,
	"剧毒术士": 40,
	"虚空假面": 41,
	"冥魂大帝": 42,
	"死亡先知": 43,
	"幻影刺客": 44,
	"帕格纳": 45,
	"圣堂刺客": 46,
	"冥界亚龙": 47,
	"露娜": 48,
	"龙骑士": 49,
	"戴泽": 50,
	"发条技师":51,
	"拉席克": 52,
	"先知": 53,
	"噬魂鬼": 54,
	"黑暗贤者": 55,
	"克林克兹": 56,
	"全能骑士": 57,
	"魅惑魔女": 58,
	"哈斯卡": 59,
	"暗夜魔王": 60,
	"育母蜘蛛": 61,
	"赏金猎人": 62,
	"编织者": 63,
	"杰奇洛": 64,
	"蝙蝠骑士": 65,
	"陈": 66,
	"幽鬼": 67,
	"远古冰魄": 68,
	"末日使者": 69,
	"熊战士": 70,
	"裂魂人": 71,
	"矮人直升机": 72,
	"炼金术士": 73,
	"祈求者": 74,
	"沉默术士": 75,
	"黑鸟": 76,
	"狼人": 77,
	"酒仙": 78,
	"暗影恶魔": 79,
	"德鲁伊" :80,
	"混沌骑士" :81,
	"米波" :82,
	"树精卫士" :83,
	"蓝胖" :84,
	"尸王" :85,
	"拉比克" :86,
	"干扰者" :87,
	"司夜刺客" :88,
	"娜迦海妖" :89,
	"光法" :90,
	"小精灵" :91,
	"维萨吉" :92,
	"斯拉克" :93,
	"美杜莎" :94,
	"巨魔战将" :95,
	"半人马战士" :96,
	"马格纳斯" :97,
	"伐木机":98,
	"钢背兽" :99,
	"巨牙海民" :100,
	"天怒" :101,
	"亚巴顿" :102,
	"上古巨神" :103,
	"军团指挥官" :104,
	"炸弹人" :105,
	"灰烬之灵" :106,
	"大地之灵" :107,
	"孽主":108,
	"恐怖利刃":109,
	"凤凰":110,
	"神谕者":111,
	"寒冬飞龙":112,
	"天穹守望者":113,
	"大圣":114,
	"邪影芳灵":119,
	"石鳞剑士":120,
	"天涯墨客":121,
	"虚空之灵":126,
	"电炎绝手":128,
	"马尔斯":129,
}
var HeroIDENName = map[uint32]string{
	1:   "npc_dota_hero_antimage",
	2:   "npc_dota_hero_axe",
	3:   "npc_dota_hero_bane",
	4:   "npc_dota_hero_bloodseeker",
	5:   "npc_dota_hero_crystal_maiden",
	6:   "npc_dota_hero_drow_ranger",
	7:   "npc_dota_hero_earthshaker",
	8:   "npc_dota_hero_juggernaut",
	9:   "npc_dota_hero_mirana",
	10:  "npc_dota_hero_morphling",
	11:  "npc_dota_hero_nevermore",
	12:  "npc_dota_hero_phantom_lancer",
	13:  "npc_dota_hero_puck",
	14:  "npc_dota_hero_pudge",
	15:  "npc_dota_hero_razor",
	16:  "npc_dota_hero_sand_king",
	17:  "npc_dota_hero_storm_spirit",
	18:  "npc_dota_hero_sven",
	19:  "npc_dota_hero_tiny",
	20:  "npc_dota_hero_vengefulspirit",
	21:  "npc_dota_hero_windrunner",
	22:  "npc_dota_hero_zuus",
	23:  "npc_dota_hero_kunkka",
	25:  "npc_dota_hero_lina",
	26:  "npc_dota_hero_lion",
	27:  "npc_dota_hero_shadow_shaman",
	28:  "npc_dota_hero_slardar",
	29:  "npc_dota_hero_tidehunter",
	30:  "npc_dota_hero_witch_doctor",
	31:  "npc_dota_hero_lich",
	32:  "npc_dota_hero_riki",
	33:  "npc_dota_hero_enigma",
	34:  "npc_dota_hero_tinker",
	35:  "npc_dota_hero_sniper",
	36:  "npc_dota_hero_necrolyte",
	37:  "npc_dota_hero_warlock",
	38:  "npc_dota_hero_beastmaster",
	39:  "npc_dota_hero_queenofpain",
	40:  "npc_dota_hero_venomancer",
	41:  "npc_dota_hero_faceless_void",
	42:  "npc_dota_hero_skeleton_king",
	43:  "npc_dota_hero_death_prophet",
	44:  "npc_dota_hero_phantom_assassin",
	45:  "npc_dota_hero_pugna",
	46:  "npc_dota_hero_templar_assassin",
	47:  "npc_dota_hero_viper",
	48:  "npc_dota_hero_luna",
	49:  "npc_dota_hero_dragon_knight",
	50:  "npc_dota_hero_dazzle",
	51:  "npc_dota_hero_rattletrap",
	52:  "npc_dota_hero_leshrac",
	53:  "npc_dota_hero_furion",
	54:  "npc_dota_hero_life_stealer",
	55:  "npc_dota_hero_dark_seer",
	56:  "npc_dota_hero_clinkz",
	57:  "npc_dota_hero_omniknight",
	58:  "npc_dota_hero_enchantress",
	59:  "npc_dota_hero_huskar",
	60:  "npc_dota_hero_night_stalker",
	61:  "npc_dota_hero_broodmother",
	62:  "npc_dota_hero_bounty_hunter",
	63:  "npc_dota_hero_weaver",
	64:  "npc_dota_hero_jakiro",
	65:  "npc_dota_hero_batrider",
	66:  "npc_dota_hero_chen",
	67:  "npc_dota_hero_spectre",
	68:  "npc_dota_hero_ancient_apparition",
	69:  "npc_dota_hero_doom_bringer",
	70:  "npc_dota_hero_ursa",
	71:  "npc_dota_hero_spirit_breaker",
	72:  "npc_dota_hero_gyrocopter",
	73:  "npc_dota_hero_alchemist",
	74:  "npc_dota_hero_invoker",
	75:  "npc_dota_hero_silencer",
	76:  "npc_dota_hero_obsidian_destroyer",
	77:  "npc_dota_hero_lycan",
	78:  "npc_dota_hero_brewmaster",
	79:  "npc_dota_hero_shadow_demon",
	80:  "npc_dota_hero_lone_druid",
	81:  "npc_dota_hero_chaos_knight",
	82:  "npc_dota_hero_meepo",
	83:  "npc_dota_hero_treant",
	84:  "npc_dota_hero_ogre_magi",
	85:  "npc_dota_hero_undying",
	86:  "npc_dota_hero_rubick",
	87:  "npc_dota_hero_disruptor",
	88:  "npc_dota_hero_nyx_assassin",
	89:  "npc_dota_hero_naga_siren",
	90:  "npc_dota_hero_keeper_of_the_light",
	91:  "npc_dota_hero_wisp",
	92:  "npc_dota_hero_visage",
	93:  "npc_dota_hero_slark",
	94:  "npc_dota_hero_medusa",
	95:  "npc_dota_hero_troll_warlord",
	96:  "npc_dota_hero_centaur",
	97:  "npc_dota_hero_magnataur",
	98:  "npc_dota_hero_shredder",
	99:  "npc_dota_hero_bristleback",
	100: "npc_dota_hero_tusk",
	101: "npc_dota_hero_skywrath_mage",
	102: "npc_dota_hero_abaddon",
	103: "npc_dota_hero_elder_titan",
	104: "npc_dota_hero_legion_commander",
	105: "npc_dota_hero_techies",
	106: "npc_dota_hero_ember_spirit",
	107: "npc_dota_hero_earth_spirit",
	108: "npc_dota_hero_abyssal_underlord",
	109: "npc_dota_hero_terrorblade",
	110: "npc_dota_hero_phoenix",
	111: "npc_dota_hero_oracle",
	112: "npc_dota_hero_winter_wyvern",
	113: "npc_dota_hero_arc_warden",
	114: "npc_dota_hero_monkey_king",
	119: "npc_dota_hero_dark_willow",
	120: "npc_dota_hero_pangolier",
	121:"npc_dota_hero_grimstroke",
	126:"npc_dota_hero_void_spirit",
	128: "npc_dota_hero_snapfire",
	129:"npc_dota_hero_mars",
}
var HeroENNameID = map[string]uint32{
	"npc_dota_hero_antimage" 	:1   ,
	"npc_dota_hero_axe" :2   ,
	"npc_dota_hero_bane" :3,
 	"npc_dota_hero_bloodseeker" :4,
	"npc_dota_hero_crystal_maiden" :5,
	"npc_dota_hero_drow_ranger" :6,
	"npc_dota_hero_earthshaker":7,
	"npc_dota_hero_juggernaut":8,
	"npc_dota_hero_mirana":9,
	"npc_dota_hero_morphling":10,
	"npc_dota_hero_nevermore":11,
	"npc_dota_hero_phantom_lancer":12,
	"npc_dota_hero_puck":13,
	"npc_dota_hero_pudge":14,
	"npc_dota_hero_razor":15,
	"npc_dota_hero_sand_king":16,
	"npc_dota_hero_storm_spirit":17,
	"npc_dota_hero_sven":18,
	"npc_dota_hero_tiny":19,
	"npc_dota_hero_vengefulspirit":20,
	"npc_dota_hero_windrunner":21,
	"npc_dota_hero_zuus":22,
	"npc_dota_hero_kunkka":23,
	 "npc_dota_hero_lina":25,
	"npc_dota_hero_lion":26,
	 "npc_dota_hero_shadow_shaman":27,
	"npc_dota_hero_slardar":28,
	"npc_dota_hero_tidehunter":29,
	 "npc_dota_hero_witch_doctor":30,
	  "npc_dota_hero_lich" :31,
	 "npc_dota_hero_riki":32,
	 "npc_dota_hero_enigma":33,
	 "npc_dota_hero_tinker":34,
	"npc_dota_hero_sniper":35,
	"npc_dota_hero_necrolyte":36,
	"npc_dota_hero_warlock":37,
	"npc_dota_hero_beastmaster":38,
	"npc_dota_hero_queenofpain":39,
	"npc_dota_hero_venomancer":40,
	"npc_dota_hero_faceless_void":41,
	"npc_dota_hero_skeleton_king":42,
	"npc_dota_hero_death_prophet":43,
	 "npc_dota_hero_phantom_assassin":44,
	 "npc_dota_hero_pugna":45,
	 "npc_dota_hero_templar_assassin":46,
	"npc_dota_hero_viper":47,
	"npc_dota_hero_luna":48,
	"npc_dota_hero_dragon_knight":49,
	"npc_dota_hero_dazzle":50,
	"npc_dota_hero_rattletrap":51,
	"npc_dota_hero_leshrac":52,
	"npc_dota_hero_furion":53,
	 "npc_dota_hero_life_stealer":54,
	"npc_dota_hero_dark_seer":55,
	 "npc_dota_hero_clinkz":56,
	 "npc_dota_hero_omniknight":57,
	 "npc_dota_hero_enchantress":58,
	 "npc_dota_hero_huskar":59,
	"npc_dota_hero_night_stalker":60,
	"npc_dota_hero_broodmother":61,
	"npc_dota_hero_bounty_hunter":62,
	 "npc_dota_hero_weaver":63,
	"npc_dota_hero_jakiro":64,
	"npc_dota_hero_batrider":65,
	"npc_dota_hero_chen":66,
	"npc_dota_hero_spectre":67,
	 "npc_dota_hero_ancient_apparition":68,
	"npc_dota_hero_doom_bringer":69,
	 "npc_dota_hero_ursa":70,
	 "npc_dota_hero_spirit_breaker":71,
	 "npc_dota_hero_gyrocopter":72,
	"npc_dota_hero_alchemist":73,
	 "npc_dota_hero_invoker":74,
	 "npc_dota_hero_silencer":75,
	 "npc_dota_hero_obsidian_destroyer":76,
	 "npc_dota_hero_lycan":77,
	"npc_dota_hero_brewmaster":78,
	"npc_dota_hero_shadow_demon":79,
	"npc_dota_hero_lone_druid":80,
	 "npc_dota_hero_chaos_knight":81,
	 "npc_dota_hero_meepo":82,
	 "npc_dota_hero_treant":83,
	  "npc_dota_hero_ogre_magi":84,
	"npc_dota_hero_undying":85,
	 "npc_dota_hero_rubick":86,
	"npc_dota_hero_disruptor":87,
	 "npc_dota_hero_nyx_assassin":88,
	 "npc_dota_hero_naga_siren":89,
	"npc_dota_hero_keeper_of_the_light":90,
	"npc_dota_hero_wisp":90,
	"npc_dota_hero_visage":92,
	"npc_dota_hero_slark":93,
	"npc_dota_hero_medusa":94,
	"npc_dota_hero_troll_warlord":95,
	"npc_dota_hero_centaur":96,
	"npc_dota_hero_magnataur":97,
	"npc_dota_hero_shredder":98,
	"npc_dota_hero_bristleback":99,
	"npc_dota_hero_tusk":100,
	"npc_dota_hero_skywrath_mage":101,
	 "npc_dota_hero_abaddon":102,
	 "npc_dota_hero_elder_titan":103,
	"npc_dota_hero_legion_commander":104,
	"npc_dota_hero_techies":105,
	 "npc_dota_hero_ember_spirit":106,
	"npc_dota_hero_earth_spirit":107,
	"npc_dota_hero_abyssal_underlord":108,
	 "npc_dota_hero_terrorblade":109,
	 "npc_dota_hero_phoenix":110,
	 "npc_dota_hero_oracle":111,
	 "npc_dota_hero_winter_wyvern":112,
	"npc_dota_hero_arc_warden":113,
	 "npc_dota_hero_monkey_king":114,
	 "npc_dota_hero_dark_willow":119,
	"npc_dota_hero_pangolier":120,
	"npc_dota_hero_grimstroke":121,
	"npc_dota_hero_void_spirit":126,
	"npc_dota_hero_snapfire":128,
	"npc_dota_hero_mars":129,
}
/*var HeroIDIcon = map[int64]string{
	1:   "antimage.jpg",
	2:   "axe.jpg",
	3:   "bane.jpg",
	4:   "bloodseeker.jpg",
	5:   "crystalmaiden.jpg",
	6:  "drowranger.jpg",
	7:   "earthshaker.jpg",
	8:   "juggernaut.jpg",
	9:   "mirana.jpg",
	10: "morphling.jpg",
	11: "shadowfiend.jpg",
	12: "phantomlancer.jpg",
	13: "puck.jpg",
	14: "pudge.jpg",
	15: "razor.jpg",
	16: "sandking.jpg",
	17: "stormspirit.jpg",
	18: "sven.jpg",
	19: "tiny.jpg",
	20: "vengeful-spirit.jpg",
	21: "windranger.jpg",
	22: "zeus.jpg",
	23: "kunkka.jpg",
	25: "lina.jpg",
	26: "lion.jpg",
	27: "shadowshaman.jpg",
	28: "slardar.jpg",
	29: "tidehunter.jpg",
	30: "witchdoctor.jpg",
	31: "lich.jpg",
	32: "riki.jpg",
	33: "enigma.jpg",
	34: "tinker.jpg",
	35: "sniper.jpg",
	36: "necrophos.jpg",
	37: "warlock.jpg",
	38: "beastmaster.jpg",
	39: "queenofpain.jpg",
	40: "venomancer.jpg",
	41: "gyrocopter.jpg",
	42: "wraithking.jpg",
	43: "deathprophet.jpg",
	44: "phantomassassin.jpg",
	45: "pugna.jpg",
	46: "templarassassin.jpg",
	47: "viper.jpg",
	48: "luna.jpg",
	49: "dragonknight.jpg",
	50: "dazzle.jpg",
	51: "clockwerk.jpg",
	52: "leshrac.jpg",
	53: "naturesprophet.jpg",
	54: "lifestealer",
	55: "dark-seer.jpg",
	56: "clinkz.jpg",
	57: "omniknight.jpg",
	58: "enchantress.jpg",
	59: "huskar.jpg",
	60: "nightstalker.jpg",
	61: "broodmother.jpg",
	62: "bountyhunter.jpg",
	63: "weaver.jpg",
	64: "jakiro.jpg",
	65: "batrider.jpg",
	66: "chen.jpg",
	67: "spectre.jpg",
	68: "ancient.jpg",
	69: "doom.jpg",
	70: "ursa",
	71: "spiritbreaker.jpg",
	72: "gyrocopter.jpg",
	73: "alchemist.jpg",
	74: "invoker.jpg",
	75: "silencer.jpg",
	76: "outworld-devourer.jpg",
	77: "lycan.jpg",
	78: "brewmaster.jpg",
	79: "shadowdemon.jpg",
	80: "lone-druid.jpg",
	81: "chaosknight.jpg",
	82: "meepo.jpg",
	83: "treant-protector.jpg",
	84: "ogremagi.jpg",
	85: "undying.jpg",
	86: "rubick.jpg",
	87: "disruptor.jpg",
	88: "nyx-assassin.jpg",
	89: "naga-siren.jpg",
	90: "keeperofthelight.jpg",
	91: "io.jpg",
	92: "visage.jpg",
	93: "slark.jpg",
	94: "medusa.jpg",
	95: "trollwarlord.jpg",
	96: "centaurwarrunner.jpg",
	97: "magnus.jpg",
	98: "timbersaw.jpg",
	99: "bristleback.jpg",
	100: "tusk.jpg",
	101: "skywrathmage.jpg",
	102: "abaddon.jpg",
	103: "eldertitan.jpg",
	104: "legioncommander.jpg",
	105: "techies.jpg",
	106: "emberspirit.jpg",
	107: "earthspirit.jpg",
	108: "underlord.jpg",
	109: "terrorblade.jpg",
	110: "phoenix.jpg",
	111: "oracle.jpg",
	112: "winterwyvern.jpg",
	113: "arcwarden.jpg",
	114: "monkey-king.jpg",
	119: "dark-willow.jpg",
	120: "pangolier.jpg",
	121: "grimstroke.jpg",
	126:"void-spirit.jpg",
	128:"snapfire.jpg",
	129:"mars.jpg",
}*/
//散慧对剑 慧夜对剑 否决 永恒之盘 圣洁吊坠
var ItemIDIcon = map[int64]string{
	//大飞鞋 ，银月之晶，陨星锤 ，肉山盾，刷新碎片，精灵布带,凝魂之泪，慧光
	1:  "blink-dagger.jpg",           //跳刀
	2:  "blades-of-attack.jpg",       //攻击之爪
	3:  "broadsword.jpg",             //阔剑
	4:  "chainmail.jpg",              //锁子甲
	5:  "claymore.jpg",               //大剑
	6:  "helm_of_iron_will.jpg",      //铁意头盔
	7:  "javelin.jpg",                //标枪
	8:  "mithril-hammer.jpg",         //秘银锤
	9:  "platemail.jpg",              //板甲
	10: "quarterstaff.jpg",           //短棍
	11: "quelling-blade.jpg",         //补刀斧
	12: "ring_of_protection.jpg",     //守护戒指
	13: "gauntlets-of-strength.jpg",  //力量手套
	14: "slippers-of-agility.jpg",    //敏捷便鞋
	15: "mantle-of-intelligence.jpg", //智力斗篷
	16: "iron-branch.jpg",            //树枝
	17: "belt-of-strength.jpg",       //力量腰带
	18: "band-of-elvenskin.jpg",      //精灵布带
	19: "robe-of-the-magi",           //法师长袍
	20: "circlet.jpg",                //圆环
	21: "ogre-axe.jpg",               //食人魔之斧
	22: "blade-of-alacrity.jpg",      //欢欣之刃
	23: "staff-of-wizardry.jpg",      //魔力法杖
	24: "ultimate-orb.jpg",           //极限法球
	25: "gloves-of-haste.jpg",        //加速手套
	26: "morbid-mask.jpg",            //吸血面具
	27: "ring-of-regen.jpg",          //回复戒指
	28: "sages-mask.jpg",             //艺人面罩
	29: "boots-of-speed.jpg",         //速度之鞋
	30: "gem-of-true-sight.jpg",      //珍视宝石
	31: "cloak.jpg",                  //抗魔斗篷
	32: "talisman-of-evasion.jpg",    //闪避护符
	33: "cheese.jpg",                 //奶酪
	34: "magic-stick.jpg",            //魔棒

	36: "magic-wand.jpg",         //大魔棒
	37: "ghost-scepter.jpg",      //幽魂权杖
	38: "clarity.jpg",            //净化药水
	39: "healing-salve.jpg",      //大药
	40: "dust-of-appearance.jpg", //显影之尘
	41: "bottle.jpg",             //瓶子
	42: "observer-ward.jpg",      //假眼
	43: "sentry-ward.jpg",        //真眼
	44: "tango.jpg",              //吃数
	45: "animal-courier.jpg",     //动物信使
	46: "town-portal-scrol.jpg",  //TP
	48: "boots-of-travel.jpg",    //飞鞋

	50: "phase-boots.jpg",      //相位鞋
	51: "demon-edge.jpg",       //恶魔刀锋
	52: "eaglesong.jpg",        //鹰歌弓
	53: "reaver.jpg",           //掠夺
	54: "sacred-relic.jpg",     //圣者遗物
	55: "hyperstone.jpg",       //振奋宝石
	56: "ring-of-health.jpg",   //治疗指环
	57: "void-stone.jpg",       //虚无宝石
	58: "mystic-staff.jpg",     //神秘法杖
	59: "energy-booster.jpg",   //能量之球
	60: "point-booster.jpg",    //精气之球
	61: "vitality-booster.jpg", //活力之球
	63: "power-treads.jpg",     //假腿

	65: "hand-of-midas.jpg", //点金手

	67:  "oblivion-staff.jpg",           //空明帐
	69:  "perseverance.jpg",             //坚韧球
	71:"poor-mans-shield.jpg",//穷鬼盾
	73:  "bracer.jpg",                   //护腕
	75:  "wraith-band.jpg",              //怨灵系带
	77:  "null-talisman.jpg",            //空灵挂件
	79:  "mekansm.jpg",                  //梅肯
	81:  "vladmirs-offering.jpg",        //祭品
	84:  "flying_courier_lg.png",        //动物信使
	86:  "buckler.jpg",                  //玄冥盾牌
	88:  "ring-of-basilius.jpg",         //圣殿
	90:  "pipe-of-insight.jpg",          //笛子
	92:  "urn-of-shadows.jpg",           //骨灰
	94:  "headdress.jpg",                //回复头巾
	96:  "scythe-of-vyse.jpg",           //羊刀
	98:  "orchid-malevolence.jpg",       //紫苑
	250: "bloodthorn.jpg",               //血棘
	252: "echo-sabre.jpg",               //回音战刃
	100: "euls-scepter-of-divinity.jpg", //吹风
	232: "aether-lens.jpg",              //以太之镜
	102: "force-staff.jpg",              //推推
	263: "hurricane-pike.jpg",           //大推推
	104: "dagon.jpg",                    //大根
	201: "dagon-level-2.jpg",
	202: "dagon-level-3.jpg",
	203: "dagon-level-4.jpg",
	204: "dagon-level-5.jpg",
	106: "necronomicon.jpg", //死灵书
	193: "necronomicon-level-2.jpg",
	194: "necronomicon-level-3.jpg",
	108: "aghanims-scepter.jpg",        //A帐
	110: "refresher-orb.jpg",           //刷新球
	117: "aegis-of-the-immortal.jpg",   //肉山盾
	112: "assault-cuirass.jpg",         //强袭装甲
	114: "heart-of-tarrasque.jpg",      //龙心
	116: "black-king-bar.jpg",          //BKB
	118: "",                            //
	119: "shivas-guard.jpg",            //西瓦
	121: "bloodstone.jpg",              //血精石
	123: "linkens-sphere.jpg",          //林肯
	226: "lotus-orb.jpg",               //莲花
	125: "vanguard.jpg",                //先锋盾
	242: "crimson-guard.jpg",           //赤红甲
	127: "blade-mail.jpg",              //刃甲
	129: "soul-booster.jpg",            //镇魂石
	131: "hood-of-defiance.jpg",        //挑战
	133: "divine-rapier.jpg",           //圣剑
	135: "monkey-king-bar.jpg",         //金箍棒
	137: "radiance.jpg",                //辉耀
	139: "butterfly.jpg",               //蝴蝶
	141: "daedalus.jpg",                //大炮
	143: "skull-basher.jpg",            //碎颅锤
	145: "battle-fury.jpg",             //狂战斧
	147: "manta-style.jpg",             //分身斧
	149: "crystalys.jpg",               //水晶剑
	236: "dragon-lance.jpg",            //魔龙枪
	151: "armlet-of-mordiggian.jpg",    //臂章
	152: "shadow-blade.jpg",            //隐刀
	249: "silver-edge.jpg",             //大隐刀，白银之锋
	154: "sange-and-yasha.jpg",         //双刀
	156: "satanic.jpg",                 //撒旦
	158: "mjollnir.jpg",                //大电锤
	160: "eye-of-skadi.jpg",            //冰眼
	162: "sange.jpg",                   //散华
	164: "helm-of-the-dominator.jpg",   //支配
	166: "maelstrom.jpg",               //电锤
	168: "desolator.jpg",               //暗灭
	170: "yasha.jpg",                   //单刀夜叉
	172: "mask-of-madness.jpg",         //疯脸
	174: "diffusal-blade.jpg",          //散失
	196: "diffusal-blade-2.jpg",        //
	176: "ethereal-blade.jpg",          //虚灵刀
	178: "soul-ring.jpg",               //魂戒
	180: "arcane-boots.jpg",            //秘法
	235: "octarine-core.jpg",           //玲珑心
	181: "orb-of-venom.jpg",            //毒球
	240: "blight-stone.jpg",            //枯萎之石
	185: "drum-of-endurance.jpg",       //战鼓
	187: "medallion-of-courage.jpg",    //勇气勋章
	229: "solar-crest.jpg",             //炎阳纹章
	188: "smoke-of-deceit.jpg",         //雾
	257: "tome-of-knowledge.jpg",       //经验书
	190: "veil-of-discord.jpg",         //纷争面纱
	231: "guardian-greaves.jpg",        //大鞋
	206: "rod-of-atos.jpg",             //阿托斯
	239: "",                            //打野爪
	208: "abyssal-blade.jpg",           //深渊之刃
	210: "heavens-halberd.jpg",         //天堂
	212: "ring-of-aquila.jpg",          //天鹰
	214: "tranquil-boots.jpg",          //绿鞋
	215: "shadow-amulet.jpg",           //暗影护符
	254: "glimmer-cape.jpg",            //微光披风
	256: "aeon-disk.jpg",               //免死金牌永恒纸盘
	225: "nullifier.jpg",               //否决
	267: "spirit-vessel.jpg",           //大骨灰
	244: "wind-lace.jpg",               //风灵之纹
	259: "kaya.jpg",                    //慧光
	260: "refresher-shard.jpg",         //刷新碎片
	223: "meteor-hammer.jpg",           //陨星锤
	220: "boots-of-travel-level-2.jpg", //大飞鞋
	265: "infused-raindrops.jpg",       //凝魂之泪
	247: "moon-shard.jpg",              //银月之晶
	273: "kaya-and-sange.jpg",//散慧对剑
	277:"yasha-and-kaya.jpg",//慧夜对剑
	287:"keen-optic.jpg",//基恩镜片
	288:"grove-bow.jpg", //林野长弓
	289:"quickening-charm",//加速护符
	290:"philosophers-stone.jpg",//贤者之石
	297:"vampire-fangs.jpg",//吸血鬼獠牙
	298:"craggy-coat",//崎岖外衣
	299:"greater-faerie-fire",//高级仙灵之火
	300:"timeless-relic.jpg",//永恒遗物
	304:"ironwood-tree.jpg",//铁树之木
	306:"pupils-gift.jpg",//学徒之礼
	308:"repair-kit.jpg",//维修器具
	309:"mind-breaker.jpg",//智灭

	311:"spell-prism.jpg",//法术棱镜

	325:"princes-knife.jpg",//亲王短刀
	326:"spider-legs.jpg",//网虫腿

	330:"witless-shako.jpg",//无知小帽
	331:"vambrace.jpg",//臂甲
	334:"imp-claw.jpg",//魔童之刃
	335:"flicker.jpg",//闪灵
	336:"telescope.jpg",//望远镜
	
	349:"arcane-ring.jpg",//奥数指环
	357:"nether-shawl.jpg",//幽冥披巾
	355:"broom-handle.jpg",//扫把帚
	358:"dragon-scale.jpg",//炎龙之鳞
	
	354:"ocean-heart.jpg",//海洋之心
	359:"essence-ring.jpg",//精华指环
	356:"trusty-shovel.jpg",//可靠铁铲
	360:"clumsy-net.jpg",//渔网
	
	378:"orb-of-destruction.jpg",//毁灭灵球
	

	361:"enchanted-quiver.jpg",//魔力箭袋
	362:"ninja-gear.jpg",//忍者用具
	365:"magic-lamp.jpg",//魔力明灯
	364:"havoc-hammer.jpg",//浩劫巨锤
	375:"faded-broach.jpg",//暗淡胸针
	377:"minotaur-horn.jpg",//灵犀角
	379:"the-leveller.jpg",//平世剑
	381:"titan-sliver.jpg",//巨神残铁

	

	

	363:"illusionists-cape.jpg",//幻术师披风
	
	376:"paladin-sword.jpg",//骑士剑
}

