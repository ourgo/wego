/*project wego*/ 
/*main package*/
package main

import	(
	"fmt"
	"net/http"
	"./router/user"
	"./router/web"
)

// 定义中间件(参数为http.HandlerFunc 不是 http.HandleFunc)
// 该函数返回一个http.Handler函数
//
func middleWeare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		// do somthings
		fmt.Println("middle logs")
		// 传递给下一个http接收(下一个路由处理函数)
		next.ServeHTTP(w,r)
	})
}

func main(){
	/*main code*/
	http.Handle("/user",middleWeare(http.HandlerFunc(user.Router)))
	http.Handle("/",middleWeare(http.HandlerFunc(web.Index)))
	fmt.Println("服务启动在8000端口")
	http.ListenAndServe(":8000",nil)

}
