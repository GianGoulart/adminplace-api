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
	rotas.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")
	rotas.HandleFunc("/user", controllers.GetAllUser).Methods("GET")
	rotas.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	rotas.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	rotas.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	//Integration routes
	rotas.HandleFunc("/integration/{id}", controllers.GetIntegrationByID).Methods("GET")
	rotas.HandleFunc("/integration", controllers.GetAllIntegration).Methods("GET")
	rotas.HandleFunc("/integration", controllers.CreateIntegration).Methods("POST")
	rotas.HandleFunc("/integration", controllers.UpdateIntegration).Methods("PUT")
	rotas.HandleFunc("/integration/{id}", controllers.DeleteIntegration).Methods("DELETE")

	//Welcome routes
	rotas.HandleFunc("/welcome/{id}", controllers.GetWelcomeByID).Methods("GET")
	rotas.HandleFunc("/welcome", controllers.GetAllWelcome).Methods("GET")
	rotas.HandleFunc("/welcome", controllers.CreateWelcome).Methods("POST")
	rotas.HandleFunc("/welcome", controllers.UpdateWelcome).Methods("PUT")
	rotas.HandleFunc("/welcome/{id}", controllers.DeleteWelcome).Methods("DELETE")

	//Employee routes
	rotas.HandleFunc("/employee/{id}", controllers.GetEmployeeByID).Methods("GET")
	rotas.HandleFunc("/employee", controllers.GetAllEmployee).Methods("GET")
	rotas.HandleFunc("/employee", controllers.CreateEmployee).Methods("POST")
	rotas.HandleFunc("/employee", controllers.UpdateEmployee).Methods("PUT")
	rotas.HandleFunc("/employee/{id}", controllers.DeleteEmployee).Methods("DELETE")

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
