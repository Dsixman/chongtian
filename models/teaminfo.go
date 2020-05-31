package models

import(
"fmt"
	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson"
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
)
func SaveTeamInfo(icon string,teamname string,teamtagname string,teamid uint32){
	var teaminfo obj.Team
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	dbteam:=db.C("teaminfo")
	teaminfo.TeamIcon=icon
	teaminfo.TeamName=teamname
	teaminfo.TeamTagName=teamtagname
	teaminfo.TeamId=teamid
	err:=dbteam.Insert(&teaminfo)
	if err!=nil{
		fmt.Printf("战队信息插入错误：%v\n", err)
	}
}

func GetAllTeamInfo() []obj.Team{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    db := session.DB("dota2_big_data")
	dbteam:=db.C("teaminfo")
	var allteaminfo []obj.Team
	dbteam.Find(bson.M{}).All(&allteaminfo)
	return allteaminfo
}

func GetTeamHeroPool(teamName string,version string) []obj.Player{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    db := session.DB("dota2_big_data")
    heropooldb:=db.C("player_info")
    var result []obj.Player
    heropooldb.Find(bson.M{"$or":[]bson.M{bson.M{"team_name":teamName},bson.M{"team_name_tag":teamName}},"player_state":"正式选手","match_player_heroes.hero_play_count.version":version}).Select(bson.M{"match_player_heroes":1,"alternate_player_heroes":1,"old_club_player_heroes":1,"rank_player_heroes":1,"player_register_string_id":1,"player_dota2_register_num_id":1,"position":1}).All(&result)
    for _,value:=range result{
    	for _,v1:=range value.MatchPlayerHeroes.HeroPlayCount{
    		v1.HeroIcon=HeroNameIcon[v1.Hero]
    	}
    	if value.AlternatePlayerHeroes!=nil{
    		for _,v2:=range value.AlternatePlayerHeroes.HeroPlayCount{
    			v2.HeroIcon=HeroNameIcon[v2.Hero]
    		}
    	}
    	if value.OldClubPlayerHeroes!=nil{
    		for _,v3:=range value.AlternatePlayerHeroes.HeroPlayCount{
    			v3.HeroIcon=HeroNameIcon[v3.Hero]
    		}
    	}
    	if len(value.RankPlayerHeroes)>0{
    		for _,v4:=range value.RankPlayerHeroes{
    			v4.HeroIcon=HeroNameIcon[v4.Hero]
    		}
    	}
    }
    return result
}

func GetTeamInfo(teamName string) *obj.Team{
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
    db := session.DB("dota2_big_data")
	dbteam:=db.C("teaminfo")
	var teaminfo obj.Team
	err:=dbteam.Find(bson.M{"$or":[]bson.M{bson.M{"team_name":teamName},bson.M{"team_tag_name":teamName}}}).One(&teaminfo)
	if err!=nil{
		fmt.Printf("战队信息查询出现错误:%v\n",err)
	}
	return &teaminfo
}