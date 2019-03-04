package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id int `orm:"auto" json:"id" description:(自增id)`
	Uid string
	Name string `orm:"size(7)" json:"name" description:(用户昵称)`
	UserName string `orm:"size(30)" description:(登录帐号)`
	Password string `"description:(用户密码)`
	Role int `orm:"default(1)"`
	Created time.Time `orm:"auto_now_add;type(date)"`
}

func Login(username string, password string) *User {
	o := orm.NewOrm()
	var user User
	password = GetMD5Hash(password)
	err := o.QueryTable("user").Filter("UserName", username).One(&user, "Id", "Name")
	if (err != nil) {
		fmt.Println(err)
	}

	return &user
}

func Signup(user User) *User {
	o := orm.NewOrm()

	uid, uiderr := uuid.NewV4()

	if (uiderr == nil) {
		user.Uid = uid.String()
	}

	user.Password = GetMD5Hash(user.Password)
	fmt.Println(user.Password)

	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println(err)
	}

	return &user
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func init() {
	orm.RegisterModel(new(User))
	//orm.RunSyncdb("default", true, true)  // 自动建表
}