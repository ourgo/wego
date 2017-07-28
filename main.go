/*project wego*/ 
/*main package*/
package main

import	(
	"fmt"
	"net/http"
)

// 展示首页
func showIndex(write http.ResponseWriter,req *http.Request){
	// 解析路由参数
	req.ParseForm()

	//输出hello wego
	fmt.Fprintf(write,"hello wego")

}

func main(){
	/*main code*/
	http.HandleFunc("/",showIndex)
	fmt.Println("服务启动在8000端口")
	http.ListenAndServe(":8000",nil)

}
