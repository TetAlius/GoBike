package frontend

import (
	"html/template"
	"log"
	"net/http"

	"appengine"

	"strconv"

	"github.com/TetAlius/GoBike/backend"
	"github.com/TetAlius/GoBike/backend/maped"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./frontend/resources/html/register.html", "./frontend/resources/html/meta.html", "./frontend/resources/html/login.html")
	if err != nil {
		log.Fatal("Parse file error: ", err)
	}
	t.Execute(w, map[string]string{"PageTitle": "GoBike - Register"})
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("Trying to register a new User")
	result := backend.AvailableUserToRegister(c, r.FormValue("Username"), r.FormValue("Email"))
	c.Infof("Available user: " + strconv.FormatBool(result))
	if result {
		user := maped.User{}
		user.Username = r.FormValue("Username")
		user.Email = r.FormValue("Email")
		registerResult := backend.RegisterUser(c, user)
		if registerResult {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/register", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/register", http.StatusFound)
	}
}
