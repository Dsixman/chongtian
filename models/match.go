package models

import(
	"fmt"
	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson"
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
	"strings"
	"sort"
)

type TenMinData struct{
	HeroName string `json:"hero_name"`
	PlayerId uint32 `json:"player_id"`
	PlayerName string `json:"player_name"`
	HeroIcon string `json:"hero_icon"`
	GameTeam uint32 `josn:"game_team"`
	Pre5MinLastHit uint32 `json:"pre_5min_lasthit"`
	Pre10MinLastHit uint32 `json:"pre_10min_lasthit"`
	Pre5MinGold *obj.GoldType `json:"pre_5min_gold"`
	Pre10MinGold *obj.GoldType `json:"pre_10min_gold"`
	Pre10MinLevelUpTime []string `json:"pre_10min_levelup_time"`
	Pre10minItem []*obj.ItemDetails `json:"pre_10min_item"`
	ConsumableItem []*obj.PlayerConsumable `json:"consumable"`
	AbilityUpgrades []*obj.Ability `json:"ability_upgrades" `
	Pre10MinKDAInfo []*obj.KDADetails `josn:"pre_10min_kda_info"`
	DeathCounts int `json:"death_counts"`
}
type MidLaneHeroData struct{
	MatchId uint64 `bson:"match_id" json:"match_id"`
	MidHero1 *TenMinData `bson:"mid_hero_1" json:"mid_hero_1"`
	MidHero2 *TenMinData `bson:"mid_hero_2" json:"mid_hero_2"`
	MidHero3 *TenMinData `bson:"mid_hero_3" json:"mid_hero_3"`
	MidHero4 *TenMinData `bson:"mid_hero_3 json:"mid_hero_4"`
}
type EdgeLaneHeroData struct{
	MatchId uint64 `bson:"match_id" json:"match_id"`
	Team1DeathCounts int `bson:"team1_death_counts" json:"team1_death_counts"`
	Team2DeathCounts int `bson:"team2_death_counts" json:"team2_death_counts"`
	SameSideHero1 *TenMinData `json:"same_side_hero_1"`
	SameSideHero2 *TenMinData `json:"same_side_hero_2"`
	SameSideHero3 *TenMinData `json:"same_side_hero_3"`
	SameSideHero4 *TenMinData `json:"same_side_hero_4"`
	SameSideHero5 *TenMinData `json:"same_side_hero_5"`
	SameSideHero6 *TenMinData `json:"same_side_hero_6"`
}
type AgainstKda struct{
	DeathTarget string `bson:"death_target" json:"death_target"`
	Assists []string `bson:"assists" json:"assists"`
}
type SameLaneHeroData struct{
	MidHeroesData []*MidLaneHeroData `json:"mid_hero_data"`
	EdgeHeroesData []*EdgeLaneHeroData `json:"edge_heroes_data"`
	OtherEdgeHeroesData []*EdgeLaneHeroData `other_edge_heroes_data`
}

type ReqTenMinData struct{
	TopLane string `json:"toplane"`
	MidpLane string `json:"midlane"`
	MidpLane2 string `json:"midlane2"`
	BottomLane string `json:"bottomlane"`
	TopLaneSelHero string `json:toplaneselhero`
	MidLaneSelHero string `json:midlaneselhero`
	BottomLaneSelHero string `json:bottomlaneselhero`
	SideLane string `json:"sidelane"`
	SoftLane string `json:"softlane"`
	HardLane string `json:"hardlane"`
	Version string `josn:"version"`
}
func GetTenMinData (req *ReqTenMinData) *SameLaneHeroData{
	var topallres []obj.CMsgMatchDetails
	var bottomallres []obj.CMsgMatchDetails
	var midallres []obj.CMsgMatchDetails
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	ColAbility:=db.C("abilities") 
	version:=req.Version
	toplane:=req.TopLane
	midlane:=req.MidpLane
	midlane2:=req.MidpLane2
	bottomlane:=req.BottomLane
	toplaneselhero:=req.TopLaneSelHero
	midlaneselhero:=req.MidLaneSelHero
	bottomlaneselhero:=req.BottomLaneSelHero

	var topselherosli []string
	var midselherosli []string
	var bottomselherosli []string
	var LaneHeroData SameLaneHeroData
	toperr:=ColCMsgMatchDetails.Find(bson.M{"version":version,"$or":[]bson.M{bson.M{"top_lane":toplane},bson.M{"bottom_lane":toplane}}}).All(&topallres)
	bottomerr:=ColCMsgMatchDetails.Find(bson.M{"version":version,"$or":[]bson.M{bson.M{"top_lane":bottomlane},bson.M{"bottom_lane":bottomlane}}}).All(&bottomallres)
	miderr:=ColCMsgMatchDetails.Find(bson.M{"version":version,"$or":[]bson.M{bson.M{"mid_lane":midlane},bson.M{"mid_lane":midlane2}}}).All(&midallres)
	if toperr!=nil{
		fmt.Printf("toperr:%v\n", toperr)
	}
	if miderr!=nil{
		fmt.Printf("miderr:%v\n", miderr)
	}
	if bottomerr!=nil{
		fmt.Printf("bottomerr:%v\n", bottomerr)
	}
	ability:=&obj.AllNpcAbility{}
	err:=ColAbility.Find(bson.M{"Version":version}).One(&ability)
	if err!=nil{
		fmt.Printf("技能查询错误:%v\n", err)
	}
	if toplaneselhero!=""{
		topselherosli=strings.Split(toplaneselhero," ")
		for _,v:=range topallres{
			res:=&EdgeLaneHeroData{}
			res.MatchId=v.MatchId
			
			for _,v1:=range v.PlayersHeroesDets{
				for k2,v2:=range topselherosli {
					if v2==v1.HeroName{
						edgehero:=&TenMinData{}	
						Pre10MinKDAInfo:=v1.DeathDets					
						edgehero.HeroName=v1.HeroName
						edgehero.PlayerName=v1.Player
						edgehero.Pre5MinLastHit=v1.LastHit.Pre5Count
						edgehero.Pre10MinLastHit=v1.LastHit.Pre10Count
						edgehero.Pre5MinGold=v1.Gold.Pre5MinGold
						edgehero.Pre10MinGold=v1.Gold.Pre10MinGold
						edgehero.ConsumableItem=v1.Consumable
						edgehero.HeroIcon=v1.HeroIcon
						lvluptimeslice:=v1.LevelUpTimes
						Looplvl:
						for key,value:=range lvluptimeslice{
							if value>600{
								lvlk:=key-1
								edgehero.Pre10MinLevelUpTime=v1.LevelUpTimeStr[0:lvlk]
								edgehero.AbilityUpgrades=v1.AbilityData[0:key]
								break Looplvl
							}
						}
						itemslice:=v1.Item
						Loopitem:
						for itemkey,itemvalue:=range itemslice{
							if itemvalue.PurchaseTimeNum>600{
								edgehero.Pre10minItem=itemslice[0:itemkey]
								break Loopitem
							}
						}
						Loopkda:
						for kdakey,kdavlaue:=range Pre10MinKDAInfo{
							if kdavlaue.Time>600{
								edgehero.Pre10MinKDAInfo=Pre10MinKDAInfo[0:kdakey]
								break Loopkda
							}
						}
						if k2==0{
							res.SameSideHero1=edgehero
						}
						if k2==1{
							res.SameSideHero2=edgehero
						}
						if k2==2{
							res.SameSideHero3=edgehero
						}
						if k2==3{
							res.SameSideHero4=edgehero
						}
						if k2==4{
							res.SameSideHero5=edgehero
						}
						if k2==5{
							res.SameSideHero6=edgehero
						}
					}
				}
			}
			
			LaneHeroData.EdgeHeroesData=append(LaneHeroData.EdgeHeroesData,res)
			
		}	
	}

	if midlaneselhero!=""{
		midselherosli=strings.Split(midlaneselhero," ")
		for _,v:=range midallres{
			res:=&MidLaneHeroData{}
			res.MatchId=v.MatchId
			for _,v1:=range v.PlayersHeroesDets{
				for k2,v2:=range midselherosli {
					if v2==v1.HeroName{
						mid:=&TenMinData{}			
						mid.HeroName=v1.HeroName
						Pre10MinKDAInfo:=v1.DeathDets
						mid.PlayerName=v1.Player
						mid.Pre5MinLastHit=v1.LastHit.Pre5Count
						mid.Pre10MinLastHit=v1.LastHit.Pre10Count
						mid.Pre5MinGold=v1.Gold.Pre5MinGold
						mid.Pre10MinGold=v1.Gold.Pre10MinGold
						mid.ConsumableItem=v1.Consumable
						mid.HeroIcon=v1.HeroIcon
						lvluptimeslice:=v1.LevelUpTimes
						Looplvl1:
						for key,value:=range lvluptimeslice{
							if value>600{
								lvlk:=key-1
								mid.Pre10MinLevelUpTime=v1.LevelUpTimeStr[0:lvlk]
								mid.AbilityUpgrades=v1.AbilityData[0:key]
								break Looplvl1
							}
						}
						itemslice:=v1.Item
						Loopitem1:
						for itemkey,itemvalue:=range itemslice{
							if itemvalue.PurchaseTimeNum>600{
								mid.Pre10minItem=itemslice[0:itemkey]
								break Loopitem1
							}
						}
						Loopkda1:
						for kdakey,kdavlaue:=range Pre10MinKDAInfo{
							if kdavlaue.Time>600{
								mid.Pre10MinKDAInfo=Pre10MinKDAInfo[0:kdakey]
								break Loopkda1
							}
						}
						if k2==0{
							res.MidHero1=mid
						}
						if k2==1{
							res.MidHero2=mid
						}
						if k2==2{
							res.MidHero3=mid
						}
						if k2==3{
							res.MidHero4=mid
						}
					}
				}
			}
			LaneHeroData.MidHeroesData=append(LaneHeroData.MidHeroesData,res)
		}
	}
	if bottomlaneselhero!=""{
		bottomselherosli=strings.Split(bottomlaneselhero," ")
		for _,v:=range midallres{
			res:=&EdgeLaneHeroData{}
			res.MatchId=v.MatchId	
			for _,v1:=range v.PlayersHeroesDets{
				for k2,v2:=range bottomselherosli {
					if v2==v1.HeroName{
						edgehero:=&TenMinData{}
						edgehero.HeroName=v1.HeroName
						edgehero.PlayerName=v1.Player
						Pre10MinKDAInfo:=v1.DeathDets
						edgehero.Pre5MinLastHit=v1.LastHit.Pre5Count
						edgehero.Pre10MinLastHit=v1.LastHit.Pre10Count
						edgehero.Pre5MinGold=v1.Gold.Pre5MinGold
						edgehero.Pre10MinGold=v1.Gold.Pre10MinGold
						edgehero.HeroIcon=v1.HeroIcon
						edgehero.ConsumableItem=v1.Consumable
						lvluptimeslice:=v1.LevelUpTimes
						Looplvl2:
						for key,value:=range lvluptimeslice{
							if value>600{
								lvlk:=key-1
								edgehero.Pre10MinLevelUpTime=v1.LevelUpTimeStr[0:lvlk]
								edgehero.AbilityUpgrades=v1.AbilityData[0:key]
								break Looplvl2
							}
						}
						itemslice:=v1.Item
						Loopitem2:
						for itemkey,itemvalue:=range itemslice{
							if itemvalue.PurchaseTimeNum>600{
								edgehero.Pre10minItem=itemslice[0:itemkey]
								break Loopitem2
							}
						}
						Loopkda2:
						for kdakey,kdavlaue:=range Pre10MinKDAInfo{
							if kdavlaue.Time>600{
								edgehero.Pre10MinKDAInfo=Pre10MinKDAInfo[0:kdakey]
								break Loopkda2
							}
						}
						if k2==0{
							res.SameSideHero1=edgehero
						}
						if k2==1{
							res.SameSideHero2=edgehero
						}
						if k2==2{
							res.SameSideHero3=edgehero
						}
						if k2==3{
							res.SameSideHero4=edgehero
						}
						if k2==4{
							res.SameSideHero5=edgehero
						}
						if k2==5{
							res.SameSideHero6=edgehero
						}
					}
				}
			}
			LaneHeroData.OtherEdgeHeroesData=append(LaneHeroData.OtherEdgeHeroesData,res)
		}		
	}
	
	return &LaneHeroData
	//CDotaGameInfo
}
type MatchHeroData struct{
	WinData []obj.CMsgMatchDetails `bson:"win_data" json:"win_data"`
	LoseData []obj.CMsgMatchDetails `bson:"lose_data" json:"lose_data"`
}
func GetMatchHeroData (heroname string,version string,position string) *MatchHeroData{
	var windata []obj.CMsgMatchDetails
	var losedata []obj.CMsgMatchDetails
	var matchHeroData MatchHeroData

	session := mongodb.CloneSession()
	defer session.Close()  
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	//fmt.Printf("position:%v\n", position)
	if position!=""{
		err:=ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"hero_name":heroname,"game_result":"win","position":position}}}).All(&windata)
		if err!=nil{
			fmt.Printf("查询比赛中的英雄数据出错：%v\n",err)
		}	
		err2:=ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"hero_name":heroname,"game_result":"lose","position":position}}}).All(&losedata)
		if err2!=nil{
			fmt.Printf("查询比赛中的英雄数据出错：%v\n",err2)
		}
	}else{
		fmt.Printf("position:%v\n", position)
		err:=ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"hero_name":heroname,"game_result":"win"}}}).All(&windata)
		if err!=nil{
			fmt.Printf("查询比赛中的英雄数据出错：%v\n",err)
		}	
		err2:=ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"hero_name":heroname,"game_result":"win"}}}).All(&losedata)
		if err2!=nil{
			fmt.Printf("查询比赛中的英雄数据出错：%v\n",err2)
		}
	}	
	matchHeroData.WinData=windata
	matchHeroData.LoseData=losedata
	return &matchHeroData
}
func GetPlayerSameHero (heroname string,version string,playerid uint32) *MatchHeroData{
	var windata []obj.CMsgMatchDetails
	var losedata []obj.CMsgMatchDetails
	var matchHeroData MatchHeroData

	session := mongodb.CloneSession()
	defer session.Close()  
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	//fmt.Printf("position:%v\n", position)
	err:=ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"hero_name":heroname,"game_result":"win","account_id":playerid}}}).All(&windata)
	if err!=nil{
		fmt.Printf("查询比赛中的英雄数据出错：%v\n",err)
	}	
	err2:=ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"hero_name":heroname,"game_result":"lose","account_id":playerid}}}).All(&losedata)
	if err2!=nil{
		fmt.Printf("查询比赛中的英雄数据出错：%v\n",err2)
	}
	matchHeroData.WinData=windata
	matchHeroData.LoseData=losedata
	return &matchHeroData
}
func GetMatchDetails(matchid uint64)(*obj.CDotaGameInfo,*obj.CMsgDOTAMatch,*obj.CMsgMatchDetails){
	session := mongodb.CloneSession()
	defer session.Close()  
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	var MatchDetails obj.CMsgMatchDetails
	var MatchBaseData obj.CMsgDOTAMatch
	var GameInfo obj.CDotaGameInfo
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	ColCMsgDOTAMatch:=db.C("CMsgDOTAMatch")
	ColCDotaGameInfo:=db.C("CDotaGameInfo")
	err:=ColCMsgMatchDetails.Find(bson.M{"match_id":matchid}).One(&MatchDetails)
	if err!=nil{
		fmt.Printf("没有找到这场比赛：%v\n", err)
	}
	err2:=ColCMsgDOTAMatch.Find(bson.M{"match_id":matchid}).One(&MatchBaseData)
	if err2!=nil{
		fmt.Printf("没有找到这场比赛：%v\n", err2)
	}
	err3:=ColCDotaGameInfo.Find(bson.M{"match_id":matchid}).One(&GameInfo)
	if err3!=nil{
		fmt.Printf("没有找到这场比赛：%v\n", err3)
	}
	return &GameInfo,&MatchBaseData,&MatchDetails
}
type LineUpGoldDets struct{
	HeroName string `bson:"hero_name" json:"hero_name"`
	GoldLessThan3000Counts int `bson:"gold_less_than_3000_counts" json:"gold_less_than_3000_counts"`
	Gold3000to3100Counts int `josn:"gold_3000_3100_counts"`
	Gold3100to3200Counts int `josn:"gold_3100_3200_counts"`
	Gold3200to3300Counts int `josn:"gold_3200_3300_counts"`
	Gold3300to3400Counts int `josn:"gold_3300_3400_counts"`
	Gold3400to3500Counts int `josn:"gold_3400_3500_counts"`
	Gold3500to3600Counts int `josn:"gold_3500_3600_counts"`
	Gold3600to3700Counts int `josn:"gold_3600_3700_counts"`
	Gold3700to3800Counts int `josn:"gold_3700_3800_counts"`
	Gold3800to3900Counts int `josn:"gold_3800_3900_counts"`
	Gold3900to4000Counts int `josn:"gold_3900_4000_counts"`
	GoldMoreThan4000Counts int `josn:"gold_more_than_4000_counts"`
	GoldMaximumProportion string `json:"gold_maximum_proportion"`//最大
	GoldSecondMaxProportion string `json:"gold_second_max_proportion"`
	Death1Counts int `json:"death_1counts"`
	Death2Counts int `json:"death_2counts"`
	Death3Counts int `json:"death_3counts"`
	DeathMoreThan3Counts int `json:"death_more_than_3counts"`
}
type LineUpOverView struct{
	LineUpHeroes string `bson:"line_up_heroes" json:"line_up_heroes"`
	Counts int `bson:"counts" json:"counts"`
	Hero1GoldDets *LineUpGoldDets `bson:"hero1_gold_dets" json:"hero1_gold_dets"`
	Hero2GoldDets *LineUpGoldDets `bson:"hero2_gold_dets" json:"hero2_gold_dets"`
	Hero3GoldDets *LineUpGoldDets `bson:"hero3_gold_dets" json:"hero3_gold_dets"`
	Hero4GoldDets *LineUpGoldDets `bson:"hero4_gold_dets" json:"hero4_gold_dets"`
	Hero5GoldDets *LineUpGoldDets `bson:"hero5_gold_dets" json:"hero5_gold_dets"`
	Hero6GoldDets *LineUpGoldDets `bson:"hero6_gold_dets" json:"hero6_gold_dets"`
}
func GetSideLineUpOverView(version string,lineUp string) *LineUpOverView{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColSideLineUpData:=db.C("SideLineUpData")
	var results []obj.SideLineUp
	var count int
	newHero1GoldDets:=&LineUpGoldDets{}
	newHero2GoldDets:=&LineUpGoldDets{}
	newHero3GoldDets:=&LineUpGoldDets{}
	newHero4GoldDets:=&LineUpGoldDets{}
	newHero5GoldDets:=&LineUpGoldDets{}
	newHero6GoldDets:=&LineUpGoldDets{}
	ColSideLineUpData.Find(bson.M{"version":version,"line_up_heroes":lineUp}).All(&results)
	newLineUpOverView:=&LineUpOverView{}
	newLineUpOverView.LineUpHeroes=lineUp
	//strings.Split(lineUp,"vs")
	
	newLineUpOverView.Hero1GoldDets=newHero1GoldDets
	newLineUpOverView.Hero2GoldDets=newHero2GoldDets
	newLineUpOverView.Hero3GoldDets=newHero3GoldDets
	newLineUpOverView.Hero4GoldDets=newHero4GoldDets
	newLineUpOverView.Hero5GoldDets=newHero5GoldDets
	newLineUpOverView.Hero6GoldDets=newHero6GoldDets
	newLineUpOverView.Hero1GoldDets=newHero1GoldDets

	sidearr:=strings.Split(lineUp,"vs")
	side1heroes:=strings.Split(sidearr[0], " ")
	side2heroes:=strings.Split(sidearr[1], " ")
	newLineUpOverView.Hero1GoldDets.HeroName=side1heroes[0]
	newLineUpOverView.Hero4GoldDets.HeroName=side2heroes[0]
	if len(side1heroes)==2{
		newLineUpOverView.Hero2GoldDets.HeroName=side1heroes[1]
	}
	if len(side1heroes)==3{
		newLineUpOverView.Hero2GoldDets.HeroName=side1heroes[1]
		newLineUpOverView.Hero3GoldDets.HeroName=side1heroes[2]
	}

	if len(side2heroes)==2{
		newLineUpOverView.Hero5GoldDets.HeroName=side2heroes[1]
	}
	if len(side2heroes)==3{
		newLineUpOverView.Hero5GoldDets.HeroName=side2heroes[1]
		newLineUpOverView.Hero6GoldDets.HeroName=side2heroes[2]
	}
	for _,v:=range results{
		count=count+len(v.LineUpDets)		
		for _,v1:=range v.LineUpDets{
			for _,v2:=range v1.AgainstHeroes{
				if v2.HeroName==newLineUpOverView.Hero1GoldDets.HeroName{
					if v2.Pre10MinGold.AllGold<3000.0{
						newLineUpOverView.Hero1GoldDets.GoldLessThan3000Counts=newLineUpOverView.Hero1GoldDets.GoldLessThan3000Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3000.0&&v2.Pre10MinGold.AllGold<3100.0{
						newLineUpOverView.Hero1GoldDets.Gold3000to3100Counts=newLineUpOverView.Hero1GoldDets.Gold3000to3100Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3100.0&&v2.Pre10MinGold.AllGold<3200.0{
						newLineUpOverView.Hero1GoldDets.Gold3100to3200Counts=newLineUpOverView.Hero1GoldDets.Gold3100to3200Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3200.0&&v2.Pre10MinGold.AllGold<3300.0{
						newLineUpOverView.Hero1GoldDets.Gold3200to3300Counts=newLineUpOverView.Hero1GoldDets.Gold3200to3300Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3300.0&&v2.Pre10MinGold.AllGold<3400.0{
						newLineUpOverView.Hero1GoldDets.Gold3300to3400Counts=newLineUpOverView.Hero1GoldDets.Gold3300to3400Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3400.0&&v2.Pre10MinGold.AllGold<3500.0{
						newLineUpOverView.Hero1GoldDets.Gold3400to3500Counts=newLineUpOverView.Hero1GoldDets.Gold3400to3500Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3500.0&&v2.Pre10MinGold.AllGold<3600.0{
						newLineUpOverView.Hero1GoldDets.Gold3500to3600Counts=newLineUpOverView.Hero1GoldDets.Gold3500to3600Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3600.0&&v2.Pre10MinGold.AllGold<3700.0{
						newLineUpOverView.Hero1GoldDets.Gold3600to3700Counts=newLineUpOverView.Hero1GoldDets.Gold3600to3700Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3700.0&&v2.Pre10MinGold.AllGold<3800.0{
						newLineUpOverView.Hero1GoldDets.Gold3700to3800Counts=newLineUpOverView.Hero1GoldDets.Gold3700to3800Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3800.0&&v2.Pre10MinGold.AllGold<3900.0{
						newLineUpOverView.Hero1GoldDets.Gold3800to3900Counts=newLineUpOverView.Hero1GoldDets.Gold3800to3900Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3900.0&&v2.Pre10MinGold.AllGold<4000.0{
						newLineUpOverView.Hero1GoldDets.Gold3900to4000Counts=newLineUpOverView.Hero1GoldDets.Gold3900to4000Counts+1
					}
					if v2.Pre10MinGold.AllGold>=4000{
						newLineUpOverView.Hero1GoldDets.GoldMoreThan4000Counts=newLineUpOverView.Hero1GoldDets.GoldMoreThan4000Counts+1
					}
					if v2.HeroDeathCount==1{
						newLineUpOverView.Hero1GoldDets.Death1Counts=newLineUpOverView.Hero1GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount==2{
						newLineUpOverView.Hero1GoldDets.Death2Counts=newLineUpOverView.Hero1GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount==3{
						newLineUpOverView.Hero1GoldDets.Death3Counts=newLineUpOverView.Hero1GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount>3{
						newLineUpOverView.Hero1GoldDets.DeathMoreThan3Counts=newLineUpOverView.Hero1GoldDets.DeathMoreThan3Counts+1
					}
				}
				if v2.HeroName==newLineUpOverView.Hero2GoldDets.HeroName{
					if v2.Pre10MinGold.AllGold<3000.0{
						newLineUpOverView.Hero2GoldDets.GoldLessThan3000Counts=newLineUpOverView.Hero2GoldDets.GoldLessThan3000Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3000.0&&v2.Pre10MinGold.AllGold<3100.0{
						newLineUpOverView.Hero2GoldDets.Gold3000to3100Counts=newLineUpOverView.Hero2GoldDets.Gold3000to3100Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3100.0&&v2.Pre10MinGold.AllGold<3200.0{
						newLineUpOverView.Hero2GoldDets.Gold3100to3200Counts=newLineUpOverView.Hero2GoldDets.Gold3100to3200Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3200.0&&v2.Pre10MinGold.AllGold<3300.0{
						newLineUpOverView.Hero2GoldDets.Gold3200to3300Counts=newLineUpOverView.Hero2GoldDets.Gold3200to3300Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3300.0&&v2.Pre10MinGold.AllGold<3400.0{
						newLineUpOverView.Hero2GoldDets.Gold3300to3400Counts=newLineUpOverView.Hero2GoldDets.Gold3300to3400Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3400.0&&v2.Pre10MinGold.AllGold<3500.0{
						newLineUpOverView.Hero2GoldDets.Gold3400to3500Counts=newLineUpOverView.Hero2GoldDets.Gold3400to3500Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3500.0&&v2.Pre10MinGold.AllGold<3600.0{
						newLineUpOverView.Hero2GoldDets.Gold3500to3600Counts=newLineUpOverView.Hero2GoldDets.Gold3500to3600Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3600.0&&v2.Pre10MinGold.AllGold<3700.0{
						newLineUpOverView.Hero2GoldDets.Gold3600to3700Counts=newLineUpOverView.Hero2GoldDets.Gold3600to3700Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3700.0&&v2.Pre10MinGold.AllGold<3800.0{
						newLineUpOverView.Hero2GoldDets.Gold3700to3800Counts=newLineUpOverView.Hero2GoldDets.Gold3700to3800Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3800.0&&v2.Pre10MinGold.AllGold<3900.0{
						newLineUpOverView.Hero2GoldDets.Gold3800to3900Counts=newLineUpOverView.Hero2GoldDets.Gold3800to3900Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3900.0&&v2.Pre10MinGold.AllGold<4000.0{
						newLineUpOverView.Hero2GoldDets.Gold3900to4000Counts=newLineUpOverView.Hero2GoldDets.Gold3900to4000Counts+1
					}
					if v2.Pre10MinGold.AllGold>=4000{
						newLineUpOverView.Hero2GoldDets.GoldMoreThan4000Counts=newLineUpOverView.Hero2GoldDets.GoldMoreThan4000Counts+1
					}
					if v2.HeroDeathCount==1{
						newLineUpOverView.Hero2GoldDets.Death1Counts=newLineUpOverView.Hero2GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount==2{
						newLineUpOverView.Hero2GoldDets.Death2Counts=newLineUpOverView.Hero2GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount==3{
						newLineUpOverView.Hero2GoldDets.Death3Counts=newLineUpOverView.Hero2GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount>3{
						newLineUpOverView.Hero2GoldDets.DeathMoreThan3Counts=newLineUpOverView.Hero2GoldDets.DeathMoreThan3Counts+1
					}	
				}
				if v2.HeroName==newLineUpOverView.Hero3GoldDets.HeroName{
						if v2.Pre10MinGold.AllGold<3000.0{
							newLineUpOverView.Hero3GoldDets.GoldLessThan3000Counts=newLineUpOverView.Hero3GoldDets.GoldLessThan3000Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3000.0&&v2.Pre10MinGold.AllGold<3100.0{
							newLineUpOverView.Hero3GoldDets.Gold3000to3100Counts=newLineUpOverView.Hero3GoldDets.Gold3000to3100Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3100.0&&v2.Pre10MinGold.AllGold<3200.0{
							newLineUpOverView.Hero3GoldDets.Gold3100to3200Counts=newLineUpOverView.Hero3GoldDets.Gold3100to3200Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3200.0&&v2.Pre10MinGold.AllGold<3300.0{
							newLineUpOverView.Hero3GoldDets.Gold3200to3300Counts=newLineUpOverView.Hero3GoldDets.Gold3200to3300Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3300.0&&v2.Pre10MinGold.AllGold<3400.0{
							newLineUpOverView.Hero3GoldDets.Gold3300to3400Counts=newLineUpOverView.Hero3GoldDets.Gold3300to3400Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3400.0&&v2.Pre10MinGold.AllGold<3500.0{
							newLineUpOverView.Hero3GoldDets.Gold3400to3500Counts=newLineUpOverView.Hero3GoldDets.Gold3400to3500Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3500.0&&v2.Pre10MinGold.AllGold<3600.0{
							newLineUpOverView.Hero3GoldDets.Gold3500to3600Counts=newLineUpOverView.Hero3GoldDets.Gold3500to3600Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3600.0&&v2.Pre10MinGold.AllGold<3700.0{
							newLineUpOverView.Hero3GoldDets.Gold3600to3700Counts=newLineUpOverView.Hero3GoldDets.Gold3600to3700Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3700.0&&v2.Pre10MinGold.AllGold<3800.0{
							newLineUpOverView.Hero3GoldDets.Gold3700to3800Counts=newLineUpOverView.Hero3GoldDets.Gold3700to3800Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3800.0&&v2.Pre10MinGold.AllGold<3900.0{
							newLineUpOverView.Hero3GoldDets.Gold3800to3900Counts=newLineUpOverView.Hero3GoldDets.Gold3800to3900Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3900.0&&v2.Pre10MinGold.AllGold<4000.0{
							newLineUpOverView.Hero3GoldDets.Gold3900to4000Counts=newLineUpOverView.Hero3GoldDets.Gold3900to4000Counts+1
						}
						if v2.Pre10MinGold.AllGold>=4000{
							newLineUpOverView.Hero3GoldDets.GoldMoreThan4000Counts=newLineUpOverView.Hero3GoldDets.GoldMoreThan4000Counts+1
						}
						if v2.HeroDeathCount==1{
							newLineUpOverView.Hero3GoldDets.Death1Counts=newLineUpOverView.Hero3GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount==2{
							newLineUpOverView.Hero3GoldDets.Death2Counts=newLineUpOverView.Hero3GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount==3{
							newLineUpOverView.Hero3GoldDets.Death3Counts=newLineUpOverView.Hero3GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount>3{
							newLineUpOverView.Hero3GoldDets.DeathMoreThan3Counts=newLineUpOverView.Hero3GoldDets.DeathMoreThan3Counts+1
						}	
				}
				if v2.HeroName==newLineUpOverView.Hero4GoldDets.HeroName{
						if v2.Pre10MinGold.AllGold<3000.0{
							newLineUpOverView.Hero4GoldDets.GoldLessThan3000Counts=newLineUpOverView.Hero4GoldDets.GoldLessThan3000Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3000.0&&v2.Pre10MinGold.AllGold<3100.0{
							newLineUpOverView.Hero4GoldDets.Gold3000to3100Counts=newLineUpOverView.Hero4GoldDets.Gold3000to3100Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3100.0&&v2.Pre10MinGold.AllGold<3200.0{
							newLineUpOverView.Hero4GoldDets.Gold3100to3200Counts=newLineUpOverView.Hero4GoldDets.Gold3100to3200Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3200.0&&v2.Pre10MinGold.AllGold<3300.0{
							newLineUpOverView.Hero4GoldDets.Gold3200to3300Counts=newLineUpOverView.Hero4GoldDets.Gold3200to3300Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3300.0&&v2.Pre10MinGold.AllGold<3400.0{
							newLineUpOverView.Hero4GoldDets.Gold3300to3400Counts=newLineUpOverView.Hero4GoldDets.Gold3300to3400Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3400.0&&v2.Pre10MinGold.AllGold<3500.0{
							newLineUpOverView.Hero4GoldDets.Gold3400to3500Counts=newLineUpOverView.Hero4GoldDets.Gold3400to3500Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3500.0&&v2.Pre10MinGold.AllGold<3600.0{
							newLineUpOverView.Hero4GoldDets.Gold3500to3600Counts=newLineUpOverView.Hero4GoldDets.Gold3500to3600Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3600.0&&v2.Pre10MinGold.AllGold<3700.0{
							newLineUpOverView.Hero4GoldDets.Gold3600to3700Counts=newLineUpOverView.Hero4GoldDets.Gold3600to3700Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3700.0&&v2.Pre10MinGold.AllGold<3800.0{
							newLineUpOverView.Hero4GoldDets.Gold3700to3800Counts=newLineUpOverView.Hero4GoldDets.Gold3700to3800Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3800.0&&v2.Pre10MinGold.AllGold<3900.0{
							newLineUpOverView.Hero4GoldDets.Gold3800to3900Counts=newLineUpOverView.Hero4GoldDets.Gold3800to3900Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3900.0&&v2.Pre10MinGold.AllGold<4000.0{
							newLineUpOverView.Hero4GoldDets.Gold3900to4000Counts=newLineUpOverView.Hero4GoldDets.Gold3900to4000Counts+1
						}
						if v2.Pre10MinGold.AllGold>=4000{
							newLineUpOverView.Hero4GoldDets.GoldMoreThan4000Counts=newLineUpOverView.Hero4GoldDets.GoldMoreThan4000Counts+1
						}
						if v2.HeroDeathCount==1{
							newLineUpOverView.Hero4GoldDets.Death1Counts=newLineUpOverView.Hero4GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount==2{
							newLineUpOverView.Hero4GoldDets.Death2Counts=newLineUpOverView.Hero4GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount==3{
							newLineUpOverView.Hero4GoldDets.Death3Counts=newLineUpOverView.Hero4GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount>3{
							newLineUpOverView.Hero4GoldDets.DeathMoreThan3Counts=newLineUpOverView.Hero4GoldDets.DeathMoreThan3Counts+1
						}	
				}
				if v2.HeroName==newLineUpOverView.Hero5GoldDets.HeroName{
					if v2.Pre10MinGold.AllGold<3000.0{
						newLineUpOverView.Hero5GoldDets.GoldLessThan3000Counts=newLineUpOverView.Hero5GoldDets.GoldLessThan3000Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3000.0&&v2.Pre10MinGold.AllGold<3100.0{
						newLineUpOverView.Hero5GoldDets.Gold3000to3100Counts=newLineUpOverView.Hero5GoldDets.Gold3000to3100Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3100.0&&v2.Pre10MinGold.AllGold<3200.0{
						newLineUpOverView.Hero5GoldDets.Gold3100to3200Counts=newLineUpOverView.Hero5GoldDets.Gold3100to3200Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3200.0&&v2.Pre10MinGold.AllGold<3300.0{
						newLineUpOverView.Hero5GoldDets.Gold3200to3300Counts=newLineUpOverView.Hero5GoldDets.Gold3200to3300Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3300.0&&v2.Pre10MinGold.AllGold<3400.0{
						newLineUpOverView.Hero5GoldDets.Gold3300to3400Counts=newLineUpOverView.Hero5GoldDets.Gold3300to3400Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3400.0&&v2.Pre10MinGold.AllGold<3500.0{
						newLineUpOverView.Hero5GoldDets.Gold3400to3500Counts=newLineUpOverView.Hero5GoldDets.Gold3400to3500Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3500.0&&v2.Pre10MinGold.AllGold<3600.0{
						newLineUpOverView.Hero5GoldDets.Gold3500to3600Counts=newLineUpOverView.Hero5GoldDets.Gold3500to3600Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3600.0&&v2.Pre10MinGold.AllGold<3700.0{
						newLineUpOverView.Hero5GoldDets.Gold3600to3700Counts=newLineUpOverView.Hero5GoldDets.Gold3600to3700Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3700.0&&v2.Pre10MinGold.AllGold<3800.0{
						newLineUpOverView.Hero5GoldDets.Gold3700to3800Counts=newLineUpOverView.Hero5GoldDets.Gold3700to3800Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3800.0&&v2.Pre10MinGold.AllGold<3900.0{
						newLineUpOverView.Hero5GoldDets.Gold3800to3900Counts=newLineUpOverView.Hero5GoldDets.Gold3800to3900Counts+1
					}
					if v2.Pre10MinGold.AllGold>=3900.0&&v2.Pre10MinGold.AllGold<4000.0{
						newLineUpOverView.Hero5GoldDets.Gold3900to4000Counts=newLineUpOverView.Hero5GoldDets.Gold3900to4000Counts+1
					}
					if v2.Pre10MinGold.AllGold>=4000{
						newLineUpOverView.Hero5GoldDets.GoldMoreThan4000Counts=newLineUpOverView.Hero5GoldDets.GoldMoreThan4000Counts+1
					}
					if v2.HeroDeathCount==1{
						newLineUpOverView.Hero5GoldDets.Death1Counts=newLineUpOverView.Hero5GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount==2{
						newLineUpOverView.Hero5GoldDets.Death2Counts=newLineUpOverView.Hero5GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount==3{
						newLineUpOverView.Hero5GoldDets.Death3Counts=newLineUpOverView.Hero5GoldDets.Death1Counts+1
					}
					if v2.HeroDeathCount>3{
						newLineUpOverView.Hero5GoldDets.DeathMoreThan3Counts=newLineUpOverView.Hero5GoldDets.DeathMoreThan3Counts+1
					}	
				}
				if v2.HeroName==newLineUpOverView.Hero6GoldDets.HeroName{
						if v2.Pre10MinGold.AllGold<3000.0{
							newLineUpOverView.Hero6GoldDets.GoldLessThan3000Counts=newLineUpOverView.Hero6GoldDets.GoldLessThan3000Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3000.0&&v2.Pre10MinGold.AllGold<3100.0{
							newLineUpOverView.Hero6GoldDets.Gold3000to3100Counts=newLineUpOverView.Hero6GoldDets.Gold3000to3100Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3100.0&&v2.Pre10MinGold.AllGold<3200.0{
							newLineUpOverView.Hero6GoldDets.Gold3100to3200Counts=newLineUpOverView.Hero6GoldDets.Gold3100to3200Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3200.0&&v2.Pre10MinGold.AllGold<3300.0{
							newLineUpOverView.Hero6GoldDets.Gold3200to3300Counts=newLineUpOverView.Hero6GoldDets.Gold3200to3300Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3300.0&&v2.Pre10MinGold.AllGold<3400.0{
							newLineUpOverView.Hero6GoldDets.Gold3300to3400Counts=newLineUpOverView.Hero6GoldDets.Gold3300to3400Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3400.0&&v2.Pre10MinGold.AllGold<3500.0{
							newLineUpOverView.Hero6GoldDets.Gold3400to3500Counts=newLineUpOverView.Hero6GoldDets.Gold3400to3500Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3500.0&&v2.Pre10MinGold.AllGold<3600.0{
							newLineUpOverView.Hero6GoldDets.Gold3500to3600Counts=newLineUpOverView.Hero6GoldDets.Gold3500to3600Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3600.0&&v2.Pre10MinGold.AllGold<3700.0{
							newLineUpOverView.Hero6GoldDets.Gold3600to3700Counts=newLineUpOverView.Hero6GoldDets.Gold3600to3700Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3700.0&&v2.Pre10MinGold.AllGold<3800.0{
							newLineUpOverView.Hero6GoldDets.Gold3700to3800Counts=newLineUpOverView.Hero6GoldDets.Gold3700to3800Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3800.0&&v2.Pre10MinGold.AllGold<3900.0{
							newLineUpOverView.Hero6GoldDets.Gold3800to3900Counts=newLineUpOverView.Hero6GoldDets.Gold3800to3900Counts+1
						}
						if v2.Pre10MinGold.AllGold>=3900.0&&v2.Pre10MinGold.AllGold<4000.0{
							newLineUpOverView.Hero6GoldDets.Gold3900to4000Counts=newLineUpOverView.Hero6GoldDets.Gold3900to4000Counts+1
						}
						if v2.Pre10MinGold.AllGold>=4000{
							newLineUpOverView.Hero6GoldDets.GoldMoreThan4000Counts=newLineUpOverView.Hero6GoldDets.GoldMoreThan4000Counts+1
						}
						if v2.HeroDeathCount==1{
							newLineUpOverView.Hero6GoldDets.Death1Counts=newLineUpOverView.Hero6GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount==2{
							newLineUpOverView.Hero6GoldDets.Death2Counts=newLineUpOverView.Hero6GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount==3{
							newLineUpOverView.Hero6GoldDets.Death3Counts=newLineUpOverView.Hero6GoldDets.Death1Counts+1
						}
						if v2.HeroDeathCount>3{
							newLineUpOverView.Hero6GoldDets.DeathMoreThan3Counts=newLineUpOverView.Hero6GoldDets.DeathMoreThan3Counts+1
						}	
				}
				
			}
		}		
	}
	newLineUpOverView.Counts=count
	//fmt.Printf("newLineUpOverView:%v\n", newLineUpOverView)
	type HeroGoldDets struct{
		Descrition string `json:"description"`
		Count int `json:"count"`	
	}
	var hero1GoldDetsmap=map[string]int{
		"经济少于3000":newLineUpOverView.Hero1GoldDets.GoldLessThan3000Counts,
		"经济在3000到3100":newLineUpOverView.Hero1GoldDets.Gold3000to3100Counts,
		"经济在3100到3200":newLineUpOverView.Hero1GoldDets.Gold3100to3200Counts,
		"经济在3200到3300":newLineUpOverView.Hero1GoldDets.Gold3200to3300Counts,
		"经济在3300到3400":newLineUpOverView.Hero1GoldDets.Gold3300to3400Counts,
		"经济在3400到3500":newLineUpOverView.Hero1GoldDets.Gold3400to3500Counts,
		"经济在3500到3600":newLineUpOverView.Hero1GoldDets.Gold3500to3600Counts,
		"经济在3600到3700":newLineUpOverView.Hero1GoldDets.Gold3600to3700Counts,
		"经济在3700到3800":newLineUpOverView.Hero1GoldDets.Gold3700to3800Counts,
		"经济在3800到3900":newLineUpOverView.Hero1GoldDets.Gold3800to3900Counts,
		"经济在3900到4000":newLineUpOverView.Hero1GoldDets.Gold3900to4000Counts,
		"经济超过4000":newLineUpOverView.Hero1GoldDets.GoldMoreThan4000Counts,
	}
	var hero2GoldDetsmap=map[string]int{
		"经济少于3000":newLineUpOverView.Hero2GoldDets.GoldLessThan3000Counts,
		"经济在3000到3100":newLineUpOverView.Hero2GoldDets.Gold3000to3100Counts,
		"经济在3100到3200":newLineUpOverView.Hero2GoldDets.Gold3100to3200Counts,
		"经济在3200到3300":newLineUpOverView.Hero2GoldDets.Gold3200to3300Counts,
		"经济在3300到3400":newLineUpOverView.Hero2GoldDets.Gold3300to3400Counts,
		"经济在3400到3500":newLineUpOverView.Hero2GoldDets.Gold3400to3500Counts,
		"经济在3500到3600":newLineUpOverView.Hero2GoldDets.Gold3500to3600Counts,
		"经济在3600到3700":newLineUpOverView.Hero2GoldDets.Gold3600to3700Counts,
		"经济在3700到3800":newLineUpOverView.Hero2GoldDets.Gold3700to3800Counts,
		"经济在3800到3900":newLineUpOverView.Hero2GoldDets.Gold3800to3900Counts,
		"经济在3900到4000":newLineUpOverView.Hero2GoldDets.Gold3900to4000Counts,
		"经济超过4000":newLineUpOverView.Hero2GoldDets.GoldMoreThan4000Counts,
	}
	var hero3GoldDetsmap map[string]int
	if  newLineUpOverView.Hero3GoldDets.HeroName!=""{
		hero3GoldDetsmap=map[string]int{
			"经济少于3000":newLineUpOverView.Hero3GoldDets.GoldLessThan3000Counts,
			"经济在3000到3100":newLineUpOverView.Hero3GoldDets.Gold3000to3100Counts,
			"经济在3100到3200":newLineUpOverView.Hero3GoldDets.Gold3100to3200Counts,
			"经济在3200到3300":newLineUpOverView.Hero3GoldDets.Gold3200to3300Counts,
			"经济在3300到3400":newLineUpOverView.Hero3GoldDets.Gold3300to3400Counts,
			"经济在3400到3500":newLineUpOverView.Hero3GoldDets.Gold3400to3500Counts,
			"经济在3500到3600":newLineUpOverView.Hero3GoldDets.Gold3500to3600Counts,
			"经济在3600到3700":newLineUpOverView.Hero3GoldDets.Gold3600to3700Counts,
			"经济在3700到3800":newLineUpOverView.Hero3GoldDets.Gold3700to3800Counts,
			"经济在3800到3900":newLineUpOverView.Hero3GoldDets.Gold3800to3900Counts,
			"经济在3900到4000":newLineUpOverView.Hero3GoldDets.Gold3900to4000Counts,
			"经济超过4000":newLineUpOverView.Hero3GoldDets.GoldMoreThan4000Counts,
		}
	} 
	var hero4GoldDetsmap map[string]int
	if  newLineUpOverView.Hero4GoldDets.HeroName!=""{
		hero4GoldDetsmap=map[string]int{
			"经济少于3000":newLineUpOverView.Hero4GoldDets.GoldLessThan3000Counts,
			"经济在3000到3100":newLineUpOverView.Hero4GoldDets.Gold3000to3100Counts,
			"经济在3100到3200":newLineUpOverView.Hero4GoldDets.Gold3100to3200Counts,
			"经济在3200到3300":newLineUpOverView.Hero4GoldDets.Gold3200to3300Counts,
			"经济在3300到3400":newLineUpOverView.Hero4GoldDets.Gold3300to3400Counts,
			"经济在3400到3500":newLineUpOverView.Hero4GoldDets.Gold3400to3500Counts,
			"经济在3500到3600":newLineUpOverView.Hero4GoldDets.Gold3500to3600Counts,
			"经济在3600到3700":newLineUpOverView.Hero4GoldDets.Gold3600to3700Counts,
			"经济在3700到3800":newLineUpOverView.Hero4GoldDets.Gold3700to3800Counts,
			"经济在3800到3900":newLineUpOverView.Hero4GoldDets.Gold3800to3900Counts,
			"经济在3900到4000":newLineUpOverView.Hero4GoldDets.Gold3900to4000Counts,
			"经济超过4000":newLineUpOverView.Hero4GoldDets.GoldMoreThan4000Counts,
		}
	}
	var hero5GoldDetsmap map[string]int
	if  newLineUpOverView.Hero5GoldDets.HeroName!=""{
		hero5GoldDetsmap=map[string]int{
			"经济少于3000":newLineUpOverView.Hero5GoldDets.GoldLessThan3000Counts,
			"经济在3000到3100":newLineUpOverView.Hero5GoldDets.Gold3000to3100Counts,
			"经济在3100到3200":newLineUpOverView.Hero5GoldDets.Gold3100to3200Counts,
			"经济在3200到3300":newLineUpOverView.Hero5GoldDets.Gold3200to3300Counts,
			"经济在3300到3400":newLineUpOverView.Hero5GoldDets.Gold3300to3400Counts,
			"经济在3400到3500":newLineUpOverView.Hero5GoldDets.Gold3400to3500Counts,
			"经济在3500到3600":newLineUpOverView.Hero5GoldDets.Gold3500to3600Counts,
			"经济在3600到3700":newLineUpOverView.Hero5GoldDets.Gold3600to3700Counts,
			"经济在3700到3800":newLineUpOverView.Hero5GoldDets.Gold3700to3800Counts,
			"经济在3800到3900":newLineUpOverView.Hero5GoldDets.Gold3800to3900Counts,
			"经济在3900到4000":newLineUpOverView.Hero5GoldDets.Gold3900to4000Counts,
			"经济超过4000":newLineUpOverView.Hero5GoldDets.GoldMoreThan4000Counts,
		}
	}
	var hero6GoldDetsmap map[string]int
	if  newLineUpOverView.Hero6GoldDets.HeroName!=""{
		hero6GoldDetsmap=map[string]int{
			"经济少于3000":newLineUpOverView.Hero6GoldDets.GoldLessThan3000Counts,
			"经济在3000到3100":newLineUpOverView.Hero6GoldDets.Gold3000to3100Counts,
			"经济在3100到3200":newLineUpOverView.Hero6GoldDets.Gold3100to3200Counts,
			"经济在3200到3300":newLineUpOverView.Hero6GoldDets.Gold3200to3300Counts,
			"经济在3300到3400":newLineUpOverView.Hero6GoldDets.Gold3300to3400Counts,
			"经济在3400到3500":newLineUpOverView.Hero6GoldDets.Gold3400to3500Counts,
			"经济在3500到3600":newLineUpOverView.Hero6GoldDets.Gold3500to3600Counts,
			"经济在3600到3700":newLineUpOverView.Hero6GoldDets.Gold3600to3700Counts,
			"经济在3700到3800":newLineUpOverView.Hero6GoldDets.Gold3700to3800Counts,
			"经济在3800到3900":newLineUpOverView.Hero6GoldDets.Gold3800to3900Counts,
			"经济在3900到4000":newLineUpOverView.Hero6GoldDets.Gold3900to4000Counts,
			"经济超过4000":newLineUpOverView.Hero6GoldDets.GoldMoreThan4000Counts,
		}
	}
 
	var Hero1GoldArr []HeroGoldDets
	var Hero2GoldArr []HeroGoldDets
	var Hero3GoldArr []HeroGoldDets
	var Hero4GoldArr []HeroGoldDets
	var Hero5GoldArr []HeroGoldDets
	var Hero6GoldArr []HeroGoldDets
	for k, v := range hero1GoldDetsmap {
		Hero1GoldArr = append(Hero1GoldArr, HeroGoldDets{k, v})
	}
	sort.Slice(Hero1GoldArr, func(i, j int) bool {
		return Hero1GoldArr[i].Count > Hero1GoldArr[j].Count  // 降序
	})
	newLineUpOverView.Hero1GoldDets.GoldMaximumProportion=Hero1GoldArr[0].Descrition
	if Hero1GoldArr[1].Count==1{
		newLineUpOverView.Hero1GoldDets.GoldSecondMaxProportion=Hero1GoldArr[1].Descrition
	}
	

	for k, v := range hero2GoldDetsmap {
		Hero2GoldArr = append(Hero2GoldArr, HeroGoldDets{k, v})
	}
	sort.Slice(Hero2GoldArr, func(i, j int) bool {
		return Hero2GoldArr[i].Count > Hero2GoldArr[j].Count  // 降序
	})
	newLineUpOverView.Hero2GoldDets.GoldMaximumProportion=Hero2GoldArr[0].Descrition
	newLineUpOverView.Hero2GoldDets.GoldSecondMaxProportion=Hero2GoldArr[1].Descrition
	
	if hero3GoldDetsmap!=nil{
		for k, v := range hero3GoldDetsmap {
			Hero3GoldArr = append(Hero3GoldArr, HeroGoldDets{k, v})
		}
		sort.Slice(Hero3GoldArr, func(i, j int) bool {
			return Hero3GoldArr[i].Count > Hero3GoldArr[j].Count  // 降序
		})
		newLineUpOverView.Hero3GoldDets.GoldMaximumProportion=Hero3GoldArr[0].Descrition
		newLineUpOverView.Hero3GoldDets.GoldSecondMaxProportion=Hero3GoldArr[1].Descrition
	}
	if hero4GoldDetsmap!=nil{
		for k, v := range hero4GoldDetsmap {
			Hero4GoldArr = append(Hero4GoldArr, HeroGoldDets{k, v})
		}
		sort.Slice(Hero4GoldArr, func(i, j int) bool {
			return Hero4GoldArr[i].Count > Hero4GoldArr[j].Count  // 降序
		})
		newLineUpOverView.Hero4GoldDets.GoldMaximumProportion=Hero4GoldArr[0].Descrition
		newLineUpOverView.Hero4GoldDets.GoldSecondMaxProportion=Hero4GoldArr[1].Descrition
	}
	if hero5GoldDetsmap!=nil{
		for k, v := range hero5GoldDetsmap {
			Hero5GoldArr = append(Hero5GoldArr, HeroGoldDets{k, v})
		}
		sort.Slice(Hero5GoldArr, func(i, j int) bool {
			return Hero5GoldArr[i].Count > Hero5GoldArr[j].Count  // 降序
		})
		newLineUpOverView.Hero5GoldDets.GoldMaximumProportion=Hero5GoldArr[0].Descrition
		newLineUpOverView.Hero5GoldDets.GoldSecondMaxProportion=Hero5GoldArr[1].Descrition
	}
	if hero6GoldDetsmap!=nil{
		for k, v := range hero6GoldDetsmap {
			Hero6GoldArr = append(Hero6GoldArr, HeroGoldDets{k, v})
		}
		sort.Slice(Hero6GoldArr, func(i, j int) bool {
			return Hero6GoldArr[i].Count > Hero6GoldArr[j].Count  // 降序
		})
		newLineUpOverView.Hero6GoldDets.GoldMaximumProportion=Hero6GoldArr[0].Descrition
		newLineUpOverView.Hero6GoldDets.GoldSecondMaxProportion=Hero6GoldArr[1].Descrition
	}
	return newLineUpOverView		
}
/*获取战队对线情况*/
func GetTeamMainPositionLineUp(version string,team string) ([]string,[]string){
	session := mongodb.CloneSession()
	defer session.Close()  
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	Colplayer:=db.C("player_info")
	ColCSideLineUpData:=db.C("SideLineUpData")
	var hardresults []obj.SideLineUp
	var softresults []obj.SideLineUp
	var hardresp []string
	var softresp []string
	var hardplayer obj.Player
	var softplayer obj.Player
	fmt.Printf("team:%v\n", team)
	playererr:=Colplayer.Find(bson.M{"team_name_tag":team,"position":"三号位","player_state":"正式选手"}).One(&hardplayer)
	if playererr!=nil{
		fmt.Printf("三号位搜索错误：%v\n",playererr)
	}
	hardplayerid:=hardplayer.PlayerDota2NumId

	playererr1:=Colplayer.Find(bson.M{"team_name_tag":team,"position":"一号位","player_state":"正式选手"}).One(&softplayer)
	if playererr1!=nil{
		fmt.Printf("一号位搜索错误：%v\n",playererr1)
	}
	softplayerid:=softplayer.PlayerDota2NumId

	hardsideerr:=ColCSideLineUpData.Find(bson.M{"$or":[]bson.M{bson.M{"version":version,"soft_lane_team":team,"soft_main_position_player_id":hardplayerid},bson.M{"version":version,"hard_lane_team":team,"hard_main_position_player_id":hardplayerid}}}).All(&hardresults)
	if hardsideerr!=nil{
		fmt.Printf("三号位边路数据搜索错误：%v\n",hardsideerr)
	}
	for _,v:=range hardresults{
		if v.SoftLaneTeam==team{
			hardresp=append(hardresp,v.LineUpHeroes)
		}else{
			lanearr:=strings.Split(v.LineUpHeroes, "vs")
			lanestr:=lanearr[1]+"vs"+lanearr[0]
			hardresp=append(hardresp,lanestr)
		}	
	}

	softsideerr:=ColCSideLineUpData.Find(bson.M{"$or":[]bson.M{bson.M{"version":version,"soft_lane_team":team,"soft_main_position_player_id":softplayerid},bson.M{"version":version,"hard_lane_team":team,"hard_main_position_player_id":softplayerid}}}).All(&softresults)
	if softsideerr!=nil{
		fmt.Printf("一号位边路数据搜索错误：%v\n",softsideerr)
	}
	for _,v:=range softresults{
		if v.SoftLaneTeam==team{
			softresp=append(softresp,v.LineUpHeroes)
		}else{
			lanearr:=strings.Split(v.LineUpHeroes, "vs")
			lanestr:=lanearr[1]+"vs"+lanearr[0]
			softresp=append(softresp,lanestr)
		}	
	}
	return hardresp,softresp
}

func GetSideLineUpDets(version string,lineUp string,team string) []obj.SideLineUp{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColSideLineUpData:=db.C("SideLineUpData")
	var respSideLineUpData []obj.SideLineUp
	var teamSideLineUpData []obj.SideLineUp
	var otherTeamSideLineUpData []obj.SideLineUp
	//var otherTeamSideLineUpData1 []obj.SideLineUp
	if team!=""{
		err:=ColSideLineUpData.Find(bson.M{"$or":[]bson.M{bson.M{"version":version,"line_up_heroes":lineUp,"soft_lane_team":team},bson.M{"version":version,"line_up_heroes":lineUp,"hard_lane_team":team}}}).All(&teamSideLineUpData)	
		if err!=nil{
			fmt.Printf("err:%v\n", err)
		}
		err1:=ColSideLineUpData.Find(bson.M{"version":version,"line_up_heroes":lineUp,"soft_lane_team":bson.M{"$ne":team},"hard_lane_team":bson.M{"$ne":team}}).All(&otherTeamSideLineUpData)
		if err1!=nil{
			fmt.Printf("err1:%v\n", err1)
		}
		respSideLineUpData=append(teamSideLineUpData,otherTeamSideLineUpData...)
	}else{
		err:=ColSideLineUpData.Find(bson.M{"version":version,"line_up_heroes":lineUp}).All(&respSideLineUpData)
		if err!=nil{
			fmt.Printf("err:%v\n", err)
		}
	}
	return respSideLineUpData
}




