package rules_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/rules"
	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

//Router é a exportação das rotas
func Router() *mux.Router {
	rotas := mux.NewRouter()

	//Rules routes
	rotas.HandleFunc("/rules", rules.GetAllRules).Methods("GET")
	rotas.HandleFunc("/rules", rules.CreateRule).Methods("POST")
	rotas.HandleFunc("/rules", rules.UpdateRule).Methods("PUT")
	rotas.HandleFunc("/rules/{id}", rules.GetRuleByAny).Methods("GET")
	rotas.HandleFunc("/crons", rules.GetAllCrons).Methods("GET")
	rotas.HandleFunc("/crons", rules.CreateCron).Methods("POST")

	return rotas
}

//ExecuteRequest executa a request em test
func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	Router().ServeHTTP(res, req)

	return res
}

//CheckResponseCode verifica o response esperado com o recebido
func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

//ClearTable limpa a tabela passada no parametro
func ClearTable(table string) (int64, error) {

	conn := settings.NewConn().ConnectDB().DB
	sql := fmt.Sprintf("delete from " + table + " where 1 = 1 ")
	res, err := conn.Exec(sql)
	if err != nil {
		return 0, err
	}
	sql = fmt.Sprintf("ALTER TABLE " + table + " AUTO_INCREMENT = 1")
	_, err = conn.Exec(sql)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//AddRules adiciona regras para teste
func AddRules(count int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB
	var id int64
	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO rules ( type_rules, description) VALUES ('action'," + strconv.Itoa(i) + ")")
		res, err := conn.Exec(statement)
		if err != nil {
			return 0, err
		}
		id, _ = res.LastInsertId()
	}

	return id, nil
}
