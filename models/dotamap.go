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
	1:   "antimage.jpg",
	2:   "axe.jpg",
	3:   "bane.jpg",
	4:   "bloodseeker.jpg",
	5:   "crystalmaiden.jpg",
	6:   "drowranger.jpg",
	7:   "earthshaker.jpg",
	8:   "juggernaut.jpg",
	9:   "mirana.jpg",
	10:  "../static/img/morphling.jpg",
	11:  "../static/img/shadowfiend.jpg",
	12:  "../static/img/phantomlancer.jpg",
	13:  "../static/img/puck.jpg",
	14:  "../static/img/pudge.jpg",
	15:  "../static/img/razor.jpg",
	16:  "../static/img/sandking.jpg",
	17:  "../static/img/stormspirit.jpg",
	18:  "../static/img/sven.jpg",
	19:  "../static/img/tiny.jpg",
	20:  "../static/img/vengeful-spirit.jpg",
	21:  "../static/img/windranger.jpg",
	22:  "../static/img/zeus.jpg",
	23:  "../static/img/kunkka.jpg",
	25:  "../static/img/lina.jpg",
	26:  "../static/img/lion.jpg",
	27:  "../static/img/shadowshaman.jpg",
	28:  "../static/img/slardar.jpg",
	29:  "../static/img/tidehunter.jpg",
	30:  "../static/img/witchdoctor.jpg",
	31:  "../static/img/lich.jpg",
	32:  "../static/img/riki.jpg",
	33:  "../static/img/enigma.jpg",
	34:  "../static/img/tinker.jpg",
	35:  "../static/img/sniper.jpg",
	36:  "../static/img/necrophos.jpg",
	37:  "../static/img/warlock.jpg",
	38:  "../static/img/beastmaster.jpg",
	39:  "../static/img/queenofpain.jpg",
	40:  "../static/img/venomancer.jpg",
	41:  "../static/img/gyrocopter.jpg",
	42:  "../static/img/wraithking.jpg",
	43:  "../static/img/deathprophet.jpg",
	44:  "../static/img/phantomassassin.jpg",
	45:  "../static/img/pugna.jpg",
	46:  "../static/img/templarassassin.jpg",
	47:  "../static/img/viper.jpg",
	48:  "../static/img/luna.jpg",
	49:  "../static/img/dragonknight.jpg",
	50:  "../static/img/dazzle.jpg",
	51:  "../static/img/clockwerk.jpg",
	52:  "../static/img/leshrac.jpg",
	53:  "../static/img/naturesprophet.jpg",
	54:  "../static/img/lifestealer",
	55:  "../static/img/dark-seer.jpg",
	56:  "../static/img/clinkz.jpg",
	57:  "../static/img/omniknight.jpg",
	58:  "../static/img/enchantress.jpg",
	59:  "../static/img/huskar.jpg",
	60:  "../static/img/nightstalker.jpg",
	61:  "../static/img/broodmother.jpg",
	62:  "../static/img/bountyhunter.jpg",
	63:  "../static/img/weaver.jpg",
	64:  "../static/img/jakiro.jpg",
	65:  "../static/img/batrider.jpg",
	66:  "../static/img/chen.jpg",
	67:  "../static/img/spectre.jpg",
	68:  "../static/img/ancient.jpg",
	69:  "../static/img/doom.jpg",
	70:  "../static/img/ursa",
	71:  "../static/img/spiritbreaker.jpg",
	72:  "../static/img/gyrocopter.jpg",
	73:  "../static/img/alchemist.jpg",
	74:  "../static/img/invoker.jpg",
	75:  "../static/img/silencer.jpg",
	76:  "../static/img/outworld-devourer.jpg",
	77:  "../static/img/lycan.jpg",
	78:  "../static/img/brewmaster.jpg",
	79:  "../static/img/shadowdemon.jpg",
	80:  "../static/img/lone-druid.jpg",
	81:  "../static/img/chaosknight.jpg",
	82:  "../static/img/meepo.jpg",
	83:  "../static/img/treant-protector.jpg",
	84:  "../static/img/ogremagi.jpg",
	85:  "../static/img/undying.jpg",
	86:  "../static/img/rubick.jpg",
	87:  "../static/img/disruptor.jpg",
	88:  "../static/img/nyx-assassin.jpg",
	89:  "../static/img/naga-siren.jpg",
	90:  "../static/img/keeperofthelight.jpg",
	91:  "../static/img/io.jpg",
	92:  "../static/img/visage.jpg",
	93:  "../static/img/slark.jpg",
	94:  "../static/img/medusa.jpg",
	95:  "../static/img/trollwarlord.jpg",
	96:  "../static/img/centaurwarrunner.jpg",
	97:  "../static/img/magnus.jpg",
	98:  "../static/img/timbersaw.jpg",
	99:  "../static/img/bristleback.jpg",
	100: "../static/img/tusk.jpg",
	101: "../static/img/skywrathmage.jpg",
	102: "../static/img/abaddon.jpg",
	103: "../static/img/eldertitan.jpg",
	104: "../static/img/legioncommander.jpg",
	105: "../static/img/techies.jpg",
	106: "../static/img/emberspirit.jpg",
	107: "../static/img/earthspirit.jpg",
	108: "../static/img/underlord.jpg",
	109: "../static/img/terrorblade.jpg",
	110: "../static/img/phoenix.jpg",
	111: "../static/img/oracle.jpg",
	112: "../static/img/winterwyvern.jpg",
	113: "../static/img/arcwarden.jpg",
	114: "../static/img/monkey-king.jpg",
	119: "../static/img/dark-willow.jpg",
	120: "../static/img/pangolier.jpg",
}

var ItemIDIcon = map[int64]string{
	//大飞鞋 ，银月之晶，陨星锤 ，肉山盾，刷新碎片，精灵布带,凝魂之泪，慧光
	1:  "../static/img/blink-dagger.jpg",           //跳刀
	2:  "../static/img/blades-of-attack.jpg",       //攻击之爪
	3:  "../static/img/broadsword.jpg",             //阔剑
	4:  "../static/img/chainmail.jpg",              //锁子甲
	5:  "../static/img/claymore.jpg",               //大剑
	6:  "../static/img/helm_of_iron_will.jpg",      //铁意头盔
	7:  "../static/img/javelin.jpg",                //标枪
	8:  "../static/img/mithril-hammer.jpg",         //秘银锤
	9:  "../static/img/platemail.jpg",              //板甲
	10: "../static/img/quarterstaff.jpg",           //短棍
	11: "../static/img/quelling-blade.jpg",         //补刀斧
	12: "../static/img/ring_of_protection.jpg",     //守护戒指
	13: "../static/img/gauntlets-of-strength.jpg",  //力量手套
	14: "../static/img/slippers-of-agility.jpg",    //敏捷便鞋
	15: "../static/img/mantle-of-intelligence.jpg", //智力斗篷
	16: "../static/img/iron-branch.jpg",            //树枝
	17: "../static/img/belt-of-strength.jpg",       //力量腰带
	18: "../static/img/band-of-elvenskin.jpg",      //精灵布带
	19: "../static/img/robe-of-the-magi",           //法师长袍
	20: "../static/img/circlet.jpg",                //圆环
	21: "../static/img/ogre-axe.jpg",               //食人魔之斧
	22: "../static/img/blade-of-alacrity.jpg",      //欢欣之刃
	23: "../static/img/staff-of-wizardry.jpg",      //魔力法杖
	24: "../static/img/ultimate-orb.jpg",           //极限法球
	25: "../static/img/gloves-of-haste.jpg",        //加速手套
	26: "../static/img/morbid-mask.jpg",            //吸血面具
	27: "../static/img/ring-of-regen.jpg",          //回复戒指
	28: "../static/img/sages-mask.jpg",             //艺人面罩
	29: "../static/img/boots-of-speed.jpg",         //速度之鞋
	30: "../static/img/gem-of-true-sight.jpg",      //珍视宝石
	31: "../static/img/cloak.jpg",                  //抗魔斗篷
	32: "../static/img/talisman-of-evasion.jpg",    //闪避护符
	33: "../static/img/cheese.jpg",                 //奶酪
	34: "../static/img/magic-stick.jpg",            //魔棒

	36: "../static/img/magic-wand.jpg",         //大魔棒
	37: "../static/img/ghost-scepter.jpg",      //幽魂权杖
	38: "../static/img/clarity.jpg",            //净化药水
	39: "../static/img/healing-salve.jpg",      //大药
	40: "../static/img/dust-of-appearance.jpg", //显影之尘
	41: "../static/img/bottle.jpg",             //瓶子
	42: "../static/img/observer-ward.jpg",      //假眼
	43: "../static/img/sentry-ward.jpg",        //真眼
	44: "../static/img/tango.jpg",              //吃数
	45: "../static/img/animal-courier.jpg",     //动物信使
	46: "../static/img/town-portal-scrol.jpg",  //TP
	48: "../static/img/boots-of-travel.jpg",    //飞鞋

	50: "../static/img/phase-boots.jpg",      //相位鞋
	51: "../static/img/demon-edge.jpg",       //恶魔刀锋
	52: "../static/img/eaglesong.jpg",        //鹰歌弓
	53: "../static/img/reaver.jpg",           //掠夺
	54: "../static/img/sacred-relic.jpg",     //圣者遗物
	55: "../static/img/hyperstone.jpg",       //振奋宝石
	56: "../static/img/ring-of-health.jpg",   //治疗指环
	57: "../static/img/void-stone.jpg",       //虚无宝石
	58: "../static/img/mystic-staff.jpg",     //神秘法杖
	59: "../static/img/energy-booster.jpg",   //能量之球
	60: "../static/img/point-booster.jpg",    //精气之球
	61: "../static/img/vitality-booster.jpg", //活力之球
	63: "../static/img/power-treads.jpg",     //假腿

	65: "../static/img/hand-of-midas.jpg", //点金手

	67:  "../static/img/oblivion-staff.jpg",           //空明帐
	69:  "../static/img/perseverance.jpg",             //坚韧球
	73:  "../static/img/bracer.jpg",                   //护腕
	75:  "../static/img/wraith-band.jpg",              //怨灵系带
	77:  "../static/img/null-talisman.jpg",            //空灵挂件
	79:  "../static/img/mekansm.jpg",                  //梅肯
	81:  "../static/img/vladmirs-offering.jpg",        //祭品
	84:  "../static/img/flying_courier_lg.png",        //动物信使
	86:  "../static/img/buckler.jpg",                  //玄冥盾牌
	88:  "../static/img/ring-of-basilius.jpg",         //圣殿
	90:  "../static/img/pipe-of-insight.jpg",          //笛子
	92:  "../static/img/urn-of-shadows.jpg",           //骨灰
	94:  "../static/img/headdress.jpg",                //回复头巾
	96:  "../static/img/scythe-of-vyse.jpg",           //羊刀
	98:  "../static/img/orchid-malevolence.jpg",       //紫苑
	250: "../static/img/bloodthorn.jpg",               //血棘
	252: "../static/img/echo-sabre.jpg",               //回音战刃
	100: "../static/img/euls-scepter-of-divinity.jpg", //吹风
	232: "../static/img/aether-lens.jpg",              //以太之镜
	102: "../static/img/force-staff.jpg",              //推推
	263: "../static/img/hurricane-pike.jpg",           //大推推
	104: "../static/img/dagon.jpg",                    //大根
	201: "../static/img/dagon-level-2.jpg",
	202: "../static/img/dagon-level-3.jpg",
	203: "../static/img/dagon-level-4.jpg",
	204: "../static/img/dagon-level-5.jpg",
	106: "../static/img/necronomicon.jpg", //死灵书
	193: "../static/img/necronomicon-level-2.jpg",
	194: "../static/img/necronomicon-level-3.jpg",
	108: "../static/img/aghanims-scepter.jpg",        //A帐
	110: "../static/img/refresher-orb.jpg",           //刷新球
	117: "../static/img/aegis-of-the-immortal.jpg",   //肉山盾
	112: "../static/img/assault-cuirass.jpg",         //强袭装甲
	114: "../static/img/heart-of-tarrasque.jpg",      //龙心
	116: "../static/img/black-king-bar.jpg",          //BKB
	118: "../static/img/",                            //
	119: "../static/img/shivas-guard.jpg",            //西瓦
	121: "../static/img/bloodstone.jpg",              //血精石
	123: "../static/img/linkens-sphere.jpg",          //林肯
	226: "../static/img/lotus-orb.jpg",               //莲花
	125: "../static/img/vanguard.jpg",                //先锋盾
	242: "../static/img/crimson-guard.jpg",           //赤红甲
	127: "../static/img/blade-mail.jpg",              //刃甲
	129: "../static/img/soul-booster.jpg",            //镇魂石
	131: "../static/img/hood-of-defiance.jpg",        //挑战
	133: "../static/img/divine-rapier.jpg",           //圣剑
	135: "../static/img/monkey-king-bar.jpg",         //金箍棒
	137: "../static/img/radiance.jpg",                //辉耀
	139: "../static/img/butterfly.jpg",               //蝴蝶
	141: "../static/img/daedalus.jpg",                //大炮
	143: "../static/img/skull-basher.jpg",            //碎颅锤
	145: "../static/img/battle-fury.jpg",             //狂战斧
	147: "../static/img/manta-style.jpg",             //分身斧
	149: "../static/img/crystalys.jpg",               //水晶剑
	236: "../static/img/dragon-lance.jpg",            //魔龙枪
	151: "../static/img/armlet-of-mordiggian.jpg",    //臂章
	152: "../static/img/shadow-blade.jpg",            //隐刀
	249: "../static/img/silver-edge.jpg",             //大隐刀，白银之锋
	154: "../static/img/sange-and-yasha.jpg",         //双刀
	156: "../static/img/satanic.jpg",                 //撒旦
	158: "../static/img/mjollnir.jpg",                //大电锤
	160: "../static/img/eye-of-skadi.jpg",            //冰眼
	162: "../static/img/sange.jpg",                   //散华
	164: "../static/img/helm-of-the-dominator.jpg",   //支配
	166: "../static/img/maelstrom.jpg",               //电锤
	168: "../static/img/desolator.jpg",               //暗灭
	170: "../static/img/yasha.jpg",                   //单刀夜叉
	172: "../static/img/mask-of-madness.jpg",         //疯脸
	174: "../static/img/diffusal-blade.jpg",          //散失
	196: "../static/img/diffusal-blade-2.jpg",        //
	176: "../static/img/ethereal-blade.jpg",          //虚灵刀
	178: "../static/img/soul-ring.jpg",               //魂戒
	180: "../static/img/arcane-boots.jpg",            //秘法
	235: "../static/img/octarine-core.jpg",           //玲珑心
	181: "../static/img/orb-of-venom.jpg",            //毒球
	240: "../static/img/blight-stone.jpg",            //枯萎之石
	185: "../static/img/drum-of-endurance.jpg",       //战鼓
	187: "../static/img/medallion-of-courage.jpg",    //勇气勋章
	229: "../static/img/solar-crest.jpg",             //炎阳纹章
	188: "../static/img/smoke-of-deceit.jpg",         //雾
	257: "../static/img/tome-of-knowledge.jpg",       //经验书
	190: "../static/img/veil-of-discord.jpg",         //纷争面纱
	231: "../static/img/guardian-greaves.jpg",        //大鞋
	206: "../static/img/rod-of-atos.jpg",             //阿托斯
	239: "../static/img/",                            //打野爪
	208: "../static/img/abyssal-blade.jpg",           //深渊之刃
	210: "../static/img/heavens-halberd.jpg",         //天堂
	212: "../static/img/ring-of-aquila.jpg",          //天鹰
	214: "../static/img/tranquil-boots.jpg",          //绿鞋
	215: "../static/img/shadow-amulet.jpg",           //暗影护符
	254: "../static/img/glimmer-cape.jpg",            //微光披风
	256: "../static/img/aeon-disk.jpg",               //免死金牌
	225: "../static/img/nullifier.jpg",               //否决
	267: "../static/img/spirit-vessel.jpg",           //大骨灰
	244: "../static/img/wind-lace.jpg",               //风灵之纹
	259: "../static/img/kaya.jpg",                    //慧光
	260: "../static/img/refresher-shard.jpg",         //刷新碎片
	223: "../static/img/meteor-hammer.jpg",           //陨星锤
	220: "../static/img/boots-of-travel-level-2.jpg", //大飞鞋
	265: "../static/img/infused-raindrops.jpg",       //凝魂之泪
	247: "../static/img/moon-shard.jpg",              //银月之晶
}
