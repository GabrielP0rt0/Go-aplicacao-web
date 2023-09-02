package routes

import (
	"module/controllers"
	"net/http"
)

func RouteOn() {
	http.HandleFunc("/", controllers.Index)
}
