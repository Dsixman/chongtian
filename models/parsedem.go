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
)
func Parse (version string,demurl string,fname string) string{
	var resultstr string
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	c := db.C("CDotaGameInfo")
	c2 := db.C("CMsgDOTAMatch")
	c3 := db.C("CMsgMatchDetails")
	playerdb:=db.C("player_info")//all_player_info
	fnamearr:=strings.Split(fname,".")
	matchid:=fnamearr[0]
	count,_:=c.Find(bson.M{"match_id":matchid}).Count()
	if count==0{
		//cteam_player := db.C("team_player")

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
		//var NewPlayer1,NewPlayer2,NewPlayer3,NewPlayer4,NewPlayer5,NewPlayer6,NewPlayer7,NewPlayer8,NewPlayer9,NewPlayer10 obj.Player
		//var NewPlayer =[]obj.Player{&NewPlayer1,&NewPlayer2,&NewPlayer3,&NewPlayer4,&NewPlayer5,&NewPlayer6,&NewPlayer7,&NewPlayer8,&NewPlayer9,&NewPlayer10}
	/*	var NewHeroCount1,NewHeroCount2,NewHeroCount3,NewHeroCount4,NewHeroCount5,NewHeroCount6,NewHeroCount7,NewHeroCount8,NewHeroCount9,NewHeroCount10 obj.HeroCount
		var NewHeroCountArr=[]*obj.HeroCount{&NewHeroCount1,&NewHeroCount2,&NewHeroCount3,&NewHeroCount4,&NewHeroCount5,&NewHeroCount6,&NewHeroCount7,&NewHeroCount8,&NewHeroCount9,&NewHeroCount10}
		*/
		var HeroIdArr []uint32
		var AbilityArr=make([][]uint32,10)
		var NewCMsgMatchDetails obj.CMsgMatchDetails
		var player1,player2,player3,player4,player5,player6,player7,player8,player9,player10 obj.PlayersHeroInfo
		var playersarr =[]*obj.PlayersHeroInfo{&player1,&player2,&player3,&player4,&player5,&player6,&player7,&player8,&player9,&player10}
		//playersarr[0].AccountId=1
		//fmt.Printf("playersarr[0].AccountId:%v\n", playersarr[0].AccountId)
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

		var heroslice []int
		var yslane string="" //优势路
		var midlane string=""
		var lslane string=""

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

		
		var TalentGold = map[uint32]float64{
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
		}
		var goldtalenttime int
		var talentgold float64
		var talentgold1 float64
		//var istalentgold bool
		var allgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var runegold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var deathgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var buybackgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var killwardgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var lasthitgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var combatgold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var killtowergold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}
		var roshangold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}

		var killcouriergold = []float64{0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0} //击杀信使获得的金钱*/
		var denies = []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var timei = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var gtimei = []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		var towertimei = []int64{0, 0}
		var radiantkilltower uint32
		var direkilltower uint32

		//fmt.Printf("version:%v\n", version)
		p.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error {
			//version:=version
			NewCDotaGameInfo.Version=version
			//fmt.Printf("NewCDotaGameInfo.Version:%v\n",NewCDotaGameInfo.Version)
			NewCDotaGameInfo.MatchId=m.GetGameInfo().GetDota().GetMatchId()
			NewCDotaGameInfo.GameMode=m.GetGameInfo().GetDota().GetGameMode()
			//NewCDotaGameInfo.PlayerInfo=m.GetGameInfo().GetDota().GetPlayerInfo()
			NewCDotaGameInfo.Leagueid=m.GetGameInfo().GetDota().GetLeagueid()
			NewCDotaGameInfo.PicksBans=m.GetGameInfo().GetDota().GetPicksBans()
			//NewCDotaGameInfo.RadiantTeamId=m.GetGameInfo().GetDota().GetRadiantTeamId()
			//NewCDotaGameInfo.DireTeamId=m.GetGameInfo().GetDota().GetDireTeamId()
			NewCDotaGameInfo.RadiantTeamTag=m.GetGameInfo().GetDota().GetRadiantTeamTag()
			NewCDotaGameInfo.DireTeamTag=m.GetGameInfo().GetDota().GetDireTeamTag()		
			NewCDotaGameInfo.EndTime=m.GetGameInfo().GetDota().GetEndTime()
			winnerid := m.GetGameInfo().GetDota().GetGameWinner()
			if winnerid==2 {
				if m.GetGameInfo().GetDota().GetRadiantTeamTag() != "" {
					NewCDotaGameInfo.GameWinner=NewCDotaGameInfo.RadiantTeamTag
					NewCDotaGameInfo.GameLoser=NewCDotaGameInfo.DireTeamTag
					for _,value:=range NewCDotaGameInfo.PicksBans{
						if *value.IsPick==true&&*value.Team==2{
							NewCDotaGameInfo.GameWinnerBp=append(NewCDotaGameInfo.GameWinnerBp,value)
						}
						if *value.IsPick==true&&*value.Team==3{
							NewCDotaGameInfo.GameLoserBp=append(NewCDotaGameInfo.GameLoserBp,value)
						}
					}
				}else{
					NewCDotaGameInfo.GameWinner="天辉"
					NewCDotaGameInfo.GameLoser="夜魇"
					for _,value:=range NewCDotaGameInfo.PicksBans{
						if *value.IsPick==true&&*value.Team==2{
							NewCDotaGameInfo.GameWinnerBp=append(NewCDotaGameInfo.GameWinnerBp,value)
						}
						if *value.IsPick==true&&*value.Team==3{
							NewCDotaGameInfo.GameLoserBp=append(NewCDotaGameInfo.GameLoserBp,value)
						}
					}
				}
			}else{
				if m.GetGameInfo().GetDota().GetRadiantTeamTag() != "" {
					NewCDotaGameInfo.GameWinner=NewCDotaGameInfo.DireTeamTag
					NewCDotaGameInfo.GameLoser=NewCDotaGameInfo.RadiantTeamTag
					for _,value:=range NewCDotaGameInfo.PicksBans{
						if *value.IsPick==true&&*value.Team==3{
							NewCDotaGameInfo.GameWinnerBp=append(NewCDotaGameInfo.GameWinnerBp,value)
						}
						if *value.IsPick==true&&*value.Team==2{
							NewCDotaGameInfo.GameLoserBp=append(NewCDotaGameInfo.GameLoserBp,value)
						}
					}
				}else{
					NewCDotaGameInfo.GameWinner="夜魇"
					NewCDotaGameInfo.GameLoser="天辉"
					for _,value:=range NewCDotaGameInfo.PicksBans{
						if *value.IsPick==true&&*value.Team==3{
							NewCDotaGameInfo.GameWinnerBp=append(NewCDotaGameInfo.GameWinnerBp,value)
						}
						if *value.IsPick==true&&*value.Team==2{
							NewCDotaGameInfo.GameLoserBp=append(NewCDotaGameInfo.GameLoserBp,value)
						}
					}
				}
			}
			return nil
		})
		
		p.Callbacks.OnCMsgDOTAMatch(func(m *dota.CMsgDOTAMatch) error {
			NewCMsgDOTAMatch.Duration = m.GetDuration()
			tm := time.Unix(int64(m.GetStartTime()), 0)
	   		timestr:=fmt.Sprintf("%s",tm.Format("2006-01-02 15:04:05"))
			NewCMsgDOTAMatch.StartTime = timestr
			NewCMsgDOTAMatch.Players = m.GetPlayers()
			for _,value:=range NewCMsgDOTAMatch.Players{
				HeroIdArr=append(HeroIdArr,value.GetHeroId())//在切片没有初始化时不能用HeroIdArr[key]=1形式赋值
			}
			NewCMsgDOTAMatch.TowerStatus = m.GetTowerStatus()
			NewCMsgDOTAMatch.BarracksStatus = m.GetBarracksStatus()
			NewCMsgDOTAMatch.FirstBloodTime = m.GetFirstBloodTime()
			NewCMsgDOTAMatch.AverageSkill = m.GetAverageSkill()
			NewCMsgDOTAMatch.RadiantTeamId = m.GetRadiantTeamId()
			NewCMsgDOTAMatch.DireTeamId = m.GetDireTeamId()
			NewCMsgDOTAMatch.RadiantTeamName = m.GetRadiantTeamName()
			NewCMsgDOTAMatch.DireTeamName = m.GetDireTeamName()
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
						playersarr[playerskey].LevelUpTimes=playersvalue.GetLevelUpTimes()
						playersarr[playerskey].AbilityUpgrades=playersvalue.GetAbilityUpgrades()
						AbilityArr[playerskey]=playersvalue.GetAbilityUpgrades()
						//fmt.Printf("AbilityArr[playerskey]%v\n", AbilityArr[playerskey])
						//AbilityArr[][]=append(AbilityArr[],playersvalue.GetAbilityUpgrades())
						purchaseInfo:=playersvalue.GetInventorySnapshot()
						items:=playersvalue.GetItems()
						for _,itemvalue:=range purchaseInfo{
							if itemvalue.GetGameTime()==0|| itemvalue.GetGameTime()==300 || itemvalue.GetGameTime()==600 || itemvalue.GetGameTime()==900 || itemvalue.GetGameTime()==1200{
								playersarr[playerskey].InventorySnapshot=append(playersarr[playerskey].InventorySnapshot,itemvalue)
							}
						}
						for _,item:=range items{
							if item.GetItemId()!=16&&item.GetItemId()!=17&&item.GetItemId()!=18&&item.GetItemId()!=19&&item.GetItemId()!=20&&item.GetItemId()!=40&&item.GetItemId()!=42&&item.GetItemId()!=43&&item.GetItemId()!=46{
								playersarr[playerskey].Item=append(playersarr[playerskey].Item,item)
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
						playersarr[playerskey2].LevelUpTimes=playersvalue.GetLevelUpTimes()
						playersarr[playerskey2].AbilityUpgrades=playersvalue.GetAbilityUpgrades()
						purchaseInfo:=playersvalue.GetInventorySnapshot()
						items:=playersvalue.GetItems()
						for _,itemvalue:=range purchaseInfo{
							if itemvalue.GetGameTime()==0|| itemvalue.GetGameTime()==300 || itemvalue.GetGameTime()==600 || itemvalue.GetGameTime()==900 || itemvalue.GetGameTime()==1200{
								playersarr[playerskey2].InventorySnapshot=append(playersarr[playerskey2].InventorySnapshot,itemvalue)
							}
						}
						for _,item:=range items{
							if item.GetItemId()!=16&&item.GetItemId()!=17&&item.GetItemId()!=18&&item.GetItemId()!=19&&item.GetItemId()!=20&&item.GetItemId()!=40&&item.GetItemId()!=42&&item.GetItemId()!=43&&item.GetItemId()!=46{
								playersarr[playerskey2].Item=append(playersarr[playerskey2].Item,item)
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
				msgtime = msgtime + 1
				prev5 = float64(starttime) + 300.0 + float64(int32(NewCMsgDOTAMatch.PreGameDuration))
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
				if (msgtype == 4 || msgtype == 10) && (math.Floor(float64(time)) <= math.Ceil(prev10)) {
					target := int(m.GetTargetName())
					x := m.GetLocationX()
					y := m.GetLocationY()
					initlocation := ""
					
					if(x >-1700.0 && x<0&&y>1250.0) ||(x>-1700.0&& x<0&&y<700.0)||(x<1000.0&&y<700.0)||(x<1000.0&&y>1250.0){
						initlocation = "中路"
						/*if target==13{
							fmt.Printf("time:%v target:%v x:%v y:%v initlocation%v\n", time,target,x,y,initlocation)
						}*/
					}
					if x < 0 && y > 0 {
						initlocation = "上路"
					}
					if x > 0 && y < 0 {
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
								}
								if  (hero9initlocation!=playersarr[8].InitLane){
									playersarr[8].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
									lanetimei9=lanetimei1+9
								}	
							}
						case heroslice[9]:
							if hero10initlocation != initlocation {
								hero10initlocation = initlocation
								if (playersarr[9].InitLane=="") {
									playersarr[9].InitLane=initlocation
								}
								if  (hero10initlocation!=playersarr[9].InitLane){
									playersarr[9].ChangeLaneTime=strconv.FormatFloat(math.Ceil(float64(time)), 'f', -1, 32)
									lanetimei10=lanetimei10+1
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
					for key, value := range heroslice {
						//
						istargethero := m.GetIsTargetHero()
						istargettower := m.GetIsTargetBuilding()
						damagesource := int(m.GetDamageSourceName())
						attackerteam := m.GetAttackerTeam()
						target := m.GetTargetName()
						targetteam := m.GetTargetTeam()
						//正补反补
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
								//*dota2lasthit[key] = *dota2lasthit[key] + strconv.FormatInt(herolasthit, 10) + " " + strconv.FormatInt(herodenies, 10) + ","
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
								//*dota2lasthit[key] = *dota2lasthit[key] + strconv.FormatInt(herolasthit, 10) + " " + strconv.FormatInt(herodenies, 10) + ","
								timei[key] = timei[key] + 1
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
								//*dota2lasthit[key] = *dota2lasthit[key] + strconv.FormatInt(herolasthit, 10) + " " + strconv.FormatInt(herodenies, 10) + ","
								timei[key] = timei[key] + 1
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
								//*dota2lasthit[key] = *dota2lasthit[key] + strconv.FormatInt(herolasthit, 10) + " " + strconv.FormatInt(herodenies, 10)
								timei[key] = timei[key] + 1
							}
							if attackerteam == targetteam {
								denies[key] = denies[key] + 1
							}
						}
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
					for key1, value1 := range heroslice {
						if target == value1 {
							if math.Floor(float64(time)) > prev5 && gtimei[key1] < 1 {
								allgold[key1] = allgold[key1] + 90.0*5.0
								allgoldstruct:=&obj.GoldType{}
								allgoldstruct.AllGold=allgold[key1]
								allgoldstruct.LastHitGold=lasthitgold[key1]
								allgoldstruct.CombatGold=combatgold[key1]
								playersarr[key1].Gold=&obj.HeroGold{}
								playersarr[key1].Gold.Pre5MinGold=allgoldstruct
								gtimei[key1] = gtimei[key1] + 1
							}
							if math.Floor(float64(time)) > prev10 && gtimei[key1] < 2 {
								gtimei[key1] = gtimei[key1] + 1
								allgold[key1] = allgold[key1] + 90.0*5.0
								allgoldstruct:=&obj.GoldType{}
								allgoldstruct.AllGold=allgold[key1]
								allgoldstruct.LastHitGold=lasthitgold[key1]
								allgoldstruct.CombatGold=combatgold[key1]
								playersarr[key1].Gold.Pre10MinGold=allgoldstruct
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								killwardgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							if math.Floor(float64(time)) > prev15 && gtimei[key1] < 3 {
								gtimei[key1] = gtimei[key1] + 1
								allgold[key1] = allgold[key1] + 90.0*5.0
								allgoldstruct:=&obj.GoldType{}
								allgoldstruct.AllGold=allgold[key1]
								allgoldstruct.LastHitGold=lasthitgold[key1]
								allgoldstruct.CombatGold=combatgold[key1]
								playersarr[key1].Gold.Pre15MinGold=allgoldstruct
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								killwardgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							if math.Floor(float64(time)) > prev20 && gtimei[key1] < 4 {
								gtimei[key1] = gtimei[key1] + 1
								allgold[key1] = allgold[key1] + 90.0*5.0
								allgoldstruct:=&obj.GoldType{}
								allgoldstruct.AllGold=allgold[key1]
								allgoldstruct.LastHitGold=lasthitgold[key1]
								allgoldstruct.CombatGold=combatgold[key1]
								playersarr[key1].Gold.Pre20MinGold=allgoldstruct
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								killwardgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							if math.Floor(float64(time)) > prev25 && gtimei[key1] < 5 {
								gtimei[key1] = gtimei[key1] + 1
								allgold[key1] = allgold[key1] + 90.0*5.0
								allgoldstruct:=&obj.GoldType{}
								allgoldstruct.AllGold=allgold[key1]
								allgoldstruct.LastHitGold=lasthitgold[key1]
								allgoldstruct.CombatGold=combatgold[key1]
								playersarr[key1].Gold.Pre25MinGold=allgoldstruct
								runegold[key1] = 0
								deathgold[key1] = 0
								buybackgold[key1] = 0
								killwardgold[key1] = 0
								lasthitgold[key1] = 0
								killtowergold[key1] = 0
								roshangold[key1] = 0
								killcouriergold[key1] = 0
							}
							u := float64(int32(m.GetValue()))
							allgold[key1] = allgold[key1] + u
							if m.GetGoldReason() == 0 {
								runegold[key1] = runegold[key1] + u
							}
							if m.GetGoldReason() == 1 {
								deathgold[key1] = deathgold[key1] + u
							}
							if m.GetGoldReason() == 2 {
								buybackgold[key1] = buybackgold[key1] + u
							}
							if m.GetGoldReason() == 6 {
								killwardgold[key1] = killwardgold[key1] + u
							}
							if m.GetGoldReason() == 13 {
								lasthitgold[key1] = lasthitgold[key1] + u
							}
							if m.GetGoldReason() == 11 {
								killtowergold[key1] = killtowergold[key1] + u
							}
							if m.GetGoldReason() == 12 {
								combatgold[key1] = combatgold[key1] + u
							}
							if m.GetGoldReason() == 14 {
								roshangold[key1] = roshangold[key1] + u
							}
							if m.GetGoldReason() == 15 {
								killcouriergold[key1] = killcouriergold[key1] + u
							}
						}

					}
				}
			}
			return nil
		})
		p.Start()
		
		NewCMsgDOTAMatch.MatchId=NewCDotaGameInfo.MatchId
		NewCMsgMatchDetails.MatchId=NewCDotaGameInfo.MatchId
		NewCMsgMatchDetails.RadiantTeamTag=NewCDotaGameInfo.RadiantTeamTag
		NewCMsgMatchDetails.DireTeamTag=NewCDotaGameInfo.DireTeamTag
		var ysdirelane string=""
		var lsdirelane string=""
		var diremidlane string=""
		for key,value:=range playersarr{
			if value.InitLane==""{
				value.InitLane="打野"
			}
			value.Player=*NewCMsgDOTAMatch.Players[key].PlayerName
			heroid:=*NewCMsgDOTAMatch.Players[key].HeroId
			value.HeroName=HeroIDName[heroid]
			ability:=value.AbilityUpgrades
			lvl:=value.LevelUpTimes
			for abilitykey,abilityvalue:=range ability {
				if abilityvalue== 5955 || abilityvalue== 5956 || abilityvalue== 6007 || abilityvalue== 6008 ||abilityvalue== 5957 || abilityvalue== 6026 || abilityvalue== 6301 || abilityvalue == 6318 ||abilityvalue == 6446||abilityvalue==8005 {
					goldtalenttime=int(lvl[abilitykey-1])
					talentgold = TalentGold[abilityvalue]
					goldtalenttime=int(lvl[abilitykey-1])
					talentgold = TalentGold[abilityvalue]
					if goldtalenttime>300 && goldtalenttime<600{
						goldtime1:=600-goldtalenttime
						talentgold1 = float64(goldtime1) * (talentgold / 60.0)
						Pre10MinGold:=playersarr[key].Gold.Pre10MinGold.AllGold
						Pre15MinGold:=playersarr[key].Gold.Pre15MinGold.AllGold 
						playersarr[key].Gold.Pre10MinGold.AllGold=Pre10MinGold+talentgold1*float64(goldtime1)
						playersarr[key].Gold.Pre15MinGold.AllGold=Pre15MinGold+talentgold1*5.0
						if playersarr[key].Gold.Pre20MinGold!=nil{
							Pre20MinGold:=playersarr[key].Gold.Pre20MinGold.AllGold
							playersarr[key].Gold.Pre20MinGold.AllGold=Pre20MinGold+talentgold1*10.0
						}
						if playersarr[key].Gold.Pre25MinGold!=nil{
							Pre25MinGold:=playersarr[key].Gold.Pre25MinGold.AllGold
							playersarr[key].Gold.Pre25MinGold.AllGold=Pre25MinGold+talentgold1*15.0
						}
					}
					if goldtalenttime>600 && goldtalenttime<900{
						goldtime1:=900-goldtalenttime
						talentgold1 = float64(goldtime1) * (talentgold / 60.0)
						Pre15MinGold:=playersarr[key].Gold.Pre15MinGold.AllGold
						playersarr[key].Gold.Pre15MinGold.AllGold=Pre15MinGold+talentgold1*float64(goldtime1)
						if playersarr[key].Gold.Pre20MinGold!=nil{
							Pre20MinGold:=playersarr[key].Gold.Pre20MinGold.AllGold
							playersarr[key].Gold.Pre20MinGold.AllGold=Pre20MinGold+talentgold1*5.0
						}
						if playersarr[key].Gold.Pre25MinGold!=nil{
							Pre25MinGold:=playersarr[key].Gold.Pre25MinGold.AllGold
							playersarr[key].Gold.Pre25MinGold.AllGold=Pre25MinGold+talentgold1*10.0
						}					
					}
					if goldtalenttime>900 && goldtalenttime<1200{
						goldtime1:=1200-goldtalenttime
						talentgold1 = float64(goldtime1) * (talentgold / 60.0)
						Pre20MinGold:=playersarr[key].Gold.Pre20MinGold.AllGold
						playersarr[key].Gold.Pre20MinGold.AllGold=Pre20MinGold+talentgold1*float64(goldtime1)
						if playersarr[key].Gold.Pre25MinGold!=nil{
							Pre25MinGold:=playersarr[key].Gold.Pre25MinGold.AllGold
							playersarr[key].Gold.Pre25MinGold.AllGold=Pre25MinGold+talentgold1*5.0
						}	
					}
					if goldtalenttime>1200 && goldtalenttime<1500{
						goldtime1:=1500-goldtalenttime
						talentgold1 = float64(goldtime1) * (talentgold / 60.0)
						if playersarr[key].Gold.Pre25MinGold!=nil{
							Pre25MinGold:=playersarr[key].Gold.Pre25MinGold.AllGold
							playersarr[key].Gold.Pre25MinGold.AllGold=Pre25MinGold+talentgold1*float64(goldtime1)
						}					
					}
				}
			}
			if (value.InitLane=="上路"){	
				if(value.GameTeam==2){
					yslane=yslane+value.HeroName
				}else{
					ysdirelane=ysdirelane+value.HeroName
				}	
			}
			if (value.InitLane=="下路"){
				
				if(value.GameTeam==2){
					lslane=lslane+value.HeroName
				}else{
					lsdirelane=lsdirelane+value.HeroName
				}		
			}
			if (value.InitLane=="中路"){
				if(value.GameTeam==2){
					midlane=midlane+value.HeroName
				}else{
					diremidlane=diremidlane+value.HeroName
				}	
			}

			count,playerdberr:=playerdb.Find(bson.M{"player_dota2_register_num_id":value.AccountId}).Count()
			if playerdberr!=nil{
				fmt.Printf("playerdberr%v\n", playerdberr)
			}
			if count==0{
				newplayer:=obj.Player{}
				newPlayerCommonHero:=obj.PlayerCommonHero{}
				newHeroCount:=obj.HeroCount{}
				newplayer.PlayerState="正式选手"
				newplayer.MatchPlayerId=value.Player
				newplayer.PlayerDota2NumId=value.AccountId
				if key<4{
					newplayer.TeamName=NewCMsgDOTAMatch.RadiantTeamName
					newplayer.TeamNameTag=NewCDotaGameInfo.RadiantTeamTag
					newPlayerCommonHero.ClubTeam=NewCMsgDOTAMatch.RadiantTeamName
					newPlayerCommonHero.ClubTeamTag=NewCDotaGameInfo.RadiantTeamTag
				}else{
					newplayer.TeamName=NewCMsgDOTAMatch.RadiantTeamName
					newplayer.TeamNameTag=NewCDotaGameInfo.RadiantTeamTag
					newPlayerCommonHero.ClubTeam=NewCMsgDOTAMatch.DireTeamName
					newPlayerCommonHero.ClubTeamTag=NewCDotaGameInfo.DireTeamTag
				}
				newHeroCount.Hero=value.HeroName
				newHeroCount.Count=1
				newHeroCount.Version=version
				newPlayerCommonHero.HeroPlayCount=append(newPlayerCommonHero.HeroPlayCount,&newHeroCount)
				if NewCDotaGameInfo.GameMode==2{
					newplayer.MatchPlayerHeroes=&newPlayerCommonHero
				}else{
					newplayer.RankPlayerHeroes=append(newplayer.RankPlayerHeroes,&newHeroCount)
				}
				playerinserterr := playerdb.Insert(&newplayer)
				if playerinserterr!=nil{
					fmt.Printf("playerinserterr :%v\n", playerinserterr)
				}
				/*Updateerr:=playerdb.Update(bson.M{"Id_":1},bson.M{"$push":bson.M{&newplayer}})
				if Updateerr!=nil{
					fmt.Printf("Updateerr:%v\n", Updateerr)
				}*/
			}else{
				if NewCDotaGameInfo.GameMode==2{
					newplayer:=obj.Player{}
					err:=playerdb.Find(bson.M{"player_dota2_register_num_id":value.AccountId}).Select(bson.M{"player_state":1}).Select(bson.M{"match_player_heroes":1}).One(&newplayer)
					if err!=nil{
						fmt.Printf("select hero_play_count err:%v\n", err)
					}
					heroplaycount:=newplayer.MatchPlayerHeroes.HeroPlayCount
					Loop:
					for _,v:=range heroplaycount{
						if v.Version==version&&v.Hero==value.HeroName{
							v.Count=v.Count+1
							break Loop	
						}
					}
					if key<4{
						if (newplayer.PlayerState=="正式选手"&&newplayer.MatchPlayerHeroes.ClubTeam==NewCMsgDOTAMatch.RadiantTeamName) || (newplayer.PlayerState=="正式选手"&&NewCMsgDOTAMatch.RadiantTeamName==""){
							updateerr:=playerdb.Update(bson.M{"player_dota2_register_num_id":value.AccountId},bson.M{"$set":bson.M{"match_player_heroes.hero_play_count":heroplaycount}})
							if updateerr!=nil{
								fmt.Printf("playerdberr%v\n", updateerr)
							}	
						}
						/*if newplayer.PlayerState=="正式选手"&&newplayer.MatchPlayerHeroes.ClubTeam!=newplayer.TeamName{
							newplayer.OldClubPlayerHeroes=newplayer.MatchPlayerHeroes
							newplayer.MatchPlayerHeroes.ClubTeam=NewCMsgDOTAMatch.RadiantTeamName
							newplayer.MatchPlayerHeroes.ClubTeamTag=NewCDotaGameInfo.RadiantTeamTag
							newHeroCounts:=obj.HeroCount
							newHeroCounts.Hero=value.HeroName
							newHeroCounts.Version=version
							newHeroCounts.Count=1
							newplayer.MatchPlayerHeroes.HeroPlayCount=[]*obj.HeroCount{&newHeroCounts}
							Updateerr:=playerdb.Update(bson.M{"all_player_info.player_dota2_register_num_id":value.AccountId},bson.M{"$push":bson.M{&newplayer.MatchPlayerHeroes}})
						}*/	
					}else{
						if (newplayer.PlayerState=="正式选手"&&newplayer.MatchPlayerHeroes.ClubTeam==NewCMsgDOTAMatch.RadiantTeamName) || (newplayer.PlayerState=="正式选手"&&NewCMsgDOTAMatch.RadiantTeamName==""){
							updateerr:=playerdb.Update(bson.M{"player_dota2_register_num_id":value.AccountId},bson.M{"$set":bson.M{"match_player_heroes.hero_play_count":heroplaycount}})	
							if updateerr!=nil{
								fmt.Printf("playerdberr%v\n", updateerr)
							}
						}
						/*if newplayer.PlayerState=="正式选手"&&newplayer.MatchPlayerHeroes.ClubTeam!=newplayer.TeamName{
							newplayer.OldClubPlayerHeroes=newplayer.MatchPlayerHeroes
							newplayer.MatchPlayerHeroes.ClubTeam=NewCMsgDOTAMatch.DireTeamName
							newplayer.MatchPlayerHeroes.ClubTeamTag=NewCDotaGameInfo.DireTeamTag
							newHeroCounts:=obj.HeroCount
							newHeroCounts.Hero=value.HeroName
							newHeroCounts.Version=version
							newHeroCounts.Count=1
							newplayer.MatchPlayerHeroes.HeroPlayCount=[]*obj.HeroCount{&newHeroCounts}
							Updateerr:=playerdb.Update(bson.M{"all_player_info.player_dota2_register_num_id":value.AccountId},bson.M{"$push":bson.M{&newplayer.MatchPlayerHeroes}})
						}*/
					}
				}else{
					newplayer:=obj.Player{}
					err:=playerdb.Find(bson.M{"player_dota2_register_num_id":value.AccountId}).Select(bson.M{"match_player_heroes.hero_play_count":1}).One(&newplayer)
					if err!=nil{
						fmt.Printf("select hero_play_count err:%v\n", err)
					}
					heroplaycount:=newplayer.RankPlayerHeroes
					loop:
					for _,v:=range heroplaycount{
						if v.Version==version&&v.Hero==value.HeroName{
							v.Count=v.Count+1
							break loop
						}
					}
					updateerr:=playerdb.Update(bson.M{"player_dota2_register_num_id":value.AccountId},bson.M{"$set":bson.M{"rank_player_heroes":heroplaycount}})
					if updateerr!=nil{
						fmt.Printf("playerdberr%v\n", updateerr)
					}
				}
			}

		}
		yslane=yslane+"vs"+ysdirelane
		lslane=lslane+"vs"+lsdirelane
		midlane=midlane+"vs"+diremidlane
		NewCMsgMatchDetails.AgainstHeroes=yslane+" "+lslane+" "+midlane
		NewCMsgMatchDetails.PlayersHeroesDets=playersarr
		inserterr := c.Insert(&NewCDotaGameInfo)
		inserterr2 := c2.Insert(&NewCMsgDOTAMatch)
		inserterr3 := c3.Insert(&NewCMsgMatchDetails)
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
		resultstr="录像分析结果已存在"
	}
	return resultstr
}
