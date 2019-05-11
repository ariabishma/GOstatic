package main

import (
	"fmt"
	"net/http"
)
import "./controller"

// func index(w http.ResponseWriter , r *http.Request)  {
// 	fmt.Fprint(w,"Hello World , Welcome To My First Web App Develop With Golang")
// }

// func About(w http.ResponseWriter , r *http.Request){
// 	var Fpath = path.Join("views","about.html")
// 	var tmpl , er = template.ParseFiles(Fpath)

// 	if er != nil {
// 		fmt.Fprint(w,"error")
// 		return	
// 	}
// 	var data = map[string]interface{}{
// 		"title" : "bishma",
// 		"name" : "m.aria bishma fauzan",
// 	}
// 	tmpl.Execute(w,data)
// }

func main()  {
/*
	Routing 
*/

	//static
	http.Handle("/asset/",http.StripPrefix("/asset/",http.FileServer(http.Dir("assets"))))

	//file
	http.HandleFunc("/about",controller.About)
	http.HandleFunc("/",controller.Index)
	fmt.Println("Server Starting on http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}