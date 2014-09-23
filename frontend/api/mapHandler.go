package frontend

import (
	"net/http"
)

func mapHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./frontend/resources/html/map.html")
}
