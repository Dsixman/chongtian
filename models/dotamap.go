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
	10:  "morphling.jpg",
	11:  "shadowfiend.jpg",
	12:  "phantomlancer.jpg",
	13:  "puck.jpg",
	14:  "pudge.jpg",
	15:  "razor.jpg",
	16:  "sandking.jpg",
	17:  "stormspirit.jpg",
	18:  "sven.jpg",
	19:  "tiny.jpg",
	20:  "vengeful-spirit.jpg",
	21:  "windranger.jpg",
	22:  "zeus.jpg",
	23:  "kunkka.jpg",
	25:  "lina.jpg",
	26:  "lion.jpg",
	27:  "shadowshaman.jpg",
	28:  "slardar.jpg",
	29:  "tidehunter.jpg",
	30:  "witchdoctor.jpg",
	31:  "lich.jpg",
	32:  "riki.jpg",
	33:  "enigma.jpg",
	34:  "tinker.jpg",
	35:  "sniper.jpg",
	36:  "necrophos.jpg",
	37:  "warlock.jpg",
	38:  "beastmaster.jpg",
	39:  "queenofpain.jpg",
	40:  "venomancer.jpg",
	41:  "gyrocopter.jpg",
	42:  "wraithking.jpg",
	43:  "deathprophet.jpg",
	44:  "phantomassassin.jpg",
	45:  "pugna.jpg",
	46:  "templarassassin.jpg",
	47:  "viper.jpg",
	48:  "luna.jpg",
	49:  "dragonknight.jpg",
	50:  "dazzle.jpg",
	51:  "clockwerk.jpg",
	52:  "leshrac.jpg",
	53:  "naturesprophet.jpg",
	54:  "lifestealer",
	55:  "dark-seer.jpg",
	56:  "clinkz.jpg",
	57:  "omniknight.jpg",
	58:  "enchantress.jpg",
	59:  "huskar.jpg",
	60:  "nightstalker.jpg",
	61:  "broodmother.jpg",
	62:  "bountyhunter.jpg",
	63:  "weaver.jpg",
	64:  "jakiro.jpg",
	65:  "batrider.jpg",
	66:  "chen.jpg",
	67:  "spectre.jpg",
	68:  "ancient.jpg",
	69:  "doom.jpg",
	70:  "ursa",
	71:  "spiritbreaker.jpg",
	72:  "gyrocopter.jpg",
	73:  "alchemist.jpg",
	74:  "invoker.jpg",
	75:  "silencer.jpg",
	76:  "outworld-devourer.jpg",
	77:  "lycan.jpg",
	78:  "brewmaster.jpg",
	79:  "shadowdemon.jpg",
	80:  "lone-druid.jpg",
	81:  "chaosknight.jpg",
	82:  "meepo.jpg",
	83:  "treant-protector.jpg",
	84:  "ogremagi.jpg",
	85:  "undying.jpg",
	86:  "rubick.jpg",
	87:  "disruptor.jpg",
	88:  "nyx-assassin.jpg",
	89:  "naga-siren.jpg",
	90:  "keeperofthelight.jpg",
	91:  "io.jpg",
	92:  "visage.jpg",
	93:  "slark.jpg",
	94:  "medusa.jpg",
	95:  "trollwarlord.jpg",
	96:  "centaurwarrunner.jpg",
	97:  "magnus.jpg",
	98:  "timbersaw.jpg",
	99:  "bristleback.jpg",
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
}

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
	256: "aeon-disk.jpg",               //免死金牌
	225: "nullifier.jpg",               //否决
	267: "spirit-vessel.jpg",           //大骨灰
	244: "wind-lace.jpg",               //风灵之纹
	259: "kaya.jpg",                    //慧光
	260: "refresher-shard.jpg",         //刷新碎片
	223: "meteor-hammer.jpg",           //陨星锤
	220: "boots-of-travel-level-2.jpg", //大飞鞋
	265: "infused-raindrops.jpg",       //凝魂之泪
	247: "moon-shard.jpg",              //银月之晶
}
