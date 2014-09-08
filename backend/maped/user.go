package maped

import (
	"time"

	"appengine/datastore"
)

// User the user model
type User struct {
	Username    string          // The user name
	Since       time.Time       // The day when the account was created
	AlreadyDone []datastore.Key // Already done routes by the user
	Email       string          // The user email
	Name        string          // The name
	Surname     string          // the surname
	BirthDate   string          // Happy birthday!
	Password    string          // It's a secret!
	Active      bool            // Is the account active?
	HashLink    string          //For activating purposes
}
