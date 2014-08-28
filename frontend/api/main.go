package frontend

import (
	"fmt"
	"net/http"

	"github.com/GoBike/backend"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, backend.TestConnection())
}
