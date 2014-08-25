package frontend

import (
	"fmt"
	"net/http"

	"text/template"
	"time"

	"appengine"
	"appengine/datastore"
)

type TestDatabase struct {
	Name       string
	Number     int
	Date       time.Time
	Content    string
	MoarNumber int64
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/query", handlerQuery)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// [START new_context]
	c := appengine.NewContext(r)
	g := TestDatabase{
		Name:       "Menti",
		Number:     23,
		Date:       time.Now(),
		Content:    "Hola ke ase",
		MoarNumber: 232323,
	}

	key := datastore.NewIncompleteKey(c, "Testing", getId(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Señorin añadido")

}

func handlerQuery(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Testing").Ancestor(getId(c)).Order("-Date").Limit(10)
	testing := make([]TestDatabase, 0, 10)

	if _, err := q.GetAll(c, &testing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := guestbookTemplate.Execute(w, testing); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

var guestbookTemplate = template.Must(template.New("test").Parse(`
<html>
  <head>
    <title>Go Testing Database</title>
  </head>
  <body>
    {{range .}}
      {{with .Name}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`))

func getId(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Testing", "default_testing", 0, nil)
}
