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
	ConCaduser      string
	ConGpessoas     string
}

// Configuracoes centraliza as opções da api
func Configuracoes() Config {
	r := Config{
		PageAccessToken: os.Getenv("PAGEACCESSTOKEN"), //"Bearer DQVJ0b3V3ZAGpZAUEs5blZAPZA0I0S2MyeVdlLThSYnhoenFFV0V6SkJ0YU9UR01aNWgyUHBONDNkUmRLLXFqQWRsODBHMDdHWjNURE9WNi00SWNWNWNMaHRxUi1TaU5vNGV1U3o0Mk9DZAVpockcxZATV5M20tcXBwZAEdKM3ZAxUXh4cHNZAVl9xdXAxTDdPTGdFSThYWlo3eFNvYzBqODRHS201WlFoWjY0ckt4cjE5VXIwYmxjQ3hiLU5wZAC1yYmVOUjh6dkdMOFVB",
		VerifyToken:     os.Getenv("VERIFYTOKEN"),     //"JARVIS",
		AppSecret:       os.Getenv("APPSECRET"),       //"def44fb9938372b6bade48a93b8f91dc",
		GraphURL:        os.Getenv("GRAPHURL")}        //"https://graph.facebook.com/v2.6/"

	return r
}
