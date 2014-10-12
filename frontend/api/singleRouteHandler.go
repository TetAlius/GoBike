package frontend

import (
	"html/template"
	"log"
	"net/http"
)

func singleRouteHandler(w http.ResponseWriter, r *http.Request) {
	//title := r.URL.Path[len("/route/"):]
	t, err := template.ParseFiles("./frontend/resources/html/routeDetails.html", "./frontend/resources/html/meta.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}
	t.Execute(w, map[string]string{"PageTitle": "GoBike - Route"})
}
