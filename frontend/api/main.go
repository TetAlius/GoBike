package frontend

import (
	"net/http"
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
	http.HandleFunc("/insert", insertRoutesHandler)
<<<<<<< HEAD
    http.HandleFunc("/uploadPhoto", upload)
=======
	http.HandleFunc("/serve/", serve)
	http.HandleFunc("/uploadPhoto", upload)
>>>>>>> 9c3313ae64a719148b18bc8159c37be6fef22214
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerPost", registerPostHandler)
	http.HandleFunc("/activateUser", activateUser)
	http.HandleFunc("/map", mapHandler)

	http.HandleFunc("/addRoute", addRouteHandler)
}
