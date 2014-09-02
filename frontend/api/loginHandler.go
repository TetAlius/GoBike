package frontend

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

	coso := r.FormValue("id")

	fmt.Fprintf(w, coso)
}
