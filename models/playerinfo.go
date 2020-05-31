package models

import(
	"fmt"
	"strings"
	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson"
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
)

func GetPlayerInfo (playerNumId uint32) *obj.Player {
	//fmt.Printf("playerNumId%v\n",playerNumId)
	var playerInfo obj.Player
	//var playerInfo obj.PlayersHeroInfo
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	dbplayer:=db.C("player_info")
	err:=dbplayer.Find(bson.M{"player_dota2_register_num_id":playerNumId}).One(&playerInfo)
	if err!=nil{
		fmt.Printf("查找选手数据出现错误2:%v\n", err)
	}
	//fmt.Printf("PlayerInfo:%v\n",playerInfo)
	return &playerInfo
}
func AddBasePlayerInfo(playerNumId uint32,playerStringName string,secondNumId uint32,thirdNumId uint32,forthNumId uint32,position string,playerState string){
	var playerInfo obj.Player
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	dbplayer:=db.C("player_info")
	err:=dbplayer.Find(bson.M{"player_dota2_register_num_id":playerNumId}).One(&playerInfo)
	if err!=nil{
		fmt.Printf("查找选手数据出现错误1:%v\n", err)
	}
	if playerInfo.PlayerDota2NumId==0{
		playerInfo.MatchPlayerId=playerStringName
		playerInfo.PlayerDota2NumId=playerNumId
		playerInfo.SecondNumId=secondNumId
		playerInfo.ThirdNumId=playerInfo.ThirdNumId
		playerInfo.FourthNumId=forthNumId
		playerInfo.Position=position
		playerInfo.PlayerState=playerState
	}
}

func AlterPlayerInfo(teamname string,teamtagname string,playerNumId uint32,playerStringName string,secondNumId uint32,thirdNumId uint32,forthNumId uint32,position string,playerState string,ismovedata string){
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    playerInfo:=obj.Player{}
	db := session.DB("dota2_big_data")
	dbplayer:=db.C("player_info")
	count,playerdberr:=dbplayer.Find(bson.M{"player_dota2_register_num_id":playerNumId}).Count()
	if playerdberr!=nil{
		fmt.Printf("查找选手数据出现错误：%v\n",playerdberr)
	}
	if count==0{
		playerInfo.MatchPlayerId=playerStringName
		playerInfo.PlayerDota2NumId=playerNumId
		playerInfo.TeamName=teamname
		playerInfo.TeamNameTag=teamtagname
		playerInfo.Position=position
		playerInfo.PlayerState=playerState
		inserterr :=dbplayer.Insert(&playerInfo)
		if inserterr!=nil{
			fmt.Printf("选手数据插入出错:%v\n",inserterr)
		}
	}else{
		err:=dbplayer.Find(bson.M{"player_dota2_register_num_id":playerNumId}).One(&playerInfo)
		if err!=nil{
			fmt.Printf("find player err:%v\n", err)
		}
		dbplayer.Update(bson.M{"player_dota2_register_num_id":playerNumId},bson.M{"$set":bson.M{"team_name":teamname,"team_name_tag":teamtagname,"player_register_string_id":playerStringName,"second_num_id": secondNumId,"third_num_id":thirdNumId,"forth_num_id":forthNumId,"position":position,"player_state":playerState}})	
		//fmt.Printf("playerInfo:%v\n",playerInfo)
		olddata:=playerInfo.MatchPlayerHeroes
		if ismovedata=="是"{ 
			dbplayer.Update(bson.M{"player_dota2_register_num_id":playerNumId},bson.M{"$set":bson.M{"old_club_player_heroes":olddata}})
		}
	}
	
}
func GetLineUpAVG(version string,midhero1 string,midhero2 string) (uint32,uint32){
		session := mongodb.CloneSession()//调用这个获得session
		defer session.Close()  //一定要记得释放
	    session.SetMode(mgo.Monotonic, true)
		db := session.DB("dota2_big_data")
		ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
		var results []obj.CMsgMatchDetails
		var midlane=midhero1+"vs"+midhero2
		var midlane2=midhero2+"vs"+midhero1
		var hero1alllasthit uint32
		var hero2alllasthit uint32
		var hero1avglasthit uint32
		var hero2avglasthit uint32
		var hero1lasthit []uint32
		var hero2lasthit []uint32
		ColCMsgMatchDetails.Find(bson.M{"$or":[]bson.M{bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"mid_lane":midlane}}},bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"mid_lane":midlane2}}}}}).All(&results)
		for _,value:=range results{
			for _,value2:=range value.PlayersHeroesDets {
				if (value2.HeroName==midhero1){
					hero1lasthit=append(hero1lasthit,value2.LastHit.Pre10Count)
				}
				if (value2.HeroName==midhero2){
					hero2lasthit=append(hero2lasthit,value2.LastHit.Pre10Count)
				}
			}
		}
		for key2,value2:=range hero1lasthit{
			if key2+1!=len(hero1lasthit){
				hero1alllasthit=hero1alllasthit+value2
			}else{
				hero1alllasthit=hero1alllasthit+value2
				hero1avglasthit=hero1alllasthit/uint32(len(hero1lasthit))
			}
		}
		for key3,value3:=range hero2lasthit{
			if key3+1!=len(hero1lasthit){
				hero2alllasthit=hero2alllasthit+value3
			}else{
				hero2alllasthit=hero2alllasthit+value3
				hero2avglasthit=hero2alllasthit/uint32(len(hero2lasthit))
			}
		}
		return hero1avglasthit,hero2avglasthit
}

func GetMidUseHero(version string,playerid uint32) []string{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	var results []obj.CMsgMatchDetails
	var resp []string
	ColCMsgMatchDetails.Find(bson.M{"version":version,"is_mid_solo":1,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"account_id":playerid,"init_lane":"中路"}}}).All(&results)
	for _,v:=range results{ 		
		for _,v1:=range v.PlayersHeroesDets{
			if v1.AccountId==playerid{
				if v1.GameTeam==2{
					if len(resp)==0{
						resp=append(resp,v.MidLane)
					}else{
						resploop:
						for respk,respv:=range resp{
							if respv==v.MidLane{
								break resploop
							}else{
								resplen:=len(resp)
								if resplen-1==respk&&respv!=v.MidLane{
									resp=append(resp,v.MidLane)
									break resploop
								}
							}
						}
					}
				}else{
					midlanearr:=strings.Split(v.MidLane, "vs")
					midlane:=midlanearr[1]+"vs"+midlanearr[0]
					if len(resp)==0{
						resp=append(resp,midlane)
					}else{
						resploop2:
						for respk,respv:=range resp{
							if respv==midlane{
								break resploop2
							}else{
								resplen:=len(resp)
								if resplen-1==respk&&respv!=midlane{
									resp=append(resp,midlane)
									break resploop2
								}
							}
						}
					}
				}
			}
		}
	}
	return resp
}

//中单英雄使用
func GetMidUseHeroDetails(version string,playerid uint32) map[string][]*MidLaneHeroData{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	var results []obj.CMsgMatchDetails
	resp:=make(map[string][]*MidLaneHeroData)
	ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"account_id":playerid,"init_lane":"中路"}}}).All(&results)
	for _,v:=range results{
		playerUseHeroes:=&MidLaneHeroData{}		
		playerUseHeroes.MatchId=v.MatchId
		midlane:=[]*MidLaneHeroData{}
		Pre10MinKDAInfo:=v.KDADets
		var hero1,hero2 string
		var midstate string
		var midcount int=0
		Playersloop:
		for _,playerv:=range v.PlayersHeroesDets{
			if playerv.InitLane=="中路"{
				midcount=midcount+1
				if midcount==3{
					midstate="双中"
					break Playersloop
				}
				if midcount==1||(midcount==2&&playerv.GameTeam==3){
					tenMinData:=&TenMinData{}
					if playerv.AccountId==playerid{
						hero1=playerv.HeroName
					}else{
						hero2=playerv.HeroName
					}
					tenMinData.GameTeam=playerv.GameTeam
					tenMinData.HeroIcon=playerv.HeroIcon
					tenMinData.HeroName=playerv.HeroName
					tenMinData.PlayerName=playerv.Player
					tenMinData.PlayerId=playerv.AccountId
					tenMinData.Pre5MinGold=playerv.Gold.Pre5MinGold
					tenMinData.Pre5MinLastHit=playerv.LastHit.Pre5Count
					tenMinData.Pre10MinGold=playerv.Gold.Pre10MinGold
					tenMinData.Pre10MinLastHit=playerv.LastHit.Pre10Count
					tenMinData.ConsumableItem=playerv.Consumable

					Loopkda:
						for kdakey,kdavlaue:=range Pre10MinKDAInfo{
							if kdavlaue.Time>600{
								tenMinData.Pre10MinKDAInfo=Pre10MinKDAInfo[0:kdakey]
								break Loopkda
							}
						}
					lvluptimeslice:=playerv.LevelUpTimes
						Looplvl:
						for key,value:=range lvluptimeslice{
							if value>600{
								lvlk:=key-1
								tenMinData.Pre10MinLevelUpTime=playerv.LevelUpTimeStr[0:lvlk]
								tenMinData.AbilityUpgrades=playerv.AbilityData[0:key]
								break Looplvl
							}
						}
					itemslice:=playerv.Item
						Loopitem:
						for itemkey,itemvalue:=range itemslice{
							if itemvalue.PurchaseTimeNum>600{
								tenMinData.Pre10minItem=itemslice[0:itemkey]
								break Loopitem
							}
						}
					if playerv.AccountId==playerid{
						playerUseHeroes.MidHero1=tenMinData
					}else{
						playerUseHeroes.MidHero2=tenMinData
					}
				}else{
					midstate="双中"
					break Playersloop
				}
			}			
		}
		if midstate!="双中"{
			midlane=append(midlane,playerUseHeroes)
			var against=hero1+"vs"+hero2
			if len(resp)==0{
				resp[against]=midlane
			}else{
				var maplen=len(resp)
				resploop:
				for respk:=range resp{
					maplen=maplen-1
					if respk==against{
						resp[respk]=append(resp[respk],playerUseHeroes)
						break resploop
					}
					if respk!=against&&maplen==0{
						resp[against]=midlane
						break resploop
					}
				}
			}
		}	
	}
	return resp
}

func GetMidLineUpOverView(version string,playerid uint32){
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	ColCMsgMatchDetails:=db.C("CMsgMatchDetails")
	var results []obj.CMsgMatchDetails
	//resp:=make(map[string][]*MidLaneHeroData)
	ColCMsgMatchDetails.Find(bson.M{"version":version,"players_heroes_dets":bson.M{"$elemMatch":bson.M{"account_id":playerid,"init_lane":"中路"}}}).All(&results)
}
