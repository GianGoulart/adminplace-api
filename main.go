package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-oci8"
	_ "github.com/qodrorid/godaemon"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)

func main() {
	rotas := mux.NewRouter()
	rotas.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	//Webhook
	rotas.HandleFunc("/webhook/{id}", controllers.GetWebhook).Methods("GET")
	rotas.HandleFunc("/webhook/{id}", controllers.PostWebhook).Methods("POST")

	//Autenticação
	rotas.HandleFunc("/auth", controllers.Authenticate).Methods("POST")

	//Enviar Email
	rotas.HandleFunc("/sendEmail/{email}", controllers.SendEmail).Methods("GET")

	//User routes
	rotas.HandleFunc("/user/{id}", controllers.GetUserByID).Methods("GET")
	rotas.HandleFunc("/user/search", controllers.GetUserByAny).Methods("POST")
	rotas.HandleFunc("/user", controllers.GetAllUser).Methods("GET")
	rotas.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	rotas.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	rotas.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")

	//Integration routes
	rotas.HandleFunc("/integration/{id}", controllers.GetIntegrationByID).Methods("GET")
	rotas.HandleFunc("/integration/search", controllers.GetIntegrationByAny).Methods("POST")
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
	rotas.HandleFunc("/employee/search", controllers.GetEmployeeByAny).Methods("POST")
	rotas.HandleFunc("/employee", controllers.GetAllEmployee).Methods("GET")
	rotas.HandleFunc("/employee", controllers.CreateEmployee).Methods("POST")
	rotas.HandleFunc("/employee", controllers.UpdateEmployee).Methods("PUT")
	rotas.HandleFunc("/employee/{id}", controllers.DeleteEmployee).Methods("DELETE")
	rotas.HandleFunc("/employee/welcome/{bool}", controllers.GetEmployeeByWelcome).Methods("GET")

	//Message routes
	rotas.HandleFunc("/sendMessage", controllers.SendMessage).Methods("POST")
	rotas.HandleFunc("/sendGroupMessage", controllers.SendGroupMessage).Methods("POST")
	rotas.HandleFunc("/message/{id}", controllers.GetMessageByID).Methods("GET")
	rotas.HandleFunc("/message/search", controllers.GetMessageBatchByAny).Methods("POST")
	rotas.HandleFunc("/message/{user}/lastMessage", controllers.GetMessageByUser).Methods("GET")
	rotas.HandleFunc("/message", controllers.CreateMessage).Methods("POST")
	//rotas.HandleFunc("/message/{id}/receive", controllers.UpdateReceivedMessage).Methods("PUT")
	//rotas.HandleFunc("/message/{id}/read", controllers.UpdateReadedMessage).Methods("PUT")

	//Batch routes
	rotas.HandleFunc("/batch/{id}", controllers.GetMessageBatchByID).Methods("GET")
	rotas.HandleFunc("/batch", controllers.CreateMessageBatch).Methods("POST")

	//Group routes
	rotas.HandleFunc("/group/{id}", controllers.GetGroupByID).Methods("GET")
	rotas.HandleFunc("/group", controllers.GetAllGroup).Methods("GET")
	rotas.HandleFunc("/group/{id}", controllers.DeleteGroupMembers).Methods("DELETE")

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

	cr := cron.New()
	cr.AddFunc("0 0 06 * * *", controllers.SendWelcomeMessage)
	cr.Start()

	http.ListenAndServe(fmt.Sprintf(":%d", Port), c.Handler(rotas))
}
