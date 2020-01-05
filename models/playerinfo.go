package models

import(
	"fmt"
	"gopkg.in/mgo.v2" 
	"gopkg.in/mgo.v2/bson"
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
)

func GetPlayerInfo (playerNumId uint32) *obj.PlayersHeroInfo {
	var playerInfo obj.PlayersHeroInfo
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	dbplayer:=db.C("player_info")
	err:=dbplayer.Find(bson.M{"player_dota2_num_id":playerNumId}).One(&playerInfo)
	if err!=nil{
		fmt.Printf("GetPlayerInfo err:%v\n", err)
	}
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
		fmt.Printf("GetPlayerInfo err:%v\n", err)
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

func AlterPlayerInfo(playerNumId uint32,playerStringName string,secondNumId uint32,thirdNumId uint32,forthNumId uint32,position string,playerState string){
	//var playerInfo obj.TeamPlayerInfo
	session := mongodb.CloneSession()//调用这个获得session
	defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	db := session.DB("dota2_big_data")
	dbplayer:=db.C("player_info")
	dbplayer.Update(bson.M{"player_dota2_register_num_id":playerNumId},bson.M{"$set":bson.M{"player_register_string_id":playerStringName,"second_num_id": secondNumId,"third_num_id":thirdNumId,"forth_num_id":forthNumId,"position":position,"player_state":playerState}})	
}