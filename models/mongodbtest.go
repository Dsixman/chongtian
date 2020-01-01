package models
import(
	"replayanaly/models/mongodb"
	//"replayanaly/models/obj"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2" 
	"fmt"
)
//查询测试
type Test1 struct{
	Version string `bson:"Version"`
	Player []k `bson:"player"`
}
type k struct{
	Name string `bson:"name"`
}
func SelectTest (){
	//var NewCMsgMatchDetails obj.CMsgMatchDetails
	//var count=obj.HeroLastHit{}
	//var count=&NewCDotaGameInfo.LastHit
	//var AllNpcAbilities *obj.AllNpcAbilities
	//var Ability1 *obj.Ability
	//fmt.Printf("PlayersHeroesDets:%v\n",PlayersHeroesDets)
	var test []Test1
	session := mongodb.CloneSession()//调用这个获得session
    defer session.Close()  //一定要记得释放
    session.SetMode(mgo.Monotonic, true)
	c := session.DB("dota2_big_data").C("test1")
	result:=c.Find(bson.M{"Version":"7.22h"}).All(&test)//players_heroes_dets &count这里查询出来的结果 ability1只有找到antimage_mana_break字段，其他字段为空
	fmt.Printf("result type:%T\n",result)
	fmt.Printf("result:%v\n",result)
	fmt.Printf("test%v\n", test)
	/*for _,value:=range AllNpcAbilities.Abilities{
		if value.Name=="antimage_mana_break"{
			Ability1=value
		}
	}
	fmt.Printf("Ability1:%v\n",Ability1.DamageType)*/
}