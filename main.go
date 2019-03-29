package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/employees"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/groups"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/integrations"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/messages"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/rules"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/users"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/utils"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/welcome"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-oci8"
	_ "github.com/qodrorid/godaemon"
	"github.com/robfig/cron"
	"github.com/rs/cors"
)

func main() {
	rotas := mux.NewRouter()
	rotas.HandleFunc("/health", utils.HealthCheck).Methods("GET")

	//Webhook
	rotas.HandleFunc("/webhook/{id}", utils.GetWebhook).Methods("GET")
	rotas.HandleFunc("/webhook/{id}", utils.PostWebhook).Methods("POST")

	//Autenticação
	rotas.HandleFunc("/auth", users.Authenticate).Methods("POST")

	//Enviar Email
	rotas.HandleFunc("/sendEmail/{email}", users.SendEmail).Methods("GET")

	//User routes
	rotas.HandleFunc("/user/{id}", users.GetUserByID).Methods("GET")
	rotas.HandleFunc("/user/search", users.GetUserByAny).Methods("POST")
	rotas.HandleFunc("/user", users.GetAllUser).Methods("GET")
	rotas.HandleFunc("/user", users.CreateUser).Methods("POST")
	rotas.HandleFunc("/user", users.UpdateUser).Methods("PUT")
	rotas.HandleFunc("/user/{id}", users.DeleteUser).Methods("DELETE")

	//Integration routes
	rotas.HandleFunc("/integration/{id}", integrations.GetIntegrationByID).Methods("GET")
	rotas.HandleFunc("/integration/search", integrations.GetIntegrationByAny).Methods("POST")
	rotas.HandleFunc("/integration", integrations.GetAllIntegration).Methods("GET")
	rotas.HandleFunc("/integration", integrations.CreateIntegration).Methods("POST")
	rotas.HandleFunc("/integration", integrations.UpdateIntegration).Methods("PUT")
	rotas.HandleFunc("/integration/{id}", integrations.DeleteIntegration).Methods("DELETE")

	//Welcome routes
	rotas.HandleFunc("/welcome/{id}", welcome.GetWelcomeByID).Methods("GET")
	rotas.HandleFunc("/welcome", welcome.GetAllWelcome).Methods("GET")
	rotas.HandleFunc("/welcome", welcome.CreateWelcome).Methods("POST")
	rotas.HandleFunc("/welcome", welcome.UpdateWelcome).Methods("PUT")
	rotas.HandleFunc("/welcome/{id}", welcome.DeleteWelcome).Methods("DELETE")

	//Employee routes
	rotas.HandleFunc("/employee/{id}", employees.GetEmployeeByID).Methods("GET")
	rotas.HandleFunc("/employee/search", employees.GetEmployeeByAny).Methods("POST")
	rotas.HandleFunc("/employee", employees.GetAllEmployee).Methods("GET")
	rotas.HandleFunc("/employee", employees.CreateEmployee).Methods("POST")
	rotas.HandleFunc("/employee", employees.UpdateEmployee).Methods("PUT")
	rotas.HandleFunc("/employee/{id}", employees.DeleteEmployee).Methods("DELETE")
	rotas.HandleFunc("/employee/welcome/{bool}", employees.GetEmployeeByWelcome).Methods("GET")

	//Message routes
	rotas.HandleFunc("/sendMessage", messages.SendMessage).Methods("POST")
	rotas.HandleFunc("/sendGroupMessage", messages.SendGroupMessage).Methods("POST")
	rotas.HandleFunc("/message/{id}", messages.GetMessageByID).Methods("GET")
	rotas.HandleFunc("/message/search", messages.GetMessageBatchByAny).Methods("POST")
	rotas.HandleFunc("/message/{user}/lastMessage", messages.GetMessageByUser).Methods("GET")
	rotas.HandleFunc("/message", messages.CreateMessage).Methods("POST")

	//Rules routes
	rotas.HandleFunc("/rules", rules.GetAllRules).Methods("GET")
	rotas.HandleFunc("/rules", rules.CreateRule).Methods("POST")
	rotas.HandleFunc("/rules", rules.UpdateRule).Methods("PUT")
	rotas.HandleFunc("/rules/{id}", rules.GetRuleByAny).Methods("GET")

	//Batch routes
	rotas.HandleFunc("/batch/{id}", messages.GetMessageBatchByID).Methods("GET")
	rotas.HandleFunc("/batch", messages.CreateMessageBatch).Methods("POST")

	//Group routes
	rotas.HandleFunc("/group/{id}", groups.GetGroupByID).Methods("GET")
	rotas.HandleFunc("/group", groups.GetAllGroup).Methods("GET")
	rotas.HandleFunc("/group/{id}", groups.DeleteGroupMembers).Methods("DELETE")

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
	cr.AddFunc("0 0 06 * * *", welcome.SendWelcomeMessage)
	cr.Start()

	http.ListenAndServe(fmt.Sprintf(":%d", Port), c.Handler(rotas))
}
