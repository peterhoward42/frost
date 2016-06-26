package main

import (
	"fmt"
	"github.com/peterhoward42/frost/server/routing"
	"net/http"
	"os"
)

func main() {
	routing.SetUpRouting()

	// We must respect the port allocated by the hosting environment, on AWS Beanstalk
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	port = ":" + port
	
	fmt.Printf("Serving on port %v", port)
	http.ListenAndServe(port, nil)
}
