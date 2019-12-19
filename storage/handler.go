package function

import (
	"net/http"

	"github.com/Qolzam/openfaas-firebase-storage/storage/handlers"
	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

// Handler function
func Handle(w http.ResponseWriter, r *http.Request) {

	// Server Routing
	if router == nil {

		router = httprouter.New()
		router.POST("/:dir", handlers.UploadeHandle())
		router.GET("/:dir/:name", handlers.GetFileHandle())
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "'X-Requested-With, X-HTTP-Method-Override, Accept, Content-Type,access-control-allow-origin, access-control-allow-headers")

	router.ServeHTTP(w, r)
}
