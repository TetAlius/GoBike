package frontend

import (
	"fmt"
	"html/template"
	"net/http"

	"log"

	"appengine"

	"github.com/GoBike/backend"
)

func init() {

	//Resources
	cssFileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("resources/css/")))
	http.Handle("/css/", cssFileServer)

	jsFileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("resources/js/")))
	http.Handle("/js/", jsFileServer)

	fontsFileServer := http.StripPrefix("/fonts/", http.FileServer(http.Dir("resources/fonts/")))
	http.Handle("/fonts/", fontsFileServer)

	http.HandleFunc("/", handler)
	http.HandleFunc("/routes", routesHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, backend.TestConnection())
}

func routesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	t, err := template.ParseFiles("resources/html/all-routes.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}

	routes, err := backend.GetAllRoutes(c)
	if err != nil {
		log.Fatal("GetAllRoutes error: ", err)
	}
	t.Execute(w, routes)
}
