package main

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	"fmt"
)

type User struct {
	Username string
	Password string
}

func FindOne(m map[string]string) {
	// 连接数据库
	Client,_ := mgo.Dial("localhost")
	// 关闭数据库
	defer Client.Clone()
	// 设置连接模式
	Client.SetMode(mgo.Monotonic,true)
	// 切换连接的数据库DB()和集合C()
	users := Client.DB("test").C("users")
	// 查找数据
	// 定义接收
	result := &User{}
	users.Find(m).One(result)
	fmt.Println(result)
}


func main() {
	mm := make(map[string]string)
	mm["username"]="xiaoli"
	FindOne(mm)
}