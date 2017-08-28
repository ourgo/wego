package web

import (
	"net/http"
	"html/template"
	"fmt"
)

func Index(w http.ResponseWriter,r *http.Request) {
	//w.Write([]byte("hello index"))
	t,err := template.ParseFiles("web/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w,nil)
}