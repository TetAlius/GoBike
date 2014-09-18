package frontend

import (
	"net/http"

	"github.com/TetAlius/GoBike/backend"
)

func init() {
	//Resources
	cssFileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("./frontend/resources/css/")))
	http.Handle("/css/", cssFileServer)
	jsFileServer := http.StripPrefix("/js/", http.FileServer(http.Dir("./frontend/resources/js/")))
	http.Handle("/js/", jsFileServer)
	fontsFileServer := http.StripPrefix("/fonts/", http.FileServer(http.Dir("./frontend/resources/fonts/")))
	http.Handle("/fonts/", fontsFileServer)

	//Handlers
	http.HandleFunc("/", routesHandler)
	http.HandleFunc("/google", loginGoogleHandler)
	http.HandleFunc("/route/", singleRouteHandler)
	http.HandleFunc("/insert", backend.InsertRoutesHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerPost", registerPostHandler)
	http.HandleFunc("/activateUser", activateUser)

	http.HandleFunc("/addRoute", addRouteHandler)
}
