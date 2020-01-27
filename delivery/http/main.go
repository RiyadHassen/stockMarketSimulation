package main

import (
	"html/template"
	"net/http"
)
var tmpl = template.Must(template.ParseGlob("../FinalDSP/ui/web/templates/*"))
func index(writer http.ResponseWriter, request *http.Request)  {
	tmpl.ExecuteTemplate(writer,"index.layout",nil)
}

func main()  {

	fs := http.FileServer(http.Dir("../FinalDSP/ui/web/assets"))
	http.Handle("/assets/",http.StripPrefix("/assets/",fs))
	http.HandleFunc("/",index)
	http.ListenAndServe(":8080",nil)
}