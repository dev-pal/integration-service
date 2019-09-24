package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dev-pal/integration-service/pkg/api/rest"
)

func main() {
	router := rest.Handler()

	fmt.Println("The dev-pal integration-service server is on tap now: https://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
