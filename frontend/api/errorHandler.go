package frontend

import (
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	http.ServeFile(w, r, "./frontend/resources/html/notfound.html")
}
