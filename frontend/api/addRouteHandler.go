package frontend

import (
	"html/template"
	"log"
	"net/http"
)

func addRouteHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./frontend/resources/html/addroute.html", "./frontend/resources/html/meta.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}
	t.Execute(w, map[string]string{"PageTitle": "GoBike - Add Route"})
}
