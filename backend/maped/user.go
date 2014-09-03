package maped

import (
	"time"

	"appengine/datastore"
)

// User the users
type User struct {
	Username    string    // the user name
	Since       time.Time // the day when the account was created
	AlreadyDone []datastore.Key
	Email       string
	Name        string
	Surname     string
	BirthDate   string
	Password    string
	Active      bool
	HashLink    string
}
