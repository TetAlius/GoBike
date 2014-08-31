package frontend

import (
	"html/template"
	"net/http"
	"time"

	"log"

	"appengine"
	"appengine/datastore"

	"github.com/GoBike/backend"
	"github.com/GoBike/backend/maped"
)

func init() {

	//Resources
	cssFileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("frontend/resources/css/")))
	http.Handle("/css/", cssFileServer)

	jsFileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("frontend/resources/js/")))
	http.Handle("/js/", jsFileServer)

	fontsFileServer := http.StripPrefix("/fonts/", http.FileServer(http.Dir("frontend/resources/fonts/")))
	http.Handle("/fonts/", fontsFileServer)

	http.HandleFunc("/", handler)
	http.HandleFunc("/routes", routesHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func routesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	t, err := template.ParseFiles("./all-routes.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}

	routes, err := backend.GetAllRoutes(c)
	if err != nil {
		log.Fatal("GetAllRoutes error: ", err)
	}
	t.Execute(w, routes)
}

func insertRoutesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := maped.Route{
		Title:          "Salas - Pola de Allande",
		Description:    "Pues ni puta idea, porque cuando lelgamos al puticlub \"Nenas\" nos volvimos con la rodilla debajo del brazo",
		CreationDate:   time.Now(),
		Distance:       52,
		BeginLoc:       "Salas",
		EndLoc:         "Pola de Allande",
		Difficulty:     "Depende de tu rodilla",
		Road:           true,
		Mountain:       true,
		Path:           false,
		Comments:       []string{"Mola pila", "Habia gastroenteritis", "Rompi la rodilla", "No sabia que los paragüayos hablaban"},
		Author:         "Menti",
		Maps:           "mira como mola __-/^^^^^^^\\____",
		Duration:       time.Now(),
		Slope:          1200,
		Photos:         "nah",
		Score:          "over 9000",
		Signal:         true,
		BeginTransport: true,
	}
	key := datastore.NewIncompleteKey(c, "Routes", routeKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func routeKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	return datastore.NewKey(c, "Routes", "default_route", 0, nil)
}
