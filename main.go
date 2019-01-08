package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/dt_souza/adminplace-api/controllers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-oci8"
	_ "github.com/qodrorid/godaemon"
	"github.com/rs/cors"
)

func main() {
	rotas := mux.NewRouter()
	rotas.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	//User routes
	rotas.HandleFunc("/user/{id}", controllers.SendMessage).Methods("GET")
	rotas.HandleFunc("/user", controllers.SendMessage).Methods("GET")
	rotas.HandleFunc("/user", controllers.SendMessage).Methods("POST")
	rotas.HandleFunc("/user", controllers.SendMessage).Methods("PUT")
	rotas.HandleFunc("/user/{id}", controllers.SendMessage).Methods("DELETE")

	//Integration routes
	rotas.HandleFunc("/integration/{id}", controllers.SendMessage).Methods("GET")
	rotas.HandleFunc("/integration", controllers.SendMessage).Methods("GET")
	rotas.HandleFunc("/integration", controllers.SendMessage).Methods("POST")
	rotas.HandleFunc("/integration", controllers.SendMessage).Methods("PUT")
	rotas.HandleFunc("/integration/{id}", controllers.SendMessage).Methods("DELETE")

	//Message routes
	rotas.HandleFunc("/sendMessage", controllers.SendMessage).Methods("POST")
	Port, _ := strconv.Atoi(os.Getenv("PORT"))
	if Port == 0 {
		Port = 3001
	}

	log.Println("Server running in port:", Port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
	})

	http.ListenAndServe(fmt.Sprintf(":%d", Port), c.Handler(rotas))
}
