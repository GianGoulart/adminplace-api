package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"
	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Sender  = "giangoulart1994@gmail.com"
	CharSet = "UTF-8"
)

// Authenticate rota: /user/{id}
func Authenticate(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	var obj = models.Authenticate{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &obj)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	user, err := repository.Authenticate(obj)

	responseRequest(w, user, err)
}

//SendEmail
func SendEmail(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	vars := mux.Vars(r)
	email := vars["email"]
	fmt.Println(email)
	var destinatarios = []string{"giancarlo.rodrigues@luizalabs.com"}

	var Recipient = aws.StringSlice(destinatarios)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIAJO6B5DJ6PSR3RJCA", "tWsnKy0VQ13VapVkIiYF3nP+X5mGzV286eux5cQM", ""),
	},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: Recipient,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String("mensagem"),
				},
				/*Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},*/
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String("assunto"),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return
	}

	fmt.Println(result)
}

// GetUserByID rota: /user/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := repository.GetUserByID(id)
	responseRequest(w, user, err)
}

// GetAllUser rota: /user
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := repository.GetAllUser()
	responseRequest(w, user, err)
}

//GetUserByAny
func GetUserByAny(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.User{})
	i := obj.(*models.User)
	user, err := repository.GetUserByAny(i)
	responseRequest(w, user, err)

}

// CreateUser rota: /user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.User{})
	us := obj.(*models.User)
	user, err := repository.CreateUser(us)
	responseRequest(w, user, err)
}

// UpdateUser rota: /user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.User{})
	us := obj.(*models.User)

	user, err := repository.UpdateUser(us)
	responseRequest(w, user, err)
}

// DeleteUser rota: /user/:id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := repository.DeleteUser(id)
	responseRequest(w, user, err)
}
