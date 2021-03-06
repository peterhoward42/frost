package main

import (
	"fmt"
	"github.com/peterhoward42/frost/server/routing"
	"net/http"
)

func main() {
	routing.SetUpRouting()
	port := ":5000"
	fmt.Printf("Serving on port %v", port)
	http.ListenAndServe(port, nil)
}
