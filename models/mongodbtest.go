package models
import(
	"replayanaly/models/mongodb"
	"replayanaly/models/obj"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2" 
	"fmt"
)
//查询测试
func SelectTest (){
	var NewCDotaGameInfo obj.CDotaGameInfo
	//var count=&NewCDotaGameInfo.LastHit
session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("CDotaGameInfo")
	result:=c.Find(bson.M{"matchid": 2271813303}).Select(bson.M{"last_hit.pre_5_count":1}).One(&NewCDotaGameInfo)
	fmt.Printf("result type:%T\n",result)
	fmt.Printf("result:%v\n",result)
	//fmt.Printf("NewCDotaGameInfo:%v\n",*(&NewCDotaGameInfo.LastHit))
}