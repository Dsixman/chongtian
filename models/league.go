package models 
import (
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"fmt"
)
func SaveLeagueInfo(icon string,leaguename string,leagueid uint32,seriesid uint32){
	var Newleague obj.League
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("league")
	Newleague.LeagueIcon=icon
	Newleague.LeagueName=leaguename
	Newleague.LeagueId=leagueid
	Newleague.SeriesId=seriesid
	inserterr := c.Insert(&Newleague)
	if inserterr!=nil{
		fmt.Printf("联赛信息插入错误%v\n", inserterr)
	}
}

func GetLeagueInfo(leaguename string) *obj.League{
	var Newleague obj.League
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("league")
	err := c.Find(bson.M{"league_name":leaguename}).One(&Newleague)
	if err!=nil{
		fmt.Printf("未找到此联赛数据:%v\n", err)
	}
	return &Newleague
}