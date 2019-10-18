package sql
import(
	"github.com/astaxie/beego/orm"
)
type User struct{
	Id int `orm:"column(id);pk" json:"Id"`
	Name string `json:"Name"` 
	Password string `json:"Password"`
	Privilege int `json:"Privilege"`
}
func init() {
    orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:zhengzhihui123@/dota2_admin?charset=utf8")
    orm.RegisterModel(new(User))
}
