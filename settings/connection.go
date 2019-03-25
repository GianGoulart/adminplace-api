package settings

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	// importacao do mysql
	_ "github.com/go-sql-driver/mysql"
)

//App o aplicativo
type App struct {
	DB  *sql.DB
	Env string
}

//factory
var (
	db *sql.DB
)

//InitDb represent a factory of database
func InitDb() (*App, error) {
	a := App{}
	a.Env = os.Getenv("ENV")
	connectionString := fmt.Sprintf("%s", a.GetDNS())
	var err error

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(9)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}

	return &App{DB: db}, nil
}

//OpenConnection representa a conexão com o banco de dados
func (a *App) OpenConnection() (*App, error) {
	a.Env = os.Getenv("ENV")
	connectionString := fmt.Sprintf("%s", a.GetDNS())
	var err error

	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("[db/OpenConnection] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
		return nil, err
	}

	a.DB.SetConnMaxLifetime(time.Minute * 5)
	a.DB.SetMaxIdleConns(0)
	a.DB.SetMaxOpenConns(15)

	if err != nil {
		log.Printf("[db/OpenConnection] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
		return nil, err
	}
	err = a.DB.Ping()
	if err != nil {
		log.Printf("[db/OpenConnection] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
		return nil, err
	}

	return &App{DB: db}, nil
}

//GetDNS representa a recuperação do acesso ao banco
func (a *App) GetDNS() string {
	var (
		user     string
		password string
		dbname   string
		host     string
		dbPort   int
	)
	file, err := ioutil.ReadFile("./env.json")
	if err == nil {
		jsonMap := make(map[string]interface{})
		json.Unmarshal(file, &jsonMap)

		env := os.Getenv("ENV")
		if env == "" {
			env = "development"
		}

		database := jsonMap[env].(map[string]interface{})

		user = fmt.Sprintf("%v", database["DBUSER"])
		password = fmt.Sprintf("%v", database["DBPASSWORD"])
		dbname = fmt.Sprintf("%v", database["DBNAME"])
		host = fmt.Sprintf("%v", database["DBHOST"])
		dbPort, _ = strconv.Atoi(fmt.Sprintf("%v", database["DBPORT"]))
	} else {
		user = "root"                    //os.Getenv("DBUSER")
		password = "#Gian2803"           //os.Getenv("DBPASSWORD")
		dbname = "WorkplaceDB"           //os.Getenv("DBNAME")
		host = "localhost"               //os.Getenv("DBHOST")
		dbPort, _ = strconv.Atoi("3306") //os.Getenv("DBPORT"))
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, host, dbPort, dbname)
	return connectionString
}
