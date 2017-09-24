package web

import (
	"net/http"
	"html/template"
	"github.com/astaxie/beego/logs"
)

func Index(w http.ResponseWriter,r *http.Request) {
	tmpl,err := template.ParseFiles("web/index.html")
	if err != nil {
		logs.Info("Parse file failed .")
	}
	tmpl.Execute(w,"wego")
}
