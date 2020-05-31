package models
import(
	"fmt"
	"log"
	"os"
	"time"
	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
	"replayanaly/models/mongodb"
	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson"
	"replayanaly/models/obj"
	"math"
	"strconv"
	"strings"
	"sort"
	"bytes"	
)
func Parse (version string,demurl string,fname string) string{
	var resultstr string
	var playerisexist string


	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	c := db.C("CDotaGameInfo")
	c2 := db.C("CMsgDOTAMatch")
	c3 := db.C("CMsgMatchDetails")
	c4:=db.C("MidLineUp")
	c5:=db.C("SideLineUpData")
	playerdb:=db.C("player_info")//all_player_info
	abilitydb:=db.C("abilities")
	allability:=&obj.AllNpcAbility{}

	findAbilityErr:=abilitydb.Find(bson.M{"Version":version}).One(&allability)
	if findAbilityErr!=nil{
		fmt.Printf("技能查询错误:%v\n", findAbilityErr)
	}
	fnamearr:=strings.Split(fname,".")
	matchid:=fnamearr[0]
	intmatchid,transerr:=strconv.Atoi(matchid)
	if transerr!=nil{
		fmt.Printf("id字符串转int出现错误:%v\n", transerr)
	}
	//c.Find(bson.M{"name": "Tom"}).Count()
	count,_:=c.Find(bson.M{"match_id":intmatchid}).Count()
	//fmt.Printf("matchid:%v count：%v\n",intmatchid,count)
	if count==0{
		f, err := os.Open(demurl)
		if err != nil {
			//log.Fatalf("unable to open file: %s", err)
			fmt.Printf("unable to open file: %s", err)
		}
		defer f.Close()
		p, parseerr := manta.NewStreamParser(f)
		if parseerr != nil {
			log.Fatalf("unable to create parser: %s", parseerr)
		}
		var NewCDotaGameInfo obj.CDotaGameInfo
		var NewCMsgDOTAMatch obj.CMsgDOTAMatch
		var NewCMsgMatchDetails obj.CMsgMatchDetails
		var TopSideLineUp obj.SideLineUp
		var BottomSideLineUp obj.SideLineUp
		var TopSideLineUpDets= &obj.SideLineUpDets{}
		var BottomSideLineUpDets =&obj.SideLineUpDets{}

		var player1,player2,player3,player4,player5,player6,player7,player8,player9,player10 obj.PlayersHeroInfo
		var playersarr =[]*obj.PlayersHeroInfo{&player1,&player2,&player3,&player4,&player5,&player6,&player7,&player8,&player9,&player10}
		//var KDADetails=[]*obj.KDADetails{}
		var lanetimei1=1
		var lanetimei2=1
		var lanetimei3=1
		var lanetimei4=1
		var lanetimei5=1
		var lanetimei6=1
		var lanetimei7=1
		var lanetimei8=1
		var lanetimei9=1
		var lanetimei10=1
		
		var radtopherocount int
		var radmidherocount int
		var radbottomherocount int

		var diretopherocount int
		var diremidherocount int
		var direbottomherocount int

		var heroslice []int
		var yslane string=""   //优势路
		var midlane string=""
		var lslane string=""
		var ysdirelane string=""
		var lsdirelane string=""
		var diremidlane string=""

		var starttime float32
		var msgtime int = 0
		var prev5 float64
		var prev10 float64
		var prev15 float64
		var prev20 float64
		var prev25 float64
		var prev30 float64
		//var itemarr []*obj.ItemPurchase
		//描述英雄位置信息
		var hero1initlocation string
		var hero2initlocation string
		var hero3initlocation string
		var hero4initlocation string
		var hero5initlocation string
		var hero6initlocation string
		var hero7initlocation string
		var hero8initlocation string
		var hero9initlocation string
		var hero10initlocation string		
/*		var TalentGold = map[uint32]float64{
			6446: 420.0,
			6318: 300.0,
			6301: 240.0,
			8005:210.0,
			6026: 180.0,
			5957: 150.0,
			6008: 120.0,
			6007: 90.0,
			5955:30.0,
			5956:60.0,
		}*/
		/*var goldtalenttime int
		var talentgold float64
		var talentgold1 float64*/
		var allgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var runegold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var deathgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var buybackgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var itembackgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var lasthitgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var creepsgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var combatgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var killtowergold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var roshangold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var pregold float64

		var killcouriergold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0} //击杀信使获得的金钱*/
		var denies = []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var timei = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var gtimei = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var towertimei = []int64{0, 0}
		var radiantkilltower uint32
		var direkilltower uint32
		var winnerid int32
		var midstatus1 obj.MidAgainst
		//var topsideLineUp obj.SideLineUp
		//var bottomsideLineUp obj.SideLineUp
		//var	radiantmid,diremid string
		p.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error {
			NewCDotaGameInfo.Version=version
			NewCDotaGameInfo.MatchId=m.GetGameInfo().GetDota().GetMatchId()
			fmt.Printf("matchid:%v\n", NewCDotaGameInfo.MatchId)
			NewCDotaGameInfo.GameMode=m.GetGameInfo().GetDota().GetGameMode()
			NewCDotaGameInfo.Leagueid=m.GetGameInfo().GetDota().GetLeagueid()
			picksbans:=m.GetGameInfo().GetDota().GetPicksBans()
			//NewCDotaGameInfo.PicksBans=m.GetGameInfo().GetDota().GetPicksBans()
			NewCDotaGameInfo.RadiantTeamTag=m.GetGameInfo().GetDota().GetRadiantTeamTag()
			NewCDotaGameInfo.DireTeamTag=m.GetGameInfo().GetDota().GetDireTeamTag()		
			NewCDotaGameInfo.EndTime=m.GetGameInfo().GetDota().GetEndTime()
			winnerid = m.GetGameInfo().GetDota().GetGameWinner()
			for _,bpvalue:=range picksbans{
				bpinfo:=&obj.BPInfo{}
				bpinfo.HeroId=bpvalue.GetHeroId()
				bpinfo.IsPick=bpvalue.GetIsPick()
				bpinfo.Team=bpvalue.GetTeam()
				NewCDotaGameInfo.PicksBans=append(NewCDotaGameInfo.PicksBans,bpinfo)
			} 
			return nil
		})
		
		p.Callbacks.OnCMsgDOTAMatch(func(m *dota.CMsgDOTAMatch) error {
			NewCMsgDOTAMatch.Duration = m.GetDuration()
			NewCMsgDOTAMatch.DurationStr=obj.Durformat(int64(m.GetDuration()))
			tm := time.Unix(int64(m.GetStartTime()), 0)
	   		timestr:=fmt.Sprintf("%s",tm.Format("2006-01-02 15:04:05"))
			NewCMsgDOTAMatch.StartTime = timestr
			players:=m.GetPlayers()
			for _,value:=range players{
				p:=&obj.DOTAMatch_Player{}
				p.Player=value
				NewCMsgDOTAMatch.Players=append(NewCMsgDOTAMatch.Players,p)
			}
			for _,value2:=range NewCMsgDOTAMatch.Players{
				item0:=value2.Player.GetItem_0()
				itemicon0:=ItemIDIcon[item0]
				item1:=value2.Player.GetItem_1()
				itemicon1:=ItemIDIcon[item1]
				item2:=value2.Player.GetItem_2()
				itemicon2:=ItemIDIcon[item2]
				item3:=value2.Player.GetItem_3()
				itemicon3:=ItemIDIcon[item3]
				item4:=value2.Player.GetItem_4()
				itemicon4:=ItemIDIcon[item4]
				item5:=value2.Player.GetItem_5()
				itemicon5:=ItemIDIcon[item5]
				item6:=value2.Player.GetItem_6()
				itemicon6:=ItemIDIcon[item6]
				item7:=value2.Player.GetItem_7()
				itemicon7:=ItemIDIcon[item7]
				item8:=value2.Player.GetItem_8()
				itemicon8:=ItemIDIcon[item8]
				item9:=value2.Player.GetItem_9()
				itemicon9:=ItemIDIcon[item9]
				/*itemicon1,itemicon2,itemicon3,itemicon4,itemicon5,itemicon6,itemicon7,itemicon8,itemicon9*/
				value2.Item=append(value2.Item,itemicon0)
				value2.Item=append(value2.Item,itemicon1)
				value2.Item=append(value2.Item,itemicon2)
				value2.Item=append(value2.Item,itemicon3)
				value2.Item=append(value2.Item,itemicon4)
				value2.Item=append(value2.Item,itemicon5)
				value2.Item=append(value2.Item,itemicon6)
				value2.Item=append(value2.Item,itemicon7)
				value2.Item=append(value2.Item,itemicon8)
				value2.Item=append(value2.Item,itemicon9)
				buffsarr:=value2.Player.GetPermanentBuffs()
				for _,buffvalue:=range buffsarr{
					buffs:=&obj.PlayerPermanentBuff{}
					id:=buffvalue.GetPermanentBuff()
					buffs.BuffId=id
					buffs.PermanentBuff=ItemIDIcon[id]
					//fmt.Printf("buffsid:%v\n", id)
					buffs.Count=buffvalue.GetStackCount()
					value2.Buffs=append(value2.Buffs,buffs)
				}
			}
			//NewCMsgDOTAMatch.Players = m.GetPlayers()
			NewCMsgDOTAMatch.RadiantTeamScore=m.GetRadiantTeamScore()
			NewCMsgDOTAMatch.DireTeamScore=m.GetDireTeamScore()
			/*for _,value:=range NewCMsgDOTAMatch.Players{
				HeroIdArr=append(HeroIdArr,value.GetHeroId())//在切片没有初始化时不能用HeroIdArr[key]=1形式赋值
			}*/
			NewCMsgDOTAMatch.TowerStatus = m.GetTowerStatus()
			NewCMsgDOTAMatch.BarracksStatus = m.GetBarracksStatus()
			NewCMsgDOTAMatch.FirstBloodTime = m.GetFirstBloodTime()
			NewCMsgDOTAMatch.AverageSkill = m.GetAverageSkill()
			NewCMsgDOTAMatch.RadiantTeamId = m.GetRadiantTeamId()
			NewCMsgDOTAMatch.DireTeamId = m.GetDireTeamId()	
			if m.GetRadiantTeamName()!=""{
				NewCMsgDOTAMatch.RadiantTeamName = m.GetRadiantTeamName()
			}else{
				NewCMsgDOTAMatch.RadiantTeamName ="天辉"
			}

			if m.GetDireTeamName()!=""{
				NewCMsgDOTAMatch.DireTeamName = m.GetDireTeamName()
			}else{
				NewCMsgDOTAMatch.DireTeamName="夜魇"
			}
			
			NewCMsgDOTAMatch.MatchSeqNum = m.GetMatchSeqNum()
			NewCMsgDOTAMatch.RadiantTeamTag = m.GetRadiantTeamTag()
			NewCMsgDOTAMatch.DireTeamTag = m.GetDireTeamTag()
			NewCMsgDOTAMatch.SeriesId = m.GetSeriesId()
			NewCMsgDOTAMatch.SeriesType = m.GetSeriesType()
			NewCMsgDOTAMatch.Engine = m.GetEngine()
			NewCMsgDOTAMatch.PrivateMetadataKey = m.GetPrivateMetadataKey()
			NewCMsgDOTAMatch.MatchOutcome = m.GetMatchOutcome()
			NewCMsgDOTAMatch.TournamentId = m.GetTournamentId()
			NewCMsgDOTAMatch.TournamentRound = m.GetTournamentRound()
			NewCMsgDOTAMatch.PreGameDuration = m.GetPreGameDuration()
			return nil
		})	
		p.Callbacks.OnCDOTAMatchMetadataFile(func(m *dota.CDOTAMatchMetadataFile) error {
			arr:=m.GetMetadata().GetTeams()
			for _,value:=range arr{			
				if(value.GetDotaTeam()==2){
					var dotateam uint32=value.GetDotaTeam()
					players:=value.GetPlayers()
					for playerskey,playersvalue:=range players{	
						playersarr[playerskey].AccountId=playersvalue.GetAccountId()
						playersarr[playerskey].GameTeam=dotateam
						playersarr[playerskey].PlayerSlot=playersvalue.GetPlayerSlot()
						LevelUpTimes:=playersvalue.GetLevelUpTimes()
						playersarr[playerskey].LevelUpTimes=LevelUpTimes
						for _,lvluptime:=range LevelUpTimes{
							uptime:=obj.Durformat(int64(lvluptime))
							playersarr[playerskey].LevelUpTimeStr=append(playersarr[playerskey].LevelUpTimeStr,uptime)
						}
						abilityUpgrades:=playersvalue.GetAbilityUpgrades()
						playersarr[playerskey].AbilityUpgrades=abilityUpgrades
						for _,abilityid:=range abilityUpgrades{
							LoopAbility:
							for _,allvalue:=range allability.Ability{
								abilitystr:=strconv.Itoa(int(abilityid))
								if abilitystr==allvalue.Id{
									playersarr[playerskey].AbilityData=append(playersarr[playerskey].AbilityData,allvalue)
									break LoopAbility					
								}
							}
						}
						//AbilityArr[playerskey]=playersvalue.GetAbilityUpgrades()
						purchaseInfo:=playersvalue.GetInventorySnapshot()
						items:=playersvalue.GetItems()
						for _,itemvalue:=range purchaseInfo{
							gametime:=itemvalue.GetGameTime()
							if gametime%60==0{
								InventorySnapshot:=&obj.PlayerInventorySnapshot{}
								InventorySnapshot.GameTime=itemvalue.GetGameTime()
								InventorySnapshot.GameTimeStr=obj.Durformat(int64(InventorySnapshot.GameTime))
								InventorySnapshot.Kills=itemvalue.GetKills()
								InventorySnapshot.Deaths=itemvalue.GetDeaths()
								InventorySnapshot.Assists=itemvalue.GetAssists()
								InventorySnapshot.Level=itemvalue.GetLevel()
								itemidarr:=itemvalue.GetItemId()
								for _,itemid:=range itemidarr{
									itemicon:=ItemIDIcon[itemid]
									InventorySnapshot.Items=append(InventorySnapshot.Items,itemicon)
								}
								playersarr[playerskey].InventorySnapshot=append(playersarr[playerskey].InventorySnapshot,InventorySnapshot)
							}
						}
						for _,item:=range items{
							//if item.GetItemId()!=16&&item.GetItemId()!=17&&item.GetItemId()!=18&&item.GetItemId()!=19&&item.GetItemId()!=20&&item.GetItemId()!=40&&item.GetItemId()!=42&&item.GetItemId()!=43&&item.GetItemId()!=46{
							itemid:=item.GetItemId()
							itemdetails:=&obj.ItemDetails{}
							if item.GetPurchaseTime()<=300&&item.GetItemId()!=46{
								if item.GetPurchaseTime()<=0{
									itemdetails.PurchaseTimeNum=item.GetPurchaseTime()			
									itemdetails.ItemId=itemid
									itemdetails.ItemIcon=ItemIDIcon[itemid]
									PurchaseTime:=item.GetPurchaseTime()
									itemdetails.PurchaseTime=obj.Durformat(int64(PurchaseTime))
									playersarr[playerskey].Item=append(playersarr[playerskey].Item,itemdetails)
								}else{
									if itemid!=12&&itemid!=13&&itemid!=14&&itemid!=15&&itemid!=17&&itemid!=18&&itemid!=19&&itemid!=20&&itemid!=25&&itemid!=26&&itemid!=27&&itemid!=28&&itemid!=29&&itemid!=40&&itemid!=42&&itemid!=43&&itemid!=46&&itemid!=2{
										itemdetails.PurchaseTimeNum=item.GetPurchaseTime()			
										itemdetails.ItemId=itemid
										itemdetails.ItemIcon=ItemIDIcon[itemid]
										PurchaseTime:=item.GetPurchaseTime()
										itemdetails.PurchaseTime=obj.Durformat(int64(PurchaseTime))
										playersarr[playerskey].Item=append(playersarr[playerskey].Item,itemdetails)
									}
								}
								if itemid==38||itemid==39||itemid==44||itemid==216||itemid==237||itemid==43{
									consumable:=&obj.PlayerConsumable{}
									itemicon:=ItemIDIcon[itemid]
									consumable.ItemIcon=itemicon
									consumable.Count=1
									consumable.ItemId=itemid
									consumable.TimeSlice=append(consumable.TimeSlice,item.GetPurchaseTime())
									if len(playersarr[playerskey].Consumable)>0{
										Loopcon:
										for consumableKey,consumableValue:=range playersarr[playerskey].Consumable{
											if itemicon==consumableValue.ItemIcon{
												consumableValue.Count=consumableValue.Count+1
												break Loopcon
											}else{
												if consumableKey==(len(playersarr[playerskey].Consumable)-1){
													playersarr[playerskey].Consumable=append(playersarr[playerskey].Consumable,consumable)
												}	
											}
										}
									}else{
										playersarr[playerskey].Consumable=append(playersarr[playerskey].Consumable,consumable)
									}
								}
							}
							if item.GetPurchaseTime()>300&&itemid!=12&&itemid!=13&&itemid!=14&&itemid!=15&&itemid!=17&&itemid!=18&&itemid!=19&&itemid!=20&&itemid!=25&&itemid!=26&&itemid!=27&&itemid!=28&&itemid!=29&&itemid!=40&&itemid!=42&&itemid!=43&&itemid!=46&&itemid!=2{
								itemdetails.PurchaseTimeNum=item.GetPurchaseTime()			
								itemdetails.ItemId=itemid
								itemdetails.ItemIcon=ItemIDIcon[itemid]
								PurchaseTime:=item.GetPurchaseTime()
								itemdetails.PurchaseTime=obj.Durformat(int64(PurchaseTime))
								if PurchaseTime<=600{
									playersarr[playerskey].Item=append(playersarr[playerskey].Item,itemdetails)	
								}
								if PurchaseTime>600&&itemdetails.ItemId!=38{
									playersarr[playerskey].Item=append(playersarr[playerskey].Item,itemdetails)
								}
								if PurchaseTime<=600{
									if itemid==38||itemid==39||itemid==44||itemid==216||itemid==237||itemid==43{
									consumable:=&obj.PlayerConsumable{}
									itemicon:=ItemIDIcon[itemid]
									consumable.ItemIcon=itemicon
									consumable.Count=1
									consumable.ItemId=itemid
									consumable.TimeSlice=append(consumable.TimeSlice,PurchaseTime)
										if len(playersarr[playerskey].Consumable)>0{
											Loopcon1:
											for consumableKey,consumableValue:=range playersarr[playerskey].Consumable{
												if itemicon==consumableValue.ItemIcon{
													consumableValue.Count=consumableValue.Count+1
													break Loopcon1
												}else{
													if consumableKey==(len(playersarr[playerskey].Consumable)-1){
														playersarr[playerskey].Consumable=append(playersarr[playerskey].Consumable,consumable)
													}	
												}
											}
										}else{
											playersarr[playerskey].Consumable=append(playersarr[playerskey].Consumable,consumable)
										}
									}
								}

							}
						}
					}
				}
				if (value.GetDotaTeam()==3){
					var dotateam uint32=value.GetDotaTeam()
					players:=value.GetPlayers()
					for playerskey,playersvalue:=range players{	
						playerskey2:=playerskey+5
						playersarr[playerskey2].AccountId=playersvalue.GetAccountId()
						playersarr[playerskey2].GameTeam=dotateam
						playersarr[playerskey2].PlayerSlot=playersvalue.GetPlayerSlot()
						LevelUpTimes:=playersvalue.GetLevelUpTimes()
						playersarr[playerskey2].LevelUpTimes=LevelUpTimes
						for _,lvluptime:=range LevelUpTimes{
							uptime:=obj.Durformat(int64(lvluptime))
							playersarr[playerskey2].LevelUpTimeStr=append(playersarr[playerskey2].LevelUpTimeStr,uptime)
						}
						abilityUpgrades:=playersvalue.GetAbilityUpgrades()
						playersarr[playerskey2].AbilityUpgrades=abilityUpgrades
						for _,abilityid:=range abilityUpgrades{
							LoopAbility1:
							for _,allvalue:=range allability.Ability{
								abilitystr:=strconv.Itoa(int(abilityid))
								if abilitystr==allvalue.Id{
									playersarr[playerskey2].AbilityData=append(playersarr[playerskey2].AbilityData,allvalue)
									break LoopAbility1				
								}
							}
						}
						purchaseInfo:=playersvalue.GetInventorySnapshot()
						items:=playersvalue.GetItems()
						for _,itemvalue:=range purchaseInfo{
							gametime:=itemvalue.GetGameTime()
							if gametime%60==0{
								InventorySnapshot:=&obj.PlayerInventorySnapshot{}
								InventorySnapshot.GameTime=itemvalue.GetGameTime()
								InventorySnapshot.GameTimeStr=obj.Durformat(int64(InventorySnapshot.GameTime))
								InventorySnapshot.Kills=itemvalue.GetKills()
								InventorySnapshot.Deaths=itemvalue.GetDeaths()
								InventorySnapshot.Assists=itemvalue.GetAssists()
								InventorySnapshot.Level=itemvalue.GetLevel()
								itemidarr:=itemvalue.GetItemId()
								for _,itemid:=range itemidarr{
									itemicon:=ItemIDIcon[itemid]
									InventorySnapshot.Items=append(InventorySnapshot.Items,itemicon)
								}
								playersarr[playerskey2].InventorySnapshot=append(playersarr[playerskey2].InventorySnapshot,InventorySnapshot)
							}
						}
						for _,item:=range items{
							itemid:=item.GetItemId()
							itemdetails:=&obj.ItemDetails{}
							if item.GetPurchaseTime()<=300&&itemid!=46{
								if item.GetPurchaseTime()<0{
									itemdetails.PurchaseTimeNum=item.GetPurchaseTime()
									itemdetails.ItemId=itemid
									itemdetails.ItemIcon=ItemIDIcon[itemid]
									PurchaseTime:=item.GetPurchaseTime()
									itemdetails.PurchaseTime=obj.Durformat(int64(PurchaseTime))
									playersarr[playerskey2].Item=append(playersarr[playerskey2].Item,itemdetails)
								}else{
									if itemid!=12&&itemid!=13&&itemid!=14&&itemid!=15&&itemid!=17&&itemid!=18&&itemid!=19&&itemid!=20&&itemid!=25&&itemid!=26&&itemid!=27&&itemid!=28&&itemid!=29&&itemid!=40&&itemid!=42&&itemid!=43&&itemid!=46&&itemid!=2{
										itemdetails.PurchaseTimeNum=item.GetPurchaseTime()
										itemdetails.ItemId=itemid
										itemdetails.ItemIcon=ItemIDIcon[itemid]
										PurchaseTime:=item.GetPurchaseTime()
										itemdetails.PurchaseTime=obj.Durformat(int64(PurchaseTime))
										playersarr[playerskey2].Item=append(playersarr[playerskey2].Item,itemdetails)
									}
								}
								if itemid==38||itemid==39||itemid==44||itemid==216||itemid==237||itemid==43{
									consumable:=&obj.PlayerConsumable{}
									itemicon:=ItemIDIcon[itemid]
									consumable.ItemIcon=itemicon
									consumable.Count=1
									consumable.ItemId=itemid
									consumable.TimeSlice=append(consumable.TimeSlice,item.GetPurchaseTime())
									if len(playersarr[playerskey2].Consumable)>0{
										Loopcon2:
										for consumableKey,consumableValue:=range playersarr[playerskey2].Consumable{
											if itemicon==consumableValue.ItemIcon{
												consumableValue.Count=consumableValue.Count+1
												break Loopcon2
											}else{
												if consumableKey==(len(playersarr[playerskey2].Consumable)-1){
													playersarr[playerskey2].Consumable=append(playersarr[playerskey2].Consumable,consumable)
												}	
											}
										}
									}else{
										playersarr[playerskey2].Consumable=append(playersarr[playerskey2].Consumable,consumable)
									}
								}
							}
							if item.GetPurchaseTime()>300&&itemid!=12&&itemid!=13&&itemid!=14&&itemid!=15&&itemid!=17&&itemid!=18&&itemid!=19&&itemid!=20&&itemid!=25&&itemid!=26&&itemid!=27&&itemid!=28&&itemid!=29&&itemid!=40&&itemid!=42&&itemid!=43&&itemid!=46&&itemid!=2{
								itemdetails.PurchaseTimeNum=item.GetPurchaseTime()
								itemdetails.ItemIcon=ItemIDIcon[itemid]
								itemdetails.ItemId=itemid
								PurchaseTime:=item.GetPurchaseTime()
								itemdetails.PurchaseTime=obj.Durformat(int64(PurchaseTime))
								if itemdetails.PurchaseTimeNum<=600{
									playersarr[playerskey2].Item=append(playersarr[playerskey2].Item,itemdetails)
								}
								if itemdetails.ItemId!=38&&itemdetails.PurchaseTimeNum>600{
									playersarr[playerskey2].Item=append(playersarr[playerskey2].Item,itemdetails)
								}
								
								if PurchaseTime<=600{
									if itemid==38||itemid==39||itemid==44||itemid==216||itemid==237||itemid==43{
									consumable:=&obj.PlayerConsumable{}
									itemicon:=ItemIDIcon[itemid]
									consumable.ItemIcon=itemicon
									consumable.Count=1
									consumable.ItemId=itemid
									consumable.TimeSlice=append(consumable.TimeSlice,PurchaseTime)
										if len(playersarr[playerskey2].Consumable)>0{
											Loopcon3:
											for consumableKey,consumableValue:=range playersarr[playerskey2].Consumable{
												if itemicon==consumableValue.ItemIcon{
													consumableValue.Count=consumableValue.Count+1
													break Loopcon3
												}else{
													if consumableKey==(len(playersarr[playerskey2].Consumable)-1){
														playersarr[playerskey2].Consumable=append(playersarr[playerskey2].Consumable,consumable)
													}	
												}
											}
										}else{
											playersarr[playerskey2].Consumable=append(playersarr[playerskey2].Consumable,consumable)
										}
									}
								}
							}
						}	
					}
				}
			}
			return nil
		})

		p.Callbacks.OnCMsgDOTACombatLogEntry(func(m *dota.CMsgDOTACombatLogEntry) error {
			msgtype := m.GetType()
			time := m.GetTimestamp()
			if msgtype == 8 && msgtime == 0 {
				starttime = m.GetTimestamp()
				//fmt.Printf("starttime:%v\n", starttime)
				msgtime = msgtime + 1
				prev5 = float64(starttime) + 300.0 + 90.0
				prev10 = math.Floor(prev5) + 300.0
				prev15 = math.Floor(prev10) + 300.0
				prev20 = math.Floor(prev15) + 300.0
				prev25 = math.Floor(prev20) + 300.0
				prev30 = math.Floor(prev25) + 300.0
			}
			if len(heroslice) < 10 {
				if msgtype == 8 {
					target := m.GetTargetName()
					heroslice = append(heroslice, int(target))
				}
			} else {
				sort.Ints(heroslice)
			}
			if time > starttime {
				//分路
				if time>=(starttime+100.0){
					if msgtype == 10 && (math.Floor(float64(time)) <= math.Ceil(prev10)) {
						target := int(m.GetTargetName())
						x := m.GetLocationX()
						y := m.GetLocationY()
						initlocation := ""
						
						if(x>-1850.0&&x<1000.0&&y>-1750.0&&y<700.0){
	                		initlocation = "中路"
	            		}
						if x < 0 && y > 0&&initlocation!="中路"{
								initlocation = "上路"
						}
						if x > 0 && y < 0 && initlocation!="中路"{
							initlocation = "下路"
						} 
						if initlocation==""{
							initlocation="打野"
						}
						if lanetimei1<=2||lanetimei2<=2||lanetimei3<=2||lanetimei4<=2||lanetimei5<=2||lanetimei6<=2||lanetimei7<=2||lanetimei8<=2||lanetimei9<=2||lanetimei10<=2{
							switch target {
							case heroslice[0]:
								if hero1initlocation != initlocation{
									hero1initlocation = initlocation
									if (playersarr[0].InitLane=="") {
										playersarr[0].InitLane=initlocation
										if playersarr[0].InitLane=="上路"{
											radtopherocount=radtopherocount+1
										}
										if playersarr[0].InitLane=="下路"{
											radbottomherocount=radbottomherocount+1
										}
										if playersarr[0].InitLane=="中路"{
											radmidherocount=radmidherocount+1
										}
									}
									 
									if  (hero1initlocation!=playersarr[0].InitLane){
										playersarr[0].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei1=lanetimei1+1
									}
								}
							case heroslice[1]:
								if hero2initlocation != initlocation {
									hero2initlocation = initlocation
				
									if (playersarr[1].InitLane=="") {
										playersarr[1].InitLane=initlocation
										if playersarr[1].InitLane=="上路"{
											radtopherocount=radtopherocount+1
										}
										if playersarr[1].InitLane=="下路"{
											radbottomherocount=radbottomherocount+1
										}
										if playersarr[1].InitLane=="中路"{
											radmidherocount=radmidherocount+1
										}
									}
									
									if  (hero2initlocation!=playersarr[1].InitLane){
										playersarr[1].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei2=lanetimei2+1
									}					
								}
							case heroslice[2]:
								if hero3initlocation != initlocation {
									hero3initlocation = initlocation

									if (playersarr[2].InitLane=="") {
										playersarr[2].InitLane=initlocation
										if playersarr[2].InitLane=="上路"{
											radtopherocount=radtopherocount+1
										}
										if playersarr[2].InitLane=="下路"{
											radbottomherocount=radbottomherocount+1
										}
										if playersarr[2].InitLane=="中路"{
											radmidherocount=radmidherocount+1
										}
									}
									
									if  (hero3initlocation!=playersarr[2].InitLane){
										playersarr[2].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei3=lanetimei3+1
									}	
								}
							case heroslice[3]:
								if hero4initlocation != initlocation {
									hero4initlocation = initlocation
									
									if (playersarr[3].InitLane=="") {
										playersarr[3].InitLane=initlocation
										if playersarr[3].InitLane=="上路"{
											radtopherocount=radtopherocount+1
										}
										if playersarr[3].InitLane=="下路"{
											radbottomherocount=radbottomherocount+1
										}
										if playersarr[3].InitLane=="中路"{
											radmidherocount=radmidherocount+1
										}
									}

									
									if  (hero4initlocation!=playersarr[3].InitLane){
										playersarr[3].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei4=lanetimei4+1
									}	
								}
							case heroslice[4]:
								if hero5initlocation != initlocation {
									hero5initlocation = initlocation
									if (playersarr[4].InitLane=="") {
										playersarr[4].InitLane=initlocation
										if playersarr[4].InitLane=="上路"{
											radtopherocount=radtopherocount+1
										}
										if playersarr[4].InitLane=="下路"{
											radbottomherocount=radbottomherocount+1
										}
										if playersarr[4].InitLane=="中路"{
											radmidherocount=radmidherocount+1
										}
									}
									if  (hero5initlocation!=playersarr[4].InitLane){
										playersarr[4].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei5=lanetimei5+1
									}	
								}
							case heroslice[5]:
								if hero6initlocation != initlocation {
									hero6initlocation = initlocation
									if (playersarr[5].InitLane=="") {
										playersarr[5].InitLane=initlocation
										if playersarr[5].InitLane=="上路"{
											diretopherocount=diretopherocount+1
										}
										if playersarr[5].InitLane=="下路"{
											direbottomherocount=direbottomherocount+1
										}
										if playersarr[5].InitLane=="中路"{
											diremidherocount=diremidherocount+1
										}
									}
									
									if  (hero6initlocation!=playersarr[5].InitLane){
										playersarr[5].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei6=lanetimei6+1
									}	
								}
							case heroslice[6]:
								if hero7initlocation != initlocation {
									hero7initlocation = initlocation
									if (playersarr[6].InitLane=="") {
										playersarr[6].InitLane=initlocation
										if playersarr[6].InitLane=="上路"{
											diretopherocount=diretopherocount+1
										}
										if playersarr[6].InitLane=="下路"{
											direbottomherocount=direbottomherocount+1
										}
										if playersarr[6].InitLane=="中路"{
											diremidherocount=diremidherocount+1
										}
									}

									if  (hero7initlocation!=playersarr[6].InitLane){
										playersarr[6].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei7=lanetimei7+1
									}	
								}
							case heroslice[7]:
								if hero8initlocation != initlocation {
									hero8initlocation = initlocation
									if (playersarr[7].InitLane=="") {
										playersarr[7].InitLane=initlocation
										if playersarr[7].InitLane=="上路"{
											diretopherocount=diretopherocount+1
										}
										if playersarr[7].InitLane=="下路"{
											direbottomherocount=direbottomherocount+1
										}
										if playersarr[7].InitLane=="中路"{
											diremidherocount=diremidherocount+1
										}
									}

									
									if  (hero8initlocation!=playersarr[7].InitLane){
										playersarr[7].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei8=lanetimei8+1
									}	
								}
							case heroslice[8]:
								if hero9initlocation != initlocation {
									hero9initlocation = initlocation
									if (playersarr[8].InitLane=="") {
										playersarr[8].InitLane=initlocation
										if playersarr[8].InitLane=="上路"{
											diretopherocount=diretopherocount+1
										}
										if playersarr[8].InitLane=="下路"{
											direbottomherocount=direbottomherocount+1
										}
										if playersarr[8].InitLane=="中路"{
											diremidherocount=diremidherocount+1
										}
									}		

									if  (hero9initlocation!=playersarr[8].InitLane){
										playersarr[8].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei9=lanetimei1+1
									}	
								}
							case heroslice[9]:
								if hero10initlocation != initlocation {
									hero10initlocation = initlocation
									if (playersarr[9].InitLane=="") {
										playersarr[9].InitLane=initlocation
										if playersarr[9].InitLane=="上路"{
											diretopherocount=diretopherocount+1
										}
										if playersarr[9].InitLane=="下路"{
											direbottomherocount=direbottomherocount+1
										}
										if playersarr[9].InitLane=="中路"{
											diremidherocount=diremidherocount+1
										}
									}

									if  (hero10initlocation!=playersarr[9].InitLane){
										playersarr[9].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
										lanetimei10=lanetimei10+1
									}	
								}
							}
						}
					}	
				}
				
				//补刀与反补和 塔状态
				if msgtype == 4 {
					istargettower1 := m.GetIsTargetBuilding()
					//attackerteam1 := m.GetAttackerTeam()
					targetteam1 := m.GetTargetTeam()
					istargethero1:=m.GetIsTargetHero()
					for key, value := range heroslice {
						istargethero := m.GetIsTargetHero()
						istargettower := m.GetIsTargetBuilding()
						damagesource := int(m.GetDamageSourceName())
						attackerteam := m.GetAttackerTeam()
						target := m.GetTargetName()
						targetteam := m.GetTargetTeam()
						targetsource:=m.GetTargetSourceName()
						//正补反补
						if target==99&&damagesource == value{								
									playersarr[key].KillCourierCount=playersarr[key].KillCourierCount+1
						}
						if target==99&&targetsource == uint32(value){
									deathtime:=float64(time-starttime-90.0)
									deathtimeuint:=uint32(int(math.Floor(deathtime)))
									playersarr[key].CourierDeathTime=append(playersarr[key].CourierDeathTime,deathtimeuint)
						}
						if istargethero == false && target != 279 && damagesource == value && target != 184 {
							if math.Floor(float64(time)) > prev5 && timei[key] < 1 && !istargettower {
								var herolasthit, herodenies uint32
								herolasthit = m.GetLastHits()
								herodenies = denies[key]
								if attackerteam != targetteam {
									herolasthit = m.GetLastHits() - 1
								} else {
									herodenies = denies[key] - 1
								}

								playersarr[key].LastHit=&obj.HeroLastHit{}
								playersarr[key].Denis=&obj.HeroLastHit{}
								playersarr[key].LastHit.Pre5Count=herolasthit
								playersarr[key].Denis.Pre5Count=herodenies
								//*dota2lasthit[key] = strconv.FormatInt(herolasthit, 10) + " " + strconv.FormatInt(herodenies, 10) + ","
								if playersarr[key].Gold==nil{
									playersarr[key].Gold=&obj.HeroGold{}
									if playersarr[key].Gold.Pre5MinGold==nil{
										playersarr[key].Gold.Pre5MinGold=&obj.GoldType{}	
									}
									playersarr[key].Gold.Pre5MinGold.CreepsGold=creepsgold[key]//野生生物
								}else{
									if playersarr[key].Gold.Pre5MinGold==nil{
										playersarr[key].Gold.Pre5MinGold=&obj.GoldType{}	
									}
									playersarr[key].Gold.Pre5MinGold.CreepsGold=creepsgold[key]//野生生物
								}
								creepsgold[key]=0
								timei[key] = timei[key] + 1
							}
							if math.Floor(float64(time)) > prev10 && timei[key] < 2 && !istargettower {
								var herolasthit, herodenies uint32
								herolasthit = m.GetLastHits()
								herodenies = denies[key]
								if attackerteam != targetteam {
									herolasthit = m.GetLastHits() - 1
								} else {
									herodenies = denies[key] - 1
								}
								playersarr[key].LastHit.Pre10Count=herolasthit
								playersarr[key].Denis.Pre10Count=herodenies
								if playersarr[key].Gold.Pre10MinGold==nil{
									playersarr[key].Gold.Pre10MinGold=&obj.GoldType{}	
								}
								playersarr[key].Gold.Pre10MinGold.CreepsGold=creepsgold[key]//野生生物
								creepsgold[key]=0
								timei[key] = timei[key] + 1
							}
							if math.Floor(float64(time)) > prev15 && timei[key] < 3 && !istargettower {
								var herolasthit, herodenies uint32
								herolasthit = m.GetLastHits()
								herodenies = denies[key]
								if attackerteam != targetteam {
									herolasthit = m.GetLastHits() - 1
								} else {
									herodenies = denies[key] - 1
								}
								playersarr[key].LastHit.Pre15Count=herolasthit
								playersarr[key].Denis.Pre15Count=herodenies
								if playersarr[key].Gold.Pre15MinGold==nil{
									playersarr[key].Gold.Pre15MinGold=&obj.GoldType{}	
								}
								playersarr[key].Gold.Pre15MinGold.CreepsGold=creepsgold[key]//野生生物
								creepsgold[key]=0
								timei[key] = timei[key] + 1
								//playersarr[key].Gold.Pre15MinGold.CreepsGold=creepsgold[key]
							}
							if math.Floor(float64(time)) > prev20 && timei[key] < 4 && !istargettower {
								var herolasthit, herodenies uint32
								herolasthit = m.GetLastHits()
								herodenies = denies[key]
								if attackerteam != targetteam {
									herolasthit = m.GetLastHits() - 1
								} else {
									herodenies = denies[key] - 1
								}
								playersarr[key].LastHit.Pre20Count=herolasthit
								playersarr[key].Denis.Pre20Count=herodenies
								if playersarr[key].Gold.Pre20MinGold==nil{
									playersarr[key].Gold.Pre20MinGold=&obj.GoldType{}	
								}
								playersarr[key].Gold.Pre20MinGold.CreepsGold=creepsgold[key]//野生生物
								creepsgold[key]=0
								timei[key] = timei[key] + 1
								//playersarr[key].Gold.Pre20MinGold.CreepsGold=creepsgold[key]
							}
							if math.Floor(float64(time)) > prev25 && timei[key] < 5 && !istargettower {
								var herolasthit, herodenies uint32
								herolasthit = m.GetLastHits()
								herodenies = denies[key]
								if attackerteam != targetteam {
									herolasthit = m.GetLastHits() - 1
								} else {
									herodenies = denies[key] - 1
								}
								playersarr[key].LastHit.Pre25Count=herolasthit
								playersarr[key].Denis.Pre25Count=herodenies
								if playersarr[key].Gold.Pre25MinGold==nil{
									playersarr[key].Gold.Pre25MinGold=&obj.GoldType{}	
								}
								playersarr[key].Gold.Pre25MinGold.CreepsGold=creepsgold[key]//野生生物
								creepsgold[key]=0
								timei[key] = timei[key] + 1
								//playersarr[key].Gold.Pre25MinGold.CreepsGold=creepsgold[key]
							}
							if attackerteam == targetteam {
								denies[key] = denies[key] + 1
							}

							if targetteam!=2&&targetteam!=3{
								creepsgold[key]=creepsgold[key]+pregold
							}
						}
					}
					if istargethero1{
						NewKDADetails:=&obj.KDADetails{}
						NewKDADetails.Attacker=strconv.Itoa(int(m.GetAttackerName()))
						NewKDADetails.DeathTarget=strconv.Itoa(int(m.GetTargetName()))
						assistant:=m.GetAssistPlayers()
						for _,asvalue:=range assistant{
							indexstr:=strconv.Itoa(int(asvalue))
							NewKDADetails.Assistant=append(NewKDADetails.Assistant,indexstr)
						}
						killtime:=time-90.0-starttime
						killtimeint:=int64(math.Floor(float64(killtime)))
						NewKDADetails.Time=killtimeint
						NewKDADetails.TimeStr=obj.Durformat(killtimeint)
						NewCMsgMatchDetails.KDADets=append(NewCMsgMatchDetails.KDADets,NewKDADetails)
					}
					if istargettower1 {
						if math.Floor(float64(time)) > prev10 && towertimei[0] < 1 && targetteam1 == 3 {
							NewCMsgMatchDetails.RadiantDestoryTower.Pre10MinTower = radiantkilltower
							towertimei[0] = towertimei[0] + 1
						}
						if math.Floor(float64(time)) > prev20 && towertimei[0] < 2 && targetteam1 == 3 {
							NewCMsgMatchDetails.RadiantDestoryTower.Pre20MinTower = radiantkilltower
							towertimei[0] = towertimei[0] + 1
						}
						if math.Floor(float64(time)) > prev30 && towertimei[0] < 3 && targetteam1 == 3 {
							NewCMsgMatchDetails.RadiantDestoryTower.Pre30MinTower = radiantkilltower
							towertimei[0] = towertimei[0] + 1
						}
						if math.Floor(float64(time)) > prev10 && towertimei[1] < 1 && targetteam1 == 2 {
							NewCMsgMatchDetails.DireDestoryTower.Pre10MinTower = direkilltower
							towertimei[1] = towertimei[1] + 1
						}
						if math.Floor(float64(time)) > prev20 && towertimei[1] < 2 && targetteam1 == 2 {
							NewCMsgMatchDetails.DireDestoryTower.Pre20MinTower = direkilltower
							towertimei[1] = towertimei[1] + 1
						}
						if math.Floor(float64(time)) > prev30 && towertimei[1] < 3 && targetteam1 == 2 {
							NewCMsgMatchDetails.DireDestoryTower.Pre30MinTower = direkilltower
							towertimei[1] = towertimei[1] + 1
						}
						if targetteam1 == 3 {
							radiantkilltower = radiantkilltower + 1
						}
						if targetteam1 == 2 {
							direkilltower = direkilltower + 1
						}
					}
				}
				//金钱
				if msgtype == 8 {
					target := int(m.GetTargetName())
					//premsgtype=msgtype
					for key1, value1 := range heroslice {
						if target == value1 {
							if math.Floor(float64(time)) > prev5 && gtimei[key1] < 1 {

								if playersarr[key1].Gold==nil{
									playersarr[key1].Gold=&obj.HeroGold{}
									if playersarr[key1].Gold.Pre5MinGold==nil{
										playersarr[key1].Gold.Pre5MinGold=&obj.GoldType{}
										playersarr[key1].Gold.Pre5MinGold.AllGold=math.Floor(allgold[key1])
										playersarr[key1].Gold.Pre5MinGold.LastHitGold=lasthitgold[key1]
										playersarr[key1].Gold.Pre5MinGold.CombatGold=combatgold[key1]
									}else{
										playersarr[key1].Gold.Pre5MinGold.AllGold=math.Floor(allgold[key1])
										playersarr[key1].Gold.Pre5MinGold.LastHitGold=lasthitgold[key1]
										playersarr[key1].Gold.Pre5MinGold.CombatGold=combatgold[key1]
									}
								}else{
									if playersarr[key1].Gold.Pre5MinGold==nil{
										playersarr[key1].Gold.Pre5MinGold=&obj.GoldType{}
										playersarr[key1].Gold.Pre5MinGold.AllGold=math.Floor(allgold[key1])
										playersarr[key1].Gold.Pre5MinGold.LastHitGold=lasthitgold[key1]
										playersarr[key1].Gold.Pre5MinGold.CombatGold=combatgold[key1]
									}else{
										playersarr[key1].Gold.Pre5MinGold.AllGold=math.Floor(allgold[key1])
										playersarr[key1].Gold.Pre5MinGold.LastHitGold=lasthitgold[key1]
										playersarr[key1].Gold.Pre5MinGold.CombatGold=combatgold[key1]
									}
								}
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								itembackgold[key1] = 0	
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
								gtimei[key1] = gtimei[key1] + 1
							}
							if math.Floor(float64(time)) > prev10 && gtimei[key1] < 2 {
								gtimei[key1] = gtimei[key1] + 1
								//math.Floor(allgold[key1]) = math.Floor(allgold[key1])
								if playersarr[key1].Gold.Pre10MinGold==nil{
									playersarr[key1].Gold.Pre10MinGold=&obj.GoldType{}
									playersarr[key1].Gold.Pre10MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre10MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre10MinGold.CombatGold=combatgold[key1]
								}else{
									playersarr[key1].Gold.Pre10MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre10MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre10MinGold.CombatGold=combatgold[key1]
								}
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								itembackgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							if math.Floor(float64(time)) > prev15 && gtimei[key1] < 3 {
								gtimei[key1] = gtimei[key1] + 1
								if playersarr[key1].Gold.Pre15MinGold==nil{
									playersarr[key1].Gold.Pre15MinGold=&obj.GoldType{}
									playersarr[key1].Gold.Pre15MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre15MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre15MinGold.CombatGold=combatgold[key1]
								}else{
									playersarr[key1].Gold.Pre15MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre15MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre15MinGold.CombatGold=combatgold[key1]
								}
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								itembackgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							if math.Floor(float64(time)) > prev20 && gtimei[key1] < 4 {
								gtimei[key1] = gtimei[key1] + 1
								if playersarr[key1].Gold.Pre20MinGold==nil{
									playersarr[key1].Gold.Pre20MinGold=&obj.GoldType{}
									playersarr[key1].Gold.Pre20MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre20MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre20MinGold.CombatGold=combatgold[key1]
								}else{
									playersarr[key1].Gold.Pre20MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre20MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre20MinGold.CombatGold=combatgold[key1]
								}
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								itembackgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							if math.Floor(float64(time)) > prev25 && gtimei[key1] < 5 {
								gtimei[key1] = gtimei[key1] + 1
								if playersarr[key1].Gold.Pre25MinGold==nil{
									playersarr[key1].Gold.Pre25MinGold=&obj.GoldType{}
									playersarr[key1].Gold.Pre25MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre25MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre25MinGold.CombatGold=combatgold[key1]
								}else{
									playersarr[key1].Gold.Pre25MinGold.AllGold=math.Floor(allgold[key1])
									playersarr[key1].Gold.Pre25MinGold.LastHitGold=lasthitgold[key1]
									playersarr[key1].Gold.Pre25MinGold.CombatGold=combatgold[key1]
								}
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								itembackgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							u := float64(int32(m.GetValue()))
							if m.GetGoldReason()!=6{
								allgold[key1] = math.Floor(allgold[key1]) + u
								pregold=u
							} 		
							goldreason:=m.GetGoldReason()
							if goldreason == 0 {//吃符和排眼的钱 reason为0
								runegold[key1] = runegold[key1] + u
							}
							if goldreason == 1 {
								deathgold[key1] = deathgold[key1] + u
							}
							if goldreason == 2 {
								buybackgold[key1] = buybackgold[key1] + u
							}
							if goldreason == 6 {//6是买了装备后卖回去的钱
								itembackgold[key1] = itembackgold[key1] + u
							}
							if goldreason == 13 {
								lasthitgold[key1] = lasthitgold[key1] + u
							}
							if goldreason == 11 {
								killtowergold[key1] = killtowergold[key1] + u
							}
							if goldreason == 12 {
								combatgold[key1] = combatgold[key1] + u
							}
							if goldreason == 14 {
								roshangold[key1] = roshangold[key1] + u
							}
							if goldreason == 15 {
								killcouriergold[key1] = killcouriergold[key1] + u
							}
							
						}
					}
				}
			}
			return nil
		})
		p.Start()
		
		TopSideLineUp.Version=version
		BottomSideLineUp.Version=version

		NewCMsgDOTAMatch.MatchId=NewCDotaGameInfo.MatchId
		NewCMsgMatchDetails.MatchId=NewCDotaGameInfo.MatchId
		if NewCDotaGameInfo.RadiantTeamTag==""{
			NewCDotaGameInfo.RadiantTeamTag=NewCMsgDOTAMatch.RadiantTeamName	
		}
		if NewCDotaGameInfo.DireTeamTag==""{
			NewCDotaGameInfo.DireTeamTag=NewCMsgDOTAMatch.DireTeamName
		}
		
		if winnerid==2{
			NewCDotaGameInfo.GameWinner=NewCDotaGameInfo.RadiantTeamTag
			NewCDotaGameInfo.GameLoser=NewCDotaGameInfo.DireTeamTag
		}else{
			NewCDotaGameInfo.GameWinner=NewCDotaGameInfo.DireTeamTag 
			NewCDotaGameInfo.GameLoser=NewCDotaGameInfo.RadiantTeamTag
		}
		NewCMsgDOTAMatch.RadiantTeamTag=NewCDotaGameInfo.RadiantTeamTag
		NewCMsgDOTAMatch.DireTeamTag=NewCDotaGameInfo.DireTeamTag
		NewCMsgMatchDetails.RadiantTeamTag=NewCDotaGameInfo.RadiantTeamTag
		NewCMsgMatchDetails.DireTeamTag=NewCDotaGameInfo.DireTeamTag
		NewCMsgMatchDetails.RadiantTeamName=NewCDotaGameInfo.RadiantTeamName
		NewCMsgMatchDetails.DireTeamName=NewCDotaGameInfo.DireTeamName
		NewCMsgMatchDetails.Duration=NewCMsgDOTAMatch.Duration
		NewCMsgMatchDetails.DurationStr=NewCMsgDOTAMatch.DurationStr
		for key,value:=range playersarr{
			for _,kdavalue:=range NewCMsgMatchDetails.KDADets{
				if value.HeroName==kdavalue.DeathTarget{
					value.DeathDets=append(value.DeathDets,kdavalue)
				}
			}
			if value.InitLane==""{
				value.InitLane="打野"
			}
			if winnerid==2{
				if key<=4{
					value.GameResult="win"
				}else{
					value.GameResult="lose"
				}		
			}else{
				if key<=4{
					value.GameResult="lose"
				}else{
					value.GameResult="win"
				}
			}
			value.Player=*NewCMsgDOTAMatch.Players[key].Player.PlayerName
			heroid:=*NewCMsgDOTAMatch.Players[key].Player.HeroId
			value.HeroName=HeroIDName[heroid]
			value.HeroIcon=HeroNameIcon[value.HeroName] 
			NewCMsgDOTAMatch.Players[key].HeroIcon=HeroNameIcon[value.HeroName]

			//var tangoIronBranch int=0
			var fiveConsumablegold float64
			var tenConsumablegold float64
			var tangoIronBranchGold float64
			var fiveallIronBranch int=0
			var fiveitemUseIronBranch int=0
			var tenallIronBranch int=0
			var tenitemUseIronBranch int=0
			for _,ivalue:=range value.Item{
				if ivalue.ItemId==16&&ivalue.PurchaseTimeNum<=300{
					fiveallIronBranch=fiveallIronBranch+1
				}
				if (ivalue.ItemId==36||ivalue.ItemId==94||ivalue.ItemId==86||ivalue.ItemId==88)&&ivalue.PurchaseTimeNum<=300{
					fiveitemUseIronBranch=fiveitemUseIronBranch+1
				}
				if ivalue.PurchaseTimeNum>300&&ivalue.PurchaseTimeNum<=600&&ivalue.ItemId==16{
					tenallIronBranch=tenallIronBranch+1
				}
				if ivalue.PurchaseTimeNum>300&&ivalue.PurchaseTimeNum<=600&&(ivalue.ItemId==36||ivalue.ItemId==94||ivalue.ItemId==86||ivalue.ItemId==88){
					tenitemUseIronBranch=tenitemUseIronBranch+1
				} 
			}
			if fiveallIronBranch!=0{
				if fiveitemUseIronBranch==0&&tenitemUseIronBranch==0{
					tangoIronBranchGold=float64(fiveallIronBranch)*50.0
				}
				if fiveallIronBranch==fiveitemUseIronBranch{
					tangoIronBranchGold=0
				}
				if fiveitemUseIronBranch!=0&&fiveallIronBranch>fiveitemUseIronBranch{
					if tenitemUseIronBranch!=0&&tenallIronBranch==tenitemUseIronBranch{
						tangoIronBranchGold=float64(fiveallIronBranch-fiveitemUseIronBranch)*50.0
					}
					if tenitemUseIronBranch!=0&&tenallIronBranch!=tenitemUseIronBranch{
						tangoIronBranchGold= float64(tenitemUseIronBranch-tenallIronBranch)*50.0
					}
				}
			}
			for _,cvalue:=range value.Consumable{
				for _,ctime:=range cvalue.TimeSlice{
					if ctime<=300{
						if cvalue.ItemId==38{
							fiveConsumablegold=fiveConsumablegold+50.0
							tenConsumablegold=tenConsumablegold+50.0
						}
						if cvalue.ItemId==39{
							fiveConsumablegold=fiveConsumablegold+110.0
							tenConsumablegold=tenConsumablegold+110.0
						}
						if cvalue.ItemId==44{
							fiveConsumablegold=fiveConsumablegold+90.0
							tenConsumablegold=tenConsumablegold+90.0
						}
						if cvalue.ItemId==216{
							fiveConsumablegold=fiveConsumablegold+70.0
							tenConsumablegold=tenConsumablegold+70.0
						}
					}
					if ctime>300{
						if cvalue.ItemId==38{
							tenConsumablegold=tenConsumablegold+50.0
						}
						if cvalue.ItemId==39{
							tenConsumablegold=tenConsumablegold+110.0
						}
						if cvalue.ItemId==44{
							tenConsumablegold=tenConsumablegold+90.0
						}
						if cvalue.ItemId==216{
							tenConsumablegold=tenConsumablegold+70.0
						}
					}
				}
			}
			//ability:=value.AbilityUpgrades
			lvl:=value.LevelUpTimes
			//假设有信使死亡
			var allsalary float64
			var nowlvluptime uint32

			for lvlkey,nextlvluptime:=range lvl{
				 persecsalary:=(83.0+(2.0*float64(lvlkey+1)))/60.0
				 if lvlkey!=0{
				 	nowlvluptime=lvl[lvlkey-1]
				 }
				 perlvlsalary:=persecsalary*float64(nextlvluptime-nowlvluptime)
				 allsalary=allsalary+perlvlsalary
				 
				 if nextlvluptime>300&&nowlvluptime<=300{ 	
				 	value.Gold.Pre5MinGold.AllGold=math.Ceil(value.Gold.Pre5MinGold.AllGold+allsalary-persecsalary*float64(nextlvluptime-300)-fiveConsumablegold-tangoIronBranchGold)
				 }
				 if nextlvluptime>600&&nowlvluptime<=600{
				 	value.Gold.Pre10MinGold.AllGold=math.Ceil(value.Gold.Pre10MinGold.AllGold+allsalary-persecsalary*float64(nextlvluptime-600)-tenConsumablegold)
				 }
				 if nextlvluptime>900&&nowlvluptime<=900{
				 	value.Gold.Pre15MinGold.AllGold=math.Ceil(value.Gold.Pre15MinGold.AllGold+allsalary-persecsalary*float64(nextlvluptime-900))
				 }
				 if nextlvluptime>1200&&nowlvluptime<=1200{
				 	value.Gold.Pre20MinGold.AllGold=math.Ceil(value.Gold.Pre20MinGold.AllGold+allsalary-persecsalary*float64(nextlvluptime-1200))
				 }

				 if nextlvluptime>1500&&nowlvluptime<=1500{
				 	value.Gold.Pre25MinGold.AllGold=math.Ceil(value.Gold.Pre25MinGold.AllGold+allsalary-persecsalary*float64(nextlvluptime-1500))
				 }
			}
			newplayer:=obj.Player{}

			var playerquery string
			count,playerdberr:=playerdb.Find(bson.M{"player_dota2_register_num_id":value.AccountId}).Count()
			if count!=0{
				playerquery="player_dota2_register_num_id"
			}
			count2,playerdberr2:=playerdb.Find(bson.M{"second_num_id":value.AccountId}).Count()
			if count2!=0{
				playerquery="second_num_id"
			}
			count3,playerdberr3:=playerdb.Find(bson.M{"third_num_id":value.AccountId}).Count()
			if count3!=0{
				playerquery="third_num_id"
			}
			count4,playerdberr4:=playerdb.Find(bson.M{"fourth_num_id":value.AccountId}).Count()
			if count4!=0{
				playerquery="fourth_num_id"
			}
			if playerdberr!=nil{
				fmt.Printf("playerdberr:%v\n", playerdberr)
			}
			if playerdberr2!=nil{
				fmt.Printf("playerdberr2:%v\n", playerdberr2)
			}
			if playerdberr3!=nil{
				fmt.Printf("playerdberr3:%v\n", playerdberr3)
			}
			if playerdberr4!=nil{
				fmt.Printf("playerdberr4:%v\n", playerdberr4)
			}
			

			findplayererr:=playerdb.Find(bson.M{playerquery:value.AccountId}).One(&newplayer)
			//fmt.Printf("账号：%v 选手id名：%v\n", value.AccountId,value.Player)
			if findplayererr!=nil{
				fmt.Printf("查找选手出现错误：%v AccountId：%v\n", findplayererr,value.AccountId)
				var bt bytes.Buffer
				s1:="无法找到选手记录"
				s2:=value.AccountId
				s3:="\n"
				bt.WriteString(s1)
				bt.WriteString(s2)
				bt.WriteString(s3)
				playerisexist=playerisexist+bt.String()
			}
			value.Position=newplayer.Position
			if (value.InitLane=="上路"){	
				newAgainstHero:=&obj.AgainstHero{}
				newAgainstHero.HeroName=value.HeroName
				newAgainstHero.PlayerId=value.AccountId
				newAgainstHero.PlayerName=value.Player
				newAgainstHero.Pre10LastHit=value.LastHit.Pre10Count
				newAgainstHero.Pre10Denis=value.Denis.Pre10Count
				newAgainstHero.Pre10MinGold=value.Gold.Pre10MinGold
				newAgainstHero.Pre5LastHit=value.LastHit.Pre5Count
				newAgainstHero.Pre5Denis=value.Denis.Pre5Count
				Loopvaluelvl1:
				for vallvluptimekey,vallvluptimevalue:=range value.LevelUpTimes{
					if vallvluptimevalue>600{
						lvlk:=vallvluptimekey-1
						newAgainstHero.Pre10MinLevelUpTime=value.LevelUpTimes[0:lvlk]
						newAgainstHero.Pre10MinLevelUpTimeStr=value.LevelUpTimeStr[0:lvlk]
						newAgainstHero.AbilityUpgrades=value.AbilityData[0:vallvluptimekey]
						break Loopvaluelvl1
					}
				}
				//newAgainstHero.Pre10minItem=value.Item
				Loopvalueitem1:
				for valitemkey,valitemvalue:=range value.Item{
					if valitemvalue.PurchaseTimeNum>600{
						newAgainstHero.Pre10minItem=value.Item[0:valitemkey]
						break Loopvalueitem1
						}
					}
				newAgainstHero.HeroIcon=value.HeroIcon
				newAgainstHero.ConsumableItem=value.Consumable
				//newAgainstHero.AbilityUpgrades=value.AbilityData
				for _,insv:=range value.InventorySnapshot{
					if insv.GameTime==600{
						newAgainstHero.HeroDeathCount=int(insv.Deaths)
					}
				}
				TopSideLineUpDets.AgainstHeroes=append(TopSideLineUpDets.AgainstHeroes,newAgainstHero)

				if(value.GameTeam==2){
					if radtopherocount==2{
						if value.Position=="四号位"||value.Position=="五号位"{
							yslane=yslane+value.HeroName							
							//topsideLineUp.Side1Heroes=yslane
						}else{
							TopSideLineUp.HardLaneMainPositionPlayerId=value.AccountId
							yslane=value.HeroName+" "+yslane
							//topsideLineUp.Side1Heroes=yslane
						}	
					}
					if radtopherocount==3{
						if value.Position=="五号位"{
							yslane=yslane+" 五号位"+value.HeroName
						}
						if value.Position!="五号位"&&value.Position!="四号位"{
							yslane=value.HeroName+"大哥位 "+yslane
							TopSideLineUp.HardLaneMainPositionPlayerId=value.AccountId
						}
						if value.Position=="四号位"{
							if yslane!=""{
								if strings.Contains(yslane,"大哥位")&&!strings.Contains(yslane,"五号位"){
									yslane=strings.Replace(yslane,"大哥位", value.HeroName, -1 )
								}
								if strings.Contains(yslane,"五号位")&&!strings.Contains(yslane,"大哥位"){
									yslane=value.HeroName+yslane
								}
								if strings.Contains(yslane,"大哥位")&&strings.Contains(yslane,"五号位"){
									yslane=strings.Replace(yslane,"大哥位", value.HeroName, -1 )
								}
							}else{
								yslane=yslane+" "+value.HeroName
							}
						}	
					}	
					if  radtopherocount==1{
						yslane=yslane+value.HeroName
						TopSideLineUp.HardLaneMainPositionPlayerId=value.AccountId
					}
				}else{
					if diretopherocount==1{
						ysdirelane=ysdirelane+value.HeroName
						TopSideLineUp.SoftLaneMainPositionPlayerId=value.AccountId
					}
					if diretopherocount==2{
						if value.Position=="四号位"||value.Position=="五号位"{
							ysdirelane=ysdirelane+value.HeroName
						}else{
							ysdirelane=value.HeroName+" "+ysdirelane
							TopSideLineUp.SoftLaneMainPositionPlayerId=value.AccountId
						}
					}
					if diretopherocount==3{
						if value.Position=="五号位"{
							ysdirelane=ysdirelane+" 五号位"+value.HeroName
						}
						if value.Position!="五号位"&&value.Position!="四号位"{
							ysdirelane=value.HeroName+"大哥位 "+ysdirelane
						}
						if value.Position=="四号位"{
							if ysdirelane!=""{
								if strings.Contains(ysdirelane,"大哥位")&&!strings.Contains(ysdirelane,"五号位"){
									ysdirelane=strings.Replace(ysdirelane,"大哥位", value.HeroName, -1 )
								}
								if strings.Contains(ysdirelane,"五号位")&&!strings.Contains(ysdirelane,"大哥位"){
									ysdirelane=value.HeroName+ysdirelane
								}
								if strings.Contains(ysdirelane,"大哥位")&&strings.Contains(ysdirelane,"五号位"){
									ysdirelane=strings.Replace(ysdirelane,"大哥位", value.HeroName, -1 )
								}
							}else{
								ysdirelane=ysdirelane+value.HeroName
							}
						}	
					}
				}	
			}
			if (value.InitLane=="下路"){
				newAgainstHero:=&obj.AgainstHero{}
				newAgainstHero.HeroName=value.HeroName
				newAgainstHero.PlayerId=value.AccountId
				newAgainstHero.PlayerName=value.Player
				newAgainstHero.Pre10LastHit=value.LastHit.Pre10Count
				newAgainstHero.Pre10Denis=value.Denis.Pre10Count
				newAgainstHero.Pre10MinGold=value.Gold.Pre10MinGold
				newAgainstHero.Pre5LastHit=value.LastHit.Pre5Count
				newAgainstHero.Pre5Denis=value.Denis.Pre5Count
				//lvluptimeslice:=playerv.LevelUpTimes
				Loopvaluelvl:
				for vallvluptimekey,vallvluptimevalue:=range value.LevelUpTimes{
					if vallvluptimevalue>600{
						lvlk:=vallvluptimekey-1
						newAgainstHero.Pre10MinLevelUpTime=value.LevelUpTimes[0:lvlk]
						newAgainstHero.Pre10MinLevelUpTimeStr=value.LevelUpTimeStr[0:lvlk]
						newAgainstHero.AbilityUpgrades=value.AbilityData[0:vallvluptimekey]
						break Loopvaluelvl
					}
				}
				//newAgainstHero.Pre10minItem=value.Item
				Loopvalueitem:
				for valitemkey,valitemvalue:=range value.Item{
					if valitemvalue.PurchaseTimeNum>600{
						newAgainstHero.Pre10minItem=value.Item[0:valitemkey]
						break Loopvalueitem
						}
					}
				newAgainstHero.HeroIcon=value.HeroIcon
				newAgainstHero.ConsumableItem=value.Consumable
				//newAgainstHero.AbilityUpgrades=value.AbilityData
				for _,insv:=range value.InventorySnapshot{
					if insv.GameTime==600{
						newAgainstHero.HeroDeathCount=int(insv.Deaths)
					}
				}
				BottomSideLineUpDets.AgainstHeroes=append(BottomSideLineUpDets.AgainstHeroes,newAgainstHero)

				if(value.GameTeam==2){	
					if radbottomherocount==1{
						lslane=lslane+value.HeroName
						BottomSideLineUp.SoftLaneMainPositionPlayerId=value.AccountId
						
					}
					if  radbottomherocount==2{
						if value.Position=="四号位"||value.Position=="五号位"{
							lslane=lslane+value.HeroName
						}else{
							lslane=value.HeroName+" "+lslane
							BottomSideLineUp.SoftLaneMainPositionPlayerId=value.AccountId
						}
					}
					if radbottomherocount==3{
						if value.Position=="五号位"{
							lslane=lslane+" 五号位"+value.HeroName
						}
						if value.Position!="五号位"&&value.Position!="四号位"{
							lslane=value.HeroName+"大哥位 "+lslane
							BottomSideLineUp.SoftLaneMainPositionPlayerId=value.AccountId
						}
						if value.Position=="四号位"{
							if lslane!=""{
								if strings.Contains(lslane,"大哥位")&&!strings.Contains(lslane,"五号位"){
									lslane=strings.Replace(lslane,"大哥位", value.HeroName, -1 )
								}
								if strings.Contains(lslane,"五号位")&&!strings.Contains(lslane,"大哥位"){
									lslane=value.HeroName+lslane		
								}
								if strings.Contains(lslane,"大哥位")&&strings.Contains(lslane,"五号位"){
									lslane=strings.Replace(lslane,"大哥位", value.HeroName, -1 )
								}
							}else{
								lslane=lslane+value.HeroName
							}
						}
					}
				}else{
					if direbottomherocount==1{
						lsdirelane=lsdirelane+value.HeroName
						BottomSideLineUp.HardLaneMainPositionPlayerId=value.AccountId
					}
					if  direbottomherocount==2{
						if value.Position=="四号位"||value.Position=="五号位"{
							lsdirelane=lsdirelane+value.HeroName
						}else{
							lsdirelane=value.HeroName+" "+lsdirelane
							BottomSideLineUp.HardLaneMainPositionPlayerId=value.AccountId
						}
					}
					if direbottomherocount==3{
						if value.Position=="五号位"{
							lsdirelane=lsdirelane+" 五号位"+value.HeroName
						}
						if value.Position!="五号位"&&value.Position!="四号位"{
							lsdirelane=value.HeroName+"大哥位 "+lsdirelane
							BottomSideLineUp.HardLaneMainPositionPlayerId=value.AccountId
						}
						if value.Position=="四号位"{
							if lsdirelane!=""{
								if strings.Contains(lsdirelane,"大哥位")&&!strings.Contains(lsdirelane,"五号位"){
									lsdirelane=strings.Replace(lsdirelane,"大哥位", value.HeroName, -1 )
								}
								if strings.Contains(lsdirelane,"五号位")&&!strings.Contains(lsdirelane,"大哥位"){
									lsdirelane=value.HeroName+lsdirelane		
								}
								if strings.Contains(lsdirelane,"大哥位")&&strings.Contains(lsdirelane,"五号位"){
									lsdirelane=strings.Replace(lsdirelane,"大哥位", value.HeroName, -1 )
								}
							}else{
								lsdirelane=lsdirelane+value.HeroName
							}
						}
					}	
				}		
			}
			if (value.InitLane=="中路"){
				if(value.GameTeam==2){
					if radmidherocount==1{						
						midlane=midlane+value.HeroName
						midstatus1.Hero1PlayerId=value.AccountId
						midstatus1.Hero1Name=value.HeroName
						midstatus1.Hero1Icon=value.HeroIcon
						midstatus1.Hero1Pre5LastHit=value.LastHit.Pre5Count
						midstatus1.Hero1Pre10LastHit=value.LastHit.Pre10Count
						midstatus1.Hero1Pre5Denis=value.Denis.Pre5Count
						midstatus1.Hero1Pre10Denis=value.Denis.Pre10Count
						midstatus1.Hero1ConsumableItem=value.Consumable
						//Pre10MinLevelUpTime:=value.LevelUpTimes
						//Pre10MinKDAInfo:=NewCMsgMatchDetails.KDADets
						lvluptimeslice:=value.LevelUpTimes
						Looplvl:
						for lvlkey1,lvlvalue1:=range lvluptimeslice{
							if lvlvalue1>600{
								lvlk:=lvlkey1-1
								midstatus1.Hero1Pre10MinLevelUpTimeStr=value.LevelUpTimeStr[0:lvlk]
								midstatus1.Hero1AbilityUpgrades=value.AbilityData[0:lvlkey1]
								break Looplvl
							}
						}
						itemslice:=value.Item
						Loopitem:
						for itemkey,itemvalue:=range itemslice{
							if itemvalue.PurchaseTimeNum>600{
								midstatus1.Hero1Pre10minItem=itemslice[0:itemkey]
								break Loopitem
							}
						}
					}
					if radmidherocount==2{
						if value.Position=="四号位"||value.Position=="五号位"{
							midlane=midlane+value.HeroName
						}else{
							midlane=value.HeroName+" "+midlane
						}
					}
				}else{
					if diremidherocount==1{
						diremidlane=diremidlane+value.HeroName
						//diremid=value.HeroName
						midstatus1.Hero2PlayerId=value.AccountId
						midstatus1.Hero2Icon=value.HeroIcon
						midstatus1.Hero2Name=value.HeroName
						midstatus1.Hero2Pre5LastHit=value.LastHit.Pre5Count
						midstatus1.Hero2Pre10LastHit=value.LastHit.Pre10Count
						midstatus1.Hero2Pre5Denis=value.Denis.Pre5Count
						midstatus1.Hero2Pre10Denis=value.Denis.Pre10Count
						midstatus1.Hero2ConsumableItem=value.Consumable

						lvluptimeslice:=value.LevelUpTimes
						Looplvl3:
						for lvlkey1,lvlvalue1:=range lvluptimeslice{
							if lvlvalue1>600{
								lvlk:=lvlkey1-1
								midstatus1.Hero2Pre10MinLevelUpTimeStr=value.LevelUpTimeStr[0:lvlk]
								midstatus1.Hero2AbilityUpgrades=value.AbilityData[0:lvlkey1]
								break Looplvl3
							}
						}
						itemslice:=value.Item
						Loopitem3:
						for itemkey,itemvalue:=range itemslice{
							if itemvalue.PurchaseTimeNum>600{
								midstatus1.Hero2Pre10minItem=itemslice[0:itemkey]
								break Loopitem3
							}
						}
					}

					if diremidherocount==2{
						if value.Position=="四号位"||value.Position=="五号位"{
							diremidlane=diremidlane+value.HeroName
						}else{
							diremidlane=value.HeroName+" "+diremidlane
						}
					}
					
				}	
			}
			
			for _,bpvalue:=range NewCDotaGameInfo.PicksBans{
				hero_name:=HeroIDName[bpvalue.HeroId]
				bpvalue.HeroIcon=HeroNameIcon[hero_name]
				if hero_name==value.HeroName{
					bpvalue.PlayerName=value.Player
					bpvalue.Position=value.Position
				}
			}
			leagueherocount:=&obj.MatchHeroCount{}
			rankherocount:=&obj.MatchHeroCount{}
			newherocount:=&obj.HeroCount{}
			newherocount.RankHeroCount=rankherocount
			newherocount.LeagueHeroCount=leagueherocount
			newherocount.Version=version
			newherocount.Hero=value.HeroName
			newherocount.HeroIcon=value.HeroIcon
			newPlayerCommonHero:=obj.PlayerCommonHero{}
			var ratequery string	
			if value.GameResult=="win"{
				if NewCDotaGameInfo.GameMode==2{
					newherocount.LeagueHeroCount.AllCount=newherocount.LeagueHeroCount.AllCount+1
					newherocount.LeagueHeroCount.WinCount=newherocount.LeagueHeroCount.WinCount+1
					ratequery="match_player_heroes.hero_play_count.$.league.win_count"
				}else{
					newherocount.RankHeroCount.AllCount=newherocount.RankHeroCount.AllCount+1
					newherocount.RankHeroCount.WinCount=newherocount.RankHeroCount.WinCount+1
					ratequery="match_player_heroes.hero_play_count.$.rank.win_count"
				}	
			}else{
				if NewCDotaGameInfo.GameMode==2{
					newherocount.LeagueHeroCount.AllCount=newherocount.LeagueHeroCount.AllCount+1
					newherocount.LeagueHeroCount.LoseCount=newherocount.LeagueHeroCount.LoseCount+1
					ratequery="match_player_heroes.hero_play_count.$.league.lose_count"
				}else{
					newherocount.RankHeroCount.AllCount=newherocount.RankHeroCount.AllCount+1
					newherocount.RankHeroCount.LoseCount=newherocount.RankHeroCount.LoseCount+1
					ratequery="match_player_heroes.hero_play_count.$.rank.lose_count"
				}
			}
			if NewCDotaGameInfo.GameMode==2{
					//遍历天辉选手
					if key<=4{
						newPlayerCommonHero.HeroPlayCount=append(newPlayerCommonHero.HeroPlayCount,newherocount)
						if newplayer.MatchPlayerHeroes!=nil{
							//确定选手有没有更换俱乐部											
							if (newplayer.PlayerState=="正式选手"&&newplayer.TeamName==NewCMsgDOTAMatch.RadiantTeamName)||(newplayer.PlayerState=="替补"){
								if len(newplayer.MatchPlayerHeroes.HeroPlayCount)>0{
									resultcount,err:=playerdb.Find(bson.M{playerquery:value.AccountId,"match_player_heroes.hero_play_count":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}}).Count()
									if err!=nil{
										fmt.Printf("find hero pool err:%v\n",err)
									}
									if resultcount!=0{	
										updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId,"match_player_heroes.hero_play_count":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}},bson.M{"$inc":bson.M{"match_player_heroes.hero_play_count.$.league.all_count":1,ratequery:1}})
										//rateupdateerr:=playerdb.Update(bson.M{playerquery:value.AccountId,"match_player_heroes.hero_play_count":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}},bson.M{"$inc":bson.M{ratequery:1}})
										if updateerr!=nil{
											fmt.Printf("updateerr1:%v\n", updateerr)
										}	
										/*if rateupdateerr!=nil{
											fmt.Printf("rateupdateerr:%v\n", rateupdateerr)
										}*/
									}else{
										updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$push":bson.M{"match_player_heroes.hero_play_count":newherocount}})
										if updateerr!=nil{
											fmt.Printf("value.AccountId:%v updateerr2:%v\n", value.AccountId,updateerr)
										}
									}
								}else{
									updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$push":bson.M{"match_player_heroes.hero_play_count":newherocount}})
									if updateerr!=nil{
										fmt.Printf("value.AccountId:%v updateerr2:%v\n", value.AccountId,updateerr)
									}
								}
							}
						}else{							
							updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$set":bson.M{"match_player_heroes":newPlayerCommonHero}})
							if updateerr!=nil{
								fmt.Printf("value.AccountId:%v 选手数据更新错误：%v\n", value.AccountId,updateerr)
							}
						}
						
					}else{
					//遍历夜魇选手
						newPlayerCommonHero.HeroPlayCount=append(newPlayerCommonHero.HeroPlayCount,newherocount)
						if newplayer.MatchPlayerHeroes!=nil{
							//确定选手有没有更换俱乐部											
							if (newplayer.PlayerState=="正式选手"&&newplayer.TeamName==NewCMsgDOTAMatch.DireTeamName)||(newplayer.PlayerState=="替补"){
								if len(newplayer.MatchPlayerHeroes.HeroPlayCount)>0{
									//fmt.Printf("name：%v id:%v\n", value.Player,value.AccountId)
									resultcount,err:=playerdb.Find(bson.M{playerquery:value.AccountId,"match_player_heroes.hero_play_count":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}}).Count()
									if err!=nil{
										fmt.Printf("find hero pool err:%v\n",err)
									}
									if resultcount!=0{
										updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId,"match_player_heroes.hero_play_count":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}},bson.M{"$inc":bson.M{"match_player_heroes.hero_play_count.$.league.all_count":1,ratequery:1}})
										if updateerr!=nil{
											fmt.Printf("updateerr1:%v\n", updateerr)
										}	
									}else{
										updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$push":bson.M{"match_player_heroes.hero_play_count":newherocount}})
										if updateerr!=nil{
											fmt.Printf("value.AccountId:%v updateerr2:%v\n", value.AccountId,updateerr)
										}
									}
								}else{
									updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$push":bson.M{"match_player_heroes.hero_play_count":newherocount}})
									if updateerr!=nil{
										fmt.Printf("value.AccountId:%v updateerr2:%v\n", value.AccountId,updateerr)
									}
								}
							}
						}else{							
							updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$set":bson.M{"match_player_heroes":newPlayerCommonHero}})
							if updateerr!=nil{
								fmt.Printf("value.AccountId:%v 选手数据更新错误：%v\n", value.AccountId,updateerr)
							}
						}
					}
					//天梯模式
			}else{
				resultcount,err:=playerdb.Find(bson.M{playerquery:value.AccountId,"rank_player_heroes":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}}).Count()
							if err!=nil{
								fmt.Printf("find hero pool err:%v\n",err)
							}
							if resultcount!=0{
								updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId,"rank_player_heroes":bson.M{"$elemMatch":bson.M{"hero":value.HeroName,"version":version}}},bson.M{"$inc":bson.M{"rank_player_heroes.$.rank.all_count":1,ratequery:1}})
								if updateerr!=nil{
									fmt.Printf("updateerr1:%v\n", updateerr)
								}	
							}else{
								updateerr:=playerdb.Update(bson.M{playerquery:value.AccountId},bson.M{"$push":bson.M{"rank_player_heroes":newherocount}})
								if updateerr!=nil{
									fmt.Printf("updateerr2:%v\n", updateerr)
								}
							}
			}
		}
		var NewMidLineUp obj.MidLineUp
		for _,kdavalue:=range NewCMsgMatchDetails.KDADets{
			for herokey,herovalue:=range heroslice{
				herostr:=strconv.Itoa(herovalue)
				attacker:=kdavalue.Attacker
				if herostr==attacker{
					kdavalue.Attacker=playersarr[herokey].HeroName
				}
				target:=kdavalue.DeathTarget
				if herostr==target{
					kdavalue.DeathTarget=playersarr[herokey].HeroName
					if kdavalue.DeathTarget==midstatus1.Hero1Name{
						midstatus1.Hero1DeathCount=midstatus1.Hero1DeathCount+1
					}
					if kdavalue.DeathTarget==midstatus1.Hero2Name{
						midstatus1.Hero2DeathCount=midstatus1.Hero2DeathCount+1
					}
				}
			}

			for ask,asv:=range kdavalue.Assistant{
				if asv=="0"{

					kdavalue.Assistant[ask]=playersarr[0].HeroName
				}
				if asv=="1"{
					kdavalue.Assistant[ask]=playersarr[1].HeroName
				}
				if asv=="2"{
					kdavalue.Assistant[ask]=playersarr[2].HeroName
				}
				if asv=="3"{
					kdavalue.Assistant[ask]=playersarr[3].HeroName
				}
				if asv=="4"{
					kdavalue.Assistant[ask]=playersarr[4].HeroName
				}
				if asv=="5"{
					kdavalue.Assistant[ask]=playersarr[5].HeroName
				}
				if asv=="6"{
					kdavalue.Assistant[ask]=playersarr[6].HeroName
				}
				if asv=="7"{
					kdavalue.Assistant[ask]=playersarr[7].HeroName
				}
				if asv=="8"{
					kdavalue.Assistant[ask]=playersarr[8].HeroName
				}
				if asv=="9"{
					kdavalue.Assistant[ask]=playersarr[9].HeroName
					
				}
			}
		}

		midstatus1.MatchId=NewCDotaGameInfo.MatchId

		if diremidherocount==1&&radmidherocount==1{
			NewCMsgMatchDetails.IsMidSolo=1
			lineupcount,lineuperr:=c4.Find(bson.M{"version":version,"mid_hero1":midstatus1.Hero1Name,"hero1_player_id":midstatus1.Hero1PlayerId,"hero2_player_id":midstatus1.Hero2PlayerId,"mid_heor2":midstatus1.Hero2Name}).Count()
			lineupcount2,lineuperr2:=c4.Find(bson.M{"version":version,"mid_hero1":midstatus1.Hero2Name,"hero1_player_id":midstatus1.Hero2PlayerId,"hero2_player_id":midstatus1.Hero1PlayerId,"mid_heor2":midstatus1.Hero1Name}).Count()
			if lineupcount==0&&lineupcount2==0{
				NewMidLineUp.Version=version
				NewMidLineUp.Hero1PlayerId=midstatus1.Hero1PlayerId
				NewMidLineUp.Hero2PlayerId=midstatus1.Hero2PlayerId
				NewMidLineUp.MidHero1=midstatus1.Hero1Name
				NewMidLineUp.MidHero2=midstatus1.Hero2Name
				NewMidLineUp.AgainstDets=append(NewMidLineUp.AgainstDets,&midstatus1)
				c4.Insert(&NewMidLineUp)
			}else{
				count,err:=c4.Find(bson.M{"version":version,"mid_hero1":midstatus1.Hero1Name,"hero1_player_id":midstatus1.Hero1PlayerId,"hero2_player_id":midstatus1.Hero2PlayerId,"mid_heor2":midstatus1.Hero2Name}).Count()
				if err!=nil{
				fmt.Printf("中单对线搜索错误1：%v\n", lineuperr)
					}
				if count!=0{
					updateerr:=c4.Update(bson.M{"version":version,"mid_hero1":midstatus1.Hero1Name,"hero1_player_id":midstatus1.Hero1PlayerId,"hero2_player_id":midstatus1.Hero2PlayerId,"mid_heor2":midstatus1.Hero2Name},bson.M{"$push":bson.M{"aginst_dets":midstatus1}})
					fmt.Printf("中单对线更新错误1:%v\n", updateerr)
				}else{
					updateerr2:=c4.Update(bson.M{"version":version,"mid_hero1":midstatus1.Hero2Name,"hero1_player_id":midstatus1.Hero2PlayerId,"hero2_player_id":midstatus1.Hero1PlayerId,"mid_heor2":midstatus1.Hero1Name},bson.M{"$push":bson.M{"aginst_dets":midstatus1}})
					fmt.Printf("中单对线更新错误12:%v\n", updateerr2)
				}			
			}
			if lineuperr!=nil{
				fmt.Printf("中单对线搜索错误1：%v\n", lineuperr)
			}
			if lineuperr2!=nil{
				fmt.Printf("中单对线搜索错误1：%v\n", lineuperr2)
			}
		}else{
			NewCMsgMatchDetails.IsMidSolo=0
		}


		yslane=ysdirelane+"vs"+yslane //优势路 vs 劣势路
		yslane=strings.Replace(yslane,"大哥位","", -1 )
		yslane=strings.Replace(yslane,"五号位","",-1)
		lslane=lslane+"vs"+lsdirelane //优势路 vs 劣势路
		lslane=strings.Replace(lslane,"大哥位","", -1 )
		lslane=strings.Replace(lslane,"五号位","",-1)

		midlane=midlane+"vs"+diremidlane//天辉先写
		NewCMsgMatchDetails.AgainstHeroes=yslane+" "+lslane+" "+midlane
		NewCMsgMatchDetails.MidLane=midlane
		NewCMsgMatchDetails.BottomLane=lslane
		NewCMsgMatchDetails.TopLane=yslane

		NewCDotaGameInfo.MidLane=midlane
		NewCDotaGameInfo.BottomLane=lslane
		NewCDotaGameInfo.TopLane=yslane

		NewCDotaGameInfo.RadiantTeamName=NewCMsgDOTAMatch.RadiantTeamName
		NewCDotaGameInfo.DireTeamName=NewCMsgDOTAMatch.DireTeamName
		NewCMsgMatchDetails.Version=version
		NewCMsgDOTAMatch.Version=version

		TopSideLineUp.LineUpHeroes=yslane
		TopSideLineUp.SoftLaneTeam=NewCDotaGameInfo.DireTeamTag
		TopSideLineUp.HardLaneTeam=NewCDotaGameInfo.RadiantTeamTag

		BottomSideLineUp.LineUpHeroes=lslane
		BottomSideLineUp.SoftLaneTeam=NewCDotaGameInfo.RadiantTeamTag
		BottomSideLineUp.HardLaneTeam=NewCDotaGameInfo.DireTeamTag

		NewCMsgMatchDetails.PlayersHeroesDets=playersarr
		if winnerid==2 {
			for _,pvalue:=range NewCDotaGameInfo.PicksBans{
				if pvalue.Team==2{
					NewCDotaGameInfo.GameWinnerBp=append(NewCDotaGameInfo.GameWinnerBp,pvalue)
				}
				if pvalue.Team==3{
					NewCDotaGameInfo.GameLoserBp=append(NewCDotaGameInfo.GameLoserBp,pvalue)
				}
			}
		}else{
			for _,pvalue:=range NewCDotaGameInfo.PicksBans{
				if pvalue.Team==3{
					NewCDotaGameInfo.GameWinnerBp=append(NewCDotaGameInfo.GameWinnerBp,pvalue)
				}
				if pvalue.Team==2{
					NewCDotaGameInfo.GameLoserBp=append(NewCDotaGameInfo.GameLoserBp,pvalue)
				}
			}
		}
		BottomSideLineUpDets.MatchId=NewCMsgDOTAMatch.MatchId
		TopSideLineUpDets.MatchId=NewCMsgDOTAMatch.MatchId
		if playerisexist==""{
			inserterr := c.Insert(&NewCDotaGameInfo)
			inserterr2 := c2.Insert(&NewCMsgDOTAMatch)
			inserterr3 := c3.Insert(&NewCMsgMatchDetails)
			topsidelinecount,sideuperr:=c5.Find(bson.M{"version":version,"soft_lane_team":TopSideLineUp.SoftLaneTeam,"hard_lane_team":TopSideLineUp.HardLaneTeam,"soft_main_position_player_id":TopSideLineUp.SoftLaneMainPositionPlayerId,"hard_main_position_player_id":TopSideLineUp.HardLaneMainPositionPlayerId,"line_up_heroes":TopSideLineUp.LineUpHeroes}).Count()
			if sideuperr!=nil{
				fmt.Printf("sideuperr:%v",sideuperr)
			}
			if topsidelinecount==0{
				TopSideLineUp.LineUpDets=append(TopSideLineUp.LineUpDets,TopSideLineUpDets)
				inserterr5 := c5.Insert(&TopSideLineUp)
				if inserterr5!=nil{
					fmt.Printf("对线数据插入错误：%v\n",inserterr5 )
				}
			}else{
				updateerr:=playerdb.Update(bson.M{"version":version,"soft_lane_team":TopSideLineUp.SoftLaneTeam,"hard_lane_team":TopSideLineUp.HardLaneTeam,"soft_main_position_player_id":TopSideLineUp.SoftLaneMainPositionPlayerId,"hard_main_position_player_id":TopSideLineUp.HardLaneMainPositionPlayerId,"line_up_heroes":TopSideLineUp.LineUpHeroes},bson.M{"$push":bson.M{"side_line_up_dets":TopSideLineUpDets}})
				if updateerr!=nil{
					fmt.Printf("对线数据更新错误:%v\n", updateerr)
				}
			}

			bottomsidelinecount,sideuperr:=c5.Find(bson.M{"version":version,"soft_lane_team":BottomSideLineUp.SoftLaneTeam,"hard_lane_team":BottomSideLineUp.HardLaneTeam,"soft_main_position_player_id":BottomSideLineUp.SoftLaneMainPositionPlayerId,"hard_main_position_player_id":BottomSideLineUp.HardLaneMainPositionPlayerId,"line_up_heroes":BottomSideLineUp.LineUpHeroes}).Count()
			if sideuperr!=nil{
				fmt.Printf("sideuperr:%v",sideuperr)
			}
			if bottomsidelinecount==0{
				BottomSideLineUp.LineUpDets=append(BottomSideLineUp.LineUpDets,BottomSideLineUpDets)
				inserterr5 := c5.Insert(&BottomSideLineUp)
				if inserterr5!=nil{
					fmt.Printf("对线数据插入错误：%v\n",inserterr5 )
				}
			}else{
				updateerr:=playerdb.Update(bson.M{"version":version,"soft_lane_team":BottomSideLineUp.SoftLaneTeam,"hard_lane_team":BottomSideLineUp.HardLaneTeam,"soft_main_position_player_id":BottomSideLineUp.SoftLaneMainPositionPlayerId,"hard_main_position_player_id":BottomSideLineUp.HardLaneMainPositionPlayerId,"line_up_heroes":BottomSideLineUp.LineUpHeroes},bson.M{"$push":bson.M{"side_line_up_dets":BottomSideLineUpDets}})
				if updateerr!=nil{
					fmt.Printf("对线数据更新错误:%v\n", updateerr)
				}
			}


			if inserterr!=nil{
				fmt.Printf("数据库插入错误：%v\n",inserterr)
			}
			if inserterr2!=nil{
				fmt.Printf("CMsgDOTAMatch插入错误：%v\n",inserterr2)
			}
			if inserterr3!=nil{
				fmt.Printf("CMsgMatchDetails插入错误：%v\n",inserterr3)
			}
			resultstr="录像分析完毕"
		}else{
			resultstr=playerisexist
		}
	}else{
		resultstr="录像分析结果已存在"
	}
	return resultstr
}
