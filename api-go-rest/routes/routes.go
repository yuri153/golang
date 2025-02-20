package routes

import (
	"golang/api-go-rest/controllers"
	"golang/api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
