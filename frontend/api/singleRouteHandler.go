package frontend

import (
	"fmt"
	"net/http"
)

func singleRouteHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/route/"):]
	/*
	   t, err := template.ParseFiles("./frontend/resources/html/route-template.html")
	   if err != nil {
	     log.Fatal("Parse file error: ", err)
	   }
	   t.Execute(w, title)*/
	fmt.Fprintf(w, title)
}
