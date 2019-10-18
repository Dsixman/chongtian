package models
import (
"github.com/mojocn/base64Captcha"
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"
 "replayanaly/models/sql"
	"crypto/md5"
    //"encoding/hex"
    "fmt"
)

func LoginValidate(username string,password string, idkey string, verifyValue string) (string ,string,*sql.User){
	o := orm.NewOrm()
	o.Using("default")
	var Userinfo sql.User
	var loginstate string
	var errcode string
	data := []byte(password)
	has := md5.Sum(data)
	Password := fmt.Sprintf("%x", has) 
	//User := new(user)
	verifyResult := base64Captcha.VerifyCaptcha(idkey, verifyValue)
	fmt.Printf("verifyResult:%v\n",verifyResult)
	/*err:=o.QueryTable("user").Filter("name", username).One(&Userinfo)
		if err!=nil{
			fmt.Printf("queryerr:%v\n",err)
		}
		fmt.Printf("查询到的用户信息:%v\n",&Userinfo)*/
	
	if verifyResult {
		exist := o.QueryTable("user").Filter("name", username).Filter("password", Password).Exist()
		if exist {
			loginstate="true"
			err:=o.QueryTable("user").Filter("name", username).One(&Userinfo)
			if err!=nil{
				fmt.Printf("queryerr:%v\n",err)
			}
		}else{	
			errcode="用户不存在"
			loginstate="false"
		}
	}else{
		errcode="验证码错误"
		fmt.Printf("verifyerr: %v\n",verifyResult)
	}
	fmt.Printf("Userinfo:%v\n",&Userinfo)
	return errcode,loginstate,&Userinfo
}

func Test(username,password string) *sql.User {
	o := orm.NewOrm()
	o.Using("default")
	var Userinfo sql.User
	data := []byte(password)
	has := md5.Sum(data)
	Password := fmt.Sprintf("%x", has) 
	fmt.Printf("Password:%v\n",Password)
	exist := o.QueryTable("user").Filter("name", username).Filter("password", Password).Exist()
	fmt.Printf("exist:%v\n",exist)//返回的是布尔型
	err:=o.QueryTable("user").Filter("name", username).One(&Userinfo)
		if err!=nil{
			fmt.Printf("queryerr:%v\n",err)
		}
	fmt.Printf("查询到的用户信息:%v\n",&Userinfo)
	return &Userinfo
}



