package db

import (
	"labix.org/v2/mgo"
)
//
//type User struct {
//	Username string
//	Password string
//}

Client,_ := mgo.Dial("localhost")


//func main() {
//
//	defer Client.Clone()
//
//	Client.SetMode(mgo.Monotonic,true)
//
//
//	db := Client.DB("test").C("users")
//
//	db.Insert(&User{"xiaoli","123456"})
//
//	result := &User{}
//
//	db.Find(bson.M{"username":"xiaoli"}).One(result)
//
//	fmt.Println(result)
//
//}