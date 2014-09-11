package backend

import (
	"appengine"
	"appengine/datastore"
	"github.com/TetAlius/GoBike/backend/maped"
	"net/http"
	"time"
)

func registerUser(user maped.User) {

}

// AvailableUserToRegister checks if the username or email are being used in a user
func AvailableUserToRegister(context appengine.Context, username string, email string) bool {
	query := datastore.NewQuery("Users").
		Filter("Username =", username).
		Filter("Active =", true)
	var users []maped.User
	_, err := query.GetAll(context, &users)

	if len(users) > 0 || err != nil {
		return false
	}

	query = datastore.NewQuery("Users").
		Filter("Email =", email).
		Filter("Active =", true)
	_, err = query.GetAll(context, &users)
	if len(users) > 0 || err != nil {
		return false
	}
	return true
}

// RegisterUser register a new User in the datastore, if the operation
// fails then ir will return a 'false'
func RegisterUser(context appengine.Context, user maped.User) bool {
	context.Infof("Registering " + user.Username)
	key := datastore.NewIncompleteKey(context, "Users", nil)
	context.Infof("Getting the datastore key")
	_, err := datastore.Put(context, key, &user)
	if err != nil {
		context.Errorf(err.Error())
		return false
	}
	return true
}

// ActivateUser activates a user with the hashLink provided
func ActivateUser(context appengine.Context, hashLink string) bool {
	query := datastore.NewQuery("Users").
		Filter("HashLink =", hashLink)
	var users []maped.User
	keys, err := query.GetAll(context, &users)
	if err != nil {
		context.Errorf("Could not recover the keys: %v", err)
		return false
	}
	users[0].Active = true
	datastore.Put(context, keys[0], &users[0])
	return true
}

//FUENTE: http://www.mschoebel.info/2014/03/09/snippet-golang-webapp-login-logout.html

//Login tries to login and store the session of an user
func Login(context appengine.Context, response http.ResponseWriter, username string, password string) bool {
	if checkCredentials(context, username, password) {
		setSession(username, response)
		return true
	} else {
		return false
	}
}

func checkCredentials(context appengine.Context, name string, password string) bool {
	q := datastore.NewQuery("Users").
		Filter("Username =", name)

	var user []maped.User
	_, err := q.GetAll(context, &user)

	if err != nil {
		return false
	}

	if len(user) == 1 {
		if user[0].Password == password && user[0].Active {
			return true
		}
	}
	return false

}

func setSession(username string, response http.ResponseWriter) {
	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{
		Name:       "session",
		Value:      username,
		Path:       "/",
		Domain:     "http://gobycicle.appspot.com/",
		Expires:    expire,
		RawExpires: expire.Format(time.UnixDate),
		MaxAge:     86400,
		Secure:     true,
		HttpOnly:   true,
	}

	http.SetCookie(response, &cookie)
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
