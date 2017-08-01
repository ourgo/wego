package main

import (
	//"labix.org/v2/mgo/bson"
	"fmt"
	mapToStruct "github.com/goinggo/mapstructure"
	"labix.org/v2/mgo/bson"
)

type user struct {
	Username string
	Password string
}

func main() {

	mm := make(map[string]string)
	mm["username"]="xiaoli"

	u := &user{}
	bm := bson.M{}
	mapToStruct.Decode(mm,u)
	bs1,_ := bson.Marshal(u)
	bson.Unmarshal(bs1,bm)
	fmt.Println(bm)
}