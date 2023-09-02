package main

import (
	"module/routes"
	"net/http"
)

func main() {
	routes.RouteOn()
	http.ListenAndServe(":4200", nil)
}
