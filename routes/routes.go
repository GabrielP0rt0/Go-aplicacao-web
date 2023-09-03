package routes

import (
	"module/controllers"
	"net/http"
)

func RouteOn() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
