package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	res := httptest.NewRecorder()
	Router().ServeHTTP(res, req)

	return res
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func clearTable(table string) (int64, error) {

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

func addRules(count int) (int64, error) {
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
