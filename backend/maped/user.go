package maped

import (
	"time"
)

// User the users
type User struct {
	name        string    // the user name
	since       time.Time // the day when the account was created
	alreadyDone []Route
}
