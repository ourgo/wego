package user

import (
	"net/http"
	"../../model/user"
	"fmt"
)

func Router(w http.ResponseWriter,r *http.Request) {
	// 向页面输出
	w.Write([]byte("hello wego"))
	// 定义map数据(模拟查询)
	m := make(map[string]string)
	m["username"]="xiaoli"
	// 查询并得到返回数据
	mm,_ := user.FindOne(m)
	// 输出返回数据
	fmt.Println(mm)
}