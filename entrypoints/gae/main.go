package gae

import (
	"github.com/peterhoward42/frost/server/routing"
)

// The Google App Engine contract for an app, requires only that request handlers are registered
// to the net/http package as part of code load-time initialisation.
func init() {
	routing.SetUpRouting()
}
