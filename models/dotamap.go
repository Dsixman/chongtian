package models

import (
//"fmt"
)

var HeroIDName = map[uint32]string{
	1:   "敌法师",
	2:   "斧王",
	3:   "霍乱之源",
	4:   "血魔",
	5:   "冰女",
	6:   "黑暗游侠",
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
	129:"马尔斯",
}
var HeroNameID = map[string]uint32{
	"敌法师":1   ,
	"斧王":2,
	"霍乱之源":3,
	"血魔":4,
	"冰女":5,
	"黑暗游侠":6,
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
	129:"npc_dota_hero_mars",
	127: "npc_dota_hero_target_dummy",
}
var HeroIDIcon = map[int64]string{
	1:   "../static/img/hero/antimage.jpg",
	2:   "../static/img/hero/axe.jpg",
	3:   "../static/img/hero/bane.jpg",
	4:   "../static/img/hero/bloodseeker.jpg",
	5:   "../static/img/hero/crystalmaiden.jpg",
	6:  "../static/img/hero/drowranger.jpg",
	7:   "../static/img/hero/earthshaker.jpg",
	8:   "../static/img/hero/juggernaut.jpg",
	9:   "../static/img/hero/mirana.jpg",
	10: "../static/img/hero/morphling.jpg",
	11: "../static/img/hero/shadowfiend.jpg",
	12: "../static/img/hero/phantomlancer.jpg",
	13: "../static/img/hero/puck.jpg",
	14: "../static/img/hero/pudge.jpg",
	15: "../static/img/hero/razor.jpg",
	16: "../static/img/hero/sandking.jpg",
	17: "../static/img/hero/stormspirit.jpg",
	18: "../static/img/hero/sven.jpg",
	19: "../static/img/hero/tiny.jpg",
	20: "../static/img/hero/vengeful-spirit.jpg",
	21: "../static/img/hero/windranger.jpg",
	22: "../static/img/hero/zeus.jpg",
	23: "../static/img/hero/kunkka.jpg",
	25: "../static/img/hero/lina.jpg",
	26: "../static/img/hero/lion.jpg",
	27: "../static/img/hero/shadowshaman.jpg",
	28: "../static/img/hero/slardar.jpg",
	29: "../static/img/hero/tidehunter.jpg",
	30: "../static/img/hero/witchdoctor.jpg",
	31: "../static/img/hero/lich.jpg",
	32: "../static/img/hero/riki.jpg",
	33: "../static/img/hero/enigma.jpg",
	34: "../static/img/hero/tinker.jpg",
	35: "../static/img/hero/sniper.jpg",
	36: "../static/img/hero/necrophos.jpg",
	37: "../static/img/hero/warlock.jpg",
	38: "../static/img/hero/beastmaster.jpg",
	39: "../static/img/hero/queenofpain.jpg",
	40: "../static/img/hero/venomancer.jpg",
	41: "../static/img/hero/gyrocopter.jpg",
	42: "../static/img/hero/wraithking.jpg",
	43: "../static/img/hero/deathprophet.jpg",
	44: "../static/img/hero/phantomassassin.jpg",
	45: "../static/img/hero/pugna.jpg",
	46: "../static/img/hero/templarassassin.jpg",
	47: "../static/img/hero/viper.jpg",
	48: "../static/img/hero/luna.jpg",
	49: "../static/img/hero/dragonknight.jpg",
	50: "../static/img/hero/dazzle.jpg",
	51: "../static/img/hero/clockwerk.jpg",
	52: "../static/img/hero/leshrac.jpg",
	53: "../static/img/hero/naturesprophet.jpg",
	54: "../static/img/hero/lifestealer",
	55: "../static/img/hero/dark-seer.jpg",
	56: "../static/img/hero/clinkz.jpg",
	57: "../static/img/hero/omniknight.jpg",
	58: "../static/img/hero/enchantress.jpg",
	59: "../static/img/hero/huskar.jpg",
	60: "../static/img/hero/nightstalker.jpg",
	61: "../static/img/hero/broodmother.jpg",
	62: "../static/img/hero/bountyhunter.jpg",
	63: "../static/img/hero/weaver.jpg",
	64: "../static/img/hero/jakiro.jpg",
	65: "../static/img/hero/batrider.jpg",
	66: "../static/img/hero/chen.jpg",
	67: "../static/img/hero/spectre.jpg",
	68: "../static/img/hero/ancient.jpg",
	69: "../static/img/hero/doom.jpg",
	70: "../static/img/hero/ursa",
	71: "../static/img/hero/spiritbreaker.jpg",
	72: "../static/img/hero/gyrocopter.jpg",
	73: "../static/img/hero/alchemist.jpg",
	74: "../static/img/hero/invoker.jpg",
	75: "../static/img/hero/silencer.jpg",
	76: "../static/img/hero/outworld-devourer.jpg",
	77: "../static/img/hero/lycan.jpg",
	78: "../static/img/hero/brewmaster.jpg",
	79: "../static/img/hero/shadowdemon.jpg",
	80: "../static/img/hero/lone-druid.jpg",
	81: "../static/img/hero/chaosknight.jpg",
	82: "../static/img/hero/meepo.jpg",
	83: "../static/img/hero/treant-protector.jpg",
	84: "../static/img/hero/ogremagi.jpg",
	85: "../static/img/hero/undying.jpg",
	86: "../static/img/hero/rubick.jpg",
	87: "../static/img/hero/disruptor.jpg",
	88: "../static/img/hero/nyx-assassin.jpg",
	89: "../static/img/hero/naga-siren.jpg",
	90: "../static/img/hero/keeperofthelight.jpg",
	91: "../static/img/hero/io.jpg",
	92: "../static/img/hero/visage.jpg",
	93: "../static/img/hero/slark.jpg",
	94: "../static/img/hero/medusa.jpg",
	95: "../static/img/hero/trollwarlord.jpg",
	96: "../static/img/hero/centaurwarrunner.jpg",
	97: "../static/img/hero/magnus.jpg",
	98: "../static/img/hero/timbersaw.jpg",
	99: "../static/img/hero/bristleback.jpg",
	100: "../static/img/hero/tusk.jpg",
	101: "../static/img/hero/skywrathmage.jpg",
	102: "../static/img/hero/abaddon.jpg",
	103: "../static/img/hero/eldertitan.jpg",
	104: "../static/img/hero/legioncommander.jpg",
	105: "../static/img/hero/techies.jpg",
	106: "../static/img/hero/emberspirit.jpg",
	107: "../static/img/hero/earthspirit.jpg",
	108: "../static/img/hero/underlord.jpg",
	109: "../static/img/hero/terrorblade.jpg",
	110: "../static/img/hero/phoenix.jpg",
	111: "../static/img/hero/oracle.jpg",
	112: "../static/img/hero/winterwyvern.jpg",
	113: "../static/img/hero/arcwarden.jpg",
	114: "../static/img/hero/monkey-king.jpg",
	119: "../static/img/hero/dark-willow.jpg",
	120: "../static/img/hero/pangolier.jpg",
}

var ItemIDIcon = map[int64]string{
	//大飞鞋 ，银月之晶，陨星锤 ，肉山盾，刷新碎片，精灵布带,凝魂之泪，慧光
	1:  "../static/img/item/blink-dagger.jpg",           //跳刀
	2:  "../static/img/item/blades-of-attack.jpg",       //攻击之爪
	3:  "../static/img/item/broadsword.jpg",             //阔剑
	4:  "../static/img/item/chainmail.jpg",              //锁子甲
	5:  "../static/img/item/claymore.jpg",               //大剑
	6:  "../static/img/item/helm_of_iron_will.jpg",      //铁意头盔
	7:  "../static/img/item/javelin.jpg",                //标枪
	8:  "../static/img/item/mithril-hammer.jpg",         //秘银锤
	9:  "../static/img/item/platemail.jpg",              //板甲
	10: "../static/img/item/quarterstaff.jpg",           //短棍
	11: "../static/img/item/quelling-blade.jpg",         //补刀斧
	12: "../static/img/item/ring_of_protection.jpg",     //守护戒指
	13: "../static/img/item/gauntlets-of-strength.jpg",  //力量手套
	14: "../static/img/item/slippers-of-agility.jpg",    //敏捷便鞋
	15: "../static/img/item/mantle-of-intelligence.jpg", //智力斗篷
	16: "../static/img/item/iron-branch.jpg",            //树枝
	17: "../static/img/item/belt-of-strength.jpg",       //力量腰带
	18: "../static/img/item/band-of-elvenskin.jpg",      //精灵布带
	19: "../static/img/item/robe-of-the-magi",           //法师长袍
	20: "../static/img/item/circlet.jpg",                //圆环
	21: "../static/img/item/ogre-axe.jpg",               //食人魔之斧
	22: "../static/img/item/blade-of-alacrity.jpg",      //欢欣之刃
	23: "../static/img/item/staff-of-wizardry.jpg",      //魔力法杖
	24: "../static/img/item/ultimate-orb.jpg",           //极限法球
	25: "../static/img/item/gloves-of-haste.jpg",        //加速手套
	26: "../static/img/item/morbid-mask.jpg",            //吸血面具
	27: "../static/img/item/ring-of-regen.jpg",          //回复戒指
	28: "../static/img/item/sages-mask.jpg",             //艺人面罩
	29: "../static/img/item/boots-of-speed.jpg",         //速度之鞋
	30: "../static/img/item/gem-of-true-sight.jpg",      //珍视宝石
	31: "../static/img/item/cloak.jpg",                  //抗魔斗篷
	32: "../static/img/item/talisman-of-evasion.jpg",    //闪避护符
	33: "../static/img/item/cheese.jpg",                 //奶酪
	34: "../static/img/item/magic-stick.jpg",            //魔棒

	36: "../static/img/item/magic-wand.jpg",         //大魔棒
	37: "../static/img/item/ghost-scepter.jpg",      //幽魂权杖
	38: "../static/img/item/clarity.jpg",            //净化药水
	39: "../static/img/item/healing-salve.jpg",      //大药
	40: "../static/img/item/dust-of-appearance.jpg", //显影之尘
	41: "../static/img/item/bottle.jpg",             //瓶子
	42: "../static/img/item/observer-ward.jpg",      //假眼
	43: "../static/img/item/sentry-ward.jpg",        //真眼
	44: "../static/img/item/tango.jpg",              //吃数
	45: "../static/img/item/animal-courier.jpg",     //动物信使
	46: "../static/img/item/town-portal-scrol.jpg",  //TP
	48: "../static/img/item/boots-of-travel.jpg",    //飞鞋

	50: "../static/img/item/phase-boots.jpg",      //相位鞋
	51: "../static/img/item/demon-edge.jpg",       //恶魔刀锋
	52: "../static/img/item/eaglesong.jpg",        //鹰歌弓
	53: "../static/img/item/reaver.jpg",           //掠夺
	54: "../static/img/item/sacred-relic.jpg",     //圣者遗物
	55: "../static/img/item/hyperstone.jpg",       //振奋宝石
	56: "../static/img/item/ring-of-health.jpg",   //治疗指环
	57: "../static/img/item/void-stone.jpg",       //虚无宝石
	58: "../static/img/item/mystic-staff.jpg",     //神秘法杖
	59: "../static/img/item/energy-booster.jpg",   //能量之球
	60: "../static/img/item/point-booster.jpg",    //精气之球
	61: "../static/img/item/vitality-booster.jpg", //活力之球
	63: "../static/img/item/power-treads.jpg",     //假腿

	65: "../static/img/item/hand-of-midas.jpg", //点金手

	67:  "../static/img/item/oblivion-staff.jpg",           //空明帐
	69:  "../static/img/item/perseverance.jpg",             //坚韧球
	73:  "../static/img/item/bracer.jpg",                   //护腕
	75:  "../static/img/item/wraith-band.jpg",              //怨灵系带
	77:  "../static/img/item/null-talisman.jpg",            //空灵挂件
	79:  "../static/img/item/mekansm.jpg",                  //梅肯
	81:  "../static/img/item/vladmirs-offering.jpg",        //祭品
	84:  "../static/img/item/flying_courier_lg.png",        //动物信使
	86:  "../static/img/item/buckler.jpg",                  //玄冥盾牌
	88:  "../static/img/item/ring-of-basilius.jpg",         //圣殿
	90:  "../static/img/item/pipe-of-insight.jpg",          //笛子
	92:  "../static/img/item/urn-of-shadows.jpg",           //骨灰
	94:  "../static/img/item/headdress.jpg",                //回复头巾
	96:  "../static/img/item/scythe-of-vyse.jpg",           //羊刀
	98:  "../static/img/item/orchid-malevolence.jpg",       //紫苑
	250: "../static/img/item/bloodthorn.jpg",               //血棘
	252: "../static/img/item/echo-sabre.jpg",               //回音战刃
	100: "../static/img/item/euls-scepter-of-divinity.jpg", //吹风
	232: "../static/img/item/aether-lens.jpg",              //以太之镜
	102: "../static/img/item/force-staff.jpg",              //推推
	263: "../static/img/item/hurricane-pike.jpg",           //大推推
	104: "../static/img/item/dagon.jpg",                    //大根
	201: "../static/img/item/dagon-level-2.jpg",
	202: "../static/img/item/dagon-level-3.jpg",
	203: "../static/img/item/dagon-level-4.jpg",
	204: "../static/img/item/dagon-level-5.jpg",
	106: "../static/img/item/necronomicon.jpg", //死灵书
	193: "../static/img/item/necronomicon-level-2.jpg",
	194: "../static/img/item/necronomicon-level-3.jpg",
	108: "../static/img/item/aghanims-scepter.jpg",        //A帐
	110: "../static/img/item/refresher-orb.jpg",           //刷新球
	117: "../static/img/item/aegis-of-the-immortal.jpg",   //肉山盾
	112: "../static/img/item/assault-cuirass.jpg",         //强袭装甲
	114: "../static/img/item/heart-of-tarrasque.jpg",      //龙心
	116: "../static/img/item/black-king-bar.jpg",          //BKB
	118: "../static/img/item/",                            //
	119: "../static/img/item/shivas-guard.jpg",            //西瓦
	121: "../static/img/item/bloodstone.jpg",              //血精石
	123: "../static/img/item/linkens-sphere.jpg",          //林肯
	226: "../static/img/item/lotus-orb.jpg",               //莲花
	125: "../static/img/item/vanguard.jpg",                //先锋盾
	242: "../static/img/item/crimson-guard.jpg",           //赤红甲
	127: "../static/img/item/blade-mail.jpg",              //刃甲
	129: "../static/img/item/soul-booster.jpg",            //镇魂石
	131: "../static/img/item/hood-of-defiance.jpg",        //挑战
	133: "../static/img/item/divine-rapier.jpg",           //圣剑
	135: "../static/img/item/monkey-king-bar.jpg",         //金箍棒
	137: "../static/img/item/radiance.jpg",                //辉耀
	139: "../static/img/item/butterfly.jpg",               //蝴蝶
	141: "../static/img/item/daedalus.jpg",                //大炮
	143: "../static/img/item/skull-basher.jpg",            //碎颅锤
	145: "../static/img/item/battle-fury.jpg",             //狂战斧
	147: "../static/img/item/manta-style.jpg",             //分身斧
	149: "../static/img/item/crystalys.jpg",               //水晶剑
	236: "../static/img/item/dragon-lance.jpg",            //魔龙枪
	151: "../static/img/item/armlet-of-mordiggian.jpg",    //臂章
	152: "../static/img/item/shadow-blade.jpg",            //隐刀
	249: "../static/img/item/silver-edge.jpg",             //大隐刀，白银之锋
	154: "../static/img/item/sange-and-yasha.jpg",         //双刀
	156: "../static/img/item/satanic.jpg",                 //撒旦
	158: "../static/img/item/mjollnir.jpg",                //大电锤
	160: "../static/img/item/eye-of-skadi.jpg",            //冰眼
	162: "../static/img/item/sange.jpg",                   //散华
	164: "../static/img/item/helm-of-the-dominator.jpg",   //支配
	166: "../static/img/item/maelstrom.jpg",               //电锤
	168: "../static/img/item/desolator.jpg",               //暗灭
	170: "../static/img/item/yasha.jpg",                   //单刀夜叉
	172: "../static/img/item/mask-of-madness.jpg",         //疯脸
	174: "../static/img/item/diffusal-blade.jpg",          //散失
	196: "../static/img/item/diffusal-blade-2.jpg",        //
	176: "../static/img/item/ethereal-blade.jpg",          //虚灵刀
	178: "../static/img/item/soul-ring.jpg",               //魂戒
	180: "../static/img/item/arcane-boots.jpg",            //秘法
	235: "../static/img/item/octarine-core.jpg",           //玲珑心
	181: "../static/img/item/orb-of-venom.jpg",            //毒球
	240: "../static/img/item/blight-stone.jpg",            //枯萎之石
	185: "../static/img/item/drum-of-endurance.jpg",       //战鼓
	187: "../static/img/item/medallion-of-courage.jpg",    //勇气勋章
	229: "../static/img/item/solar-crest.jpg",             //炎阳纹章
	188: "../static/img/item/smoke-of-deceit.jpg",         //雾
	257: "../static/img/item/tome-of-knowledge.jpg",       //经验书
	190: "../static/img/item/veil-of-discord.jpg",         //纷争面纱
	231: "../static/img/item/guardian-greaves.jpg",        //大鞋
	206: "../static/img/item/rod-of-atos.jpg",             //阿托斯
	239: "../static/img/item/",                            //打野爪
	208: "../static/img/item/abyssal-blade.jpg",           //深渊之刃
	210: "../static/img/item/heavens-halberd.jpg",         //天堂
	212: "../static/img/item/ring-of-aquila.jpg",          //天鹰
	214: "../static/img/item/tranquil-boots.jpg",          //绿鞋
	215: "../static/img/item/shadow-amulet.jpg",           //暗影护符
	254: "../static/img/item/glimmer-cape.jpg",            //微光披风
	256: "../static/img/item/aeon-disk.jpg",               //免死金牌
	225: "../static/img/item/nullifier.jpg",               //否决
	267: "../static/img/item/spirit-vessel.jpg",           //大骨灰
	244: "../static/img/item/wind-lace.jpg",               //风灵之纹
	259: "../static/img/item/kaya.jpg",                    //慧光
	260: "../static/img/item/refresher-shard.jpg",         //刷新碎片
	223: "../static/img/item/meteor-hammer.jpg",           //陨星锤
	220: "../static/img/item/boots-of-travel-level-2.jpg", //大飞鞋
	265: "../static/img/item/infused-raindrops.jpg",       //凝魂之泪
	247: "../static/img/item/moon-shard.jpg",              //银月之晶
}

var AbilityIcon= map[string]string{
	"antimage_mana_break":"../static/img/abilities/anti-mage-mana-break.jpg",
	"antimage_blink":"../static/img/abilities/anti-mage-blink.jpg",
	"antimage_counterspell":"../static/img/abilities/anti-mage-counterspel.jpg",
	"antimage_mana_void":"../static/img/abilities/anti-mage-mana-void.jpg",
	//axe
	"axe_berserkers_call":"../static/img/abilities/axe-berserkers-call.jpg",
	"axe_battle_hunger":"../static/img/abilities/axe-battle-hunger.jpg",
	"axe_counter_helix":"../static/img/abilities/axe-counter-helix.jpg",
	"axe_culling_blade":"../static/img/abilities/axe-culling-blade.jpg",
	//bane
	"bane_enfeeble":"../static/img/abilities/bane-enfeeble.jpg",
	"bane_brain_sap":"../static/img/abilities/bane-brain-sap.jpg",
	"bane_nightmare":"../static/img/abilities/bane-nightmare.jpg",
	"bane_fiends_grip":"../static/img/abilities/bane-fiends-grip.jpg",
	//bloodseeker
	"bloodseeker_bloodrage":"../static/img/abilities/bloodseeker-bloodrage.jpg",
	"bloodseeker_blood_bath":"../static/img/abilities/bloodseeker-blood-rite.jpg",
	"bloodseeker_thirst":"../static/img/abilities/bloodseeker-thirst.jpg",
	"bloodseeker_rupture":"../static/img/abilities/bloodseeker-rupture.jpg",
	//冰女crystal_maiden
	"crystal_maiden_crystal_nova":"../static/img/abilities/crystal-maiden-crystal-nova.jpg",
	"crystal_maiden_frostbite":"../static/img/abilities/crystal-maiden-frostbite.jpg",
	"crystal_maiden_brilliance_aura":"../static/img/abilities/crystal-maiden-arcane-aura.jpg",
	"crystal_maiden_freezing_field":"../static/img/abilities/crystal-maiden-freezing-field.jpg",
	//卓尔游侠drow_ranger
	"drow_ranger_frost_arrows":"../static/img/abilities/drow-ranger-frost-arrows.jpg",
	"drow_ranger_wave_of_silence":"../static/img/abilities/drow-ranger-gust.jpg",
	"drow_ranger_trueshot":"../static/img/abilities/drow-ranger-marksmanship.jpg",
	"drow_ranger_marksmanship":"../static/img/abilities/drow-ranger-precision-aura.jpg",
	//撼地者earthshaker
	"earthshaker_fissure":"../static/img/abilities/earthshaker-fissure.jpg",
	"earthshaker_enchant_totem":"../static/img/abilities/earthshaker-enchant-totem.jpg",
	"earthshaker_aftershock":"../static/img/abilities/earthshaker-aftershock.jpg",
	"earthshaker_echo_slam":"../static/img/abilities/earthshaker-echo-slam.jpg",
	//juggernaut剑圣
	"juggernaut_blade_fury":"../static/img/abilities/juggernaut-blade-fury.jpg",
	"juggernaut_healing_ward":"../static/img/abilities/juggernaut-healing-ward.jpg",
	"juggernaut_blade_dance":"../static/img/abilities/juggernaut-blade-dance.jpg",
	"juggernaut_omni_slash":"../static/img/abilities/juggernaut-omnislash.jpg",
	//米拉娜mirana_starfall
	"mirana_starfall":"../static/img/abilities/mirana-starstorm.jpg",
	"mirana_arrow":"../static/img/abilities/mirana-sacred-arrow.jpg",
	"mirana_leap":"../static/img/abilities/mirana-leap.jpg",
	"mirana_invis":"../static/img/abilities/mirana-moonlight-shadow.jpg",
	//影魔
	"nevermore_shadowraze1":"../static/img/abilities/shadow-fiend-shadowraze.jpg",
	"nevermore_necromastery":"../static/img/abilities/shadow-fiend-necromastery.jpg",
	"nevermore_dark_lord":"../static/img/abilities/shadow-fiend-presence-of-the-dark-lord.jpg",
	"nevermore_requiem":"../static/img/abilities/shadow-fiend-requiem-of-souls.jpg",
	//水人
	"morphling_waveform":"../static/img/abilities/morphling-waveform.jpg",
	"morphling_adaptive_strike_agi":"../static/img/abilities/morphling-adaptive-strike-agility.jpg",
	"morphling_adaptive_strike_str":"../static/img/abilities/morphling-adaptive-strike-strength.jpg",
	"morphling_morph_agi":"../static/img/abilities/morphling-attribute-shift-agility-gain.jpg",
	"morphling_morph_str":"../static/img/abilities/morphling_morph_str_lg.jpg",
	"morphling_replicate":"../static/img/abilities/morphling_morph_replicate_lg.jpg",
	"morphling_morph_replicate":"../static/img/abilities/morphling-morph.jpg",
	"morphling_morph":"",
	//幻影长矛手phantom_lancer_spirit_lance
	"phantom_lancer_spirit_lance":"../static/img/abilities/phantom-lancer-spirit-lance.jpg",
	"phantom_lancer_doppelwalk":"../static/img/abilities/phantom-lancer-doppelganger.jpg",
	"phantom_lancer_phantom_edge":"../static/img/abilities/phantom-lancer-phantom-rush.jpg",
	"phantom_lancer_juxtapose":"../static/img/abilities/phantom-lancer-juxtapose.jpg",
	//puck
	"puck_illusory_orb":"../static/img/abilities/puck-illusory-orb.jpg",
	"puck_waning_rift":"../static/img/abilities/puck-waning-rift.jpg",
	"puck_phase_shift":"../static/img/abilities/puck-phase-shift.jpg",
	"puck_ethereal_jaunt":"",
	"puck_dream_coil":"../static/img/abilities/puck-dream-coil.jpg",
	//帕吉pudge
	"pudge_meat_hook":"../static/img/abilities/pudge-meat-hook.jpg",
	"pudge_rot":"../static/img/abilities/pudge-rot.jpg",
	"pudge_flesh_heap":"../static/img/abilities/pudge-flesh-heap.jpg",
	"pudge_dismember":"../static/img/abilities/pudge-dismember.jpg",
	//razor
	"razor_plasma_field":"../static/img/abilities/razor-plasma-field.jpg",
	"razor_static_link":"../static/img/abilities/razor-static-link.jpg",
	"razor_unstable_current":"../static/img/abilities/razor-unstable-current.jpg",
	"razor_eye_of_the_storm":"../static/img/abilities/razor-eye-of-the-storm.jpg",
	//sandking
	"sandking_burrowstrik":"../static/img/abilities/sand-king-burrowstrike.jpg",
	"sandking_sand_storm":"../static/img/abilities/sand-king-sand-storm.jpg",
	"sandking_caustic_finale":"../static/img/abilities/sand-king-caustic-finale.jpg",
	"sandking_epicenter":"../static/img/abilities/sand-king-epicenter.jpg",
	//storm
	"storm_spirit_static_remnant":"../static/img/abilities/storm-spirit-static-remnant.jpg",
	"storm_spirit_electric_vortex":"../static/img/abilities/storm-spirit-electric-vortex.jpg",
	"storm_spirit_overload":"../static/img/abilities/storm-spirit-overload.jpg",
	"storm_spirit_ball_lightning":"../static/img/abilities/storm-spirit-ball-lightning.jpg",
	//sven
	"sven_storm_bolt":"../static/img/abilities/sven-storm-hammer.jpg",
	"sven_great_cleave":"../static/img/abilities/sven-great-cleave.jpg",
	"sven_warcry":"../static/img/abilities/sven-warcry.jpg",
	"sven_gods_strength":"../static/img/abilities/sven-gods-strength.jpg",
	//tiny
	"tiny_avalanche":"../static/img/abilities/tiny-avalanche.jpg",
	"tiny_toss":"../static/img/abilities/tiny-toss.jpg",
	"tiny_craggy_exterior":"../static/img/abilities/tiny-tree-grab.jpg",
	"tiny_tree_channel":"../static/img/abilities/tiny-tree-grab.jpg",
	"tiny_grow":"../static/img/abilities/tiny-grow.jpg",
	"tiny_toss_tree":"../static/img/abilities/tiny-tree-grab.jpg",
	//复仇之魂
	"vengefulspirit_magic_missile":"../static/img/abilities/vengeful-spirit-magic-missile.jpg",
	"vengefulspirit_wave_of_terror":"../static/img/abilities/vengeful-spirit-wave-of-terror.jpg",
	"vengefulspirit_command_aura":"../static/img/abilities/vengeful-spirit-vengeance-aura.jpg",
	"vengefulspirit_nether_swap":"../static/img/abilities/vengeful-spirit-nether-swap.jpg",
	//风行者
	"windrunner_shackleshot":"../static/img/abilities/windranger-shackleshot.jpg",
	"windrunner_powershot":"../static/img/abilities/windranger-powershot.jpg",
	"windrunner_windrun":"../static/img/abilities/windranger-windrun.jpg",
	"windrunner_focusfire":"../static/img/abilities/windranger-focus-fire.jpg",
	//宙斯
	"zuus_arc_lightning":"../static/img/abilities/zeus-arc-lightning.jpg",
	"zuus_lightning_bolt":"../static/img/abilities/zeus-lightning-bolt.jpg",
	"zuus_static_field":"../static/img/abilities/zeus-static-field.jpg",
	"zuus_cloud":"../static/img/abilities/zuus_cloud_lg.jpg",
	"zuus_thundergods_wrath":"../static/img/abilities/zeus-thundergods-wrath.jpg",
	//昆卡
	"kunkka_torrent":"../static/img/abilities/kunkka-torrent.jpg",
	"kunkka_tidebringer":"../static/img/abilities/kunkka-tidebringer.jpg",
	"kunkka_x_marks_the_spot":"../static/img/abilities/kunkka-x-marks-the-spot.jpg",
	"kunkka_ghostship":"../static/img/abilities/kunkka-ghostship.jpg",
	//丽娜
	"lina_dragon_slave":"../static/img/abilities/lina-dragon-slave.jpg",
	"lina_light_strike_array":"../static/img/abilities/lina-light-strike-array.jpg",
	"lina_fiery_soul":"../static/img/abilities/lina-fiery-soul.jpg",
	"lina_laguna_blade":"../static/img/abilities/lina-laguna-blade.jpg",
	//lich
	"lich_frost_nova":"../static/img/abilities/lich-frost-blast.jpg",
	"lich_frost_shield":"../static/img/abilities/lich-frost-shield.jpg",
	"lich_sinister_gaze":"../static/img/abilities/lich-sinister-gaze.jpg",
	"lich_chain_frost":"../static/img/abilities/lich-chain-frost.jpg",
	//lion
	"lion_impale":"../static/img/abilities/lion-earth-spike.jpg",
	"lion_voodoo":"../static/img/abilities/lion-hex.jpg",
	"lion_mana_drain":"../static/img/abilities/lion-mana-drain.jpg",
	"lion_finger_of_death":"../static/img/abilities/lion-finger-of-death.jpg",
	//shadow_shaman
	"shadow_shaman_ether_shock":"../static/img/abilities/shadow-shaman-ether-shock.jpg",
	"shadow_shaman_voodoo":"../static/img/abilities/shadow-shaman-hex.jpg",
	"shadow_shaman_shackles":"../static/img/abilities/shadow-shaman-shackles.jpg",
	"shadow_shaman_mass_serpent_ward":"../static/img/abilities/shadow-shaman-mass-serpent-ward.jpg",
	//大鱼
	"slardar_sprint":"../static/img/abilities/slardar-guardian-sprint.jpg",
	"slardar_slithereen_crush":"../static/img/abilities/slardar-slithereen-crush.jpg",
	"slardar_bash":"../static/img/abilities/slardar-bash-of-the-deep.jpg",
	"slardar_amplify_damage":"../static/img/abilities/slardar-corrosive-haze.jpg",
	//潮汐猎人
	"tidehunter_gush":"../static/img/abilities/tidehunter-gush.jpg",
	"tidehunter_kraken_shell":"../static/img/abilities/tidehunter-anchor-smash.jpg",
	"tidehunter_anchor_smash":"../static/img/abilities/tidehunter-kraken-shell.jpg",
	"tidehunter_ravage":"../static/img/abilities/tidehunter-ravage.jpg",
	//巫医
	"witch_doctor_paralyzing_cask":"../static/img/abilities/witch-doctor-paralyzing.jpg",
	"witch_doctor_voodoo_restoration":"../static/img/abilities/witch-doctor-voodoo-restoration.jpg",
	"witch_doctor_maledict":"../static/img/abilities/witch-doctor-maledict-5140.jpg",
	"witch_doctor_death_ward":"../static/img/abilities/witch-doctor-death-ward.jpg",
	//力丸
	"riki_smoke_screen":"../static/img/abilities/wriki-smoke-screen.jpg",
	"riki_blink_strike":"../static/img/abilities/riki-blink-strike.jpg",
	"riki_permanent_invisibility":"../static/img/abilities/riki-cloak-and-dagger.jpg",
	"riki_tricks_of_the_trade":"../static/img/abilities/riki-tricks-of-the-trade.jpg",
	//谜团enigma
	"enigma_malefice":"../static/img/abilities/enigma-malefice.jpg",
	"enigma_demonic_conversion":"../static/img/abilities/enigma-demonic-conversion.jpg",
	"enigma_midnight_pulse":"../static/img/abilities/enigma-midnight-pulse.jpg",
	"enigma_black_hole":"../static/img/abilities/enigma-black-hole.jpg",
	//tinker
	"tinker_laser":"../static/img/abilities/tinker-laser.jpg",
	"tinker_heat_seeking_missile":"../static/img/abilities/tinker-heat-seeking-missile.jpg",
	"tinker_march_of_the_machines":"../static/img/abilities/tinker-march-of-the-machines.jpg",
	"tinker_rearm":"../static/img/abilities/tinker-rearm.jpg",
	//火枪
	"sniper_shrapnel":"../static/img/abilities/sniper-shrapnel.jpg",
	"sniper_headshot":"../static/img/abilities/sniper-headshot.jpg",
	"sniper_take_aim":"../static/img/abilities/sniper-take-aim.jpg",
	"sniper_assassinate":"../static/img/abilities/sniper-assassinate.jpg",
	//瘟疫法师
	"necrolyte_death_pulse":"../static/img/abilities/necrophos-death-pulse.jpg",
	"necrolyte_sadist":"../static/img/abilities/necrophos-ghost-shroud.jpg",
	"necrolyte_heartstopper_aura":"../static/img/abilities/necrophos-heartstopper-aura.jpg",
	"necrolyte_reapers_scythe":"../static/img/abilities/necrophos-reapers.jpg",
	//warlock
	"warlock_fatal_bonds":"../static/img/abilities/warlock-fatal-bonds.jpg",
	"warlock_shadow_word":"../static/img/abilities/warlock-shadow-word.jpg",
	"warlock_upheaval":"../static/img/abilities/warlock-upheaval.jpg",
	"warlock_rain_of_chaos":"../static/img/abilities/warlock-chaotic-offering.jpg",
	//兽王
	"beastmaster_wild_axes":"../static/img/abilities/beastmaster-wild-axes.jpg",
	"beastmaster_call_of_the_wild_boar":"../static/img/abilities/beastmaster-call-of-the-wild-boar.jpg",
	"beastmaster_call_of_the_wild_hawk":"../static/img/abilities/beastmaster-call-of-the-wild-hawk.jpg",
	"beastmaster_inner_beast":"../static/img/abilities/beastmaster-inner-beast.jpg",
	"beastmaster_primal_roar":"../static/img/abilities/beastmaster-primal-roar.jpg",
	//痛苦女王
	"queenofpain_shadow_strike":"../static/img/abilities/queen-of-pain-shadow-strike.jpg",
	"queenofpain_blink":"../static/img/abilities/queen-of-pain-blink.jpg",
	"queenofpain_scream_of_pain":"../static/img/abilities/queen-of-pain-scream-of-pain.jpg",
	"queenofpain_sonic_wave":"../static/img/abilities/queen-of-pain-sonic-wave.jpg",
	//剧毒
	"venomancer_venomous_gale":"../static/img/abilities/venomancer-venomous.jpg",
	"venomancer_poison_sting":"../static/img/abilities/venomancer-poison-sting.jpg",
	"venomancer_plague_ward":"../static/img/abilities/venomancer-plague-ward.jpg",
	"venomancer_poison_nova":"../static/img/abilities/venomancer-poison-nova.jpg",
	//虚空
	"faceless_void_time_walk":"../static/img/abilities/faceless-void-time-walk.jpg",
	"faceless_void_time_dilation":"../static/img/abilities/faceless-void-time-dilation.jpg",
	"faceless_void_time_lock":"../static/img/abilities/faceless-void-time-lock.jpg",
	"faceless_void_chronosphere":"../static/img/abilities/faceless-void-chronosphere.jpg",
	//骷髅王
	"skeleton_king_hellfire_blast":"../static/img/abilities/wraith-king-wraithfire-blast.jpg",
	"skeleton_king_vampiric_aura":"../static/img/abilities/wraith-king-vampiric-aura.jpg",
	"skeleton_king_mortal_strike":"../static/img/abilities/wraith-king-mortal-strike.jpg",
	"skeleton_king_reincarnation":"../static/img/abilities/wraith-king-reincarnation.jpg",
	//死亡先知
	"death_prophet_carrion_swarm":"../static/img/abilities/death-prophet-crypt-swarm.jpg",
	"death_prophet_silence":"../static/img/abilities/death-prophet-silence.jpg",
	"death_prophet_spirit_siphon":"../static/img/abilities/death-prophet-spirit-siphon.jpg",
	"death_prophet_exorcism":"../static/img/abilities/death-prophet-exorcism.jpg",
	//幻影刺客
	"phantom_assassin_stifling_dagger":"../static/img/abilities/phantom-assassin-stifling-dagger.jpg",
	"phantom_assassin_phantom_strike":"../static/img/abilities/phantom-assassin-phantom-strike.jpg",
	"phantom_assassin_blur":"../static/img/abilities/phantom-assassin-blur.jpg",
	"phantom_assassin_coup_de_grace":"../static/img/abilities/phantom-assassin-coup-de-grace.jpg",
	//帕格纳
	"pugna_nether_blast":"../static/img/abilities/pugna-nether-blast.jpg",
	"pugna_decrepify":"../static/img/abilities/pugna-decrepify.jpg",
	"pugna_nether_ward":"../static/img/abilities/pugna-nether-ward.jpg",
	"pugna_life_drain":"../static/img/abilities/pugna-life-drain.jpg",
	//TA
	"templar_assassin_refraction":"../static/img/abilities/templar-assassin-refraction.jpg",
	"templar_assassin_meld":"../static/img/abilities/templar-assassin-meld.jpg",
	"templar_assassin_psi_blades":"../static/img/abilities/templar-assassin-psi-blades.jpg",
	"templar_assassin_trap":"../static/img/abilities/templar-assassin-psionic-trap.jpg",
	"templar_assassin_psionic_trap":"../static/img/abilities/templar-assassin-psionic-trap.jpg",
	//viper
	"viper_poison_attack":"../static/img/abilities/viper-poison-attack.jpg",
	"viper_nethertoxin":"../static/img/abilities/viper-nethertoxin.jpg",
	"viper_corrosive_skin":"../static/img/abilities/viper-corrosive-skin.jpg",
	"viper_viper_strike":"../static/img/abilities/viper-viper-strike.jpg",
	//luna
	"luna_lucent_beam":"../static/img/abilities/luna-lucent-beam.jpg",
	"luna_moon_glaive":"../static/img/abilities/luna-moon-glaives.jpg",
	"luna_lunar_blessing":"../static/img/abilities/luna-lunar-blessing.jpg",
	"luna_eclipse":"../static/img/abilities/luna-eclipse.jpg",
	//龙骑士
	"dragon_knight_breathe_fire":"../static/img/abilities/dragon-knight-breathe-fire.jpg",
	"dragon_knight_dragon_tail":"../static/img/abilities/dragon-knight-dragon-tail.jpg",
	"dragon_knight_dragon_blood":"../static/img/abilities/dragon-knight-dragon-blood.jpg",
	"dragon_knight_elder_dragon_form":"../static/img/abilities/dragon-knight-elder-dragon-form.jpg",
	//戴泽
	"dazzle_poison_touch":"../static/img/abilities/dazzle-poison-touch.jpg",
	"dazzle_shallow_grave":"../static/img/abilities/dazzle-shallow-grave.jpg",
	"dazzle_shadow_wave":"../static/img/abilities/dazzle-shadow-wave.jpg",
	"dazzle_bad_juju":"../static/img/abilities/dazzle-bad-juju.jpg",
	//发条
	"rattletrap_battery_assault":"../static/img/abilities/clockwerk-battery-assault.jpg",
	"rattletrap_power_cogs":"../static/img/abilities/clockwerk-power-cogs.jpg",
	"rattletrap_rocket_flare":"../static/img/abilities/clockwerk-rocket-flare.jpg",
	"rattletrap_hookshot":"../static/img/abilities/clockwerk-hookshot.jpg",
	//拉席克
	"leshrac_split_earth":"../static/img/abilities/leshrac-split-earth.jpg",
	"leshrac_diabolic_edict":"../static/img/abilities/leshrac-diabolic-edict.jpg",
	"leshrac_lightning_storm":"../static/img/abilities/leshrac-lightning-storm.jpg",
	"leshrac_pulse_nova":"../static/img/abilities/leshrac-pulse-nova.jpg",
	//先知
	"furion_sprout":"../static/img/abilities/natures-prophet-sprout.jpg",
	"furion_teleportation":"../static/img/abilities/natures-prophet-teleportation.jpg",
	"furion_force_of_nature":"../static/img/abilities/natures-prophet-natures-call.jpg",
	"furion_wrath_of_nature":"../static/img/abilities/natures-prophet-wrath-of-nature.jpg",
	//小狗
	"life_stealer_rage":"../static/img/abilities/lifestealer-rage.jpg",
	"life_stealer_feast":"../static/img/abilities/lifestealer-feast.jpg",
	"life_stealer_open_wounds":"../static/img/abilities/lifestealer-open-wounds.jpg",
	"life_stealer_infest":"../static/img/abilities/lifestealer-infest.jpg",
	//黑暗贤者
	"dark_seer_vacuum":"../static/img/abilities/dark-seer-vacuum.jpg",
	"dark_seer_ion_shell":"../static/img/abilities/dark-seer-ion-shell.jpg",
	"dark_seer_surge":"../static/img/abilities/dark-seer-surge.jpg",
	"dark_seer_wall_of_replica":"../static/img/abilities/dark-seer-wall-of-replica.jpg",
	//克林克兹
	"clinkz_strafe":"../static/img/abilities/clinkz-strafe.jpg",
	"clinkz_searing_arrows":"../static/img/abilities/clinkz-searing-arrows.jpg",
	"clinkz_wind_walk":"../static/img/abilities/clinkz-skeleton-walk.jpg",
	"clinkz_burning_army":"../static/img/abilities/clinkz-burning-army.jpg",
	//全能
	"omniknight_purification":"../static/img/abilities/omniknight-purification.jpg",
	"omniknight_repel":"../static/img/abilities/omniknight-heavenly-grace.jpg",
	"omniknight_degen_aura":"../static/img/abilities/omniknight-degen-aura.jpg",
	"omniknight_guardian_angel":"../static/img/abilities/omniknight-guardian-angel.jpg",
	//小鹿
	"enchantress_untouchable":"../static/img/abilities/enchantress-untouchable.jpg",
	"enchantress_enchant":"../static/img/abilities/enchantress-enchant.jpg",
	"enchantress_natures_attendants":"../static/img/abilities/enchantress-natures-attendants.jpg",
	"enchantress_impetus":"../static/img/abilities/enchantress-impetus.jpg",
	//哈斯卡
	"huskar_inner_fire":"../static/img/abilities/huskar-inner-fire.jpg",
	"huskar_burning_spear":"../static/img/abilities/huskar-burning-spear.jpg",
	"huskar_berserkers_blood":"../static/img/abilities/huskar-berserkers-blood.jpg",
	"huskar_life_break":"../static/img/abilities/huskar-life-break.jpg",
	//夜魔
	"night_stalker_void":"../static/img/abilities/night-stalker-void.jpg",
	"night_stalker_crippling_fear":"../static/img/abilities/night-stalker-crippling-fear.jpg",
	"night_stalker_hunter_in_the_night":"../static/img/abilities/night-stalker-hunter-in-the-night.jpg",
	"night_stalker_darkness":"../static/img/abilities/night-stalker-dark-ascension.jpg",
	//育母蜘蛛
	"broodmother_spawn_spiderlings":"../static/img/abilities/broodmother-spawn-spiderlings.jpg",
	"broodmother_spin_web":"../static/img/abilities/broodmother-spin-web.jpg",
	"broodmother_incapacitating_bite":"../static/img/abilities/broodmother-incapacitating-bite.jpg",
	"broodmother_insatiable_hunger":"../static/img/abilities/broodmother-insatiable-hunger.jpg",
	//赏金
	"bounty_hunter_shuriken_toss":"../static/img/abilities/bounty-hunter-shuriken-toss.jpg",
	"bounty_hunter_jinada":"../static/img/abilities/bounty-hunter-jinada.jpg",
	"bounty_hunter_wind_walk":"../static/img/abilities/bounty-hunter-shadow-walk.jpg",
	"bounty_hunter_track":"../static/img/abilities/bounty-hunter-track.jpg",
	//蚂蚁
	"weaver_the_swarm":"../static/img/abilities/weaver-the-swarm.jpg",
	"weaver_shukuchi":"../static/img/abilities/weaver-shukuchi.jpg",
	"weaver_geminate_attack":"../static/img/abilities/weaver-geminate-attack.jpg",
	"weaver_time_lapse":"../static/img/abilities/weaver-time-lapse.jpg",
	//双头龙
	"jakiro_dual_breath":"../static/img/abilities/jakiro-dual-breath.jpg",
	"jakiro_ice_path":"../static/img/abilities/jakiro-ice-path.jpg",
	"jakiro_liquid_fire":"../static/img/abilities/jakiro-liquid-fire.jpg",
	"jakiro_macropyre":"../static/img/abilities/jakiro-macropyre.jpg",
	//双蝙蝠
	"batrider_sticky_napalm":"../static/img/abilities/batrider-sticky-napalm.jpg",
	"batrider_flamebreak":"../static/img/abilities/batrider-flamebreak.jpg",
	"batrider_firefly":"../static/img/abilities/batrider-firefly.jpg",
	"batrider_flaming_lasso":"../static/img/abilities/batrider-flaming-lasso.jpg",
	//chen
	"chen_penitence":"../static/img/abilities/chen-penitence.jpg",
	"chen_holy_persuasion":"../static/img/abilities/chen-holy-persuasion.jpg",
	"chen_divine_favor":"../static/img/abilities/chen-divine-favor.jpg",
	"chen_hand_of_god":"../static/img/abilities/chen-hand-of-god.jpg",
	//幽鬼
	"spectre_spectral_dagger":"../static/img/abilities/spectre-spectral-dagger.jpg",
	"spectre_desolate":"../static/img/abilities/spectre-desolate.jpg",
	"spectre_dispersion":"../static/img/abilities/spectre-dispersion.jpg",
	"spectre_haunt_single":"../static/img/abilities/spectre-haunt.jpg",
	"spectre_haunt":"../static/img/abilities/spectre-haunt.jpg",

	//doom
	"doom_bringer_devour":"../static/img/abilities/doom-devour.jpg",
	"doom_bringer_scorched_earth":"../static/img/abilities/doom-scorched-earth.jpg",
	"doom_bringer_infernal_blade":"../static/img/abilities/doom-infernal.jpg",
	"doom_bringer_doom":"../static/img/abilities/doom-doom.jpg",
	//冰魂
	"ancient_apparition_cold_feet":"../static/img/abilities/ancient-apparition-cold-feet.jpg",
	"ancient_apparition_ice_vortex":"../static/img/abilities/ancient-apparition-ice-vortex.jpg",
	"ancient_apparition_chilling_touch":"../static/img/abilities/ancient-apparition-chilling-touch.jpg",
	"ancient_apparition_ice_blast":"../static/img/abilities/ancient-apparition-release.jpg",
	//熊战士
	"ursa_earthshock":"../static/img/abilities/ursa-earthshock.jpg",
	"ursa_overpower":"../static/img/abilities/ursa-overpower.jpg",
	"ursa_fury_swipes":"../static/img/abilities/ursa-fury-swipes.jpg",
	"ursa_enrage":"../static/img/abilities/ursa-enrage.jpg",
	//猎魂人
	"spirit_breaker_charge_of_darkness":"../static/img/abilities/spirit-breaker-charge-of-darkness.jpg",
	"spirit_breaker_bulldoze":"../static/img/abilities/spirit-breaker-bulldoze.jpg",
	"spirit_breaker_greater_bash":"../static/img/abilities/spirit-breaker-greater-bash.jpg",
	"spirit_breaker_nether_strike":"../static/img/abilities/spirit-breaker-nether-strike.jpg",
	//飞机
	"gyrocopter_rocket_barrage":"../static/img/abilities/gyrocopter-rocket-barrage.jpg",
	"gyrocopter_homing_missile":"../static/img/abilities/gyrocopter-homing-missile.jpg",
	"gyrocopter_flak_cannon":"../static/img/abilities/gyrocopter-flak-cannon.jpg",
	"gyrocopter_call_down":"../static/img/abilities/gyrocopter-call-down.jpg",
	//炼金
	"alchemist_acid_spray":"../static/img/abilities/alchemist-acid-spray.jpg",
	"alchemist_unstable_concoction":"../static/img/abilities/galchemist-unstable-concoction.jpg",
	"alchemist_goblins_greed":"../static/img/abilities/alchemist-greevils-greed.jpg",
	"alchemist_chemical_rage":"../static/img/abilities/alchemist-chemical-rage.jpg",
	//祈求者
	"invoker_quas":"../static/img/abilities/alchemist-acid-spray.jpg",
	"invoker_wex":"../static/img/abilities/galchemist-unstable-concoction.jpg",
	"invoker_exort":"../static/img/abilities/alchemist-greevils-greed.jpg",
	"invoker_invoke":"../static/img/abilities/alchemist-chemical-rage.jpg",
	"invoker_cold_snap":"../static/img/abilities/invoker_cold_snap_lg.png",
	"invoker_ghost_walk":"../static/img/abilities/invoker_ghost_walk_lg.png",
	"invoker_tornado":"../static/img/abilities/invoker_tornado_lg.png",
	"invoker_emp":"../static/img/abilities/invoker_emp_lg.png",
	"invoker_alacrity":"../static/img/abilities/invoker_alacrity_lg.png",
	"invoker_chaos_meteor":"../static/img/abilities/invoker_chaos_meteor_lg.png",
	"invoker_sun_strik":"../static/img/abilities/invoker_sun_strike_lg.png",
	"invoker_forge_spirit":"../static/img/abilities/invoker_forge_spirit_lg.png",
	"invoker_deafening_blast":"../static/img/abilities/invoker_deafening_blast_lg.png",
	//沉默
	"silencer_curse_of_the_silent":"../static/img/abilities/silencer-arcane-curse.jpg",
	"silencer_glaives_of_wisdom":"../static/img/abilities/silencer-glaives-of-wisdom.jpg",
	"silencer_global_silence":"../static/img/abilities/silencer-global-silence.jpg",
	"silencer_last_word":"../static/img/abilities/silencer-last-word.jpg",
	//黑鸟
	"obsidian_destroyer_arcane_orb":"../static/img/abilities/outworld-devourer-arcane-orb.jpg",
	"obsidian_destroyer_astral_imprisonment":"../static/img/abilities/outworld-devourer-astral-imprisonment.jpg",
	"obsidian_destroyer_equilibrium":"../static/img/abilities/outworld-devourer-equilibrium.jpg",
	"obsidian_destroyer_sanity_eclipse":"../static/img/abilities/outworld-devourer-sanitys-eclipse.jpg",
	//狼人
	"lycan_summon_wolves":"../static/img/abilities/lycan-summon-wolves.jpg",
	"lycan_howl":"../static/img/abilities/lycan-howl.jpg",
	"lycan_feral_impulse":"../static/img/abilities/lycan-feral-impulse.jpg",
	"lycan_shapeshift":"../static/img/abilities/lycan-shapeshift.jpg",
	//酒仙
	"brewmaster_thunder_clap":"../static/img/abilities/brewmaster-thunder-clap.jpg",
	"brewmaster_cinder_brew":"../static/img/abilities/brewmaster-cinder-brew.jpg",
	"brewmaster_drunken_brawler":"../static/img/abilities/brewmaster-drunken-brawler.jpg",
	"brewmaster_primal_split":"../static/img/abilities/brewmaster-primal-split.jpg",
	//暗影恶魔
	"shadow_demon_disruption":"../static/img/abilities/shadow-demon-disruption.jpg",
	"shadow_demon_soul_catcher":"../static/img/abilities/shadow-demon-soul-catcher.jpg",
	"shadow_demon_shadow_poison":"../static/img/abilities/shadow-demon-shadow-poison.jpg",
	"shadow_demon_demonic_purge":"../static/img/abilities/shadow-demon-demonic-purge.jpg",
	//德鲁伊
	"lone_druid_spirit_bear":"../static/img/abilities/lone-druid-summon-spirit.jpg",
	"lone_druid_spirit_link":"../static/img/abilities/lone-druid-spirit-link.jpg",
	"lone_druid_savage_roar":"../static/img/abilities/lone-druid-savage.jpg",
	"lone_druid_true_form":"../static/img/abilities/lone-druid-true-form.jpg",
	"lone_druid_true_form_druid":"../static/img/abilities/lone_druid_true_form_druid_lg.jpg",
	//米波
	"meepo_earthbind":"../static/img/abilities/meepo-earthbind.jpg",
	"meepo_poof":"../static/img/abilities/meepo-poof.jpg",
	"meepo_ransack":"../static/img/abilities/meepo-ransack.jpg",
	"meepo_divided_we_stand":"../static/img/abilities/meepo-divided-we-stand.jpg",
	//大树
	"treant_natures_guise":"../static/img/abilities/treant-protector-natures-guise.jpg",
	"treant_leech_seed":"../static/img/abilities/treant-protector-leech-seed.jpg",
	"treant_living_armor":"../static/img/abilities/treant-protector-living-armor.jpg",
	"treant_overgrowth":"../static/img/abilities/treant-protector-overgrowth.jpg",
	//蓝胖
	"ogre_magi_fireblast":"../static/img/abilities/ogre-magi-fireblast.jpg",
	"ogre_magi_ignite":"../static/img/abilities/ogre-magi-ignite.jpg",
	"ogre_magi_bloodlust":"../static/img/abilities/ogre-magi-bloodlust.jpg",
	"ogre_magi_unrefined_fireblast":"../static/img/abilities/ogre-magi-fireblast.jpg",
	"ogre_magi_multicast":"../static/img/abilities/ogre-magi-multicast.jpg",
	//尸王
	"undying_decay":"../static/img/abilities/undying-decay.jpg",
	"undying_soul_rip":"../static/img/abilities/undying-soul-rip.jpg",
	"undying_tombstone":"../static/img/abilities/undying-tombstone.jpg",
	"undying_flesh_golem":"../static/img/abilities/undying-flesh-golem.jpg",
	//拉比克
	"rubick_telekinesis":"../static/img/abilities/rubick-telekinesis.jpg",
	"rubick_fade_bolt":"../static/img/abilities/rubick-fade-bolt.jpg",
	"rubick_arcane_supremacy":"../static/img/abilities/rubick-arcane-supremacy.jpg",
	"rubick_spell_steal":"../static/img/abilities/rubick-spell-steal.jpg",
	//萨尔
	"disruptor_thunder_strike":"../static/img/abilities/disruptor-thunder-strike.jpg",
	"disruptor_glimpse":"../static/img/abilities/disruptor-glimpse.jpg",
	"disruptor_kinetic_field":"../static/img/abilities/disruptor-kinetic-field.jpg",
	"disruptor_static_storm":"../static/img/abilities/disruptor-static-storm.jpg",
	//小强
	"nyx_assassin_impale":"../static/img/abilities/nyx-assassin-impale.jpg",
	"nyx_assassin_mana_burn":"../static/img/abilities/nyx-assassin-mana-burn.jpg",
	"nyx_assassin_spiked_carapace":"../static/img/abilities/nyx-assassin-spiked-carapace.jpg",
	"nyx_assassin_burrow":"../static/img/abilities/nyx_assassin_burrow_lg.png",
	"nyx_assassin_vendetta":"../static/img/abilities/nyx-assassin-vendetta.jpg",
	//小娜迦
	"naga_siren_mirror_image":"../static/img/abilities/naga-siren-mirror-image.jpg",
	"naga_siren_ensnare":"../static/img/abilities/naga-siren-ensnare.jpg",
	"naga_siren_rip_tide":"../static/img/abilities/naga-siren-rip-tide.jpg",
	"naga_siren_song_of_the_siren":"../static/img/abilities/naga-siren-song-of-the-siren.png",
	//光法
	"keeper_of_the_light_illuminate":"../static/img/abilities/keeper-of-the-light-illuminate.jpg",
	"keeper_of_the_light_blinding_light":"../static/img/abilities/keeper-of-the-light-blinding-light.jpg",
	"keeper_of_the_light_chakra_magic":"../static/img/abilities/keeper-of-the-light-chakra-magic.jpg",
	"keeper_of_the_light_will_o_wisp":"../static/img/abilities/keeper-of-the-light-will-o-wisp.jpg",
	//小精灵
	"wisp_tether":"../static/img/abilities/io-tether.jpg",
	"wisp_spirits":"../static/img/abilities/io-spirits.jpg",
	"wisp_overcharge":"../static/img/abilities/io-overcharge.jpg",
	"wisp_relocate":"../static/img/abilities/io-relocate.png",
	//维萨吉
	"visage_grave_chill":"../static/img/abilities/visage-grave-chill.jpg",
	"visage_soul_assumption":"../static/img/abilities/visage-soul-assumption.jpg",
	"visage_gravekeepers_cloak":"../static/img/abilities/visage-gravekeepers-cloak.jpg",
	"visage_stone_form_self_cast":"visage_stone_form_self_cast_lg.png",
	"isage_summon_familiars":"../static/img/abilities/visage-summon-familiars.png",
	//小鱼人
	"slark_dark_pact":"../static/img/abilities/slark-dark-pact.jpg",
	"slark_pounce":"../static/img/abilities/slark-pounce.jpg",
	"slark_essence_shift":"../static/img/abilities/slark-essence-shift.jpg",
	"slark_shadow_dance":"../static/img/abilities/slark-shadow-dance.jpg",
	//美杜莎
	"medusa_split_shot":"../static/img/abilities/medusa-split-shot.jpg",
	"medusa_mystic_snake":"../static/img/abilities/medusa-mystic-snake.jpg",
	"medusa_mana_shield":"../static/img/abilities/medusa-mana-shield.jpg",
	"medusa_stone_gaze":"../static/img/abilities/medusa-stone-gaze.jpg",
	//巨魔
	"troll_warlord_berserkers_rage":"../static/img/abilities/troll-warlord-berserkers-raget.jpg",
	"troll_warlord_whirling_axes_ranged":"../static/img/abilities/troll-warlord-whirling-axes-ranged.jpg",
	"troll_warlord_whirling_axes_melee":"../static/img/abilities/troll_warlord_whirling_axes_melee_lg.png",
	"troll_warlord_fervor":"../static/img/abilities/troll-warlord-fervor.jpg",
	"troll_warlord_battle_trance":"../static/img/abilities/troll-warlord-battle-trance.jpg",
	//人马
	"centaur_hoof_stomp":"../static/img/abilities/centaur-warrunner-hoof-stomp.jpg",
	"centaur_double_edge":"../static/img/abilities/centaur-warrunner-double-edge.jpg",
	"centaur_return":"../static/img/abilities/centaur-warrunner-retaliate.jpg",
	"centaur_stampede":"../static/img/abilities/centaur-warrunner-stampede.jpg",
	//猛犸
	"magnataur_shockwave":"../static/img/abilities/magnus-shockwave.jpg",
	"magnataur_empower":"../static/img/abilities/magnus-empower.jpg",
	"magnataur_skewer":"../static/img/abilities/magnus-skewer.jpg",
	"magnataur_reverse_polarity":"../static/img/abilities/magnus-reverse-polarity.jpg",
	//伐木机
	"shredder_whirling_death":"../static/img/abilities/timbersaw-whirling-death.jpg",
	"shredder_timber_chain":"../static/img/abilities/timbersaw-timber-chain.jpg",
	"shredder_reactive_armor":"../static/img/abilities/timbersaw-reactive-armor.jpg",
	"shredder_chakram":"../static/img/abilities/timbersaw-chakram.jpg",
	//刚被
	"bristleback_viscous_nasal_goo":"../static/img/abilities/bristleback-viscous-nasal.jpg",
	"bristleback_quill_spray":"../static/img/abilities/bristleback-quill-spray.jpg",
	"bristleback_bristleback":"../static/img/abilities/bristleback-bristleback.jpg",
	"bristleback_warpath":"../static/img/abilities/bristleback-warpath.jpg",
	//海民
	"tusk_ice_shards":"../static/img/abilities/tusk-ice-shards.jpg",
	"tusk_snowball":"../static/img/abilities/tusk-snowball.jpg",
	"tusk_tag_team":"../static/img/abilities/tusk-tag-team.jpg",
	"tusk_walrus_kick":"../static/img/abilities/tusk_walrus_kick_lg.png",
	"tusk_walrus_punch":"../static/img/abilities/tusk-walrus-punch.jpg",
	//天鹰
	"skywrath_mage_arcane_bolt":"../static/img/abilities/skywrath-mage-arcane-bolt.jpg",
	"skywrath_mage_concussive_shot":"../static/img/abilities/skywrath-mage-concussive-shot.jpg",
	"skywrath_mage_ancient_seal":"../static/img/abilities/skywrath-mage-ancient-seal.jpg",
	"skywrath_mage_mystic_flare":"../static/img/abilities/skywrath-mage-mystic-flare.jpg",
	//亚巴顿
	"abaddon_death_coil":"../static/img/abilities/abaddon-mist-coil.jpg",
	"abaddon_aphotic_shield":"../static/img/abilities/abaddon-aphotic-shield.jpg",
	"abaddon_frostmourne":"../static/img/abilities/abaddon-curse-of-avernus.jpg",
	"abaddon_borrowed_time":"../static/img/abilities/abaddon-borrowed-time.jpg",
	//大牛
	"elder_titan_echo_stomp":"../static/img/abilities/elder-titan-echo-stomp.jpg",
	"elder_titan_ancestral_spirit":"../static/img/abilities/elder-titan-astral-spirit.jpg",
	"elder_titan_natural_order":"../static/img/abilities/elder-titan-natural-order.jpg",
	"elder_titan_earth_splitter":"../static/img/abilities/elder-titan-earth-splitter.jpg",
	//军团
	"legion_commander_overwhelming_odds":"../static/img/abilities/legion-commander-overwhelming-odds.jpg",
	"legion_commander_press_the_attack":"../static/img/abilities/legion-commander-press-the-attack.jpg",
	"legion_commander_moment_of_courage":"../static/img/abilities/legion-commander-moment-of-courage.jpg",
	"legion_commander_duel":"../static/img/abilities/legion-commander-duel.jpg",
	//火猫
	"ember_spirit_searing_chains":"../static/img/abilities/legion-commander-overwhelming-odds.jpg",
	"ember_spirit_sleight_of_fist":"../static/img/abilities/legion-commander-press-the-attack.jpg",
	"ember_spirit_flame_guard":"../static/img/abilities/legion-commander-moment-of-courage.jpg",
	//"ember_spirit_activate_fire_remnant":"",
	"ember_spirit_fire_remnant":"../static/img/abilities/legion-commander-duel.jpg",
	//大地之灵
	"earth_spirit_boulder_smash":"../static/img/abilities/earth-spirit-boulder-smash.jpg",
	"earth_spirit_rolling_boulder":"../static/img/abilities/earth-spirit-rolling-boulder.jpg",
	"earth_spirit_geomagnetic_grip":"../static/img/abilities/earth-spirit-geomagnetic-grip.jpg",
	"earth_spirit_petrify":"earth_spirit_petrify_lg.png",
	"earth_spirit_magnetize":"../static/img/abilities/earth-spirit-magnetize.jpg",
	//恐怖利刃
	"terrorblade_reflection":"../static/img/abilities/terrorblade-reflection.jpg",
	"terrorblade_conjure_image":"../static/img/abilities/terrorblade-conjure-image.jpg",
	"terrorblade_metamorphosis":"../static/img/abilities/terrorblade-metamorphosis.jpg",
	"terrorblade_sunder":"../static/img/abilities/terrorblade-sunder.jpg",
	//凤凰
	"phoenix_icarus_dive":"../static/img/abilities/phoenix-icarus-dive.jpg",
	"phoenix_fire_spirits":"../static/img/abilities/phoenix-fire-spirits.jpg",
	"phoenix_sun_ray":"../static/img/abilities/phoenix-sun-ray.jpg",
	"phoenix_supernova":"../static/img/abilities/phoenix-supernova.jpg",
	//神谕者
	"oracle_fortunes_end":"../static/img/abilities/oracle-fortunes-end.jpg",
	"oracle_fates_edict":"../static/img/abilities/oracle-fates-edict.jpg",
	"oracle_purifying_flames":"../static/img/abilities/oracle-purifying-flames.jpg",
	"oracle_false_promise":"../static/img/abilities/oracle-false-promise.jpg",
	//炸弹人
	"techies_land_mines":"../static/img/abilities/techies-proximity-mines.jpg",
	"techies_stasis_trap":"../static/img/abilities/techies-stasis-trap.jpg",
	"techies_suicide":"../static/img/abilities/techies-blast-off.jpg",
	"techies_remote_mines":"../static/img/abilities/techies-remote-mines.jpg",
	//冰龙
	"winter_wyvern_arctic_burn":"../static/img/abilities/winter-wyvern-arctic-burn.jpg",
	"winter_wyvern_splinter_blast":"../static/img/abilities/winter-wyvern-splinter-blast.jpg",
	"winter_wyvern_cold_embrace":"../static/img/abilities/winter-wyvern-cold-embrace.jpg",
	"winter_wyvern_winters_curse":"../static/img/abilities/winter-wyvern-winters-curse.jpg",
	//天穹
	"arc_warden_flux":"../static/img/abilities/arc-warden-flux.jpg",
	"arc_warden_magnetic_field":"../static/img/abilities/arc-warden-magnetic-field.jpg",
	"arc_warden_spark_wraith":"../static/img/abilities/arc-warden-spark-wraith.jpg",
	"arc_warden_tempest_double":"../static/img/abilities/arc-warden-tempest-double.jpg",
	//大屁股
	"abyssal_underlord_firestorm":"../static/img/abilities/underlord-firestorm.jpg",
	"abyssal_underlord_pit_of_malice":"../static/img/abilities/underlord-pit-of-malice.jpg",
	"abyssal_underlord_atrophy_aura":"../static/img/abilities/underlord-atrophy-aura.jpg",
	"abyssal_underlord_dark_rift":"../static/img/abilities/underlord-dark-rift.jpg",
	//大圣
	"monkey_king_boundless_strike":"../static/img/abilities/monkey-king-boundless-strike.jpg",
	"monkey_king_tree_dance":"../static/img/abilities/monkey-king-tree-dance.jpg",
	"monkey_king_jingu_mastery":"../static/img/abilities/monkey-king-jingu-mastery.jpg",
	"monkey_king_wukongs_command":"../static/img/abilities/monkey-king-wukongs-command.jpg",
	//石灵剑士
	"pangolier_swashbuckle":"../static/img/abilities/pangolier-swashbuckle.jpg",
	"pangolier_shield_crash":"../static/img/abilities/pangolier-shield-crash.jpg",
	"pangolier_lucky_shot":"../static/img/abilities/pangolier-lucky-shot.jpg",
	"pangolier_gyroshell":"../static/img/abilities/pangolier-rolling-thunder.jpg",
	//
	"dark_willow_bramble_maze":"../static/img/abilities/dark-willow-bramble-maze.jpg",
	"dark_willow_shadow_realm":"../static/img/abilities/dark-willow-shadow-realm.jpg",
	"dark_willow_cursed_crown":"../static/img/abilities/dark-willow-cursed-crown.jpg",
	"dark_willow_bedlam":"../static/img/abilities/dark-willow-bedlam.jpg",
	"dark_willow_terrorize":"../static/img/abilities/dark-willow-terrorize.jpg",
	//石灵剑士
	"mars_spear":"../static/img/abilities/mars-spear-of-mars.jpg",
	"mars_gods_rebuke":"../static/img/abilities/mars-gods-rebuke.jpg",
	"mars_bulwark":"../static/img/abilities/mars-bulwark.jpg",
	"mars_arena_of_blood":"../static/img/abilities/mars-arena-of-blood.jpg",
}
