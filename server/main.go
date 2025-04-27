package main

import (
	"fmt"
	"log"
	"net/http"

	routes "github.com/tarvt/ticket-booking/server/router"
)

// go mod init ticket-booking
// go run main.go
func main() {
	r := routes.Router()
	fmt.Println("starting the server on port 9000...")
	log.Fatal(http.ListenAndServe(":9000", r))

}
