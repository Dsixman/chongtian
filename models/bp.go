package models
import(
	"replayanaly/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/dotabuff/manta/dota"
	"replayanaly/models/obj"
	"strings"
)

func getMatchBp(matchid uint64) []*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent {
	//var bp2 CDotaGameInfo. 
	var NewCDotaGameInfo obj.CDotaGameInfo
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	err:=c.Find(bson.M{"matchid": matchid}).Select(bson.M{"PicksBans":1}).One(&NewCDotaGameInfo)
	if err!=nil{
		fmt.Printf("select bp has err:%v\n",err)
	}
	bp:=NewCDotaGameInfo.PicksBans
	return bp
}
/*err = mndb.
    C("virtualcategoryprototypes").
    Find(bson.M{
        "children":   nil,
        "deleted_by": nil,
        "$or": []bson.M{
            {"indexCycleNo": bson.M{"$exists": false}},
            {"indexCycleNo": bson.M{"$lt": cycle}},
        },
    }).
    One(vc)*/
func getBp(team string ,iswin string,hero string) *[][]*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent{
	var NewCDotaGameInfo []obj.CDotaGameInfo
	var teambp [][]*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent
	//var wincondition string
	var heroarr=strings.Split(hero, " ")
	var heroidarr []uint32
	for _,value:=range heroarr{
		heroid:=HeroNameID[value]
		heroidarr=append(heroidarr,heroid)
	}
	
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	if iswin=="win"{
		if len(heroarr)==1{
			c.Find(bson.M{"game_winner":team,"game_winner_bp.hero_id":heroidarr[0]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==2{
			c.Find(bson.M{"game_winner":team}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[0]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[1]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==3{
			c.Find(bson.M{"game_winner":team,"game_winner_bp.hero_id":heroidarr[0]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[1]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[2]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==4{
			c.Find(bson.M{"game_winner":team}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[1]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[2]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[3]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==5{
			c.Find(bson.M{"game_winner":team}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[1]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[2]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[3]}).Select(bson.M{"game_winner_bp.hero_id":heroidarr[4]}).All(&NewCDotaGameInfo)
		}	
	}
	if iswin=="lose"{
		if len(heroarr)==1{
			c.Find(bson.M{"game_loser":team,"game_loser_bp.hero_id":heroidarr[0]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==2{
			c.Find(bson.M{"game_loser":team}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[0]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[1]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==3{
			c.Find(bson.M{"game_loser":team}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[0]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[1]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[2]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==4{
			c.Find(bson.M{"game_loser":team}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[0]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[1]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[2]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[3]}).All(&NewCDotaGameInfo)
		}
		if len(heroarr)==5{
			c.Find(bson.M{"game_loser":team}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[0]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[1]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[2]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[3]}).Select(bson.M{"game_loser_bp.hero_id":heroidarr[4]}).All(&NewCDotaGameInfo)
		}	
	}
	for  _,value2:=range NewCDotaGameInfo{
		teambp=append(teambp,value2.PicksBans)
	}
	return &teambp
}


func getTeamBp(team string) *[][]*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent{
	var NewCDotaGameInfo []obj.CDotaGameInfo
	var teambp [][]*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	err:=c.Find(bson.M{"$or": []bson.M{bson.M{"radiant_team_tag":team},bson.M{"dire_team_tag": team}}}).Select(bson.M{"PicksBans":1}).All(&NewCDotaGameInfo)
	if err!=nil{
		fmt.Printf("select bp has err:%v\n",err)
	}
	for _,value:=range NewCDotaGameInfo{
		teambp=append(teambp,value.PicksBans)
	}
	return &teambp
}

func getHeroBp(hero string) *[][]*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent{
	var NewCDotaGameInfo []obj.CDotaGameInfo
	var teambp [][]*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	heroid:=HeroNameID[hero]
	heroname:=HeroIDENName[heroid]
	err:=c.Find(bson.M{"picks_bans.hero_name":heroname}).All(&NewCDotaGameInfo)
	for _,value:=range NewCDotaGameInfo{
		teambp=append(teambp,value.PicksBans)
	}
	if err!=nil{
		fmt.Printf("get hero bp has err:%v\n",err)
	}
	return &teambp
}
//队伍各成员英雄池及各英雄使用次数
func getTeamPlayer(team string) *obj.Player{
	var NewTeamPlayer obj.Player
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("player_info")
	err:=c.Find(bson.M{"match_player_heroes.club_team":team}).All(&NewTeamPlayer)
	if err!=nil{
		fmt.Printf("get hero bp has err:%v\n",err)
	}
	return &NewTeamPlayer	
}

