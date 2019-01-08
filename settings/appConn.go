package settings

import (
	"database/sql"
	"log"
)

// ConnBusiness representa as regras de conneção com o banco.
type ConnBusiness interface {
	ConnectDB() *conn
}

type conn struct {
	DB *sql.DB
}

// NewConn retorna uma nova instancia de conn.
func NewConn() ConnBusiness {
	return &conn{}
}

func (c *conn) ConnectDB() *conn {
	conn, err := InitDb()
	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão. Erro: %s", err.Error())
	}

	c.DB = conn.DB

	return c
}
