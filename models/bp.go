package models
import(
	"replayanaly/models/mongodb"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
	//"github.com/dotabuff/manta/dota"
	"replayanaly/models/obj"
	//"strings"
)

func GetMatchBp(matchid uint64) []*obj.BPInfo{
	//var bp2 CDotaGameInfo. 
	var NewCDotaGameInfo obj.CDotaGameInfo
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	err:=c.Find(bson.M{"matchid": matchid}).Select(bson.M{"picks_bans":1}).One(&NewCDotaGameInfo)
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

func GetTeamBp(team string,version string) *[]obj.CDotaGameInfo{
	var NewCDotaGameInfo []obj.CDotaGameInfo
	//var teambp [][]*obj.BPInfo
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	err:=c.Find(bson.M{"$or": []bson.M{bson.M{"radiant_team_tag":team},bson.M{"dire_team_tag":team}},"version":version}).All(&NewCDotaGameInfo)
	if err!=nil{
		fmt.Printf("select bp has err:%v\n",err)
	}
	/*for _,value:=range NewCDotaGameInfo{
		teambp=append(teambp,value.PicksBans)
	}*/
	// fmt.Printf("teambp:%v\n",teambp)
	return &NewCDotaGameInfo
}

func GetHeroBp(hero string) *[]obj.CDotaGameInfo{
	var NewCDotaGameInfo []obj.CDotaGameInfo
	session := mongodb.CloneSession()
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	heroid:=HeroNameID[hero]
	heroname:=HeroIDENName[heroid]
	err:=c.Find(bson.M{"picks_bans.hero_name":heroname}).All(&NewCDotaGameInfo)
	if err!=nil{
		fmt.Printf("get hero bp has err:%v\n",err)
	}
	return &NewCDotaGameInfo
}
// //队伍各成员英雄池及各英雄使用次数
// func GetTeamPlayer(team string) *obj.Player{
// 	var NewTeamPlayer obj.Player
// 	session := mongodb.CloneSession()
// 	defer session.Close()  //一定要记得释放
//     session.SetMode(mgo.Monotonic, true)
// 	c := session.DB("dota2_big_data").C("player_info")
// 	err:=c.Find(bson.M{"match_player_heroes.club_team":team}).All(&NewTeamPlayer)
// 	if err!=nil{
// 		fmt.Printf("get hero bp has err:%v\n",err)
// 	}
// 	return &NewTeamPlayer	
// }

