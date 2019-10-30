package models
import(
	"fmt"
	"log"
	"os"
	"time"
	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
	"math"
	"strconv"
	"sort"
	"gopkg.in/mgo.v2" 
)
func Parse (version,demurl string){
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	c := db.C("CDotaGameInfo")
	c2 := db.C("CMsgDOTAMatch")
	c3 := db.C("CMsgMatchDetails")
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
		version:=version
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
			}else{

				NewCDotaGameInfo.GameWinner="天辉"
				NewCDotaGameInfo.GameLoser="夜魇"
			}
		}else{
			if m.GetGameInfo().GetDota().GetRadiantTeamTag() != "" {
				NewCDotaGameInfo.GameWinner=NewCDotaGameInfo.DireTeamTag
				NewCDotaGameInfo.GameLoser=NewCDotaGameInfo.RadiantTeamTag
			}else{
				NewCDotaGameInfo.GameWinner="夜魇"
				NewCDotaGameInfo.GameLoser="天辉"
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
					for _,itemvalue:=range purchaseInfo{
						if itemvalue.GetGameTime()==0|| itemvalue.GetGameTime()==300 || itemvalue.GetGameTime()==600 || itemvalue.GetGameTime()==900 || itemvalue.GetGameTime()==1200{
							playersarr[playerskey].InventorySnapshot=append(playersarr[playerskey].InventorySnapshot,itemvalue)
						}
					}	
				}
			}
			if (value.GetDotaTeam()==3){
				var dotateam uint32=value.GetDotaTeam()

				players:=value.GetPlayers()
				for playerskey,playersvalue:=range players{	
					playersarr[playerskey+5].AccountId=playersvalue.GetAccountId()
					playersarr[playerskey+5].GameTeam=dotateam
					playersarr[playerskey+5].PlayerSlot=playersvalue.GetPlayerSlot()
					playersarr[playerskey+5].LevelUpTimes=playersvalue.GetLevelUpTimes()
					AbilityArr[playerskey+5]=playersvalue.GetAbilityUpgrades()
					//fmt.Printf("playersarr[playerskey].AbilityUpgrades：%v\n",playersarr[playerskey].AbilityUpgrades)
					purchaseInfo:=playersvalue.GetInventorySnapshot()
					for _,itemvalue:=range purchaseInfo{
						if itemvalue.GetGameTime()==0|| itemvalue.GetGameTime()==300 || itemvalue.GetGameTime()==600 || itemvalue.GetGameTime()==900 || itemvalue.GetGameTime()==1200{
							playersarr[playerskey+5].InventorySnapshot=append(playersarr[playerskey+5].InventorySnapshot,itemvalue)
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
							playersarr[key1].Gold.Pre5MinGold=allgold[key1]
							gtimei[key1] = gtimei[key1] + 1
						}
						if math.Floor(float64(time)) > prev10 && gtimei[key1] < 2 {
							gtimei[key1] = gtimei[key1] + 1
							allgold[key1] = allgold[key1] + 90.0*5.0
							playersarr[key1].Gold.Pre10MinGold=allgold[key1]
							runegold[key1] = 0
							deathgold[key1] = 0
							buybackgold[key1] = 0
							lasthitgold[key1] = 0
							killtowergold[key1] = 0
							roshangold[key1] = 0
							killcouriergold[key1] = 0
						}
						if math.Floor(float64(time)) > prev15 && gtimei[key1] < 3 {
							/*for i := 0; i < 10; i++ {
								abilityorder1 := playersarr[i].AbilityUpgrades
								for i2 := 0; i2 < len(abilityorder1); i2++ {
									value := abilityorder1[i2]
									if value == 5955 || value == 5956 || value == 6007 || value == 6008 || value == 5957 || value == 6026 || value == 6301 || value == 6318 || value == 6446||value==8005 {
										goldtalenttime=int(playersarr[i].LevelUpTimes[i2-1])
										talentgold = TalentGold[value]
										if istalentgold == false {
											goldtime1 := prev15 - float64(goldtalenttime)
											talentgold1 = goldtime1 * (talentgold / 60.0)
											istalentgold = true
										} else {
											talentgold1 = 5.0 * talentgold
										}
									}
								}
							}*/
							gtimei[key1] = gtimei[key1] + 1
							allgold[key1] = allgold[key1] + 90.0*5.0
							playersarr[key1].Gold.Pre15MinGold=allgold[key1]
							runegold[key1] = 0
							deathgold[key1] = 0
							buybackgold[key1] = 0
							lasthitgold[key1] = 0
							killtowergold[key1] = 0
							roshangold[key1] = 0
							killcouriergold[key1] = 0
						}
						if math.Floor(float64(time)) > prev20 && gtimei[key1] < 4 {
							gtimei[key1] = gtimei[key1] + 1
							allgold[key1] = allgold[key1] + 9.0*5.0
							playersarr[key1].Gold.Pre20MinGold=allgold[key1]
							runegold[key1] = 0
							deathgold[key1] = 0
							buybackgold[key1] = 0
							lasthitgold[key1] = 0
							killtowergold[key1] = 0
							roshangold[key1] = 0
							killcouriergold[key1] = 0
						}
						if math.Floor(float64(time)) > prev25 && gtimei[key1] < 5 {
							gtimei[key1] = gtimei[key1] + 1
							allgold[key1] = allgold[key1] + 90.0*5.0
							playersarr[key1].Gold.Pre25MinGold=allgold[key1]
							runegold[key1] = 0
							deathgold[key1] = 0
							buybackgold[key1] = 0
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
			}
		}
		if goldtalenttime>300 && goldtalenttime<600{
			goldtime1:=600-goldtalenttime
			
			talentgold1 = float64(goldtime1) * (talentgold / 60.0)
			playersarr[key].Gold.Pre10MinGold=playersarr[key].Gold.Pre10MinGold+talentgold1*float64(goldtime1)
			playersarr[key].Gold.Pre15MinGold=playersarr[key].Gold.Pre15MinGold+talentgold1*5.0
			playersarr[key].Gold.Pre20MinGold=playersarr[key].Gold.Pre20MinGold+talentgold1*5.0
			playersarr[key].Gold.Pre25MinGold=playersarr[key].Gold.Pre10MinGold+talentgold1*5.0
		}
		if goldtalenttime>600 && goldtalenttime<900{
			goldtime1:=900-goldtalenttime
			
			talentgold1 = float64(goldtime1) * (talentgold / 60.0)
//			playersarr[key].Gold.Pre10MinGold=playersarr[key].Gold.Pre10MinGold+talentgold1
			playersarr[key].Gold.Pre15MinGold=playersarr[key].Gold.Pre15MinGold+talentgold1*float64(goldtime1)
			playersarr[key].Gold.Pre20MinGold=playersarr[key].Gold.Pre20MinGold+talentgold1*5.0
			playersarr[key].Gold.Pre25MinGold=playersarr[key].Gold.Pre10MinGold+talentgold1*5.0
		}
		if goldtalenttime>900 && goldtalenttime<1200{
			goldtime1:=900-goldtalenttime
			
			talentgold1 = float64(goldtime1) * (talentgold / 60.0)
			playersarr[key].Gold.Pre20MinGold=playersarr[key].Gold.Pre20MinGold+talentgold1*float64(goldtime1)
			playersarr[key].Gold.Pre25MinGold=playersarr[key].Gold.Pre10MinGold+talentgold1*5.0
		}
		if goldtalenttime>1200 && goldtalenttime<1500{
			goldtime1:=1500-goldtalenttime
			
			talentgold1 = float64(goldtime1) * (talentgold / 60.0)
			playersarr[key].Gold.Pre25MinGold=playersarr[key].Gold.Pre10MinGold+talentgold1*float64(goldtime1)
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
		
	}
	yslane=yslane+"vs"+ysdirelane
	lslane=lslane+"vs"+lsdirelane
	midlane=midlane+"vs"+diremidlane
	NewCMsgMatchDetails.AgainstHeroes=yslane+" "+lslane+" "+midlane
	NewCMsgMatchDetails.MatchId=NewCDotaGameInfo.MatchId
	NewCMsgMatchDetails.RadiantTeamTag=NewCDotaGameInfo.RadiantTeamTag
	NewCMsgMatchDetails.DireTeamTag=NewCDotaGameInfo.DireTeamTag
	NewCMsgMatchDetails.PlayersHeroesDets=playersarr
	//fmt.Printf("NewCDotaGameInfo:%v\n",NewCDotaGameInfo)

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
}
