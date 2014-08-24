package backend

import "appengine"

func main() {}
hostname, err := appengine.ModuleHostname(c, "my-backend", "", "")
if err != nil {

}
url = "http://" + hostname + "/"
}
