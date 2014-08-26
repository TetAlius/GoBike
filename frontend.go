package main

import (
	"fmt"
	util "github.com/TetAlius/Gobike/util"
	"html/template"
	"log"
	"net/http"

	//"appengine"
	//"appengine/datastore"
)

func init() {
	http.HandleFunc("/route/", routeHandler)
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/allroutes", allRoutesHandler)

}

func allRoutesHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("resources/html/allroutes.html")

	// if err != nil {
	// log.Fatal("Parse file error:", err)
	// }
	err = util.Parse("routesexample.json")
	if err != nil {
		log.Fatal(err)
	}

	// routes := make([]util.AllRoutes, 0, 10)
	// t.Execute(w, util.AllRoutes)
	// t.Execute(w, routes)
	//fmt.Fprintf(w, util.Routes[0].Name)
	//fmt.Fprintf(w, util.Routes.Routes[0].Name)
	t.Execute(w, util.Routes.Routes)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//	c := appengine.NewContext(r)
	fmt.Fprintf(w, "HELLOOO")
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("resources/html/routetemplate.html")

	if err != nil {
		log.Fatal("Parse file error:", err)
	}
	route := r.URL.Path[7:]

	t.Execute(w, map[string]string{"title": route})

}
