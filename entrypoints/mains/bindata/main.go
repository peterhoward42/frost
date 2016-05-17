package main

import (
	"net/http"
	"github.com/peterhoward42/frost/server/routing"
)

func main() {
	routing.SetUpRouting()
	http.ListenAndServe(":47066", nil)
}
