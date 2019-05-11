package controller

import (
	"fmt"
	"net/http"
	"html/template"
	"path" 
) 

type D map[string]interface{}

var tmpl, err = template.ParseGlob("views/*")


func Bishma() {
	var filepath = path.Join("views", "index.html")
	fmt.Println(filepath)
}

func About(w http.ResponseWriter , r *http.Request){
	// var Fpath = path.Join("views","*")
	// var tmpl , er = template.ParseFiles(Fpath)
	
	// if er != nil {
	// 	fmt.Fprint(w,"error")
	// 	return	
	// }

	var Data = D{
		"title" : "bishma",
		"name" : "m.aria bishma fauzassssn",
	}
	err = tmpl.ExecuteTemplate(w, "about", Data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
	// tmpl.ExecuteTemplate(w,"index",Data)
}

func Index(w http.ResponseWriter , r *http.Request)  {
	// var filepath = path.Join("views", "index.html")
	// var tmpl, err = template.ParseFiles(filepath)
	
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var Data =D{
		"title" : "Learning Golang Web",
		"name"  : "mochammad aria bishma fauzan",
	}

	tmpl.ExecuteTemplate(w,"index",Data)

	// err = tmpl.Execute(w, Data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	// fmt.Fprint(w,"Hello World , Welcome To My First Web App Develop With Golang")
}
