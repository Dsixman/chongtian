package models
import (
"github.com/mojocn/base64Captcha"
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
 "replayanaly/models/sql"
	//"crypto/md5"
    //"encoding/hex"
    "fmt"
)

func LoginValidate(username string,password string, idkey string, verifyValue string) (string ,*sql.User){
	o := orm.NewOrm()
	o.Using("default")
	var Userinfo sql.User
	var loginstate string
	//User := new(user)
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	if verifyResult {
	
	exist := o.QueryTable("user").Filter("name", username).Filter("password", password).Exist()
		if exist {
			loginstate="true"
			_,err:=o.QueryTable("user").Filter("name", username).All(&Userinfo)
			if err!=nil{
				fmt.Printf("queryerr:%v\n",err)
			}
		}else{
			loginstate="false"
		}
	}else{
		fmt.Printf("verifyerr: %v",verifyResult)
	}
	
	return loginstate,&Userinfo
}
func checkIsLogin(){}


