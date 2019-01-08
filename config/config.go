package config

import (
	"os"
)

// Config monta o retorno padrão
type Config struct {
	PageAccessToken string
	VerifyToken     string
	AppSecret       string
	GraphURL        string
}

// Configuracoes centraliza as opções da api
func Configuracoes() Config {
	r := Config{
		PageAccessToken: os.Getenv("PAGEACCESSTOKEN"),
		VerifyToken:     os.Getenv("VERIFYTOKEN"),
		AppSecret:       os.Getenv("APPSECRET"),
		GraphURL:        os.Getenv("GRAPHURL")}

	return r
}
