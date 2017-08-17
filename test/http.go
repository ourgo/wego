package main

import (
	"net/http"
	"fmt"
)

// 中间件
func logMiddle(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		fmt.Println("middle log")
		next.ServeHTTP(w,r)
	})
}
// 中间件2
func hook(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("hook log",r.URL.Path)
		next.ServeHTTP(w,r)
	})
}

// 首页
func index(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/"{
		return
	}
	w.Write([]byte("hello"))
	fmt.Println("index log")
}

func main() {

	http.Handle("/api",logMiddle(hook(http.HandlerFunc(index))))
	http.Handle("/",hook(http.HandlerFunc(index)))

	http.ListenAndServe(":8000",nil)
}