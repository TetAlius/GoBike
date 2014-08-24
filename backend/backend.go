package backend

import "appengine"

func main() {
	hostname, err := appengine.ModuleHostname(c, "my-backend", "", "")
	url = "http://" + hostname + "/"
}
