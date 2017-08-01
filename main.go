/*project wego*/ 
/*main package*/
package main

import	(
	"fmt"
	"net/http"
	"./router/user"
	"./router/web"
)

func main(){
	/*main code*/
	http.HandleFunc("/",web.Index)
	http.HandleFunc("/user",user.Router)
	fmt.Println("服务启动在8000端口")
	http.ListenAndServe(":8000",nil)

}
