package frost

import (
	"github.com/peterhoward42/frost/server"
)

func init() {
	server.InitialiseGlobalGuiTemplate()
	server.SetUpRouting()
}
