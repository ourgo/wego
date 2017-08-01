package user

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Username string
	Password string
}

func FindOne(m map[string]string) (bson.M,map[string]string) {
	// 连接数据库
	Client,_ := mgo.Dial("localhost")
	// 关闭数据库
	defer Client.Close()
	// 设置连接模式
	Client.SetMode(mgo.Monotonic,true)
	// 切换连接的数据库DB()和集合C()
	users := Client.DB("test").C("users")
	// 定义接收
	result := &User{}
	// 查找数据
	err := users.Find(m).One(result)
	// 错误处理
	if err != nil{
		merr := make(map[string]string)
		merr["code"]="0"
		merr["msg"]="没有该数据"
		return nil,merr
	}
	// 定义数据转换结构
	mm := bson.M{}
	// 转换数据
	data,_ := bson.Marshal(result)
	// 转换为map
	bson.Unmarshal(data,mm)
	// 返回数据
	return mm,nil
}


func user() {
	mm := make(map[string]string)
	mm["username"]="xiaoli"
	FindOne(mm)
}