package frontend

import (
	"appengine"
	"github.com/TetAlius/GoBike/backend"
	"html/template"
	"log"
	"net/http"
)

func routesHandler(w http.ResponseWriter, r *http.Request) {
	// 404 page
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	c := appengine.NewContext(r)

	t, err := template.ParseFiles("./frontend/resources/html/all-routes.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}

	routes, err := backend.GetAllRoutes(c)
	if err != nil {
		log.Fatal("GetAllRoutes error: ", err)
	}
	t.Execute(w, routes)
}
