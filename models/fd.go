package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Fd struct {
	Id int `orm:"auto" json:"id" description:(自增id)`
	Url string `json:"url"`
	Role int `orm:"default(1)"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func ItemsList() []*Fd {
	o := orm.NewOrm()
	fdcollection := make([]*Fd, 0)
	queryTable := o.QueryTable("fd")
	_, err := queryTable.All(&fdcollection)

	if err != nil {
		fmt.Println(err)
	}

	return fdcollection
}

func ItemDetailSet(fd Fd) *Fd {
	o := orm.NewOrm()
	_, err := o.Insert(&fd)

	if err != nil {
		fmt.Println(err)
	}

	return &fd
}

func init() {
	orm.RegisterModel(new(Fd))
	//orm.RunSyncdb("default", true, true)  // 自动建表
}